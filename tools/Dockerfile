FROM debian:12
COPY cache/usr/bin/ /usr/bin/
COPY cache/usr/boot/ /usr/boot/
COPY cache/usr/conf/ /usr/conf/
COPY cache/usr/lib/ /usr/lib/
COPY cache/usr/include/google/ /usr/include/google/
ENV PATH="/usr/bin/node_modules/npm/bin:$PATH"
SHELL ["bash", "-c"]
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources ;\
    apt update ;\
    apt install -y /usr/bin/adoptium-ca.deb --fix-broken --no-install-recommends --no-install-suggests ;\
    apt install -y /usr/bin/jre.deb --fix-broken --no-install-recommends --no-install-suggests ;\
    npm config set registry https://registry.npmmirror.com ;\
    npm install -g ts-proto@1.155.1 typescript@5 ;\
    apt autoremove -y ;\
    apt clean all -y ;\
    rm -rf /usr/bin/*.deb /usr/lib/x86_64-linux-gnu/perl-base /var/lib/apt/lists /root/.npm /usr/share/{locale,doc,man,fonts} /usr/lib/node_modules/typescript/lib/tsserver.js /usr/lib/node_modules/ts-proto/node_modules/dprint-node/{dprint-node.win32-x64-msvc.node,dprint-node.linux-x64-musl.node,dprint-node.linux-arm64-musl.node,dprint-node.linux-arm64-gnu.node,dprint-node.darwin-x64.node,dprint-node.darwin-arm64.node}
CMD ["/usr/bin/protoc"]
