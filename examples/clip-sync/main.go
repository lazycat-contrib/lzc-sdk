package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	lzcapis "gitee.com/linakesi/lzc-sdk/lang/go"
	pbdevices "gitee.com/linakesi/lzc-sdk/lang/go/common"
	pbusers "gitee.com/linakesi/lzc-sdk/lang/go/common"
)

// 在后台定期，维护拓扑状态：UID<-->DeviceID的关系 (有可能变动,比如同一个设备上，切换了用户)
// 1. 列举所有UID
// 2. 列举UID对应的所有DeviceID
// 3. 对每个DeviceID注册监听函数
//
//
// 在执行回调函数时，同步剪贴板到其他设备上
// 4. 收到新内容，
// 5. 查询到对应DeviceID的UID
// 6. 此UID的剪贴板内容与上次是否有变化，无变化则不进行后续处理。
// 7. 获取UID对应的设备列表，依次检查跳过关闭同步剪贴板功能的设备。
// 8. 将剪贴板内容依次同步到剩余设备中

type SyncClipDaemon struct {
	gw *lzcapis.APIGateway

	devices     sync.Map // map[peer.id]*Device
	lastcontent sync.Map //map[uid]content
}

func (d *SyncClipDaemon) LastConentChanged(uid string, content string) bool {
	_v, loaded := d.lastcontent.LoadOrStore(uid, content)
	if !loaded {
		return true
	}
	if _v.(string) != content {
		d.lastcontent.Store(uid, content)
		return true
	}
	return false
}

func (sd *SyncClipDaemon) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := sd.oneTick(ctx)
			if err != nil {
				fmt.Println("TICK ERR:", err)
			}
			time.Sleep(time.Second * 30)
		}
	}
}

func (sd *SyncClipDaemon) oneTick(ctx context.Context) error {
	repl, err := sd.gw.Users.ListUIDs(ctx, &pbusers.ListUIDsRequest{})
	if err != nil {
		return err
	}

	for _, uid := range repl.Uids {
		repl, err := sd.gw.Devices.ListEndDevices(ctx, &pbdevices.ListEndDeviceRequest{Uid: uid})
		if err != nil {
			fmt.Printf("List User %q Devices failed:", err)
			continue
		}

		for _, dinfo := range repl.Devices {
			_dp, loaded := sd.devices.LoadOrStore(dinfo.PeerId, NewDevice(sd.gw, dinfo.PeerId))
			_dp.(*Device).UpdateInfo(uid, dinfo.DeviceApiUrl, dinfo.IsOnline)
			if !loaded {
				fmt.Printf("Found New Device: %s@%s IsOnline:%v\n", uid, dinfo.PeerId, dinfo.IsOnline)
				go _dp.(*Device).RegisterCallback(ctx, sd.OnNewClipContent)
			}
		}
	}
	return nil
}

func (sd *SyncClipDaemon) userDevices(uid string) []*Device {
	var all []*Device
	sd.devices.Range(func(_, _d interface{}) bool {
		d := _d.(*Device)
		if d.uid != uid {
			return true
		}
		all = append(all, d)
		return true
	})
	return all
}

func (sd *SyncClipDaemon) OnNewClipContent(uid string, from string, content string) {
	if !sd.LastConentChanged(uid, content) {
		return
	}
	fmt.Printf("New CLIPBOARD from %q %q %q\n", uid, from, content)
	for _, d := range sd.userDevices(uid) {
		if d.id == from {
			continue
		}
		if !d.isonline {
			fmt.Println("Skip writing clipboard for offline device ", d.id)
			continue
		}
		go func(d *Device) {
			err := d.WriteClipboard(context.Background(), content)
			if err != nil {
				fmt.Println(d.id, "SYNC ERROR:", err, d.id)
			} else {
				fmt.Println("SYNC CLIPBOARD TO ", d.id, "DONE")
			}
		}(d)
	}
}

func main() {
	apigw, err := lzcapis.NewAPIGateway(context.Background())
	if err != nil {
		fmt.Println("无法创建LZCAPI", err)
		return
	}
	sd := &SyncClipDaemon{gw: apigw}
	sd.Start(context.Background())
}
