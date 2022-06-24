FROM golang:alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.bfsu.edu.cn/g' /etc/apk/repositories
RUN apk add tzdata git\
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone


ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GOPRIVATE=gitee.com/linakesi

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /usr/bin/lzc-clip

CMD ["/usr/bin/lzc-clip"]
