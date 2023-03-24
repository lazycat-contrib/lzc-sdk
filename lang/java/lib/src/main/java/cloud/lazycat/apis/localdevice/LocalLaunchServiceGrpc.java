package cloud.lazycat.apis.localdevice;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: localdevice/local-launch.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class LocalLaunchServiceGrpc {

  private LocalLaunchServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.localdevice.LocalLaunchService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply> getPinAppMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PinApp",
      requestType = cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest.class,
      responseType = cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply> getPinAppMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest, cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply> getPinAppMethod;
    if ((getPinAppMethod = LocalLaunchServiceGrpc.getPinAppMethod) == null) {
      synchronized (LocalLaunchServiceGrpc.class) {
        if ((getPinAppMethod = LocalLaunchServiceGrpc.getPinAppMethod) == null) {
          LocalLaunchServiceGrpc.getPinAppMethod = getPinAppMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest, cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PinApp"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply.getDefaultInstance()))
              .setSchemaDescriptor(new LocalLaunchServiceMethodDescriptorSupplier("PinApp"))
              .build();
        }
      }
    }
    return getPinAppMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply> getUnPinAppMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UnPinApp",
      requestType = cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest.class,
      responseType = cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply> getUnPinAppMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest, cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply> getUnPinAppMethod;
    if ((getUnPinAppMethod = LocalLaunchServiceGrpc.getUnPinAppMethod) == null) {
      synchronized (LocalLaunchServiceGrpc.class) {
        if ((getUnPinAppMethod = LocalLaunchServiceGrpc.getUnPinAppMethod) == null) {
          LocalLaunchServiceGrpc.getUnPinAppMethod = getUnPinAppMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest, cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UnPinApp"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply.getDefaultInstance()))
              .setSchemaDescriptor(new LocalLaunchServiceMethodDescriptorSupplier("UnPinApp"))
              .build();
        }
      }
    }
    return getUnPinAppMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> getOpenAppMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "OpenApp",
      requestType = cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest.class,
      responseType = cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> getOpenAppMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest, cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> getOpenAppMethod;
    if ((getOpenAppMethod = LocalLaunchServiceGrpc.getOpenAppMethod) == null) {
      synchronized (LocalLaunchServiceGrpc.class) {
        if ((getOpenAppMethod = LocalLaunchServiceGrpc.getOpenAppMethod) == null) {
          LocalLaunchServiceGrpc.getOpenAppMethod = getOpenAppMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest, cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "OpenApp"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply.getDefaultInstance()))
              .setSchemaDescriptor(new LocalLaunchServiceMethodDescriptorSupplier("OpenApp"))
              .build();
        }
      }
    }
    return getOpenAppMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> getOpenUnsafeAppMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "OpenUnsafeApp",
      requestType = cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest.class,
      responseType = cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> getOpenUnsafeAppMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest, cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> getOpenUnsafeAppMethod;
    if ((getOpenUnsafeAppMethod = LocalLaunchServiceGrpc.getOpenUnsafeAppMethod) == null) {
      synchronized (LocalLaunchServiceGrpc.class) {
        if ((getOpenUnsafeAppMethod = LocalLaunchServiceGrpc.getOpenUnsafeAppMethod) == null) {
          LocalLaunchServiceGrpc.getOpenUnsafeAppMethod = getOpenUnsafeAppMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest, cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "OpenUnsafeApp"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply.getDefaultInstance()))
              .setSchemaDescriptor(new LocalLaunchServiceMethodDescriptorSupplier("OpenUnsafeApp"))
              .build();
        }
      }
    }
    return getOpenUnsafeAppMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply> getOpenAppMethodMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "OpenAppMethod",
      requestType = cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest.class,
      responseType = cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest,
      cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply> getOpenAppMethodMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest, cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply> getOpenAppMethodMethod;
    if ((getOpenAppMethodMethod = LocalLaunchServiceGrpc.getOpenAppMethodMethod) == null) {
      synchronized (LocalLaunchServiceGrpc.class) {
        if ((getOpenAppMethodMethod = LocalLaunchServiceGrpc.getOpenAppMethodMethod) == null) {
          LocalLaunchServiceGrpc.getOpenAppMethodMethod = getOpenAppMethodMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest, cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "OpenAppMethod"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply.getDefaultInstance()))
              .setSchemaDescriptor(new LocalLaunchServiceMethodDescriptorSupplier("OpenAppMethod"))
              .build();
        }
      }
    }
    return getOpenAppMethodMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static LocalLaunchServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LocalLaunchServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LocalLaunchServiceStub>() {
        @java.lang.Override
        public LocalLaunchServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LocalLaunchServiceStub(channel, callOptions);
        }
      };
    return LocalLaunchServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static LocalLaunchServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LocalLaunchServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LocalLaunchServiceBlockingStub>() {
        @java.lang.Override
        public LocalLaunchServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LocalLaunchServiceBlockingStub(channel, callOptions);
        }
      };
    return LocalLaunchServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static LocalLaunchServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LocalLaunchServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LocalLaunchServiceFutureStub>() {
        @java.lang.Override
        public LocalLaunchServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LocalLaunchServiceFutureStub(channel, callOptions);
        }
      };
    return LocalLaunchServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class LocalLaunchServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 创建快捷方式
     * </pre>
     */
    public void pinApp(cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPinAppMethod(), responseObserver);
    }

    /**
     * <pre>
     * 删除快捷方式
     * </pre>
     */
    public void unPinApp(cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUnPinAppMethod(), responseObserver);
    }

    /**
     * <pre>
     * 打开指定的懒猫云应用
     * </pre>
     */
    public void openApp(cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getOpenAppMethod(), responseObserver);
    }

    /**
     * <pre>
     * 使用不安全模式打开指定APP
     * App页面由两个webview组成：一个control-view,一个content-view
     * 其中app渲染在control-view中，content-view的内容由App代码控制
     * </pre>
     */
    public void openUnsafeApp(cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getOpenUnsafeAppMethod(), responseObserver);
    }

    /**
     * <pre>
     * 当前设备支持的应用打开方式
     * </pre>
     */
    public void openAppMethod(cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getOpenAppMethodMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getPinAppMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest,
                cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply>(
                  this, METHODID_PIN_APP)))
          .addMethod(
            getUnPinAppMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest,
                cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply>(
                  this, METHODID_UN_PIN_APP)))
          .addMethod(
            getOpenAppMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest,
                cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply>(
                  this, METHODID_OPEN_APP)))
          .addMethod(
            getOpenUnsafeAppMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest,
                cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply>(
                  this, METHODID_OPEN_UNSAFE_APP)))
          .addMethod(
            getOpenAppMethodMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest,
                cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply>(
                  this, METHODID_OPEN_APP_METHOD)))
          .build();
    }
  }

  /**
   */
  public static final class LocalLaunchServiceStub extends io.grpc.stub.AbstractAsyncStub<LocalLaunchServiceStub> {
    private LocalLaunchServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LocalLaunchServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LocalLaunchServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 创建快捷方式
     * </pre>
     */
    public void pinApp(cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPinAppMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 删除快捷方式
     * </pre>
     */
    public void unPinApp(cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUnPinAppMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 打开指定的懒猫云应用
     * </pre>
     */
    public void openApp(cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getOpenAppMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 使用不安全模式打开指定APP
     * App页面由两个webview组成：一个control-view,一个content-view
     * 其中app渲染在control-view中，content-view的内容由App代码控制
     * </pre>
     */
    public void openUnsafeApp(cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getOpenUnsafeAppMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 当前设备支持的应用打开方式
     * </pre>
     */
    public void openAppMethod(cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getOpenAppMethodMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class LocalLaunchServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<LocalLaunchServiceBlockingStub> {
    private LocalLaunchServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LocalLaunchServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LocalLaunchServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 创建快捷方式
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply pinApp(cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPinAppMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 删除快捷方式
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply unPinApp(cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUnPinAppMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 打开指定的懒猫云应用
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply openApp(cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getOpenAppMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 使用不安全模式打开指定APP
     * App页面由两个webview组成：一个control-view,一个content-view
     * 其中app渲染在control-view中，content-view的内容由App代码控制
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply openUnsafeApp(cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getOpenUnsafeAppMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 当前设备支持的应用打开方式
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply openAppMethod(cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getOpenAppMethodMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class LocalLaunchServiceFutureStub extends io.grpc.stub.AbstractFutureStub<LocalLaunchServiceFutureStub> {
    private LocalLaunchServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LocalLaunchServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LocalLaunchServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 创建快捷方式
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply> pinApp(
        cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPinAppMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 删除快捷方式
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply> unPinApp(
        cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUnPinAppMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 打开指定的懒猫云应用
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> openApp(
        cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getOpenAppMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 使用不安全模式打开指定APP
     * App页面由两个webview组成：一个control-view,一个content-view
     * 其中app渲染在control-view中，content-view的内容由App代码控制
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply> openUnsafeApp(
        cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getOpenUnsafeAppMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 当前设备支持的应用打开方式
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply> openAppMethod(
        cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getOpenAppMethodMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_PIN_APP = 0;
  private static final int METHODID_UN_PIN_APP = 1;
  private static final int METHODID_OPEN_APP = 2;
  private static final int METHODID_OPEN_UNSAFE_APP = 3;
  private static final int METHODID_OPEN_APP_METHOD = 4;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final LocalLaunchServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(LocalLaunchServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_PIN_APP:
          serviceImpl.pinApp((cloud.lazycat.apis.localdevice.LocalLaunch.PinAppRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.PinAppReply>) responseObserver);
          break;
        case METHODID_UN_PIN_APP:
          serviceImpl.unPinApp((cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.UnPinAppReply>) responseObserver);
          break;
        case METHODID_OPEN_APP:
          serviceImpl.openApp((cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply>) responseObserver);
          break;
        case METHODID_OPEN_UNSAFE_APP:
          serviceImpl.openUnsafeApp((cloud.lazycat.apis.localdevice.LocalLaunch.OpenUnsafeAppRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppReply>) responseObserver);
          break;
        case METHODID_OPEN_APP_METHOD:
          serviceImpl.openAppMethod((cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.LocalLaunch.OpenAppMethodReply>) responseObserver);
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

  private static abstract class LocalLaunchServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    LocalLaunchServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.localdevice.LocalLaunch.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("LocalLaunchService");
    }
  }

  private static final class LocalLaunchServiceFileDescriptorSupplier
      extends LocalLaunchServiceBaseDescriptorSupplier {
    LocalLaunchServiceFileDescriptorSupplier() {}
  }

  private static final class LocalLaunchServiceMethodDescriptorSupplier
      extends LocalLaunchServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    LocalLaunchServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (LocalLaunchServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new LocalLaunchServiceFileDescriptorSupplier())
              .addMethod(getPinAppMethod())
              .addMethod(getUnPinAppMethod())
              .addMethod(getOpenAppMethod())
              .addMethod(getOpenUnsafeAppMethod())
              .addMethod(getOpenAppMethodMethod())
              .build();
        }
      }
    }
    return result;
  }
}
