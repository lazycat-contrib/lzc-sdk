package lzc.ssdp;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: ssdp/ssdp.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class SSDPServiceGrpc {

  private SSDPServiceGrpc() {}

  public static final String SERVICE_NAME = "lzc.ssdp.SSDPService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<lzc.ssdp.Ssdp.SearchRequest,
      lzc.ssdp.Ssdp.Service> getSearchMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Search",
      requestType = lzc.ssdp.Ssdp.SearchRequest.class,
      responseType = lzc.ssdp.Ssdp.Service.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<lzc.ssdp.Ssdp.SearchRequest,
      lzc.ssdp.Ssdp.Service> getSearchMethod() {
    io.grpc.MethodDescriptor<lzc.ssdp.Ssdp.SearchRequest, lzc.ssdp.Ssdp.Service> getSearchMethod;
    if ((getSearchMethod = SSDPServiceGrpc.getSearchMethod) == null) {
      synchronized (SSDPServiceGrpc.class) {
        if ((getSearchMethod = SSDPServiceGrpc.getSearchMethod) == null) {
          SSDPServiceGrpc.getSearchMethod = getSearchMethod =
              io.grpc.MethodDescriptor.<lzc.ssdp.Ssdp.SearchRequest, lzc.ssdp.Ssdp.Service>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Search"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.ssdp.Ssdp.SearchRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.ssdp.Ssdp.Service.getDefaultInstance()))
              .setSchemaDescriptor(new SSDPServiceMethodDescriptorSupplier("Search"))
              .build();
        }
      }
    }
    return getSearchMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SSDPServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SSDPServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SSDPServiceStub>() {
        @java.lang.Override
        public SSDPServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SSDPServiceStub(channel, callOptions);
        }
      };
    return SSDPServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SSDPServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SSDPServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SSDPServiceBlockingStub>() {
        @java.lang.Override
        public SSDPServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SSDPServiceBlockingStub(channel, callOptions);
        }
      };
    return SSDPServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SSDPServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SSDPServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SSDPServiceFutureStub>() {
        @java.lang.Override
        public SSDPServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SSDPServiceFutureStub(channel, callOptions);
        }
      };
    return SSDPServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class SSDPServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void search(lzc.ssdp.Ssdp.SearchRequest request,
        io.grpc.stub.StreamObserver<lzc.ssdp.Ssdp.Service> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSearchMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSearchMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                lzc.ssdp.Ssdp.SearchRequest,
                lzc.ssdp.Ssdp.Service>(
                  this, METHODID_SEARCH)))
          .build();
    }
  }

  /**
   */
  public static final class SSDPServiceStub extends io.grpc.stub.AbstractAsyncStub<SSDPServiceStub> {
    private SSDPServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SSDPServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SSDPServiceStub(channel, callOptions);
    }

    /**
     */
    public void search(lzc.ssdp.Ssdp.SearchRequest request,
        io.grpc.stub.StreamObserver<lzc.ssdp.Ssdp.Service> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getSearchMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class SSDPServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<SSDPServiceBlockingStub> {
    private SSDPServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SSDPServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SSDPServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public java.util.Iterator<lzc.ssdp.Ssdp.Service> search(
        lzc.ssdp.Ssdp.SearchRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getSearchMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class SSDPServiceFutureStub extends io.grpc.stub.AbstractFutureStub<SSDPServiceFutureStub> {
    private SSDPServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SSDPServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SSDPServiceFutureStub(channel, callOptions);
    }
  }

  private static final int METHODID_SEARCH = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final SSDPServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(SSDPServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SEARCH:
          serviceImpl.search((lzc.ssdp.Ssdp.SearchRequest) request,
              (io.grpc.stub.StreamObserver<lzc.ssdp.Ssdp.Service>) responseObserver);
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

  private static abstract class SSDPServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    SSDPServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return lzc.ssdp.Ssdp.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("SSDPService");
    }
  }

  private static final class SSDPServiceFileDescriptorSupplier
      extends SSDPServiceBaseDescriptorSupplier {
    SSDPServiceFileDescriptorSupplier() {}
  }

  private static final class SSDPServiceMethodDescriptorSupplier
      extends SSDPServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    SSDPServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (SSDPServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new SSDPServiceFileDescriptorSupplier())
              .addMethod(getSearchMethod())
              .build();
        }
      }
    }
    return result;
  }
}
