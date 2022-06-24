FROM node:16-alpine3.12

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache openssh rsync \
  && echo "root:root" | chpasswd

EXPOSE 22

COPY sshd_config /etc/ssh/sshd_config

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]