package cloud.lazycat.apis.common;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Task status 
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: common/filetrans.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class FileTransferServiceGrpc {

  private FileTransferServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.common.FileTransferService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID,
      cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp> getCreateQueueMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateQueue",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskQueueID.class,
      responseType = cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID,
      cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp> getCreateQueueMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID, cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp> getCreateQueueMethod;
    if ((getCreateQueueMethod = FileTransferServiceGrpc.getCreateQueueMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getCreateQueueMethod = FileTransferServiceGrpc.getCreateQueueMethod) == null) {
          FileTransferServiceGrpc.getCreateQueueMethod = getCreateQueueMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskQueueID, cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateQueue"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskQueueID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("CreateQueue"))
              .build();
        }
      }
    }
    return getCreateQueueMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Filetrans.TaskQueueListResp> getListQueueMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListQueue",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.common.Filetrans.TaskQueueListResp.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.common.Filetrans.TaskQueueListResp> getListQueueMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.common.Filetrans.TaskQueueListResp> getListQueueMethod;
    if ((getListQueueMethod = FileTransferServiceGrpc.getListQueueMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getListQueueMethod = FileTransferServiceGrpc.getListQueueMethod) == null) {
          FileTransferServiceGrpc.getListQueueMethod = getListQueueMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.common.Filetrans.TaskQueueListResp>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListQueue"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskQueueListResp.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("ListQueue"))
              .build();
        }
      }
    }
    return getListQueueMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
      cloud.lazycat.apis.common.Filetrans.Task> getQueryQueueMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryQueue",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq.class,
      responseType = cloud.lazycat.apis.common.Filetrans.Task.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
      cloud.lazycat.apis.common.Filetrans.Task> getQueryQueueMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq, cloud.lazycat.apis.common.Filetrans.Task> getQueryQueueMethod;
    if ((getQueryQueueMethod = FileTransferServiceGrpc.getQueryQueueMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getQueryQueueMethod = FileTransferServiceGrpc.getQueryQueueMethod) == null) {
          FileTransferServiceGrpc.getQueryQueueMethod = getQueryQueueMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq, cloud.lazycat.apis.common.Filetrans.Task>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryQueue"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.Task.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("QueryQueue"))
              .build();
        }
      }
    }
    return getQueryQueueMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
      com.google.protobuf.Empty> getClearQueueMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClearQueue",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
      com.google.protobuf.Empty> getClearQueueMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq, com.google.protobuf.Empty> getClearQueueMethod;
    if ((getClearQueueMethod = FileTransferServiceGrpc.getClearQueueMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getClearQueueMethod = FileTransferServiceGrpc.getClearQueueMethod) == null) {
          FileTransferServiceGrpc.getClearQueueMethod = getClearQueueMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClearQueue"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("ClearQueue"))
              .build();
        }
      }
    }
    return getClearQueueMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq,
      com.google.protobuf.Empty> getConfigQueueMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ConfigQueue",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq,
      com.google.protobuf.Empty> getConfigQueueMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq, com.google.protobuf.Empty> getConfigQueueMethod;
    if ((getConfigQueueMethod = FileTransferServiceGrpc.getConfigQueueMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getConfigQueueMethod = FileTransferServiceGrpc.getConfigQueueMethod) == null) {
          FileTransferServiceGrpc.getConfigQueueMethod = getConfigQueueMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ConfigQueue"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("ConfigQueue"))
              .build();
        }
      }
    }
    return getConfigQueueMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID,
      com.google.protobuf.Empty> getPauseQueueMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PauseQueue",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskQueueID.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID,
      com.google.protobuf.Empty> getPauseQueueMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID, com.google.protobuf.Empty> getPauseQueueMethod;
    if ((getPauseQueueMethod = FileTransferServiceGrpc.getPauseQueueMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getPauseQueueMethod = FileTransferServiceGrpc.getPauseQueueMethod) == null) {
          FileTransferServiceGrpc.getPauseQueueMethod = getPauseQueueMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskQueueID, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PauseQueue"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskQueueID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("PauseQueue"))
              .build();
        }
      }
    }
    return getPauseQueueMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID,
      com.google.protobuf.Empty> getStartQuqueMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StartQuque",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskQueueID.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID,
      com.google.protobuf.Empty> getStartQuqueMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueID, com.google.protobuf.Empty> getStartQuqueMethod;
    if ((getStartQuqueMethod = FileTransferServiceGrpc.getStartQuqueMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getStartQuqueMethod = FileTransferServiceGrpc.getStartQuqueMethod) == null) {
          FileTransferServiceGrpc.getStartQuqueMethod = getStartQuqueMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskQueueID, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StartQuque"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskQueueID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("StartQuque"))
              .build();
        }
      }
    }
    return getStartQuqueMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
      cloud.lazycat.apis.common.Filetrans.QueueMessageResp> getQueryQueueMessageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryQueueMessage",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq.class,
      responseType = cloud.lazycat.apis.common.Filetrans.QueueMessageResp.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
      cloud.lazycat.apis.common.Filetrans.QueueMessageResp> getQueryQueueMessageMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq, cloud.lazycat.apis.common.Filetrans.QueueMessageResp> getQueryQueueMessageMethod;
    if ((getQueryQueueMessageMethod = FileTransferServiceGrpc.getQueryQueueMessageMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getQueryQueueMessageMethod = FileTransferServiceGrpc.getQueryQueueMessageMethod) == null) {
          FileTransferServiceGrpc.getQueryQueueMessageMethod = getQueryQueueMessageMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq, cloud.lazycat.apis.common.Filetrans.QueueMessageResp>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryQueueMessage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.QueueMessageResp.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("QueryQueueMessage"))
              .build();
        }
      }
    }
    return getQueryQueueMessageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskCreateRequest,
      cloud.lazycat.apis.common.Filetrans.Task> getCreateTaskMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateTask",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskCreateRequest.class,
      responseType = cloud.lazycat.apis.common.Filetrans.Task.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskCreateRequest,
      cloud.lazycat.apis.common.Filetrans.Task> getCreateTaskMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskCreateRequest, cloud.lazycat.apis.common.Filetrans.Task> getCreateTaskMethod;
    if ((getCreateTaskMethod = FileTransferServiceGrpc.getCreateTaskMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getCreateTaskMethod = FileTransferServiceGrpc.getCreateTaskMethod) == null) {
          FileTransferServiceGrpc.getCreateTaskMethod = getCreateTaskMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskCreateRequest, cloud.lazycat.apis.common.Filetrans.Task>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateTask"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskCreateRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.Task.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("CreateTask"))
              .build();
        }
      }
    }
    return getCreateTaskMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskCreateRequests,
      cloud.lazycat.apis.common.Filetrans.Task> getCreateTasksMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateTasks",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskCreateRequests.class,
      responseType = cloud.lazycat.apis.common.Filetrans.Task.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskCreateRequests,
      cloud.lazycat.apis.common.Filetrans.Task> getCreateTasksMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskCreateRequests, cloud.lazycat.apis.common.Filetrans.Task> getCreateTasksMethod;
    if ((getCreateTasksMethod = FileTransferServiceGrpc.getCreateTasksMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getCreateTasksMethod = FileTransferServiceGrpc.getCreateTasksMethod) == null) {
          FileTransferServiceGrpc.getCreateTasksMethod = getCreateTasksMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskCreateRequests, cloud.lazycat.apis.common.Filetrans.Task>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateTasks"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskCreateRequests.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.Task.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("CreateTasks"))
              .build();
        }
      }
    }
    return getCreateTasksMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID,
      cloud.lazycat.apis.common.Filetrans.Task> getQueryTaskMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryTask",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskID.class,
      responseType = cloud.lazycat.apis.common.Filetrans.Task.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID,
      cloud.lazycat.apis.common.Filetrans.Task> getQueryTaskMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID, cloud.lazycat.apis.common.Filetrans.Task> getQueryTaskMethod;
    if ((getQueryTaskMethod = FileTransferServiceGrpc.getQueryTaskMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getQueryTaskMethod = FileTransferServiceGrpc.getQueryTaskMethod) == null) {
          FileTransferServiceGrpc.getQueryTaskMethod = getQueryTaskMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskID, cloud.lazycat.apis.common.Filetrans.Task>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryTask"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.Task.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("QueryTask"))
              .build();
        }
      }
    }
    return getQueryTaskMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID,
      com.google.protobuf.Empty> getResumeTaskMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ResumeTask",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskID.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID,
      com.google.protobuf.Empty> getResumeTaskMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID, com.google.protobuf.Empty> getResumeTaskMethod;
    if ((getResumeTaskMethod = FileTransferServiceGrpc.getResumeTaskMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getResumeTaskMethod = FileTransferServiceGrpc.getResumeTaskMethod) == null) {
          FileTransferServiceGrpc.getResumeTaskMethod = getResumeTaskMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskID, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ResumeTask"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("ResumeTask"))
              .build();
        }
      }
    }
    return getResumeTaskMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID,
      com.google.protobuf.Empty> getPauseTaskMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PauseTask",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskID.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID,
      com.google.protobuf.Empty> getPauseTaskMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID, com.google.protobuf.Empty> getPauseTaskMethod;
    if ((getPauseTaskMethod = FileTransferServiceGrpc.getPauseTaskMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getPauseTaskMethod = FileTransferServiceGrpc.getPauseTaskMethod) == null) {
          FileTransferServiceGrpc.getPauseTaskMethod = getPauseTaskMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskID, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PauseTask"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("PauseTask"))
              .build();
        }
      }
    }
    return getPauseTaskMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID,
      com.google.protobuf.Empty> getDeleteTaskMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteTask",
      requestType = cloud.lazycat.apis.common.Filetrans.TaskID.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID,
      com.google.protobuf.Empty> getDeleteTaskMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.common.Filetrans.TaskID, com.google.protobuf.Empty> getDeleteTaskMethod;
    if ((getDeleteTaskMethod = FileTransferServiceGrpc.getDeleteTaskMethod) == null) {
      synchronized (FileTransferServiceGrpc.class) {
        if ((getDeleteTaskMethod = FileTransferServiceGrpc.getDeleteTaskMethod) == null) {
          FileTransferServiceGrpc.getDeleteTaskMethod = getDeleteTaskMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.common.Filetrans.TaskID, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteTask"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.common.Filetrans.TaskID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new FileTransferServiceMethodDescriptorSupplier("DeleteTask"))
              .build();
        }
      }
    }
    return getDeleteTaskMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static FileTransferServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileTransferServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileTransferServiceStub>() {
        @java.lang.Override
        public FileTransferServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileTransferServiceStub(channel, callOptions);
        }
      };
    return FileTransferServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static FileTransferServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileTransferServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileTransferServiceBlockingStub>() {
        @java.lang.Override
        public FileTransferServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileTransferServiceBlockingStub(channel, callOptions);
        }
      };
    return FileTransferServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static FileTransferServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FileTransferServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FileTransferServiceFutureStub>() {
        @java.lang.Override
        public FileTransferServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FileTransferServiceFutureStub(channel, callOptions);
        }
      };
    return FileTransferServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * Task status 
   * </pre>
   */
  public static abstract class FileTransferServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 创建队列
     * </pre>
     */
    public void createQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateQueueMethod(), responseObserver);
    }

    /**
     * <pre>
     * 列出所有 QueueID
     * </pre>
     */
    public void listQueue(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.TaskQueueListResp> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListQueueMethod(), responseObserver);
    }

    /**
     * <pre>
     * 通过队列的 ID 和 Status 获取任务列表
     * </pre>
     */
    public void queryQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryQueueMethod(), responseObserver);
    }

    /**
     * <pre>
     * 通过队列的 ID 和 Status 清除任务
     * </pre>
     */
    public void clearQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getClearQueueMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据队列 ID 设置队列的速率并发等设置
     * </pre>
     */
    public void configQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getConfigQueueMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据队列 ID 暂停队列
     * </pre>
     */
    public void pauseQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPauseQueueMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据队列 ID 开始队列
     * </pre>
     */
    public void startQuque(cloud.lazycat.apis.common.Filetrans.TaskQueueID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStartQuqueMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据队列 ID 获取队列动态信息（比如 pending 状态的增加了 task1）
     * </pre>
     */
    public void queryQueueMessage(cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.QueueMessageResp> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryQueueMessageMethod(), responseObserver);
    }

    /**
     * <pre>
     * 暂时不支持创建Task时创建任务，需要提前创建好任务。queue_id不存在则报错
     * </pre>
     */
    public void createTask(cloud.lazycat.apis.common.Filetrans.TaskCreateRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateTaskMethod(), responseObserver);
    }

    /**
     */
    public void createTasks(cloud.lazycat.apis.common.Filetrans.TaskCreateRequests request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateTasksMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据 ID 获取单个任务状态
     * </pre>
     */
    public void queryTask(cloud.lazycat.apis.common.Filetrans.TaskID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryTaskMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据 ID 开始单个任务
     * </pre>
     */
    public void resumeTask(cloud.lazycat.apis.common.Filetrans.TaskID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getResumeTaskMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据 ID 暂停单个任务
     * </pre>
     */
    public void pauseTask(cloud.lazycat.apis.common.Filetrans.TaskID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPauseTaskMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据 ID 删除单个任务
     * </pre>
     */
    public void deleteTask(cloud.lazycat.apis.common.Filetrans.TaskID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteTaskMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getCreateQueueMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskQueueID,
                cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp>(
                  this, METHODID_CREATE_QUEUE)))
          .addMethod(
            getListQueueMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.common.Filetrans.TaskQueueListResp>(
                  this, METHODID_LIST_QUEUE)))
          .addMethod(
            getQueryQueueMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
                cloud.lazycat.apis.common.Filetrans.Task>(
                  this, METHODID_QUERY_QUEUE)))
          .addMethod(
            getClearQueueMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
                com.google.protobuf.Empty>(
                  this, METHODID_CLEAR_QUEUE)))
          .addMethod(
            getConfigQueueMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq,
                com.google.protobuf.Empty>(
                  this, METHODID_CONFIG_QUEUE)))
          .addMethod(
            getPauseQueueMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskQueueID,
                com.google.protobuf.Empty>(
                  this, METHODID_PAUSE_QUEUE)))
          .addMethod(
            getStartQuqueMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskQueueID,
                com.google.protobuf.Empty>(
                  this, METHODID_START_QUQUE)))
          .addMethod(
            getQueryQueueMessageMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq,
                cloud.lazycat.apis.common.Filetrans.QueueMessageResp>(
                  this, METHODID_QUERY_QUEUE_MESSAGE)))
          .addMethod(
            getCreateTaskMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskCreateRequest,
                cloud.lazycat.apis.common.Filetrans.Task>(
                  this, METHODID_CREATE_TASK)))
          .addMethod(
            getCreateTasksMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskCreateRequests,
                cloud.lazycat.apis.common.Filetrans.Task>(
                  this, METHODID_CREATE_TASKS)))
          .addMethod(
            getQueryTaskMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskID,
                cloud.lazycat.apis.common.Filetrans.Task>(
                  this, METHODID_QUERY_TASK)))
          .addMethod(
            getResumeTaskMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskID,
                com.google.protobuf.Empty>(
                  this, METHODID_RESUME_TASK)))
          .addMethod(
            getPauseTaskMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskID,
                com.google.protobuf.Empty>(
                  this, METHODID_PAUSE_TASK)))
          .addMethod(
            getDeleteTaskMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.common.Filetrans.TaskID,
                com.google.protobuf.Empty>(
                  this, METHODID_DELETE_TASK)))
          .build();
    }
  }

  /**
   * <pre>
   * Task status 
   * </pre>
   */
  public static final class FileTransferServiceStub extends io.grpc.stub.AbstractAsyncStub<FileTransferServiceStub> {
    private FileTransferServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileTransferServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileTransferServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * 创建队列
     * </pre>
     */
    public void createQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateQueueMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 列出所有 QueueID
     * </pre>
     */
    public void listQueue(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.TaskQueueListResp> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListQueueMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 通过队列的 ID 和 Status 获取任务列表
     * </pre>
     */
    public void queryQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getQueryQueueMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 通过队列的 ID 和 Status 清除任务
     * </pre>
     */
    public void clearQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getClearQueueMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据队列 ID 设置队列的速率并发等设置
     * </pre>
     */
    public void configQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getConfigQueueMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据队列 ID 暂停队列
     * </pre>
     */
    public void pauseQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPauseQueueMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据队列 ID 开始队列
     * </pre>
     */
    public void startQuque(cloud.lazycat.apis.common.Filetrans.TaskQueueID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStartQuqueMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据队列 ID 获取队列动态信息（比如 pending 状态的增加了 task1）
     * </pre>
     */
    public void queryQueueMessage(cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.QueueMessageResp> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getQueryQueueMessageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 暂时不支持创建Task时创建任务，需要提前创建好任务。queue_id不存在则报错
     * </pre>
     */
    public void createTask(cloud.lazycat.apis.common.Filetrans.TaskCreateRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateTaskMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createTasks(cloud.lazycat.apis.common.Filetrans.TaskCreateRequests request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getCreateTasksMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据 ID 获取单个任务状态
     * </pre>
     */
    public void queryTask(cloud.lazycat.apis.common.Filetrans.TaskID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getQueryTaskMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据 ID 开始单个任务
     * </pre>
     */
    public void resumeTask(cloud.lazycat.apis.common.Filetrans.TaskID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getResumeTaskMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据 ID 暂停单个任务
     * </pre>
     */
    public void pauseTask(cloud.lazycat.apis.common.Filetrans.TaskID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPauseTaskMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据 ID 删除单个任务
     * </pre>
     */
    public void deleteTask(cloud.lazycat.apis.common.Filetrans.TaskID request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteTaskMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * Task status 
   * </pre>
   */
  public static final class FileTransferServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<FileTransferServiceBlockingStub> {
    private FileTransferServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileTransferServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileTransferServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 创建队列
     * </pre>
     */
    public cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp createQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateQueueMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 列出所有 QueueID
     * </pre>
     */
    public cloud.lazycat.apis.common.Filetrans.TaskQueueListResp listQueue(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListQueueMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 通过队列的 ID 和 Status 获取任务列表
     * </pre>
     */
    public java.util.Iterator<cloud.lazycat.apis.common.Filetrans.Task> queryQueue(
        cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getQueryQueueMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 通过队列的 ID 和 Status 清除任务
     * </pre>
     */
    public com.google.protobuf.Empty clearQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getClearQueueMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据队列 ID 设置队列的速率并发等设置
     * </pre>
     */
    public com.google.protobuf.Empty configQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getConfigQueueMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据队列 ID 暂停队列
     * </pre>
     */
    public com.google.protobuf.Empty pauseQueue(cloud.lazycat.apis.common.Filetrans.TaskQueueID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPauseQueueMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据队列 ID 开始队列
     * </pre>
     */
    public com.google.protobuf.Empty startQuque(cloud.lazycat.apis.common.Filetrans.TaskQueueID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStartQuqueMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据队列 ID 获取队列动态信息（比如 pending 状态的增加了 task1）
     * </pre>
     */
    public java.util.Iterator<cloud.lazycat.apis.common.Filetrans.QueueMessageResp> queryQueueMessage(
        cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getQueryQueueMessageMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 暂时不支持创建Task时创建任务，需要提前创建好任务。queue_id不存在则报错
     * </pre>
     */
    public cloud.lazycat.apis.common.Filetrans.Task createTask(cloud.lazycat.apis.common.Filetrans.TaskCreateRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateTaskMethod(), getCallOptions(), request);
    }

    /**
     */
    public java.util.Iterator<cloud.lazycat.apis.common.Filetrans.Task> createTasks(
        cloud.lazycat.apis.common.Filetrans.TaskCreateRequests request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getCreateTasksMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据 ID 获取单个任务状态
     * </pre>
     */
    public java.util.Iterator<cloud.lazycat.apis.common.Filetrans.Task> queryTask(
        cloud.lazycat.apis.common.Filetrans.TaskID request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getQueryTaskMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据 ID 开始单个任务
     * </pre>
     */
    public com.google.protobuf.Empty resumeTask(cloud.lazycat.apis.common.Filetrans.TaskID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getResumeTaskMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据 ID 暂停单个任务
     * </pre>
     */
    public com.google.protobuf.Empty pauseTask(cloud.lazycat.apis.common.Filetrans.TaskID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPauseTaskMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据 ID 删除单个任务
     * </pre>
     */
    public com.google.protobuf.Empty deleteTask(cloud.lazycat.apis.common.Filetrans.TaskID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteTaskMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * Task status 
   * </pre>
   */
  public static final class FileTransferServiceFutureStub extends io.grpc.stub.AbstractFutureStub<FileTransferServiceFutureStub> {
    private FileTransferServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileTransferServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FileTransferServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 创建队列
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp> createQueue(
        cloud.lazycat.apis.common.Filetrans.TaskQueueID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateQueueMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 列出所有 QueueID
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Filetrans.TaskQueueListResp> listQueue(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListQueueMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 通过队列的 ID 和 Status 清除任务
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> clearQueue(
        cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getClearQueueMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据队列 ID 设置队列的速率并发等设置
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> configQueue(
        cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getConfigQueueMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据队列 ID 暂停队列
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> pauseQueue(
        cloud.lazycat.apis.common.Filetrans.TaskQueueID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPauseQueueMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据队列 ID 开始队列
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> startQuque(
        cloud.lazycat.apis.common.Filetrans.TaskQueueID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStartQuqueMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 暂时不支持创建Task时创建任务，需要提前创建好任务。queue_id不存在则报错
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.common.Filetrans.Task> createTask(
        cloud.lazycat.apis.common.Filetrans.TaskCreateRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateTaskMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据 ID 开始单个任务
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> resumeTask(
        cloud.lazycat.apis.common.Filetrans.TaskID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getResumeTaskMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据 ID 暂停单个任务
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> pauseTask(
        cloud.lazycat.apis.common.Filetrans.TaskID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPauseTaskMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据 ID 删除单个任务
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> deleteTask(
        cloud.lazycat.apis.common.Filetrans.TaskID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteTaskMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_QUEUE = 0;
  private static final int METHODID_LIST_QUEUE = 1;
  private static final int METHODID_QUERY_QUEUE = 2;
  private static final int METHODID_CLEAR_QUEUE = 3;
  private static final int METHODID_CONFIG_QUEUE = 4;
  private static final int METHODID_PAUSE_QUEUE = 5;
  private static final int METHODID_START_QUQUE = 6;
  private static final int METHODID_QUERY_QUEUE_MESSAGE = 7;
  private static final int METHODID_CREATE_TASK = 8;
  private static final int METHODID_CREATE_TASKS = 9;
  private static final int METHODID_QUERY_TASK = 10;
  private static final int METHODID_RESUME_TASK = 11;
  private static final int METHODID_PAUSE_TASK = 12;
  private static final int METHODID_DELETE_TASK = 13;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final FileTransferServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(FileTransferServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE_QUEUE:
          serviceImpl.createQueue((cloud.lazycat.apis.common.Filetrans.TaskQueueID) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.FileTaskQueueResp>) responseObserver);
          break;
        case METHODID_LIST_QUEUE:
          serviceImpl.listQueue((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.TaskQueueListResp>) responseObserver);
          break;
        case METHODID_QUERY_QUEUE:
          serviceImpl.queryQueue((cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task>) responseObserver);
          break;
        case METHODID_CLEAR_QUEUE:
          serviceImpl.clearQueue((cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_CONFIG_QUEUE:
          serviceImpl.configQueue((cloud.lazycat.apis.common.Filetrans.TaskQueueConfigReq) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_PAUSE_QUEUE:
          serviceImpl.pauseQueue((cloud.lazycat.apis.common.Filetrans.TaskQueueID) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_START_QUQUE:
          serviceImpl.startQuque((cloud.lazycat.apis.common.Filetrans.TaskQueueID) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_QUERY_QUEUE_MESSAGE:
          serviceImpl.queryQueueMessage((cloud.lazycat.apis.common.Filetrans.TaskQueueQueryReq) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.QueueMessageResp>) responseObserver);
          break;
        case METHODID_CREATE_TASK:
          serviceImpl.createTask((cloud.lazycat.apis.common.Filetrans.TaskCreateRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task>) responseObserver);
          break;
        case METHODID_CREATE_TASKS:
          serviceImpl.createTasks((cloud.lazycat.apis.common.Filetrans.TaskCreateRequests) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task>) responseObserver);
          break;
        case METHODID_QUERY_TASK:
          serviceImpl.queryTask((cloud.lazycat.apis.common.Filetrans.TaskID) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.common.Filetrans.Task>) responseObserver);
          break;
        case METHODID_RESUME_TASK:
          serviceImpl.resumeTask((cloud.lazycat.apis.common.Filetrans.TaskID) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_PAUSE_TASK:
          serviceImpl.pauseTask((cloud.lazycat.apis.common.Filetrans.TaskID) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_DELETE_TASK:
          serviceImpl.deleteTask((cloud.lazycat.apis.common.Filetrans.TaskID) request,
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

  private static abstract class FileTransferServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    FileTransferServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.common.Filetrans.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("FileTransferService");
    }
  }

  private static final class FileTransferServiceFileDescriptorSupplier
      extends FileTransferServiceBaseDescriptorSupplier {
    FileTransferServiceFileDescriptorSupplier() {}
  }

  private static final class FileTransferServiceMethodDescriptorSupplier
      extends FileTransferServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    FileTransferServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (FileTransferServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new FileTransferServiceFileDescriptorSupplier())
              .addMethod(getCreateQueueMethod())
              .addMethod(getListQueueMethod())
              .addMethod(getQueryQueueMethod())
              .addMethod(getClearQueueMethod())
              .addMethod(getConfigQueueMethod())
              .addMethod(getPauseQueueMethod())
              .addMethod(getStartQuqueMethod())
              .addMethod(getQueryQueueMessageMethod())
              .addMethod(getCreateTaskMethod())
              .addMethod(getCreateTasksMethod())
              .addMethod(getQueryTaskMethod())
              .addMethod(getResumeTaskMethod())
              .addMethod(getPauseTaskMethod())
              .addMethod(getDeleteTaskMethod())
              .build();
        }
      }
    }
    return result;
  }
}
