// swift-tools-version: 5.6
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "Cloud_Lazycat_Apis",

    products: [
        // Products define the executables and libraries a package produces, and make them visible to other packages.
        .library(
            name: "Cloud_Lazycat_Apis",
            targets: ["common", "localdevice", "sys"]
        ),
    ],
    dependencies: [
        // Dependencies declare other packages that this package depends on.
        .package(name: "SwiftProtobuf", url: "https://github.com/apple/swift-protobuf.git", from: "1.19.1"),
    ],
    targets: [
        // Targets are the basic building blocks of a package. A target can define a module or a test suite.
        // Targets can depend on other targets in this package, and on products in packages this package depends on.
        .target(
            name: "common",
            dependencies: ["SwiftProtobuf", "sys"]
        ),
        .target(
            name: "localdevice",
            dependencies: ["SwiftProtobuf"]
        ),
        .target(
            name: "sys",
            dependencies: ["SwiftProtobuf"]
        ),
    ]
)
