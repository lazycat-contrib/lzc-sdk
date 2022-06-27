package cloud.lazycat.localdevice.apis;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: localdevice/dialog/dialog.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class DialogManagerGrpc {

  private DialogManagerGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.localdevice.apis.DialogManager";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.QuestionRequest,
      cloud.lazycat.localdevice.apis.Dialog.QuestionResult> getQuestionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Question",
      requestType = cloud.lazycat.localdevice.apis.Dialog.QuestionRequest.class,
      responseType = cloud.lazycat.localdevice.apis.Dialog.QuestionResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.QuestionRequest,
      cloud.lazycat.localdevice.apis.Dialog.QuestionResult> getQuestionMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.QuestionRequest, cloud.lazycat.localdevice.apis.Dialog.QuestionResult> getQuestionMethod;
    if ((getQuestionMethod = DialogManagerGrpc.getQuestionMethod) == null) {
      synchronized (DialogManagerGrpc.class) {
        if ((getQuestionMethod = DialogManagerGrpc.getQuestionMethod) == null) {
          DialogManagerGrpc.getQuestionMethod = getQuestionMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.localdevice.apis.Dialog.QuestionRequest, cloud.lazycat.localdevice.apis.Dialog.QuestionResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Question"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Dialog.QuestionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Dialog.QuestionResult.getDefaultInstance()))
              .setSchemaDescriptor(new DialogManagerMethodDescriptorSupplier("Question"))
              .build();
        }
      }
    }
    return getQuestionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest,
      com.google.protobuf.Empty> getMessageBoxMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "MessageBox",
      requestType = cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest,
      com.google.protobuf.Empty> getMessageBoxMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest, com.google.protobuf.Empty> getMessageBoxMethod;
    if ((getMessageBoxMethod = DialogManagerGrpc.getMessageBoxMethod) == null) {
      synchronized (DialogManagerGrpc.class) {
        if ((getMessageBoxMethod = DialogManagerGrpc.getMessageBoxMethod) == null) {
          DialogManagerGrpc.getMessageBoxMethod = getMessageBoxMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "MessageBox"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new DialogManagerMethodDescriptorSupplier("MessageBox"))
              .build();
        }
      }
    }
    return getMessageBoxMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.PasswordRequest,
      cloud.lazycat.localdevice.apis.Dialog.PasswordResult> getPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Password",
      requestType = cloud.lazycat.localdevice.apis.Dialog.PasswordRequest.class,
      responseType = cloud.lazycat.localdevice.apis.Dialog.PasswordResult.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.PasswordRequest,
      cloud.lazycat.localdevice.apis.Dialog.PasswordResult> getPasswordMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.localdevice.apis.Dialog.PasswordRequest, cloud.lazycat.localdevice.apis.Dialog.PasswordResult> getPasswordMethod;
    if ((getPasswordMethod = DialogManagerGrpc.getPasswordMethod) == null) {
      synchronized (DialogManagerGrpc.class) {
        if ((getPasswordMethod = DialogManagerGrpc.getPasswordMethod) == null) {
          DialogManagerGrpc.getPasswordMethod = getPasswordMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.localdevice.apis.Dialog.PasswordRequest, cloud.lazycat.localdevice.apis.Dialog.PasswordResult>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Password"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Dialog.PasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.localdevice.apis.Dialog.PasswordResult.getDefaultInstance()))
              .setSchemaDescriptor(new DialogManagerMethodDescriptorSupplier("Password"))
              .build();
        }
      }
    }
    return getPasswordMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DialogManagerStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DialogManagerStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DialogManagerStub>() {
        @java.lang.Override
        public DialogManagerStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DialogManagerStub(channel, callOptions);
        }
      };
    return DialogManagerStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DialogManagerBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DialogManagerBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DialogManagerBlockingStub>() {
        @java.lang.Override
        public DialogManagerBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DialogManagerBlockingStub(channel, callOptions);
        }
      };
    return DialogManagerBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static DialogManagerFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<DialogManagerFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<DialogManagerFutureStub>() {
        @java.lang.Override
        public DialogManagerFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new DialogManagerFutureStub(channel, callOptions);
        }
      };
    return DialogManagerFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class DialogManagerImplBase implements io.grpc.BindableService {

    /**
     */
    public void question(cloud.lazycat.localdevice.apis.Dialog.QuestionRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Dialog.QuestionResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQuestionMethod(), responseObserver);
    }

    /**
     */
    public void messageBox(cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getMessageBoxMethod(), responseObserver);
    }

    /**
     */
    public void password(cloud.lazycat.localdevice.apis.Dialog.PasswordRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Dialog.PasswordResult> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPasswordMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getQuestionMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.localdevice.apis.Dialog.QuestionRequest,
                cloud.lazycat.localdevice.apis.Dialog.QuestionResult>(
                  this, METHODID_QUESTION)))
          .addMethod(
            getMessageBoxMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_MESSAGE_BOX)))
          .addMethod(
            getPasswordMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.localdevice.apis.Dialog.PasswordRequest,
                cloud.lazycat.localdevice.apis.Dialog.PasswordResult>(
                  this, METHODID_PASSWORD)))
          .build();
    }
  }

  /**
   */
  public static final class DialogManagerStub extends io.grpc.stub.AbstractAsyncStub<DialogManagerStub> {
    private DialogManagerStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DialogManagerStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DialogManagerStub(channel, callOptions);
    }

    /**
     */
    public void question(cloud.lazycat.localdevice.apis.Dialog.QuestionRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Dialog.QuestionResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQuestionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void messageBox(cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getMessageBoxMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void password(cloud.lazycat.localdevice.apis.Dialog.PasswordRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Dialog.PasswordResult> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPasswordMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class DialogManagerBlockingStub extends io.grpc.stub.AbstractBlockingStub<DialogManagerBlockingStub> {
    private DialogManagerBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DialogManagerBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DialogManagerBlockingStub(channel, callOptions);
    }

    /**
     */
    public cloud.lazycat.localdevice.apis.Dialog.QuestionResult question(cloud.lazycat.localdevice.apis.Dialog.QuestionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQuestionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty messageBox(cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getMessageBoxMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.localdevice.apis.Dialog.PasswordResult password(cloud.lazycat.localdevice.apis.Dialog.PasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPasswordMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class DialogManagerFutureStub extends io.grpc.stub.AbstractFutureStub<DialogManagerFutureStub> {
    private DialogManagerFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DialogManagerFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new DialogManagerFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.localdevice.apis.Dialog.QuestionResult> question(
        cloud.lazycat.localdevice.apis.Dialog.QuestionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQuestionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> messageBox(
        cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getMessageBoxMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.localdevice.apis.Dialog.PasswordResult> password(
        cloud.lazycat.localdevice.apis.Dialog.PasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPasswordMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_QUESTION = 0;
  private static final int METHODID_MESSAGE_BOX = 1;
  private static final int METHODID_PASSWORD = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DialogManagerImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DialogManagerImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_QUESTION:
          serviceImpl.question((cloud.lazycat.localdevice.apis.Dialog.QuestionRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Dialog.QuestionResult>) responseObserver);
          break;
        case METHODID_MESSAGE_BOX:
          serviceImpl.messageBox((cloud.lazycat.localdevice.apis.Dialog.MessageBoxRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_PASSWORD:
          serviceImpl.password((cloud.lazycat.localdevice.apis.Dialog.PasswordRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.localdevice.apis.Dialog.PasswordResult>) responseObserver);
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

  private static abstract class DialogManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    DialogManagerBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.localdevice.apis.Dialog.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("DialogManager");
    }
  }

  private static final class DialogManagerFileDescriptorSupplier
      extends DialogManagerBaseDescriptorSupplier {
    DialogManagerFileDescriptorSupplier() {}
  }

  private static final class DialogManagerMethodDescriptorSupplier
      extends DialogManagerBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    DialogManagerMethodDescriptorSupplier(String methodName) {
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
      synchronized (DialogManagerGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DialogManagerFileDescriptorSupplier())
              .addMethod(getQuestionMethod())
              .addMethod(getMessageBoxMethod())
              .addMethod(getPasswordMethod())
              .build();
        }
      }
    }
    return result;
  }
}
