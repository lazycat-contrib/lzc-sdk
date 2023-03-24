package cloud.lazycat.apis.localdevice;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: localdevice/clipboard.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class ClipboardManagerGrpc {

  private ClipboardManagerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.localdevice.ClipboardManager";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest,
      cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> getReadMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Read",
      requestType = cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest,
      cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> getReadMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest, cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> getReadMethod;
    if ((getReadMethod = ClipboardManagerGrpc.getReadMethod) == null) {
      synchronized (ClipboardManagerGrpc.class) {
        if ((getReadMethod = ClipboardManagerGrpc.getReadMethod) == null) {
          ClipboardManagerGrpc.getReadMethod = getReadMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest, cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Read"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClipboardManagerMethodDescriptorSupplier("Read"))
              .build();
        }
      }
    }
    return getReadMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest,
      cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse> getWriteMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Write",
      requestType = cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest,
      cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse> getWriteMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest, cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse> getWriteMethod;
    if ((getWriteMethod = ClipboardManagerGrpc.getWriteMethod) == null) {
      synchronized (ClipboardManagerGrpc.class) {
        if ((getWriteMethod = ClipboardManagerGrpc.getWriteMethod) == null) {
          ClipboardManagerGrpc.getWriteMethod = getWriteMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest, cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Write"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClipboardManagerMethodDescriptorSupplier("Write"))
              .build();
        }
      }
    }
    return getWriteMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest,
      cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> getWatchMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Watch",
      requestType = cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest,
      cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> getWatchMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest, cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> getWatchMethod;
    if ((getWatchMethod = ClipboardManagerGrpc.getWatchMethod) == null) {
      synchronized (ClipboardManagerGrpc.class) {
        if ((getWatchMethod = ClipboardManagerGrpc.getWatchMethod) == null) {
          ClipboardManagerGrpc.getWatchMethod = getWatchMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest, cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Watch"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClipboardManagerMethodDescriptorSupplier("Watch"))
              .build();
        }
      }
    }
    return getWatchMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ClipboardManagerStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ClipboardManagerStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ClipboardManagerStub>() {
        @java.lang.Override
        public ClipboardManagerStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ClipboardManagerStub(channel, callOptions);
        }
      };
    return ClipboardManagerStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ClipboardManagerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ClipboardManagerBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ClipboardManagerBlockingStub>() {
        @java.lang.Override
        public ClipboardManagerBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ClipboardManagerBlockingStub(channel, callOptions);
        }
      };
    return ClipboardManagerBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ClipboardManagerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ClipboardManagerFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ClipboardManagerFutureStub>() {
        @java.lang.Override
        public ClipboardManagerFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ClipboardManagerFutureStub(channel, callOptions);
        }
      };
    return ClipboardManagerFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class ClipboardManagerImplBase implements io.grpc.BindableService {

    /**
     */
    public void read(cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReadMethod(), responseObserver);
    }

    /**
     */
    public void write(cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getWriteMethod(), responseObserver);
    }

    /**
     */
    public void watch(cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getWatchMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getReadMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest,
                cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse>(
                  this, METHODID_READ)))
          .addMethod(
            getWriteMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest,
                cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse>(
                  this, METHODID_WRITE)))
          .addMethod(
            getWatchMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest,
                cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse>(
                  this, METHODID_WATCH)))
          .build();
    }
  }

  /**
   */
  public static final class ClipboardManagerStub extends io.grpc.stub.AbstractAsyncStub<ClipboardManagerStub> {
    private ClipboardManagerStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ClipboardManagerStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ClipboardManagerStub(channel, callOptions);
    }

    /**
     */
    public void read(cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReadMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void write(cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getWriteMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void watch(cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getWatchMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class ClipboardManagerBlockingStub extends io.grpc.stub.AbstractBlockingStub<ClipboardManagerBlockingStub> {
    private ClipboardManagerBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ClipboardManagerBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ClipboardManagerBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse read(cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReadMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse write(cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getWriteMethod(), getCallOptions(), request);
    }

    /**
     */
    public java.util.Iterator<cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> watch(
        cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getWatchMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class ClipboardManagerFutureStub extends io.grpc.stub.AbstractFutureStub<ClipboardManagerFutureStub> {
    private ClipboardManagerFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ClipboardManagerFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ClipboardManagerFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse> read(
        cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReadMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse> write(
        cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getWriteMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_READ = 0;
  private static final int METHODID_WRITE = 1;
  private static final int METHODID_WATCH = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ClipboardManagerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ClipboardManagerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_READ:
          serviceImpl.read((cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse>) responseObserver);
          break;
        case METHODID_WRITE:
          serviceImpl.write((cloud.lazycat.apis.localdevice.Clipboard.WriteClipRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.WriteClipResponse>) responseObserver);
          break;
        case METHODID_WATCH:
          serviceImpl.watch((cloud.lazycat.apis.localdevice.Clipboard.ReadClipRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Clipboard.ReadClipResponse>) responseObserver);
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

  private static abstract class ClipboardManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ClipboardManagerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.localdevice.Clipboard.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ClipboardManager");
    }
  }

  private static final class ClipboardManagerFileDescriptorSupplier
      extends ClipboardManagerBaseDescriptorSupplier {
    ClipboardManagerFileDescriptorSupplier() {}
  }

  private static final class ClipboardManagerMethodDescriptorSupplier
      extends ClipboardManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    ClipboardManagerMethodDescriptorSupplier(String methodName) {
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
      synchronized (ClipboardManagerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ClipboardManagerFileDescriptorSupplier())
              .addMethod(getReadMethod())
              .addMethod(getWriteMethod())
              .addMethod(getWatchMethod())
              .build();
        }
      }
    }
    return result;
  }
}
