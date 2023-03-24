package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/devopt.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class DevOptServiceGrpc {

  private DevOptServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.DevOptService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.Devopt.DeveloperOptions> getGetDeveloperOptionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDeveloperOptions",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.Devopt.DeveloperOptions.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.Devopt.DeveloperOptions> getGetDeveloperOptionsMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.Devopt.DeveloperOptions> getGetDeveloperOptionsMethod;
    if ((getGetDeveloperOptionsMethod = DevOptServiceGrpc.getGetDeveloperOptionsMethod) == null) {
      synchronized (DevOptServiceGrpc.class) {
        if ((getGetDeveloperOptionsMethod = DevOptServiceGrpc.getGetDeveloperOptionsMethod) == null) {
          DevOptServiceGrpc.getGetDeveloperOptionsMethod = getGetDeveloperOptionsMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.Devopt.DeveloperOptions>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDeveloperOptions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Devopt.DeveloperOptions.getDefaultInstance()))
              .setSchemaDescriptor(new DevOptServiceMethodDescriptorSupplier("GetDeveloperOptions"))
              .build();
        }
      }
    }
    return getGetDeveloperOptionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Devopt.DeveloperOptions,
      com.google.protobuf.Empty> getSetDeveloperOptionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetDeveloperOptions",
      requestType = cloud.lazycat.apis.sys.Devopt.DeveloperOptions.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Devopt.DeveloperOptions,
      com.google.protobuf.Empty> getSetDeveloperOptionsMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Devopt.DeveloperOptions, com.google.protobuf.Empty> getSetDeveloperOptionsMethod;
    if ((getSetDeveloperOptionsMethod = DevOptServiceGrpc.getSetDeveloperOptionsMethod) == null) {
      synchronized (DevOptServiceGrpc.class) {
        if ((getSetDeveloperOptionsMethod = DevOptServiceGrpc.getSetDeveloperOptionsMethod) == null) {
          DevOptServiceGrpc.getSetDeveloperOptionsMethod = getSetDeveloperOptionsMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Devopt.DeveloperOptions, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetDeveloperOptions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Devopt.DeveloperOptions.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new DevOptServiceMethodDescriptorSupplier("SetDeveloperOptions"))
              .build();
        }
      }
    }
    return getSetDeveloperOptionsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DevOptServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevOptServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevOptServiceStub>() {
        @java.lang.Override
        public DevOptServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevOptServiceStub(channel, callOptions);
        }
      };
    return DevOptServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DevOptServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevOptServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevOptServiceBlockingStub>() {
        @java.lang.Override
        public DevOptServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevOptServiceBlockingStub(channel, callOptions);
        }
      };
    return DevOptServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DevOptServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DevOptServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DevOptServiceFutureStub>() {
        @java.lang.Override
        public DevOptServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DevOptServiceFutureStub(channel, callOptions);
        }
      };
    return DevOptServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class DevOptServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void getDeveloperOptions(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Devopt.DeveloperOptions> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDeveloperOptionsMethod(), responseObserver);
    }

    /**
     */
    public void setDeveloperOptions(cloud.lazycat.apis.sys.Devopt.DeveloperOptions request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetDeveloperOptionsMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetDeveloperOptionsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.Devopt.DeveloperOptions>(
                  this, METHODID_GET_DEVELOPER_OPTIONS)))
          .addMethod(
            getSetDeveloperOptionsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Devopt.DeveloperOptions,
                com.google.protobuf.Empty>(
                  this, METHODID_SET_DEVELOPER_OPTIONS)))
          .build();
    }
  }

  /**
   */
  public static final class DevOptServiceStub extends io.grpc.stub.AbstractAsyncStub<DevOptServiceStub> {
    private DevOptServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevOptServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevOptServiceStub(channel, callOptions);
    }

    /**
     */
    public void getDeveloperOptions(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Devopt.DeveloperOptions> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDeveloperOptionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void setDeveloperOptions(cloud.lazycat.apis.sys.Devopt.DeveloperOptions request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetDeveloperOptionsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class DevOptServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<DevOptServiceBlockingStub> {
    private DevOptServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevOptServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevOptServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.sys.Devopt.DeveloperOptions getDeveloperOptions(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDeveloperOptionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty setDeveloperOptions(cloud.lazycat.apis.sys.Devopt.DeveloperOptions request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetDeveloperOptionsMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class DevOptServiceFutureStub extends io.grpc.stub.AbstractFutureStub<DevOptServiceFutureStub> {
    private DevOptServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DevOptServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DevOptServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Devopt.DeveloperOptions> getDeveloperOptions(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDeveloperOptionsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> setDeveloperOptions(
        cloud.lazycat.apis.sys.Devopt.DeveloperOptions request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetDeveloperOptionsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_DEVELOPER_OPTIONS = 0;
  private static final int METHODID_SET_DEVELOPER_OPTIONS = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DevOptServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DevOptServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_DEVELOPER_OPTIONS:
          serviceImpl.getDeveloperOptions((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Devopt.DeveloperOptions>) responseObserver);
          break;
        case METHODID_SET_DEVELOPER_OPTIONS:
          serviceImpl.setDeveloperOptions((cloud.lazycat.apis.sys.Devopt.DeveloperOptions) request,
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

  private static abstract class DevOptServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DevOptServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.Devopt.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("DevOptService");
    }
  }

  private static final class DevOptServiceFileDescriptorSupplier
      extends DevOptServiceBaseDescriptorSupplier {
    DevOptServiceFileDescriptorSupplier() {}
  }

  private static final class DevOptServiceMethodDescriptorSupplier
      extends DevOptServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DevOptServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (DevOptServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DevOptServiceFileDescriptorSupplier())
              .addMethod(getGetDeveloperOptionsMethod())
              .addMethod(getSetDeveloperOptionsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
