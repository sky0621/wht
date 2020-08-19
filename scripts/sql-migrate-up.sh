#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
echo "${SCRIPT_DIR}"
cd "${SCRIPT_DIR}" && cd ../db

env=${1:-}

if [ -z "${env}" ]; then
  echo -n "Select env(local/localtest/dev): "
  read env
fi

# https://github.com/rubenv/sql-migrate
#go get -v github.com/rubenv/sql-migrate/...
sql-migrate up -env="${env}"
