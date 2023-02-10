// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: localdevice/device.proto
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

public struct Cloud_Lazycat_Apis_Localdevice_DeviceInfo {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var os: String = String()

  public var cpu: String = String()

  public var name: String = String()

  public var documentRootDir: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Cloud_Lazycat_Apis_Localdevice_DeviceInfo: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "cloud.lazycat.apis.localdevice"

extension Cloud_Lazycat_Apis_Localdevice_DeviceInfo: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".DeviceInfo"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "OS"),
    2: .same(proto: "CPU"),
    3: .same(proto: "name"),
    4: .same(proto: "documentRootDir"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.os) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.cpu) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.documentRootDir) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.os.isEmpty {
      try visitor.visitSingularStringField(value: self.os, fieldNumber: 1)
    }
    if !self.cpu.isEmpty {
      try visitor.visitSingularStringField(value: self.cpu, fieldNumber: 2)
    }
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 3)
    }
    if !self.documentRootDir.isEmpty {
      try visitor.visitSingularStringField(value: self.documentRootDir, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Cloud_Lazycat_Apis_Localdevice_DeviceInfo, rhs: Cloud_Lazycat_Apis_Localdevice_DeviceInfo) -> Bool {
    if lhs.os != rhs.os {return false}
    if lhs.cpu != rhs.cpu {return false}
    if lhs.name != rhs.name {return false}
    if lhs.documentRootDir != rhs.documentRootDir {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
