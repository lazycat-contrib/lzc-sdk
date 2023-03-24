package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/dirmon.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class DirMonitorGrpc {

  private DirMonitorGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.DirMonitor";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest,
      cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse> getFindNewMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindNew",
      requestType = cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest.class,
      responseType = cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest,
      cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse> getFindNewMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest, cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse> getFindNewMethod;
    if ((getFindNewMethod = DirMonitorGrpc.getFindNewMethod) == null) {
      synchronized (DirMonitorGrpc.class) {
        if ((getFindNewMethod = DirMonitorGrpc.getFindNewMethod) == null) {
          DirMonitorGrpc.getFindNewMethod = getFindNewMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest, cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindNew"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse.getDefaultInstance()))
              .setSchemaDescriptor(new DirMonitorMethodDescriptorSupplier("FindNew"))
              .build();
        }
      }
    }
    return getFindNewMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DirMonitorStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DirMonitorStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DirMonitorStub>() {
        @java.lang.Override
        public DirMonitorStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DirMonitorStub(channel, callOptions);
        }
      };
    return DirMonitorStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DirMonitorBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DirMonitorBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DirMonitorBlockingStub>() {
        @java.lang.Override
        public DirMonitorBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DirMonitorBlockingStub(channel, callOptions);
        }
      };
    return DirMonitorBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DirMonitorFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DirMonitorFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DirMonitorFutureStub>() {
        @java.lang.Override
        public DirMonitorFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DirMonitorFutureStub(channel, callOptions);
        }
      };
    return DirMonitorFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class DirMonitorImplBase implements io.grpc.BindableService {

    /**
     */
    public void findNew(cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindNewMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getFindNewMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest,
                cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse>(
                  this, METHODID_FIND_NEW)))
          .build();
    }
  }

  /**
   */
  public static final class DirMonitorStub extends io.grpc.stub.AbstractAsyncStub<DirMonitorStub> {
    private DirMonitorStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DirMonitorStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DirMonitorStub(channel, callOptions);
    }

    /**
     */
    public void findNew(cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindNewMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class DirMonitorBlockingStub extends io.grpc.stub.AbstractBlockingStub<DirMonitorBlockingStub> {
    private DirMonitorBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DirMonitorBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DirMonitorBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse findNew(cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindNewMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class DirMonitorFutureStub extends io.grpc.stub.AbstractFutureStub<DirMonitorFutureStub> {
    private DirMonitorFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DirMonitorFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DirMonitorFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse> findNew(
        cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindNewMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_FIND_NEW = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DirMonitorImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DirMonitorImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_FIND_NEW:
          serviceImpl.findNew((cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Dirmon.DirMonitorFindNewResponse>) responseObserver);
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

  private static abstract class DirMonitorBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DirMonitorBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.Dirmon.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("DirMonitor");
    }
  }

  private static final class DirMonitorFileDescriptorSupplier
      extends DirMonitorBaseDescriptorSupplier {
    DirMonitorFileDescriptorSupplier() {}
  }

  private static final class DirMonitorMethodDescriptorSupplier
      extends DirMonitorBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DirMonitorMethodDescriptorSupplier(String methodName) {
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
      synchronized (DirMonitorGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DirMonitorFileDescriptorSupplier())
              .addMethod(getFindNewMethod())
              .build();
        }
      }
    }
    return result;
  }
}
