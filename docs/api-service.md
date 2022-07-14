API Service
==========


如何创建一个新的API Service?
==========

API Service分为GRPC和HTTP两种，这里说明如何编写以及注册GRPC服务并通过apiGateway曝露到lzcapis.

1. 在合适的protos/目录下创建服务的proto文件
2. 执行./build.sh生成相关语言的桩代码
3. 更新lzc-apis-protos版本

以上完成接口定义，接下来需要实现具体服务。

以盒子中的服务为例，

1. go mod init sample-service
2. go get -u gitee.com/linakesi/lzc-apis-protos 安装最新的protos代码，或使用go mod edit -replace的方式使用本地未提交版本的protos
3. 编写实现代码。本质是按照proto文件实现grpc server并通过/lzcapp/run/lzc-apis.socket找到系统的apiGateway，并注册自身。

```golang
import (
    pbgw "gitee.com/linakesi/lzc-apis-protos/lang/go/common"
    myService "gitee.com/linakesi/lzc-apis-protos/lang/go/common" //引入上面提交的新服务
)
...

type myService {
}

	s := grpc.NewServer() //创建一个grpc server对象
	myService.RegisterMyServiceServer(s, &serviceImpl{}) //在此grpc server上注册一个服务实现

	l, err := net.Listen("tcp", "0.0.0.0:3999")
	if err != nil {
		return "", err
	}
	go s.Serve(l) //监听并提供grpc服务


    // 向/lzcapp/run/lzc-apis.socket拨号
	conn, err := lzcapis.NewAPIConn()
	if err != nil {
		fmt.Println("Can't connect to lzc-api.socket", err)
		return
	}
    // 创建apiGateway对象。lzc-api上的其他服务也可以通过此conn直接获取到
	gw, err := pbgw.NewAPIGatewayClient(conn)

    // 向gateway注册上面刚创建的grpc服务。
	gw.RegisterService(context.TODO(), &pbgw.ServiceInfo{
		ServiceType: pbgw.ServiceType_GRPC,
		ServiceName: myService.MyService_ServiceDesc.ServiceName, //固定字段，由proto文件自动生成
		ServiceAddr: "http://myservice-myservice.lzcapp:3999", //根据lzcdns规范曝露服务地址。
    })

```

4. 编写app描述文件。关键是配置正确的服务名称，以便上述`ServiceAddr`中的内部`lzc dns`地址和端口是正确的。注意如果没有特殊需求，不要在
   api service中通过`ingress`的方式曝露任何端口到外部。

```
//docker-compose.yml
x-lazycat-app:
  id: myservice

services:
  dns:
    pull_policy: build
    build: ./src


//src/Dockerfile
FROM golang:1.16-alpine AS build

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download -x

COPY *.go ./

RUN go build -o /my-service

FROM alpine:latest
COPY --from=build /my-service /my-service

ENTRYPOINT [ "/my-service" ]

```

5. 部署此服务到盒子. `lzc-cli deploy`
