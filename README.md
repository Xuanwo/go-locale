# go-locale

[![Build Status](https://travis-ci.com/Xuanwo/go-locale.svg?branch=master)](https://travis-ci.com/Xuanwo/go-locale)
[![GoDoc](https://godoc.org/github.com/Xuanwo/go-locale?status.svg)](https://godoc.org/github.com/Xuanwo/go-locale)
[![Go Report Card](https://goreportcard.com/badge/github.com/Xuanwo/go-locale)](https://goreportcard.com/report/github.com/Xuanwo/go-locale)
[![codecov](https://codecov.io/gh/Xuanwo/go-locale/branch/master/graph/badge.svg)](https://codecov.io/gh/Xuanwo/go-locale)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/Xuanwo/go-locale/blob/master/LICENSE)

`go-locale` is a Golang lib for cross platform locale detection.

## Usage

```go
import (
    "github.com/Xuanwo/go-locale"
)

func main() {
	tag, err := locale.Detect()
    if err != nil {
        log.Fatal(err)
    }
    // Have fun with language.Tag!
}
```

## Acknowledgments

Inspired by [jibber_jabber](https://github.com/cloudfoundry-attic/jibber_jabber)