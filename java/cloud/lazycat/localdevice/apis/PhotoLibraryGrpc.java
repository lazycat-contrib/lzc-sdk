package cloud.lazycat.localdevice.apis;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: localdevice/photo/photo.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PhotoLibraryGrpc {

  private PhotoLibraryGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.localdevice.apis.PhotoLibrary";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest,
      cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply> getListAlbumsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListAlbums",
      requestType = cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest.class,
      responseType = cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest,
      cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply> getListAlbumsMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest, cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply> getListAlbumsMethod;
    if ((getListAlbumsMethod = PhotoLibraryGrpc.getListAlbumsMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getListAlbumsMethod = PhotoLibraryGrpc.getListAlbumsMethod) == null) {
          PhotoLibraryGrpc.getListAlbumsMethod = getListAlbumsMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest, cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListAlbums"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("ListAlbums"))
              .build();
        }
      }
    }
    return getListAlbumsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest,
      com.google.protobuf.Empty> getPutPhotoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PutPhoto",
      requestType = cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest,
      com.google.protobuf.Empty> getPutPhotoMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest, com.google.protobuf.Empty> getPutPhotoMethod;
    if ((getPutPhotoMethod = PhotoLibraryGrpc.getPutPhotoMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getPutPhotoMethod = PhotoLibraryGrpc.getPutPhotoMethod) == null) {
          PhotoLibraryGrpc.getPutPhotoMethod = getPutPhotoMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PutPhoto"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("PutPhoto"))
              .build();
        }
      }
    }
    return getPutPhotoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest,
      cloud.lazycat.localdevice.apis.Photo.PhotoMeta> getListPhotoMetasMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListPhotoMetas",
      requestType = cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest.class,
      responseType = cloud.lazycat.localdevice.apis.Photo.PhotoMeta.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest,
      cloud.lazycat.localdevice.apis.Photo.PhotoMeta> getListPhotoMetasMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest, cloud.lazycat.localdevice.apis.Photo.PhotoMeta> getListPhotoMetasMethod;
    if ((getListPhotoMetasMethod = PhotoLibraryGrpc.getListPhotoMetasMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getListPhotoMetasMethod = PhotoLibraryGrpc.getListPhotoMetasMethod) == null) {
          PhotoLibraryGrpc.getListPhotoMetasMethod = getListPhotoMetasMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest, cloud.lazycat.localdevice.apis.Photo.PhotoMeta>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListPhotoMetas"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Photo.PhotoMeta.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("ListPhotoMetas"))
              .build();
        }
      }
    }
    return getListPhotoMetasMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PhotoLibraryStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PhotoLibraryStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PhotoLibraryStub>() {
        @java.lang.Override
        public PhotoLibraryStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PhotoLibraryStub(channel, callOptions);
        }
      };
    return PhotoLibraryStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PhotoLibraryBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PhotoLibraryBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PhotoLibraryBlockingStub>() {
        @java.lang.Override
        public PhotoLibraryBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PhotoLibraryBlockingStub(channel, callOptions);
        }
      };
    return PhotoLibraryBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PhotoLibraryFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PhotoLibraryFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PhotoLibraryFutureStub>() {
        @java.lang.Override
        public PhotoLibraryFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PhotoLibraryFutureStub(channel, callOptions);
        }
      };
    return PhotoLibraryFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class PhotoLibraryImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 列举所有的系统相册
     * </pre>
     */
    public void listAlbums(cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListAlbumsMethod(), responseObserver);
    }

    /**
     * <pre>
     * 存储一张图片到某个相册中
     * </pre>
     */
    public void putPhoto(cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPutPhotoMethod(), responseObserver);
    }

    /**
     * <pre>
     * 枚举具体相册中的图片元信息
     * </pre>
     */
    public void listPhotoMetas(cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Photo.PhotoMeta> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListPhotoMetasMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListAlbumsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest,
                cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply>(
                  this, METHODID_LIST_ALBUMS)))
          .addMethod(
            getPutPhotoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_PUT_PHOTO)))
          .addMethod(
            getListPhotoMetasMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest,
                cloud.lazycat.localdevice.apis.Photo.PhotoMeta>(
                  this, METHODID_LIST_PHOTO_METAS)))
          .build();
    }
  }

  /**
   */
  public static final class PhotoLibraryStub extends io.grpc.stub.AbstractAsyncStub<PhotoLibraryStub> {
    private PhotoLibraryStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PhotoLibraryStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PhotoLibraryStub(channel, callOptions);
    }

    /**
     * <pre>
     * 列举所有的系统相册
     * </pre>
     */
    public void listAlbums(cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListAlbumsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 存储一张图片到某个相册中
     * </pre>
     */
    public void putPhoto(cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPutPhotoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 枚举具体相册中的图片元信息
     * </pre>
     */
    public void listPhotoMetas(cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Photo.PhotoMeta> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getListPhotoMetasMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class PhotoLibraryBlockingStub extends io.grpc.stub.AbstractBlockingStub<PhotoLibraryBlockingStub> {
    private PhotoLibraryBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PhotoLibraryBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PhotoLibraryBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 列举所有的系统相册
     * </pre>
     */
    public cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply listAlbums(cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListAlbumsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 存储一张图片到某个相册中
     * </pre>
     */
    public com.google.protobuf.Empty putPhoto(cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPutPhotoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 枚举具体相册中的图片元信息
     * </pre>
     */
    public java.util.Iterator<cloud.lazycat.localdevice.apis.Photo.PhotoMeta> listPhotoMetas(
        cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getListPhotoMetasMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class PhotoLibraryFutureStub extends io.grpc.stub.AbstractFutureStub<PhotoLibraryFutureStub> {
    private PhotoLibraryFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PhotoLibraryFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PhotoLibraryFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 列举所有的系统相册
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply> listAlbums(
        cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListAlbumsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 存储一张图片到某个相册中
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> putPhoto(
        cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPutPhotoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_ALBUMS = 0;
  private static final int METHODID_PUT_PHOTO = 1;
  private static final int METHODID_LIST_PHOTO_METAS = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final PhotoLibraryImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(PhotoLibraryImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LIST_ALBUMS:
          serviceImpl.listAlbums((cloud.lazycat.localdevice.apis.Photo.ListAlbumsRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Photo.ListAlbumsReply>) responseObserver);
          break;
        case METHODID_PUT_PHOTO:
          serviceImpl.putPhoto((cloud.lazycat.localdevice.apis.Photo.PutPhotoRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_LIST_PHOTO_METAS:
          serviceImpl.listPhotoMetas((cloud.lazycat.localdevice.apis.Photo.ListPhotoMetasRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Photo.PhotoMeta>) responseObserver);
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

  private static abstract class PhotoLibraryBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PhotoLibraryBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.localdevice.apis.Photo.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PhotoLibrary");
    }
  }

  private static final class PhotoLibraryFileDescriptorSupplier
      extends PhotoLibraryBaseDescriptorSupplier {
    PhotoLibraryFileDescriptorSupplier() {}
  }

  private static final class PhotoLibraryMethodDescriptorSupplier
      extends PhotoLibraryBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    PhotoLibraryMethodDescriptorSupplier(String methodName) {
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
      synchronized (PhotoLibraryGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PhotoLibraryFileDescriptorSupplier())
              .addMethod(getListAlbumsMethod())
              .addMethod(getPutPhotoMethod())
              .addMethod(getListPhotoMetasMethod())
              .build();
        }
      }
    }
    return result;
  }
}
