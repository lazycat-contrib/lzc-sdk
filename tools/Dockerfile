FROM archlinux:base-devel-20230716.0.165339
RUN echo 'Server = https://mirrors.tuna.tsinghua.edu.cn/archlinux/$repo/os/$arch' > /etc/pacman.d/mirrorlist

# RUN pacman-db-upgrade
# RUN pacman -Syyu --noconfirm
RUN pacman -Sy --noconfirm grpc go nodejs wget

RUN pacman -Sy --noconfirm jdk8-openjdk maven npm

RUN wget https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/1.47.0/protoc-gen-grpc-java-1.47.0-linux-x86_64.exe \
    -O /usr/bin/protoc-gen-grpc-java
RUN chmod a+x /usr/bin/protoc-gen-grpc-java

ENV GOPATH=/go PATH=/go/bin/:$PATH
ENV GOPROXY=https://goproxy.cn

RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.3
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

RUN npm config set registry https://registry.npmmirror.com/

RUN npm install -g ts-proto@1.155.1

CMD ["/usr/bin/protoc"]
