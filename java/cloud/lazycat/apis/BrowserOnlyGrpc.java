package cloud.lazycat.apis;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * this service is provied by frontend server, backend code shouldn't use it.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: browseronly/browseronly.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class BrowserOnlyGrpc {

  private BrowserOnlyGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.BrowserOnly";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.Browseronly.UserInfo> getQueryUserInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryUserInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.Browseronly.UserInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.Browseronly.UserInfo> getQueryUserInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.Browseronly.UserInfo> getQueryUserInfoMethod;
    if ((getQueryUserInfoMethod = BrowserOnlyGrpc.getQueryUserInfoMethod) == null) {
      synchronized (BrowserOnlyGrpc.class) {
        if ((getQueryUserInfoMethod = BrowserOnlyGrpc.getQueryUserInfoMethod) == null) {
          BrowserOnlyGrpc.getQueryUserInfoMethod = getQueryUserInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.Browseronly.UserInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryUserInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.Browseronly.UserInfo.getDefaultInstance()))
              .setSchemaDescriptor(new BrowserOnlyMethodDescriptorSupplier("QueryUserInfo"))
              .build();
        }
      }
    }
    return getQueryUserInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.Browseronly.AppInfo> getQueryAppInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryAppInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.Browseronly.AppInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.Browseronly.AppInfo> getQueryAppInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.Browseronly.AppInfo> getQueryAppInfoMethod;
    if ((getQueryAppInfoMethod = BrowserOnlyGrpc.getQueryAppInfoMethod) == null) {
      synchronized (BrowserOnlyGrpc.class) {
        if ((getQueryAppInfoMethod = BrowserOnlyGrpc.getQueryAppInfoMethod) == null) {
          BrowserOnlyGrpc.getQueryAppInfoMethod = getQueryAppInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.Browseronly.AppInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryAppInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.Browseronly.AppInfo.getDefaultInstance()))
              .setSchemaDescriptor(new BrowserOnlyMethodDescriptorSupplier("QueryAppInfo"))
              .build();
        }
      }
    }
    return getQueryAppInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getPairAllDevicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PairAllDevices",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getPairAllDevicesMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getPairAllDevicesMethod;
    if ((getPairAllDevicesMethod = BrowserOnlyGrpc.getPairAllDevicesMethod) == null) {
      synchronized (BrowserOnlyGrpc.class) {
        if ((getPairAllDevicesMethod = BrowserOnlyGrpc.getPairAllDevicesMethod) == null) {
          BrowserOnlyGrpc.getPairAllDevicesMethod = getPairAllDevicesMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PairAllDevices"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BrowserOnlyMethodDescriptorSupplier("PairAllDevices"))
              .build();
        }
      }
    }
    return getPairAllDevicesMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static BrowserOnlyStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyStub>() {
        @java.lang.Override
        public BrowserOnlyStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BrowserOnlyStub(channel, callOptions);
        }
      };
    return BrowserOnlyStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static BrowserOnlyBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyBlockingStub>() {
        @java.lang.Override
        public BrowserOnlyBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BrowserOnlyBlockingStub(channel, callOptions);
        }
      };
    return BrowserOnlyBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static BrowserOnlyFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyFutureStub>() {
        @java.lang.Override
        public BrowserOnlyFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BrowserOnlyFutureStub(channel, callOptions);
        }
      };
    return BrowserOnlyFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static abstract class BrowserOnlyImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 查询当前登陆用户对应信息
     * </pre>
     */
    public void queryUserInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.Browseronly.UserInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryUserInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * 查询当前访问的lzcapp对应信息
     * </pre>
     */
    public void queryAppInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.Browseronly.AppInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryAppInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * 对devices.proto:_PairAllDeivces的自动封装
     * </pre>
     */
    public void pairAllDevices(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPairAllDevicesMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getQueryUserInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.Browseronly.UserInfo>(
                  this, METHODID_QUERY_USER_INFO)))
          .addMethod(
            getQueryAppInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.Browseronly.AppInfo>(
                  this, METHODID_QUERY_APP_INFO)))
          .addMethod(
            getPairAllDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_PAIR_ALL_DEVICES)))
          .build();
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class BrowserOnlyStub extends io.grpc.stub.AbstractAsyncStub<BrowserOnlyStub> {
    private BrowserOnlyStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BrowserOnlyStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BrowserOnlyStub(channel, callOptions);
    }

    /**
     * <pre>
     * 查询当前登陆用户对应信息
     * </pre>
     */
    public void queryUserInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.Browseronly.UserInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryUserInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 查询当前访问的lzcapp对应信息
     * </pre>
     */
    public void queryAppInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.Browseronly.AppInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryAppInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 对devices.proto:_PairAllDeivces的自动封装
     * </pre>
     */
    public void pairAllDevices(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getPairAllDevicesMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class BrowserOnlyBlockingStub extends io.grpc.stub.AbstractBlockingStub<BrowserOnlyBlockingStub> {
    private BrowserOnlyBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BrowserOnlyBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BrowserOnlyBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 查询当前登陆用户对应信息
     * </pre>
     */
    public cloud.lazycat.apis.Browseronly.UserInfo queryUserInfo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryUserInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 查询当前访问的lzcapp对应信息
     * </pre>
     */
    public cloud.lazycat.apis.Browseronly.AppInfo queryAppInfo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryAppInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 对devices.proto:_PairAllDeivces的自动封装
     * </pre>
     */
    public java.util.Iterator<com.google.protobuf.Empty> pairAllDevices(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getPairAllDevicesMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class BrowserOnlyFutureStub extends io.grpc.stub.AbstractFutureStub<BrowserOnlyFutureStub> {
    private BrowserOnlyFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BrowserOnlyFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BrowserOnlyFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 查询当前登陆用户对应信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.Browseronly.UserInfo> queryUserInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryUserInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 查询当前访问的lzcapp对应信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.Browseronly.AppInfo> queryAppInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryAppInfoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_QUERY_USER_INFO = 0;
  private static final int METHODID_QUERY_APP_INFO = 1;
  private static final int METHODID_PAIR_ALL_DEVICES = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final BrowserOnlyImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(BrowserOnlyImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_QUERY_USER_INFO:
          serviceImpl.queryUserInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.Browseronly.UserInfo>) responseObserver);
          break;
        case METHODID_QUERY_APP_INFO:
          serviceImpl.queryAppInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.Browseronly.AppInfo>) responseObserver);
          break;
        case METHODID_PAIR_ALL_DEVICES:
          serviceImpl.pairAllDevices((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
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

  private static abstract class BrowserOnlyBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    BrowserOnlyBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.Browseronly.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("BrowserOnly");
    }
  }

  private static final class BrowserOnlyFileDescriptorSupplier
      extends BrowserOnlyBaseDescriptorSupplier {
    BrowserOnlyFileDescriptorSupplier() {}
  }

  private static final class BrowserOnlyMethodDescriptorSupplier
      extends BrowserOnlyBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    BrowserOnlyMethodDescriptorSupplier(String methodName) {
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
      synchronized (BrowserOnlyGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new BrowserOnlyFileDescriptorSupplier())
              .addMethod(getQueryUserInfoMethod())
              .addMethod(getQueryAppInfoMethod())
              .addMethod(getPairAllDevicesMethod())
              .build();
        }
      }
    }
    return result;
  }
}
