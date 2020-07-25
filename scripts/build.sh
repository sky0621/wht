#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}
cd ${SCRIPT_DIR} && cd ../

project=${1:-}
if [[ -z "${project}" ]]; then
  echo -n "need project"
  exit 1
fi

gcloud builds submit --tag gcr.io/${project}/wht:latest .
