#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}
cd ${SCRIPT_DIR} && cd ../src

rm -rf ./adapter/rdb/boiled/*

# https://github.com/volatiletech/sqlboiler
#go get -u github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
# Do not forget the trailing /v4
#go get -u github.com/volatiletech/sqlboiler/v4
# Assuming you're going to use the null package for its additional null types
# Do not forget the trailing /v8
#go get -u github.com/volatiletech/null/v8

sqlboiler --wipe psql
