#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}

project=${1:-}
if [[ -z "${project}" ]]; then
  echo -n "need project"
  exit 1
fi

gcloud run deploy wht \
  --image gcr.io/${project}/wht:latest \
  --platform managed \
  --project ${project} \
  --allow-unauthenticated \
  --region asia-northeast1
