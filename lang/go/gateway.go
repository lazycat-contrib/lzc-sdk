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

	Box              common.BoxServiceClient
	Users            common.UserManagerClient
	Devices          common.EndDeviceServiceClient
	Permisions       common.PermissionManagerClient
	PeripheralDevice common.PeripheralDeviceServiceClient
	FileTransfer     common.FileTransferServiceClient
	PkgManager       sys.PackageManagerClient
	AccessControler  sys.AccessControlerServiceClient
	Btrfs            sys.BtrfsUtilClient
	DirMonitor       sys.DirMonitorClient
	Message          common.MessageServiceClient

	OSSnapshot sys.OSSnapshotServiceClient
}

type DeviceProxy struct {
	conn     *grpc.ClientConn
	metaCred *metadataCredentials

	authToken *AuthToken

	Config       localdevice.UserConfigClient
	Device       localdevice.DeviceServiceClient
	Clipboard    localdevice.ClipboardManagerClient
	Dialog       localdevice.DialogManagerClient
	PhotoLibrary localdevice.PhotoLibraryClient
	Network      localdevice.NetworkManagerClient
	Permission   localdevice.PermissionManagerClient
	FileHandler  common.FileHandlerClient
}

func (d *DeviceProxy) GetAuthToken(ctx context.Context) (*AuthToken, error) {
	return d.metaCred.getAuthToken(ctx)
}

func (d *DeviceProxy) Close() error {
	if d.conn != nil {
		return d.conn.Close()
	}
	return nil
}

func (gw *APIGateway) NewDeviceProxy(apiurl string) (*DeviceProxy, error) {
	parsedUrl, err := url.Parse(apiurl)
	if err != nil {
		return nil, err
	}

	metaCred := newMetadataCredentials(parsedUrl.Host, gw.cred)
	conn, err := grpc.Dial(parsedUrl.Host, grpc.WithPerRPCCredentials(metaCred), gw.cred)
	if err != nil {
		return nil, err
	}

	return &DeviceProxy{
		conn:     conn,
		metaCred: metaCred,

		Config:       localdevice.NewUserConfigClient(conn),
		Device:       localdevice.NewDeviceServiceClient(conn),
		Clipboard:    localdevice.NewClipboardManagerClient(conn),
		Dialog:       localdevice.NewDialogManagerClient(conn),
		PhotoLibrary: localdevice.NewPhotoLibraryClient(conn),
		Network:      localdevice.NewNetworkManagerClient(conn),
		Permission:   localdevice.NewPermissionManagerClient(conn),
		FileHandler:  common.NewFileHandlerClient(conn),
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

		OSSnapshot:      sys.NewOSSnapshotServiceClient(conn),
		PkgManager:      sys.NewPackageManagerClient(conn),
		AccessControler: sys.NewAccessControlerServiceClient(conn),
		Btrfs:           sys.NewBtrfsUtilClient(conn),
		DirMonitor:      sys.NewDirMonitorClient(conn),

		Box:              common.NewBoxServiceClient(conn),
		Users:            common.NewUserManagerClient(conn),
		Devices:          common.NewEndDeviceServiceClient(conn),
		Permisions:       common.NewPermissionManagerClient(conn),
		PeripheralDevice: common.NewPeripheralDeviceServiceClient(conn),
		FileTransfer:     common.NewFileTransferServiceClient(conn),
		Message:          common.NewMessageServiceClient(conn),
	}, nil
}

func (a *APIGateway) Close() error {
	return a.conn.Close()
}
