import io.grpc.ChannelCredentials;
import io.grpc.TlsChannelCredentials;

import java.io.File;
import java.io.IOException;

public abstract class CredUtil {

    private static final String CAPath = "/lzcapp/run/certs/box.crt";
    private static final String APPCertPath = "/lzcapp/run/certs/app.crt";
    private static final String APPKeyPath = "/lzcapp/run/certs/app.key";
    public static final String APISocetPath = "/lzcapp/run/sys/lzc-api.socket";


    public static ChannelCredentials buildClientCredOption(String caCrt, String appKey, String appCrt) throws IOException {
        TlsChannelCredentials.Builder tlsBuilder = TlsChannelCredentials.newBuilder();
        tlsBuilder.keyManager(new File(APPCertPath), new File(APPKeyPath));
        tlsBuilder.trustManager(new File(CAPath));
        return tlsBuilder.build();
    }

   public static ChannelCredentials DialCred() throws IOException {
       return buildClientCredOption(APPCertPath, APPCertPath, CAPath);
   }

}

