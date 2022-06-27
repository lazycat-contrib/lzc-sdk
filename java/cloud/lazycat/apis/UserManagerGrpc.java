package cloud.lazycat.apis;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: users/users.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class UserManagerGrpc {

  private UserManagerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.UserManager";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.Users.ListUsersRequest,
      cloud.lazycat.apis.Users.ListUsersReply> getListUsersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListUsers",
      requestType = cloud.lazycat.apis.Users.ListUsersRequest.class,
      responseType = cloud.lazycat.apis.Users.ListUsersReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.Users.ListUsersRequest,
      cloud.lazycat.apis.Users.ListUsersReply> getListUsersMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.Users.ListUsersRequest, cloud.lazycat.apis.Users.ListUsersReply> getListUsersMethod;
    if ((getListUsersMethod = UserManagerGrpc.getListUsersMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getListUsersMethod = UserManagerGrpc.getListUsersMethod) == null) {
          UserManagerGrpc.getListUsersMethod = getListUsersMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.Users.ListUsersRequest, cloud.lazycat.apis.Users.ListUsersReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListUsers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.Users.ListUsersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.Users.ListUsersReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("ListUsers"))
              .build();
        }
      }
    }
    return getListUsersMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserManagerStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserManagerStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserManagerStub>() {
        @java.lang.Override
        public UserManagerStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserManagerStub(channel, callOptions);
        }
      };
    return UserManagerStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserManagerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserManagerBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserManagerBlockingStub>() {
        @java.lang.Override
        public UserManagerBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserManagerBlockingStub(channel, callOptions);
        }
      };
    return UserManagerBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UserManagerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserManagerFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserManagerFutureStub>() {
        @java.lang.Override
        public UserManagerFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserManagerFutureStub(channel, callOptions);
        }
      };
    return UserManagerFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class UserManagerImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public void listUsers(cloud.lazycat.apis.Users.ListUsersRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.Users.ListUsersReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListUsersMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListUsersMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.Users.ListUsersRequest,
                cloud.lazycat.apis.Users.ListUsersReply>(
                  this, METHODID_LIST_USERS)))
          .build();
    }
  }

  /**
   */
  public static final class UserManagerStub extends io.grpc.stub.AbstractAsyncStub<UserManagerStub> {
    private UserManagerStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserManagerStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserManagerStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public void listUsers(cloud.lazycat.apis.Users.ListUsersRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.Users.ListUsersReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListUsersMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class UserManagerBlockingStub extends io.grpc.stub.AbstractBlockingStub<UserManagerBlockingStub> {
    private UserManagerBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserManagerBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserManagerBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public cloud.lazycat.apis.Users.ListUsersReply listUsers(cloud.lazycat.apis.Users.ListUsersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListUsersMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class UserManagerFutureStub extends io.grpc.stub.AbstractFutureStub<UserManagerFutureStub> {
    private UserManagerFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserManagerFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserManagerFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 枚举当前登陆用户所有的设备信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.Users.ListUsersReply> listUsers(
        cloud.lazycat.apis.Users.ListUsersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListUsersMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_USERS = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final UserManagerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(UserManagerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LIST_USERS:
          serviceImpl.listUsers((cloud.lazycat.apis.Users.ListUsersRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.Users.ListUsersReply>) responseObserver);
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

  private static abstract class UserManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserManagerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.Users.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("UserManager");
    }
  }

  private static final class UserManagerFileDescriptorSupplier
      extends UserManagerBaseDescriptorSupplier {
    UserManagerFileDescriptorSupplier() {}
  }

  private static final class UserManagerMethodDescriptorSupplier
      extends UserManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    UserManagerMethodDescriptorSupplier(String methodName) {
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
      synchronized (UserManagerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserManagerFileDescriptorSupplier())
              .addMethod(getListUsersMethod())
              .build();
        }
      }
    }
    return result;
  }
}
