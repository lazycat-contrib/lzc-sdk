#!/bin/sh
set -ex

SHELLDIR="$(dirname "$(realpath "$0")")"

cd "$SHELLDIR"

mkdir -p debug/ssh
if [ -f debug/ssh/id_ed25519 ]; then
  rm -f debug/ssh/id_ed25519
  rm -f debug/ssh/id_ed25519.pub
fi
ssh-keygen -t ed25519 -f debug/ssh/id_ed25519 -q -P ""
chmod 0400 debug/ssh/id_ed25519
cat debug/ssh/id_ed25519.pub > debug/ssh/authorized_keys
