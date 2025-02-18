#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -eux -o pipefail

# Go to the root of the boundary repo
root_dir="$(git rev-parse --show-toplevel)"
pushd "${root_dir}" > /dev/null

# make docker image
export IMAGE_TAG_DEV="${IMAGE_NAME}"
make build-ui docker-build-dev

# make the cli to be used by the test runner
export GOOS=linux
make build
zip -j ${ARTIFACT_PATH}/boundary.zip bin/boundary

popd > /dev/null
