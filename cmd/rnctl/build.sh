#!/bin/bash

# At the root directory
cd $GOPATH/src/k8s.io/release

# Run dep to get the dependencies
dep ensure -vendor-only

# Run gazelle to auto-generate BUILD files
bazel run //:gazelle

# Build the program
bazel build cmd/rnctl:rnctl
