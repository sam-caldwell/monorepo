alphaNumericIdentifier
======================

## Description
A reusable alpha numeric identifier with no spaces, used by `virtualization` project and others where
the project needs an identifier consisting only of letters and numbers (starting with a letter) that
includes upper and lower case characters.

## Usage
```go
package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/types/alphaNumericIdentifier"
)

var identifier Identifier

idenfier.Set("myName1")
fmt.Println(idenfier.Get())

```
