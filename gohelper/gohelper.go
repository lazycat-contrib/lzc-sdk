package gohelper

import (
	"strings"

	"gitee.com/linakesi/lzc-apis-protos/devices"
	"gitee.com/linakesi/lzc-apis-protos/localdevice/clipboard"
	"gitee.com/linakesi/lzc-apis-protos/localdevice/dialog"
	"gitee.com/linakesi/lzc-apis-protos/permissions"
	"gitee.com/linakesi/lzc-apis-protos/users"
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

	Clipboard clipboard.ClipboardManagerClient
	Dialog    dialog.DialogManagerClient
}

func (gw *APIGateway) NewDeviceProxy(deviceapiurl string) (*DeviceProxy, error) {
	apiurl := strings.TrimPrefix(deviceapiurl, "https://")

	conn, err := grpc.Dial(apiurl, gw.cred)
	if err != nil {
		return nil, err
	}

	return &DeviceProxy{
		conn: conn,

		Clipboard: clipboard.NewClipboardManagerClient(conn),
		Dialog:    dialog.NewDialogManagerClient(conn),
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
