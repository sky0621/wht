#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
# shellcheck disable=SC2086
echo ${SCRIPT_DIR}
cd "${SCRIPT_DIR}" && cd ../

env_api_endpoint=${1:-}

if [ -z "${env_api_endpoint}" ]; then
  echo "no api-endpoint, so get by secret"

  api_endpoint=$(gcloud secrets versions access latest --secret="api-endpoint")
  if [[ -z "${api_endpoint}" ]]; then
    echo -n "need api-endpoint"
    exit 1
  fi
  echo "${api_endpoint}"

  env_api_endpoint="${api_endpoint}"
fi

cd ./view
rm -fr node_modules
rm -fr frontendgen
yarn install

export WHT_API_ENDPOINT="${env_api_endpoint}"
yarn run nuxt-ts generate
