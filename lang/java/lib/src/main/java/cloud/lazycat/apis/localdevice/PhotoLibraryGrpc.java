package cloud.lazycat.apis.localdevice;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: localdevice/photo.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PhotoLibraryGrpc {

  private PhotoLibraryGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.localdevice.PhotoLibrary";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest,
      cloud.lazycat.apis.localdevice.Photo.Album> getMakeAlbumMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "MakeAlbum",
      requestType = cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Photo.Album.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest,
      cloud.lazycat.apis.localdevice.Photo.Album> getMakeAlbumMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest, cloud.lazycat.apis.localdevice.Photo.Album> getMakeAlbumMethod;
    if ((getMakeAlbumMethod = PhotoLibraryGrpc.getMakeAlbumMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getMakeAlbumMethod = PhotoLibraryGrpc.getMakeAlbumMethod) == null) {
          PhotoLibraryGrpc.getMakeAlbumMethod = getMakeAlbumMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest, cloud.lazycat.apis.localdevice.Photo.Album>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "MakeAlbum"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.Album.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("MakeAlbum"))
              .build();
        }
      }
    }
    return getMakeAlbumMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest,
      cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply> getListAlbumsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListAlbums",
      requestType = cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest,
      cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply> getListAlbumsMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest, cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply> getListAlbumsMethod;
    if ((getListAlbumsMethod = PhotoLibraryGrpc.getListAlbumsMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getListAlbumsMethod = PhotoLibraryGrpc.getListAlbumsMethod) == null) {
          PhotoLibraryGrpc.getListAlbumsMethod = getListAlbumsMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest, cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListAlbums"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("ListAlbums"))
              .build();
        }
      }
    }
    return getListAlbumsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest,
      cloud.lazycat.apis.localdevice.Photo.PutPhotoReply> getPutPhotoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PutPhoto",
      requestType = cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Photo.PutPhotoReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest,
      cloud.lazycat.apis.localdevice.Photo.PutPhotoReply> getPutPhotoMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest, cloud.lazycat.apis.localdevice.Photo.PutPhotoReply> getPutPhotoMethod;
    if ((getPutPhotoMethod = PhotoLibraryGrpc.getPutPhotoMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getPutPhotoMethod = PhotoLibraryGrpc.getPutPhotoMethod) == null) {
          PhotoLibraryGrpc.getPutPhotoMethod = getPutPhotoMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest, cloud.lazycat.apis.localdevice.Photo.PutPhotoReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PutPhoto"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.PutPhotoReply.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("PutPhoto"))
              .build();
        }
      }
    }
    return getPutPhotoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest,
      cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply> getDeletePhotoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeletePhoto",
      requestType = cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest,
      cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply> getDeletePhotoMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest, cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply> getDeletePhotoMethod;
    if ((getDeletePhotoMethod = PhotoLibraryGrpc.getDeletePhotoMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getDeletePhotoMethod = PhotoLibraryGrpc.getDeletePhotoMethod) == null) {
          PhotoLibraryGrpc.getDeletePhotoMethod = getDeletePhotoMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest, cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeletePhoto"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("DeletePhoto"))
              .build();
        }
      }
    }
    return getDeletePhotoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest,
      cloud.lazycat.apis.localdevice.Photo.PhotoMeta> getListPhotoMetasMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListPhotoMetas",
      requestType = cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Photo.PhotoMeta.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest,
      cloud.lazycat.apis.localdevice.Photo.PhotoMeta> getListPhotoMetasMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest, cloud.lazycat.apis.localdevice.Photo.PhotoMeta> getListPhotoMetasMethod;
    if ((getListPhotoMetasMethod = PhotoLibraryGrpc.getListPhotoMetasMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getListPhotoMetasMethod = PhotoLibraryGrpc.getListPhotoMetasMethod) == null) {
          PhotoLibraryGrpc.getListPhotoMetasMethod = getListPhotoMetasMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest, cloud.lazycat.apis.localdevice.Photo.PhotoMeta>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListPhotoMetas"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.PhotoMeta.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("ListPhotoMetas"))
              .build();
        }
      }
    }
    return getListPhotoMetasMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest,
      cloud.lazycat.apis.localdevice.Photo.PhotoMeta> getQueryPhotoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryPhoto",
      requestType = cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Photo.PhotoMeta.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest,
      cloud.lazycat.apis.localdevice.Photo.PhotoMeta> getQueryPhotoMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest, cloud.lazycat.apis.localdevice.Photo.PhotoMeta> getQueryPhotoMethod;
    if ((getQueryPhotoMethod = PhotoLibraryGrpc.getQueryPhotoMethod) == null) {
      synchronized (PhotoLibraryGrpc.class) {
        if ((getQueryPhotoMethod = PhotoLibraryGrpc.getQueryPhotoMethod) == null) {
          PhotoLibraryGrpc.getQueryPhotoMethod = getQueryPhotoMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest, cloud.lazycat.apis.localdevice.Photo.PhotoMeta>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryPhoto"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Photo.PhotoMeta.getDefaultInstance()))
              .setSchemaDescriptor(new PhotoLibraryMethodDescriptorSupplier("QueryPhoto"))
              .build();
        }
      }
    }
    return getQueryPhotoMethod;
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
     */
    public void makeAlbum(cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.Album> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getMakeAlbumMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列举所有的系统相册
     * </pre>
     */
    public void listAlbums(cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListAlbumsMethod(), responseObserver);
    }

    /**
     * <pre>
     * 存储一张图片到某个相册中
     * </pre>
     */
    public void putPhoto(cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PutPhotoReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPutPhotoMethod(), responseObserver);
    }

    /**
     */
    public void deletePhoto(cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeletePhotoMethod(), responseObserver);
    }

    /**
     * <pre>
     * 枚举具体相册中的图片元信息
     * </pre>
     */
    public void listPhotoMetas(cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PhotoMeta> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListPhotoMetasMethod(), responseObserver);
    }

    /**
     */
    public void queryPhoto(cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PhotoMeta> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryPhotoMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getMakeAlbumMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest,
                cloud.lazycat.apis.localdevice.Photo.Album>(
                  this, METHODID_MAKE_ALBUM)))
          .addMethod(
            getListAlbumsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest,
                cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply>(
                  this, METHODID_LIST_ALBUMS)))
          .addMethod(
            getPutPhotoMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest,
                cloud.lazycat.apis.localdevice.Photo.PutPhotoReply>(
                  this, METHODID_PUT_PHOTO)))
          .addMethod(
            getDeletePhotoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest,
                cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply>(
                  this, METHODID_DELETE_PHOTO)))
          .addMethod(
            getListPhotoMetasMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest,
                cloud.lazycat.apis.localdevice.Photo.PhotoMeta>(
                  this, METHODID_LIST_PHOTO_METAS)))
          .addMethod(
            getQueryPhotoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest,
                cloud.lazycat.apis.localdevice.Photo.PhotoMeta>(
                  this, METHODID_QUERY_PHOTO)))
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
     */
    public void makeAlbum(cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.Album> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getMakeAlbumMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列举所有的系统相册
     * </pre>
     */
    public void listAlbums(cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListAlbumsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 存储一张图片到某个相册中
     * </pre>
     */
    public void putPhoto(cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PutPhotoReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getPutPhotoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deletePhoto(cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeletePhotoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 枚举具体相册中的图片元信息
     * </pre>
     */
    public void listPhotoMetas(cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PhotoMeta> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getListPhotoMetasMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void queryPhoto(cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PhotoMeta> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryPhotoMethod(), getCallOptions()), request, responseObserver);
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
     */
    public cloud.lazycat.apis.localdevice.Photo.Album makeAlbum(cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getMakeAlbumMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列举所有的系统相册
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply listAlbums(cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListAlbumsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 存储一张图片到某个相册中
     * </pre>
     */
    public java.util.Iterator<cloud.lazycat.apis.localdevice.Photo.PutPhotoReply> putPhoto(
        cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getPutPhotoMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply deletePhoto(cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeletePhotoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 枚举具体相册中的图片元信息
     * </pre>
     */
    public java.util.Iterator<cloud.lazycat.apis.localdevice.Photo.PhotoMeta> listPhotoMetas(
        cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getListPhotoMetasMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.localdevice.Photo.PhotoMeta queryPhoto(cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryPhotoMethod(), getCallOptions(), request);
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
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.Photo.Album> makeAlbum(
        cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getMakeAlbumMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列举所有的系统相册
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply> listAlbums(
        cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListAlbumsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply> deletePhoto(
        cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeletePhotoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.Photo.PhotoMeta> queryPhoto(
        cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryPhotoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_MAKE_ALBUM = 0;
  private static final int METHODID_LIST_ALBUMS = 1;
  private static final int METHODID_PUT_PHOTO = 2;
  private static final int METHODID_DELETE_PHOTO = 3;
  private static final int METHODID_LIST_PHOTO_METAS = 4;
  private static final int METHODID_QUERY_PHOTO = 5;

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
        case METHODID_MAKE_ALBUM:
          serviceImpl.makeAlbum((cloud.lazycat.apis.localdevice.Photo.MakeAlbumRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.Album>) responseObserver);
          break;
        case METHODID_LIST_ALBUMS:
          serviceImpl.listAlbums((cloud.lazycat.apis.localdevice.Photo.ListAlbumsRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.ListAlbumsReply>) responseObserver);
          break;
        case METHODID_PUT_PHOTO:
          serviceImpl.putPhoto((cloud.lazycat.apis.localdevice.Photo.PutPhotoRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PutPhotoReply>) responseObserver);
          break;
        case METHODID_DELETE_PHOTO:
          serviceImpl.deletePhoto((cloud.lazycat.apis.localdevice.Photo.DeletePhotoRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.DeletePhotoReply>) responseObserver);
          break;
        case METHODID_LIST_PHOTO_METAS:
          serviceImpl.listPhotoMetas((cloud.lazycat.apis.localdevice.Photo.ListPhotoMetasRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PhotoMeta>) responseObserver);
          break;
        case METHODID_QUERY_PHOTO:
          serviceImpl.queryPhoto((cloud.lazycat.apis.localdevice.Photo.QueryPhotoRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Photo.PhotoMeta>) responseObserver);
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
      return cloud.lazycat.apis.localdevice.Photo.getDescriptor();
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
              .addMethod(getMakeAlbumMethod())
              .addMethod(getListAlbumsMethod())
              .addMethod(getPutPhotoMethod())
              .addMethod(getDeletePhotoMethod())
              .addMethod(getListPhotoMetasMethod())
              .addMethod(getQueryPhotoMethod())
              .build();
        }
      }
    }
    return result;
  }
}
