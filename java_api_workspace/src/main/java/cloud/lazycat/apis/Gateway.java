package cloud.lazycat.apis;

import cloud.lazycat.apis.common.EndDeviceServiceGrpc;
import cloud.lazycat.apis.common.PermissionManagerGrpc;
import cloud.lazycat.apis.common.UserManagerGrpc;
import cloud.lazycat.apis.localdevice.ClipboardManagerGrpc;
import cloud.lazycat.apis.localdevice.DialogManagerGrpc;
import cloud.lazycat.apis.localdevice.NetworkManagerGrpc;
import cloud.lazycat.apis.localdevice.PhotoLibraryGrpc;
import io.grpc.ChannelCredentials;
import io.grpc.Grpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

import java.io.IOException;

public class Gateway {

    private static final String unixProtocol = "unix:";
    private static final String httpProtocol = "https://";


    public static DeviceProxy newDeviceProxy(String deviceApiUrl) throws IOException {
        if (deviceApiUrl.startsWith(httpProtocol)) {
            deviceApiUrl = deviceApiUrl.substring(httpProtocol.length());
        }
        int splitIndex = deviceApiUrl.lastIndexOf(":");
        String host = deviceApiUrl.substring(0, splitIndex);
        String portStr = deviceApiUrl.substring(splitIndex + 1);
        int port = Integer.parseInt(portStr);


        ManagedChannel channel = Grpc.newChannelBuilderForAddress(host,
                port, CredUtil.DialCred()).build();

        DeviceProxy deviceProxy = new DeviceProxy();
        deviceProxy.channel = channel;
        deviceProxy.clipboard = ClipboardManagerGrpc.newBlockingStub(channel);
        deviceProxy.dialog = DialogManagerGrpc.newBlockingStub(channel);
        deviceProxy.network = NetworkManagerGrpc.newBlockingStub(channel);
        deviceProxy.photoLibrary = PhotoLibraryGrpc.newBlockingStub(channel);

        return deviceProxy;
    }

    public static ManagedChannel NewApiConn() throws IOException {
        ChannelCredentials cred = CredUtil.DialCred();
        ManagedChannelBuilder<?> managedChannelBuilder =
                Grpc.newChannelBuilder(unixProtocol + CredUtil.APISocetPath, cred);
        return managedChannelBuilder.build();
    }

    public static APIGateWay newAPIGateway() throws IOException {
        APIGateWay apiGateWay = new APIGateWay();
        apiGateWay.cred = CredUtil.DialCred();;
        ManagedChannel channel =
                Grpc.newChannelBuilder(unixProtocol + CredUtil.APISocetPath, apiGateWay.cred).build();
        apiGateWay.user = UserManagerGrpc.newBlockingStub(channel);
        apiGateWay.endDevice = EndDeviceServiceGrpc.newBlockingStub(channel);
        apiGateWay.permission = PermissionManagerGrpc.newBlockingStub(channel);
        return apiGateWay;
    }

   public static class APIGateWay {
        public ManagedChannel channel;
        public ChannelCredentials cred;
        public UserManagerGrpc.UserManagerBlockingStub user;
        public EndDeviceServiceGrpc.EndDeviceServiceBlockingStub endDevice;
        public PermissionManagerGrpc.PermissionManagerBlockingStub permission;

    }

    public static class DeviceProxy {
        public ManagedChannel channel;
        public ClipboardManagerGrpc.ClipboardManagerBlockingStub clipboard;
        public DialogManagerGrpc.DialogManagerBlockingStub dialog;
        public PhotoLibraryGrpc.PhotoLibraryBlockingStub photoLibrary;
        public NetworkManagerGrpc.NetworkManagerBlockingStub network;
    }
}
