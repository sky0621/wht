#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
echo "${SCRIPT_DIR}"
cd "${SCRIPT_DIR}" && cd ../

go run ./tools/gqlgen/global-object-id-gen/main.go

cd ./src
rm -f ./adapter/web/dummyresolvers/*.go
rm -f ./adapter/web/graphql_generated.go
rm -f ./adapter/web/gqlmodel/graphql_generated.go

subCmd=${1:-}

# https://gqlgen.com/
# https://github.com/99designs/gqlgen
#go get -u github.com/99designs/gqlgen@v0.11.3
gqlgen "${subCmd}"
