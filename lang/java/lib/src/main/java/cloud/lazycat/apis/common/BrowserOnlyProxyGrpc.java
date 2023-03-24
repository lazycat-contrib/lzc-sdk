package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * this service is provied by frontend server, backend code shouldn't use it.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/browseronly.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class BrowserOnlyProxyGrpc {

  private BrowserOnlyProxyGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.BrowserOnlyProxy";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Browseronly.SessionInfo> getQuerySessionInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QuerySessionInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.common.Browseronly.SessionInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Browseronly.SessionInfo> getQuerySessionInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.common.Browseronly.SessionInfo> getQuerySessionInfoMethod;
    if ((getQuerySessionInfoMethod = BrowserOnlyProxyGrpc.getQuerySessionInfoMethod) == null) {
      synchronized (BrowserOnlyProxyGrpc.class) {
        if ((getQuerySessionInfoMethod = BrowserOnlyProxyGrpc.getQuerySessionInfoMethod) == null) {
          BrowserOnlyProxyGrpc.getQuerySessionInfoMethod = getQuerySessionInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.common.Browseronly.SessionInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QuerySessionInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Browseronly.SessionInfo.getDefaultInstance()))
              .setSchemaDescriptor(new BrowserOnlyProxyMethodDescriptorSupplier("QuerySessionInfo"))
              .build();
        }
      }
    }
    return getQuerySessionInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Browseronly.AppInfo> getQueryAppInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryAppInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.common.Browseronly.AppInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Browseronly.AppInfo> getQueryAppInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.common.Browseronly.AppInfo> getQueryAppInfoMethod;
    if ((getQueryAppInfoMethod = BrowserOnlyProxyGrpc.getQueryAppInfoMethod) == null) {
      synchronized (BrowserOnlyProxyGrpc.class) {
        if ((getQueryAppInfoMethod = BrowserOnlyProxyGrpc.getQueryAppInfoMethod) == null) {
          BrowserOnlyProxyGrpc.getQueryAppInfoMethod = getQueryAppInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.common.Browseronly.AppInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryAppInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Browseronly.AppInfo.getDefaultInstance()))
              .setSchemaDescriptor(new BrowserOnlyProxyMethodDescriptorSupplier("QueryAppInfo"))
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
    if ((getPairAllDevicesMethod = BrowserOnlyProxyGrpc.getPairAllDevicesMethod) == null) {
      synchronized (BrowserOnlyProxyGrpc.class) {
        if ((getPairAllDevicesMethod = BrowserOnlyProxyGrpc.getPairAllDevicesMethod) == null) {
          BrowserOnlyProxyGrpc.getPairAllDevicesMethod = getPairAllDevicesMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PairAllDevices"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BrowserOnlyProxyMethodDescriptorSupplier("PairAllDevices"))
              .build();
        }
      }
    }
    return getPairAllDevicesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Browseronly.APIServerInfo> getQueryAPIServerInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryAPIServerInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.common.Browseronly.APIServerInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Browseronly.APIServerInfo> getQueryAPIServerInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.common.Browseronly.APIServerInfo> getQueryAPIServerInfoMethod;
    if ((getQueryAPIServerInfoMethod = BrowserOnlyProxyGrpc.getQueryAPIServerInfoMethod) == null) {
      synchronized (BrowserOnlyProxyGrpc.class) {
        if ((getQueryAPIServerInfoMethod = BrowserOnlyProxyGrpc.getQueryAPIServerInfoMethod) == null) {
          BrowserOnlyProxyGrpc.getQueryAPIServerInfoMethod = getQueryAPIServerInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.common.Browseronly.APIServerInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryAPIServerInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Browseronly.APIServerInfo.getDefaultInstance()))
              .setSchemaDescriptor(new BrowserOnlyProxyMethodDescriptorSupplier("QueryAPIServerInfo"))
              .build();
        }
      }
    }
    return getQueryAPIServerInfoMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static BrowserOnlyProxyStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyProxyStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyProxyStub>() {
        @java.lang.Override
        public BrowserOnlyProxyStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BrowserOnlyProxyStub(channel, callOptions);
        }
      };
    return BrowserOnlyProxyStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static BrowserOnlyProxyBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyProxyBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyProxyBlockingStub>() {
        @java.lang.Override
        public BrowserOnlyProxyBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BrowserOnlyProxyBlockingStub(channel, callOptions);
        }
      };
    return BrowserOnlyProxyBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static BrowserOnlyProxyFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyProxyFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BrowserOnlyProxyFutureStub>() {
        @java.lang.Override
        public BrowserOnlyProxyFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BrowserOnlyProxyFutureStub(channel, callOptions);
        }
      };
    return BrowserOnlyProxyFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static abstract class BrowserOnlyProxyImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 查询当前登陆用户对应信息
     * </pre>
     */
    public void querySessionInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.SessionInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQuerySessionInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * 查询当前访问的lzcapp对应信息
     * </pre>
     */
    public void queryAppInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.AppInfo> responseObserver) {
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

    /**
     */
    public void queryAPIServerInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.APIServerInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryAPIServerInfoMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getQuerySessionInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.common.Browseronly.SessionInfo>(
                  this, METHODID_QUERY_SESSION_INFO)))
          .addMethod(
            getQueryAppInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.common.Browseronly.AppInfo>(
                  this, METHODID_QUERY_APP_INFO)))
          .addMethod(
            getPairAllDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_PAIR_ALL_DEVICES)))
          .addMethod(
            getQueryAPIServerInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.common.Browseronly.APIServerInfo>(
                  this, METHODID_QUERY_APISERVER_INFO)))
          .build();
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class BrowserOnlyProxyStub extends io.grpc.stub.AbstractAsyncStub<BrowserOnlyProxyStub> {
    private BrowserOnlyProxyStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BrowserOnlyProxyStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BrowserOnlyProxyStub(channel, callOptions);
    }

    /**
     * <pre>
     * 查询当前登陆用户对应信息
     * </pre>
     */
    public void querySessionInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.SessionInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQuerySessionInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 查询当前访问的lzcapp对应信息
     * </pre>
     */
    public void queryAppInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.AppInfo> responseObserver) {
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

    /**
     */
    public void queryAPIServerInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.APIServerInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryAPIServerInfoMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class BrowserOnlyProxyBlockingStub extends io.grpc.stub.AbstractBlockingStub<BrowserOnlyProxyBlockingStub> {
    private BrowserOnlyProxyBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BrowserOnlyProxyBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BrowserOnlyProxyBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 查询当前登陆用户对应信息
     * </pre>
     */
    public cloud.lazycat.apis.common.Browseronly.SessionInfo querySessionInfo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQuerySessionInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 查询当前访问的lzcapp对应信息
     * </pre>
     */
    public cloud.lazycat.apis.common.Browseronly.AppInfo queryAppInfo(com.google.protobuf.Empty request) {
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

    /**
     */
    public cloud.lazycat.apis.common.Browseronly.APIServerInfo queryAPIServerInfo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryAPIServerInfoMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class BrowserOnlyProxyFutureStub extends io.grpc.stub.AbstractFutureStub<BrowserOnlyProxyFutureStub> {
    private BrowserOnlyProxyFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BrowserOnlyProxyFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BrowserOnlyProxyFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 查询当前登陆用户对应信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Browseronly.SessionInfo> querySessionInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQuerySessionInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 查询当前访问的lzcapp对应信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Browseronly.AppInfo> queryAppInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryAppInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Browseronly.APIServerInfo> queryAPIServerInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryAPIServerInfoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_QUERY_SESSION_INFO = 0;
  private static final int METHODID_QUERY_APP_INFO = 1;
  private static final int METHODID_PAIR_ALL_DEVICES = 2;
  private static final int METHODID_QUERY_APISERVER_INFO = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final BrowserOnlyProxyImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(BrowserOnlyProxyImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_QUERY_SESSION_INFO:
          serviceImpl.querySessionInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.SessionInfo>) responseObserver);
          break;
        case METHODID_QUERY_APP_INFO:
          serviceImpl.queryAppInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.AppInfo>) responseObserver);
          break;
        case METHODID_PAIR_ALL_DEVICES:
          serviceImpl.pairAllDevices((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_QUERY_APISERVER_INFO:
          serviceImpl.queryAPIServerInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Browseronly.APIServerInfo>) responseObserver);
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

  private static abstract class BrowserOnlyProxyBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    BrowserOnlyProxyBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.Browseronly.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("BrowserOnlyProxy");
    }
  }

  private static final class BrowserOnlyProxyFileDescriptorSupplier
      extends BrowserOnlyProxyBaseDescriptorSupplier {
    BrowserOnlyProxyFileDescriptorSupplier() {}
  }

  private static final class BrowserOnlyProxyMethodDescriptorSupplier
      extends BrowserOnlyProxyBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    BrowserOnlyProxyMethodDescriptorSupplier(String methodName) {
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
      synchronized (BrowserOnlyProxyGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new BrowserOnlyProxyFileDescriptorSupplier())
              .addMethod(getQuerySessionInfoMethod())
              .addMethod(getQueryAppInfoMethod())
              .addMethod(getPairAllDevicesMethod())
              .addMethod(getQueryAPIServerInfoMethod())
              .build();
        }
      }
    }
    return result;
  }
}
