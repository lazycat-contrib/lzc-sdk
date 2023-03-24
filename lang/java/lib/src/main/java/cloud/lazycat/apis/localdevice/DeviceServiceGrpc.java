package cloud.lazycat.apis.localdevice;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: localdevice/device.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class DeviceServiceGrpc {

  private DeviceServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.localdevice.DeviceService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.localdevice.Device.DeviceInfo> getQueryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Query",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.localdevice.Device.DeviceInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.localdevice.Device.DeviceInfo> getQueryMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.localdevice.Device.DeviceInfo> getQueryMethod;
    if ((getQueryMethod = DeviceServiceGrpc.getQueryMethod) == null) {
      synchronized (DeviceServiceGrpc.class) {
        if ((getQueryMethod = DeviceServiceGrpc.getQueryMethod) == null) {
          DeviceServiceGrpc.getQueryMethod = getQueryMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.localdevice.Device.DeviceInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Query"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Device.DeviceInfo.getDefaultInstance()))
              .setSchemaDescriptor(new DeviceServiceMethodDescriptorSupplier("Query"))
              .build();
        }
      }
    }
    return getQueryMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DeviceServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DeviceServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DeviceServiceStub>() {
        @java.lang.Override
        public DeviceServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DeviceServiceStub(channel, callOptions);
        }
      };
    return DeviceServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DeviceServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DeviceServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DeviceServiceBlockingStub>() {
        @java.lang.Override
        public DeviceServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DeviceServiceBlockingStub(channel, callOptions);
        }
      };
    return DeviceServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DeviceServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DeviceServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DeviceServiceFutureStub>() {
        @java.lang.Override
        public DeviceServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DeviceServiceFutureStub(channel, callOptions);
        }
      };
    return DeviceServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class DeviceServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void query(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Device.DeviceInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getQueryMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.localdevice.Device.DeviceInfo>(
                  this, METHODID_QUERY)))
          .build();
    }
  }

  /**
   */
  public static final class DeviceServiceStub extends io.grpc.stub.AbstractAsyncStub<DeviceServiceStub> {
    private DeviceServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DeviceServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DeviceServiceStub(channel, callOptions);
    }

    /**
     */
    public void query(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Device.DeviceInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class DeviceServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<DeviceServiceBlockingStub> {
    private DeviceServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DeviceServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DeviceServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.localdevice.Device.DeviceInfo query(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class DeviceServiceFutureStub extends io.grpc.stub.AbstractFutureStub<DeviceServiceFutureStub> {
    private DeviceServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DeviceServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DeviceServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.Device.DeviceInfo> query(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_QUERY = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DeviceServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DeviceServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_QUERY:
          serviceImpl.query((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Device.DeviceInfo>) responseObserver);
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

  private static abstract class DeviceServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DeviceServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.localdevice.Device.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("DeviceService");
    }
  }

  private static final class DeviceServiceFileDescriptorSupplier
      extends DeviceServiceBaseDescriptorSupplier {
    DeviceServiceFileDescriptorSupplier() {}
  }

  private static final class DeviceServiceMethodDescriptorSupplier
      extends DeviceServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DeviceServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (DeviceServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DeviceServiceFileDescriptorSupplier())
              .addMethod(getQueryMethod())
              .build();
        }
      }
    }
    return result;
  }
}
