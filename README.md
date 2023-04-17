[![GitHub Release](https://img.shields.io/github/v/release/{{.repository.owner}}/{{.repository.name}})](https://github.com/{{.repository.owner}}/{{.repository.name}}/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/{{.repository.owner}}/{{.repository.name}}.svg)](https://pkg.go.dev/github.com/{{.repository.owner}}/{{.repository.name}})
[![go.mod](https://img.shields.io/github/go-mod/go-version/{{.repository.owner}}/{{.repository.name}})](go.mod)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://img.shields.io/github/license/{{.repository.owner}}/{{.repository.name}})
[![Build Status](https://img.shields.io/github/actions/workflow/status/{{.repository.owner}}/{{.repository.name}}/build.yml?branch=main)](https://github.com/{{.repository.owner}}/{{.repository.name}}/actions?query=workflow%3ABuild+branch%3Amain)
[![CodeQL](https://github.com/{{.repository.owner}}/{{.repository.name}}/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/{{.repository.owner}}/{{.repository.name}}/actions/workflows/codeql.yml)

# {{.project}}

{{.description}}

## History and project status

This module is still `in active development` and the API is still subject to breaking changes.

## Installation

Use go get to retrieve the library to add it to your GOPATH workspace, or project's Go module dependencies.

```bash
go get -u github.com/{{.repository.owner}}/{{.repository.name}}
```

To update the library use go get -u to retrieve the latest version of it.

```bash
go get -u github.com/{{.repository.owner}}/{{.repository.name}}
```

You could specify a concrete version of this module as It's shown on the below. Replace x.y.z by the desired version.

```bash
module github.com/<org>/<repository>
require ( 
  github.com/{{.repository.owner}}/{{.repository.name}} vX.Y.Z
)
```

## Overview of packages

The library is composed by:

* `package`: .....

## Getting started

### Pre-requisites

* Go 1.19+

### Examples

A rich and growing set of examples of usage of this module can be found in folder `examples`.


### Contributing

See the [contributing](https://github.com/{{.repository.owner}}/{{.repository.name}}/blob/main/CONTRIBUTING.md) documentation.


