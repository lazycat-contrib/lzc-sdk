package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/OS_upgrader.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class OSUpgradeServiceGrpc {

  private OSUpgradeServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.OSUpgradeService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo> getLocalMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Local",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo> getLocalMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo> getLocalMethod;
    if ((getLocalMethod = OSUpgradeServiceGrpc.getLocalMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getLocalMethod = OSUpgradeServiceGrpc.getLocalMethod) == null) {
          OSUpgradeServiceGrpc.getLocalMethod = getLocalMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Local"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Local"))
              .build();
        }
      }
    }
    return getLocalMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo> getRemoteMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Remote",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo> getRemoteMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo> getRemoteMethod;
    if ((getRemoteMethod = OSUpgradeServiceGrpc.getRemoteMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getRemoteMethod = OSUpgradeServiceGrpc.getRemoteMethod) == null) {
          OSUpgradeServiceGrpc.getRemoteMethod = getRemoteMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Remote"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Remote"))
              .build();
        }
      }
    }
    return getRemoteMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSUpgrader.SystemVersion,
      com.google.protobuf.Empty> getSelectMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Select",
      requestType = cloud.lazycat.apis.sys.OSUpgrader.SystemVersion.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSUpgrader.SystemVersion,
      com.google.protobuf.Empty> getSelectMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSUpgrader.SystemVersion, com.google.protobuf.Empty> getSelectMethod;
    if ((getSelectMethod = OSUpgradeServiceGrpc.getSelectMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getSelectMethod = OSUpgradeServiceGrpc.getSelectMethod) == null) {
          OSUpgradeServiceGrpc.getSelectMethod = getSelectMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSUpgrader.SystemVersion, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Select"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSUpgrader.SystemVersion.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Select"))
              .build();
        }
      }
    }
    return getSelectMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getStartMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Start",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getStartMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getStartMethod;
    if ((getStartMethod = OSUpgradeServiceGrpc.getStartMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getStartMethod = OSUpgradeServiceGrpc.getStartMethod) == null) {
          OSUpgradeServiceGrpc.getStartMethod = getStartMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Start"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Start"))
              .build();
        }
      }
    }
    return getStartMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getPauseMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Pause",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getPauseMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getPauseMethod;
    if ((getPauseMethod = OSUpgradeServiceGrpc.getPauseMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getPauseMethod = OSUpgradeServiceGrpc.getPauseMethod) == null) {
          OSUpgradeServiceGrpc.getPauseMethod = getPauseMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Pause"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Pause"))
              .build();
        }
      }
    }
    return getPauseMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo> getProgressMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Progress",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo> getProgressMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo> getProgressMethod;
    if ((getProgressMethod = OSUpgradeServiceGrpc.getProgressMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getProgressMethod = OSUpgradeServiceGrpc.getProgressMethod) == null) {
          OSUpgradeServiceGrpc.getProgressMethod = getProgressMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Progress"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Progress"))
              .build();
        }
      }
    }
    return getProgressMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getInstallMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Install",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getInstallMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getInstallMethod;
    if ((getInstallMethod = OSUpgradeServiceGrpc.getInstallMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getInstallMethod = OSUpgradeServiceGrpc.getInstallMethod) == null) {
          OSUpgradeServiceGrpc.getInstallMethod = getInstallMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Install"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Install"))
              .build();
        }
      }
    }
    return getInstallMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSUpgrader.SystemVersion,
      com.google.protobuf.Empty> getSwitchMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Switch",
      requestType = cloud.lazycat.apis.sys.OSUpgrader.SystemVersion.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSUpgrader.SystemVersion,
      com.google.protobuf.Empty> getSwitchMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSUpgrader.SystemVersion, com.google.protobuf.Empty> getSwitchMethod;
    if ((getSwitchMethod = OSUpgradeServiceGrpc.getSwitchMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getSwitchMethod = OSUpgradeServiceGrpc.getSwitchMethod) == null) {
          OSUpgradeServiceGrpc.getSwitchMethod = getSwitchMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSUpgrader.SystemVersion, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Switch"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSUpgrader.SystemVersion.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Switch"))
              .build();
        }
      }
    }
    return getSwitchMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getPruneMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Prune",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getPruneMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getPruneMethod;
    if ((getPruneMethod = OSUpgradeServiceGrpc.getPruneMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getPruneMethod = OSUpgradeServiceGrpc.getPruneMethod) == null) {
          OSUpgradeServiceGrpc.getPruneMethod = getPruneMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Prune"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Prune"))
              .build();
        }
      }
    }
    return getPruneMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getRebootMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Reboot",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getRebootMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getRebootMethod;
    if ((getRebootMethod = OSUpgradeServiceGrpc.getRebootMethod) == null) {
      synchronized (OSUpgradeServiceGrpc.class) {
        if ((getRebootMethod = OSUpgradeServiceGrpc.getRebootMethod) == null) {
          OSUpgradeServiceGrpc.getRebootMethod = getRebootMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Reboot"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSUpgradeServiceMethodDescriptorSupplier("Reboot"))
              .build();
        }
      }
    }
    return getRebootMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static OSUpgradeServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OSUpgradeServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OSUpgradeServiceStub>() {
        @java.lang.Override
        public OSUpgradeServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OSUpgradeServiceStub(channel, callOptions);
        }
      };
    return OSUpgradeServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static OSUpgradeServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OSUpgradeServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OSUpgradeServiceBlockingStub>() {
        @java.lang.Override
        public OSUpgradeServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OSUpgradeServiceBlockingStub(channel, callOptions);
        }
      };
    return OSUpgradeServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static OSUpgradeServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OSUpgradeServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OSUpgradeServiceFutureStub>() {
        @java.lang.Override
        public OSUpgradeServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OSUpgradeServiceFutureStub(channel, callOptions);
        }
      };
    return OSUpgradeServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class OSUpgradeServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 获取当前系统的版本状态
     * </pre>
     */
    public void local(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLocalMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取远程系统的版本状态
     * </pre>
     */
    public void remote(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRemoteMethod(), responseObserver);
    }

    /**
     * <pre>
     * 选择远程某个版本，获取到大小准备下载
     * </pre>
     */
    public void select(cloud.lazycat.apis.sys.OSUpgrader.SystemVersion request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSelectMethod(), responseObserver);
    }

    /**
     * <pre>
     * 开始下载
     * </pre>
     */
    public void start(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStartMethod(), responseObserver);
    }

    /**
     * <pre>
     * 暂停下载
     * </pre>
     */
    public void pause(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPauseMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取下载进度
     * </pre>
     */
    public void progress(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getProgressMethod(), responseObserver);
    }

    /**
     * <pre>
     * 安装当前已下好的版本
     * </pre>
     */
    public void install(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getInstallMethod(), responseObserver);
    }

    /**
     * <pre>
     * 切换到某个版本（需要重启生效）
     * </pre>
     */
    public void switch_(cloud.lazycat.apis.sys.OSUpgrader.SystemVersion request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSwitchMethod(), responseObserver);
    }

    /**
     * <pre>
     * 清理（阻塞）
     * </pre>
     */
    public void prune(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPruneMethod(), responseObserver);
    }

    /**
     * <pre>
     * 重启
     * </pre>
     */
    public void reboot(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRebootMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getLocalMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo>(
                  this, METHODID_LOCAL)))
          .addMethod(
            getRemoteMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo>(
                  this, METHODID_REMOTE)))
          .addMethod(
            getSelectMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSUpgrader.SystemVersion,
                com.google.protobuf.Empty>(
                  this, METHODID_SELECT)))
          .addMethod(
            getStartMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_START)))
          .addMethod(
            getPauseMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_PAUSE)))
          .addMethod(
            getProgressMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo>(
                  this, METHODID_PROGRESS)))
          .addMethod(
            getInstallMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_INSTALL)))
          .addMethod(
            getSwitchMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSUpgrader.SystemVersion,
                com.google.protobuf.Empty>(
                  this, METHODID_SWITCH)))
          .addMethod(
            getPruneMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_PRUNE)))
          .addMethod(
            getRebootMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_REBOOT)))
          .build();
    }
  }

  /**
   */
  public static final class OSUpgradeServiceStub extends io.grpc.stub.AbstractAsyncStub<OSUpgradeServiceStub> {
    private OSUpgradeServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OSUpgradeServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OSUpgradeServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 获取当前系统的版本状态
     * </pre>
     */
    public void local(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLocalMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取远程系统的版本状态
     * </pre>
     */
    public void remote(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRemoteMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 选择远程某个版本，获取到大小准备下载
     * </pre>
     */
    public void select(cloud.lazycat.apis.sys.OSUpgrader.SystemVersion request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSelectMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 开始下载
     * </pre>
     */
    public void start(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStartMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 暂停下载
     * </pre>
     */
    public void pause(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPauseMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取下载进度
     * </pre>
     */
    public void progress(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getProgressMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 安装当前已下好的版本
     * </pre>
     */
    public void install(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getInstallMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 切换到某个版本（需要重启生效）
     * </pre>
     */
    public void switch_(cloud.lazycat.apis.sys.OSUpgrader.SystemVersion request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSwitchMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 清理（阻塞）
     * </pre>
     */
    public void prune(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPruneMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 重启
     * </pre>
     */
    public void reboot(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRebootMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class OSUpgradeServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<OSUpgradeServiceBlockingStub> {
    private OSUpgradeServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OSUpgradeServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OSUpgradeServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 获取当前系统的版本状态
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo local(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLocalMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取远程系统的版本状态
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo remote(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRemoteMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 选择远程某个版本，获取到大小准备下载
     * </pre>
     */
    public com.google.protobuf.Empty select(cloud.lazycat.apis.sys.OSUpgrader.SystemVersion request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSelectMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 开始下载
     * </pre>
     */
    public com.google.protobuf.Empty start(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStartMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 暂停下载
     * </pre>
     */
    public com.google.protobuf.Empty pause(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPauseMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取下载进度
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo progress(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getProgressMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 安装当前已下好的版本
     * </pre>
     */
    public com.google.protobuf.Empty install(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getInstallMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 切换到某个版本（需要重启生效）
     * </pre>
     */
    public com.google.protobuf.Empty switch_(cloud.lazycat.apis.sys.OSUpgrader.SystemVersion request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSwitchMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 清理（阻塞）
     * </pre>
     */
    public com.google.protobuf.Empty prune(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPruneMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 重启
     * </pre>
     */
    public com.google.protobuf.Empty reboot(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRebootMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class OSUpgradeServiceFutureStub extends io.grpc.stub.AbstractFutureStub<OSUpgradeServiceFutureStub> {
    private OSUpgradeServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OSUpgradeServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OSUpgradeServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 获取当前系统的版本状态
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo> local(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLocalMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取远程系统的版本状态
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo> remote(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRemoteMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 选择远程某个版本，获取到大小准备下载
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> select(
        cloud.lazycat.apis.sys.OSUpgrader.SystemVersion request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSelectMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 开始下载
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> start(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStartMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 暂停下载
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> pause(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPauseMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取下载进度
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo> progress(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getProgressMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 安装当前已下好的版本
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> install(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getInstallMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 切换到某个版本（需要重启生效）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> switch_(
        cloud.lazycat.apis.sys.OSUpgrader.SystemVersion request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSwitchMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 清理（阻塞）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> prune(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPruneMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 重启
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> reboot(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRebootMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LOCAL = 0;
  private static final int METHODID_REMOTE = 1;
  private static final int METHODID_SELECT = 2;
  private static final int METHODID_START = 3;
  private static final int METHODID_PAUSE = 4;
  private static final int METHODID_PROGRESS = 5;
  private static final int METHODID_INSTALL = 6;
  private static final int METHODID_SWITCH = 7;
  private static final int METHODID_PRUNE = 8;
  private static final int METHODID_REBOOT = 9;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final OSUpgradeServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(OSUpgradeServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LOCAL:
          serviceImpl.local((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.LocalSystemVersionInfo>) responseObserver);
          break;
        case METHODID_REMOTE:
          serviceImpl.remote((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.RemoteSystemVersionInfo>) responseObserver);
          break;
        case METHODID_SELECT:
          serviceImpl.select((cloud.lazycat.apis.sys.OSUpgrader.SystemVersion) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_START:
          serviceImpl.start((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_PAUSE:
          serviceImpl.pause((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_PROGRESS:
          serviceImpl.progress((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSUpgrader.UpgradeProgressInfo>) responseObserver);
          break;
        case METHODID_INSTALL:
          serviceImpl.install((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SWITCH:
          serviceImpl.switch_((cloud.lazycat.apis.sys.OSUpgrader.SystemVersion) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_PRUNE:
          serviceImpl.prune((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_REBOOT:
          serviceImpl.reboot((com.google.protobuf.Empty) request,
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

  private static abstract class OSUpgradeServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    OSUpgradeServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.OSUpgrader.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("OSUpgradeService");
    }
  }

  private static final class OSUpgradeServiceFileDescriptorSupplier
      extends OSUpgradeServiceBaseDescriptorSupplier {
    OSUpgradeServiceFileDescriptorSupplier() {}
  }

  private static final class OSUpgradeServiceMethodDescriptorSupplier
      extends OSUpgradeServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    OSUpgradeServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (OSUpgradeServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new OSUpgradeServiceFileDescriptorSupplier())
              .addMethod(getLocalMethod())
              .addMethod(getRemoteMethod())
              .addMethod(getSelectMethod())
              .addMethod(getStartMethod())
              .addMethod(getPauseMethod())
              .addMethod(getProgressMethod())
              .addMethod(getInstallMethod())
              .addMethod(getSwitchMethod())
              .addMethod(getPruneMethod())
              .addMethod(getRebootMethod())
              .build();
        }
      }
    }
    return result;
  }
}
