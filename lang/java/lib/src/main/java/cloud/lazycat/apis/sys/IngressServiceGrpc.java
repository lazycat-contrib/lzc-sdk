package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * 这里 Access Control 的逻辑是：
 * 1. 默认用户可以访问所有 app
 * 2. 如果调用了 AllowAdd ，则该用户进入白名单模式，只有被管理员允许访问的应用才可以访问
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/ingress.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class IngressServiceGrpc {

  private IngressServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.IngressService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowRequest,
      com.google.protobuf.Empty> getAllowAddMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AllowAdd",
      requestType = cloud.lazycat.apis.sys.Ingress.IngressAllowRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowRequest,
      com.google.protobuf.Empty> getAllowAddMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowRequest, com.google.protobuf.Empty> getAllowAddMethod;
    if ((getAllowAddMethod = IngressServiceGrpc.getAllowAddMethod) == null) {
      synchronized (IngressServiceGrpc.class) {
        if ((getAllowAddMethod = IngressServiceGrpc.getAllowAddMethod) == null) {
          IngressServiceGrpc.getAllowAddMethod = getAllowAddMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Ingress.IngressAllowRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AllowAdd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new IngressServiceMethodDescriptorSupplier("AllowAdd"))
              .build();
        }
      }
    }
    return getAllowAddMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowRequest,
      com.google.protobuf.Empty> getAllowDelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AllowDel",
      requestType = cloud.lazycat.apis.sys.Ingress.IngressAllowRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowRequest,
      com.google.protobuf.Empty> getAllowDelMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowRequest, com.google.protobuf.Empty> getAllowDelMethod;
    if ((getAllowDelMethod = IngressServiceGrpc.getAllowDelMethod) == null) {
      synchronized (IngressServiceGrpc.class) {
        if ((getAllowDelMethod = IngressServiceGrpc.getAllowDelMethod) == null) {
          IngressServiceGrpc.getAllowDelMethod = getAllowDelMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Ingress.IngressAllowRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AllowDel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new IngressServiceMethodDescriptorSupplier("AllowDel"))
              .build();
        }
      }
    }
    return getAllowDelMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest,
      com.google.protobuf.Empty> getAllowDelAllMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AllowDelAll",
      requestType = cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest,
      com.google.protobuf.Empty> getAllowDelAllMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest, com.google.protobuf.Empty> getAllowDelAllMethod;
    if ((getAllowDelAllMethod = IngressServiceGrpc.getAllowDelAllMethod) == null) {
      synchronized (IngressServiceGrpc.class) {
        if ((getAllowDelAllMethod = IngressServiceGrpc.getAllowDelAllMethod) == null) {
          IngressServiceGrpc.getAllowDelAllMethod = getAllowDelAllMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AllowDelAll"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new IngressServiceMethodDescriptorSupplier("AllowDelAll"))
              .build();
        }
      }
    }
    return getAllowDelAllMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest,
      cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse> getAllowGetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AllowGet",
      requestType = cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest.class,
      responseType = cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest,
      cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse> getAllowGetMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest, cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse> getAllowGetMethod;
    if ((getAllowGetMethod = IngressServiceGrpc.getAllowGetMethod) == null) {
      synchronized (IngressServiceGrpc.class) {
        if ((getAllowGetMethod = IngressServiceGrpc.getAllowGetMethod) == null) {
          IngressServiceGrpc.getAllowGetMethod = getAllowGetMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest, cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AllowGet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse.getDefaultInstance()))
              .setSchemaDescriptor(new IngressServiceMethodDescriptorSupplier("AllowGet"))
              .build();
        }
      }
    }
    return getAllowGetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest,
      cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse> getAllowListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AllowList",
      requestType = cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest.class,
      responseType = cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest,
      cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse> getAllowListMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest, cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse> getAllowListMethod;
    if ((getAllowListMethod = IngressServiceGrpc.getAllowListMethod) == null) {
      synchronized (IngressServiceGrpc.class) {
        if ((getAllowListMethod = IngressServiceGrpc.getAllowListMethod) == null) {
          IngressServiceGrpc.getAllowListMethod = getAllowListMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest, cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AllowList"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse.getDefaultInstance()))
              .setSchemaDescriptor(new IngressServiceMethodDescriptorSupplier("AllowList"))
              .build();
        }
      }
    }
    return getAllowListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getAllowClearMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AllowClear",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getAllowClearMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getAllowClearMethod;
    if ((getAllowClearMethod = IngressServiceGrpc.getAllowClearMethod) == null) {
      synchronized (IngressServiceGrpc.class) {
        if ((getAllowClearMethod = IngressServiceGrpc.getAllowClearMethod) == null) {
          IngressServiceGrpc.getAllowClearMethod = getAllowClearMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AllowClear"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new IngressServiceMethodDescriptorSupplier("AllowClear"))
              .build();
        }
      }
    }
    return getAllowClearMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest,
      cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse> getAllowManageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AllowManage",
      requestType = cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest.class,
      responseType = cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest,
      cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse> getAllowManageMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest, cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse> getAllowManageMethod;
    if ((getAllowManageMethod = IngressServiceGrpc.getAllowManageMethod) == null) {
      synchronized (IngressServiceGrpc.class) {
        if ((getAllowManageMethod = IngressServiceGrpc.getAllowManageMethod) == null) {
          IngressServiceGrpc.getAllowManageMethod = getAllowManageMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest, cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AllowManage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse.getDefaultInstance()))
              .setSchemaDescriptor(new IngressServiceMethodDescriptorSupplier("AllowManage"))
              .build();
        }
      }
    }
    return getAllowManageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest,
      cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse> getGetAppLastAccessTimeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAppLastAccessTime",
      requestType = cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest.class,
      responseType = cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest,
      cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse> getGetAppLastAccessTimeMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest, cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse> getGetAppLastAccessTimeMethod;
    if ((getGetAppLastAccessTimeMethod = IngressServiceGrpc.getGetAppLastAccessTimeMethod) == null) {
      synchronized (IngressServiceGrpc.class) {
        if ((getGetAppLastAccessTimeMethod = IngressServiceGrpc.getGetAppLastAccessTimeMethod) == null) {
          IngressServiceGrpc.getGetAppLastAccessTimeMethod = getGetAppLastAccessTimeMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest, cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAppLastAccessTime"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse.getDefaultInstance()))
              .setSchemaDescriptor(new IngressServiceMethodDescriptorSupplier("GetAppLastAccessTime"))
              .build();
        }
      }
    }
    return getGetAppLastAccessTimeMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static IngressServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<IngressServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<IngressServiceStub>() {
        @java.lang.Override
        public IngressServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new IngressServiceStub(channel, callOptions);
        }
      };
    return IngressServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static IngressServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<IngressServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<IngressServiceBlockingStub>() {
        @java.lang.Override
        public IngressServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new IngressServiceBlockingStub(channel, callOptions);
        }
      };
    return IngressServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static IngressServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<IngressServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<IngressServiceFutureStub>() {
        @java.lang.Override
        public IngressServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new IngressServiceFutureStub(channel, callOptions);
        }
      };
    return IngressServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * 这里 Access Control 的逻辑是：
   * 1. 默认用户可以访问所有 app
   * 2. 如果调用了 AllowAdd ，则该用户进入白名单模式，只有被管理员允许访问的应用才可以访问
   * </pre>
   */
  public static abstract class IngressServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 允许一个用户对指定 app 的访问
     * </pre>
     */
    public void allowAdd(cloud.lazycat.apis.sys.Ingress.IngressAllowRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAllowAddMethod(), responseObserver);
    }

    /**
     * <pre>
     * 阻止一个用户对指定 app 的访问权限
     * </pre>
     */
    public void allowDel(cloud.lazycat.apis.sys.Ingress.IngressAllowRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAllowDelMethod(), responseObserver);
    }

    /**
     * <pre>
     * 阻止一个用户对所有 app 的访问权限
     * </pre>
     */
    public void allowDelAll(cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAllowDelAllMethod(), responseObserver);
    }

    /**
     * <pre>
     * 查询一个用户是否被允许访问指定 app
     * </pre>
     */
    public void allowGet(cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAllowGetMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列出一个用户的所有被允许访问的 app 列表
     * </pre>
     */
    public void allowList(cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAllowListMethod(), responseObserver);
    }

    /**
     * <pre>
     * 清空白名单
     * </pre>
     */
    public void allowClear(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAllowClearMethod(), responseObserver);
    }

    /**
     * <pre>
     * 启用/禁用白名单
     * </pre>
     */
    public void allowManage(cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAllowManageMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取指定 app 最后一次被访问的时刻
     * </pre>
     */
    public void getAppLastAccessTime(cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAppLastAccessTimeMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getAllowAddMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Ingress.IngressAllowRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_ALLOW_ADD)))
          .addMethod(
            getAllowDelMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Ingress.IngressAllowRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_ALLOW_DEL)))
          .addMethod(
            getAllowDelAllMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_ALLOW_DEL_ALL)))
          .addMethod(
            getAllowGetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest,
                cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse>(
                  this, METHODID_ALLOW_GET)))
          .addMethod(
            getAllowListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest,
                cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse>(
                  this, METHODID_ALLOW_LIST)))
          .addMethod(
            getAllowClearMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_ALLOW_CLEAR)))
          .addMethod(
            getAllowManageMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest,
                cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse>(
                  this, METHODID_ALLOW_MANAGE)))
          .addMethod(
            getGetAppLastAccessTimeMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest,
                cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse>(
                  this, METHODID_GET_APP_LAST_ACCESS_TIME)))
          .build();
    }
  }

  /**
   * <pre>
   * 这里 Access Control 的逻辑是：
   * 1. 默认用户可以访问所有 app
   * 2. 如果调用了 AllowAdd ，则该用户进入白名单模式，只有被管理员允许访问的应用才可以访问
   * </pre>
   */
  public static final class IngressServiceStub extends io.grpc.stub.AbstractAsyncStub<IngressServiceStub> {
    private IngressServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected IngressServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new IngressServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 允许一个用户对指定 app 的访问
     * </pre>
     */
    public void allowAdd(cloud.lazycat.apis.sys.Ingress.IngressAllowRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAllowAddMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 阻止一个用户对指定 app 的访问权限
     * </pre>
     */
    public void allowDel(cloud.lazycat.apis.sys.Ingress.IngressAllowRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAllowDelMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 阻止一个用户对所有 app 的访问权限
     * </pre>
     */
    public void allowDelAll(cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAllowDelAllMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 查询一个用户是否被允许访问指定 app
     * </pre>
     */
    public void allowGet(cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAllowGetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列出一个用户的所有被允许访问的 app 列表
     * </pre>
     */
    public void allowList(cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAllowListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 清空白名单
     * </pre>
     */
    public void allowClear(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAllowClearMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 启用/禁用白名单
     * </pre>
     */
    public void allowManage(cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAllowManageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取指定 app 最后一次被访问的时刻
     * </pre>
     */
    public void getAppLastAccessTime(cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAppLastAccessTimeMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * 这里 Access Control 的逻辑是：
   * 1. 默认用户可以访问所有 app
   * 2. 如果调用了 AllowAdd ，则该用户进入白名单模式，只有被管理员允许访问的应用才可以访问
   * </pre>
   */
  public static final class IngressServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<IngressServiceBlockingStub> {
    private IngressServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected IngressServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new IngressServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 允许一个用户对指定 app 的访问
     * </pre>
     */
    public com.google.protobuf.Empty allowAdd(cloud.lazycat.apis.sys.Ingress.IngressAllowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAllowAddMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 阻止一个用户对指定 app 的访问权限
     * </pre>
     */
    public com.google.protobuf.Empty allowDel(cloud.lazycat.apis.sys.Ingress.IngressAllowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAllowDelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 阻止一个用户对所有 app 的访问权限
     * </pre>
     */
    public com.google.protobuf.Empty allowDelAll(cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAllowDelAllMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 查询一个用户是否被允许访问指定 app
     * </pre>
     */
    public cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse allowGet(cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAllowGetMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列出一个用户的所有被允许访问的 app 列表
     * </pre>
     */
    public cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse allowList(cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAllowListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 清空白名单
     * </pre>
     */
    public com.google.protobuf.Empty allowClear(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAllowClearMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 启用/禁用白名单
     * </pre>
     */
    public cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse allowManage(cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAllowManageMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取指定 app 最后一次被访问的时刻
     * </pre>
     */
    public cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse getAppLastAccessTime(cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAppLastAccessTimeMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * 这里 Access Control 的逻辑是：
   * 1. 默认用户可以访问所有 app
   * 2. 如果调用了 AllowAdd ，则该用户进入白名单模式，只有被管理员允许访问的应用才可以访问
   * </pre>
   */
  public static final class IngressServiceFutureStub extends io.grpc.stub.AbstractFutureStub<IngressServiceFutureStub> {
    private IngressServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected IngressServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new IngressServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 允许一个用户对指定 app 的访问
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> allowAdd(
        cloud.lazycat.apis.sys.Ingress.IngressAllowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAllowAddMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 阻止一个用户对指定 app 的访问权限
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> allowDel(
        cloud.lazycat.apis.sys.Ingress.IngressAllowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAllowDelMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 阻止一个用户对所有 app 的访问权限
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> allowDelAll(
        cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAllowDelAllMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 查询一个用户是否被允许访问指定 app
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse> allowGet(
        cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAllowGetMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列出一个用户的所有被允许访问的 app 列表
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse> allowList(
        cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAllowListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 清空白名单
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> allowClear(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAllowClearMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 启用/禁用白名单
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse> allowManage(
        cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAllowManageMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取指定 app 最后一次被访问的时刻
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse> getAppLastAccessTime(
        cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAppLastAccessTimeMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_ALLOW_ADD = 0;
  private static final int METHODID_ALLOW_DEL = 1;
  private static final int METHODID_ALLOW_DEL_ALL = 2;
  private static final int METHODID_ALLOW_GET = 3;
  private static final int METHODID_ALLOW_LIST = 4;
  private static final int METHODID_ALLOW_CLEAR = 5;
  private static final int METHODID_ALLOW_MANAGE = 6;
  private static final int METHODID_GET_APP_LAST_ACCESS_TIME = 7;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final IngressServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(IngressServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_ALLOW_ADD:
          serviceImpl.allowAdd((cloud.lazycat.apis.sys.Ingress.IngressAllowRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_ALLOW_DEL:
          serviceImpl.allowDel((cloud.lazycat.apis.sys.Ingress.IngressAllowRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_ALLOW_DEL_ALL:
          serviceImpl.allowDelAll((cloud.lazycat.apis.sys.Ingress.IngressAllowDelAllRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_ALLOW_GET:
          serviceImpl.allowGet((cloud.lazycat.apis.sys.Ingress.IngressAllowGetRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowGetResponse>) responseObserver);
          break;
        case METHODID_ALLOW_LIST:
          serviceImpl.allowList((cloud.lazycat.apis.sys.Ingress.IngressAllowListRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowListResponse>) responseObserver);
          break;
        case METHODID_ALLOW_CLEAR:
          serviceImpl.allowClear((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_ALLOW_MANAGE:
          serviceImpl.allowManage((cloud.lazycat.apis.sys.Ingress.IngressAllowManageRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAllowManageResponse>) responseObserver);
          break;
        case METHODID_GET_APP_LAST_ACCESS_TIME:
          serviceImpl.getAppLastAccessTime((cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Ingress.IngressAppLastAccessTimeResponse>) responseObserver);
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

  private static abstract class IngressServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    IngressServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.Ingress.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("IngressService");
    }
  }

  private static final class IngressServiceFileDescriptorSupplier
      extends IngressServiceBaseDescriptorSupplier {
    IngressServiceFileDescriptorSupplier() {}
  }

  private static final class IngressServiceMethodDescriptorSupplier
      extends IngressServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    IngressServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (IngressServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new IngressServiceFileDescriptorSupplier())
              .addMethod(getAllowAddMethod())
              .addMethod(getAllowDelMethod())
              .addMethod(getAllowDelAllMethod())
              .addMethod(getAllowGetMethod())
              .addMethod(getAllowListMethod())
              .addMethod(getAllowClearMethod())
              .addMethod(getAllowManageMethod())
              .addMethod(getGetAppLastAccessTimeMethod())
              .build();
        }
      }
    }
    return result;
  }
}
