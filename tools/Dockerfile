FROM debian:12 as build
COPY cache/usr/bin/ /usr/bin/
COPY cache/usr/boot/ /usr/boot/
COPY cache/usr/conf/ /usr/conf/
COPY cache/usr/lib/ /usr/lib/
COPY cache/usr/include/google/ /usr/include/google/
ENV PATH="/usr/bin/node_modules/npm/bin:$PATH"
SHELL ["bash", "-c"]
RUN <<EOF
sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources
apt update
apt-mark hold libfreetype6 libfontconfig1 libasound2
apt install /usr/bin/jre.deb /usr/bin/adoptium-ca.deb --fix-broken --no-install-recommends --no-install-suggests --no-upgrade
npm config set registry https://registry.npmmirror.com
npm install -g ts-proto@1.155.1 typescript@5
apt autoremove -y
apt clean all -y
rm -rf /usr/bin/*.deb /usr/lib/x86_64-linux-gnu/perl-base /var/lib/apt/lists /root/.npm /usr/share/{locale,doc,man,fonts} /usr/lib/node_modules/typescript/lib/tsserver.js /usr/lib/node_modules/ts-proto/node_modules/dprint-node/{dprint-node.win32-x64-msvc.node,dprint-node.linux-x64-musl.node,dprint-node.linux-arm64-musl.node,dprint-node.linux-arm64-gnu.node,dprint-node.darwin-x64.node,dprint-node.darwin-arm64.node}
EOF

FROM swift:5.10 as swift
COPY cache/usr/swift-protobuf.tar.gz /
RUN  <<EOF
tar -xf /swift-protobuf.tar.gz --strip-components 1
swift build -c release --static-swift-stdlib
strip .build/release/protoc-gen-swift
EOF

FROM debian:12
RUN --mount=type=bind,from=build,source=/,target=/build --mount=type=bind,from=swift,source=/,target=/swift --mount=type=bind,from=awesomecontainers/static-rsync:3.2.7,source=/,target=/rsync <<EOF
/rsync/bin/rsync -a --delete /build/ / # better than --squash
cp /swift/.build/release/protoc-gen-swift /usr/bin/
EOF
ENV PATH="/usr/bin/node_modules/npm/bin:$PATH"
CMD ["/usr/bin/protoc"]
