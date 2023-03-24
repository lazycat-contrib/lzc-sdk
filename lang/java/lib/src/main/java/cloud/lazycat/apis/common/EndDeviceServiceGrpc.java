package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * 懒猫云客户端设备管理服务
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/end_device.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class EndDeviceServiceGrpc {

  private EndDeviceServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.EndDeviceService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest,
      cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply> getListEndDevicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListEndDevices",
      requestType = cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest.class,
      responseType = cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest,
      cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply> getListEndDevicesMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest, cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply> getListEndDevicesMethod;
    if ((getListEndDevicesMethod = EndDeviceServiceGrpc.getListEndDevicesMethod) == null) {
      synchronized (EndDeviceServiceGrpc.class) {
        if ((getListEndDevicesMethod = EndDeviceServiceGrpc.getListEndDevicesMethod) == null) {
          EndDeviceServiceGrpc.getListEndDevicesMethod = getListEndDevicesMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest, cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListEndDevices"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply.getDefaultInstance()))
              .setSchemaDescriptor(new EndDeviceServiceMethodDescriptorSupplier("ListEndDevices"))
              .build();
        }
      }
    }
    return getListEndDevicesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest,
      com.google.protobuf.Empty> getPairAllEndDevicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PairAllEndDevices_",
      requestType = cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest,
      com.google.protobuf.Empty> getPairAllEndDevicesMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest, com.google.protobuf.Empty> getPairAllEndDevicesMethod;
    if ((getPairAllEndDevicesMethod = EndDeviceServiceGrpc.getPairAllEndDevicesMethod) == null) {
      synchronized (EndDeviceServiceGrpc.class) {
        if ((getPairAllEndDevicesMethod = EndDeviceServiceGrpc.getPairAllEndDevicesMethod) == null) {
          EndDeviceServiceGrpc.getPairAllEndDevicesMethod = getPairAllEndDevicesMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PairAllEndDevices_"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new EndDeviceServiceMethodDescriptorSupplier("PairAllEndDevices_"))
              .build();
        }
      }
    }
    return getPairAllEndDevicesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest,
      com.google.protobuf.Empty> getRemoveEndDeviceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RemoveEndDevice",
      requestType = cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest,
      com.google.protobuf.Empty> getRemoveEndDeviceMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest, com.google.protobuf.Empty> getRemoveEndDeviceMethod;
    if ((getRemoveEndDeviceMethod = EndDeviceServiceGrpc.getRemoveEndDeviceMethod) == null) {
      synchronized (EndDeviceServiceGrpc.class) {
        if ((getRemoveEndDeviceMethod = EndDeviceServiceGrpc.getRemoveEndDeviceMethod) == null) {
          EndDeviceServiceGrpc.getRemoveEndDeviceMethod = getRemoveEndDeviceMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RemoveEndDevice"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new EndDeviceServiceMethodDescriptorSupplier("RemoveEndDevice"))
              .build();
        }
      }
    }
    return getRemoveEndDeviceMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static EndDeviceServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EndDeviceServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EndDeviceServiceStub>() {
        @java.lang.Override
        public EndDeviceServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EndDeviceServiceStub(channel, callOptions);
        }
      };
    return EndDeviceServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static EndDeviceServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EndDeviceServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EndDeviceServiceBlockingStub>() {
        @java.lang.Override
        public EndDeviceServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EndDeviceServiceBlockingStub(channel, callOptions);
        }
      };
    return EndDeviceServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static EndDeviceServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EndDeviceServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EndDeviceServiceFutureStub>() {
        @java.lang.Override
        public EndDeviceServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EndDeviceServiceFutureStub(channel, callOptions);
        }
      };
    return EndDeviceServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * 懒猫云客户端设备管理服务
   * </pre>
   */
  public static abstract class EndDeviceServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public void listEndDevices(cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListEndDevicesMethod(), responseObserver);
    }

    /**
     * <pre>
     * 将发起请求的设备与登陆用户的其他所有设备建立其虚拟网络隧道
     * 以便发起请求的浏览器可以绕过盒子直接访问其他节点上的设备API，比如剪贴板、文件拷贝等
     * 后端代码本身就与任意设备建立好了虚拟网络隧道，因此不需要使用此API。
     * 前端代码可以调用browseronly.proto:PairAllDeivce自动设置相关参数。
     * </pre>
     */
    public void pairAllEndDevices(cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPairAllEndDevicesMethod(), responseObserver);
    }

    /**
     * <pre>
     * 移除指定uid下面的指定设备(如果设备id为空，将会移除所有的设备）
     * </pre>
     */
    public void removeEndDevice(cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRemoveEndDeviceMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListEndDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest,
                cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply>(
                  this, METHODID_LIST_END_DEVICES)))
          .addMethod(
            getPairAllEndDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_PAIR_ALL_END_DEVICES_)))
          .addMethod(
            getRemoveEndDeviceMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_REMOVE_END_DEVICE)))
          .build();
    }
  }

  /**
   * <pre>
   * 懒猫云客户端设备管理服务
   * </pre>
   */
  public static final class EndDeviceServiceStub extends io.grpc.stub.AbstractAsyncStub<EndDeviceServiceStub> {
    private EndDeviceServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EndDeviceServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EndDeviceServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public void listEndDevices(cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListEndDevicesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 将发起请求的设备与登陆用户的其他所有设备建立其虚拟网络隧道
     * 以便发起请求的浏览器可以绕过盒子直接访问其他节点上的设备API，比如剪贴板、文件拷贝等
     * 后端代码本身就与任意设备建立好了虚拟网络隧道，因此不需要使用此API。
     * 前端代码可以调用browseronly.proto:PairAllDeivce自动设置相关参数。
     * </pre>
     */
    public void pairAllEndDevices(cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getPairAllEndDevicesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 移除指定uid下面的指定设备(如果设备id为空，将会移除所有的设备）
     * </pre>
     */
    public void removeEndDevice(cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRemoveEndDeviceMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * 懒猫云客户端设备管理服务
   * </pre>
   */
  public static final class EndDeviceServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<EndDeviceServiceBlockingStub> {
    private EndDeviceServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EndDeviceServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EndDeviceServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply listEndDevices(cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListEndDevicesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 将发起请求的设备与登陆用户的其他所有设备建立其虚拟网络隧道
     * 以便发起请求的浏览器可以绕过盒子直接访问其他节点上的设备API，比如剪贴板、文件拷贝等
     * 后端代码本身就与任意设备建立好了虚拟网络隧道，因此不需要使用此API。
     * 前端代码可以调用browseronly.proto:PairAllDeivce自动设置相关参数。
     * </pre>
     */
    public java.util.Iterator<com.google.protobuf.Empty> pairAllEndDevices(
        cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getPairAllEndDevicesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 移除指定uid下面的指定设备(如果设备id为空，将会移除所有的设备）
     * </pre>
     */
    public com.google.protobuf.Empty removeEndDevice(cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRemoveEndDeviceMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * 懒猫云客户端设备管理服务
   * </pre>
   */
  public static final class EndDeviceServiceFutureStub extends io.grpc.stub.AbstractFutureStub<EndDeviceServiceFutureStub> {
    private EndDeviceServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EndDeviceServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EndDeviceServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply> listEndDevices(
        cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListEndDevicesMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 移除指定uid下面的指定设备(如果设备id为空，将会移除所有的设备）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> removeEndDevice(
        cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRemoveEndDeviceMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_END_DEVICES = 0;
  private static final int METHODID_PAIR_ALL_END_DEVICES_ = 1;
  private static final int METHODID_REMOVE_END_DEVICE = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final EndDeviceServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(EndDeviceServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LIST_END_DEVICES:
          serviceImpl.listEndDevices((cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.EndDeviceOuterClass.ListEndDeviceReply>) responseObserver);
          break;
        case METHODID_PAIR_ALL_END_DEVICES_:
          serviceImpl.pairAllEndDevices((cloud.lazycat.apis.common.EndDeviceOuterClass.PairEndDeviceRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_REMOVE_END_DEVICE:
          serviceImpl.removeEndDevice((cloud.lazycat.apis.common.EndDeviceOuterClass.RemoveEndDeviceRequest) request,
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

  private static abstract class EndDeviceServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    EndDeviceServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.EndDeviceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("EndDeviceService");
    }
  }

  private static final class EndDeviceServiceFileDescriptorSupplier
      extends EndDeviceServiceBaseDescriptorSupplier {
    EndDeviceServiceFileDescriptorSupplier() {}
  }

  private static final class EndDeviceServiceMethodDescriptorSupplier
      extends EndDeviceServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    EndDeviceServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (EndDeviceServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new EndDeviceServiceFileDescriptorSupplier())
              .addMethod(getListEndDevicesMethod())
              .addMethod(getPairAllEndDevicesMethod())
              .addMethod(getRemoveEndDeviceMethod())
              .build();
        }
      }
    }
    return result;
  }
}
