#!/bin/bash

: "${PROJECT:?is required}"
: "${SERVICE:?is required}"
: "${TARGET_HOST:?is required}"

env_variables=$(mktemp)
trap "rm $env_variables" 0

cat <<EOS >> $env_variables
service: ${SERVICE}
env_variables:
  TARGET_HOST: ${TARGET_HOST}
EOS

cat app_base.yaml $env_variables > app.yaml

gcloud app deploy --project="$PROJECT" app.yaml
