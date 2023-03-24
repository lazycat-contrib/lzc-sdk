package lzc.dlna;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * 目前支持搜索DLNA的Render设备，并投送媒体文件
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: dlna/dlna.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class RemoteMediaPlayerServiceGrpc {

  private RemoteMediaPlayerServiceGrpc() {}

  public static final String SERVICE_NAME = "lzc.dlna.RemoteMediaPlayerService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      lzc.dlna.Dlna.ScanRMPResponse> getScanRMPMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ScanRMP",
      requestType = com.google.protobuf.Empty.class,
      responseType = lzc.dlna.Dlna.ScanRMPResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      lzc.dlna.Dlna.ScanRMPResponse> getScanRMPMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, lzc.dlna.Dlna.ScanRMPResponse> getScanRMPMethod;
    if ((getScanRMPMethod = RemoteMediaPlayerServiceGrpc.getScanRMPMethod) == null) {
      synchronized (RemoteMediaPlayerServiceGrpc.class) {
        if ((getScanRMPMethod = RemoteMediaPlayerServiceGrpc.getScanRMPMethod) == null) {
          RemoteMediaPlayerServiceGrpc.getScanRMPMethod = getScanRMPMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, lzc.dlna.Dlna.ScanRMPResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ScanRMP"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.dlna.Dlna.ScanRMPResponse.getDefaultInstance()))
              .setSchemaDescriptor(new RemoteMediaPlayerServiceMethodDescriptorSupplier("ScanRMP"))
              .build();
        }
      }
    }
    return getScanRMPMethod;
  }

  private static volatile io.grpc.MethodDescriptor<lzc.dlna.Dlna.SubscribeRequest,
      lzc.dlna.Dlna.RMPStatus> getSubscribeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Subscribe",
      requestType = lzc.dlna.Dlna.SubscribeRequest.class,
      responseType = lzc.dlna.Dlna.RMPStatus.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<lzc.dlna.Dlna.SubscribeRequest,
      lzc.dlna.Dlna.RMPStatus> getSubscribeMethod() {
    io.grpc.MethodDescriptor<lzc.dlna.Dlna.SubscribeRequest, lzc.dlna.Dlna.RMPStatus> getSubscribeMethod;
    if ((getSubscribeMethod = RemoteMediaPlayerServiceGrpc.getSubscribeMethod) == null) {
      synchronized (RemoteMediaPlayerServiceGrpc.class) {
        if ((getSubscribeMethod = RemoteMediaPlayerServiceGrpc.getSubscribeMethod) == null) {
          RemoteMediaPlayerServiceGrpc.getSubscribeMethod = getSubscribeMethod =
              io.grpc.MethodDescriptor.<lzc.dlna.Dlna.SubscribeRequest, lzc.dlna.Dlna.RMPStatus>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Subscribe"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.dlna.Dlna.SubscribeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.dlna.Dlna.RMPStatus.getDefaultInstance()))
              .setSchemaDescriptor(new RemoteMediaPlayerServiceMethodDescriptorSupplier("Subscribe"))
              .build();
        }
      }
    }
    return getSubscribeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<lzc.dlna.Dlna.DoActionRequest,
      com.google.protobuf.Empty> getDoActionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DoAction",
      requestType = lzc.dlna.Dlna.DoActionRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<lzc.dlna.Dlna.DoActionRequest,
      com.google.protobuf.Empty> getDoActionMethod() {
    io.grpc.MethodDescriptor<lzc.dlna.Dlna.DoActionRequest, com.google.protobuf.Empty> getDoActionMethod;
    if ((getDoActionMethod = RemoteMediaPlayerServiceGrpc.getDoActionMethod) == null) {
      synchronized (RemoteMediaPlayerServiceGrpc.class) {
        if ((getDoActionMethod = RemoteMediaPlayerServiceGrpc.getDoActionMethod) == null) {
          RemoteMediaPlayerServiceGrpc.getDoActionMethod = getDoActionMethod =
              io.grpc.MethodDescriptor.<lzc.dlna.Dlna.DoActionRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DoAction"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.dlna.Dlna.DoActionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new RemoteMediaPlayerServiceMethodDescriptorSupplier("DoAction"))
              .build();
        }
      }
    }
    return getDoActionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<lzc.dlna.Dlna.GetPositionInfoRequest,
      lzc.dlna.Dlna.GetPositionInfoResponse> getGetPositionInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPositionInfo",
      requestType = lzc.dlna.Dlna.GetPositionInfoRequest.class,
      responseType = lzc.dlna.Dlna.GetPositionInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<lzc.dlna.Dlna.GetPositionInfoRequest,
      lzc.dlna.Dlna.GetPositionInfoResponse> getGetPositionInfoMethod() {
    io.grpc.MethodDescriptor<lzc.dlna.Dlna.GetPositionInfoRequest, lzc.dlna.Dlna.GetPositionInfoResponse> getGetPositionInfoMethod;
    if ((getGetPositionInfoMethod = RemoteMediaPlayerServiceGrpc.getGetPositionInfoMethod) == null) {
      synchronized (RemoteMediaPlayerServiceGrpc.class) {
        if ((getGetPositionInfoMethod = RemoteMediaPlayerServiceGrpc.getGetPositionInfoMethod) == null) {
          RemoteMediaPlayerServiceGrpc.getGetPositionInfoMethod = getGetPositionInfoMethod =
              io.grpc.MethodDescriptor.<lzc.dlna.Dlna.GetPositionInfoRequest, lzc.dlna.Dlna.GetPositionInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPositionInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.dlna.Dlna.GetPositionInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  lzc.dlna.Dlna.GetPositionInfoResponse.getDefaultInstance()))
              .setSchemaDescriptor(new RemoteMediaPlayerServiceMethodDescriptorSupplier("GetPositionInfo"))
              .build();
        }
      }
    }
    return getGetPositionInfoMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static RemoteMediaPlayerServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<RemoteMediaPlayerServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<RemoteMediaPlayerServiceStub>() {
        @java.lang.Override
        public RemoteMediaPlayerServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new RemoteMediaPlayerServiceStub(channel, callOptions);
        }
      };
    return RemoteMediaPlayerServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static RemoteMediaPlayerServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<RemoteMediaPlayerServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<RemoteMediaPlayerServiceBlockingStub>() {
        @java.lang.Override
        public RemoteMediaPlayerServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new RemoteMediaPlayerServiceBlockingStub(channel, callOptions);
        }
      };
    return RemoteMediaPlayerServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static RemoteMediaPlayerServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<RemoteMediaPlayerServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<RemoteMediaPlayerServiceFutureStub>() {
        @java.lang.Override
        public RemoteMediaPlayerServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new RemoteMediaPlayerServiceFutureStub(channel, callOptions);
        }
      };
    return RemoteMediaPlayerServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * 目前支持搜索DLNA的Render设备，并投送媒体文件
   * </pre>
   */
  public static abstract class RemoteMediaPlayerServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public void scanRMP(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<lzc.dlna.Dlna.ScanRMPResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getScanRMPMethod(), responseObserver);
    }

    /**
     * <pre>
     * 立刻返回选择的当前RMP状态，并在状态变动时重新发送
     * </pre>
     */
    public void subscribe(lzc.dlna.Dlna.SubscribeRequest request,
        io.grpc.stub.StreamObserver<lzc.dlna.Dlna.RMPStatus> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubscribeMethod(), responseObserver);
    }

    /**
     */
    public void doAction(lzc.dlna.Dlna.DoActionRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDoActionMethod(), responseObserver);
    }

    /**
     */
    public void getPositionInfo(lzc.dlna.Dlna.GetPositionInfoRequest request,
        io.grpc.stub.StreamObserver<lzc.dlna.Dlna.GetPositionInfoResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPositionInfoMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getScanRMPMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                lzc.dlna.Dlna.ScanRMPResponse>(
                  this, METHODID_SCAN_RMP)))
          .addMethod(
            getSubscribeMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                lzc.dlna.Dlna.SubscribeRequest,
                lzc.dlna.Dlna.RMPStatus>(
                  this, METHODID_SUBSCRIBE)))
          .addMethod(
            getDoActionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                lzc.dlna.Dlna.DoActionRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_DO_ACTION)))
          .addMethod(
            getGetPositionInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                lzc.dlna.Dlna.GetPositionInfoRequest,
                lzc.dlna.Dlna.GetPositionInfoResponse>(
                  this, METHODID_GET_POSITION_INFO)))
          .build();
    }
  }

  /**
   * <pre>
   * 目前支持搜索DLNA的Render设备，并投送媒体文件
   * </pre>
   */
  public static final class RemoteMediaPlayerServiceStub extends io.grpc.stub.AbstractAsyncStub<RemoteMediaPlayerServiceStub> {
    private RemoteMediaPlayerServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected RemoteMediaPlayerServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new RemoteMediaPlayerServiceStub(channel, callOptions);
    }

    /**
     */
    public void scanRMP(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<lzc.dlna.Dlna.ScanRMPResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getScanRMPMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 立刻返回选择的当前RMP状态，并在状态变动时重新发送
     * </pre>
     */
    public void subscribe(lzc.dlna.Dlna.SubscribeRequest request,
        io.grpc.stub.StreamObserver<lzc.dlna.Dlna.RMPStatus> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getSubscribeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void doAction(lzc.dlna.Dlna.DoActionRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDoActionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPositionInfo(lzc.dlna.Dlna.GetPositionInfoRequest request,
        io.grpc.stub.StreamObserver<lzc.dlna.Dlna.GetPositionInfoResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPositionInfoMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * 目前支持搜索DLNA的Render设备，并投送媒体文件
   * </pre>
   */
  public static final class RemoteMediaPlayerServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<RemoteMediaPlayerServiceBlockingStub> {
    private RemoteMediaPlayerServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected RemoteMediaPlayerServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new RemoteMediaPlayerServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public lzc.dlna.Dlna.ScanRMPResponse scanRMP(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getScanRMPMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 立刻返回选择的当前RMP状态，并在状态变动时重新发送
     * </pre>
     */
    public java.util.Iterator<lzc.dlna.Dlna.RMPStatus> subscribe(
        lzc.dlna.Dlna.SubscribeRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getSubscribeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty doAction(lzc.dlna.Dlna.DoActionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDoActionMethod(), getCallOptions(), request);
    }

    /**
     */
    public lzc.dlna.Dlna.GetPositionInfoResponse getPositionInfo(lzc.dlna.Dlna.GetPositionInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPositionInfoMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * 目前支持搜索DLNA的Render设备，并投送媒体文件
   * </pre>
   */
  public static final class RemoteMediaPlayerServiceFutureStub extends io.grpc.stub.AbstractFutureStub<RemoteMediaPlayerServiceFutureStub> {
    private RemoteMediaPlayerServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected RemoteMediaPlayerServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new RemoteMediaPlayerServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<lzc.dlna.Dlna.ScanRMPResponse> scanRMP(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getScanRMPMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> doAction(
        lzc.dlna.Dlna.DoActionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDoActionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<lzc.dlna.Dlna.GetPositionInfoResponse> getPositionInfo(
        lzc.dlna.Dlna.GetPositionInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPositionInfoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SCAN_RMP = 0;
  private static final int METHODID_SUBSCRIBE = 1;
  private static final int METHODID_DO_ACTION = 2;
  private static final int METHODID_GET_POSITION_INFO = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final RemoteMediaPlayerServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(RemoteMediaPlayerServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SCAN_RMP:
          serviceImpl.scanRMP((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<lzc.dlna.Dlna.ScanRMPResponse>) responseObserver);
          break;
        case METHODID_SUBSCRIBE:
          serviceImpl.subscribe((lzc.dlna.Dlna.SubscribeRequest) request,
              (io.grpc.stub.StreamObserver<lzc.dlna.Dlna.RMPStatus>) responseObserver);
          break;
        case METHODID_DO_ACTION:
          serviceImpl.doAction((lzc.dlna.Dlna.DoActionRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GET_POSITION_INFO:
          serviceImpl.getPositionInfo((lzc.dlna.Dlna.GetPositionInfoRequest) request,
              (io.grpc.stub.StreamObserver<lzc.dlna.Dlna.GetPositionInfoResponse>) responseObserver);
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

  private static abstract class RemoteMediaPlayerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    RemoteMediaPlayerServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return lzc.dlna.Dlna.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("RemoteMediaPlayerService");
    }
  }

  private static final class RemoteMediaPlayerServiceFileDescriptorSupplier
      extends RemoteMediaPlayerServiceBaseDescriptorSupplier {
    RemoteMediaPlayerServiceFileDescriptorSupplier() {}
  }

  private static final class RemoteMediaPlayerServiceMethodDescriptorSupplier
      extends RemoteMediaPlayerServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    RemoteMediaPlayerServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (RemoteMediaPlayerServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new RemoteMediaPlayerServiceFileDescriptorSupplier())
              .addMethod(getScanRMPMethod())
              .addMethod(getSubscribeMethod())
              .addMethod(getDoActionMethod())
              .addMethod(getGetPositionInfoMethod())
              .build();
        }
      }
    }
    return result;
  }
}
