#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}
cd ${SCRIPT_DIR} && cd ../src/adapter/controller/gqlmodel

# https://github.com/vektah/dataloaden
#go get -u github.com/vektah/dataloaden@v0.3.0
dataloaden ContentLoader int64 []github.com/sky0621/wht/adapter/controller/gqlmodel.Content
