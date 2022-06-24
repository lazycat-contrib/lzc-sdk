#!/bin/sh
set -ex

# generate host keys if not present
ssh-keygen -A

/docker-entrypoint.sh nginx

# do not detach (-D), log to stderr (-e), passthrough other arguments
exec /usr/sbin/sshd -D -e "$@"
