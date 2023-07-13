package gohelper

import (
	"context"
	"net/url"

	"gitee.com/linakesi/lzc-sdk/lang/go/common"
	"gitee.com/linakesi/lzc-sdk/lang/go/localdevice"
	"gitee.com/linakesi/lzc-sdk/lang/go/sys"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	Ingress          sys.IngressServiceClient
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
}

func (d *DeviceProxy) GetAuthToken() (*AuthToken, error) {
	return d.metaCred.getAuthToken()
}

func (d *DeviceProxy) Close() error { return d.conn.Close() }

func (gw *APIGateway) NewDeviceProxy(apiurl string) (*DeviceProxy, error) {
	parsedUrl, err := url.Parse(apiurl)
	if err != nil {
		return nil, err
	}

	cred := gw.cred
	if parsedUrl.Scheme == "http" {
		cred = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	unauthedConn, err := grpc.Dial(parsedUrl.Host, cred)
	if err != nil {
		return nil, err
	}

	metaCred := newMetadataCredentials(unauthedConn)
	conn, err := grpc.Dial(
		parsedUrl.Host, cred,
		grpc.WithPerRPCCredentials(metaCred),
	)
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

		OSSnapshot: sys.NewOSSnapshotServiceClient(conn),
		PkgManager: sys.NewPackageManagerClient(conn),
		Ingress:    sys.NewIngressServiceClient(conn),
		Btrfs:      sys.NewBtrfsUtilClient(conn),
		DirMonitor: sys.NewDirMonitorClient(conn),

		Box:              common.NewBoxServiceClient(conn),
		Users:            common.NewUserManagerClient(conn),
		Devices:          common.NewEndDeviceServiceClient(conn),
		Permisions:       common.NewPermissionManagerClient(conn),
		PeripheralDevice: common.NewPeripheralDeviceServiceClient(conn),
		FileTransfer:     common.NewFileTransferServiceClient(conn),
		Message:          common.NewMessageServiceClient(conn),
	}, nil
}
