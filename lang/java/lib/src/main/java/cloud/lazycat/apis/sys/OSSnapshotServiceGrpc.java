package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/OS_snapshot.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class OSSnapshotServiceGrpc {

  private OSSnapshotServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.OSSnapshotService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
      com.google.protobuf.Empty> getBackupPoolAddMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "BackupPoolAdd",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
      com.google.protobuf.Empty> getBackupPoolAddMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest, com.google.protobuf.Empty> getBackupPoolAddMethod;
    if ((getBackupPoolAddMethod = OSSnapshotServiceGrpc.getBackupPoolAddMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getBackupPoolAddMethod = OSSnapshotServiceGrpc.getBackupPoolAddMethod) == null) {
          OSSnapshotServiceGrpc.getBackupPoolAddMethod = getBackupPoolAddMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "BackupPoolAdd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("BackupPoolAdd"))
              .build();
        }
      }
    }
    return getBackupPoolAddMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
      com.google.protobuf.Empty> getBackupPoolDelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "BackupPoolDel",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
      com.google.protobuf.Empty> getBackupPoolDelMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest, com.google.protobuf.Empty> getBackupPoolDelMethod;
    if ((getBackupPoolDelMethod = OSSnapshotServiceGrpc.getBackupPoolDelMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getBackupPoolDelMethod = OSSnapshotServiceGrpc.getBackupPoolDelMethod) == null) {
          OSSnapshotServiceGrpc.getBackupPoolDelMethod = getBackupPoolDelMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "BackupPoolDel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("BackupPoolDel"))
              .build();
        }
      }
    }
    return getBackupPoolDelMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse> getBackupPoolListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "BackupPoolList",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse> getBackupPoolListMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse> getBackupPoolListMethod;
    if ((getBackupPoolListMethod = OSSnapshotServiceGrpc.getBackupPoolListMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getBackupPoolListMethod = OSSnapshotServiceGrpc.getBackupPoolListMethod) == null) {
          OSSnapshotServiceGrpc.getBackupPoolListMethod = getBackupPoolListMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "BackupPoolList"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("BackupPoolList"))
              .build();
        }
      }
    }
    return getBackupPoolListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
      com.google.protobuf.Empty> getDatasetAddMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DatasetAdd",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
      com.google.protobuf.Empty> getDatasetAddMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest, com.google.protobuf.Empty> getDatasetAddMethod;
    if ((getDatasetAddMethod = OSSnapshotServiceGrpc.getDatasetAddMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getDatasetAddMethod = OSSnapshotServiceGrpc.getDatasetAddMethod) == null) {
          OSSnapshotServiceGrpc.getDatasetAddMethod = getDatasetAddMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DatasetAdd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("DatasetAdd"))
              .build();
        }
      }
    }
    return getDatasetAddMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
      com.google.protobuf.Empty> getDatasetDelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DatasetDel",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
      com.google.protobuf.Empty> getDatasetDelMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest, com.google.protobuf.Empty> getDatasetDelMethod;
    if ((getDatasetDelMethod = OSSnapshotServiceGrpc.getDatasetDelMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getDatasetDelMethod = OSSnapshotServiceGrpc.getDatasetDelMethod) == null) {
          OSSnapshotServiceGrpc.getDatasetDelMethod = getDatasetDelMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DatasetDel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("DatasetDel"))
              .build();
        }
      }
    }
    return getDatasetDelMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> getDatasetListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DatasetList",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> getDatasetListMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> getDatasetListMethod;
    if ((getDatasetListMethod = OSSnapshotServiceGrpc.getDatasetListMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getDatasetListMethod = OSSnapshotServiceGrpc.getDatasetListMethod) == null) {
          OSSnapshotServiceGrpc.getDatasetListMethod = getDatasetListMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DatasetList"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("DatasetList"))
              .build();
        }
      }
    }
    return getDatasetListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest,
      com.google.protobuf.Empty> getDatasetBackupDelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DatasetBackupDel",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest,
      com.google.protobuf.Empty> getDatasetBackupDelMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest, com.google.protobuf.Empty> getDatasetBackupDelMethod;
    if ((getDatasetBackupDelMethod = OSSnapshotServiceGrpc.getDatasetBackupDelMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getDatasetBackupDelMethod = OSSnapshotServiceGrpc.getDatasetBackupDelMethod) == null) {
          OSSnapshotServiceGrpc.getDatasetBackupDelMethod = getDatasetBackupDelMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DatasetBackupDel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("DatasetBackupDel"))
              .build();
        }
      }
    }
    return getDatasetBackupDelMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> getDatasetBackupListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DatasetBackupList",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest.class,
      responseType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> getDatasetBackupListMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> getDatasetBackupListMethod;
    if ((getDatasetBackupListMethod = OSSnapshotServiceGrpc.getDatasetBackupListMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getDatasetBackupListMethod = OSSnapshotServiceGrpc.getDatasetBackupListMethod) == null) {
          OSSnapshotServiceGrpc.getDatasetBackupListMethod = getDatasetBackupListMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DatasetBackupList"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("DatasetBackupList"))
              .build();
        }
      }
    }
    return getDatasetBackupListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
      com.google.protobuf.Empty> getSnapshotAddMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapshotAdd",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
      com.google.protobuf.Empty> getSnapshotAddMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest, com.google.protobuf.Empty> getSnapshotAddMethod;
    if ((getSnapshotAddMethod = OSSnapshotServiceGrpc.getSnapshotAddMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getSnapshotAddMethod = OSSnapshotServiceGrpc.getSnapshotAddMethod) == null) {
          OSSnapshotServiceGrpc.getSnapshotAddMethod = getSnapshotAddMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapshotAdd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("SnapshotAdd"))
              .build();
        }
      }
    }
    return getSnapshotAddMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
      com.google.protobuf.Empty> getSnapshotDelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapshotDel",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
      com.google.protobuf.Empty> getSnapshotDelMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest, com.google.protobuf.Empty> getSnapshotDelMethod;
    if ((getSnapshotDelMethod = OSSnapshotServiceGrpc.getSnapshotDelMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getSnapshotDelMethod = OSSnapshotServiceGrpc.getSnapshotDelMethod) == null) {
          OSSnapshotServiceGrpc.getSnapshotDelMethod = getSnapshotDelMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapshotDel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("SnapshotDel"))
              .build();
        }
      }
    }
    return getSnapshotDelMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> getSnapshotListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapshotList",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest.class,
      responseType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> getSnapshotListMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> getSnapshotListMethod;
    if ((getSnapshotListMethod = OSSnapshotServiceGrpc.getSnapshotListMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getSnapshotListMethod = OSSnapshotServiceGrpc.getSnapshotListMethod) == null) {
          OSSnapshotServiceGrpc.getSnapshotListMethod = getSnapshotListMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapshotList"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("SnapshotList"))
              .build();
        }
      }
    }
    return getSnapshotListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
      com.google.protobuf.Empty> getSnapshotRestoreMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapshotRestore",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
      com.google.protobuf.Empty> getSnapshotRestoreMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest, com.google.protobuf.Empty> getSnapshotRestoreMethod;
    if ((getSnapshotRestoreMethod = OSSnapshotServiceGrpc.getSnapshotRestoreMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getSnapshotRestoreMethod = OSSnapshotServiceGrpc.getSnapshotRestoreMethod) == null) {
          OSSnapshotServiceGrpc.getSnapshotRestoreMethod = getSnapshotRestoreMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapshotRestore"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("SnapshotRestore"))
              .build();
        }
      }
    }
    return getSnapshotRestoreMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> getSnapshotBackupAddMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapshotBackupAdd",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest.class,
      responseType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> getSnapshotBackupAddMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> getSnapshotBackupAddMethod;
    if ((getSnapshotBackupAddMethod = OSSnapshotServiceGrpc.getSnapshotBackupAddMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getSnapshotBackupAddMethod = OSSnapshotServiceGrpc.getSnapshotBackupAddMethod) == null) {
          OSSnapshotServiceGrpc.getSnapshotBackupAddMethod = getSnapshotBackupAddMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapshotBackupAdd"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("SnapshotBackupAdd"))
              .build();
        }
      }
    }
    return getSnapshotBackupAddMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest,
      com.google.protobuf.Empty> getSnapshotBackupDelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapshotBackupDel",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest,
      com.google.protobuf.Empty> getSnapshotBackupDelMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest, com.google.protobuf.Empty> getSnapshotBackupDelMethod;
    if ((getSnapshotBackupDelMethod = OSSnapshotServiceGrpc.getSnapshotBackupDelMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getSnapshotBackupDelMethod = OSSnapshotServiceGrpc.getSnapshotBackupDelMethod) == null) {
          OSSnapshotServiceGrpc.getSnapshotBackupDelMethod = getSnapshotBackupDelMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapshotBackupDel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("SnapshotBackupDel"))
              .build();
        }
      }
    }
    return getSnapshotBackupDelMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> getSnapshotBackupListMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapshotBackupList",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest.class,
      responseType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> getSnapshotBackupListMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> getSnapshotBackupListMethod;
    if ((getSnapshotBackupListMethod = OSSnapshotServiceGrpc.getSnapshotBackupListMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getSnapshotBackupListMethod = OSSnapshotServiceGrpc.getSnapshotBackupListMethod) == null) {
          OSSnapshotServiceGrpc.getSnapshotBackupListMethod = getSnapshotBackupListMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapshotBackupList"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("SnapshotBackupList"))
              .build();
        }
      }
    }
    return getSnapshotBackupListMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> getSnapshotBackupRestoreMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SnapshotBackupRestore",
      requestType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest.class,
      responseType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> getSnapshotBackupRestoreMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> getSnapshotBackupRestoreMethod;
    if ((getSnapshotBackupRestoreMethod = OSSnapshotServiceGrpc.getSnapshotBackupRestoreMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getSnapshotBackupRestoreMethod = OSSnapshotServiceGrpc.getSnapshotBackupRestoreMethod) == null) {
          OSSnapshotServiceGrpc.getSnapshotBackupRestoreMethod = getSnapshotBackupRestoreMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest, cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SnapshotBackupRestore"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("SnapshotBackupRestore"))
              .build();
        }
      }
    }
    return getSnapshotBackupRestoreMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse> getGetStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetStatus",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse> getGetStatusMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse> getGetStatusMethod;
    if ((getGetStatusMethod = OSSnapshotServiceGrpc.getGetStatusMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getGetStatusMethod = OSSnapshotServiceGrpc.getGetStatusMethod) == null) {
          OSSnapshotServiceGrpc.getGetStatusMethod = getGetStatusMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("GetStatus"))
              .build();
        }
      }
    }
    return getGetStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getStopTransferMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StopTransfer",
      requestType = com.google.protobuf.Empty.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      com.google.protobuf.Empty> getStopTransferMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, com.google.protobuf.Empty> getStopTransferMethod;
    if ((getStopTransferMethod = OSSnapshotServiceGrpc.getStopTransferMethod) == null) {
      synchronized (OSSnapshotServiceGrpc.class) {
        if ((getStopTransferMethod = OSSnapshotServiceGrpc.getStopTransferMethod) == null) {
          OSSnapshotServiceGrpc.getStopTransferMethod = getStopTransferMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StopTransfer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new OSSnapshotServiceMethodDescriptorSupplier("StopTransfer"))
              .build();
        }
      }
    }
    return getStopTransferMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static OSSnapshotServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OSSnapshotServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OSSnapshotServiceStub>() {
        @java.lang.Override
        public OSSnapshotServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OSSnapshotServiceStub(channel, callOptions);
        }
      };
    return OSSnapshotServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static OSSnapshotServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OSSnapshotServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OSSnapshotServiceBlockingStub>() {
        @java.lang.Override
        public OSSnapshotServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OSSnapshotServiceBlockingStub(channel, callOptions);
        }
      };
    return OSSnapshotServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static OSSnapshotServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<OSSnapshotServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<OSSnapshotServiceFutureStub>() {
        @java.lang.Override
        public OSSnapshotServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new OSSnapshotServiceFutureStub(channel, callOptions);
        }
      };
    return OSSnapshotServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class OSSnapshotServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 添加快照备份池（必须是已存在的 ZFS 池）
     * </pre>
     */
    public void backupPoolAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBackupPoolAddMethod(), responseObserver);
    }

    /**
     * <pre>
     * 移除快照备份池（对应 ZFS 池不会被删除）
     * </pre>
     */
    public void backupPoolDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBackupPoolDelMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列举已注册的快照备份池
     * </pre>
     */
    public void backupPoolList(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getBackupPoolListMethod(), responseObserver);
    }

    /**
     * <pre>
     * 创建数据集并挂载到指定路径（若已存在，不会重复创建）
     * </pre>
     */
    public void datasetAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDatasetAddMethod(), responseObserver);
    }

    /**
     * <pre>
     * 删除数据集（及其所有快照）
     * </pre>
     */
    public void datasetDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDatasetDelMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列举所有数据集
     * </pre>
     */
    public void datasetList(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDatasetListMethod(), responseObserver);
    }

    /**
     * <pre>
     * 删除备份池中指定数据集（及其所有快照）
     * </pre>
     */
    public void datasetBackupDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDatasetBackupDelMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列举指定备份池中的所有数据集
     * </pre>
     */
    public void datasetBackupList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDatasetBackupListMethod(), responseObserver);
    }

    /**
     * <pre>
     * 为指定数据集创建快照（同一个数据集下的快照名称不能重复）
     * </pre>
     */
    public void snapshotAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapshotAddMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据名称删除指定数据集中的一个快照
     * </pre>
     */
    public void snapshotDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapshotDelMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列举指定数据集下所有快照
     * </pre>
     */
    public void snapshotList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapshotListMethod(), responseObserver);
    }

    /**
     * <pre>
     * 将数据集回滚到指定快照的状态（数据集中较新的快照会被丢弃）
     * </pre>
     */
    public void snapshotRestore(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapshotRestoreMethod(), responseObserver);
    }

    /**
     * <pre>
     * 备份快照到已注册的备份池（若对应数据集不存在，则会自动创建）
     * </pre>
     */
    public void snapshotBackupAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapshotBackupAddMethod(), responseObserver);
    }

    /**
     * <pre>
     * 将指定数据集快照从备份池中移除
     * </pre>
     */
    public void snapshotBackupDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapshotBackupDelMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列举指定备份池中某个数据集的快照
     * </pre>
     */
    public void snapshotBackupList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapshotBackupListMethod(), responseObserver);
    }

    /**
     * <pre>
     * 将某个快照备份还原
     * </pre>
     */
    public void snapshotBackupRestore(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSnapshotBackupRestoreMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取当前运行状态
     * </pre>
     */
    public void getStatus(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetStatusMethod(), responseObserver);
    }

    /**
     * <pre>
     * 停止当前正在进行的传输任务
     * </pre>
     */
    public void stopTransfer(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStopTransferMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getBackupPoolAddMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_BACKUP_POOL_ADD)))
          .addMethod(
            getBackupPoolDelMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_BACKUP_POOL_DEL)))
          .addMethod(
            getBackupPoolListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse>(
                  this, METHODID_BACKUP_POOL_LIST)))
          .addMethod(
            getDatasetAddMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_DATASET_ADD)))
          .addMethod(
            getDatasetDelMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_DATASET_DEL)))
          .addMethod(
            getDatasetListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse>(
                  this, METHODID_DATASET_LIST)))
          .addMethod(
            getDatasetBackupDelMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_DATASET_BACKUP_DEL)))
          .addMethod(
            getDatasetBackupListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest,
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse>(
                  this, METHODID_DATASET_BACKUP_LIST)))
          .addMethod(
            getSnapshotAddMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPSHOT_ADD)))
          .addMethod(
            getSnapshotDelMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPSHOT_DEL)))
          .addMethod(
            getSnapshotListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest,
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse>(
                  this, METHODID_SNAPSHOT_LIST)))
          .addMethod(
            getSnapshotRestoreMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPSHOT_RESTORE)))
          .addMethod(
            getSnapshotBackupAddMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest,
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse>(
                  this, METHODID_SNAPSHOT_BACKUP_ADD)))
          .addMethod(
            getSnapshotBackupDelMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SNAPSHOT_BACKUP_DEL)))
          .addMethod(
            getSnapshotBackupListMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest,
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse>(
                  this, METHODID_SNAPSHOT_BACKUP_LIST)))
          .addMethod(
            getSnapshotBackupRestoreMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest,
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse>(
                  this, METHODID_SNAPSHOT_BACKUP_RESTORE)))
          .addMethod(
            getGetStatusMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse>(
                  this, METHODID_GET_STATUS)))
          .addMethod(
            getStopTransferMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                com.google.protobuf.Empty>(
                  this, METHODID_STOP_TRANSFER)))
          .build();
    }
  }

  /**
   */
  public static final class OSSnapshotServiceStub extends io.grpc.stub.AbstractAsyncStub<OSSnapshotServiceStub> {
    private OSSnapshotServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OSSnapshotServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OSSnapshotServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 添加快照备份池（必须是已存在的 ZFS 池）
     * </pre>
     */
    public void backupPoolAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBackupPoolAddMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 移除快照备份池（对应 ZFS 池不会被删除）
     * </pre>
     */
    public void backupPoolDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBackupPoolDelMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列举已注册的快照备份池
     * </pre>
     */
    public void backupPoolList(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getBackupPoolListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 创建数据集并挂载到指定路径（若已存在，不会重复创建）
     * </pre>
     */
    public void datasetAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDatasetAddMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 删除数据集（及其所有快照）
     * </pre>
     */
    public void datasetDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDatasetDelMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列举所有数据集
     * </pre>
     */
    public void datasetList(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDatasetListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 删除备份池中指定数据集（及其所有快照）
     * </pre>
     */
    public void datasetBackupDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDatasetBackupDelMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列举指定备份池中的所有数据集
     * </pre>
     */
    public void datasetBackupList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDatasetBackupListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 为指定数据集创建快照（同一个数据集下的快照名称不能重复）
     * </pre>
     */
    public void snapshotAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapshotAddMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据名称删除指定数据集中的一个快照
     * </pre>
     */
    public void snapshotDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapshotDelMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列举指定数据集下所有快照
     * </pre>
     */
    public void snapshotList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapshotListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 将数据集回滚到指定快照的状态（数据集中较新的快照会被丢弃）
     * </pre>
     */
    public void snapshotRestore(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapshotRestoreMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 备份快照到已注册的备份池（若对应数据集不存在，则会自动创建）
     * </pre>
     */
    public void snapshotBackupAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapshotBackupAddMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 将指定数据集快照从备份池中移除
     * </pre>
     */
    public void snapshotBackupDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapshotBackupDelMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列举指定备份池中某个数据集的快照
     * </pre>
     */
    public void snapshotBackupList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapshotBackupListMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 将某个快照备份还原
     * </pre>
     */
    public void snapshotBackupRestore(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSnapshotBackupRestoreMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取当前运行状态
     * </pre>
     */
    public void getStatus(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 停止当前正在进行的传输任务
     * </pre>
     */
    public void stopTransfer(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStopTransferMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class OSSnapshotServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<OSSnapshotServiceBlockingStub> {
    private OSSnapshotServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OSSnapshotServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OSSnapshotServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 添加快照备份池（必须是已存在的 ZFS 池）
     * </pre>
     */
    public com.google.protobuf.Empty backupPoolAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBackupPoolAddMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 移除快照备份池（对应 ZFS 池不会被删除）
     * </pre>
     */
    public com.google.protobuf.Empty backupPoolDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBackupPoolDelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列举已注册的快照备份池
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse backupPoolList(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getBackupPoolListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 创建数据集并挂载到指定路径（若已存在，不会重复创建）
     * </pre>
     */
    public com.google.protobuf.Empty datasetAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDatasetAddMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 删除数据集（及其所有快照）
     * </pre>
     */
    public com.google.protobuf.Empty datasetDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDatasetDelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列举所有数据集
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse datasetList(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDatasetListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 删除备份池中指定数据集（及其所有快照）
     * </pre>
     */
    public com.google.protobuf.Empty datasetBackupDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDatasetBackupDelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列举指定备份池中的所有数据集
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse datasetBackupList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDatasetBackupListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 为指定数据集创建快照（同一个数据集下的快照名称不能重复）
     * </pre>
     */
    public com.google.protobuf.Empty snapshotAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapshotAddMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据名称删除指定数据集中的一个快照
     * </pre>
     */
    public com.google.protobuf.Empty snapshotDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapshotDelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列举指定数据集下所有快照
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse snapshotList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapshotListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 将数据集回滚到指定快照的状态（数据集中较新的快照会被丢弃）
     * </pre>
     */
    public com.google.protobuf.Empty snapshotRestore(cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapshotRestoreMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 备份快照到已注册的备份池（若对应数据集不存在，则会自动创建）
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse snapshotBackupAdd(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapshotBackupAddMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 将指定数据集快照从备份池中移除
     * </pre>
     */
    public com.google.protobuf.Empty snapshotBackupDel(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapshotBackupDelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列举指定备份池中某个数据集的快照
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse snapshotBackupList(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapshotBackupListMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 将某个快照备份还原
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse snapshotBackupRestore(cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSnapshotBackupRestoreMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取当前运行状态
     * </pre>
     */
    public cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse getStatus(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 停止当前正在进行的传输任务
     * </pre>
     */
    public com.google.protobuf.Empty stopTransfer(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStopTransferMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class OSSnapshotServiceFutureStub extends io.grpc.stub.AbstractFutureStub<OSSnapshotServiceFutureStub> {
    private OSSnapshotServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected OSSnapshotServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new OSSnapshotServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 添加快照备份池（必须是已存在的 ZFS 池）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> backupPoolAdd(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBackupPoolAddMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 移除快照备份池（对应 ZFS 池不会被删除）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> backupPoolDel(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBackupPoolDelMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列举已注册的快照备份池
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse> backupPoolList(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getBackupPoolListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 创建数据集并挂载到指定路径（若已存在，不会重复创建）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> datasetAdd(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDatasetAddMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 删除数据集（及其所有快照）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> datasetDel(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDatasetDelMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列举所有数据集
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> datasetList(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDatasetListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 删除备份池中指定数据集（及其所有快照）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> datasetBackupDel(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDatasetBackupDelMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列举指定备份池中的所有数据集
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse> datasetBackupList(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDatasetBackupListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 为指定数据集创建快照（同一个数据集下的快照名称不能重复）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapshotAdd(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapshotAddMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据名称删除指定数据集中的一个快照
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapshotDel(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapshotDelMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列举指定数据集下所有快照
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> snapshotList(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapshotListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 将数据集回滚到指定快照的状态（数据集中较新的快照会被丢弃）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapshotRestore(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapshotRestoreMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 备份快照到已注册的备份池（若对应数据集不存在，则会自动创建）
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> snapshotBackupAdd(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapshotBackupAddMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 将指定数据集快照从备份池中移除
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> snapshotBackupDel(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapshotBackupDelMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列举指定备份池中某个数据集的快照
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse> snapshotBackupList(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapshotBackupListMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 将某个快照备份还原
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse> snapshotBackupRestore(
        cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSnapshotBackupRestoreMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取当前运行状态
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse> getStatus(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetStatusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 停止当前正在进行的传输任务
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> stopTransfer(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStopTransferMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_BACKUP_POOL_ADD = 0;
  private static final int METHODID_BACKUP_POOL_DEL = 1;
  private static final int METHODID_BACKUP_POOL_LIST = 2;
  private static final int METHODID_DATASET_ADD = 3;
  private static final int METHODID_DATASET_DEL = 4;
  private static final int METHODID_DATASET_LIST = 5;
  private static final int METHODID_DATASET_BACKUP_DEL = 6;
  private static final int METHODID_DATASET_BACKUP_LIST = 7;
  private static final int METHODID_SNAPSHOT_ADD = 8;
  private static final int METHODID_SNAPSHOT_DEL = 9;
  private static final int METHODID_SNAPSHOT_LIST = 10;
  private static final int METHODID_SNAPSHOT_RESTORE = 11;
  private static final int METHODID_SNAPSHOT_BACKUP_ADD = 12;
  private static final int METHODID_SNAPSHOT_BACKUP_DEL = 13;
  private static final int METHODID_SNAPSHOT_BACKUP_LIST = 14;
  private static final int METHODID_SNAPSHOT_BACKUP_RESTORE = 15;
  private static final int METHODID_GET_STATUS = 16;
  private static final int METHODID_STOP_TRANSFER = 17;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final OSSnapshotServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(OSSnapshotServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_BACKUP_POOL_ADD:
          serviceImpl.backupPoolAdd((cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_BACKUP_POOL_DEL:
          serviceImpl.backupPoolDel((cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_BACKUP_POOL_LIST:
          serviceImpl.backupPoolList((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolListResponse>) responseObserver);
          break;
        case METHODID_DATASET_ADD:
          serviceImpl.datasetAdd((cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_DATASET_DEL:
          serviceImpl.datasetDel((cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_DATASET_LIST:
          serviceImpl.datasetList((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse>) responseObserver);
          break;
        case METHODID_DATASET_BACKUP_DEL:
          serviceImpl.datasetBackupDel((cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetBackupRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_DATASET_BACKUP_LIST:
          serviceImpl.datasetBackupList((cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupPoolRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetListResponse>) responseObserver);
          break;
        case METHODID_SNAPSHOT_ADD:
          serviceImpl.snapshotAdd((cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPSHOT_DEL:
          serviceImpl.snapshotDel((cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPSHOT_LIST:
          serviceImpl.snapshotList((cloud.lazycat.apis.sys.OSSnapshot.SnapshotDatasetRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse>) responseObserver);
          break;
        case METHODID_SNAPSHOT_RESTORE:
          serviceImpl.snapshotRestore((cloud.lazycat.apis.sys.OSSnapshot.SnapshotRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPSHOT_BACKUP_ADD:
          serviceImpl.snapshotBackupAdd((cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse>) responseObserver);
          break;
        case METHODID_SNAPSHOT_BACKUP_DEL:
          serviceImpl.snapshotBackupDel((cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_SNAPSHOT_BACKUP_LIST:
          serviceImpl.snapshotBackupList((cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupListRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotListResponse>) responseObserver);
          break;
        case METHODID_SNAPSHOT_BACKUP_RESTORE:
          serviceImpl.snapshotBackupRestore((cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotBackupTransferResponse>) responseObserver);
          break;
        case METHODID_GET_STATUS:
          serviceImpl.getStatus((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.OSSnapshot.SnapshotStatusResponse>) responseObserver);
          break;
        case METHODID_STOP_TRANSFER:
          serviceImpl.stopTransfer((com.google.protobuf.Empty) request,
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

  private static abstract class OSSnapshotServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    OSSnapshotServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.OSSnapshot.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("OSSnapshotService");
    }
  }

  private static final class OSSnapshotServiceFileDescriptorSupplier
      extends OSSnapshotServiceBaseDescriptorSupplier {
    OSSnapshotServiceFileDescriptorSupplier() {}
  }

  private static final class OSSnapshotServiceMethodDescriptorSupplier
      extends OSSnapshotServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    OSSnapshotServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (OSSnapshotServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new OSSnapshotServiceFileDescriptorSupplier())
              .addMethod(getBackupPoolAddMethod())
              .addMethod(getBackupPoolDelMethod())
              .addMethod(getBackupPoolListMethod())
              .addMethod(getDatasetAddMethod())
              .addMethod(getDatasetDelMethod())
              .addMethod(getDatasetListMethod())
              .addMethod(getDatasetBackupDelMethod())
              .addMethod(getDatasetBackupListMethod())
              .addMethod(getSnapshotAddMethod())
              .addMethod(getSnapshotDelMethod())
              .addMethod(getSnapshotListMethod())
              .addMethod(getSnapshotRestoreMethod())
              .addMethod(getSnapshotBackupAddMethod())
              .addMethod(getSnapshotBackupDelMethod())
              .addMethod(getSnapshotBackupListMethod())
              .addMethod(getSnapshotBackupRestoreMethod())
              .addMethod(getGetStatusMethod())
              .addMethod(getStopTransferMethod())
              .build();
        }
      }
    }
    return result;
  }
}
