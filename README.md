# reg-room-service

[![Go](../../actions/workflows/go.yml/badge.svg)](../../actions/workflows/go.yml)
[![CodeQL](../../actions/workflows/codeql-analysis.yml/badge.svg)](../../actions/workflows/codeql-analysis.yml)
[![Release](../../actions/workflows/release.yml/badge.svg)](../../actions/workflows/release.yml)

## Overview

A backend service that provides room management.

Implemented in go.

Command line arguments
```-config <path-to-config-file> [-migrate-database]```

## Installation

This service uses go modules to provide dependency management, see `go.mod`.

If you place this repository OUTSIDE of your gopath, `go build cmd/main.go` and
`go test ./...` will download all required dependencies by default.

Go 1.14 or later is required.
