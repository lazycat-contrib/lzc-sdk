package cloud.lazycat.apis.localdevice;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: localdevice/config.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class UserConfigGrpc {

  private UserConfigGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.localdevice.UserConfig";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest,
      cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse> getGetUserConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserConfig",
      requestType = cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest.class,
      responseType = cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest,
      cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse> getGetUserConfigMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest, cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse> getGetUserConfigMethod;
    if ((getGetUserConfigMethod = UserConfigGrpc.getGetUserConfigMethod) == null) {
      synchronized (UserConfigGrpc.class) {
        if ((getGetUserConfigMethod = UserConfigGrpc.getGetUserConfigMethod) == null) {
          UserConfigGrpc.getGetUserConfigMethod = getGetUserConfigMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest, cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserConfigMethodDescriptorSupplier("GetUserConfig"))
              .build();
        }
      }
    }
    return getGetUserConfigMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserConfigStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserConfigStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserConfigStub>() {
        @java.lang.Override
        public UserConfigStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserConfigStub(channel, callOptions);
        }
      };
    return UserConfigStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserConfigBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserConfigBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserConfigBlockingStub>() {
        @java.lang.Override
        public UserConfigBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserConfigBlockingStub(channel, callOptions);
        }
      };
    return UserConfigBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UserConfigFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserConfigFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserConfigFutureStub>() {
        @java.lang.Override
        public UserConfigFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserConfigFutureStub(channel, callOptions);
        }
      };
    return UserConfigFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class UserConfigImplBase implements io.grpc.BindableService {

    /**
     */
    public void getUserConfig(cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserConfigMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetUserConfigMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest,
                cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse>(
                  this, METHODID_GET_USER_CONFIG)))
          .build();
    }
  }

  /**
   */
  public static final class UserConfigStub extends io.grpc.stub.AbstractAsyncStub<UserConfigStub> {
    private UserConfigStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserConfigStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserConfigStub(channel, callOptions);
    }

    /**
     */
    public void getUserConfig(cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserConfigMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class UserConfigBlockingStub extends io.grpc.stub.AbstractBlockingStub<UserConfigBlockingStub> {
    private UserConfigBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserConfigBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserConfigBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse getUserConfig(cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserConfigMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class UserConfigFutureStub extends io.grpc.stub.AbstractFutureStub<UserConfigFutureStub> {
    private UserConfigFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserConfigFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserConfigFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse> getUserConfig(
        cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserConfigMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_USER_CONFIG = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final UserConfigImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(UserConfigImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_USER_CONFIG:
          serviceImpl.getUserConfig((cloud.lazycat.apis.localdevice.Config.GetUserConfigRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.localdevice.Config.GetUserConfigResponse>) responseObserver);
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

  private static abstract class UserConfigBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserConfigBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.localdevice.Config.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("UserConfig");
    }
  }

  private static final class UserConfigFileDescriptorSupplier
      extends UserConfigBaseDescriptorSupplier {
    UserConfigFileDescriptorSupplier() {}
  }

  private static final class UserConfigMethodDescriptorSupplier
      extends UserConfigBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    UserConfigMethodDescriptorSupplier(String methodName) {
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
      synchronized (UserConfigGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserConfigFileDescriptorSupplier())
              .addMethod(getGetUserConfigMethod())
              .build();
        }
      }
    }
    return result;
  }
}
