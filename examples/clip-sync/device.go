package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	lzcapis "gitee.com/linakesi/lzc-sdk/lang/go"

	pbclipboard "gitee.com/linakesi/lzc-sdk/lang/go/localdevice"
)

type Device struct {
	gw *lzcapis.APIGateway
	id string

	uid      string
	devurl   string
	isonline bool
}

type ClipCallback func(uid string, from string, content string)

func NewDevice(gw *lzcapis.APIGateway, id string) *Device {
	return &Device{gw: gw, id: id}
}

func (d *Device) UpdateInfo(uid string, devurl string, isonline bool) {
	d.uid = uid
	d.isonline = isonline
	d.devurl = devurl
}

func (d *Device) RegisterCallback(ctx context.Context, cb ClipCallback) error {
	work := func() {
		if !d.isonline {
			fmt.Println("Skip watch clipboard for offline device ", d.id)
			return
		}
		dp, err := d.gw.NewDeviceProxy(d.devurl)
		if err != nil {
			fmt.Println("Can't connect with end device", d.devurl, err)
			return
		}
		wc, err := dp.Clipboard.Watch(ctx, &pbclipboard.ReadClipRequest{})
		if err != nil {
			fmt.Println("WATCH ERRROR:", d.devurl, err)
			return
		}
		for {
			select {
			case <-wc.Context().Done():
				fmt.Println("WATCH DONE:", d.id, wc.Context().Err())
				return
			default:
				fmt.Println("BEGIN LISTEN CLIPBOARD FOR", d.id)
				data, err := wc.Recv()
				if err != nil {
					fmt.Println("WATCH RECV ERR:", d.devurl, err)
					return
				}
				go cb(d.uid, d.id, string(data.Content))
			}
		}
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			work()
			time.Sleep(time.Second * 5)
		}
	}
}

func (d *Device) WriteClipboard(ctx context.Context, content string) error {
	if !d.isonline {
		return errors.New("Device is not online")
	}
	dp, err := d.gw.NewDeviceProxy(d.devurl)
	if err != nil {
		return err
	}
	defer dp.Close()
	_, err = dp.Clipboard.Write(ctx, &pbclipboard.WriteClipRequest{Content: []byte(content)})
	return err
}
