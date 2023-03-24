package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/btrfs.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class BtrfsUtilGrpc {

  private BtrfsUtilGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.BtrfsUtil";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest,
      com.google.protobuf.Empty> getSubvolCreateMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubvolCreate",
      requestType = cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest,
      com.google.protobuf.Empty> getSubvolCreateMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest, com.google.protobuf.Empty> getSubvolCreateMethod;
    if ((getSubvolCreateMethod = BtrfsUtilGrpc.getSubvolCreateMethod) == null) {
      synchronized (BtrfsUtilGrpc.class) {
        if ((getSubvolCreateMethod = BtrfsUtilGrpc.getSubvolCreateMethod) == null) {
          BtrfsUtilGrpc.getSubvolCreateMethod = getSubvolCreateMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubvolCreate"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BtrfsUtilMethodDescriptorSupplier("SubvolCreate"))
              .build();
        }
      }
    }
    return getSubvolCreateMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest,
      cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse> getSubvolInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubvolInfo",
      requestType = cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest.class,
      responseType = cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest,
      cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse> getSubvolInfoMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest, cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse> getSubvolInfoMethod;
    if ((getSubvolInfoMethod = BtrfsUtilGrpc.getSubvolInfoMethod) == null) {
      synchronized (BtrfsUtilGrpc.class) {
        if ((getSubvolInfoMethod = BtrfsUtilGrpc.getSubvolInfoMethod) == null) {
          BtrfsUtilGrpc.getSubvolInfoMethod = getSubvolInfoMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest, cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubvolInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BtrfsUtilMethodDescriptorSupplier("SubvolInfo"))
              .build();
        }
      }
    }
    return getSubvolInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest,
      cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse> getSubvolFindNewMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubvolFindNew",
      requestType = cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest.class,
      responseType = cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest,
      cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse> getSubvolFindNewMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest, cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse> getSubvolFindNewMethod;
    if ((getSubvolFindNewMethod = BtrfsUtilGrpc.getSubvolFindNewMethod) == null) {
      synchronized (BtrfsUtilGrpc.class) {
        if ((getSubvolFindNewMethod = BtrfsUtilGrpc.getSubvolFindNewMethod) == null) {
          BtrfsUtilGrpc.getSubvolFindNewMethod = getSubvolFindNewMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest, cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubvolFindNew"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BtrfsUtilMethodDescriptorSupplier("SubvolFindNew"))
              .build();
        }
      }
    }
    return getSubvolFindNewMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static BtrfsUtilStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BtrfsUtilStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BtrfsUtilStub>() {
        @java.lang.Override
        public BtrfsUtilStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BtrfsUtilStub(channel, callOptions);
        }
      };
    return BtrfsUtilStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static BtrfsUtilBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BtrfsUtilBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BtrfsUtilBlockingStub>() {
        @java.lang.Override
        public BtrfsUtilBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BtrfsUtilBlockingStub(channel, callOptions);
        }
      };
    return BtrfsUtilBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static BtrfsUtilFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BtrfsUtilFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BtrfsUtilFutureStub>() {
        @java.lang.Override
        public BtrfsUtilFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BtrfsUtilFutureStub(channel, callOptions);
        }
      };
    return BtrfsUtilFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class BtrfsUtilImplBase implements io.grpc.BindableService {

    /**
     */
    public void subvolCreate(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubvolCreateMethod(), responseObserver);
    }

    /**
     */
    public void subvolInfo(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubvolInfoMethod(), responseObserver);
    }

    /**
     */
    public void subvolFindNew(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubvolFindNewMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSubvolCreateMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SUBVOL_CREATE)))
          .addMethod(
            getSubvolInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest,
                cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse>(
                  this, METHODID_SUBVOL_INFO)))
          .addMethod(
            getSubvolFindNewMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest,
                cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse>(
                  this, METHODID_SUBVOL_FIND_NEW)))
          .build();
    }
  }

  /**
   */
  public static final class BtrfsUtilStub extends io.grpc.stub.AbstractAsyncStub<BtrfsUtilStub> {
    private BtrfsUtilStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BtrfsUtilStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BtrfsUtilStub(channel, callOptions);
    }

    /**
     */
    public void subvolCreate(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubvolCreateMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void subvolInfo(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubvolInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void subvolFindNew(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubvolFindNewMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class BtrfsUtilBlockingStub extends io.grpc.stub.AbstractBlockingStub<BtrfsUtilBlockingStub> {
    private BtrfsUtilBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BtrfsUtilBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BtrfsUtilBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.google.protobuf.Empty subvolCreate(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubvolCreateMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse subvolInfo(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubvolInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse subvolFindNew(cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubvolFindNewMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class BtrfsUtilFutureStub extends io.grpc.stub.AbstractFutureStub<BtrfsUtilFutureStub> {
    private BtrfsUtilFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BtrfsUtilFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BtrfsUtilFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> subvolCreate(
        cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubvolCreateMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse> subvolInfo(
        cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubvolInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse> subvolFindNew(
        cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubvolFindNewMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SUBVOL_CREATE = 0;
  private static final int METHODID_SUBVOL_INFO = 1;
  private static final int METHODID_SUBVOL_FIND_NEW = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final BtrfsUtilImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(BtrfsUtilImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SUBVOL_CREATE:
          serviceImpl.subvolCreate((cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolCreateRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SUBVOL_INFO:
          serviceImpl.subvolInfo((cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolInfoResponse>) responseObserver);
          break;
        case METHODID_SUBVOL_FIND_NEW:
          serviceImpl.subvolFindNew((cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Btrfs.BtrfsSubvolFindNewResponse>) responseObserver);
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

  private static abstract class BtrfsUtilBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    BtrfsUtilBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.Btrfs.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("BtrfsUtil");
    }
  }

  private static final class BtrfsUtilFileDescriptorSupplier
      extends BtrfsUtilBaseDescriptorSupplier {
    BtrfsUtilFileDescriptorSupplier() {}
  }

  private static final class BtrfsUtilMethodDescriptorSupplier
      extends BtrfsUtilBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    BtrfsUtilMethodDescriptorSupplier(String methodName) {
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
      synchronized (BtrfsUtilGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new BtrfsUtilFileDescriptorSupplier())
              .addMethod(getSubvolCreateMethod())
              .addMethod(getSubvolInfoMethod())
              .addMethod(getSubvolFindNewMethod())
              .build();
        }
      }
    }
    return result;
  }
}
