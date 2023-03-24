package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 *应用管理
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/hc-core/hc-core.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class CoreSystemGrpc {

  private CoreSystemGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.CoreSystem";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.ListOptRequest,
      cloud.lazycat.apis.sys.HcCore.AppIdList> getListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "List",
      requestType = cloud.lazycat.apis.sys.HcCore.ListOptRequest.class,
      responseType = cloud.lazycat.apis.sys.HcCore.AppIdList.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.ListOptRequest,
      cloud.lazycat.apis.sys.HcCore.AppIdList> getListMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.ListOptRequest, cloud.lazycat.apis.sys.HcCore.AppIdList> getListMethod;
    if ((getListMethod = CoreSystemGrpc.getListMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getListMethod = CoreSystemGrpc.getListMethod) == null) {
          CoreSystemGrpc.getListMethod = getListMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.ListOptRequest, cloud.lazycat.apis.sys.HcCore.AppIdList>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "List"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.ListOptRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.AppIdList.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("List"))
              .build();
        }
      }
    }
    return getListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.RawData,
      com.google.protobuf.Empty> getApplyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Apply",
      requestType = cloud.lazycat.apis.sys.HcCore.RawData.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.RawData,
      com.google.protobuf.Empty> getApplyMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.RawData, com.google.protobuf.Empty> getApplyMethod;
    if ((getApplyMethod = CoreSystemGrpc.getApplyMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getApplyMethod = CoreSystemGrpc.getApplyMethod) == null) {
          CoreSystemGrpc.getApplyMethod = getApplyMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.RawData, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Apply"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.RawData.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("Apply"))
              .build();
        }
      }
    }
    return getApplyMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL,
      com.google.protobuf.Empty> getApply2Method;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Apply2",
      requestType = cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL,
      com.google.protobuf.Empty> getApply2Method() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL, com.google.protobuf.Empty> getApply2Method;
    if ((getApply2Method = CoreSystemGrpc.getApply2Method) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getApply2Method = CoreSystemGrpc.getApply2Method) == null) {
          CoreSystemGrpc.getApply2Method = getApply2Method =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Apply2"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("Apply2"))
              .build();
        }
      }
    }
    return getApply2Method;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.DeleteRequest,
      com.google.protobuf.Empty> getRemoveMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Remove",
      requestType = cloud.lazycat.apis.sys.HcCore.DeleteRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.DeleteRequest,
      com.google.protobuf.Empty> getRemoveMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.DeleteRequest, com.google.protobuf.Empty> getRemoveMethod;
    if ((getRemoveMethod = CoreSystemGrpc.getRemoveMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getRemoveMethod = CoreSystemGrpc.getRemoveMethod) == null) {
          CoreSystemGrpc.getRemoveMethod = getRemoveMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.DeleteRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Remove"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.DeleteRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("Remove"))
              .build();
        }
      }
    }
    return getRemoveMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest,
      com.google.protobuf.Empty> getDisableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Disable",
      requestType = cloud.lazycat.apis.sys.HcCore.QueryRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest,
      com.google.protobuf.Empty> getDisableMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest, com.google.protobuf.Empty> getDisableMethod;
    if ((getDisableMethod = CoreSystemGrpc.getDisableMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getDisableMethod = CoreSystemGrpc.getDisableMethod) == null) {
          CoreSystemGrpc.getDisableMethod = getDisableMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.QueryRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Disable"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.QueryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("Disable"))
              .build();
        }
      }
    }
    return getDisableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest,
      com.google.protobuf.Empty> getEnableMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Enable",
      requestType = cloud.lazycat.apis.sys.HcCore.QueryRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest,
      com.google.protobuf.Empty> getEnableMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest, com.google.protobuf.Empty> getEnableMethod;
    if ((getEnableMethod = CoreSystemGrpc.getEnableMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getEnableMethod = CoreSystemGrpc.getEnableMethod) == null) {
          CoreSystemGrpc.getEnableMethod = getEnableMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.QueryRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Enable"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.QueryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("Enable"))
              .build();
        }
      }
    }
    return getEnableMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest,
      cloud.lazycat.apis.sys.HcCore.AppMetadataArray> getQueryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Query",
      requestType = cloud.lazycat.apis.sys.HcCore.QueryRequest.class,
      responseType = cloud.lazycat.apis.sys.HcCore.AppMetadataArray.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest,
      cloud.lazycat.apis.sys.HcCore.AppMetadataArray> getQueryMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.QueryRequest, cloud.lazycat.apis.sys.HcCore.AppMetadataArray> getQueryMethod;
    if ((getQueryMethod = CoreSystemGrpc.getQueryMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getQueryMethod = CoreSystemGrpc.getQueryMethod) == null) {
          CoreSystemGrpc.getQueryMethod = getQueryMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.QueryRequest, cloud.lazycat.apis.sys.HcCore.AppMetadataArray>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Query"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.QueryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.AppMetadataArray.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("Query"))
              .build();
        }
      }
    }
    return getQueryMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.HcCore.AppsStatus> getStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Status",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.HcCore.AppsStatus.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.HcCore.AppsStatus> getStatusMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.HcCore.AppsStatus> getStatusMethod;
    if ((getStatusMethod = CoreSystemGrpc.getStatusMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getStatusMethod = CoreSystemGrpc.getStatusMethod) == null) {
          CoreSystemGrpc.getStatusMethod = getStatusMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.HcCore.AppsStatus>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Status"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.AppsStatus.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("Status"))
              .build();
        }
      }
    }
    return getStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.AppId,
      cloud.lazycat.apis.sys.HcCore.InstancesStatus> getAppStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AppStatus",
      requestType = cloud.lazycat.apis.sys.HcCore.AppId.class,
      responseType = cloud.lazycat.apis.sys.HcCore.InstancesStatus.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.AppId,
      cloud.lazycat.apis.sys.HcCore.InstancesStatus> getAppStatusMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.AppId, cloud.lazycat.apis.sys.HcCore.InstancesStatus> getAppStatusMethod;
    if ((getAppStatusMethod = CoreSystemGrpc.getAppStatusMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getAppStatusMethod = CoreSystemGrpc.getAppStatusMethod) == null) {
          CoreSystemGrpc.getAppStatusMethod = getAppStatusMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.AppId, cloud.lazycat.apis.sys.HcCore.InstancesStatus>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AppStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.AppId.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.InstancesStatus.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("AppStatus"))
              .build();
        }
      }
    }
    return getAppStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      cloud.lazycat.apis.sys.HcCore.StatusInfo> getInstanceStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "InstanceStatus",
      requestType = cloud.lazycat.apis.sys.HcCore.InstanceInfo.class,
      responseType = cloud.lazycat.apis.sys.HcCore.StatusInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      cloud.lazycat.apis.sys.HcCore.StatusInfo> getInstanceStatusMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo, cloud.lazycat.apis.sys.HcCore.StatusInfo> getInstanceStatusMethod;
    if ((getInstanceStatusMethod = CoreSystemGrpc.getInstanceStatusMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getInstanceStatusMethod = CoreSystemGrpc.getInstanceStatusMethod) == null) {
          CoreSystemGrpc.getInstanceStatusMethod = getInstanceStatusMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.InstanceInfo, cloud.lazycat.apis.sys.HcCore.StatusInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "InstanceStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.InstanceInfo.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.StatusInfo.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("InstanceStatus"))
              .build();
        }
      }
    }
    return getInstanceStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      com.google.protobuf.Empty> getInstanceAddMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "InstanceAdd",
      requestType = cloud.lazycat.apis.sys.HcCore.InstanceInfo.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      com.google.protobuf.Empty> getInstanceAddMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo, com.google.protobuf.Empty> getInstanceAddMethod;
    if ((getInstanceAddMethod = CoreSystemGrpc.getInstanceAddMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getInstanceAddMethod = CoreSystemGrpc.getInstanceAddMethod) == null) {
          CoreSystemGrpc.getInstanceAddMethod = getInstanceAddMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.InstanceInfo, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "InstanceAdd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.InstanceInfo.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("InstanceAdd"))
              .build();
        }
      }
    }
    return getInstanceAddMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      com.google.protobuf.Empty> getInstanceRemoveMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "InstanceRemove",
      requestType = cloud.lazycat.apis.sys.HcCore.InstanceInfo.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      com.google.protobuf.Empty> getInstanceRemoveMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo, com.google.protobuf.Empty> getInstanceRemoveMethod;
    if ((getInstanceRemoveMethod = CoreSystemGrpc.getInstanceRemoveMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getInstanceRemoveMethod = CoreSystemGrpc.getInstanceRemoveMethod) == null) {
          CoreSystemGrpc.getInstanceRemoveMethod = getInstanceRemoveMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.InstanceInfo, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "InstanceRemove"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.InstanceInfo.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("InstanceRemove"))
              .build();
        }
      }
    }
    return getInstanceRemoveMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      com.google.protobuf.Empty> getInstancePauseMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "InstancePause",
      requestType = cloud.lazycat.apis.sys.HcCore.InstanceInfo.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      com.google.protobuf.Empty> getInstancePauseMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo, com.google.protobuf.Empty> getInstancePauseMethod;
    if ((getInstancePauseMethod = CoreSystemGrpc.getInstancePauseMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getInstancePauseMethod = CoreSystemGrpc.getInstancePauseMethod) == null) {
          CoreSystemGrpc.getInstancePauseMethod = getInstancePauseMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.InstanceInfo, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "InstancePause"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.InstanceInfo.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("InstancePause"))
              .build();
        }
      }
    }
    return getInstancePauseMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      com.google.protobuf.Empty> getInstanceResumeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "InstanceResume",
      requestType = cloud.lazycat.apis.sys.HcCore.InstanceInfo.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo,
      com.google.protobuf.Empty> getInstanceResumeMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.HcCore.InstanceInfo, com.google.protobuf.Empty> getInstanceResumeMethod;
    if ((getInstanceResumeMethod = CoreSystemGrpc.getInstanceResumeMethod) == null) {
      synchronized (CoreSystemGrpc.class) {
        if ((getInstanceResumeMethod = CoreSystemGrpc.getInstanceResumeMethod) == null) {
          CoreSystemGrpc.getInstanceResumeMethod = getInstanceResumeMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.HcCore.InstanceInfo, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "InstanceResume"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.HcCore.InstanceInfo.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new CoreSystemMethodDescriptorSupplier("InstanceResume"))
              .build();
        }
      }
    }
    return getInstanceResumeMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static CoreSystemStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CoreSystemStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CoreSystemStub>() {
        @java.lang.Override
        public CoreSystemStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CoreSystemStub(channel, callOptions);
        }
      };
    return CoreSystemStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static CoreSystemBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CoreSystemBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CoreSystemBlockingStub>() {
        @java.lang.Override
        public CoreSystemBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CoreSystemBlockingStub(channel, callOptions);
        }
      };
    return CoreSystemBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static CoreSystemFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CoreSystemFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CoreSystemFutureStub>() {
        @java.lang.Override
        public CoreSystemFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CoreSystemFutureStub(channel, callOptions);
        }
      };
    return CoreSystemFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   *应用管理
   * </pre>
   */
  public static abstract class CoreSystemImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * List 列出所有已安装的应用
     * </pre>
     */
    public void list(cloud.lazycat.apis.sys.HcCore.ListOptRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppIdList> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListMethod(), responseObserver);
    }

    /**
     * <pre>
     * Apply 安装或更新一个 app
     * </pre>
     */
    public void apply(cloud.lazycat.apis.sys.HcCore.RawData request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getApplyMethod(), responseObserver);
    }

    /**
     */
    public void apply2(cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getApply2Method(), responseObserver);
    }

    /**
     * <pre>
     * Remove 移除一个 app
     * </pre>
     */
    public void remove(cloud.lazycat.apis.sys.HcCore.DeleteRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRemoveMethod(), responseObserver);
    }

    /**
     * <pre>
     * Disable 禁用一个 app
     * </pre>
     */
    public void disable(cloud.lazycat.apis.sys.HcCore.QueryRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDisableMethod(), responseObserver);
    }

    /**
     * <pre>
     * Enable 启用一个 app
     * </pre>
     */
    public void enable(cloud.lazycat.apis.sys.HcCore.QueryRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getEnableMethod(), responseObserver);
    }

    /**
     * <pre>
     * Query 查询 app 的元信息
     * </pre>
     */
    public void query(cloud.lazycat.apis.sys.HcCore.QueryRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppMetadataArray> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryMethod(), responseObserver);
    }

    /**
     * <pre>
     * 查询所有 app 的所有实例状态
     * </pre>
     */
    public void status(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppsStatus> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStatusMethod(), responseObserver);
    }

    /**
     * <pre>
     * 查询 app 的所有实例状态
     * </pre>
     */
    public void appStatus(cloud.lazycat.apis.sys.HcCore.AppId request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.InstancesStatus> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAppStatusMethod(), responseObserver);
    }

    /**
     * <pre>
     * 查询实例状态
     * </pre>
     */
    public void instanceStatus(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.StatusInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getInstanceStatusMethod(), responseObserver);
    }

    /**
     * <pre>
     * 添加一个实例
     * </pre>
     */
    public void instanceAdd(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getInstanceAddMethod(), responseObserver);
    }

    /**
     * <pre>
     * 删除一个实例
     * </pre>
     */
    public void instanceRemove(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getInstanceRemoveMethod(), responseObserver);
    }

    /**
     * <pre>
     * 暂停一个实例
     * </pre>
     */
    public void instancePause(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getInstancePauseMethod(), responseObserver);
    }

    /**
     * <pre>
     * 继续一个实例
     * </pre>
     */
    public void instanceResume(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getInstanceResumeMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.ListOptRequest,
                cloud.lazycat.apis.sys.HcCore.AppIdList>(
                  this, METHODID_LIST)))
          .addMethod(
            getApplyMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.RawData,
                com.google.protobuf.Empty>(
                  this, METHODID_APPLY)))
          .addMethod(
            getApply2Method(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL,
                com.google.protobuf.Empty>(
                  this, METHODID_APPLY2)))
          .addMethod(
            getRemoveMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.DeleteRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_REMOVE)))
          .addMethod(
            getDisableMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.QueryRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_DISABLE)))
          .addMethod(
            getEnableMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.QueryRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_ENABLE)))
          .addMethod(
            getQueryMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.QueryRequest,
                cloud.lazycat.apis.sys.HcCore.AppMetadataArray>(
                  this, METHODID_QUERY)))
          .addMethod(
            getStatusMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.HcCore.AppsStatus>(
                  this, METHODID_STATUS)))
          .addMethod(
            getAppStatusMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.AppId,
                cloud.lazycat.apis.sys.HcCore.InstancesStatus>(
                  this, METHODID_APP_STATUS)))
          .addMethod(
            getInstanceStatusMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.InstanceInfo,
                cloud.lazycat.apis.sys.HcCore.StatusInfo>(
                  this, METHODID_INSTANCE_STATUS)))
          .addMethod(
            getInstanceAddMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.InstanceInfo,
                com.google.protobuf.Empty>(
                  this, METHODID_INSTANCE_ADD)))
          .addMethod(
            getInstanceRemoveMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.InstanceInfo,
                com.google.protobuf.Empty>(
                  this, METHODID_INSTANCE_REMOVE)))
          .addMethod(
            getInstancePauseMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.InstanceInfo,
                com.google.protobuf.Empty>(
                  this, METHODID_INSTANCE_PAUSE)))
          .addMethod(
            getInstanceResumeMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.HcCore.InstanceInfo,
                com.google.protobuf.Empty>(
                  this, METHODID_INSTANCE_RESUME)))
          .build();
    }
  }

  /**
   * <pre>
   *应用管理
   * </pre>
   */
  public static final class CoreSystemStub extends io.grpc.stub.AbstractAsyncStub<CoreSystemStub> {
    private CoreSystemStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CoreSystemStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CoreSystemStub(channel, callOptions);
    }

    /**
     * <pre>
     * List 列出所有已安装的应用
     * </pre>
     */
    public void list(cloud.lazycat.apis.sys.HcCore.ListOptRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppIdList> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Apply 安装或更新一个 app
     * </pre>
     */
    public void apply(cloud.lazycat.apis.sys.HcCore.RawData request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getApplyMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void apply2(cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getApply2Method(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Remove 移除一个 app
     * </pre>
     */
    public void remove(cloud.lazycat.apis.sys.HcCore.DeleteRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRemoveMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Disable 禁用一个 app
     * </pre>
     */
    public void disable(cloud.lazycat.apis.sys.HcCore.QueryRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDisableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Enable 启用一个 app
     * </pre>
     */
    public void enable(cloud.lazycat.apis.sys.HcCore.QueryRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getEnableMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Query 查询 app 的元信息
     * </pre>
     */
    public void query(cloud.lazycat.apis.sys.HcCore.QueryRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppMetadataArray> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 查询所有 app 的所有实例状态
     * </pre>
     */
    public void status(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppsStatus> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 查询 app 的所有实例状态
     * </pre>
     */
    public void appStatus(cloud.lazycat.apis.sys.HcCore.AppId request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.InstancesStatus> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAppStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 查询实例状态
     * </pre>
     */
    public void instanceStatus(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.StatusInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getInstanceStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 添加一个实例
     * </pre>
     */
    public void instanceAdd(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getInstanceAddMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 删除一个实例
     * </pre>
     */
    public void instanceRemove(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getInstanceRemoveMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 暂停一个实例
     * </pre>
     */
    public void instancePause(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getInstancePauseMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 继续一个实例
     * </pre>
     */
    public void instanceResume(cloud.lazycat.apis.sys.HcCore.InstanceInfo request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getInstanceResumeMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   *应用管理
   * </pre>
   */
  public static final class CoreSystemBlockingStub extends io.grpc.stub.AbstractBlockingStub<CoreSystemBlockingStub> {
    private CoreSystemBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CoreSystemBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CoreSystemBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * List 列出所有已安装的应用
     * </pre>
     */
    public cloud.lazycat.apis.sys.HcCore.AppIdList list(cloud.lazycat.apis.sys.HcCore.ListOptRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Apply 安装或更新一个 app
     * </pre>
     */
    public com.google.protobuf.Empty apply(cloud.lazycat.apis.sys.HcCore.RawData request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getApplyMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty apply2(cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getApply2Method(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Remove 移除一个 app
     * </pre>
     */
    public com.google.protobuf.Empty remove(cloud.lazycat.apis.sys.HcCore.DeleteRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRemoveMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Disable 禁用一个 app
     * </pre>
     */
    public com.google.protobuf.Empty disable(cloud.lazycat.apis.sys.HcCore.QueryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDisableMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Enable 启用一个 app
     * </pre>
     */
    public com.google.protobuf.Empty enable(cloud.lazycat.apis.sys.HcCore.QueryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getEnableMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Query 查询 app 的元信息
     * </pre>
     */
    public cloud.lazycat.apis.sys.HcCore.AppMetadataArray query(cloud.lazycat.apis.sys.HcCore.QueryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 查询所有 app 的所有实例状态
     * </pre>
     */
    public cloud.lazycat.apis.sys.HcCore.AppsStatus status(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 查询 app 的所有实例状态
     * </pre>
     */
    public cloud.lazycat.apis.sys.HcCore.InstancesStatus appStatus(cloud.lazycat.apis.sys.HcCore.AppId request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAppStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 查询实例状态
     * </pre>
     */
    public cloud.lazycat.apis.sys.HcCore.StatusInfo instanceStatus(cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getInstanceStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 添加一个实例
     * </pre>
     */
    public com.google.protobuf.Empty instanceAdd(cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getInstanceAddMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 删除一个实例
     * </pre>
     */
    public com.google.protobuf.Empty instanceRemove(cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getInstanceRemoveMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 暂停一个实例
     * </pre>
     */
    public com.google.protobuf.Empty instancePause(cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getInstancePauseMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 继续一个实例
     * </pre>
     */
    public com.google.protobuf.Empty instanceResume(cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getInstanceResumeMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   *应用管理
   * </pre>
   */
  public static final class CoreSystemFutureStub extends io.grpc.stub.AbstractFutureStub<CoreSystemFutureStub> {
    private CoreSystemFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CoreSystemFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CoreSystemFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * List 列出所有已安装的应用
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.HcCore.AppIdList> list(
        cloud.lazycat.apis.sys.HcCore.ListOptRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Apply 安装或更新一个 app
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> apply(
        cloud.lazycat.apis.sys.HcCore.RawData request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getApplyMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> apply2(
        cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getApply2Method(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Remove 移除一个 app
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> remove(
        cloud.lazycat.apis.sys.HcCore.DeleteRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRemoveMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Disable 禁用一个 app
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> disable(
        cloud.lazycat.apis.sys.HcCore.QueryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDisableMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Enable 启用一个 app
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> enable(
        cloud.lazycat.apis.sys.HcCore.QueryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getEnableMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Query 查询 app 的元信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.HcCore.AppMetadataArray> query(
        cloud.lazycat.apis.sys.HcCore.QueryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 查询所有 app 的所有实例状态
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.HcCore.AppsStatus> status(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStatusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 查询 app 的所有实例状态
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.HcCore.InstancesStatus> appStatus(
        cloud.lazycat.apis.sys.HcCore.AppId request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAppStatusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 查询实例状态
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.HcCore.StatusInfo> instanceStatus(
        cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getInstanceStatusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 添加一个实例
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> instanceAdd(
        cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getInstanceAddMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 删除一个实例
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> instanceRemove(
        cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getInstanceRemoveMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 暂停一个实例
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> instancePause(
        cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getInstancePauseMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 继续一个实例
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> instanceResume(
        cloud.lazycat.apis.sys.HcCore.InstanceInfo request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getInstanceResumeMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST = 0;
  private static final int METHODID_APPLY = 1;
  private static final int METHODID_APPLY2 = 2;
  private static final int METHODID_REMOVE = 3;
  private static final int METHODID_DISABLE = 4;
  private static final int METHODID_ENABLE = 5;
  private static final int METHODID_QUERY = 6;
  private static final int METHODID_STATUS = 7;
  private static final int METHODID_APP_STATUS = 8;
  private static final int METHODID_INSTANCE_STATUS = 9;
  private static final int METHODID_INSTANCE_ADD = 10;
  private static final int METHODID_INSTANCE_REMOVE = 11;
  private static final int METHODID_INSTANCE_PAUSE = 12;
  private static final int METHODID_INSTANCE_RESUME = 13;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final CoreSystemImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(CoreSystemImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_LIST:
          serviceImpl.list((cloud.lazycat.apis.sys.HcCore.ListOptRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppIdList>) responseObserver);
          break;
        case METHODID_APPLY:
          serviceImpl.apply((cloud.lazycat.apis.sys.HcCore.RawData) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_APPLY2:
          serviceImpl.apply2((cloud.lazycat.apis.sys.PackageManagerOuterClass.PkgURL) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_REMOVE:
          serviceImpl.remove((cloud.lazycat.apis.sys.HcCore.DeleteRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_DISABLE:
          serviceImpl.disable((cloud.lazycat.apis.sys.HcCore.QueryRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_ENABLE:
          serviceImpl.enable((cloud.lazycat.apis.sys.HcCore.QueryRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_QUERY:
          serviceImpl.query((cloud.lazycat.apis.sys.HcCore.QueryRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppMetadataArray>) responseObserver);
          break;
        case METHODID_STATUS:
          serviceImpl.status((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.AppsStatus>) responseObserver);
          break;
        case METHODID_APP_STATUS:
          serviceImpl.appStatus((cloud.lazycat.apis.sys.HcCore.AppId) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.InstancesStatus>) responseObserver);
          break;
        case METHODID_INSTANCE_STATUS:
          serviceImpl.instanceStatus((cloud.lazycat.apis.sys.HcCore.InstanceInfo) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.HcCore.StatusInfo>) responseObserver);
          break;
        case METHODID_INSTANCE_ADD:
          serviceImpl.instanceAdd((cloud.lazycat.apis.sys.HcCore.InstanceInfo) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_INSTANCE_REMOVE:
          serviceImpl.instanceRemove((cloud.lazycat.apis.sys.HcCore.InstanceInfo) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_INSTANCE_PAUSE:
          serviceImpl.instancePause((cloud.lazycat.apis.sys.HcCore.InstanceInfo) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_INSTANCE_RESUME:
          serviceImpl.instanceResume((cloud.lazycat.apis.sys.HcCore.InstanceInfo) request,
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

  private static abstract class CoreSystemBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CoreSystemBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.HcCore.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("CoreSystem");
    }
  }

  private static final class CoreSystemFileDescriptorSupplier
      extends CoreSystemBaseDescriptorSupplier {
    CoreSystemFileDescriptorSupplier() {}
  }

  private static final class CoreSystemMethodDescriptorSupplier
      extends CoreSystemBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    CoreSystemMethodDescriptorSupplier(String methodName) {
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
      synchronized (CoreSystemGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new CoreSystemFileDescriptorSupplier())
              .addMethod(getListMethod())
              .addMethod(getApplyMethod())
              .addMethod(getApply2Method())
              .addMethod(getRemoveMethod())
              .addMethod(getDisableMethod())
              .addMethod(getEnableMethod())
              .addMethod(getQueryMethod())
              .addMethod(getStatusMethod())
              .addMethod(getAppStatusMethod())
              .addMethod(getInstanceStatusMethod())
              .addMethod(getInstanceAddMethod())
              .addMethod(getInstanceRemoveMethod())
              .addMethod(getInstancePauseMethod())
              .addMethod(getInstanceResumeMethod())
              .build();
        }
      }
    }
    return result;
  }
}
