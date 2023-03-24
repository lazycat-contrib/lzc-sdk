package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/OS_daemon.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class OsDaemonServiceGrpc {

  private OsDaemonServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.OsDaemonService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSDaemon.OptionKey,
      cloud.lazycat.apis.sys.OSDaemon.OptionValue> getGetOptionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetOption",
      requestType = cloud.lazycat.apis.sys.OSDaemon.OptionKey.class,
      responseType = cloud.lazycat.apis.sys.OSDaemon.OptionValue.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSDaemon.OptionKey,
      cloud.lazycat.apis.sys.OSDaemon.OptionValue> getGetOptionMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSDaemon.OptionKey, cloud.lazycat.apis.sys.OSDaemon.OptionValue> getGetOptionMethod;
    if ((getGetOptionMethod = OsDaemonServiceGrpc.getGetOptionMethod) == null) {
      synchronized (OsDaemonServiceGrpc.class) {
        if ((getGetOptionMethod = OsDaemonServiceGrpc.getGetOptionMethod) == null) {
          OsDaemonServiceGrpc.getGetOptionMethod = getGetOptionMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSDaemon.OptionKey, cloud.lazycat.apis.sys.OSDaemon.OptionValue>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetOption"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSDaemon.OptionKey.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSDaemon.OptionValue.getDefaultInstance()))
              .setSchemaDescriptor(new OsDaemonServiceMethodDescriptorSupplier("GetOption"))
              .build();
        }
      }
    }
    return getGetOptionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue,
      com.google.protobuf.Empty> getSetOptionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetOption",
      requestType = cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue,
      com.google.protobuf.Empty> getSetOptionMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue, com.google.protobuf.Empty> getSetOptionMethod;
    if ((getSetOptionMethod = OsDaemonServiceGrpc.getSetOptionMethod) == null) {
      synchronized (OsDaemonServiceGrpc.class) {
        if ((getSetOptionMethod = OsDaemonServiceGrpc.getSetOptionMethod) == null) {
          OsDaemonServiceGrpc.getSetOptionMethod = getSetOptionMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetOption"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OsDaemonServiceMethodDescriptorSupplier("SetOption"))
              .build();
        }
      }
    }
    return getSetOptionMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static OsDaemonServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OsDaemonServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OsDaemonServiceStub>() {
        @java.lang.Override
        public OsDaemonServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OsDaemonServiceStub(channel, callOptions);
        }
      };
    return OsDaemonServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static OsDaemonServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OsDaemonServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OsDaemonServiceBlockingStub>() {
        @java.lang.Override
        public OsDaemonServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OsDaemonServiceBlockingStub(channel, callOptions);
        }
      };
    return OsDaemonServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static OsDaemonServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OsDaemonServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OsDaemonServiceFutureStub>() {
        @java.lang.Override
        public OsDaemonServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OsDaemonServiceFutureStub(channel, callOptions);
        }
      };
    return OsDaemonServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class OsDaemonServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void getOption(cloud.lazycat.apis.sys.OSDaemon.OptionKey request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSDaemon.OptionValue> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetOptionMethod(), responseObserver);
    }

    /**
     */
    public void setOption(cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetOptionMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetOptionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSDaemon.OptionKey,
                cloud.lazycat.apis.sys.OSDaemon.OptionValue>(
                  this, METHODID_GET_OPTION)))
          .addMethod(
            getSetOptionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue,
                com.google.protobuf.Empty>(
                  this, METHODID_SET_OPTION)))
          .build();
    }
  }

  /**
   */
  public static final class OsDaemonServiceStub extends io.grpc.stub.AbstractAsyncStub<OsDaemonServiceStub> {
    private OsDaemonServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OsDaemonServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OsDaemonServiceStub(channel, callOptions);
    }

    /**
     */
    public void getOption(cloud.lazycat.apis.sys.OSDaemon.OptionKey request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSDaemon.OptionValue> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetOptionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void setOption(cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetOptionMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class OsDaemonServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<OsDaemonServiceBlockingStub> {
    private OsDaemonServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OsDaemonServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OsDaemonServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.sys.OSDaemon.OptionValue getOption(cloud.lazycat.apis.sys.OSDaemon.OptionKey request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetOptionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty setOption(cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetOptionMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class OsDaemonServiceFutureStub extends io.grpc.stub.AbstractFutureStub<OsDaemonServiceFutureStub> {
    private OsDaemonServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OsDaemonServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OsDaemonServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSDaemon.OptionValue> getOption(
        cloud.lazycat.apis.sys.OSDaemon.OptionKey request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetOptionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> setOption(
        cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetOptionMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_OPTION = 0;
  private static final int METHODID_SET_OPTION = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final OsDaemonServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(OsDaemonServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_OPTION:
          serviceImpl.getOption((cloud.lazycat.apis.sys.OSDaemon.OptionKey) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSDaemon.OptionValue>) responseObserver);
          break;
        case METHODID_SET_OPTION:
          serviceImpl.setOption((cloud.lazycat.apis.sys.OSDaemon.OptionKeyValue) request,
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

  private static abstract class OsDaemonServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    OsDaemonServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.OSDaemon.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("OsDaemonService");
    }
  }

  private static final class OsDaemonServiceFileDescriptorSupplier
      extends OsDaemonServiceBaseDescriptorSupplier {
    OsDaemonServiceFileDescriptorSupplier() {}
  }

  private static final class OsDaemonServiceMethodDescriptorSupplier
      extends OsDaemonServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    OsDaemonServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (OsDaemonServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new OsDaemonServiceFileDescriptorSupplier())
              .addMethod(getGetOptionMethod())
              .addMethod(getSetOptionMethod())
              .build();
        }
      }
    }
    return result;
  }
}
