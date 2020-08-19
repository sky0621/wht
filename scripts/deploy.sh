#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
# shellcheck disable=SC2086
echo ${SCRIPT_DIR}

project=$(gcloud secrets versions access latest --secret="project-id")
if [[ -z "${project}" ]]; then
  echo -n "need project"
  exit 1
fi
echo "${project}"

db_host=$(gcloud secrets versions access latest --secret="db-host")
if [[ -z "${db_host}" ]]; then
  echo -n "need db-host"
  exit 1
fi
# shellcheck disable=SC2086
echo ${db_host}

db_user=$(gcloud secrets versions access latest --secret="db-user")
if [[ -z "${db_user}" ]]; then
  echo -n "need db-user"
  exit 1
fi
echo "${db_user}"

db_pass=$(gcloud secrets versions access latest --secret="db-pass")
if [[ -z "${db_pass}" ]]; then
  echo -n "need db-pass"
  exit 1
fi
echo "${db_pass}"

db_name=$(gcloud secrets versions access latest --secret="db-name")
if [[ -z "${db_name}" ]]; then
  echo -n "need db-name"
  exit 1
fi
# shellcheck disable=SC2086
echo ${db_name}

gcloud run deploy wht \
  --image gcr.io/"${project}"/wht:latest \
  --platform managed \
  --project "${project}" \
  --region asia-northeast1 \
  --allow-unauthenticated \
  --add-cloudsql-instances "${db_host}" \
  --set-env-vars WHT_ENV=gcp \
  --set-env-vars WHT_DB_HOST="${db_host}" \
  --set-env-vars WHT_DB_USER="${db_user}" \
  --set-env-vars WHT_DB_PASS="${db_pass}" \
  --set-env-vars WHT_DB_NAME="${db_name}" \
  --set-env-vars WHT_IMAGE_CONTENTS_BUCKET="wht_image_contents_bucket" \
  --set-env-vars WHT_TRACE=true

# MEMO: use all user access
#  --allow-unauthenticated \
