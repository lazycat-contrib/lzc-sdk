package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/network_manager.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class NetworkManagerGrpc {

  private NetworkManagerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.NetworkManager";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo> getStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Status",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo> getStatusMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo> getStatusMethod;
    if ((getStatusMethod = NetworkManagerGrpc.getStatusMethod) == null) {
      synchronized (NetworkManagerGrpc.class) {
        if ((getStatusMethod = NetworkManagerGrpc.getStatusMethod) == null) {
          NetworkManagerGrpc.getStatusMethod = getStatusMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Status"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo.getDefaultInstance()))
              .setSchemaDescriptor(new NetworkManagerMethodDescriptorSupplier("Status"))
              .build();
        }
      }
    }
    return getStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getWifiScanMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "WifiScan",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getWifiScanMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getWifiScanMethod;
    if ((getWifiScanMethod = NetworkManagerGrpc.getWifiScanMethod) == null) {
      synchronized (NetworkManagerGrpc.class) {
        if ((getWifiScanMethod = NetworkManagerGrpc.getWifiScanMethod) == null) {
          NetworkManagerGrpc.getWifiScanMethod = getWifiScanMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "WifiScan"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new NetworkManagerMethodDescriptorSupplier("WifiScan"))
              .build();
        }
      }
    }
    return getWifiScanMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList> getWifiListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "WifiList",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList> getWifiListMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList> getWifiListMethod;
    if ((getWifiListMethod = NetworkManagerGrpc.getWifiListMethod) == null) {
      synchronized (NetworkManagerGrpc.class) {
        if ((getWifiListMethod = NetworkManagerGrpc.getWifiListMethod) == null) {
          NetworkManagerGrpc.getWifiListMethod = getWifiListMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "WifiList"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList.getDefaultInstance()))
              .setSchemaDescriptor(new NetworkManagerMethodDescriptorSupplier("WifiList"))
              .build();
        }
      }
    }
    return getWifiListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo,
      cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> getWifiConnectMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "WifiConnect",
      requestType = cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo.class,
      responseType = cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo,
      cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> getWifiConnectMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo, cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> getWifiConnectMethod;
    if ((getWifiConnectMethod = NetworkManagerGrpc.getWifiConnectMethod) == null) {
      synchronized (NetworkManagerGrpc.class) {
        if ((getWifiConnectMethod = NetworkManagerGrpc.getWifiConnectMethod) == null) {
          NetworkManagerGrpc.getWifiConnectMethod = getWifiConnectMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo, cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "WifiConnect"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply.getDefaultInstance()))
              .setSchemaDescriptor(new NetworkManagerMethodDescriptorSupplier("WifiConnect"))
              .build();
        }
      }
    }
    return getWifiConnectMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo,
      cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> getWifiConfigAddMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "WifiConfigAdd",
      requestType = cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo.class,
      responseType = cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo,
      cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> getWifiConfigAddMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo, cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> getWifiConfigAddMethod;
    if ((getWifiConfigAddMethod = NetworkManagerGrpc.getWifiConfigAddMethod) == null) {
      synchronized (NetworkManagerGrpc.class) {
        if ((getWifiConfigAddMethod = NetworkManagerGrpc.getWifiConfigAddMethod) == null) {
          NetworkManagerGrpc.getWifiConfigAddMethod = getWifiConfigAddMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo, cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "WifiConfigAdd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply.getDefaultInstance()))
              .setSchemaDescriptor(new NetworkManagerMethodDescriptorSupplier("WifiConfigAdd"))
              .build();
        }
      }
    }
    return getWifiConfigAddMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static NetworkManagerStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<NetworkManagerStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<NetworkManagerStub>() {
        @java.lang.Override
        public NetworkManagerStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new NetworkManagerStub(channel, callOptions);
        }
      };
    return NetworkManagerStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static NetworkManagerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<NetworkManagerBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<NetworkManagerBlockingStub>() {
        @java.lang.Override
        public NetworkManagerBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new NetworkManagerBlockingStub(channel, callOptions);
        }
      };
    return NetworkManagerBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static NetworkManagerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<NetworkManagerFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<NetworkManagerFutureStub>() {
        @java.lang.Override
        public NetworkManagerFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new NetworkManagerFutureStub(channel, callOptions);
        }
      };
    return NetworkManagerFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class NetworkManagerImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 获取网络设备的状态（是否已连接，连接了哪个）
     * </pre>
     */
    public void status(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStatusMethod(), responseObserver);
    }

    /**
     * <pre>
     * Scan 扫描附近wifi热点信息，扫描结果在内部缓存里（阻塞，可能耗费数秒）
     * </pre>
     */
    public void wifiScan(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getWifiScanMethod(), responseObserver);
    }

    /**
     * <pre>
     * List 列出内部缓存中的 wifi 列表
     * </pre>
     */
    public void wifiList(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getWifiListMethod(), responseObserver);
    }

    /**
     * <pre>
     * 连接一个 wifi 热点
     *   连接成功后会把其它配置设为非自动连接，把自己设为自动
     *   连接失败会删除已保存的配置
     * </pre>
     */
    public void wifiConnect(cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getWifiConnectMethod(), responseObserver);
    }

    /**
     * <pre>
     * 手动添加和连接一个 wifi 热点配置（用于连接隐藏网络）
     * </pre>
     */
    public void wifiConfigAdd(cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getWifiConfigAddMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getStatusMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo>(
                  this, METHODID_STATUS)))
          .addMethod(
            getWifiScanMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_WIFI_SCAN)))
          .addMethod(
            getWifiListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList>(
                  this, METHODID_WIFI_LIST)))
          .addMethod(
            getWifiConnectMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo,
                cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply>(
                  this, METHODID_WIFI_CONNECT)))
          .addMethod(
            getWifiConfigAddMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo,
                cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply>(
                  this, METHODID_WIFI_CONFIG_ADD)))
          .build();
    }
  }

  /**
   */
  public static final class NetworkManagerStub extends io.grpc.stub.AbstractAsyncStub<NetworkManagerStub> {
    private NetworkManagerStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NetworkManagerStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new NetworkManagerStub(channel, callOptions);
    }

    /**
     * <pre>
     * 获取网络设备的状态（是否已连接，连接了哪个）
     * </pre>
     */
    public void status(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Scan 扫描附近wifi热点信息，扫描结果在内部缓存里（阻塞，可能耗费数秒）
     * </pre>
     */
    public void wifiScan(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getWifiScanMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * List 列出内部缓存中的 wifi 列表
     * </pre>
     */
    public void wifiList(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getWifiListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 连接一个 wifi 热点
     *   连接成功后会把其它配置设为非自动连接，把自己设为自动
     *   连接失败会删除已保存的配置
     * </pre>
     */
    public void wifiConnect(cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getWifiConnectMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 手动添加和连接一个 wifi 热点配置（用于连接隐藏网络）
     * </pre>
     */
    public void wifiConfigAdd(cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getWifiConfigAddMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class NetworkManagerBlockingStub extends io.grpc.stub.AbstractBlockingStub<NetworkManagerBlockingStub> {
    private NetworkManagerBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NetworkManagerBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new NetworkManagerBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 获取网络设备的状态（是否已连接，连接了哪个）
     * </pre>
     */
    public cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo status(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Scan 扫描附近wifi热点信息，扫描结果在内部缓存里（阻塞，可能耗费数秒）
     * </pre>
     */
    public com.google.protobuf.Empty wifiScan(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getWifiScanMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * List 列出内部缓存中的 wifi 列表
     * </pre>
     */
    public cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList wifiList(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getWifiListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 连接一个 wifi 热点
     *   连接成功后会把其它配置设为非自动连接，把自己设为自动
     *   连接失败会删除已保存的配置
     * </pre>
     */
    public cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply wifiConnect(cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getWifiConnectMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 手动添加和连接一个 wifi 热点配置（用于连接隐藏网络）
     * </pre>
     */
    public cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply wifiConfigAdd(cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getWifiConfigAddMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class NetworkManagerFutureStub extends io.grpc.stub.AbstractFutureStub<NetworkManagerFutureStub> {
    private NetworkManagerFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected NetworkManagerFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new NetworkManagerFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 获取网络设备的状态（是否已连接，连接了哪个）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo> status(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStatusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Scan 扫描附近wifi热点信息，扫描结果在内部缓存里（阻塞，可能耗费数秒）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> wifiScan(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getWifiScanMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * List 列出内部缓存中的 wifi 列表
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList> wifiList(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getWifiListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 连接一个 wifi 热点
     *   连接成功后会把其它配置设为非自动连接，把自己设为自动
     *   连接失败会删除已保存的配置
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> wifiConnect(
        cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getWifiConnectMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 手动添加和连接一个 wifi 热点配置（用于连接隐藏网络）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply> wifiConfigAdd(
        cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getWifiConfigAddMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_STATUS = 0;
  private static final int METHODID_WIFI_SCAN = 1;
  private static final int METHODID_WIFI_LIST = 2;
  private static final int METHODID_WIFI_CONNECT = 3;
  private static final int METHODID_WIFI_CONFIG_ADD = 4;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final NetworkManagerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(NetworkManagerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_STATUS:
          serviceImpl.status((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.NetworkDeviceStatusInfo>) responseObserver);
          break;
        case METHODID_WIFI_SCAN:
          serviceImpl.wifiScan((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_WIFI_LIST:
          serviceImpl.wifiList((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.AccessPointInfoList>) responseObserver);
          break;
        case METHODID_WIFI_CONNECT:
          serviceImpl.wifiConnect((cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectInfo) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply>) responseObserver);
          break;
        case METHODID_WIFI_CONFIG_ADD:
          serviceImpl.wifiConfigAdd((cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConfigInfo) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.NetworkManagerOuterClass.WifiConnectReply>) responseObserver);
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

  private static abstract class NetworkManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    NetworkManagerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.NetworkManagerOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("NetworkManager");
    }
  }

  private static final class NetworkManagerFileDescriptorSupplier
      extends NetworkManagerBaseDescriptorSupplier {
    NetworkManagerFileDescriptorSupplier() {}
  }

  private static final class NetworkManagerMethodDescriptorSupplier
      extends NetworkManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    NetworkManagerMethodDescriptorSupplier(String methodName) {
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
      synchronized (NetworkManagerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new NetworkManagerFileDescriptorSupplier())
              .addMethod(getStatusMethod())
              .addMethod(getWifiScanMethod())
              .addMethod(getWifiListMethod())
              .addMethod(getWifiConnectMethod())
              .addMethod(getWifiConfigAddMethod())
              .build();
        }
      }
    }
    return result;
  }
}
