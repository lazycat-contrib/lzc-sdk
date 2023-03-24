package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/event.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class EventServiceGrpc {

  private EventServiceGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.EventService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest,
      cloud.lazycat.apis.sys.EventOuterClass.Event> getSubscribeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Subscribe",
      requestType = cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest.class,
      responseType = cloud.lazycat.apis.sys.EventOuterClass.Event.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest,
      cloud.lazycat.apis.sys.EventOuterClass.Event> getSubscribeMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest, cloud.lazycat.apis.sys.EventOuterClass.Event> getSubscribeMethod;
    if ((getSubscribeMethod = EventServiceGrpc.getSubscribeMethod) == null) {
      synchronized (EventServiceGrpc.class) {
        if ((getSubscribeMethod = EventServiceGrpc.getSubscribeMethod) == null) {
          EventServiceGrpc.getSubscribeMethod = getSubscribeMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest, cloud.lazycat.apis.sys.EventOuterClass.Event>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Subscribe"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.EventOuterClass.Event.getDefaultInstance()))
              .setSchemaDescriptor(new EventServiceMethodDescriptorSupplier("Subscribe"))
              .build();
        }
      }
    }
    return getSubscribeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SendRequest,
      com.google.protobuf.Empty> getSendMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Send",
      requestType = cloud.lazycat.apis.sys.EventOuterClass.SendRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SendRequest,
      com.google.protobuf.Empty> getSendMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SendRequest, com.google.protobuf.Empty> getSendMethod;
    if ((getSendMethod = EventServiceGrpc.getSendMethod) == null) {
      synchronized (EventServiceGrpc.class) {
        if ((getSendMethod = EventServiceGrpc.getSendMethod) == null) {
          EventServiceGrpc.getSendMethod = getSendMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.EventOuterClass.SendRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Send"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.EventOuterClass.SendRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new EventServiceMethodDescriptorSupplier("Send"))
              .build();
        }
      }
    }
    return getSendMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest,
      cloud.lazycat.apis.sys.EventOuterClass.Uuid> getGenPendingSendingMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GenPendingSending",
      requestType = cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest.class,
      responseType = cloud.lazycat.apis.sys.EventOuterClass.Uuid.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest,
      cloud.lazycat.apis.sys.EventOuterClass.Uuid> getGenPendingSendingMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest, cloud.lazycat.apis.sys.EventOuterClass.Uuid> getGenPendingSendingMethod;
    if ((getGenPendingSendingMethod = EventServiceGrpc.getGenPendingSendingMethod) == null) {
      synchronized (EventServiceGrpc.class) {
        if ((getGenPendingSendingMethod = EventServiceGrpc.getGenPendingSendingMethod) == null) {
          EventServiceGrpc.getGenPendingSendingMethod = getGenPendingSendingMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest, cloud.lazycat.apis.sys.EventOuterClass.Uuid>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GenPendingSending"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.EventOuterClass.Uuid.getDefaultInstance()))
              .setSchemaDescriptor(new EventServiceMethodDescriptorSupplier("GenPendingSending"))
              .build();
        }
      }
    }
    return getGenPendingSendingMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest,
      cloud.lazycat.apis.sys.EventOuterClass.Event> getSolvePendingMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SolvePending",
      requestType = cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest.class,
      responseType = cloud.lazycat.apis.sys.EventOuterClass.Event.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest,
      cloud.lazycat.apis.sys.EventOuterClass.Event> getSolvePendingMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest, cloud.lazycat.apis.sys.EventOuterClass.Event> getSolvePendingMethod;
    if ((getSolvePendingMethod = EventServiceGrpc.getSolvePendingMethod) == null) {
      synchronized (EventServiceGrpc.class) {
        if ((getSolvePendingMethod = EventServiceGrpc.getSolvePendingMethod) == null) {
          EventServiceGrpc.getSolvePendingMethod = getSolvePendingMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest, cloud.lazycat.apis.sys.EventOuterClass.Event>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SolvePending"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.EventOuterClass.Event.getDefaultInstance()))
              .setSchemaDescriptor(new EventServiceMethodDescriptorSupplier("SolvePending"))
              .build();
        }
      }
    }
    return getSolvePendingMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static EventServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EventServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EventServiceStub>() {
        @java.lang.Override
        public EventServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EventServiceStub(channel, callOptions);
        }
      };
    return EventServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static EventServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EventServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EventServiceBlockingStub>() {
        @java.lang.Override
        public EventServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EventServiceBlockingStub(channel, callOptions);
        }
      };
    return EventServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static EventServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EventServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EventServiceFutureStub>() {
        @java.lang.Override
        public EventServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EventServiceFutureStub(channel, callOptions);
        }
      };
    return EventServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class EventServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     *正常的事件订阅发送接口
     * </pre>
     */
    public void subscribe(cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Event> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubscribeMethod(), responseObserver);
    }

    /**
     */
    public void send(cloud.lazycat.apis.sys.EventOuterClass.SendRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSendMethod(), responseObserver);
    }

    /**
     * <pre>
     * 特殊的代理发送机制，少量特殊应用(box-settings在“找回密码"，"邀请用户")场景需要在未登陆盒子之前就行交互
     * 之前是出现一个操作就从lzcapp-&gt;lzc-runtime-&gt;hportal-&gt;hclient全部增加一层代理接口，维护性和灵活性都非常
     * 差。因此改成由特定lzcapp生产一个悬而未决的特殊PendingEvent. 此PendingEvent的内容主体由lzcapp创建，但
     * 少量内容由hclient根据用户交互等方式获取并提供给EventService进行合并。(比如uid/password字段)
     * 最后把一个PendingEvent转换为一个合法的Event后发送出去。
     * </pre>
     */
    public void genPendingSending(cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Uuid> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGenPendingSendingMethod(), responseObserver);
    }

    /**
     * <pre>
     * 这里再返回一个奇怪的Event是为了处理调用者无法感知事件产生的结果。
     * 比如用来由hclient代理创建用户时，会导致如果出现UID已经占用的情况无法传播给hclient
     * </pre>
     */
    public void solvePending(cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Event> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSolvePendingMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSubscribeMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest,
                cloud.lazycat.apis.sys.EventOuterClass.Event>(
                  this, METHODID_SUBSCRIBE)))
          .addMethod(
            getSendMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.EventOuterClass.SendRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_SEND)))
          .addMethod(
            getGenPendingSendingMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest,
                cloud.lazycat.apis.sys.EventOuterClass.Uuid>(
                  this, METHODID_GEN_PENDING_SENDING)))
          .addMethod(
            getSolvePendingMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest,
                cloud.lazycat.apis.sys.EventOuterClass.Event>(
                  this, METHODID_SOLVE_PENDING)))
          .build();
    }
  }

  /**
   */
  public static final class EventServiceStub extends io.grpc.stub.AbstractAsyncStub<EventServiceStub> {
    private EventServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EventServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EventServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     *正常的事件订阅发送接口
     * </pre>
     */
    public void subscribe(cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Event> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getSubscribeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void send(cloud.lazycat.apis.sys.EventOuterClass.SendRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSendMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 特殊的代理发送机制，少量特殊应用(box-settings在“找回密码"，"邀请用户")场景需要在未登陆盒子之前就行交互
     * 之前是出现一个操作就从lzcapp-&gt;lzc-runtime-&gt;hportal-&gt;hclient全部增加一层代理接口，维护性和灵活性都非常
     * 差。因此改成由特定lzcapp生产一个悬而未决的特殊PendingEvent. 此PendingEvent的内容主体由lzcapp创建，但
     * 少量内容由hclient根据用户交互等方式获取并提供给EventService进行合并。(比如uid/password字段)
     * 最后把一个PendingEvent转换为一个合法的Event后发送出去。
     * </pre>
     */
    public void genPendingSending(cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Uuid> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGenPendingSendingMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 这里再返回一个奇怪的Event是为了处理调用者无法感知事件产生的结果。
     * 比如用来由hclient代理创建用户时，会导致如果出现UID已经占用的情况无法传播给hclient
     * </pre>
     */
    public void solvePending(cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Event> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSolvePendingMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class EventServiceBlockingStub extends io.grpc.stub.AbstractBlockingStub<EventServiceBlockingStub> {
    private EventServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EventServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EventServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     *正常的事件订阅发送接口
     * </pre>
     */
    public java.util.Iterator<cloud.lazycat.apis.sys.EventOuterClass.Event> subscribe(
        cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getSubscribeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty send(cloud.lazycat.apis.sys.EventOuterClass.SendRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSendMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 特殊的代理发送机制，少量特殊应用(box-settings在“找回密码"，"邀请用户")场景需要在未登陆盒子之前就行交互
     * 之前是出现一个操作就从lzcapp-&gt;lzc-runtime-&gt;hportal-&gt;hclient全部增加一层代理接口，维护性和灵活性都非常
     * 差。因此改成由特定lzcapp生产一个悬而未决的特殊PendingEvent. 此PendingEvent的内容主体由lzcapp创建，但
     * 少量内容由hclient根据用户交互等方式获取并提供给EventService进行合并。(比如uid/password字段)
     * 最后把一个PendingEvent转换为一个合法的Event后发送出去。
     * </pre>
     */
    public cloud.lazycat.apis.sys.EventOuterClass.Uuid genPendingSending(cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGenPendingSendingMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 这里再返回一个奇怪的Event是为了处理调用者无法感知事件产生的结果。
     * 比如用来由hclient代理创建用户时，会导致如果出现UID已经占用的情况无法传播给hclient
     * </pre>
     */
    public cloud.lazycat.apis.sys.EventOuterClass.Event solvePending(cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSolvePendingMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class EventServiceFutureStub extends io.grpc.stub.AbstractFutureStub<EventServiceFutureStub> {
    private EventServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EventServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EventServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> send(
        cloud.lazycat.apis.sys.EventOuterClass.SendRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSendMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 特殊的代理发送机制，少量特殊应用(box-settings在“找回密码"，"邀请用户")场景需要在未登陆盒子之前就行交互
     * 之前是出现一个操作就从lzcapp-&gt;lzc-runtime-&gt;hportal-&gt;hclient全部增加一层代理接口，维护性和灵活性都非常
     * 差。因此改成由特定lzcapp生产一个悬而未决的特殊PendingEvent. 此PendingEvent的内容主体由lzcapp创建，但
     * 少量内容由hclient根据用户交互等方式获取并提供给EventService进行合并。(比如uid/password字段)
     * 最后把一个PendingEvent转换为一个合法的Event后发送出去。
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.EventOuterClass.Uuid> genPendingSending(
        cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGenPendingSendingMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 这里再返回一个奇怪的Event是为了处理调用者无法感知事件产生的结果。
     * 比如用来由hclient代理创建用户时，会导致如果出现UID已经占用的情况无法传播给hclient
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.EventOuterClass.Event> solvePending(
        cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSolvePendingMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SUBSCRIBE = 0;
  private static final int METHODID_SEND = 1;
  private static final int METHODID_GEN_PENDING_SENDING = 2;
  private static final int METHODID_SOLVE_PENDING = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final EventServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(EventServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SUBSCRIBE:
          serviceImpl.subscribe((cloud.lazycat.apis.sys.EventOuterClass.SubscribeRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Event>) responseObserver);
          break;
        case METHODID_SEND:
          serviceImpl.send((cloud.lazycat.apis.sys.EventOuterClass.SendRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GEN_PENDING_SENDING:
          serviceImpl.genPendingSending((cloud.lazycat.apis.sys.EventOuterClass.GenPendingSendingRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Uuid>) responseObserver);
          break;
        case METHODID_SOLVE_PENDING:
          serviceImpl.solvePending((cloud.lazycat.apis.sys.EventOuterClass.SolvePendingRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.EventOuterClass.Event>) responseObserver);
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

  private static abstract class EventServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    EventServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.EventOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("EventService");
    }
  }

  private static final class EventServiceFileDescriptorSupplier
      extends EventServiceBaseDescriptorSupplier {
    EventServiceFileDescriptorSupplier() {}
  }

  private static final class EventServiceMethodDescriptorSupplier
      extends EventServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    EventServiceMethodDescriptorSupplier(String methodName) {
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
      synchronized (EventServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new EventServiceFileDescriptorSupplier())
              .addMethod(getSubscribeMethod())
              .addMethod(getSendMethod())
              .addMethod(getGenPendingSendingMethod())
              .addMethod(getSolvePendingMethod())
              .build();
        }
      }
    }
    return result;
  }
}
