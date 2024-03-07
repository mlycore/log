# Copyright 2024 mlycore. All rights reserved.
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

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: lint
lint: check-lint
	GOBIN=$(LOCALBIN) CGO_ENABLED=0 golangci-lint run -v --timeout=5m ./...

.PHONY: check-lint
check-lint: $(CHECK_LINT) ## Download golangci-lint-setup locally if necessary.
$(CHECK_LINT): $(LOCALBIN)
	GOBIN=$(LOCALBIN) CGO_ENABLED=0 go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest%