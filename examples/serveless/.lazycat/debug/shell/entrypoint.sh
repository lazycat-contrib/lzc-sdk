#!/bin/sh
set -ex

# generate host keys if not present
ssh-keygen -A

rsync --daemon

# do not detach (-D), log to stderr (-e), passthrough other arguments
exec /usr/sbin/sshd -D -e "$@"
