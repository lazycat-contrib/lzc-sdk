package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/box.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class BoxServiceGrpc {

  private BoxServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.BoxService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Box.BoxInfo> getQueryInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.common.Box.BoxInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Box.BoxInfo> getQueryInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.common.Box.BoxInfo> getQueryInfoMethod;
    if ((getQueryInfoMethod = BoxServiceGrpc.getQueryInfoMethod) == null) {
      synchronized (BoxServiceGrpc.class) {
        if ((getQueryInfoMethod = BoxServiceGrpc.getQueryInfoMethod) == null) {
          BoxServiceGrpc.getQueryInfoMethod = getQueryInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.common.Box.BoxInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Box.BoxInfo.getDefaultInstance()))
              .setSchemaDescriptor(new BoxServiceMethodDescriptorSupplier("QueryInfo"))
              .build();
        }
      }
    }
    return getQueryInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest,
      com.google.protobuf.Empty> getChangeDisplayNameMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ChangeDisplayName",
      requestType = cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest,
      com.google.protobuf.Empty> getChangeDisplayNameMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest, com.google.protobuf.Empty> getChangeDisplayNameMethod;
    if ((getChangeDisplayNameMethod = BoxServiceGrpc.getChangeDisplayNameMethod) == null) {
      synchronized (BoxServiceGrpc.class) {
        if ((getChangeDisplayNameMethod = BoxServiceGrpc.getChangeDisplayNameMethod) == null) {
          BoxServiceGrpc.getChangeDisplayNameMethod = getChangeDisplayNameMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ChangeDisplayName"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BoxServiceMethodDescriptorSupplier("ChangeDisplayName"))
              .build();
        }
      }
    }
    return getChangeDisplayNameMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Box.ShutdownRequest,
      com.google.protobuf.Empty> getShutdownMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Shutdown",
      requestType = cloud.lazycat.apis.common.Box.ShutdownRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Box.ShutdownRequest,
      com.google.protobuf.Empty> getShutdownMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Box.ShutdownRequest, com.google.protobuf.Empty> getShutdownMethod;
    if ((getShutdownMethod = BoxServiceGrpc.getShutdownMethod) == null) {
      synchronized (BoxServiceGrpc.class) {
        if ((getShutdownMethod = BoxServiceGrpc.getShutdownMethod) == null) {
          BoxServiceGrpc.getShutdownMethod = getShutdownMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Box.ShutdownRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Shutdown"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Box.ShutdownRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new BoxServiceMethodDescriptorSupplier("Shutdown"))
              .build();
        }
      }
    }
    return getShutdownMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static BoxServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BoxServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BoxServiceStub>() {
        @java.lang.Override
        public BoxServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BoxServiceStub(channel, callOptions);
        }
      };
    return BoxServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static BoxServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BoxServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BoxServiceBlockingStub>() {
        @java.lang.Override
        public BoxServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BoxServiceBlockingStub(channel, callOptions);
        }
      };
    return BoxServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static BoxServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BoxServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BoxServiceFutureStub>() {
        @java.lang.Override
        public BoxServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BoxServiceFutureStub(channel, callOptions);
        }
      };
    return BoxServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class BoxServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void queryInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Box.BoxInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryInfoMethod(), responseObserver);
    }

    /**
     */
    public void changeDisplayName(cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getChangeDisplayNameMethod(), responseObserver);
    }

    /**
     */
    public void shutdown(cloud.lazycat.apis.common.Box.ShutdownRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getShutdownMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getQueryInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.common.Box.BoxInfo>(
                  this, METHODID_QUERY_INFO)))
          .addMethod(
            getChangeDisplayNameMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_CHANGE_DISPLAY_NAME)))
          .addMethod(
            getShutdownMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Box.ShutdownRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SHUTDOWN)))
          .build();
    }
  }

  /**
   */
  public static final class BoxServiceStub extends io.grpc.stub.AbstractAsyncStub<BoxServiceStub> {
    private BoxServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BoxServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BoxServiceStub(channel, callOptions);
    }

    /**
     */
    public void queryInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Box.BoxInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void changeDisplayName(cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getChangeDisplayNameMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void shutdown(cloud.lazycat.apis.common.Box.ShutdownRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getShutdownMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class BoxServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<BoxServiceBlockingStub> {
    private BoxServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BoxServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BoxServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.common.Box.BoxInfo queryInfo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryInfoMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty changeDisplayName(cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getChangeDisplayNameMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty shutdown(cloud.lazycat.apis.common.Box.ShutdownRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getShutdownMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class BoxServiceFutureStub extends io.grpc.stub.AbstractFutureStub<BoxServiceFutureStub> {
    private BoxServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BoxServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BoxServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Box.BoxInfo> queryInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryInfoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> changeDisplayName(
        cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getChangeDisplayNameMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> shutdown(
        cloud.lazycat.apis.common.Box.ShutdownRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getShutdownMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_QUERY_INFO = 0;
  private static final int METHODID_CHANGE_DISPLAY_NAME = 1;
  private static final int METHODID_SHUTDOWN = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final BoxServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(BoxServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_QUERY_INFO:
          serviceImpl.queryInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Box.BoxInfo>) responseObserver);
          break;
        case METHODID_CHANGE_DISPLAY_NAME:
          serviceImpl.changeDisplayName((cloud.lazycat.apis.common.Box.ChangeDisplayNameRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SHUTDOWN:
          serviceImpl.shutdown((cloud.lazycat.apis.common.Box.ShutdownRequest) request,
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

  private static abstract class BoxServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    BoxServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.Box.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("BoxService");
    }
  }

  private static final class BoxServiceFileDescriptorSupplier
      extends BoxServiceBaseDescriptorSupplier {
    BoxServiceFileDescriptorSupplier() {}
  }

  private static final class BoxServiceMethodDescriptorSupplier
      extends BoxServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    BoxServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (BoxServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new BoxServiceFileDescriptorSupplier())
              .addMethod(getQueryInfoMethod())
              .addMethod(getChangeDisplayNameMethod())
              .addMethod(getShutdownMethod())
              .build();
        }
      }
    }
    return result;
  }
}
