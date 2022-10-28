#!/bin/bash -x
go test ./example
go test -trimpath ./example
bazel test --test_output=errors //...
