package cloud.lazycat.apis.localdevice;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * 权限管理
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: localdevice/permission.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PermissionManagerGrpc {

  private PermissionManagerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.localdevice.PermissionManager";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest,
      cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> getGetPermissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPermission",
      requestType = cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest.class,
      responseType = cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest,
      cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> getGetPermissionMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest, cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> getGetPermissionMethod;
    if ((getGetPermissionMethod = PermissionManagerGrpc.getGetPermissionMethod) == null) {
      synchronized (PermissionManagerGrpc.class) {
        if ((getGetPermissionMethod = PermissionManagerGrpc.getGetPermissionMethod) == null) {
          PermissionManagerGrpc.getGetPermissionMethod = getGetPermissionMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest, cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPermission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply.getDefaultInstance()))
              .setSchemaDescriptor(new PermissionManagerMethodDescriptorSupplier("GetPermission"))
              .build();
        }
      }
    }
    return getGetPermissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest,
      cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> getRequestPermissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RequestPermission",
      requestType = cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest.class,
      responseType = cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest,
      cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> getRequestPermissionMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest, cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> getRequestPermissionMethod;
    if ((getRequestPermissionMethod = PermissionManagerGrpc.getRequestPermissionMethod) == null) {
      synchronized (PermissionManagerGrpc.class) {
        if ((getRequestPermissionMethod = PermissionManagerGrpc.getRequestPermissionMethod) == null) {
          PermissionManagerGrpc.getRequestPermissionMethod = getRequestPermissionMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest, cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RequestPermission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply.getDefaultInstance()))
              .setSchemaDescriptor(new PermissionManagerMethodDescriptorSupplier("RequestPermission"))
              .build();
        }
      }
    }
    return getRequestPermissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply> getListPermissionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListPermissions",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply> getListPermissionsMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply> getListPermissionsMethod;
    if ((getListPermissionsMethod = PermissionManagerGrpc.getListPermissionsMethod) == null) {
      synchronized (PermissionManagerGrpc.class) {
        if ((getListPermissionsMethod = PermissionManagerGrpc.getListPermissionsMethod) == null) {
          PermissionManagerGrpc.getListPermissionsMethod = getListPermissionsMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListPermissions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply.getDefaultInstance()))
              .setSchemaDescriptor(new PermissionManagerMethodDescriptorSupplier("ListPermissions"))
              .build();
        }
      }
    }
    return getListPermissionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest,
      cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse> getRequestAuthTokenMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RequestAuthToken",
      requestType = cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest.class,
      responseType = cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest,
      cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse> getRequestAuthTokenMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest, cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse> getRequestAuthTokenMethod;
    if ((getRequestAuthTokenMethod = PermissionManagerGrpc.getRequestAuthTokenMethod) == null) {
      synchronized (PermissionManagerGrpc.class) {
        if ((getRequestAuthTokenMethod = PermissionManagerGrpc.getRequestAuthTokenMethod) == null) {
          PermissionManagerGrpc.getRequestAuthTokenMethod = getRequestAuthTokenMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest, cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RequestAuthToken"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PermissionManagerMethodDescriptorSupplier("RequestAuthToken"))
              .build();
        }
      }
    }
    return getRequestAuthTokenMethod;
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
   * 权限管理
   * </pre>
   */
  public static abstract class PermissionManagerImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 检测权限
     * </pre>
     */
    public void getPermission(cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPermissionMethod(), responseObserver);
    }

    /**
     * <pre>
     * 申请权限（会弹出对话框让用户决定是否同意）
     * </pre>
     */
    public void requestPermission(cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRequestPermissionMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列举权限
     * </pre>
     */
    public void listPermissions(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListPermissionsMethod(), responseObserver);
    }

    /**
     * <pre>
     * 申请Token
     * </pre>
     */
    public void requestAuthToken(cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRequestAuthTokenMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetPermissionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest,
                cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply>(
                  this, METHODID_GET_PERMISSION)))
          .addMethod(
            getRequestPermissionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest,
                cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply>(
                  this, METHODID_REQUEST_PERMISSION)))
          .addMethod(
            getListPermissionsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply>(
                  this, METHODID_LIST_PERMISSIONS)))
          .addMethod(
            getRequestAuthTokenMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest,
                cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse>(
                  this, METHODID_REQUEST_AUTH_TOKEN)))
          .build();
    }
  }

  /**
   * <pre>
   * 权限管理
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
     * <pre>
     * 检测权限
     * </pre>
     */
    public void getPermission(cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPermissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 申请权限（会弹出对话框让用户决定是否同意）
     * </pre>
     */
    public void requestPermission(cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRequestPermissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列举权限
     * </pre>
     */
    public void listPermissions(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListPermissionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 申请Token
     * </pre>
     */
    public void requestAuthToken(cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRequestAuthTokenMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * 权限管理
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
     * <pre>
     * 检测权限
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply getPermission(cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPermissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 申请权限（会弹出对话框让用户决定是否同意）
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply requestPermission(cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRequestPermissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列举权限
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply listPermissions(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListPermissionsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 申请Token
     * </pre>
     */
    public cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse requestAuthToken(cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRequestAuthTokenMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * 权限管理
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
     * <pre>
     * 检测权限
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> getPermission(
        cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPermissionMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 申请权限（会弹出对话框让用户决定是否同意）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply> requestPermission(
        cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRequestPermissionMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列举权限
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply> listPermissions(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListPermissionsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 申请Token
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse> requestAuthToken(
        cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRequestAuthTokenMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_PERMISSION = 0;
  private static final int METHODID_REQUEST_PERMISSION = 1;
  private static final int METHODID_LIST_PERMISSIONS = 2;
  private static final int METHODID_REQUEST_AUTH_TOKEN = 3;

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
        case METHODID_GET_PERMISSION:
          serviceImpl.getPermission((cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply>) responseObserver);
          break;
        case METHODID_REQUEST_PERMISSION:
          serviceImpl.requestPermission((cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.PermissionReply>) responseObserver);
          break;
        case METHODID_LIST_PERMISSIONS:
          serviceImpl.listPermissions((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.ListPermissionsReply>) responseObserver);
          break;
        case METHODID_REQUEST_AUTH_TOKEN:
          serviceImpl.requestAuthToken((cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.PermissionOuterClass.RequestAuthTokenResponse>) responseObserver);
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
      return cloud.lazycat.apis.localdevice.PermissionOuterClass.getDescriptor();
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
              .addMethod(getGetPermissionMethod())
              .addMethod(getRequestPermissionMethod())
              .addMethod(getListPermissionsMethod())
              .addMethod(getRequestAuthTokenMethod())
              .build();
        }
      }
    }
    return result;
  }
}
