# Copyright 2018 Igor Zibarev
# Copyright 2018 The envexec Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: all
all: deps test

.PHONY: deps
deps:
	go mod tidy
	go mod vendor

.PHONY: test
test:
	go test -v ./...

.PHONY: images
images:
	docker build -f docker/scratch/Dockerfile -t hypnoglow/envexec:latest-scratch .
	docker build -f docker/alpine/Dockerfile -t hypnoglow/envexec:latest-alpine .
