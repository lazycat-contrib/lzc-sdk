package gohelper

import (
	"context"
	"net/url"

	"gitee.com/linakesi/lzc-sdk/lang/go/common"
	"gitee.com/linakesi/lzc-sdk/lang/go/localdevice"
	"gitee.com/linakesi/lzc-sdk/lang/go/sys"

	"google.golang.org/grpc"
)

type APIGateway struct {
	conn *grpc.ClientConn
	cred grpc.DialOption

	Users            common.UserManagerClient
	Devices          common.EndDeviceServiceClient
	Permisions       common.PermissionManagerClient
	PeripheralDevice common.PeripheralDeviceServiceClient
	FileTransfer     common.FileTransferServiceClient
	PkgManager       sys.PackageManagerClient
	Ingress          sys.IngressServiceClient

	OSUpgrader sys.OSUpgradeServiceClient
	OSSnapshot sys.OSSnapshotServiceClient
}

type DeviceProxy struct {
	conn *grpc.ClientConn

	Config       localdevice.UserConfigClient
	Device       localdevice.DeviceServiceClient
	Clipboard    localdevice.ClipboardManagerClient
	Dialog       localdevice.DialogManagerClient
	PhotoLibrary localdevice.PhotoLibraryClient
	Network      localdevice.NetworkManagerClient
}

func (d *DeviceProxy) Close() error { return d.conn.Close() }

func (gw *APIGateway) NewDeviceProxy(apiurl string) (*DeviceProxy, error) {
	parsedUrl, err := url.Parse(apiurl)
	if err != nil {
		return nil, err
	}

	var opts []grpc.DialOption
	if parsedUrl.Scheme == "https" {
		opts = append(opts, gw.cred)
	}

	conn, err := grpc.Dial(parsedUrl.Host, opts...)
	if err != nil {
		return nil, err
	}

	return &DeviceProxy{
		conn: conn,

		Config:       localdevice.NewUserConfigClient(conn),
		Device:       localdevice.NewDeviceServiceClient(conn),
		Clipboard:    localdevice.NewClipboardManagerClient(conn),
		Dialog:       localdevice.NewDialogManagerClient(conn),
		PhotoLibrary: localdevice.NewPhotoLibraryClient(conn),
		Network:      localdevice.NewNetworkManagerClient(conn),
	}, nil
}

func NewAPIConn(ctx context.Context) (*grpc.ClientConn, error) {
	cred, err := buildClientCredOption(CAPath, APPKeyPath, APPCertPath)
	if err != nil {
		return nil, err
	}
	return dialConn(ctx, cred)
}

func NewAPIGateway(ctx context.Context) (*APIGateway, error) {
	cred, err := buildClientCredOption(CAPath, APPKeyPath, APPCertPath)
	if err != nil {
		return nil, err
	}
	conn, err := dialConn(ctx, cred)
	if err != nil {
		return nil, err
	}
	return &APIGateway{
		cred: cred,
		conn: conn,

		OSUpgrader: sys.NewOSUpgradeServiceClient(conn),
		OSSnapshot: sys.NewOSSnapshotServiceClient(conn),
		PkgManager: sys.NewPackageManagerClient(conn),
		Ingress:    sys.NewIngressServiceClient(conn),

		Users:            common.NewUserManagerClient(conn),
		Devices:          common.NewEndDeviceServiceClient(conn),
		Permisions:       common.NewPermissionManagerClient(conn),
		PeripheralDevice: common.NewPeripheralDeviceServiceClient(conn),
		FileTransfer:     common.NewFileTransferServiceClient(conn),
	}, nil
}
