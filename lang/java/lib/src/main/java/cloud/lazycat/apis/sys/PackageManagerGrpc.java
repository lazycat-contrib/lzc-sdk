package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/package_manager.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PackageManagerGrpc {

  private PackageManagerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.PackageManager";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest,
      com.google.protobuf.Empty> getInstallMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Install",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest,
      com.google.protobuf.Empty> getInstallMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest, com.google.protobuf.Empty> getInstallMethod;
    if ((getInstallMethod = PackageManagerGrpc.getInstallMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getInstallMethod = PackageManagerGrpc.getInstallMethod) == null) {
          PackageManagerGrpc.getInstallMethod = getInstallMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Install"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("Install"))
              .build();
        }
      }
    }
    return getInstallMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest,
      com.google.protobuf.Empty> getUninstallMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Uninstall",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest,
      com.google.protobuf.Empty> getUninstallMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest, com.google.protobuf.Empty> getUninstallMethod;
    if ((getUninstallMethod = PackageManagerGrpc.getUninstallMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getUninstallMethod = PackageManagerGrpc.getUninstallMethod) == null) {
          PackageManagerGrpc.getUninstallMethod = getUninstallMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Uninstall"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("Uninstall"))
              .build();
        }
      }
    }
    return getUninstallMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid,
      com.google.protobuf.Empty> getClearCacheMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClearCache",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid,
      com.google.protobuf.Empty> getClearCacheMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid, com.google.protobuf.Empty> getClearCacheMethod;
    if ((getClearCacheMethod = PackageManagerGrpc.getClearCacheMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getClearCacheMethod = PackageManagerGrpc.getClearCacheMethod) == null) {
          PackageManagerGrpc.getClearCacheMethod = getClearCacheMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClearCache"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("ClearCache"))
              .build();
        }
      }
    }
    return getClearCacheMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse> getQueryApplicationMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryApplication",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest.class,
      responseType = cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse> getQueryApplicationMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse> getQueryApplicationMethod;
    if ((getQueryApplicationMethod = PackageManagerGrpc.getQueryApplicationMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getQueryApplicationMethod = PackageManagerGrpc.getQueryApplicationMethod) == null) {
          PackageManagerGrpc.getQueryApplicationMethod = getQueryApplicationMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryApplication"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("QueryApplication"))
              .build();
        }
      }
    }
    return getQueryApplicationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage> getQueryAppStorageUsageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryAppStorageUsage",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest.class,
      responseType = cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage> getQueryAppStorageUsageMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage> getQueryAppStorageUsageMethod;
    if ((getQueryAppStorageUsageMethod = PackageManagerGrpc.getQueryAppStorageUsageMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getQueryAppStorageUsageMethod = PackageManagerGrpc.getQueryAppStorageUsageMethod) == null) {
          PackageManagerGrpc.getQueryAppStorageUsageMethod = getQueryAppStorageUsageMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryAppStorageUsage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("QueryAppStorageUsage"))
              .build();
        }
      }
    }
    return getQueryAppStorageUsageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission,
      com.google.protobuf.Empty> getSetUserPermissionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetUserPermissions",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission,
      com.google.protobuf.Empty> getSetUserPermissionsMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission, com.google.protobuf.Empty> getSetUserPermissionsMethod;
    if ((getSetUserPermissionsMethod = PackageManagerGrpc.getSetUserPermissionsMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getSetUserPermissionsMethod = PackageManagerGrpc.getSetUserPermissionsMethod) == null) {
          PackageManagerGrpc.getSetUserPermissionsMethod = getSetUserPermissionsMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetUserPermissions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("SetUserPermissions"))
              .build();
        }
      }
    }
    return getSetUserPermissionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission> getGetUserPermissionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserPermissions",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest.class,
      responseType = cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission> getGetUserPermissionsMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission> getGetUserPermissionsMethod;
    if ((getGetUserPermissionsMethod = PackageManagerGrpc.getGetUserPermissionsMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getGetUserPermissionsMethod = PackageManagerGrpc.getGetUserPermissionsMethod) == null) {
          PackageManagerGrpc.getGetUserPermissionsMethod = getGetUserPermissionsMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserPermissions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("GetUserPermissions"))
              .build();
        }
      }
    }
    return getGetUserPermissionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid,
      com.google.protobuf.Empty> getPauseAppDownloadMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PauseAppDownload",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid,
      com.google.protobuf.Empty> getPauseAppDownloadMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid, com.google.protobuf.Empty> getPauseAppDownloadMethod;
    if ((getPauseAppDownloadMethod = PackageManagerGrpc.getPauseAppDownloadMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getPauseAppDownloadMethod = PackageManagerGrpc.getPauseAppDownloadMethod) == null) {
          PackageManagerGrpc.getPauseAppDownloadMethod = getPauseAppDownloadMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PauseAppDownload"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("PauseAppDownload"))
              .build();
        }
      }
    }
    return getPauseAppDownloadMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse> getGetActionURLMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetActionURL",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest.class,
      responseType = cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse> getGetActionURLMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse> getGetActionURLMethod;
    if ((getGetActionURLMethod = PackageManagerGrpc.getGetActionURLMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getGetActionURLMethod = PackageManagerGrpc.getGetActionURLMethod) == null) {
          PackageManagerGrpc.getGetActionURLMethod = getGetActionURLMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetActionURL"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("GetActionURL"))
              .build();
        }
      }
    }
    return getGetActionURLMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse> getListFileHandlerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListFileHandler",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest.class,
      responseType = cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest,
      cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse> getListFileHandlerMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse> getListFileHandlerMethod;
    if ((getListFileHandlerMethod = PackageManagerGrpc.getListFileHandlerMethod) == null) {
      synchronized (PackageManagerGrpc.class) {
        if ((getListFileHandlerMethod = PackageManagerGrpc.getListFileHandlerMethod) == null) {
          PackageManagerGrpc.getListFileHandlerMethod = getListFileHandlerMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest, cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListFileHandler"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PackageManagerMethodDescriptorSupplier("ListFileHandler"))
              .build();
        }
      }
    }
    return getListFileHandlerMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PackageManagerStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PackageManagerStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PackageManagerStub>() {
        @java.lang.Override
        public PackageManagerStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PackageManagerStub(channel, callOptions);
        }
      };
    return PackageManagerStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PackageManagerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PackageManagerBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PackageManagerBlockingStub>() {
        @java.lang.Override
        public PackageManagerBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PackageManagerBlockingStub(channel, callOptions);
        }
      };
    return PackageManagerBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PackageManagerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PackageManagerFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PackageManagerFutureStub>() {
        @java.lang.Override
        public PackageManagerFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PackageManagerFutureStub(channel, callOptions);
        }
      };
    return PackageManagerFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class PackageManagerImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 根据 URL 和 校验码（可选），安装应用
     * </pre>
     */
    public void install(cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getInstallMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据 AppId 卸载应用
     * </pre>
     */
    public void uninstall(cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUninstallMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据 AppId 清除缓存
     * </pre>
     */
    public void clearCache(cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getClearCacheMethod(), responseObserver);
    }

    /**
     * <pre>
     * 查询用户安装的特定应用的详情
     * </pre>
     */
    public void queryApplication(cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryApplicationMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取应用占用的存储空间详情
     * </pre>
     */
    public void queryAppStorageUsage(cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryAppStorageUsageMethod(), responseObserver);
    }

    /**
     * <pre>
     * 设置某个用户是否可以安装应用
     * </pre>
     */
    public void setUserPermissions(cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetUserPermissionsMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取某个用户安装应用的权限
     * </pre>
     */
    public void getUserPermissions(cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserPermissionsMethod(), responseObserver);
    }

    /**
     * <pre>
     * 暂停下载特定应用
     * </pre>
     */
    public void pauseAppDownload(cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPauseAppDownloadMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取用某个应用打开某个文件的链接
     * </pre>
     */
    public void getActionURL(cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetActionURLMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列出可以打开指定 MIME 类型的应用
     * </pre>
     */
    public void listFileHandler(cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListFileHandlerMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getInstallMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_INSTALL)))
          .addMethod(
            getUninstallMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_UNINSTALL)))
          .addMethod(
            getClearCacheMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid,
                com.google.protobuf.Empty>(
                  this, METHODID_CLEAR_CACHE)))
          .addMethod(
            getQueryApplicationMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest,
                cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse>(
                  this, METHODID_QUERY_APPLICATION)))
          .addMethod(
            getQueryAppStorageUsageMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest,
                cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage>(
                  this, METHODID_QUERY_APP_STORAGE_USAGE)))
          .addMethod(
            getSetUserPermissionsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission,
                com.google.protobuf.Empty>(
                  this, METHODID_SET_USER_PERMISSIONS)))
          .addMethod(
            getGetUserPermissionsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest,
                cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission>(
                  this, METHODID_GET_USER_PERMISSIONS)))
          .addMethod(
            getPauseAppDownloadMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid,
                com.google.protobuf.Empty>(
                  this, METHODID_PAUSE_APP_DOWNLOAD)))
          .addMethod(
            getGetActionURLMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest,
                cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse>(
                  this, METHODID_GET_ACTION_URL)))
          .addMethod(
            getListFileHandlerMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest,
                cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse>(
                  this, METHODID_LIST_FILE_HANDLER)))
          .build();
    }
  }

  /**
   */
  public static final class PackageManagerStub extends io.grpc.stub.AbstractAsyncStub<PackageManagerStub> {
    private PackageManagerStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PackageManagerStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PackageManagerStub(channel, callOptions);
    }

    /**
     * <pre>
     * 根据 URL 和 校验码（可选），安装应用
     * </pre>
     */
    public void install(cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getInstallMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据 AppId 卸载应用
     * </pre>
     */
    public void uninstall(cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUninstallMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据 AppId 清除缓存
     * </pre>
     */
    public void clearCache(cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getClearCacheMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 查询用户安装的特定应用的详情
     * </pre>
     */
    public void queryApplication(cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryApplicationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取应用占用的存储空间详情
     * </pre>
     */
    public void queryAppStorageUsage(cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryAppStorageUsageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 设置某个用户是否可以安装应用
     * </pre>
     */
    public void setUserPermissions(cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetUserPermissionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取某个用户安装应用的权限
     * </pre>
     */
    public void getUserPermissions(cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserPermissionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 暂停下载特定应用
     * </pre>
     */
    public void pauseAppDownload(cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPauseAppDownloadMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取用某个应用打开某个文件的链接
     * </pre>
     */
    public void getActionURL(cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetActionURLMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列出可以打开指定 MIME 类型的应用
     * </pre>
     */
    public void listFileHandler(cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListFileHandlerMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class PackageManagerBlockingStub extends io.grpc.stub.AbstractBlockingStub<PackageManagerBlockingStub> {
    private PackageManagerBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PackageManagerBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PackageManagerBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 根据 URL 和 校验码（可选），安装应用
     * </pre>
     */
    public com.google.protobuf.Empty install(cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getInstallMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据 AppId 卸载应用
     * </pre>
     */
    public com.google.protobuf.Empty uninstall(cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUninstallMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据 AppId 清除缓存
     * </pre>
     */
    public com.google.protobuf.Empty clearCache(cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getClearCacheMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 查询用户安装的特定应用的详情
     * </pre>
     */
    public cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse queryApplication(cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryApplicationMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取应用占用的存储空间详情
     * </pre>
     */
    public cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage queryAppStorageUsage(cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryAppStorageUsageMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 设置某个用户是否可以安装应用
     * </pre>
     */
    public com.google.protobuf.Empty setUserPermissions(cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetUserPermissionsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取某个用户安装应用的权限
     * </pre>
     */
    public cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission getUserPermissions(cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserPermissionsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 暂停下载特定应用
     * </pre>
     */
    public com.google.protobuf.Empty pauseAppDownload(cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPauseAppDownloadMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取用某个应用打开某个文件的链接
     * </pre>
     */
    public cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse getActionURL(cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActionURLMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列出可以打开指定 MIME 类型的应用
     * </pre>
     */
    public cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse listFileHandler(cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListFileHandlerMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class PackageManagerFutureStub extends io.grpc.stub.AbstractFutureStub<PackageManagerFutureStub> {
    private PackageManagerFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PackageManagerFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PackageManagerFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 根据 URL 和 校验码（可选），安装应用
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> install(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getInstallMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据 AppId 卸载应用
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> uninstall(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUninstallMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据 AppId 清除缓存
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> clearCache(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getClearCacheMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 查询用户安装的特定应用的详情
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse> queryApplication(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryApplicationMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取应用占用的存储空间详情
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage> queryAppStorageUsage(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryAppStorageUsageMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 设置某个用户是否可以安装应用
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> setUserPermissions(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetUserPermissionsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取某个用户安装应用的权限
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission> getUserPermissions(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserPermissionsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 暂停下载特定应用
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> pauseAppDownload(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPauseAppDownloadMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取用某个应用打开某个文件的链接
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse> getActionURL(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetActionURLMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列出可以打开指定 MIME 类型的应用
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse> listFileHandler(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListFileHandlerMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_INSTALL = 0;
  private static final int METHODID_UNINSTALL = 1;
  private static final int METHODID_CLEAR_CACHE = 2;
  private static final int METHODID_QUERY_APPLICATION = 3;
  private static final int METHODID_QUERY_APP_STORAGE_USAGE = 4;
  private static final int METHODID_SET_USER_PERMISSIONS = 5;
  private static final int METHODID_GET_USER_PERMISSIONS = 6;
  private static final int METHODID_PAUSE_APP_DOWNLOAD = 7;
  private static final int METHODID_GET_ACTION_URL = 8;
  private static final int METHODID_LIST_FILE_HANDLER = 9;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final PackageManagerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(PackageManagerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_INSTALL:
          serviceImpl.install((cloud.lazycat.apis.sys.PackageManagerOuterClass.InstallRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_UNINSTALL:
          serviceImpl.uninstall((cloud.lazycat.apis.sys.PackageManagerOuterClass.UninstallRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_CLEAR_CACHE:
          serviceImpl.clearCache((cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_QUERY_APPLICATION:
          serviceImpl.queryApplication((cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryApplicationResponse>) responseObserver);
          break;
        case METHODID_QUERY_APP_STORAGE_USAGE:
          serviceImpl.queryAppStorageUsage((cloud.lazycat.apis.sys.PackageManagerOuterClass.QueryAppStorageUsageRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.AppStorageUsage>) responseObserver);
          break;
        case METHODID_SET_USER_PERMISSIONS:
          serviceImpl.setUserPermissions((cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GET_USER_PERMISSIONS:
          serviceImpl.getUserPermissions((cloud.lazycat.apis.sys.PackageManagerOuterClass.GetUserPermissionsRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.UserPermission>) responseObserver);
          break;
        case METHODID_PAUSE_APP_DOWNLOAD:
          serviceImpl.pauseAppDownload((cloud.lazycat.apis.sys.PackageManagerOuterClass.Appid) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GET_ACTION_URL:
          serviceImpl.getActionURL((cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.GetActionURLResponse>) responseObserver);
          break;
        case METHODID_LIST_FILE_HANDLER:
          serviceImpl.listFileHandler((cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PackageManagerOuterClass.ListFileHandlerResponse>) responseObserver);
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

  private static abstract class PackageManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PackageManagerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.PackageManagerOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PackageManager");
    }
  }

  private static final class PackageManagerFileDescriptorSupplier
      extends PackageManagerBaseDescriptorSupplier {
    PackageManagerFileDescriptorSupplier() {}
  }

  private static final class PackageManagerMethodDescriptorSupplier
      extends PackageManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    PackageManagerMethodDescriptorSupplier(String methodName) {
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
      synchronized (PackageManagerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PackageManagerFileDescriptorSupplier())
              .addMethod(getInstallMethod())
              .addMethod(getUninstallMethod())
              .addMethod(getClearCacheMethod())
              .addMethod(getQueryApplicationMethod())
              .addMethod(getQueryAppStorageUsageMethod())
              .addMethod(getSetUserPermissionsMethod())
              .addMethod(getGetUserPermissionsMethod())
              .addMethod(getPauseAppDownloadMethod())
              .addMethod(getGetActionURLMethod())
              .addMethod(getListFileHandlerMethod())
              .build();
        }
      }
    }
    return result;
  }
}
