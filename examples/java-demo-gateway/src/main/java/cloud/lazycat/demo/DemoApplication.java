package cloud.lazycat.demo;

import cloud.lazycat.apis.Gateway;
import cloud.lazycat.apis.common.EndDeviceOuterClass;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;

@SpringBootApplication
@RestController
public class DemoApplication {

	public DemoApplication() throws IOException {
	}

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}



	@GetMapping("/getDeviceList")
	public Object getAlbums() throws IOException {
		Gateway.APIGateWay apiGateWay = Gateway.newAPIGateway();
		EndDeviceOuterClass.ListEndDeviceRequest.Builder builder = EndDeviceOuterClass.ListEndDeviceRequest.newBuilder();
		EndDeviceOuterClass.ListEndDeviceReply listEndDeviceReply = apiGateWay.endDevice.listEndDevices(builder.build());
		return listEndDeviceReply.getDevicesList();
	}
}