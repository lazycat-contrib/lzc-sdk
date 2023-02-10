// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: localdevice/config.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

public struct Cloud_Lazycat_Apis_Localdevice_GetUserConfigRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var boxID: String = String()

  public var userID: String = String()

  public var appID: String = String()

  public var configKey: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Localdevice_GetUserConfigResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var configValue: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Lazycat_Apis_Localdevice_GetUserConfigRequest: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Localdevice_GetUserConfigResponse: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.lazycat.apis.localdevice"

extension Cloud_Lazycat_Apis_Localdevice_GetUserConfigRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GetUserConfigRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "boxId"),
    2: .same(proto: "userId"),
    4: .same(proto: "appId"),
    3: .same(proto: "configKey"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.boxID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.userID) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.configKey) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.appID) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.boxID.isEmpty {
      try visitor.visitSingularStringField(value: self.boxID, fieldNumber: 1)
    }
    if !self.userID.isEmpty {
      try visitor.visitSingularStringField(value: self.userID, fieldNumber: 2)
    }
    if !self.configKey.isEmpty {
      try visitor.visitSingularStringField(value: self.configKey, fieldNumber: 3)
    }
    if !self.appID.isEmpty {
      try visitor.visitSingularStringField(value: self.appID, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Localdevice_GetUserConfigRequest, rhs: Cloud_Lazycat_Apis_Localdevice_GetUserConfigRequest) -> Bool {
    if lhs.boxID != rhs.boxID {return false}
    if lhs.userID != rhs.userID {return false}
    if lhs.appID != rhs.appID {return false}
    if lhs.configKey != rhs.configKey {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Localdevice_GetUserConfigResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GetUserConfigResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "configValue"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.configValue) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.configValue.isEmpty {
      try visitor.visitSingularStringField(value: self.configValue, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Localdevice_GetUserConfigResponse, rhs: Cloud_Lazycat_Apis_Localdevice_GetUserConfigResponse) -> Bool {
    if lhs.configValue != rhs.configValue {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
