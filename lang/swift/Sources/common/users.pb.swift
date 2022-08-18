// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: common/users.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

import sys

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

public struct Cloud_Lazycat_Apis_Common_UserID {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uid: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_UserInfo {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uid: String = String()

  public var nickname: String = String()

  public var avatar: String = String()

  public var role: sys.Cloud_Lazycat_Apis_Sys_Role = .normal

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_ListUIDsRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_ListUIDsReply {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uids: [String] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_ChangeRoleReqeust {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uid: String = String()

  public var role: sys.Cloud_Lazycat_Apis_Sys_Role = .normal

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_ResetPasswordRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uid: String = String()

  public var oldPassword: String = String()

  public var newPassword: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_DeleteUserRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uid: String = String()

  /// 是否清理指定uid用户的数据目录
  public var clearUserData_p: Bool = false

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_CreateUserRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uid: String = String()

  public var password: String = String()

  public var role: sys.Cloud_Lazycat_Apis_Sys_Role = .normal

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_UpdateUserInfoRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uid: String = String()

  public var nickname: String = String()

  public var avatar: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Cloud_Lazycat_Apis_Common_ForceResetPasswordRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var uid: String = String()

  public var newPassword: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Lazycat_Apis_Common_UserID: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_UserInfo: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_ListUIDsRequest: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_ListUIDsReply: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_ChangeRoleReqeust: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_ResetPasswordRequest: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_DeleteUserRequest: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_CreateUserRequest: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_UpdateUserInfoRequest: @unchecked Sendable {}
extension Cloud_Lazycat_Apis_Common_ForceResetPasswordRequest: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.lazycat.apis.common"

extension Cloud_Lazycat_Apis_Common_UserID: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".UserID"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uid"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.uid) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uid.isEmpty {
      try visitor.visitSingularStringField(value: self.uid, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_UserID, rhs: Cloud_Lazycat_Apis_Common_UserID) -> Bool {
    if lhs.uid != rhs.uid {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_UserInfo: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".UserInfo"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uid"),
    2: .same(proto: "nickname"),
    3: .same(proto: "avatar"),
    4: .same(proto: "role"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.uid) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.nickname) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.avatar) }()
      case 4: try { try decoder.decodeSingularEnumField(value: &self.role) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uid.isEmpty {
      try visitor.visitSingularStringField(value: self.uid, fieldNumber: 1)
    }
    if !self.nickname.isEmpty {
      try visitor.visitSingularStringField(value: self.nickname, fieldNumber: 2)
    }
    if !self.avatar.isEmpty {
      try visitor.visitSingularStringField(value: self.avatar, fieldNumber: 3)
    }
    if self.role != .normal {
      try visitor.visitSingularEnumField(value: self.role, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_UserInfo, rhs: Cloud_Lazycat_Apis_Common_UserInfo) -> Bool {
    if lhs.uid != rhs.uid {return false}
    if lhs.nickname != rhs.nickname {return false}
    if lhs.avatar != rhs.avatar {return false}
    if lhs.role != rhs.role {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_ListUIDsRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ListUIDsRequest"
  public static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_ListUIDsRequest, rhs: Cloud_Lazycat_Apis_Common_ListUIDsRequest) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_ListUIDsReply: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ListUIDsReply"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uids"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeRepeatedStringField(value: &self.uids) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uids.isEmpty {
      try visitor.visitRepeatedStringField(value: self.uids, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_ListUIDsReply, rhs: Cloud_Lazycat_Apis_Common_ListUIDsReply) -> Bool {
    if lhs.uids != rhs.uids {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_ChangeRoleReqeust: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ChangeRoleReqeust"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uid"),
    2: .same(proto: "role"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.uid) }()
      case 2: try { try decoder.decodeSingularEnumField(value: &self.role) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uid.isEmpty {
      try visitor.visitSingularStringField(value: self.uid, fieldNumber: 1)
    }
    if self.role != .normal {
      try visitor.visitSingularEnumField(value: self.role, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_ChangeRoleReqeust, rhs: Cloud_Lazycat_Apis_Common_ChangeRoleReqeust) -> Bool {
    if lhs.uid != rhs.uid {return false}
    if lhs.role != rhs.role {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_ResetPasswordRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ResetPasswordRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uid"),
    2: .standard(proto: "old_password"),
    3: .standard(proto: "new_password"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.uid) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.oldPassword) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.newPassword) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uid.isEmpty {
      try visitor.visitSingularStringField(value: self.uid, fieldNumber: 1)
    }
    if !self.oldPassword.isEmpty {
      try visitor.visitSingularStringField(value: self.oldPassword, fieldNumber: 2)
    }
    if !self.newPassword.isEmpty {
      try visitor.visitSingularStringField(value: self.newPassword, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_ResetPasswordRequest, rhs: Cloud_Lazycat_Apis_Common_ResetPasswordRequest) -> Bool {
    if lhs.uid != rhs.uid {return false}
    if lhs.oldPassword != rhs.oldPassword {return false}
    if lhs.newPassword != rhs.newPassword {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_DeleteUserRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".DeleteUserRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uid"),
    2: .standard(proto: "clear_user_data"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.uid) }()
      case 2: try { try decoder.decodeSingularBoolField(value: &self.clearUserData_p) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uid.isEmpty {
      try visitor.visitSingularStringField(value: self.uid, fieldNumber: 1)
    }
    if self.clearUserData_p != false {
      try visitor.visitSingularBoolField(value: self.clearUserData_p, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_DeleteUserRequest, rhs: Cloud_Lazycat_Apis_Common_DeleteUserRequest) -> Bool {
    if lhs.uid != rhs.uid {return false}
    if lhs.clearUserData_p != rhs.clearUserData_p {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_CreateUserRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CreateUserRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uid"),
    2: .same(proto: "password"),
    3: .same(proto: "role"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.uid) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.password) }()
      case 3: try { try decoder.decodeSingularEnumField(value: &self.role) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uid.isEmpty {
      try visitor.visitSingularStringField(value: self.uid, fieldNumber: 1)
    }
    if !self.password.isEmpty {
      try visitor.visitSingularStringField(value: self.password, fieldNumber: 2)
    }
    if self.role != .normal {
      try visitor.visitSingularEnumField(value: self.role, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_CreateUserRequest, rhs: Cloud_Lazycat_Apis_Common_CreateUserRequest) -> Bool {
    if lhs.uid != rhs.uid {return false}
    if lhs.password != rhs.password {return false}
    if lhs.role != rhs.role {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_UpdateUserInfoRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".UpdateUserInfoRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uid"),
    2: .same(proto: "nickname"),
    3: .same(proto: "avatar"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.uid) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.nickname) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.avatar) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uid.isEmpty {
      try visitor.visitSingularStringField(value: self.uid, fieldNumber: 1)
    }
    if !self.nickname.isEmpty {
      try visitor.visitSingularStringField(value: self.nickname, fieldNumber: 2)
    }
    if !self.avatar.isEmpty {
      try visitor.visitSingularStringField(value: self.avatar, fieldNumber: 3)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_UpdateUserInfoRequest, rhs: Cloud_Lazycat_Apis_Common_UpdateUserInfoRequest) -> Bool {
    if lhs.uid != rhs.uid {return false}
    if lhs.nickname != rhs.nickname {return false}
    if lhs.avatar != rhs.avatar {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Cloud_Lazycat_Apis_Common_ForceResetPasswordRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ForceResetPasswordRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "uid"),
    2: .standard(proto: "new_password"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.uid) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.newPassword) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.uid.isEmpty {
      try visitor.visitSingularStringField(value: self.uid, fieldNumber: 1)
    }
    if !self.newPassword.isEmpty {
      try visitor.visitSingularStringField(value: self.newPassword, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Common_ForceResetPasswordRequest, rhs: Cloud_Lazycat_Apis_Common_ForceResetPasswordRequest) -> Bool {
    if lhs.uid != rhs.uid {return false}
    if lhs.newPassword != rhs.newPassword {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
