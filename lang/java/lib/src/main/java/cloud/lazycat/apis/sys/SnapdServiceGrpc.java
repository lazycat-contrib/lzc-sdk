package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/snapd.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class SnapdServiceGrpc {

  private SnapdServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.SnapdService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest,
      com.google.protobuf.Empty> getSnapdEnableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapdEnable",
      requestType = cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest,
      com.google.protobuf.Empty> getSnapdEnableMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest, com.google.protobuf.Empty> getSnapdEnableMethod;
    if ((getSnapdEnableMethod = SnapdServiceGrpc.getSnapdEnableMethod) == null) {
      synchronized (SnapdServiceGrpc.class) {
        if ((getSnapdEnableMethod = SnapdServiceGrpc.getSnapdEnableMethod) == null) {
          SnapdServiceGrpc.getSnapdEnableMethod = getSnapdEnableMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapdEnable"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new SnapdServiceMethodDescriptorSupplier("SnapdEnable"))
              .build();
        }
      }
    }
    return getSnapdEnableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest,
      com.google.protobuf.Empty> getSnapdDisableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapdDisable",
      requestType = cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest,
      com.google.protobuf.Empty> getSnapdDisableMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest, com.google.protobuf.Empty> getSnapdDisableMethod;
    if ((getSnapdDisableMethod = SnapdServiceGrpc.getSnapdDisableMethod) == null) {
      synchronized (SnapdServiceGrpc.class) {
        if ((getSnapdDisableMethod = SnapdServiceGrpc.getSnapdDisableMethod) == null) {
          SnapdServiceGrpc.getSnapdDisableMethod = getSnapdDisableMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapdDisable"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new SnapdServiceMethodDescriptorSupplier("SnapdDisable"))
              .build();
        }
      }
    }
    return getSnapdDisableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest,
      cloud.lazycat.apis.sys.Snapd.SnapdConf> getSnapdConfGetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapdConfGet",
      requestType = cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest.class,
      responseType = cloud.lazycat.apis.sys.Snapd.SnapdConf.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest,
      cloud.lazycat.apis.sys.Snapd.SnapdConf> getSnapdConfGetMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest, cloud.lazycat.apis.sys.Snapd.SnapdConf> getSnapdConfGetMethod;
    if ((getSnapdConfGetMethod = SnapdServiceGrpc.getSnapdConfGetMethod) == null) {
      synchronized (SnapdServiceGrpc.class) {
        if ((getSnapdConfGetMethod = SnapdServiceGrpc.getSnapdConfGetMethod) == null) {
          SnapdServiceGrpc.getSnapdConfGetMethod = getSnapdConfGetMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest, cloud.lazycat.apis.sys.Snapd.SnapdConf>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapdConfGet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdConf.getDefaultInstance()))
              .setSchemaDescriptor(new SnapdServiceMethodDescriptorSupplier("SnapdConfGet"))
              .build();
        }
      }
    }
    return getSnapdConfGetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest,
      com.google.protobuf.Empty> getSnapdConfSetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapdConfSet",
      requestType = cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest,
      com.google.protobuf.Empty> getSnapdConfSetMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest, com.google.protobuf.Empty> getSnapdConfSetMethod;
    if ((getSnapdConfSetMethod = SnapdServiceGrpc.getSnapdConfSetMethod) == null) {
      synchronized (SnapdServiceGrpc.class) {
        if ((getSnapdConfSetMethod = SnapdServiceGrpc.getSnapdConfSetMethod) == null) {
          SnapdServiceGrpc.getSnapdConfSetMethod = getSnapdConfSetMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapdConfSet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new SnapdServiceMethodDescriptorSupplier("SnapdConfSet"))
              .build();
        }
      }
    }
    return getSnapdConfSetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest,
      cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse> getSnapdSnapListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapdSnapList",
      requestType = cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest.class,
      responseType = cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest,
      cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse> getSnapdSnapListMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest, cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse> getSnapdSnapListMethod;
    if ((getSnapdSnapListMethod = SnapdServiceGrpc.getSnapdSnapListMethod) == null) {
      synchronized (SnapdServiceGrpc.class) {
        if ((getSnapdSnapListMethod = SnapdServiceGrpc.getSnapdSnapListMethod) == null) {
          SnapdServiceGrpc.getSnapdSnapListMethod = getSnapdSnapListMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest, cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapdSnapList"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SnapdServiceMethodDescriptorSupplier("SnapdSnapList"))
              .build();
        }
      }
    }
    return getSnapdSnapListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
      com.google.protobuf.Empty> getSnapdSnapAddMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapdSnapAdd",
      requestType = cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
      com.google.protobuf.Empty> getSnapdSnapAddMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest, com.google.protobuf.Empty> getSnapdSnapAddMethod;
    if ((getSnapdSnapAddMethod = SnapdServiceGrpc.getSnapdSnapAddMethod) == null) {
      synchronized (SnapdServiceGrpc.class) {
        if ((getSnapdSnapAddMethod = SnapdServiceGrpc.getSnapdSnapAddMethod) == null) {
          SnapdServiceGrpc.getSnapdSnapAddMethod = getSnapdSnapAddMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapdSnapAdd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new SnapdServiceMethodDescriptorSupplier("SnapdSnapAdd"))
              .build();
        }
      }
    }
    return getSnapdSnapAddMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
      com.google.protobuf.Empty> getSnapdSnapDelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapdSnapDel",
      requestType = cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
      com.google.protobuf.Empty> getSnapdSnapDelMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest, com.google.protobuf.Empty> getSnapdSnapDelMethod;
    if ((getSnapdSnapDelMethod = SnapdServiceGrpc.getSnapdSnapDelMethod) == null) {
      synchronized (SnapdServiceGrpc.class) {
        if ((getSnapdSnapDelMethod = SnapdServiceGrpc.getSnapdSnapDelMethod) == null) {
          SnapdServiceGrpc.getSnapdSnapDelMethod = getSnapdSnapDelMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapdSnapDel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new SnapdServiceMethodDescriptorSupplier("SnapdSnapDel"))
              .build();
        }
      }
    }
    return getSnapdSnapDelMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
      com.google.protobuf.Empty> getSnapdSnapRestoreMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapdSnapRestore",
      requestType = cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
      com.google.protobuf.Empty> getSnapdSnapRestoreMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest, com.google.protobuf.Empty> getSnapdSnapRestoreMethod;
    if ((getSnapdSnapRestoreMethod = SnapdServiceGrpc.getSnapdSnapRestoreMethod) == null) {
      synchronized (SnapdServiceGrpc.class) {
        if ((getSnapdSnapRestoreMethod = SnapdServiceGrpc.getSnapdSnapRestoreMethod) == null) {
          SnapdServiceGrpc.getSnapdSnapRestoreMethod = getSnapdSnapRestoreMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapdSnapRestore"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new SnapdServiceMethodDescriptorSupplier("SnapdSnapRestore"))
              .build();
        }
      }
    }
    return getSnapdSnapRestoreMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SnapdServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SnapdServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SnapdServiceStub>() {
        @java.lang.Override
        public SnapdServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SnapdServiceStub(channel, callOptions);
        }
      };
    return SnapdServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SnapdServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SnapdServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SnapdServiceBlockingStub>() {
        @java.lang.Override
        public SnapdServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SnapdServiceBlockingStub(channel, callOptions);
        }
      };
    return SnapdServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SnapdServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SnapdServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SnapdServiceFutureStub>() {
        @java.lang.Override
        public SnapdServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SnapdServiceFutureStub(channel, callOptions);
        }
      };
    return SnapdServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class SnapdServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 启用备份
     * </pre>
     */
    public void snapdEnable(cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapdEnableMethod(), responseObserver);
    }

    /**
     * <pre>
     * 禁用备份
     * </pre>
     */
    public void snapdDisable(cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapdDisableMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取快照备份配置
     * </pre>
     */
    public void snapdConfGet(cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Snapd.SnapdConf> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapdConfGetMethod(), responseObserver);
    }

    /**
     * <pre>
     * 修改快照备份配置
     * </pre>
     */
    public void snapdConfSet(cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapdConfSetMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列举快照信息
     * </pre>
     */
    public void snapdSnapList(cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapdSnapListMethod(), responseObserver);
    }

    /**
     * <pre>
     * 手动创建快照/备份。以该方式创建的快照/备份会被自动策略忽略
     * </pre>
     */
    public void snapdSnapAdd(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapdSnapAddMethod(), responseObserver);
    }

    /**
     * <pre>
     * 手动删除快照。目标快照可以是通过自动策略创建的，也可以是手动创建的
     * </pre>
     */
    public void snapdSnapDel(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapdSnapDelMethod(), responseObserver);
    }

    /**
     * <pre>
     * 将数据回滚到指定快照
     * </pre>
     */
    public void snapdSnapRestore(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapdSnapRestoreMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSnapdEnableMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPD_ENABLE)))
          .addMethod(
            getSnapdDisableMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPD_DISABLE)))
          .addMethod(
            getSnapdConfGetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest,
                cloud.lazycat.apis.sys.Snapd.SnapdConf>(
                  this, METHODID_SNAPD_CONF_GET)))
          .addMethod(
            getSnapdConfSetMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPD_CONF_SET)))
          .addMethod(
            getSnapdSnapListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest,
                cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse>(
                  this, METHODID_SNAPD_SNAP_LIST)))
          .addMethod(
            getSnapdSnapAddMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPD_SNAP_ADD)))
          .addMethod(
            getSnapdSnapDelMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPD_SNAP_DEL)))
          .addMethod(
            getSnapdSnapRestoreMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPD_SNAP_RESTORE)))
          .build();
    }
  }

  /**
   */
  public static final class SnapdServiceStub extends io.grpc.stub.AbstractAsyncStub<SnapdServiceStub> {
    private SnapdServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SnapdServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SnapdServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 启用备份
     * </pre>
     */
    public void snapdEnable(cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapdEnableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 禁用备份
     * </pre>
     */
    public void snapdDisable(cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapdDisableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取快照备份配置
     * </pre>
     */
    public void snapdConfGet(cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Snapd.SnapdConf> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapdConfGetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 修改快照备份配置
     * </pre>
     */
    public void snapdConfSet(cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapdConfSetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列举快照信息
     * </pre>
     */
    public void snapdSnapList(cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapdSnapListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 手动创建快照/备份。以该方式创建的快照/备份会被自动策略忽略
     * </pre>
     */
    public void snapdSnapAdd(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapdSnapAddMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 手动删除快照。目标快照可以是通过自动策略创建的，也可以是手动创建的
     * </pre>
     */
    public void snapdSnapDel(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapdSnapDelMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 将数据回滚到指定快照
     * </pre>
     */
    public void snapdSnapRestore(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapdSnapRestoreMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class SnapdServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<SnapdServiceBlockingStub> {
    private SnapdServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SnapdServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SnapdServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 启用备份
     * </pre>
     */
    public com.google.protobuf.Empty snapdEnable(cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapdEnableMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 禁用备份
     * </pre>
     */
    public com.google.protobuf.Empty snapdDisable(cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapdDisableMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取快照备份配置
     * </pre>
     */
    public cloud.lazycat.apis.sys.Snapd.SnapdConf snapdConfGet(cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapdConfGetMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 修改快照备份配置
     * </pre>
     */
    public com.google.protobuf.Empty snapdConfSet(cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapdConfSetMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列举快照信息
     * </pre>
     */
    public cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse snapdSnapList(cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapdSnapListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 手动创建快照/备份。以该方式创建的快照/备份会被自动策略忽略
     * </pre>
     */
    public com.google.protobuf.Empty snapdSnapAdd(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapdSnapAddMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 手动删除快照。目标快照可以是通过自动策略创建的，也可以是手动创建的
     * </pre>
     */
    public com.google.protobuf.Empty snapdSnapDel(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapdSnapDelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 将数据回滚到指定快照
     * </pre>
     */
    public com.google.protobuf.Empty snapdSnapRestore(cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapdSnapRestoreMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class SnapdServiceFutureStub extends io.grpc.stub.AbstractFutureStub<SnapdServiceFutureStub> {
    private SnapdServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SnapdServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SnapdServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 启用备份
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapdEnable(
        cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapdEnableMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 禁用备份
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapdDisable(
        cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapdDisableMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取快照备份配置
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Snapd.SnapdConf> snapdConfGet(
        cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapdConfGetMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 修改快照备份配置
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapdConfSet(
        cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapdConfSetMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列举快照信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse> snapdSnapList(
        cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapdSnapListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 手动创建快照/备份。以该方式创建的快照/备份会被自动策略忽略
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapdSnapAdd(
        cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapdSnapAddMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 手动删除快照。目标快照可以是通过自动策略创建的，也可以是手动创建的
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapdSnapDel(
        cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapdSnapDelMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 将数据回滚到指定快照
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapdSnapRestore(
        cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapdSnapRestoreMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SNAPD_ENABLE = 0;
  private static final int METHODID_SNAPD_DISABLE = 1;
  private static final int METHODID_SNAPD_CONF_GET = 2;
  private static final int METHODID_SNAPD_CONF_SET = 3;
  private static final int METHODID_SNAPD_SNAP_LIST = 4;
  private static final int METHODID_SNAPD_SNAP_ADD = 5;
  private static final int METHODID_SNAPD_SNAP_DEL = 6;
  private static final int METHODID_SNAPD_SNAP_RESTORE = 7;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final SnapdServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(SnapdServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SNAPD_ENABLE:
          serviceImpl.snapdEnable((cloud.lazycat.apis.sys.Snapd.SnapdEnableRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPD_DISABLE:
          serviceImpl.snapdDisable((cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPD_CONF_GET:
          serviceImpl.snapdConfGet((cloud.lazycat.apis.sys.Snapd.SnapdTargetRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Snapd.SnapdConf>) responseObserver);
          break;
        case METHODID_SNAPD_CONF_SET:
          serviceImpl.snapdConfSet((cloud.lazycat.apis.sys.Snapd.SnapdConfSetRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPD_SNAP_LIST:
          serviceImpl.snapdSnapList((cloud.lazycat.apis.sys.Snapd.SnapdListSnapRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.Snapd.SnapdListSnapResponse>) responseObserver);
          break;
        case METHODID_SNAPD_SNAP_ADD:
          serviceImpl.snapdSnapAdd((cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPD_SNAP_DEL:
          serviceImpl.snapdSnapDel((cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPD_SNAP_RESTORE:
          serviceImpl.snapdSnapRestore((cloud.lazycat.apis.sys.Snapd.SnapdSnapRequest) request,
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

  private static abstract class SnapdServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    SnapdServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.Snapd.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("SnapdService");
    }
  }

  private static final class SnapdServiceFileDescriptorSupplier
      extends SnapdServiceBaseDescriptorSupplier {
    SnapdServiceFileDescriptorSupplier() {}
  }

  private static final class SnapdServiceMethodDescriptorSupplier
      extends SnapdServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    SnapdServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (SnapdServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new SnapdServiceFileDescriptorSupplier())
              .addMethod(getSnapdEnableMethod())
              .addMethod(getSnapdDisableMethod())
              .addMethod(getSnapdConfGetMethod())
              .addMethod(getSnapdConfSetMethod())
              .addMethod(getSnapdSnapListMethod())
              .addMethod(getSnapdSnapAddMethod())
              .addMethod(getSnapdSnapDelMethod())
              .addMethod(getSnapdSnapRestoreMethod())
              .build();
        }
      }
    }
    return result;
  }
}
