package gohelper

import (
	"strings"

	"gitee.com/linakesi/lzc-apis-protos/lang/go/common/devices"
	"gitee.com/linakesi/lzc-apis-protos/lang/go/common/permissions"
	"gitee.com/linakesi/lzc-apis-protos/lang/go/common/users"
	"gitee.com/linakesi/lzc-apis-protos/lang/go/localdevice/clipboard"
	"gitee.com/linakesi/lzc-apis-protos/lang/go/localdevice/dialog"
	"gitee.com/linakesi/lzc-apis-protos/lang/go/localdevice/network"
	"gitee.com/linakesi/lzc-apis-protos/lang/go/localdevice/photo"
	"google.golang.org/grpc"
)

type APIGateway struct {
	conn *grpc.ClientConn
	cred grpc.DialOption

	Users      users.UserManagerClient
	Devices    devices.DevicesClient
	Permisions permissions.PermissionManagerClient
}

type DeviceProxy struct {
	conn *grpc.ClientConn

	Clipboard    clipboard.ClipboardManagerClient
	Dialog       dialog.DialogManagerClient
	PhotoLibrary photo.PhotoLibraryClient
	Network      network.NetworkManagerClient
}

func (d *DeviceProxy) Close() error { return d.conn.Close() }

func (gw *APIGateway) NewDeviceProxy(deviceapiurl string) (*DeviceProxy, error) {
	apiurl := strings.TrimPrefix(deviceapiurl, "https://")

	conn, err := grpc.Dial(apiurl, gw.cred)
	if err != nil {
		return nil, err
	}

	return &DeviceProxy{
		conn: conn,

		Clipboard:    clipboard.NewClipboardManagerClient(conn),
		Dialog:       dialog.NewDialogManagerClient(conn),
		PhotoLibrary: photo.NewPhotoLibraryClient(conn),
		Network:      network.NewNetworkManagerClient(conn),
	}, nil
}

func NewAPIConn() (*grpc.ClientConn, error) {
	cred, err := DialCred()
	if err != nil {
		return nil, err
	}
	return grpc.Dial("unix:"+APISocketPath, cred)
}

func NewAPIGateway() (*APIGateway, error) {
	cred, err := DialCred()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial("unix:"+APISocketPath, cred)
	if err != nil {
		return nil, err
	}
	return &APIGateway{
		cred: cred,
		conn: conn,

		Users:      users.NewUserManagerClient(conn),
		Devices:    devices.NewDevicesClient(conn),
		Permisions: permissions.NewPermissionManagerClient(conn),
	}, nil
}
