#!/bin/bash
set -e

godep go test -timeout 60s ./...
GOMAXPROCS=4 godep go test -timeout 60s -race ./...
