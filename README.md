# go-locale

[![Build Status](https://travis-ci.com/Xuanwo/go-locale.svg?branch=master)](https://travis-ci.com/Xuanwo/go-locale)
[![GoDoc](https://godoc.org/github.com/Xuanwo/go-locale?status.svg)](https://godoc.org/github.com/Xuanwo/go-locale)
[![Go Report Card](https://goreportcard.com/badge/github.com/Xuanwo/go-locale)](https://goreportcard.com/report/github.com/Xuanwo/go-locale)
[![codecov](https://codecov.io/gh/Xuanwo/go-locale/branch/master/graph/badge.svg)](https://codecov.io/gh/Xuanwo/go-locale)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/Xuanwo/go-locale/blob/master/LICENSE)

`go-locale` is a Golang lib for cross platform locale detection.

## OS Support

### POSIX Compatible Systems

Check order

- Lookup env `LANGUAGE`
- Lookup env `LC_ALL`
- Lookup env `LC_MESSAGES`
- Lookup env `LANG`
- Read file `$XDG_CONFIG_HOME/locale.conf`
- Read file `$HOME/.config/locale.conf`
- Read file `/etc/locale.conf`

Support

- Linux: Ubuntu, CentOS, RHEL, Archlinux...
- [DragonFly BSD](https://www.dragonflybsd.org/)
- [FreeBSD](https://www.freebsd.org/)
- [illumos](https://illumos.org/)
- [NetBSD](https://www.netbsd.org/)
- [OpenBSD](https://www.openbsd.org/)
- [Solaris](https://www.oracle.com/solaris)

### Windows

Check order

- [Win32 OLE](https://docs.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal)

Support

- Windows Vista or Windows Server 2008 Later

### macOS X (darwin)

Check order

- All check for POSIX compatible systems
- macOS X [User Defaults System](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/UserDefaults/AboutPreferenceDomains/AboutPreferenceDomains.html)

Support

- All macOS X version

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

    tags, err := locale.DetectAll()
    if err != nil {
        log.Fatal(err)
    }
    // Get all available tags
}
```

## Acknowledgments

Inspired by [jibber_jabber](https://github.com/cloudfoundry-attic/jibber_jabber)
