package cloud.lazycat.apis;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * this service is provied by frontend server, backend code shouldn't use it.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: permissions/permissions.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PermissionManagerGrpc {

  private PermissionManagerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.PermissionManager";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.Permissions.PermissionDesc,
      cloud.lazycat.apis.Permissions.PermissionToken> getRequestPermissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RequestPermission",
      requestType = cloud.lazycat.apis.Permissions.PermissionDesc.class,
      responseType = cloud.lazycat.apis.Permissions.PermissionToken.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.Permissions.PermissionDesc,
      cloud.lazycat.apis.Permissions.PermissionToken> getRequestPermissionMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.Permissions.PermissionDesc, cloud.lazycat.apis.Permissions.PermissionToken> getRequestPermissionMethod;
    if ((getRequestPermissionMethod = PermissionManagerGrpc.getRequestPermissionMethod) == null) {
      synchronized (PermissionManagerGrpc.class) {
        if ((getRequestPermissionMethod = PermissionManagerGrpc.getRequestPermissionMethod) == null) {
          PermissionManagerGrpc.getRequestPermissionMethod = getRequestPermissionMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.Permissions.PermissionDesc, cloud.lazycat.apis.Permissions.PermissionToken>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RequestPermission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.Permissions.PermissionDesc.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.Permissions.PermissionToken.getDefaultInstance()))
              .setSchemaDescriptor(new PermissionManagerMethodDescriptorSupplier("RequestPermission"))
              .build();
        }
      }
    }
    return getRequestPermissionMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PermissionManagerStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PermissionManagerStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PermissionManagerStub>() {
        @java.lang.Override
        public PermissionManagerStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PermissionManagerStub(channel, callOptions);
        }
      };
    return PermissionManagerStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PermissionManagerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PermissionManagerBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PermissionManagerBlockingStub>() {
        @java.lang.Override
        public PermissionManagerBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PermissionManagerBlockingStub(channel, callOptions);
        }
      };
    return PermissionManagerBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PermissionManagerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PermissionManagerFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PermissionManagerFutureStub>() {
        @java.lang.Override
        public PermissionManagerFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PermissionManagerFutureStub(channel, callOptions);
        }
      };
    return PermissionManagerFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static abstract class PermissionManagerImplBase implements io.grpc.BindableService {

    /**
     */
    public void requestPermission(cloud.lazycat.apis.Permissions.PermissionDesc request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.Permissions.PermissionToken> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRequestPermissionMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getRequestPermissionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.Permissions.PermissionDesc,
                cloud.lazycat.apis.Permissions.PermissionToken>(
                  this, METHODID_REQUEST_PERMISSION)))
          .build();
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class PermissionManagerStub extends io.grpc.stub.AbstractAsyncStub<PermissionManagerStub> {
    private PermissionManagerStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PermissionManagerStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PermissionManagerStub(channel, callOptions);
    }

    /**
     */
    public void requestPermission(cloud.lazycat.apis.Permissions.PermissionDesc request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.Permissions.PermissionToken> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRequestPermissionMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class PermissionManagerBlockingStub extends io.grpc.stub.AbstractBlockingStub<PermissionManagerBlockingStub> {
    private PermissionManagerBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PermissionManagerBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PermissionManagerBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.Permissions.PermissionToken requestPermission(cloud.lazycat.apis.Permissions.PermissionDesc request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRequestPermissionMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * this service is provied by frontend server, backend code shouldn't use it.
   * </pre>
   */
  public static final class PermissionManagerFutureStub extends io.grpc.stub.AbstractFutureStub<PermissionManagerFutureStub> {
    private PermissionManagerFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PermissionManagerFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PermissionManagerFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.Permissions.PermissionToken> requestPermission(
        cloud.lazycat.apis.Permissions.PermissionDesc request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRequestPermissionMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_REQUEST_PERMISSION = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final PermissionManagerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(PermissionManagerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_REQUEST_PERMISSION:
          serviceImpl.requestPermission((cloud.lazycat.apis.Permissions.PermissionDesc) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.Permissions.PermissionToken>) responseObserver);
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

  private static abstract class PermissionManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PermissionManagerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.Permissions.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PermissionManager");
    }
  }

  private static final class PermissionManagerFileDescriptorSupplier
      extends PermissionManagerBaseDescriptorSupplier {
    PermissionManagerFileDescriptorSupplier() {}
  }

  private static final class PermissionManagerMethodDescriptorSupplier
      extends PermissionManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    PermissionManagerMethodDescriptorSupplier(String methodName) {
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
      synchronized (PermissionManagerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PermissionManagerFileDescriptorSupplier())
              .addMethod(getRequestPermissionMethod())
              .build();
        }
      }
    }
    return result;
  }
}
