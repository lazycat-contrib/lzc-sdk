package lzc.lanforward;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: lanforward/lanforward.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class LanForwardServiceGrpc {

  private LanForwardServiceGrpc() {}

  public static final String SERVICE_NAME = "lzc.lanforward.LanForwardService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<lzc.lanforward.Lanforward.ForwardRequest,
      lzc.lanforward.Lanforward.ForwardResponse> getForwardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Forward",
      requestType = lzc.lanforward.Lanforward.ForwardRequest.class,
      responseType = lzc.lanforward.Lanforward.ForwardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<lzc.lanforward.Lanforward.ForwardRequest,
      lzc.lanforward.Lanforward.ForwardResponse> getForwardMethod() {
    io.grpc.MethodDescriptor<lzc.lanforward.Lanforward.ForwardRequest, lzc.lanforward.Lanforward.ForwardResponse> getForwardMethod;
    if ((getForwardMethod = LanForwardServiceGrpc.getForwardMethod) == null) {
      synchronized (LanForwardServiceGrpc.class) {
        if ((getForwardMethod = LanForwardServiceGrpc.getForwardMethod) == null) {
          LanForwardServiceGrpc.getForwardMethod = getForwardMethod =
              io.grpc.MethodDescriptor.<lzc.lanforward.Lanforward.ForwardRequest, lzc.lanforward.Lanforward.ForwardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Forward"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.lanforward.Lanforward.ForwardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.lanforward.Lanforward.ForwardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LanForwardServiceMethodDescriptorSupplier("Forward"))
              .build();
        }
      }
    }
    return getForwardMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static LanForwardServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LanForwardServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LanForwardServiceStub>() {
        @java.lang.Override
        public LanForwardServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LanForwardServiceStub(channel, callOptions);
        }
      };
    return LanForwardServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static LanForwardServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LanForwardServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LanForwardServiceBlockingStub>() {
        @java.lang.Override
        public LanForwardServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LanForwardServiceBlockingStub(channel, callOptions);
        }
      };
    return LanForwardServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static LanForwardServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LanForwardServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LanForwardServiceFutureStub>() {
        @java.lang.Override
        public LanForwardServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LanForwardServiceFutureStub(channel, callOptions);
        }
      };
    return LanForwardServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class LanForwardServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     *将特定service的特定端口上的http服务转发到局域网可直接访问的地址上，
     *返回被代理后的url. (实际使用时，需要使用此url拼接最终服务)
     *请求者需要主动关闭此请求以便停止相关服务
     * </pre>
     */
    public void forward(lzc.lanforward.Lanforward.ForwardRequest request,
        io.grpc.stub.StreamObserver<lzc.lanforward.Lanforward.ForwardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getForwardMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getForwardMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                lzc.lanforward.Lanforward.ForwardRequest,
                lzc.lanforward.Lanforward.ForwardResponse>(
                  this, METHODID_FORWARD)))
          .build();
    }
  }

  /**
   */
  public static final class LanForwardServiceStub extends io.grpc.stub.AbstractAsyncStub<LanForwardServiceStub> {
    private LanForwardServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LanForwardServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LanForwardServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     *将特定service的特定端口上的http服务转发到局域网可直接访问的地址上，
     *返回被代理后的url. (实际使用时，需要使用此url拼接最终服务)
     *请求者需要主动关闭此请求以便停止相关服务
     * </pre>
     */
    public void forward(lzc.lanforward.Lanforward.ForwardRequest request,
        io.grpc.stub.StreamObserver<lzc.lanforward.Lanforward.ForwardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getForwardMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class LanForwardServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<LanForwardServiceBlockingStub> {
    private LanForwardServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LanForwardServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LanForwardServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     *将特定service的特定端口上的http服务转发到局域网可直接访问的地址上，
     *返回被代理后的url. (实际使用时，需要使用此url拼接最终服务)
     *请求者需要主动关闭此请求以便停止相关服务
     * </pre>
     */
    public java.util.Iterator<lzc.lanforward.Lanforward.ForwardResponse> forward(
        lzc.lanforward.Lanforward.ForwardRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getForwardMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class LanForwardServiceFutureStub extends io.grpc.stub.AbstractFutureStub<LanForwardServiceFutureStub> {
    private LanForwardServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LanForwardServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LanForwardServiceFutureStub(channel, callOptions);
    }
  }

  private static final int METHODID_FORWARD = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final LanForwardServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(LanForwardServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_FORWARD:
          serviceImpl.forward((lzc.lanforward.Lanforward.ForwardRequest) request,
              (io.grpc.stub.StreamObserver<lzc.lanforward.Lanforward.ForwardResponse>) responseObserver);
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

  private static abstract class LanForwardServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    LanForwardServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return lzc.lanforward.Lanforward.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("LanForwardService");
    }
  }

  private static final class LanForwardServiceFileDescriptorSupplier
      extends LanForwardServiceBaseDescriptorSupplier {
    LanForwardServiceFileDescriptorSupplier() {}
  }

  private static final class LanForwardServiceMethodDescriptorSupplier
      extends LanForwardServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    LanForwardServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (LanForwardServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new LanForwardServiceFileDescriptorSupplier())
              .addMethod(getForwardMethod())
              .build();
        }
      }
    }
    return result;
  }
}
