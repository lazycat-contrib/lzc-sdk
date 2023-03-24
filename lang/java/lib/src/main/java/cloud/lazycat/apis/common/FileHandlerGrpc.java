package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/file_handler.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class FileHandlerGrpc {

  private FileHandlerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.FileHandler";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest,
      cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply> getQueryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "query",
      requestType = cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest.class,
      responseType = cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest,
      cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply> getQueryMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest, cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply> getQueryMethod;
    if ((getQueryMethod = FileHandlerGrpc.getQueryMethod) == null) {
      synchronized (FileHandlerGrpc.class) {
        if ((getQueryMethod = FileHandlerGrpc.getQueryMethod) == null) {
          FileHandlerGrpc.getQueryMethod = getQueryMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest, cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "query"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply.getDefaultInstance()))
              .setSchemaDescriptor(new FileHandlerMethodDescriptorSupplier("query"))
              .build();
        }
      }
    }
    return getQueryMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest,
      cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply> getOpenMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "open",
      requestType = cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest.class,
      responseType = cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest,
      cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply> getOpenMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest, cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply> getOpenMethod;
    if ((getOpenMethod = FileHandlerGrpc.getOpenMethod) == null) {
      synchronized (FileHandlerGrpc.class) {
        if ((getOpenMethod = FileHandlerGrpc.getOpenMethod) == null) {
          FileHandlerGrpc.getOpenMethod = getOpenMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest, cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "open"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply.getDefaultInstance()))
              .setSchemaDescriptor(new FileHandlerMethodDescriptorSupplier("open"))
              .build();
        }
      }
    }
    return getOpenMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static FileHandlerStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileHandlerStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileHandlerStub>() {
        @java.lang.Override
        public FileHandlerStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileHandlerStub(channel, callOptions);
        }
      };
    return FileHandlerStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static FileHandlerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileHandlerBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileHandlerBlockingStub>() {
        @java.lang.Override
        public FileHandlerBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileHandlerBlockingStub(channel, callOptions);
        }
      };
    return FileHandlerBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static FileHandlerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileHandlerFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileHandlerFutureStub>() {
        @java.lang.Override
        public FileHandlerFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileHandlerFutureStub(channel, callOptions);
        }
      };
    return FileHandlerFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class FileHandlerImplBase implements io.grpc.BindableService {

    /**
     */
    public void query(cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryMethod(), responseObserver);
    }

    /**
     */
    public void open(cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getOpenMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getQueryMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest,
                cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply>(
                  this, METHODID_QUERY)))
          .addMethod(
            getOpenMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest,
                cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply>(
                  this, METHODID_OPEN)))
          .build();
    }
  }

  /**
   */
  public static final class FileHandlerStub extends io.grpc.stub.AbstractAsyncStub<FileHandlerStub> {
    private FileHandlerStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileHandlerStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileHandlerStub(channel, callOptions);
    }

    /**
     */
    public void query(cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void open(cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getOpenMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class FileHandlerBlockingStub extends io.grpc.stub.AbstractBlockingStub<FileHandlerBlockingStub> {
    private FileHandlerBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileHandlerBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileHandlerBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply query(cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply open(cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getOpenMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class FileHandlerFutureStub extends io.grpc.stub.AbstractFutureStub<FileHandlerFutureStub> {
    private FileHandlerFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileHandlerFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileHandlerFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply> query(
        cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply> open(
        cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getOpenMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_QUERY = 0;
  private static final int METHODID_OPEN = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final FileHandlerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(FileHandlerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_QUERY:
          serviceImpl.query((cloud.lazycat.apis.common.FileHandlerOuterClass.QueryRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.FileHandlerOuterClass.QueryReply>) responseObserver);
          break;
        case METHODID_OPEN:
          serviceImpl.open((cloud.lazycat.apis.common.FileHandlerOuterClass.OpenRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.FileHandlerOuterClass.OpenReply>) responseObserver);
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

  private static abstract class FileHandlerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    FileHandlerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.FileHandlerOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("FileHandler");
    }
  }

  private static final class FileHandlerFileDescriptorSupplier
      extends FileHandlerBaseDescriptorSupplier {
    FileHandlerFileDescriptorSupplier() {}
  }

  private static final class FileHandlerMethodDescriptorSupplier
      extends FileHandlerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    FileHandlerMethodDescriptorSupplier(String methodName) {
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
      synchronized (FileHandlerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new FileHandlerFileDescriptorSupplier())
              .addMethod(getQueryMethod())
              .addMethod(getOpenMethod())
              .build();
        }
      }
    }
    return result;
  }
}
