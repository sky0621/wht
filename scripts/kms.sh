#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}
cd ${SCRIPT_DIR} && cd ../

cmd=${1:-}

if [ -z "${cmd}" ]; then
  echo -n "Select command(mkkeyring/mkkey/enc/dec): "
  read cmd
fi

case "${cmd}" in
  "mkkeyring" ) gcloud kms keyrings create env --location global ;;
  "mkkey"     ) gcloud kms keys create app-credential --location global --keyring env --purpose encryption ;;
  "enc"       ) gcloud kms encrypt --location global --keyring env --key app-credential --plaintext-file app-credential.json --ciphertext-file app-credential.json.encrypted ;;
  "dec"       ) gcloud kms decrypt --location global --keyring env --key app-credential --ciphertext-file app-credential.json.encrypted --plaintext-file app-credential.json ;;
  *           ) echo "invalid command: ${cmd}"
                exit 1;;
esac
