package cloud.lazycat.apis;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: devices/devices.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class DevicesGrpc {

  private DevicesGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.Devices";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest,
      cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply> getListDevicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListDevices",
      requestType = cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest.class,
      responseType = cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest,
      cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply> getListDevicesMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest, cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply> getListDevicesMethod;
    if ((getListDevicesMethod = DevicesGrpc.getListDevicesMethod) == null) {
      synchronized (DevicesGrpc.class) {
        if ((getListDevicesMethod = DevicesGrpc.getListDevicesMethod) == null) {
          DevicesGrpc.getListDevicesMethod = getListDevicesMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest, cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListDevices"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply.getDefaultInstance()))
              .setSchemaDescriptor(new DevicesMethodDescriptorSupplier("ListDevices"))
              .build();
        }
      }
    }
    return getListDevicesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest,
      com.google.protobuf.Empty> getPairAllDevicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PairAllDevices_",
      requestType = cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest,
      com.google.protobuf.Empty> getPairAllDevicesMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest, com.google.protobuf.Empty> getPairAllDevicesMethod;
    if ((getPairAllDevicesMethod = DevicesGrpc.getPairAllDevicesMethod) == null) {
      synchronized (DevicesGrpc.class) {
        if ((getPairAllDevicesMethod = DevicesGrpc.getPairAllDevicesMethod) == null) {
          DevicesGrpc.getPairAllDevicesMethod = getPairAllDevicesMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PairAllDevices_"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new DevicesMethodDescriptorSupplier("PairAllDevices_"))
              .build();
        }
      }
    }
    return getPairAllDevicesMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DevicesStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevicesStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevicesStub>() {
        @java.lang.Override
        public DevicesStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevicesStub(channel, callOptions);
        }
      };
    return DevicesStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DevicesBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevicesBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevicesBlockingStub>() {
        @java.lang.Override
        public DevicesBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevicesBlockingStub(channel, callOptions);
        }
      };
    return DevicesBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DevicesFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevicesFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevicesFutureStub>() {
        @java.lang.Override
        public DevicesFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevicesFutureStub(channel, callOptions);
        }
      };
    return DevicesFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class DevicesImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public void listDevices(cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListDevicesMethod(), responseObserver);
    }

    /**
     * <pre>
     * 将发起请求的设备与登陆用户的其他所有设备建立其虚拟网络隧道
     * 以便发起请求的浏览器可以绕过盒子直接访问其他节点上的设备API，比如剪贴板、文件拷贝等
     * 后端代码本身就与任意设备建立好了虚拟网络隧道，因此不需要使用此API。
     * 前端代码可以调用browseronly.proto:PairAllDeivce自动设置相关参数。
     * </pre>
     */
    public void pairAllDevices(cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPairAllDevicesMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest,
                cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply>(
                  this, METHODID_LIST_DEVICES)))
          .addMethod(
            getPairAllDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_PAIR_ALL_DEVICES_)))
          .build();
    }
  }

  /**
   */
  public static final class DevicesStub extends io.grpc.stub.AbstractAsyncStub<DevicesStub> {
    private DevicesStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevicesStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevicesStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public void listDevices(cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListDevicesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 将发起请求的设备与登陆用户的其他所有设备建立其虚拟网络隧道
     * 以便发起请求的浏览器可以绕过盒子直接访问其他节点上的设备API，比如剪贴板、文件拷贝等
     * 后端代码本身就与任意设备建立好了虚拟网络隧道，因此不需要使用此API。
     * 前端代码可以调用browseronly.proto:PairAllDeivce自动设置相关参数。
     * </pre>
     */
    public void pairAllDevices(cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getPairAllDevicesMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class DevicesBlockingStub extends io.grpc.stub.AbstractBlockingStub<DevicesBlockingStub> {
    private DevicesBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevicesBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevicesBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply listDevices(cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListDevicesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 将发起请求的设备与登陆用户的其他所有设备建立其虚拟网络隧道
     * 以便发起请求的浏览器可以绕过盒子直接访问其他节点上的设备API，比如剪贴板、文件拷贝等
     * 后端代码本身就与任意设备建立好了虚拟网络隧道，因此不需要使用此API。
     * 前端代码可以调用browseronly.proto:PairAllDeivce自动设置相关参数。
     * </pre>
     */
    public java.util.Iterator<com.google.protobuf.Empty> pairAllDevices(
        cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getPairAllDevicesMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class DevicesFutureStub extends io.grpc.stub.AbstractFutureStub<DevicesFutureStub> {
    private DevicesFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevicesFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevicesFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply> listDevices(
        cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListDevicesMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_DEVICES = 0;
  private static final int METHODID_PAIR_ALL_DEVICES_ = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DevicesImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DevicesImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LIST_DEVICES:
          serviceImpl.listDevices((cloud.lazycat.apis.DevicesOuterClass.ListDeviceRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.DevicesOuterClass.ListDeviceReply>) responseObserver);
          break;
        case METHODID_PAIR_ALL_DEVICES_:
          serviceImpl.pairAllDevices((cloud.lazycat.apis.DevicesOuterClass.PairDeviceRequest) request,
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

  private static abstract class DevicesBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DevicesBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.DevicesOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Devices");
    }
  }

  private static final class DevicesFileDescriptorSupplier
      extends DevicesBaseDescriptorSupplier {
    DevicesFileDescriptorSupplier() {}
  }

  private static final class DevicesMethodDescriptorSupplier
      extends DevicesBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DevicesMethodDescriptorSupplier(String methodName) {
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
      synchronized (DevicesGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DevicesFileDescriptorSupplier())
              .addMethod(getListDevicesMethod())
              .addMethod(getPairAllDevicesMethod())
              .build();
        }
      }
    }
    return result;
  }
}
