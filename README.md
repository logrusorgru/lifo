lifo
====

[![GoDoc](https://godoc.org/github.com/logrusorgru/lifo?status.svg)](https://godoc.org/github.com/logrusorgru/lifo)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/lifo.svg)](https://travis-ci.org/logrusorgru/lifo)
[![Coverage Status](https://coveralls.io/repos/logrusorgru/lifo/badge.svg?branch=master)](https://coveralls.io/r/logrusorgru/lifo?branch=master)

golang lifo reader/writer buffer like truncated bytes.Buffer but lifo

### Install

```bash
go get github.com/logrusorgru/lifo
cd $GOPATH/src/github.com/logrusorgru/lifo
go test
```

### Methods

- Read
- Write
- WriteTo
- ReadByte
- WriteByte
- Len
- Next

### Example

```go
package main

import (
	"fmt"
	"github.com/logrusorgru/lifo"
)

func main() {
	b := lifo.NewBuffer(nil)

	b.Write([]byte("bye"))
	b.Write([]byte("hello"))

	hello := make([]byte, 5)
	bye := make([]byte, 3)

	b.Read(hello)
	b.Read(bye)

	fmt.Println(string(hello))
	fmt.Println(string(bye))

	// Output
	// hello
	// bye
}
```

### Licensing

Copyright &copy; 2015 Konstantin Ivanov <ivanov.konstantin@logrus.org.ru>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE.md file for more details.

