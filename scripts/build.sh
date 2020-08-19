#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
# shellcheck disable=SC2086
echo ${SCRIPT_DIR}
cd "${SCRIPT_DIR}" && cd ../

project=$(gcloud secrets versions access latest --secret="project-id")
if [[ -z "${project}" ]]; then
  echo -n "need project"
  exit 1
fi

(
  cd "${SCRIPT_DIR}"
  ./kms.sh dec
)

sed -i -e 's/app-credential.json/#app-credential.json/' .gitignore

gcloud builds submit --tag gcr.io/"${project}"/wht:latest .

sed -i -e 's/#app-credential.json/app-credential.json/' .gitignore

rm app-credential.json
