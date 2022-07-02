FROM archlinux:base-devel-20220626.0.64095
RUN echo 'Server = http://mirrors.tuna.tsinghua.edu.cn/archlinux/$repo/os/$arch' > /etc/pacman.d/mirrorlist

# RUN pacman-db-upgrade
# RUN pacman -Syyu --noconfirm
RUN pacman -Sy --noconfirm grpc go nodejs wget

RUN wget https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/1.47.0/protoc-gen-grpc-java-1.47.0-linux-x86_64.exe \
    -O /usr/bin/protoc-gen-grpc-java
RUN chmod a+x /usr/bin/protoc-gen-grpc-java

ENV GOPATH=/go PATH=/go/bin/:$PATH
ENV GOPROXY=https://goproxy.cn

RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

CMD ["/usr/bin/protoc"]
