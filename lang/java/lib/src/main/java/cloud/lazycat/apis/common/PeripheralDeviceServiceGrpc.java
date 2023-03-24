package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * 外围设备管理服务
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/peripheral_device.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PeripheralDeviceServiceGrpc {

  private PeripheralDeviceServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.PeripheralDeviceService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply> getListFilesystemsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListFilesystems",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply> getListFilesystemsMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply> getListFilesystemsMethod;
    if ((getListFilesystemsMethod = PeripheralDeviceServiceGrpc.getListFilesystemsMethod) == null) {
      synchronized (PeripheralDeviceServiceGrpc.class) {
        if ((getListFilesystemsMethod = PeripheralDeviceServiceGrpc.getListFilesystemsMethod) == null) {
          PeripheralDeviceServiceGrpc.getListFilesystemsMethod = getListFilesystemsMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListFilesystems"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply.getDefaultInstance()))
              .setSchemaDescriptor(new PeripheralDeviceServiceMethodDescriptorSupplier("ListFilesystems"))
              .build();
        }
      }
    }
    return getListFilesystemsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest,
      com.google.protobuf.Empty> getMountDiskMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "MountDisk",
      requestType = cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest,
      com.google.protobuf.Empty> getMountDiskMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest, com.google.protobuf.Empty> getMountDiskMethod;
    if ((getMountDiskMethod = PeripheralDeviceServiceGrpc.getMountDiskMethod) == null) {
      synchronized (PeripheralDeviceServiceGrpc.class) {
        if ((getMountDiskMethod = PeripheralDeviceServiceGrpc.getMountDiskMethod) == null) {
          PeripheralDeviceServiceGrpc.getMountDiskMethod = getMountDiskMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "MountDisk"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PeripheralDeviceServiceMethodDescriptorSupplier("MountDisk"))
              .build();
        }
      }
    }
    return getMountDiskMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest,
      com.google.protobuf.Empty> getMountWebDavMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "MountWebDav",
      requestType = cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest,
      com.google.protobuf.Empty> getMountWebDavMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest, com.google.protobuf.Empty> getMountWebDavMethod;
    if ((getMountWebDavMethod = PeripheralDeviceServiceGrpc.getMountWebDavMethod) == null) {
      synchronized (PeripheralDeviceServiceGrpc.class) {
        if ((getMountWebDavMethod = PeripheralDeviceServiceGrpc.getMountWebDavMethod) == null) {
          PeripheralDeviceServiceGrpc.getMountWebDavMethod = getMountWebDavMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "MountWebDav"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PeripheralDeviceServiceMethodDescriptorSupplier("MountWebDav"))
              .build();
        }
      }
    }
    return getMountWebDavMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest,
      com.google.protobuf.Empty> getUmountFilesystemMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UmountFilesystem",
      requestType = cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest,
      com.google.protobuf.Empty> getUmountFilesystemMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest, com.google.protobuf.Empty> getUmountFilesystemMethod;
    if ((getUmountFilesystemMethod = PeripheralDeviceServiceGrpc.getUmountFilesystemMethod) == null) {
      synchronized (PeripheralDeviceServiceGrpc.class) {
        if ((getUmountFilesystemMethod = PeripheralDeviceServiceGrpc.getUmountFilesystemMethod) == null) {
          PeripheralDeviceServiceGrpc.getUmountFilesystemMethod = getUmountFilesystemMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UmountFilesystem"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PeripheralDeviceServiceMethodDescriptorSupplier("UmountFilesystem"))
              .build();
        }
      }
    }
    return getUmountFilesystemMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest,
      com.google.protobuf.Empty> getMountArchiveMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "MountArchive",
      requestType = cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest,
      com.google.protobuf.Empty> getMountArchiveMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest, com.google.protobuf.Empty> getMountArchiveMethod;
    if ((getMountArchiveMethod = PeripheralDeviceServiceGrpc.getMountArchiveMethod) == null) {
      synchronized (PeripheralDeviceServiceGrpc.class) {
        if ((getMountArchiveMethod = PeripheralDeviceServiceGrpc.getMountArchiveMethod) == null) {
          PeripheralDeviceServiceGrpc.getMountArchiveMethod = getMountArchiveMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "MountArchive"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PeripheralDeviceServiceMethodDescriptorSupplier("MountArchive"))
              .build();
        }
      }
    }
    return getMountArchiveMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PeripheralDeviceServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeripheralDeviceServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeripheralDeviceServiceStub>() {
        @java.lang.Override
        public PeripheralDeviceServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeripheralDeviceServiceStub(channel, callOptions);
        }
      };
    return PeripheralDeviceServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PeripheralDeviceServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeripheralDeviceServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeripheralDeviceServiceBlockingStub>() {
        @java.lang.Override
        public PeripheralDeviceServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeripheralDeviceServiceBlockingStub(channel, callOptions);
        }
      };
    return PeripheralDeviceServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PeripheralDeviceServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PeripheralDeviceServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PeripheralDeviceServiceFutureStub>() {
        @java.lang.Override
        public PeripheralDeviceServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PeripheralDeviceServiceFutureStub(channel, callOptions);
        }
      };
    return PeripheralDeviceServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * 外围设备管理服务
   * </pre>
   */
  public static abstract class PeripheralDeviceServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 列出当前盒子已挂载和可以挂载但未挂载的文件系统
     * </pre>
     */
    public void listFilesystems(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListFilesystemsMethod(), responseObserver);
    }

    /**
     * <pre>
     * 挂载/卸载特定移动磁盘的某个分区到
     * $APPID/lzcapp/run/mnt/media/$partition_uuid 目录上
     * </pre>
     */
    public void mountDisk(cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getMountDiskMethod(), responseObserver);
    }

    /**
     * <pre>
     * 挂载 WebDav 服务到 $APPID/lzcapp/run/mnt/home/$uid 目录下，具体路径可以指定
     * </pre>
     */
    public void mountWebDav(cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getMountWebDavMethod(), responseObserver);
    }

    /**
     * <pre>
     * 通过 uuid 或 mountpoint 卸载文件系统
     * </pre>
     */
    public void umountFilesystem(cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUmountFilesystemMethod(), responseObserver);
    }

    /**
     */
    public void mountArchive(cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getMountArchiveMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListFilesystemsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply>(
                  this, METHODID_LIST_FILESYSTEMS)))
          .addMethod(
            getMountDiskMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_MOUNT_DISK)))
          .addMethod(
            getMountWebDavMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_MOUNT_WEB_DAV)))
          .addMethod(
            getUmountFilesystemMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_UMOUNT_FILESYSTEM)))
          .addMethod(
            getMountArchiveMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_MOUNT_ARCHIVE)))
          .build();
    }
  }

  /**
   * <pre>
   * 外围设备管理服务
   * </pre>
   */
  public static final class PeripheralDeviceServiceStub extends io.grpc.stub.AbstractAsyncStub<PeripheralDeviceServiceStub> {
    private PeripheralDeviceServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeripheralDeviceServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeripheralDeviceServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 列出当前盒子已挂载和可以挂载但未挂载的文件系统
     * </pre>
     */
    public void listFilesystems(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListFilesystemsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 挂载/卸载特定移动磁盘的某个分区到
     * $APPID/lzcapp/run/mnt/media/$partition_uuid 目录上
     * </pre>
     */
    public void mountDisk(cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getMountDiskMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 挂载 WebDav 服务到 $APPID/lzcapp/run/mnt/home/$uid 目录下，具体路径可以指定
     * </pre>
     */
    public void mountWebDav(cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getMountWebDavMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 通过 uuid 或 mountpoint 卸载文件系统
     * </pre>
     */
    public void umountFilesystem(cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUmountFilesystemMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void mountArchive(cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getMountArchiveMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * 外围设备管理服务
   * </pre>
   */
  public static final class PeripheralDeviceServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<PeripheralDeviceServiceBlockingStub> {
    private PeripheralDeviceServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeripheralDeviceServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeripheralDeviceServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 列出当前盒子已挂载和可以挂载但未挂载的文件系统
     * </pre>
     */
    public cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply listFilesystems(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListFilesystemsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 挂载/卸载特定移动磁盘的某个分区到
     * $APPID/lzcapp/run/mnt/media/$partition_uuid 目录上
     * </pre>
     */
    public com.google.protobuf.Empty mountDisk(cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getMountDiskMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 挂载 WebDav 服务到 $APPID/lzcapp/run/mnt/home/$uid 目录下，具体路径可以指定
     * </pre>
     */
    public com.google.protobuf.Empty mountWebDav(cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getMountWebDavMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 通过 uuid 或 mountpoint 卸载文件系统
     * </pre>
     */
    public com.google.protobuf.Empty umountFilesystem(cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUmountFilesystemMethod(), getCallOptions(), request);
    }

    /**
     */
    public java.util.Iterator<com.google.protobuf.Empty> mountArchive(
        cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getMountArchiveMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * 外围设备管理服务
   * </pre>
   */
  public static final class PeripheralDeviceServiceFutureStub extends io.grpc.stub.AbstractFutureStub<PeripheralDeviceServiceFutureStub> {
    private PeripheralDeviceServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PeripheralDeviceServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PeripheralDeviceServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 列出当前盒子已挂载和可以挂载但未挂载的文件系统
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply> listFilesystems(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListFilesystemsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 挂载/卸载特定移动磁盘的某个分区到
     * $APPID/lzcapp/run/mnt/media/$partition_uuid 目录上
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> mountDisk(
        cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getMountDiskMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 挂载 WebDav 服务到 $APPID/lzcapp/run/mnt/home/$uid 目录下，具体路径可以指定
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> mountWebDav(
        cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getMountWebDavMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 通过 uuid 或 mountpoint 卸载文件系统
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> umountFilesystem(
        cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUmountFilesystemMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_FILESYSTEMS = 0;
  private static final int METHODID_MOUNT_DISK = 1;
  private static final int METHODID_MOUNT_WEB_DAV = 2;
  private static final int METHODID_UMOUNT_FILESYSTEM = 3;
  private static final int METHODID_MOUNT_ARCHIVE = 4;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final PeripheralDeviceServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(PeripheralDeviceServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LIST_FILESYSTEMS:
          serviceImpl.listFilesystems((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.PeripheralDevice.ListFilesystemsReply>) responseObserver);
          break;
        case METHODID_MOUNT_DISK:
          serviceImpl.mountDisk((cloud.lazycat.apis.common.PeripheralDevice.MountDiskRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_MOUNT_WEB_DAV:
          serviceImpl.mountWebDav((cloud.lazycat.apis.common.PeripheralDevice.MountWebDavRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_UMOUNT_FILESYSTEM:
          serviceImpl.umountFilesystem((cloud.lazycat.apis.common.PeripheralDevice.UmountFilesystemRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_MOUNT_ARCHIVE:
          serviceImpl.mountArchive((cloud.lazycat.apis.common.PeripheralDevice.MountArchiveRequest) request,
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

  private static abstract class PeripheralDeviceServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PeripheralDeviceServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.PeripheralDevice.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PeripheralDeviceService");
    }
  }

  private static final class PeripheralDeviceServiceFileDescriptorSupplier
      extends PeripheralDeviceServiceBaseDescriptorSupplier {
    PeripheralDeviceServiceFileDescriptorSupplier() {}
  }

  private static final class PeripheralDeviceServiceMethodDescriptorSupplier
      extends PeripheralDeviceServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    PeripheralDeviceServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (PeripheralDeviceServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PeripheralDeviceServiceFileDescriptorSupplier())
              .addMethod(getListFilesystemsMethod())
              .addMethod(getMountDiskMethod())
              .addMethod(getMountWebDavMethod())
              .addMethod(getUmountFilesystemMethod())
              .addMethod(getMountArchiveMethod())
              .build();
        }
      }
    }
    return result;
  }
}
