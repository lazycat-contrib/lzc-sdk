package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/gateway.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class APIGatewayGrpc {

  private APIGatewayGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.APIGateway";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Gateway.GatewayInfo> getQueryGatewayInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryGatewayInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.common.Gateway.GatewayInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Gateway.GatewayInfo> getQueryGatewayInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.common.Gateway.GatewayInfo> getQueryGatewayInfoMethod;
    if ((getQueryGatewayInfoMethod = APIGatewayGrpc.getQueryGatewayInfoMethod) == null) {
      synchronized (APIGatewayGrpc.class) {
        if ((getQueryGatewayInfoMethod = APIGatewayGrpc.getQueryGatewayInfoMethod) == null) {
          APIGatewayGrpc.getQueryGatewayInfoMethod = getQueryGatewayInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.common.Gateway.GatewayInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryGatewayInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Gateway.GatewayInfo.getDefaultInstance()))
              .setSchemaDescriptor(new APIGatewayMethodDescriptorSupplier("QueryGatewayInfo"))
              .build();
        }
      }
    }
    return getQueryGatewayInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Gateway.ServiceInfo,
      cloud.lazycat.apis.common.Gateway.RegisterReply> getRegisterServiceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RegisterService",
      requestType = cloud.lazycat.apis.common.Gateway.ServiceInfo.class,
      responseType = cloud.lazycat.apis.common.Gateway.RegisterReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Gateway.ServiceInfo,
      cloud.lazycat.apis.common.Gateway.RegisterReply> getRegisterServiceMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Gateway.ServiceInfo, cloud.lazycat.apis.common.Gateway.RegisterReply> getRegisterServiceMethod;
    if ((getRegisterServiceMethod = APIGatewayGrpc.getRegisterServiceMethod) == null) {
      synchronized (APIGatewayGrpc.class) {
        if ((getRegisterServiceMethod = APIGatewayGrpc.getRegisterServiceMethod) == null) {
          APIGatewayGrpc.getRegisterServiceMethod = getRegisterServiceMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Gateway.ServiceInfo, cloud.lazycat.apis.common.Gateway.RegisterReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RegisterService"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Gateway.ServiceInfo.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Gateway.RegisterReply.getDefaultInstance()))
              .setSchemaDescriptor(new APIGatewayMethodDescriptorSupplier("RegisterService"))
              .build();
        }
      }
    }
    return getRegisterServiceMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest,
      cloud.lazycat.apis.common.Gateway.HTTPAccessInfo> getAccessHTTPServiceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AccessHTTPService",
      requestType = cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest.class,
      responseType = cloud.lazycat.apis.common.Gateway.HTTPAccessInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest,
      cloud.lazycat.apis.common.Gateway.HTTPAccessInfo> getAccessHTTPServiceMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest, cloud.lazycat.apis.common.Gateway.HTTPAccessInfo> getAccessHTTPServiceMethod;
    if ((getAccessHTTPServiceMethod = APIGatewayGrpc.getAccessHTTPServiceMethod) == null) {
      synchronized (APIGatewayGrpc.class) {
        if ((getAccessHTTPServiceMethod = APIGatewayGrpc.getAccessHTTPServiceMethod) == null) {
          APIGatewayGrpc.getAccessHTTPServiceMethod = getAccessHTTPServiceMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest, cloud.lazycat.apis.common.Gateway.HTTPAccessInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AccessHTTPService"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Gateway.HTTPAccessInfo.getDefaultInstance()))
              .setSchemaDescriptor(new APIGatewayMethodDescriptorSupplier("AccessHTTPService"))
              .build();
        }
      }
    }
    return getAccessHTTPServiceMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static APIGatewayStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<APIGatewayStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<APIGatewayStub>() {
        @java.lang.Override
        public APIGatewayStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new APIGatewayStub(channel, callOptions);
        }
      };
    return APIGatewayStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static APIGatewayBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<APIGatewayBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<APIGatewayBlockingStub>() {
        @java.lang.Override
        public APIGatewayBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new APIGatewayBlockingStub(channel, callOptions);
        }
      };
    return APIGatewayBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static APIGatewayFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<APIGatewayFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<APIGatewayFutureStub>() {
        @java.lang.Override
        public APIGatewayFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new APIGatewayFutureStub(channel, callOptions);
        }
      };
    return APIGatewayFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class APIGatewayImplBase implements io.grpc.BindableService {

    /**
     */
    public void queryGatewayInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.GatewayInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryGatewayInfoMethod(), responseObserver);
    }

    /**
     */
    public void registerService(cloud.lazycat.apis.common.Gateway.ServiceInfo request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.RegisterReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRegisterServiceMethod(), responseObserver);
    }

    /**
     * <pre>
     * 访问http service前，需要申请auth信息以便Gateway可以识别来源对应的
     * 真实appid。
     * 返回值尽量交由浏览器前端设置xhr对应参数。如果确实无法做到(比如img src中)，则
     * 可以通过https://$appdomain/$service_proxy/的形式由/usr/bin/lzcapp代为
     * 转发。此前缀可以通过BrowserOnlyProxy.proto接口获取到。
     * </pre>
     */
    public void accessHTTPService(cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.HTTPAccessInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAccessHTTPServiceMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getQueryGatewayInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.common.Gateway.GatewayInfo>(
                  this, METHODID_QUERY_GATEWAY_INFO)))
          .addMethod(
            getRegisterServiceMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Gateway.ServiceInfo,
                cloud.lazycat.apis.common.Gateway.RegisterReply>(
                  this, METHODID_REGISTER_SERVICE)))
          .addMethod(
            getAccessHTTPServiceMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest,
                cloud.lazycat.apis.common.Gateway.HTTPAccessInfo>(
                  this, METHODID_ACCESS_HTTPSERVICE)))
          .build();
    }
  }

  /**
   */
  public static final class APIGatewayStub extends io.grpc.stub.AbstractAsyncStub<APIGatewayStub> {
    private APIGatewayStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected APIGatewayStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new APIGatewayStub(channel, callOptions);
    }

    /**
     */
    public void queryGatewayInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.GatewayInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryGatewayInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void registerService(cloud.lazycat.apis.common.Gateway.ServiceInfo request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.RegisterReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRegisterServiceMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 访问http service前，需要申请auth信息以便Gateway可以识别来源对应的
     * 真实appid。
     * 返回值尽量交由浏览器前端设置xhr对应参数。如果确实无法做到(比如img src中)，则
     * 可以通过https://$appdomain/$service_proxy/的形式由/usr/bin/lzcapp代为
     * 转发。此前缀可以通过BrowserOnlyProxy.proto接口获取到。
     * </pre>
     */
    public void accessHTTPService(cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.HTTPAccessInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAccessHTTPServiceMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class APIGatewayBlockingStub extends io.grpc.stub.AbstractBlockingStub<APIGatewayBlockingStub> {
    private APIGatewayBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected APIGatewayBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new APIGatewayBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.common.Gateway.GatewayInfo queryGatewayInfo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryGatewayInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.common.Gateway.RegisterReply registerService(cloud.lazycat.apis.common.Gateway.ServiceInfo request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRegisterServiceMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 访问http service前，需要申请auth信息以便Gateway可以识别来源对应的
     * 真实appid。
     * 返回值尽量交由浏览器前端设置xhr对应参数。如果确实无法做到(比如img src中)，则
     * 可以通过https://$appdomain/$service_proxy/的形式由/usr/bin/lzcapp代为
     * 转发。此前缀可以通过BrowserOnlyProxy.proto接口获取到。
     * </pre>
     */
    public cloud.lazycat.apis.common.Gateway.HTTPAccessInfo accessHTTPService(cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAccessHTTPServiceMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class APIGatewayFutureStub extends io.grpc.stub.AbstractFutureStub<APIGatewayFutureStub> {
    private APIGatewayFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected APIGatewayFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new APIGatewayFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Gateway.GatewayInfo> queryGatewayInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryGatewayInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Gateway.RegisterReply> registerService(
        cloud.lazycat.apis.common.Gateway.ServiceInfo request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRegisterServiceMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 访问http service前，需要申请auth信息以便Gateway可以识别来源对应的
     * 真实appid。
     * 返回值尽量交由浏览器前端设置xhr对应参数。如果确实无法做到(比如img src中)，则
     * 可以通过https://$appdomain/$service_proxy/的形式由/usr/bin/lzcapp代为
     * 转发。此前缀可以通过BrowserOnlyProxy.proto接口获取到。
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Gateway.HTTPAccessInfo> accessHTTPService(
        cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAccessHTTPServiceMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_QUERY_GATEWAY_INFO = 0;
  private static final int METHODID_REGISTER_SERVICE = 1;
  private static final int METHODID_ACCESS_HTTPSERVICE = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final APIGatewayImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(APIGatewayImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_QUERY_GATEWAY_INFO:
          serviceImpl.queryGatewayInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.GatewayInfo>) responseObserver);
          break;
        case METHODID_REGISTER_SERVICE:
          serviceImpl.registerService((cloud.lazycat.apis.common.Gateway.ServiceInfo) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.RegisterReply>) responseObserver);
          break;
        case METHODID_ACCESS_HTTPSERVICE:
          serviceImpl.accessHTTPService((cloud.lazycat.apis.common.Gateway.AccessHTTPServiceRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Gateway.HTTPAccessInfo>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class APIGatewayBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    APIGatewayBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.Gateway.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("APIGateway");
    }
  }

  private static final class APIGatewayFileDescriptorSupplier
      extends APIGatewayBaseDescriptorSupplier {
    APIGatewayFileDescriptorSupplier() {}
  }

  private static final class APIGatewayMethodDescriptorSupplier
      extends APIGatewayBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    APIGatewayMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (APIGatewayGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new APIGatewayFileDescriptorSupplier())
              .addMethod(getQueryGatewayInfoMethod())
              .addMethod(getRegisterServiceMethod())
              .addMethod(getAccessHTTPServiceMethod())
              .build();
        }
      }
    }
    return result;
  }
}
