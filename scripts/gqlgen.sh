#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}
cd ${SCRIPT_DIR} && cd ../src

subCmd=${1:-}

# https://gqlgen.com/
# https://github.com/99designs/gqlgen
#go get -u github.com/99designs/gqlgen@v0.11.3
gqlgen ${subCmd}
