<img src="assets\taxediologolandscape.jpg" alt="drawing" width="200"/>

<h1 align="center">
  ISO 639 Language Codes
</h1>

<h3 align="center">
  <!-- <a href="https://pkg.go.dev/github.com/KalbiProject/Kalbi">Documentation</a> •  -->
  <a href="https://taxed.io">taxed.io</a>
</h3>

[![Go Report Card](https://goreportcard.com/badge/github.com/taxedio/iso639)](https://goreportcard.com/report/github.com/taxedio/iso639)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/taxedio/iso639/LICENCE)  
[![GitHub contributors](https://img.shields.io/github/contributors/taxedio/iso639)](https://github.com/taxedio/iso639/graphs/contributors)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/taxedio/iso639)

# Introduction

"ISO 639 is a standardized nomenclature used to classify languages. Each language is assigned a two-letter (639-1) and three-letter (639-2 and 639-3) lowercase abbreviation, amended in later versions of the nomenclature."


# Basic Example

```GO
package main

import (
	"fmt"

	"github.com/taxedio/iso639"
)

var (
	enName  string = "English"
	frName  string = "anglais"
	alph2   string = "EN"
	alph3   string = "ENG"
)

func main() {
	// English Names, codes can be found via English name (enName in Struct)
	fmt.Printf(`iso639.Alpha3Match("English") - %v`, *iso639.Alpha3Match(enName))
	fmt.Printf(`iso639.Alpha2Match("English") - %v`, *iso639.Alpha2Match(enName))

  // English Names with different casing and spaces will still return requested codes
	fmt.Printf(`iso639.Alpha3Match(" english   ") - %v`, *iso639.Alpha3Match(" english   "))
	fmt.Printf(`iso639.Alpha2Match(" english   ") - %v`, *iso639.Alpha2Match(" english   "))

  // French Names, codes can be found via English name (frName in Struct)
	fmt.Printf(`iso639.Alpha3Match("anglais") - %v`, *iso639.Alpha3Match(frName))
	fmt.Printf(`iso639.Alpha2Match("anglais") - %v`, *iso639.Alpha2Match(frName))

  // Alpha Codes, both can return the other:
	fmt.Printf(`iso639.Alpha3Match("EN") - %v`, *iso639.Alpha3Match(alph2))
	fmt.Printf(`iso639.Alpha2Match("ENG") - %v`, *iso639.Alpha2Match(alph3))

  // Diacritics, (IN DEVELOPMENT). If e is used instead of é then the values will still match:
	fmt.Printf(`iso639.Alpha2Match("éwé") - %v`, *iso639.Alpha3Match("éwé"))
	fmt.Printf(`iso639.Alpha2Match("éwé") - %v`, *iso639.Alpha2Match("éwé"))
	fmt.Printf(`iso639.Alpha3Match("ewe") - %v`, *iso639.Alpha3Match("ewe"))
	fmt.Printf(`iso639.Alpha3Match("ewe") - %v`, *iso639.Alpha2Match("ewe"))

  // Returning Structs, all of the rules and examples above work using the StructMatch
	fmt.Printf(`iso639.StructMatch("English") - %v`, *iso639.StructMatch(enName))
	fmt.Printf(`iso639.StructMatch("  english  ") - %v`, *iso639.StructMatch(" english  "))
	fmt.Printf(`iso639.StructMatch("anglais") - %v`, *iso639.StructMatch(frName))
	fmt.Printf(`iso639.StructMatch("EN") - %v`, *iso639.StructMatch(alph2))
	fmt.Printf(`iso639.StructMatch("ENG") - %v`, *iso639.StructMatch(alph3))
	fmt.Printf(`iso639.StructMatch("éwé") - %v`, *iso639.StructMatch("éwé"))
	fmt.Printf(`iso639.StructMatch("ewe  ") - %v`, *iso639.StructMatch("ewe  "))
}
```

**console**:

```stdout
// English Names, codes can be found via English name (enName in Struct)
iso639.Alpha3Match("English") - ENG
iso639.Alpha2Match("English") - EN

// English Names with different casing and spaces will still return requested codes
iso639.Alpha3Match(" english   ") - ENG
iso639.Alpha2Match(" english   ") - EN

// French Names, codes can be found via English name (frName in Struct)
iso639.Alpha3Match("anglais") - ENG
iso639.Alpha2Match("anglais") - EN

// Alpha Codes, both can return the other:
iso639.Alpha3Match("EN") - ENG
iso639.Alpha2Match("ENG") - EN

// Diacritics, (IN DEVELOPMENT). If e is used instead of é then the values will still match:
iso639.Alpha2Match("éwé") - EWE
iso639.Alpha2Match("éwé") - EE
iso639.Alpha3Match("ewe") - EWE
iso639.Alpha3Match("ewe") - EE

// Returning Structs, all of the rules and examples above work using the StructMatch
iso639.StructMatch("English") - {English anglais EN ENG XXX}
iso639.StructMatch("  english  ") - {English anglais EN ENG XXX}
iso639.StructMatch("anglais") - {English anglais EN ENG XXX}
iso639.StructMatch("EN") - {English anglais EN ENG XXX}
iso639.StructMatch("ENG") - {English anglais EN ENG XXX}
iso639.StructMatch("éwé") - {Ewe éwé EE EWE XXX}
iso639.StructMatch("ewe  ") - {Ewe éwé EE EWE XXX}
```
