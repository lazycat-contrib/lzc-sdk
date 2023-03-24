package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/media_meta.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class MediaMetaGrpc {

  private MediaMetaGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.MediaMeta";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest,
      cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> getGetUserMediaMetaMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserMediaMeta",
      requestType = cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest.class,
      responseType = cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest,
      cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> getGetUserMediaMetaMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest, cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> getGetUserMediaMetaMethod;
    if ((getGetUserMediaMetaMethod = MediaMetaGrpc.getGetUserMediaMetaMethod) == null) {
      synchronized (MediaMetaGrpc.class) {
        if ((getGetUserMediaMetaMethod = MediaMetaGrpc.getGetUserMediaMetaMethod) == null) {
          MediaMetaGrpc.getGetUserMediaMetaMethod = getGetUserMediaMetaMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest, cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserMediaMeta"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MediaMetaMethodDescriptorSupplier("GetUserMediaMeta"))
              .build();
        }
      }
    }
    return getGetUserMediaMetaMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest,
      cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> getGetUserMediaMetaBatchMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserMediaMetaBatch",
      requestType = cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest.class,
      responseType = cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest,
      cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> getGetUserMediaMetaBatchMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest, cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> getGetUserMediaMetaBatchMethod;
    if ((getGetUserMediaMetaBatchMethod = MediaMetaGrpc.getGetUserMediaMetaBatchMethod) == null) {
      synchronized (MediaMetaGrpc.class) {
        if ((getGetUserMediaMetaBatchMethod = MediaMetaGrpc.getGetUserMediaMetaBatchMethod) == null) {
          MediaMetaGrpc.getGetUserMediaMetaBatchMethod = getGetUserMediaMetaBatchMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest, cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserMediaMetaBatch"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MediaMetaMethodDescriptorSupplier("GetUserMediaMetaBatch"))
              .build();
        }
      }
    }
    return getGetUserMediaMetaBatchMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static MediaMetaStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MediaMetaStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MediaMetaStub>() {
        @java.lang.Override
        public MediaMetaStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MediaMetaStub(channel, callOptions);
        }
      };
    return MediaMetaStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static MediaMetaBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MediaMetaBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MediaMetaBlockingStub>() {
        @java.lang.Override
        public MediaMetaBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MediaMetaBlockingStub(channel, callOptions);
        }
      };
    return MediaMetaBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static MediaMetaFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MediaMetaFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MediaMetaFutureStub>() {
        @java.lang.Override
        public MediaMetaFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MediaMetaFutureStub(channel, callOptions);
        }
      };
    return MediaMetaFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class MediaMetaImplBase implements io.grpc.BindableService {

    /**
     */
    public void getUserMediaMeta(cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserMediaMetaMethod(), responseObserver);
    }

    /**
     */
    public io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest> getUserMediaMetaBatch(
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> responseObserver) {
      return io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall(getGetUserMediaMetaBatchMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetUserMediaMetaMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest,
                cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse>(
                  this, METHODID_GET_USER_MEDIA_META)))
          .addMethod(
            getGetUserMediaMetaBatchMethod(),
            io.grpc.stub.ServerCalls.asyncBidiStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest,
                cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse>(
                  this, METHODID_GET_USER_MEDIA_META_BATCH)))
          .build();
    }
  }

  /**
   */
  public static final class MediaMetaStub extends io.grpc.stub.AbstractAsyncStub<MediaMetaStub> {
    private MediaMetaStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MediaMetaStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MediaMetaStub(channel, callOptions);
    }

    /**
     */
    public void getUserMediaMeta(cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserMediaMetaMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest> getUserMediaMetaBatch(
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> responseObserver) {
      return io.grpc.stub.ClientCalls.asyncBidiStreamingCall(
          getChannel().newCall(getGetUserMediaMetaBatchMethod(), getCallOptions()), responseObserver);
    }
  }

  /**
   */
  public static final class MediaMetaBlockingStub extends io.grpc.stub.AbstractBlockingStub<MediaMetaBlockingStub> {
    private MediaMetaBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MediaMetaBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MediaMetaBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse getUserMediaMeta(cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserMediaMetaMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class MediaMetaFutureStub extends io.grpc.stub.AbstractFutureStub<MediaMetaFutureStub> {
    private MediaMetaFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MediaMetaFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MediaMetaFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse> getUserMediaMeta(
        cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserMediaMetaMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_USER_MEDIA_META = 0;
  private static final int METHODID_GET_USER_MEDIA_META_BATCH = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final MediaMetaImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(MediaMetaImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_USER_MEDIA_META:
          serviceImpl.getUserMediaMeta((cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse>) responseObserver);
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
        case METHODID_GET_USER_MEDIA_META_BATCH:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.getUserMediaMetaBatch(
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.MediaMetaOuterClass.GetUserMediaMetaResponse>) responseObserver);
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class MediaMetaBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    MediaMetaBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.MediaMetaOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("MediaMeta");
    }
  }

  private static final class MediaMetaFileDescriptorSupplier
      extends MediaMetaBaseDescriptorSupplier {
    MediaMetaFileDescriptorSupplier() {}
  }

  private static final class MediaMetaMethodDescriptorSupplier
      extends MediaMetaBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    MediaMetaMethodDescriptorSupplier(String methodName) {
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
      synchronized (MediaMetaGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new MediaMetaFileDescriptorSupplier())
              .addMethod(getGetUserMediaMetaMethod())
              .addMethod(getGetUserMediaMetaBatchMethod())
              .build();
        }
      }
    }
    return result;
  }
}
