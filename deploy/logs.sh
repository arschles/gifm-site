#!/bin/sh

set -eou pipefail

az container logs --follow -g $RES_GROUP -n $NAME
