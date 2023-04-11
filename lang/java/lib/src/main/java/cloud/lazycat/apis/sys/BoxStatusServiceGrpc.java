package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/box-status.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class BoxStatusServiceGrpc {

  private BoxStatusServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.BoxStatusService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus> getGetAllMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAll",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus> getGetAllMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus> getGetAllMethod;
    if ((getGetAllMethod = BoxStatusServiceGrpc.getGetAllMethod) == null) {
      synchronized (BoxStatusServiceGrpc.class) {
        if ((getGetAllMethod = BoxStatusServiceGrpc.getGetAllMethod) == null) {
          BoxStatusServiceGrpc.getGetAllMethod = getGetAllMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAll"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus.getDefaultInstance()))
              .setSchemaDescriptor(new BoxStatusServiceMethodDescriptorSupplier("GetAll"))
              .build();
        }
      }
    }
    return getGetAllMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks,
      com.google.protobuf.Empty> getFormatMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Format",
      requestType = cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks,
      com.google.protobuf.Empty> getFormatMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks, com.google.protobuf.Empty> getFormatMethod;
    if ((getFormatMethod = BoxStatusServiceGrpc.getFormatMethod) == null) {
      synchronized (BoxStatusServiceGrpc.class) {
        if ((getFormatMethod = BoxStatusServiceGrpc.getFormatMethod) == null) {
          BoxStatusServiceGrpc.getFormatMethod = getFormatMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Format"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BoxStatusServiceMethodDescriptorSupplier("Format"))
              .build();
        }
      }
    }
    return getFormatMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static BoxStatusServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BoxStatusServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BoxStatusServiceStub>() {
        @java.lang.Override
        public BoxStatusServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BoxStatusServiceStub(channel, callOptions);
        }
      };
    return BoxStatusServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static BoxStatusServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BoxStatusServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BoxStatusServiceBlockingStub>() {
        @java.lang.Override
        public BoxStatusServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BoxStatusServiceBlockingStub(channel, callOptions);
        }
      };
    return BoxStatusServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static BoxStatusServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BoxStatusServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BoxStatusServiceFutureStub>() {
        @java.lang.Override
        public BoxStatusServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BoxStatusServiceFutureStub(channel, callOptions);
        }
      };
    return BoxStatusServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class BoxStatusServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void getAll(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAllMethod(), responseObserver);
    }

    /**
     */
    public void format(cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFormatMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetAllMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus>(
                  this, METHODID_GET_ALL)))
          .addMethod(
            getFormatMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks,
                com.google.protobuf.Empty>(
                  this, METHODID_FORMAT)))
          .build();
    }
  }

  /**
   */
  public static final class BoxStatusServiceStub extends io.grpc.stub.AbstractAsyncStub<BoxStatusServiceStub> {
    private BoxStatusServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BoxStatusServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BoxStatusServiceStub(channel, callOptions);
    }

    /**
     */
    public void getAll(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAllMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void format(cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFormatMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class BoxStatusServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<BoxStatusServiceBlockingStub> {
    private BoxStatusServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BoxStatusServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BoxStatusServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus getAll(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAllMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty format(cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFormatMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class BoxStatusServiceFutureStub extends io.grpc.stub.AbstractFutureStub<BoxStatusServiceFutureStub> {
    private BoxStatusServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BoxStatusServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BoxStatusServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus> getAll(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAllMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> format(
        cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFormatMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_ALL = 0;
  private static final int METHODID_FORMAT = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final BoxStatusServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(BoxStatusServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_ALL:
          serviceImpl.getAll((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.BoxStatusOuterClass.BoxStatus>) responseObserver);
          break;
        case METHODID_FORMAT:
          serviceImpl.format((cloud.lazycat.apis.sys.BoxStatusOuterClass.Disks) request,
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

  private static abstract class BoxStatusServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    BoxStatusServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.BoxStatusOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("BoxStatusService");
    }
  }

  private static final class BoxStatusServiceFileDescriptorSupplier
      extends BoxStatusServiceBaseDescriptorSupplier {
    BoxStatusServiceFileDescriptorSupplier() {}
  }

  private static final class BoxStatusServiceMethodDescriptorSupplier
      extends BoxStatusServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    BoxStatusServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (BoxStatusServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new BoxStatusServiceFileDescriptorSupplier())
              .addMethod(getGetAllMethod())
              .addMethod(getFormatMethod())
              .build();
        }
      }
    }
    return result;
  }
}
