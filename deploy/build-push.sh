#!/bin/sh

set -eou pipefail

cd ..
make docker-build
make docker-push
