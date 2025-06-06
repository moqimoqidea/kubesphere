#!/usr/bin/env bash

# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script runs all *_test.go files. It is equivalent to `make test`.
# Usage: `hack/test-go.sh` or `make test`.
# Note: This script is a vestigial redirection. Please do not add "real" logic.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

kube::golang::setup_env

cd "${KUBE_ROOT}" || exit 1

ENVTEST_K8S_VERSION=1.28.x
KUBEBUILDER_ASSETS=$(setup-envtest use $ENVTEST_K8S_VERSION -p path --bin-dir="${KUBE_OUTPUT_BIN}")
export KUBEBUILDER_ASSETS

go test ./pkg/... ./cmd/... -covermode=atomic -coverprofile=coverage.txt
go -C staging/src/kubesphere.io/api test ./...
go -C staging/src/kubesphere.io/client-go test ./...