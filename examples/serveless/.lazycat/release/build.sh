#!/bin/sh
set -ex


SHELLDIR="$(dirname "$(realpath "$0")")"

cd "$SHELLDIR"

npm run build
