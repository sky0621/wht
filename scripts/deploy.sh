#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}

project=${1:-}
if [[ -z "${project}" ]]; then
  echo -n "need project"
  exit 1
fi

db_user=${2:-}
if [[ -z "${db_user}" ]]; then
  echo -n "need db_user"
  exit 1
fi

db_pass=${3:-}
if [[ -z "${db_pass}" ]]; then
  echo -n "need db_pass"
  exit 1
fi

db_name=${4:-}
if [[ -z "${db_name}" ]]; then
  echo -n "need db_name"
  exit 1
fi

db_instance_conn=${5:-}
if [[ -z "${db_instance_conn}" ]]; then
  echo -n "need db_instance_conn"
  exit 1
fi

#db_host=${5:-}
#if [[ -z "${db_host}" ]]; then
#  echo -n "need db_host"
#  exit 1
#fi

#db_port=${6:-}
#if [[ -z "${db_port}" ]]; then
#  echo -n "need db_port"
#  exit 1
#fi

gcloud run deploy wht \
  --image gcr.io/${project}/wht:latest \
  --platform managed \
  --project ${project} \
  --allow-unauthenticated \
  --region asia-northeast1 \
  --set-env-vars WHT_DB_USER=${db_user},WHT_DB_PASS=${db_pass},WHT_DB_INSTANCE_CONNECTION_NAME=${db_instance_conn},WHT_DB_NAME=${db_name}
#  --set-env-vars WHT_DB_USER=${db_user},WHT_DB_PASS=${db_pass},WHT_DB_HOST=${db_host},WHT_DB_PORT=${db_port},WHT_DB_NAME=${db_name}
