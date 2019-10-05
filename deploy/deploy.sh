#!/bin/sh

set -eou pipefail

REGION="westus"
DNS_LABEL="gifm-app"
IMAGE=$1

echo "Deploying $IMAGE to group $RES_GROUP in $REGION"

az container create \
-g $RES_GROUP \
-l $REGION \
--ip-address public \
--image $IMAGE \
--ports 3000 \
-e "AUTH_CALLBACK=$AUTH_CALLBACK" "DATABASE_URL=$DATABASE_URL" "GITHUB_KEY=$GITHUB_KEY" "GITHUB_SECRET=$GITHUB_SECRET" "SESSION_SECRET=$SESSION_SECRET" \
-n $NAME


