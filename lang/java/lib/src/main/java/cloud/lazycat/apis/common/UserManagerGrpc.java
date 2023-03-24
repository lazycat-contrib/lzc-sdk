package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/users.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class UserManagerGrpc {

  private UserManagerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.UserManager";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ListUIDsRequest,
      cloud.lazycat.apis.common.Users.ListUIDsReply> getListUIDsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListUIDs",
      requestType = cloud.lazycat.apis.common.Users.ListUIDsRequest.class,
      responseType = cloud.lazycat.apis.common.Users.ListUIDsReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ListUIDsRequest,
      cloud.lazycat.apis.common.Users.ListUIDsReply> getListUIDsMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ListUIDsRequest, cloud.lazycat.apis.common.Users.ListUIDsReply> getListUIDsMethod;
    if ((getListUIDsMethod = UserManagerGrpc.getListUIDsMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getListUIDsMethod = UserManagerGrpc.getListUIDsMethod) == null) {
          UserManagerGrpc.getListUIDsMethod = getListUIDsMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Users.ListUIDsRequest, cloud.lazycat.apis.common.Users.ListUIDsReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListUIDs"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.ListUIDsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.ListUIDsReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("ListUIDs"))
              .build();
        }
      }
    }
    return getListUIDsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.UserID,
      cloud.lazycat.apis.common.Users.UserInfo> getQueryUserInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryUserInfo",
      requestType = cloud.lazycat.apis.common.Users.UserID.class,
      responseType = cloud.lazycat.apis.common.Users.UserInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.UserID,
      cloud.lazycat.apis.common.Users.UserInfo> getQueryUserInfoMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.UserID, cloud.lazycat.apis.common.Users.UserInfo> getQueryUserInfoMethod;
    if ((getQueryUserInfoMethod = UserManagerGrpc.getQueryUserInfoMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getQueryUserInfoMethod = UserManagerGrpc.getQueryUserInfoMethod) == null) {
          UserManagerGrpc.getQueryUserInfoMethod = getQueryUserInfoMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Users.UserID, cloud.lazycat.apis.common.Users.UserInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryUserInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.UserID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.UserInfo.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("QueryUserInfo"))
              .build();
        }
      }
    }
    return getQueryUserInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.UpdateUserInfoRequest,
      com.google.protobuf.Empty> getUpdateUserInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateUserInfo",
      requestType = cloud.lazycat.apis.common.Users.UpdateUserInfoRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.UpdateUserInfoRequest,
      com.google.protobuf.Empty> getUpdateUserInfoMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.UpdateUserInfoRequest, com.google.protobuf.Empty> getUpdateUserInfoMethod;
    if ((getUpdateUserInfoMethod = UserManagerGrpc.getUpdateUserInfoMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getUpdateUserInfoMethod = UserManagerGrpc.getUpdateUserInfoMethod) == null) {
          UserManagerGrpc.getUpdateUserInfoMethod = getUpdateUserInfoMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Users.UpdateUserInfoRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateUserInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.UpdateUserInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("UpdateUserInfo"))
              .build();
        }
      }
    }
    return getUpdateUserInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ChangeRoleReqeust,
      com.google.protobuf.Empty> getChangeRoleMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ChangeRole",
      requestType = cloud.lazycat.apis.common.Users.ChangeRoleReqeust.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ChangeRoleReqeust,
      com.google.protobuf.Empty> getChangeRoleMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ChangeRoleReqeust, com.google.protobuf.Empty> getChangeRoleMethod;
    if ((getChangeRoleMethod = UserManagerGrpc.getChangeRoleMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getChangeRoleMethod = UserManagerGrpc.getChangeRoleMethod) == null) {
          UserManagerGrpc.getChangeRoleMethod = getChangeRoleMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Users.ChangeRoleReqeust, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ChangeRole"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.ChangeRoleReqeust.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("ChangeRole"))
              .build();
        }
      }
    }
    return getChangeRoleMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ResetPasswordRequest,
      com.google.protobuf.Empty> getResetPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ResetPassword",
      requestType = cloud.lazycat.apis.common.Users.ResetPasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ResetPasswordRequest,
      com.google.protobuf.Empty> getResetPasswordMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ResetPasswordRequest, com.google.protobuf.Empty> getResetPasswordMethod;
    if ((getResetPasswordMethod = UserManagerGrpc.getResetPasswordMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getResetPasswordMethod = UserManagerGrpc.getResetPasswordMethod) == null) {
          UserManagerGrpc.getResetPasswordMethod = getResetPasswordMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Users.ResetPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ResetPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.ResetPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("ResetPassword"))
              .build();
        }
      }
    }
    return getResetPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.DeleteUserRequest,
      com.google.protobuf.Empty> getDeleteUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteUser",
      requestType = cloud.lazycat.apis.common.Users.DeleteUserRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.DeleteUserRequest,
      com.google.protobuf.Empty> getDeleteUserMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.DeleteUserRequest, com.google.protobuf.Empty> getDeleteUserMethod;
    if ((getDeleteUserMethod = UserManagerGrpc.getDeleteUserMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getDeleteUserMethod = UserManagerGrpc.getDeleteUserMethod) == null) {
          UserManagerGrpc.getDeleteUserMethod = getDeleteUserMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Users.DeleteUserRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.DeleteUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("DeleteUser"))
              .build();
        }
      }
    }
    return getDeleteUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.CreateUserRequest,
      com.google.protobuf.Empty> getCreateUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateUser",
      requestType = cloud.lazycat.apis.common.Users.CreateUserRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.CreateUserRequest,
      com.google.protobuf.Empty> getCreateUserMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.CreateUserRequest, com.google.protobuf.Empty> getCreateUserMethod;
    if ((getCreateUserMethod = UserManagerGrpc.getCreateUserMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getCreateUserMethod = UserManagerGrpc.getCreateUserMethod) == null) {
          UserManagerGrpc.getCreateUserMethod = getCreateUserMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Users.CreateUserRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.CreateUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("CreateUser"))
              .build();
        }
      }
    }
    return getCreateUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ForceResetPasswordRequest,
      com.google.protobuf.Empty> getForceResetPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ForceResetPassword",
      requestType = cloud.lazycat.apis.common.Users.ForceResetPasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ForceResetPasswordRequest,
      com.google.protobuf.Empty> getForceResetPasswordMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Users.ForceResetPasswordRequest, com.google.protobuf.Empty> getForceResetPasswordMethod;
    if ((getForceResetPasswordMethod = UserManagerGrpc.getForceResetPasswordMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getForceResetPasswordMethod = UserManagerGrpc.getForceResetPasswordMethod) == null) {
          UserManagerGrpc.getForceResetPasswordMethod = getForceResetPasswordMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Users.ForceResetPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ForceResetPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Users.ForceResetPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("ForceResetPassword"))
              .build();
        }
      }
    }
    return getForceResetPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest,
      cloud.lazycat.apis.sys.PortalServer.UserInvitation> getGenUserInvitationMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GenUserInvitation",
      requestType = cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.UserInvitation.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest,
      cloud.lazycat.apis.sys.PortalServer.UserInvitation> getGenUserInvitationMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest, cloud.lazycat.apis.sys.PortalServer.UserInvitation> getGenUserInvitationMethod;
    if ((getGenUserInvitationMethod = UserManagerGrpc.getGenUserInvitationMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getGenUserInvitationMethod = UserManagerGrpc.getGenUserInvitationMethod) == null) {
          UserManagerGrpc.getGenUserInvitationMethod = getGenUserInvitationMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest, cloud.lazycat.apis.sys.PortalServer.UserInvitation>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GenUserInvitation"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.UserInvitation.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("GenUserInvitation"))
              .build();
        }
      }
    }
    return getGenUserInvitationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest,
      com.google.protobuf.Empty> getCheckPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CheckPassword",
      requestType = cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest,
      com.google.protobuf.Empty> getCheckPasswordMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest, com.google.protobuf.Empty> getCheckPasswordMethod;
    if ((getCheckPasswordMethod = UserManagerGrpc.getCheckPasswordMethod) == null) {
      synchronized (UserManagerGrpc.class) {
        if ((getCheckPasswordMethod = UserManagerGrpc.getCheckPasswordMethod) == null) {
          UserManagerGrpc.getCheckPasswordMethod = getCheckPasswordMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CheckPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new UserManagerMethodDescriptorSupplier("CheckPassword"))
              .build();
        }
      }
    }
    return getCheckPasswordMethod;
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
     *  获取所有盒子用户的uid信息(允许任意有效用户)
     * </pre>
     */
    public void listUIDs(cloud.lazycat.apis.common.Users.ListUIDsRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Users.ListUIDsReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListUIDsMethod(), responseObserver);
    }

    /**
     * <pre>
     *  根据用户uid查询用户信息(允许任意有效用户)
     * </pre>
     */
    public void queryUserInfo(cloud.lazycat.apis.common.Users.UserID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Users.UserInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryUserInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     *  更新指定uid用户的nickname和avator(允许用户修改自己的用户信息)
     * </pre>
     */
    public void updateUserInfo(cloud.lazycat.apis.common.Users.UpdateUserInfoRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateUserInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     *  修改指定uid的用户角色(管理员角色允许调用)
     * </pre>
     */
    public void changeRole(cloud.lazycat.apis.common.Users.ChangeRoleReqeust request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getChangeRoleMethod(), responseObserver);
    }

    /**
     * <pre>
     *  通过验证旧密码修改新的密码(允许用户重置自己的密码)
     * </pre>
     */
    public void resetPassword(cloud.lazycat.apis.common.Users.ResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getResetPasswordMethod(), responseObserver);
    }

    /**
     * <pre>
     * 删除用户信息(管理员角色允许调用)
     * </pre>
     */
    public void deleteUser(cloud.lazycat.apis.common.Users.DeleteUserRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteUserMethod(), responseObserver);
    }

    /**
     * <pre>
     * 创建用户信息(管理员角色允许调用)
     * </pre>
     */
    public void createUser(cloud.lazycat.apis.common.Users.CreateUserRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateUserMethod(), responseObserver);
    }

    /**
     * <pre>
     * 强制重置用户密码（管理员角色允许调用)
     * </pre>
     */
    public void forceResetPassword(cloud.lazycat.apis.common.Users.ForceResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getForceResetPasswordMethod(), responseObserver);
    }

    /**
     */
    public void genUserInvitation(cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.UserInvitation> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGenUserInvitationMethod(), responseObserver);
    }

    /**
     * <pre>
     * 检测用户密码有效性（是否能够登录）
     * </pre>
     */
    public void checkPassword(cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCheckPasswordMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListUIDsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Users.ListUIDsRequest,
                cloud.lazycat.apis.common.Users.ListUIDsReply>(
                  this, METHODID_LIST_UIDS)))
          .addMethod(
            getQueryUserInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Users.UserID,
                cloud.lazycat.apis.common.Users.UserInfo>(
                  this, METHODID_QUERY_USER_INFO)))
          .addMethod(
            getUpdateUserInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Users.UpdateUserInfoRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_UPDATE_USER_INFO)))
          .addMethod(
            getChangeRoleMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Users.ChangeRoleReqeust,
                com.google.protobuf.Empty>(
                  this, METHODID_CHANGE_ROLE)))
          .addMethod(
            getResetPasswordMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Users.ResetPasswordRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_RESET_PASSWORD)))
          .addMethod(
            getDeleteUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Users.DeleteUserRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_DELETE_USER)))
          .addMethod(
            getCreateUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Users.CreateUserRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_CREATE_USER)))
          .addMethod(
            getForceResetPasswordMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Users.ForceResetPasswordRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_FORCE_RESET_PASSWORD)))
          .addMethod(
            getGenUserInvitationMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest,
                cloud.lazycat.apis.sys.PortalServer.UserInvitation>(
                  this, METHODID_GEN_USER_INVITATION)))
          .addMethod(
            getCheckPasswordMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_CHECK_PASSWORD)))
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
     *  获取所有盒子用户的uid信息(允许任意有效用户)
     * </pre>
     */
    public void listUIDs(cloud.lazycat.apis.common.Users.ListUIDsRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Users.ListUIDsReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListUIDsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  根据用户uid查询用户信息(允许任意有效用户)
     * </pre>
     */
    public void queryUserInfo(cloud.lazycat.apis.common.Users.UserID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Users.UserInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryUserInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  更新指定uid用户的nickname和avator(允许用户修改自己的用户信息)
     * </pre>
     */
    public void updateUserInfo(cloud.lazycat.apis.common.Users.UpdateUserInfoRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateUserInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  修改指定uid的用户角色(管理员角色允许调用)
     * </pre>
     */
    public void changeRole(cloud.lazycat.apis.common.Users.ChangeRoleReqeust request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getChangeRoleMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  通过验证旧密码修改新的密码(允许用户重置自己的密码)
     * </pre>
     */
    public void resetPassword(cloud.lazycat.apis.common.Users.ResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getResetPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 删除用户信息(管理员角色允许调用)
     * </pre>
     */
    public void deleteUser(cloud.lazycat.apis.common.Users.DeleteUserRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 创建用户信息(管理员角色允许调用)
     * </pre>
     */
    public void createUser(cloud.lazycat.apis.common.Users.CreateUserRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 强制重置用户密码（管理员角色允许调用)
     * </pre>
     */
    public void forceResetPassword(cloud.lazycat.apis.common.Users.ForceResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getForceResetPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void genUserInvitation(cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.UserInvitation> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGenUserInvitationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 检测用户密码有效性（是否能够登录）
     * </pre>
     */
    public void checkPassword(cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCheckPasswordMethod(), getCallOptions()), request, responseObserver);
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
     *  获取所有盒子用户的uid信息(允许任意有效用户)
     * </pre>
     */
    public cloud.lazycat.apis.common.Users.ListUIDsReply listUIDs(cloud.lazycat.apis.common.Users.ListUIDsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListUIDsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  根据用户uid查询用户信息(允许任意有效用户)
     * </pre>
     */
    public cloud.lazycat.apis.common.Users.UserInfo queryUserInfo(cloud.lazycat.apis.common.Users.UserID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryUserInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  更新指定uid用户的nickname和avator(允许用户修改自己的用户信息)
     * </pre>
     */
    public com.google.protobuf.Empty updateUserInfo(cloud.lazycat.apis.common.Users.UpdateUserInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateUserInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  修改指定uid的用户角色(管理员角色允许调用)
     * </pre>
     */
    public com.google.protobuf.Empty changeRole(cloud.lazycat.apis.common.Users.ChangeRoleReqeust request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getChangeRoleMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  通过验证旧密码修改新的密码(允许用户重置自己的密码)
     * </pre>
     */
    public com.google.protobuf.Empty resetPassword(cloud.lazycat.apis.common.Users.ResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getResetPasswordMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 删除用户信息(管理员角色允许调用)
     * </pre>
     */
    public com.google.protobuf.Empty deleteUser(cloud.lazycat.apis.common.Users.DeleteUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteUserMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 创建用户信息(管理员角色允许调用)
     * </pre>
     */
    public com.google.protobuf.Empty createUser(cloud.lazycat.apis.common.Users.CreateUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateUserMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 强制重置用户密码（管理员角色允许调用)
     * </pre>
     */
    public com.google.protobuf.Empty forceResetPassword(cloud.lazycat.apis.common.Users.ForceResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getForceResetPasswordMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.sys.PortalServer.UserInvitation genUserInvitation(cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGenUserInvitationMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 检测用户密码有效性（是否能够登录）
     * </pre>
     */
    public com.google.protobuf.Empty checkPassword(cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCheckPasswordMethod(), getCallOptions(), request);
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
     *  获取所有盒子用户的uid信息(允许任意有效用户)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Users.ListUIDsReply> listUIDs(
        cloud.lazycat.apis.common.Users.ListUIDsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListUIDsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  根据用户uid查询用户信息(允许任意有效用户)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Users.UserInfo> queryUserInfo(
        cloud.lazycat.apis.common.Users.UserID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryUserInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  更新指定uid用户的nickname和avator(允许用户修改自己的用户信息)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> updateUserInfo(
        cloud.lazycat.apis.common.Users.UpdateUserInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateUserInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  修改指定uid的用户角色(管理员角色允许调用)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> changeRole(
        cloud.lazycat.apis.common.Users.ChangeRoleReqeust request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getChangeRoleMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  通过验证旧密码修改新的密码(允许用户重置自己的密码)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> resetPassword(
        cloud.lazycat.apis.common.Users.ResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getResetPasswordMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 删除用户信息(管理员角色允许调用)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> deleteUser(
        cloud.lazycat.apis.common.Users.DeleteUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteUserMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 创建用户信息(管理员角色允许调用)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> createUser(
        cloud.lazycat.apis.common.Users.CreateUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 强制重置用户密码（管理员角色允许调用)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> forceResetPassword(
        cloud.lazycat.apis.common.Users.ForceResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getForceResetPasswordMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.UserInvitation> genUserInvitation(
        cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGenUserInvitationMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 检测用户密码有效性（是否能够登录）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> checkPassword(
        cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCheckPasswordMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_UIDS = 0;
  private static final int METHODID_QUERY_USER_INFO = 1;
  private static final int METHODID_UPDATE_USER_INFO = 2;
  private static final int METHODID_CHANGE_ROLE = 3;
  private static final int METHODID_RESET_PASSWORD = 4;
  private static final int METHODID_DELETE_USER = 5;
  private static final int METHODID_CREATE_USER = 6;
  private static final int METHODID_FORCE_RESET_PASSWORD = 7;
  private static final int METHODID_GEN_USER_INVITATION = 8;
  private static final int METHODID_CHECK_PASSWORD = 9;

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
        case METHODID_LIST_UIDS:
          serviceImpl.listUIDs((cloud.lazycat.apis.common.Users.ListUIDsRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Users.ListUIDsReply>) responseObserver);
          break;
        case METHODID_QUERY_USER_INFO:
          serviceImpl.queryUserInfo((cloud.lazycat.apis.common.Users.UserID) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Users.UserInfo>) responseObserver);
          break;
        case METHODID_UPDATE_USER_INFO:
          serviceImpl.updateUserInfo((cloud.lazycat.apis.common.Users.UpdateUserInfoRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_CHANGE_ROLE:
          serviceImpl.changeRole((cloud.lazycat.apis.common.Users.ChangeRoleReqeust) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_RESET_PASSWORD:
          serviceImpl.resetPassword((cloud.lazycat.apis.common.Users.ResetPasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_DELETE_USER:
          serviceImpl.deleteUser((cloud.lazycat.apis.common.Users.DeleteUserRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_CREATE_USER:
          serviceImpl.createUser((cloud.lazycat.apis.common.Users.CreateUserRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_FORCE_RESET_PASSWORD:
          serviceImpl.forceResetPassword((cloud.lazycat.apis.common.Users.ForceResetPasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GEN_USER_INVITATION:
          serviceImpl.genUserInvitation((cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.UserInvitation>) responseObserver);
          break;
        case METHODID_CHECK_PASSWORD:
          serviceImpl.checkPassword((cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest) request,
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

  private static abstract class UserManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserManagerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.Users.getDescriptor();
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
              .addMethod(getListUIDsMethod())
              .addMethod(getQueryUserInfoMethod())
              .addMethod(getUpdateUserInfoMethod())
              .addMethod(getChangeRoleMethod())
              .addMethod(getResetPasswordMethod())
              .addMethod(getDeleteUserMethod())
              .addMethod(getCreateUserMethod())
              .addMethod(getForceResetPasswordMethod())
              .addMethod(getGenUserInvitationMethod())
              .addMethod(getCheckPasswordMethod())
              .build();
        }
      }
    }
    return result;
  }
}
