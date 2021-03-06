// Copyright 2016 HeadwindFly. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clevergo

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"time"
)

const (
	Version = "1.0.2"
	Logo    = `  ____ _     _______     _______ ____   ____  ___
 / ___| |   | ____\ \   / / ____|  _ \ / ___|/ _ \
| |   | |   |  _|  \ \ / /|  _| | |_) | |  _| | | |
| |___| |___| |___  \ V / | |___|  _ <| |_| | |_| |
 \____|_____|_____|  \_/  |_____|_| \_\\____|\___/ `
)

func info() {
	fmt.Printf("\x1b[36;1m%s %s\x1b[0m\n\n\x1b[32;1mStarted at %s\x1b[0m\n", Logo, Version, time.Now())
}

func ListenAndServe(addr string, handler fasthttp.RequestHandler) error {
	info()
	return fasthttp.ListenAndServe(addr, handler)
}

func ListenAndServeUNIX(addr string, mode os.FileMode, handler fasthttp.RequestHandler) error {
	info()
	return fasthttp.ListenAndServeUNIX(addr, mode, handler)
}

func ListenAndServeTLS(addr, certFile, keyFile string, handler fasthttp.RequestHandler) error {
	info()
	return fasthttp.ListenAndServeTLS(addr, certFile, keyFile, handler)
}

func ListenAndServeTLSEmbed(addr string, certData, keyData []byte, handler fasthttp.RequestHandler) error {
	info()
	return fasthttp.ListenAndServeTLSEmbed(addr, certData, keyData, handler)
}
