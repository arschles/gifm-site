#!/bin/bash

set -eoux pipefail

az container show -g $RES_GROUP -n $NAME -o table
