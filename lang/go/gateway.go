package gohelper

import (
	"context"
	"strings"

	"gitee.com/linakesi/lzc-apis-protos/lang/go/common"
	"gitee.com/linakesi/lzc-apis-protos/lang/go/localdevice"
	"gitee.com/linakesi/lzc-apis-protos/lang/go/sys"

	"google.golang.org/grpc"
)

type APIGateway struct {
	conn *grpc.ClientConn
	cred grpc.DialOption

	Users            common.UserManagerClient
	Devices          common.EndDeviceServiceClient
	Permisions       common.PermissionManagerClient
	PeripheralDevice common.PeripheralDeviceServiceClient
	PkgManager       sys.PackageManagerClient

	OSUpgrader sys.OSUpgradeServiceClient
	OSSnapshot sys.OSSnapshotServiceClient
}

type DeviceProxy struct {
	conn *grpc.ClientConn

	Clipboard    localdevice.ClipboardManagerClient
	Dialog       localdevice.DialogManagerClient
	PhotoLibrary localdevice.PhotoLibraryClient
	Network      localdevice.NetworkManagerClient
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

		Clipboard:    localdevice.NewClipboardManagerClient(conn),
		Dialog:       localdevice.NewDialogManagerClient(conn),
		PhotoLibrary: localdevice.NewPhotoLibraryClient(conn),
		Network:      localdevice.NewNetworkManagerClient(conn),
	}, nil
}

func NewAPIConn(ctx context.Context) (*grpc.ClientConn, error) {
	cred, err := DialCred()
	if err != nil {
		return nil, err
	}
	return grpc.DialContext(ctx, "unix:"+APISocketPath, grpc.WithBlock(), cred)
}

func NewAPIGateway(ctx context.Context) (*APIGateway, error) {
	cred, err := DialCred()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialContext(ctx, "unix:"+APISocketPath, grpc.WithBlock(), cred)
	if err != nil {
		return nil, err
	}

	return &APIGateway{
		cred: cred,
		conn: conn,

		OSUpgrader: sys.NewOSUpgradeServiceClient(conn),
		OSSnapshot: sys.NewOSSnapshotServiceClient(conn),
		PkgManager: sys.NewPackageManagerClient(conn),

		Users:            common.NewUserManagerClient(conn),
		Devices:          common.NewEndDeviceServiceClient(conn),
		Permisions:       common.NewPermissionManagerClient(conn),
		PeripheralDevice: common.NewPeripheralDeviceServiceClient(conn),
	}, nil
}
