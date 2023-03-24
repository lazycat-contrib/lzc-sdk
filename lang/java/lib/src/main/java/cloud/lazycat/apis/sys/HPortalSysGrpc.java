package cloud.lazycat.apis.sys;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.47.0)",
    comments = "Source: sys/portal-server/portal-server.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class HPortalSysGrpc {

  private HPortalSysGrpc() {}

  public static final String SERVICE_NAME = "cloud.lazycat.apis.sys.HPortalSys";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AuthToken,
      cloud.lazycat.apis.sys.PortalServer.LoginInfo> getQueryLoginMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryLogin",
      requestType = cloud.lazycat.apis.sys.PortalServer.AuthToken.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.LoginInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AuthToken,
      cloud.lazycat.apis.sys.PortalServer.LoginInfo> getQueryLoginMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AuthToken, cloud.lazycat.apis.sys.PortalServer.LoginInfo> getQueryLoginMethod;
    if ((getQueryLoginMethod = HPortalSysGrpc.getQueryLoginMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getQueryLoginMethod = HPortalSysGrpc.getQueryLoginMethod) == null) {
          HPortalSysGrpc.getQueryLoginMethod = getQueryLoginMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.AuthToken, cloud.lazycat.apis.sys.PortalServer.LoginInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryLogin"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.AuthToken.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.LoginInfo.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("QueryLogin"))
              .build();
        }
      }
    }
    return getQueryLoginMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest,
      cloud.lazycat.apis.sys.PortalServer.ListDeviceReply> getListDevicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListDevices",
      requestType = cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.ListDeviceReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest,
      cloud.lazycat.apis.sys.PortalServer.ListDeviceReply> getListDevicesMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest, cloud.lazycat.apis.sys.PortalServer.ListDeviceReply> getListDevicesMethod;
    if ((getListDevicesMethod = HPortalSysGrpc.getListDevicesMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getListDevicesMethod = HPortalSysGrpc.getListDevicesMethod) == null) {
          HPortalSysGrpc.getListDevicesMethod = getListDevicesMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest, cloud.lazycat.apis.sys.PortalServer.ListDeviceReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListDevices"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.ListDeviceReply.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("ListDevices"))
              .build();
        }
      }
    }
    return getListDevicesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DeviceID,
      cloud.lazycat.apis.sys.PortalServer.Device> getQueryDeviceByIDMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryDeviceByID",
      requestType = cloud.lazycat.apis.sys.PortalServer.DeviceID.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.Device.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DeviceID,
      cloud.lazycat.apis.sys.PortalServer.Device> getQueryDeviceByIDMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DeviceID, cloud.lazycat.apis.sys.PortalServer.Device> getQueryDeviceByIDMethod;
    if ((getQueryDeviceByIDMethod = HPortalSysGrpc.getQueryDeviceByIDMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getQueryDeviceByIDMethod = HPortalSysGrpc.getQueryDeviceByIDMethod) == null) {
          HPortalSysGrpc.getQueryDeviceByIDMethod = getQueryDeviceByIDMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.DeviceID, cloud.lazycat.apis.sys.PortalServer.Device>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryDeviceByID"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.DeviceID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.Device.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("QueryDeviceByID"))
              .build();
        }
      }
    }
    return getQueryDeviceByIDMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.PortalServer.BoxInfo> getQueryBoxInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryBoxInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.BoxInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.PortalServer.BoxInfo> getQueryBoxInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.PortalServer.BoxInfo> getQueryBoxInfoMethod;
    if ((getQueryBoxInfoMethod = HPortalSysGrpc.getQueryBoxInfoMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getQueryBoxInfoMethod = HPortalSysGrpc.getQueryBoxInfoMethod) == null) {
          HPortalSysGrpc.getQueryBoxInfoMethod = getQueryBoxInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.PortalServer.BoxInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryBoxInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.BoxInfo.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("QueryBoxInfo"))
              .build();
        }
      }
    }
    return getQueryBoxInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DomainCertRequest,
      cloud.lazycat.apis.sys.PortalServer.DomainCertReply> getGetDomainCertMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDomainCert",
      requestType = cloud.lazycat.apis.sys.PortalServer.DomainCertRequest.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.DomainCertReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DomainCertRequest,
      cloud.lazycat.apis.sys.PortalServer.DomainCertReply> getGetDomainCertMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DomainCertRequest, cloud.lazycat.apis.sys.PortalServer.DomainCertReply> getGetDomainCertMethod;
    if ((getGetDomainCertMethod = HPortalSysGrpc.getGetDomainCertMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getGetDomainCertMethod = HPortalSysGrpc.getGetDomainCertMethod) == null) {
          HPortalSysGrpc.getGetDomainCertMethod = getGetDomainCertMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.DomainCertRequest, cloud.lazycat.apis.sys.PortalServer.DomainCertReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDomainCert"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.DomainCertRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.DomainCertReply.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("GetDomainCert"))
              .build();
        }
      }
    }
    return getGetDomainCertMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DomainCertRequest,
      cloud.lazycat.apis.sys.PortalServer.DomainCertReply> getGetDomainSelfCertMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDomainSelfCert",
      requestType = cloud.lazycat.apis.sys.PortalServer.DomainCertRequest.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.DomainCertReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DomainCertRequest,
      cloud.lazycat.apis.sys.PortalServer.DomainCertReply> getGetDomainSelfCertMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DomainCertRequest, cloud.lazycat.apis.sys.PortalServer.DomainCertReply> getGetDomainSelfCertMethod;
    if ((getGetDomainSelfCertMethod = HPortalSysGrpc.getGetDomainSelfCertMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getGetDomainSelfCertMethod = HPortalSysGrpc.getGetDomainSelfCertMethod) == null) {
          HPortalSysGrpc.getGetDomainSelfCertMethod = getGetDomainSelfCertMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.DomainCertRequest, cloud.lazycat.apis.sys.PortalServer.DomainCertReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDomainSelfCert"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.DomainCertRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.DomainCertReply.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("GetDomainSelfCert"))
              .build();
        }
      }
    }
    return getGetDomainSelfCertMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AppCertRequest,
      cloud.lazycat.apis.sys.PortalServer.AppCertReply> getGetAppCertMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAppCert",
      requestType = cloud.lazycat.apis.sys.PortalServer.AppCertRequest.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.AppCertReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AppCertRequest,
      cloud.lazycat.apis.sys.PortalServer.AppCertReply> getGetAppCertMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AppCertRequest, cloud.lazycat.apis.sys.PortalServer.AppCertReply> getGetAppCertMethod;
    if ((getGetAppCertMethod = HPortalSysGrpc.getGetAppCertMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getGetAppCertMethod = HPortalSysGrpc.getGetAppCertMethod) == null) {
          HPortalSysGrpc.getGetAppCertMethod = getGetAppCertMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.AppCertRequest, cloud.lazycat.apis.sys.PortalServer.AppCertReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAppCert"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.AppCertRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.AppCertReply.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("GetAppCert"))
              .build();
        }
      }
    }
    return getGetAppCertMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest,
      cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply> getAllocVirtualExternalIPMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AllocVirtualExternalIP",
      requestType = cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest,
      cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply> getAllocVirtualExternalIPMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest, cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply> getAllocVirtualExternalIPMethod;
    if ((getAllocVirtualExternalIPMethod = HPortalSysGrpc.getAllocVirtualExternalIPMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getAllocVirtualExternalIPMethod = HPortalSysGrpc.getAllocVirtualExternalIPMethod) == null) {
          HPortalSysGrpc.getAllocVirtualExternalIPMethod = getAllocVirtualExternalIPMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest, cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AllocVirtualExternalIP"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("AllocVirtualExternalIP"))
              .build();
        }
      }
    }
    return getAllocVirtualExternalIPMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest,
      cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply> getFreeVirtualExternalIPMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FreeVirtualExternalIP",
      requestType = cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest,
      cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply> getFreeVirtualExternalIPMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest, cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply> getFreeVirtualExternalIPMethod;
    if ((getFreeVirtualExternalIPMethod = HPortalSysGrpc.getFreeVirtualExternalIPMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getFreeVirtualExternalIPMethod = HPortalSysGrpc.getFreeVirtualExternalIPMethod) == null) {
          HPortalSysGrpc.getFreeVirtualExternalIPMethod = getFreeVirtualExternalIPMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest, cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FreeVirtualExternalIP"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("FreeVirtualExternalIP"))
              .build();
        }
      }
    }
    return getFreeVirtualExternalIPMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest,
      com.google.protobuf.Empty> getPairDevicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PairDevices",
      requestType = cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest,
      com.google.protobuf.Empty> getPairDevicesMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest, com.google.protobuf.Empty> getPairDevicesMethod;
    if ((getPairDevicesMethod = HPortalSysGrpc.getPairDevicesMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getPairDevicesMethod = HPortalSysGrpc.getPairDevicesMethod) == null) {
          HPortalSysGrpc.getPairDevicesMethod = getPairDevicesMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PairDevices"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("PairDevices"))
              .build();
        }
      }
    }
    return getPairDevicesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.PortalServer.ListUsersReply> getListUsersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListUsers",
      requestType = com.google.protobuf.Empty.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.ListUsersReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      cloud.lazycat.apis.sys.PortalServer.ListUsersReply> getListUsersMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, cloud.lazycat.apis.sys.PortalServer.ListUsersReply> getListUsersMethod;
    if ((getListUsersMethod = HPortalSysGrpc.getListUsersMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getListUsersMethod = HPortalSysGrpc.getListUsersMethod) == null) {
          HPortalSysGrpc.getListUsersMethod = getListUsersMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, cloud.lazycat.apis.sys.PortalServer.ListUsersReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListUsers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.ListUsersReply.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("ListUsers"))
              .build();
        }
      }
    }
    return getListUsersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.UserID,
      cloud.lazycat.apis.sys.PortalServer.QueryRoleReply> getQueryRoleMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "QueryRole",
      requestType = cloud.lazycat.apis.sys.PortalServer.UserID.class,
      responseType = cloud.lazycat.apis.sys.PortalServer.QueryRoleReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.UserID,
      cloud.lazycat.apis.sys.PortalServer.QueryRoleReply> getQueryRoleMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.UserID, cloud.lazycat.apis.sys.PortalServer.QueryRoleReply> getQueryRoleMethod;
    if ((getQueryRoleMethod = HPortalSysGrpc.getQueryRoleMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getQueryRoleMethod = HPortalSysGrpc.getQueryRoleMethod) == null) {
          HPortalSysGrpc.getQueryRoleMethod = getQueryRoleMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.UserID, cloud.lazycat.apis.sys.PortalServer.QueryRoleReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "QueryRole"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.UserID.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.QueryRoleReply.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("QueryRole"))
              .build();
        }
      }
    }
    return getQueryRoleMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust,
      com.google.protobuf.Empty> getChangeRoleMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ChangeRole",
      requestType = cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust,
      com.google.protobuf.Empty> getChangeRoleMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust, com.google.protobuf.Empty> getChangeRoleMethod;
    if ((getChangeRoleMethod = HPortalSysGrpc.getChangeRoleMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getChangeRoleMethod = HPortalSysGrpc.getChangeRoleMethod) == null) {
          HPortalSysGrpc.getChangeRoleMethod = getChangeRoleMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ChangeRole"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("ChangeRole"))
              .build();
        }
      }
    }
    return getChangeRoleMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest,
      com.google.protobuf.Empty> getResetPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ResetPassword",
      requestType = cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest,
      com.google.protobuf.Empty> getResetPasswordMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest, com.google.protobuf.Empty> getResetPasswordMethod;
    if ((getResetPasswordMethod = HPortalSysGrpc.getResetPasswordMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getResetPasswordMethod = HPortalSysGrpc.getResetPasswordMethod) == null) {
          HPortalSysGrpc.getResetPasswordMethod = getResetPasswordMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ResetPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("ResetPassword"))
              .build();
        }
      }
    }
    return getResetPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest,
      com.google.protobuf.Empty> getDeleteUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteUser",
      requestType = cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest,
      com.google.protobuf.Empty> getDeleteUserMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest, com.google.protobuf.Empty> getDeleteUserMethod;
    if ((getDeleteUserMethod = HPortalSysGrpc.getDeleteUserMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getDeleteUserMethod = HPortalSysGrpc.getDeleteUserMethod) == null) {
          HPortalSysGrpc.getDeleteUserMethod = getDeleteUserMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("DeleteUser"))
              .build();
        }
      }
    }
    return getDeleteUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.CreateUserRequest,
      com.google.protobuf.Empty> getCreateUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateUser",
      requestType = cloud.lazycat.apis.sys.PortalServer.CreateUserRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.CreateUserRequest,
      com.google.protobuf.Empty> getCreateUserMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.CreateUserRequest, com.google.protobuf.Empty> getCreateUserMethod;
    if ((getCreateUserMethod = HPortalSysGrpc.getCreateUserMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getCreateUserMethod = HPortalSysGrpc.getCreateUserMethod) == null) {
          HPortalSysGrpc.getCreateUserMethod = getCreateUserMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.CreateUserRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.CreateUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("CreateUser"))
              .build();
        }
      }
    }
    return getCreateUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest,
      com.google.protobuf.Empty> getForceResetPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ForceResetPassword",
      requestType = cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest,
      com.google.protobuf.Empty> getForceResetPasswordMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest, com.google.protobuf.Empty> getForceResetPasswordMethod;
    if ((getForceResetPasswordMethod = HPortalSysGrpc.getForceResetPasswordMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getForceResetPasswordMethod = HPortalSysGrpc.getForceResetPasswordMethod) == null) {
          HPortalSysGrpc.getForceResetPasswordMethod = getForceResetPasswordMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ForceResetPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("ForceResetPassword"))
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
    if ((getGenUserInvitationMethod = HPortalSysGrpc.getGenUserInvitationMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getGenUserInvitationMethod = HPortalSysGrpc.getGenUserInvitationMethod) == null) {
          HPortalSysGrpc.getGenUserInvitationMethod = getGenUserInvitationMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest, cloud.lazycat.apis.sys.PortalServer.UserInvitation>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GenUserInvitation"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.UserInvitation.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("GenUserInvitation"))
              .build();
        }
      }
    }
    return getGenUserInvitationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.LogoutRequest,
      com.google.protobuf.Empty> getLogoutMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Logout",
      requestType = cloud.lazycat.apis.sys.PortalServer.LogoutRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.LogoutRequest,
      com.google.protobuf.Empty> getLogoutMethod() {
    io.grpc.MethodDescriptor<cloud.lazycat.apis.sys.PortalServer.LogoutRequest, com.google.protobuf.Empty> getLogoutMethod;
    if ((getLogoutMethod = HPortalSysGrpc.getLogoutMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getLogoutMethod = HPortalSysGrpc.getLogoutMethod) == null) {
          HPortalSysGrpc.getLogoutMethod = getLogoutMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.LogoutRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Logout"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.LogoutRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("Logout"))
              .build();
        }
      }
    }
    return getLogoutMethod;
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
    if ((getCheckPasswordMethod = HPortalSysGrpc.getCheckPasswordMethod) == null) {
      synchronized (HPortalSysGrpc.class) {
        if ((getCheckPasswordMethod = HPortalSysGrpc.getCheckPasswordMethod) == null) {
          HPortalSysGrpc.getCheckPasswordMethod = getCheckPasswordMethod =
              io.grpc.MethodDescriptor.<cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CheckPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new HPortalSysMethodDescriptorSupplier("CheckPassword"))
              .build();
        }
      }
    }
    return getCheckPasswordMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static HPortalSysStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<HPortalSysStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<HPortalSysStub>() {
        @java.lang.Override
        public HPortalSysStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new HPortalSysStub(channel, callOptions);
        }
      };
    return HPortalSysStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static HPortalSysBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<HPortalSysBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<HPortalSysBlockingStub>() {
        @java.lang.Override
        public HPortalSysBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new HPortalSysBlockingStub(channel, callOptions);
        }
      };
    return HPortalSysBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static HPortalSysFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<HPortalSysFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<HPortalSysFutureStub>() {
        @java.lang.Override
        public HPortalSysFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new HPortalSysFutureStub(channel, callOptions);
        }
      };
    return HPortalSysFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class HPortalSysImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * 用auth-token反向查询登陆的设备以及UID
     * </pre>
     */
    public void queryLogin(cloud.lazycat.apis.sys.PortalServer.AuthToken request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.LoginInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryLoginMethod(), responseObserver);
    }

    /**
     * <pre>
     * 根据UID返回所有的设备列表
     * </pre>
     */
    public void listDevices(cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.ListDeviceReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListDevicesMethod(), responseObserver);
    }

    /**
     */
    public void queryDeviceByID(cloud.lazycat.apis.sys.PortalServer.DeviceID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.Device> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryDeviceByIDMethod(), responseObserver);
    }

    /**
     */
    public void queryBoxInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.BoxInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryBoxInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * 获取盒子所属域名下或下一级域名的https证书。
     * 注意不是所有ACME服务器都支持泛域名。
     * </pre>
     */
    public void getDomainCert(cloud.lazycat.apis.sys.PortalServer.DomainCertRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.DomainCertReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDomainCertMethod(), responseObserver);
    }

    /**
     */
    public void getDomainSelfCert(cloud.lazycat.apis.sys.PortalServer.DomainCertRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.DomainCertReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDomainSelfCertMethod(), responseObserver);
    }

    /**
     * <pre>
     * 在部署具体app前，调用此接口获取app证书
     * APP证书格式为:
     *   Issuer: O = $BOX.ORIGIN, CN = $BOX.DOMAIN, serialNumber = $BOX.ID
     *   Subject: O = $BOX.ORIGIN, CN = $APP.DOMAIN, serialNumber = '$UID&#64;$APP.ID || $APP.ID'
     * Issuer为box.crt，通过QueryBoxInfo查询到BoxInfo.BoxCrt。并且box.crt的公钥与box.id是一一对应关系。
     * 盒子内部组件可以直接通过QueryBoxInfo来验证信任链是否合法，盒子外部系统需要通过其他机制比如libp2p.identify去验证box.crt的合法性。
     * </pre>
     */
    public void getAppCert(cloud.lazycat.apis.sys.PortalServer.AppCertRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.AppCertReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAppCertMethod(), responseObserver);
    }

    /**
     * <pre>
     * 申请额外的外部可访问的IP,并配置对应访问域名
     * </pre>
     */
    public void allocVirtualExternalIP(cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAllocVirtualExternalIPMethod(), responseObserver);
    }

    /**
     * <pre>
     * 释放虚拟IP
     * </pre>
     */
    public void freeVirtualExternalIP(cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFreeVirtualExternalIPMethod(), responseObserver);
    }

    /**
     */
    public void pairDevices(cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPairDevicesMethod(), responseObserver);
    }

    /**
     * <pre>
     *  查询所有UID
     * </pre>
     */
    public void listUsers(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.ListUsersReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListUsersMethod(), responseObserver);
    }

    /**
     * <pre>
     *  根据用户uid查询用户信息
     * </pre>
     */
    public void queryRole(cloud.lazycat.apis.sys.PortalServer.UserID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.QueryRoleReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getQueryRoleMethod(), responseObserver);
    }

    /**
     * <pre>
     *  修改指定uid的用户角色
     * </pre>
     */
    public void changeRole(cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getChangeRoleMethod(), responseObserver);
    }

    /**
     * <pre>
     *  通过验证旧密码修改新的密码
     * </pre>
     */
    public void resetPassword(cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getResetPasswordMethod(), responseObserver);
    }

    /**
     * <pre>
     *  删除用户信息
     * </pre>
     */
    public void deleteUser(cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteUserMethod(), responseObserver);
    }

    /**
     * <pre>
     *  创建用户信息
     * </pre>
     */
    public void createUser(cloud.lazycat.apis.sys.PortalServer.CreateUserRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateUserMethod(), responseObserver);
    }

    /**
     * <pre>
     *  强制重置用户密码
     * </pre>
     */
    public void forceResetPassword(cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getForceResetPasswordMethod(), responseObserver);
    }

    /**
     * <pre>
     * 生成用户注册token,以便上层实现各类用户注册机制
     * </pre>
     */
    public void genUserInvitation(cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.UserInvitation> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGenUserInvitationMethod(), responseObserver);
    }

    /**
     * <pre>
     * 注销当前用户指定设备，同时将连接断开
     * </pre>
     */
    public void logout(cloud.lazycat.apis.sys.PortalServer.LogoutRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLogoutMethod(), responseObserver);
    }

    /**
     * <pre>
     *校验用户密码是否正确
     * </pre>
     */
    public void checkPassword(cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCheckPasswordMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getQueryLoginMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.AuthToken,
                cloud.lazycat.apis.sys.PortalServer.LoginInfo>(
                  this, METHODID_QUERY_LOGIN)))
          .addMethod(
            getListDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest,
                cloud.lazycat.apis.sys.PortalServer.ListDeviceReply>(
                  this, METHODID_LIST_DEVICES)))
          .addMethod(
            getQueryDeviceByIDMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.DeviceID,
                cloud.lazycat.apis.sys.PortalServer.Device>(
                  this, METHODID_QUERY_DEVICE_BY_ID)))
          .addMethod(
            getQueryBoxInfoMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.PortalServer.BoxInfo>(
                  this, METHODID_QUERY_BOX_INFO)))
          .addMethod(
            getGetDomainCertMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.DomainCertRequest,
                cloud.lazycat.apis.sys.PortalServer.DomainCertReply>(
                  this, METHODID_GET_DOMAIN_CERT)))
          .addMethod(
            getGetDomainSelfCertMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.DomainCertRequest,
                cloud.lazycat.apis.sys.PortalServer.DomainCertReply>(
                  this, METHODID_GET_DOMAIN_SELF_CERT)))
          .addMethod(
            getGetAppCertMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.AppCertRequest,
                cloud.lazycat.apis.sys.PortalServer.AppCertReply>(
                  this, METHODID_GET_APP_CERT)))
          .addMethod(
            getAllocVirtualExternalIPMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest,
                cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply>(
                  this, METHODID_ALLOC_VIRTUAL_EXTERNAL_IP)))
          .addMethod(
            getFreeVirtualExternalIPMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest,
                cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply>(
                  this, METHODID_FREE_VIRTUAL_EXTERNAL_IP)))
          .addMethod(
            getPairDevicesMethod(),
            io.grpc.stub.ServerCalls.asyncServerStreamingCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_PAIR_DEVICES)))
          .addMethod(
            getListUsersMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.google.protobuf.Empty,
                cloud.lazycat.apis.sys.PortalServer.ListUsersReply>(
                  this, METHODID_LIST_USERS)))
          .addMethod(
            getQueryRoleMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.UserID,
                cloud.lazycat.apis.sys.PortalServer.QueryRoleReply>(
                  this, METHODID_QUERY_ROLE)))
          .addMethod(
            getChangeRoleMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust,
                com.google.protobuf.Empty>(
                  this, METHODID_CHANGE_ROLE)))
          .addMethod(
            getResetPasswordMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_RESET_PASSWORD)))
          .addMethod(
            getDeleteUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_DELETE_USER)))
          .addMethod(
            getCreateUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.CreateUserRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_CREATE_USER)))
          .addMethod(
            getForceResetPasswordMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest,
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
            getLogoutMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                cloud.lazycat.apis.sys.PortalServer.LogoutRequest,
                com.google.protobuf.Empty>(
                  this, METHODID_LOGOUT)))
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
  public static final class HPortalSysStub extends io.grpc.stub.AbstractAsyncStub<HPortalSysStub> {
    private HPortalSysStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected HPortalSysStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new HPortalSysStub(channel, callOptions);
    }

    /**
     * <pre>
     * 用auth-token反向查询登陆的设备以及UID
     * </pre>
     */
    public void queryLogin(cloud.lazycat.apis.sys.PortalServer.AuthToken request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.LoginInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryLoginMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 根据UID返回所有的设备列表
     * </pre>
     */
    public void listDevices(cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.ListDeviceReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListDevicesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void queryDeviceByID(cloud.lazycat.apis.sys.PortalServer.DeviceID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.Device> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryDeviceByIDMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void queryBoxInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.BoxInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryBoxInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 获取盒子所属域名下或下一级域名的https证书。
     * 注意不是所有ACME服务器都支持泛域名。
     * </pre>
     */
    public void getDomainCert(cloud.lazycat.apis.sys.PortalServer.DomainCertRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.DomainCertReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDomainCertMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getDomainSelfCert(cloud.lazycat.apis.sys.PortalServer.DomainCertRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.DomainCertReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDomainSelfCertMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 在部署具体app前，调用此接口获取app证书
     * APP证书格式为:
     *   Issuer: O = $BOX.ORIGIN, CN = $BOX.DOMAIN, serialNumber = $BOX.ID
     *   Subject: O = $BOX.ORIGIN, CN = $APP.DOMAIN, serialNumber = '$UID&#64;$APP.ID || $APP.ID'
     * Issuer为box.crt，通过QueryBoxInfo查询到BoxInfo.BoxCrt。并且box.crt的公钥与box.id是一一对应关系。
     * 盒子内部组件可以直接通过QueryBoxInfo来验证信任链是否合法，盒子外部系统需要通过其他机制比如libp2p.identify去验证box.crt的合法性。
     * </pre>
     */
    public void getAppCert(cloud.lazycat.apis.sys.PortalServer.AppCertRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.AppCertReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAppCertMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 申请额外的外部可访问的IP,并配置对应访问域名
     * </pre>
     */
    public void allocVirtualExternalIP(cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAllocVirtualExternalIPMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 释放虚拟IP
     * </pre>
     */
    public void freeVirtualExternalIP(cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFreeVirtualExternalIPMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void pairDevices(cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getPairDevicesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  查询所有UID
     * </pre>
     */
    public void listUsers(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.ListUsersReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListUsersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  根据用户uid查询用户信息
     * </pre>
     */
    public void queryRole(cloud.lazycat.apis.sys.PortalServer.UserID request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.QueryRoleReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getQueryRoleMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  修改指定uid的用户角色
     * </pre>
     */
    public void changeRole(cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getChangeRoleMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  通过验证旧密码修改新的密码
     * </pre>
     */
    public void resetPassword(cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getResetPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  删除用户信息
     * </pre>
     */
    public void deleteUser(cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  创建用户信息
     * </pre>
     */
    public void createUser(cloud.lazycat.apis.sys.PortalServer.CreateUserRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *  强制重置用户密码
     * </pre>
     */
    public void forceResetPassword(cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getForceResetPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 生成用户注册token,以便上层实现各类用户注册机制
     * </pre>
     */
    public void genUserInvitation(cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest request,
        io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.UserInvitation> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGenUserInvitationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * 注销当前用户指定设备，同时将连接断开
     * </pre>
     */
    public void logout(cloud.lazycat.apis.sys.PortalServer.LogoutRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLogoutMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     *校验用户密码是否正确
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
  public static final class HPortalSysBlockingStub extends io.grpc.stub.AbstractBlockingStub<HPortalSysBlockingStub> {
    private HPortalSysBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected HPortalSysBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new HPortalSysBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * 用auth-token反向查询登陆的设备以及UID
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.LoginInfo queryLogin(cloud.lazycat.apis.sys.PortalServer.AuthToken request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryLoginMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 根据UID返回所有的设备列表
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.ListDeviceReply listDevices(cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListDevicesMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.sys.PortalServer.Device queryDeviceByID(cloud.lazycat.apis.sys.PortalServer.DeviceID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryDeviceByIDMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.sys.PortalServer.BoxInfo queryBoxInfo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryBoxInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 获取盒子所属域名下或下一级域名的https证书。
     * 注意不是所有ACME服务器都支持泛域名。
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.DomainCertReply getDomainCert(cloud.lazycat.apis.sys.PortalServer.DomainCertRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDomainCertMethod(), getCallOptions(), request);
    }

    /**
     */
    public cloud.lazycat.apis.sys.PortalServer.DomainCertReply getDomainSelfCert(cloud.lazycat.apis.sys.PortalServer.DomainCertRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDomainSelfCertMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 在部署具体app前，调用此接口获取app证书
     * APP证书格式为:
     *   Issuer: O = $BOX.ORIGIN, CN = $BOX.DOMAIN, serialNumber = $BOX.ID
     *   Subject: O = $BOX.ORIGIN, CN = $APP.DOMAIN, serialNumber = '$UID&#64;$APP.ID || $APP.ID'
     * Issuer为box.crt，通过QueryBoxInfo查询到BoxInfo.BoxCrt。并且box.crt的公钥与box.id是一一对应关系。
     * 盒子内部组件可以直接通过QueryBoxInfo来验证信任链是否合法，盒子外部系统需要通过其他机制比如libp2p.identify去验证box.crt的合法性。
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.AppCertReply getAppCert(cloud.lazycat.apis.sys.PortalServer.AppCertRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAppCertMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 申请额外的外部可访问的IP,并配置对应访问域名
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply allocVirtualExternalIP(cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAllocVirtualExternalIPMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 释放虚拟IP
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply freeVirtualExternalIP(cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFreeVirtualExternalIPMethod(), getCallOptions(), request);
    }

    /**
     */
    public java.util.Iterator<com.google.protobuf.Empty> pairDevices(
        cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getPairDevicesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  查询所有UID
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.ListUsersReply listUsers(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListUsersMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  根据用户uid查询用户信息
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.QueryRoleReply queryRole(cloud.lazycat.apis.sys.PortalServer.UserID request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getQueryRoleMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  修改指定uid的用户角色
     * </pre>
     */
    public com.google.protobuf.Empty changeRole(cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getChangeRoleMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  通过验证旧密码修改新的密码
     * </pre>
     */
    public com.google.protobuf.Empty resetPassword(cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getResetPasswordMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  删除用户信息
     * </pre>
     */
    public com.google.protobuf.Empty deleteUser(cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteUserMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  创建用户信息
     * </pre>
     */
    public com.google.protobuf.Empty createUser(cloud.lazycat.apis.sys.PortalServer.CreateUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateUserMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *  强制重置用户密码
     * </pre>
     */
    public com.google.protobuf.Empty forceResetPassword(cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getForceResetPasswordMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 生成用户注册token,以便上层实现各类用户注册机制
     * </pre>
     */
    public cloud.lazycat.apis.sys.PortalServer.UserInvitation genUserInvitation(cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGenUserInvitationMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * 注销当前用户指定设备，同时将连接断开
     * </pre>
     */
    public com.google.protobuf.Empty logout(cloud.lazycat.apis.sys.PortalServer.LogoutRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLogoutMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     *校验用户密码是否正确
     * </pre>
     */
    public com.google.protobuf.Empty checkPassword(cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCheckPasswordMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class HPortalSysFutureStub extends io.grpc.stub.AbstractFutureStub<HPortalSysFutureStub> {
    private HPortalSysFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected HPortalSysFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new HPortalSysFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * 用auth-token反向查询登陆的设备以及UID
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.LoginInfo> queryLogin(
        cloud.lazycat.apis.sys.PortalServer.AuthToken request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryLoginMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 根据UID返回所有的设备列表
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.ListDeviceReply> listDevices(
        cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListDevicesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.Device> queryDeviceByID(
        cloud.lazycat.apis.sys.PortalServer.DeviceID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryDeviceByIDMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.BoxInfo> queryBoxInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryBoxInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 获取盒子所属域名下或下一级域名的https证书。
     * 注意不是所有ACME服务器都支持泛域名。
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.DomainCertReply> getDomainCert(
        cloud.lazycat.apis.sys.PortalServer.DomainCertRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDomainCertMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.DomainCertReply> getDomainSelfCert(
        cloud.lazycat.apis.sys.PortalServer.DomainCertRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDomainSelfCertMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 在部署具体app前，调用此接口获取app证书
     * APP证书格式为:
     *   Issuer: O = $BOX.ORIGIN, CN = $BOX.DOMAIN, serialNumber = $BOX.ID
     *   Subject: O = $BOX.ORIGIN, CN = $APP.DOMAIN, serialNumber = '$UID&#64;$APP.ID || $APP.ID'
     * Issuer为box.crt，通过QueryBoxInfo查询到BoxInfo.BoxCrt。并且box.crt的公钥与box.id是一一对应关系。
     * 盒子内部组件可以直接通过QueryBoxInfo来验证信任链是否合法，盒子外部系统需要通过其他机制比如libp2p.identify去验证box.crt的合法性。
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.AppCertReply> getAppCert(
        cloud.lazycat.apis.sys.PortalServer.AppCertRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAppCertMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 申请额外的外部可访问的IP,并配置对应访问域名
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply> allocVirtualExternalIP(
        cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAllocVirtualExternalIPMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 释放虚拟IP
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply> freeVirtualExternalIP(
        cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFreeVirtualExternalIPMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  查询所有UID
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.ListUsersReply> listUsers(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListUsersMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  根据用户uid查询用户信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.QueryRoleReply> queryRole(
        cloud.lazycat.apis.sys.PortalServer.UserID request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getQueryRoleMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  修改指定uid的用户角色
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> changeRole(
        cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getChangeRoleMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  通过验证旧密码修改新的密码
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> resetPassword(
        cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getResetPasswordMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  删除用户信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> deleteUser(
        cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteUserMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  创建用户信息
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> createUser(
        cloud.lazycat.apis.sys.PortalServer.CreateUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *  强制重置用户密码
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> forceResetPassword(
        cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getForceResetPasswordMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 生成用户注册token,以便上层实现各类用户注册机制
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<cloud.lazycat.apis.sys.PortalServer.UserInvitation> genUserInvitation(
        cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGenUserInvitationMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * 注销当前用户指定设备，同时将连接断开
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> logout(
        cloud.lazycat.apis.sys.PortalServer.LogoutRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLogoutMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     *校验用户密码是否正确
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> checkPassword(
        cloud.lazycat.apis.sys.PortalServer.CheckPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCheckPasswordMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_QUERY_LOGIN = 0;
  private static final int METHODID_LIST_DEVICES = 1;
  private static final int METHODID_QUERY_DEVICE_BY_ID = 2;
  private static final int METHODID_QUERY_BOX_INFO = 3;
  private static final int METHODID_GET_DOMAIN_CERT = 4;
  private static final int METHODID_GET_DOMAIN_SELF_CERT = 5;
  private static final int METHODID_GET_APP_CERT = 6;
  private static final int METHODID_ALLOC_VIRTUAL_EXTERNAL_IP = 7;
  private static final int METHODID_FREE_VIRTUAL_EXTERNAL_IP = 8;
  private static final int METHODID_PAIR_DEVICES = 9;
  private static final int METHODID_LIST_USERS = 10;
  private static final int METHODID_QUERY_ROLE = 11;
  private static final int METHODID_CHANGE_ROLE = 12;
  private static final int METHODID_RESET_PASSWORD = 13;
  private static final int METHODID_DELETE_USER = 14;
  private static final int METHODID_CREATE_USER = 15;
  private static final int METHODID_FORCE_RESET_PASSWORD = 16;
  private static final int METHODID_GEN_USER_INVITATION = 17;
  private static final int METHODID_LOGOUT = 18;
  private static final int METHODID_CHECK_PASSWORD = 19;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final HPortalSysImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(HPortalSysImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_QUERY_LOGIN:
          serviceImpl.queryLogin((cloud.lazycat.apis.sys.PortalServer.AuthToken) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.LoginInfo>) responseObserver);
          break;
        case METHODID_LIST_DEVICES:
          serviceImpl.listDevices((cloud.lazycat.apis.sys.PortalServer.ListDeviceRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.ListDeviceReply>) responseObserver);
          break;
        case METHODID_QUERY_DEVICE_BY_ID:
          serviceImpl.queryDeviceByID((cloud.lazycat.apis.sys.PortalServer.DeviceID) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.Device>) responseObserver);
          break;
        case METHODID_QUERY_BOX_INFO:
          serviceImpl.queryBoxInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.BoxInfo>) responseObserver);
          break;
        case METHODID_GET_DOMAIN_CERT:
          serviceImpl.getDomainCert((cloud.lazycat.apis.sys.PortalServer.DomainCertRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.DomainCertReply>) responseObserver);
          break;
        case METHODID_GET_DOMAIN_SELF_CERT:
          serviceImpl.getDomainSelfCert((cloud.lazycat.apis.sys.PortalServer.DomainCertRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.DomainCertReply>) responseObserver);
          break;
        case METHODID_GET_APP_CERT:
          serviceImpl.getAppCert((cloud.lazycat.apis.sys.PortalServer.AppCertRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.AppCertReply>) responseObserver);
          break;
        case METHODID_ALLOC_VIRTUAL_EXTERNAL_IP:
          serviceImpl.allocVirtualExternalIP((cloud.lazycat.apis.sys.PortalServer.AllocVEIPRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.AllocVEIPReply>) responseObserver);
          break;
        case METHODID_FREE_VIRTUAL_EXTERNAL_IP:
          serviceImpl.freeVirtualExternalIP((cloud.lazycat.apis.sys.PortalServer.FreeVEIPRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.FreeVEIPReply>) responseObserver);
          break;
        case METHODID_PAIR_DEVICES:
          serviceImpl.pairDevices((cloud.lazycat.apis.sys.PortalServer.PairDevicesRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_LIST_USERS:
          serviceImpl.listUsers((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.ListUsersReply>) responseObserver);
          break;
        case METHODID_QUERY_ROLE:
          serviceImpl.queryRole((cloud.lazycat.apis.sys.PortalServer.UserID) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.QueryRoleReply>) responseObserver);
          break;
        case METHODID_CHANGE_ROLE:
          serviceImpl.changeRole((cloud.lazycat.apis.sys.PortalServer.ChangeRoleReqeust) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_RESET_PASSWORD:
          serviceImpl.resetPassword((cloud.lazycat.apis.sys.PortalServer.ResetPasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_DELETE_USER:
          serviceImpl.deleteUser((cloud.lazycat.apis.sys.PortalServer.DeleteUserRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_CREATE_USER:
          serviceImpl.createUser((cloud.lazycat.apis.sys.PortalServer.CreateUserRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_FORCE_RESET_PASSWORD:
          serviceImpl.forceResetPassword((cloud.lazycat.apis.sys.PortalServer.ForceResetPasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_GEN_USER_INVITATION:
          serviceImpl.genUserInvitation((cloud.lazycat.apis.sys.PortalServer.GenUserInvitationRequest) request,
              (io.grpc.stub.StreamObserver<cloud.lazycat.apis.sys.PortalServer.UserInvitation>) responseObserver);
          break;
        case METHODID_LOGOUT:
          serviceImpl.logout((cloud.lazycat.apis.sys.PortalServer.LogoutRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
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

  private static abstract class HPortalSysBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    HPortalSysBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return cloud.lazycat.apis.sys.PortalServer.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("HPortalSys");
    }
  }

  private static final class HPortalSysFileDescriptorSupplier
      extends HPortalSysBaseDescriptorSupplier {
    HPortalSysFileDescriptorSupplier() {}
  }

  private static final class HPortalSysMethodDescriptorSupplier
      extends HPortalSysBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    HPortalSysMethodDescriptorSupplier(String methodName) {
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
      synchronized (HPortalSysGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new HPortalSysFileDescriptorSupplier())
              .addMethod(getQueryLoginMethod())
              .addMethod(getListDevicesMethod())
              .addMethod(getQueryDeviceByIDMethod())
              .addMethod(getQueryBoxInfoMethod())
              .addMethod(getGetDomainCertMethod())
              .addMethod(getGetDomainSelfCertMethod())
              .addMethod(getGetAppCertMethod())
              .addMethod(getAllocVirtualExternalIPMethod())
              .addMethod(getFreeVirtualExternalIPMethod())
              .addMethod(getPairDevicesMethod())
              .addMethod(getListUsersMethod())
              .addMethod(getQueryRoleMethod())
              .addMethod(getChangeRoleMethod())
              .addMethod(getResetPasswordMethod())
              .addMethod(getDeleteUserMethod())
              .addMethod(getCreateUserMethod())
              .addMethod(getForceResetPasswordMethod())
              .addMethod(getGenUserInvitationMethod())
              .addMethod(getLogoutMethod())
              .addMethod(getCheckPasswordMethod())
              .build();
        }
      }
    }
    return result;
  }
}
