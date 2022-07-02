FROM registry.lazycat.cloud/lzc/lzcapp:0.1
ARG HTTP_SERVICE_PORT

WORKDIR /app
COPY dist .
EXPOSE $HTTP_SERVICE_PORT

ENV HTTP_SERVICE_PORT ${HTTP_SERVICE_PORT}
CMD lzcapp --listen :8080 -up "/=file:///app"
