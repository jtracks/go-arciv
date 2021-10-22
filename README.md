# go-arciv
![Tests](https://github.com/jtracks/go-arciv/actions/workflows/test.yml/badge.svg)
[![Docs](https://pkg.go.dev/badge/jtracks/go-arciv)](https://pkg.go.dev/github.com/jtracks/go-arciv)
## About
Unofficial golang package for interacting with the [arXiv API](https://arxiv.org/help/api/user-manual)

## Installation

``` bash
go get github.com/jtracks/go-arciv/arciv
```

## Documentation
[arciv package go doc](https://pkg.go.dev/github.com/jtracks/go-arciv/arciv)
## Usage

``` go
package main

import (
	"fmt"
	"github.com/jtracks/go-arciv/arciv"
)

func main() {

	result, _ = arciv.Search(
		arciv.SimpleQuery{
			Search:     "electron",
			MaxResults: 5,
		})

	for i, e := range result.Entries {
		fmt.Printf("Result %v: %v\n", i+1, e.Title)
	}
}

//> Result 1: Impact of Electron-Electron Cusp on Configuration Interaction Energies
//> Result 2: Electron thermal conductivity owing to collisions between degenerate electrons
//> Result 3: Electron pairing: from metastable electron pair to bipolaron
//> Result 4: Electron Temperature Anisotropy and Electron Beam Constraints From Electron //> Kinetic Instabilities in the Solar Wind
//> Result 5: Hamiltonian of a many-electron system with single-electron and electron-pair //> states in a two-dimensional periodic potential
```

## Disclaimer
This project is entirely independent from arXiv and arxiv.org.
## Acknowledgements 
Thank you to arXiv for use of its open access interoperability.