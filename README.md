[![GitHub Release](https://img.shields.io/github/v/release/astrokube/pathify)](https://github.com/astrokube/pathify/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/astrokube/pathify.svg)](https://pkg.go.dev/github.com/astrokube/pathify)
[![go.mod](https://img.shields.io/github/go-mod/go-version/astrokube/pathify)](go.mod)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://img.shields.io/github/license/astrokube/pathify)
[![Build Status](https://img.shields.io/github/actions/workflow/status/astrokube/pathify/build.yml?branch=main)](https://github.com/astrokube/pathify/actions?query=workflow%3ABuild+branch%3Amain)
[![CodeQL](https://github.com/astrokube/pathify/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/astrokube/pathify/actions/workflows/codeql.yml)

# Pathify

The swiss knife to dea with the hassle of manipulate generic 

## History and project status

This module is already `ready-for-production` and the [astrokube organization](https://www.github.com/astrokube) already
take advantage of it for our internal projects.

## Pathify  Highlights

* **Easy integration**: It's straightforward to be integrated with your current developments. 

## Installation

Use go get to retrieve the library to add it to your GOPATH workspace, or project's Go module dependencies.

```bash
go get -u github.com/astrokube/pathify
```

To update the library use go get -u to retrieve the latest version of it.

```bash
go get -u github.com/astrokube/pathify
```

You could specify a concrete version of this module as It's shown on the below. Replace x.y.z by the desired version.

```bash
module github.com/<org>/<repository>
require ( 
  github.com/astrokube/pathify vX.Y.Z
)
```

## Getting started

### Pre-requisites

* Go 1.19+

### Examples

A rich and growing set of examples of usage of this module can be found in folder `examples`.

```go
package main

import (
	"strings"

	"github.com/astrokube/pathify"
)

var peopleArray = []any{
	map[string]any{
		"firstname": "John",
		"lastname":  "Doe",
		"age":       29,
	},
	map[string]any{
		"firstname": "Jane",
		"lastname":  "Moe",
		"age":       30,
	},
}

func main() {
	p := pathify.Load[[]any](peopleArray).Set(
		"[1].lastname", "Doe",
		"[0].firstname", "Wendy",
		"[2].firstname", "Cindy",
		"[1].firstname", strings.ToUpper,
	)
	b, _ := p.YAML()
	println(string(b))
	b, _ = p.JSON()
	println(string(b))
}
```

**Output**
```bash
- age: 29
  firstname: Wendy
  lastname: Doe
- age: 30
  firstname: JANE
  lastname: Doe
- firstname: Cindy

[{"age":29,"firstname":"Wendy","lastname":"Doe"},{"age":30,"firstname":"JANE","lastname":"Doe"},{"firstname":"Cindy"}]

```

### Contributing

See the [contributing](https://github.com/astrokube/pathify/blob/main/CONTRIBUTING.md) documentation.


