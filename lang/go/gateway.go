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
	TvOS             sys.TvOSClient
	Version          sys.VersionInfoServiceClient
}

type DeviceProxy struct {
	conn     *grpc.ClientConn
	metaCred *metadataCredentials

	authToken *AuthToken

	Config        localdevice.UserConfigClient
	Device        localdevice.DeviceServiceClient
	Dialog        localdevice.DialogManagerClient
	PhotoLibrary  localdevice.PhotoLibraryClient
	Network       localdevice.NetworkManagerClient
	Permission    localdevice.PermissionManagerClient
	FileHandler   common.FileHandlerClient
	Rim           localdevice.RimClient
	RemoteControl localdevice.RemoteControlClient
	Contact       localdevice.ContactsManagerClient
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

		Config:        localdevice.NewUserConfigClient(conn),
		Device:        localdevice.NewDeviceServiceClient(conn),
		Dialog:        localdevice.NewDialogManagerClient(conn),
		PhotoLibrary:  localdevice.NewPhotoLibraryClient(conn),
		Network:       localdevice.NewNetworkManagerClient(conn),
		Permission:    localdevice.NewPermissionManagerClient(conn),
		FileHandler:   common.NewFileHandlerClient(conn),
		Rim:           localdevice.NewRimClient(conn),
		RemoteControl: localdevice.NewRemoteControlClient(conn),
		Contact:       localdevice.NewContactsManagerClient(conn),
	}, nil
}

func NewAPIConn(ctx context.Context) (*grpc.ClientConn, error) {
	cred, err := BuildClientCredOption(CAPath, APPKeyPath, APPCertPath)
	if err != nil {
		return nil, err
	}
	return dialConn(ctx, cred)
}

func NewAPIGateway(ctx context.Context) (*APIGateway, error) {
	cred, err := BuildClientCredOption(CAPath, APPKeyPath, APPCertPath)
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

		PkgManager:      sys.NewPackageManagerClient(conn),
		AccessControler: sys.NewAccessControlerServiceClient(conn),
		Btrfs:           sys.NewBtrfsUtilClient(conn),
		DirMonitor:      sys.NewDirMonitorClient(conn),
		TvOS:            sys.NewTvOSClient(conn),

		Box:              common.NewBoxServiceClient(conn),
		Users:            common.NewUserManagerClient(conn),
		Devices:          common.NewEndDeviceServiceClient(conn),
		Permisions:       common.NewPermissionManagerClient(conn),
		PeripheralDevice: common.NewPeripheralDeviceServiceClient(conn),
		FileTransfer:     common.NewFileTransferServiceClient(conn),
		Message:          common.NewMessageServiceClient(conn),
		Version:          sys.NewVersionInfoServiceClient(conn),
	}, nil
}

func (a *APIGateway) Close() error {
	return a.conn.Close()
}
