// Copyright 2017 The Bazel Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The skylark command interprets a Skylark file.
// With no arguments, it starts a read-eval-print loop (REPL).
package main

import (
	"fmt"
	"log"

	"bytes"
	"encoding/json"

	"github.com/google/skylark"
	"github.com/kataras/go-errors"
	"os"
)

func init() {
	skylark.Universe["getenv"] = skylark.NewBuiltin("getnev", getenv)
}

func getenv(thread *skylark.Thread, fn *skylark.Builtin, args skylark.Tuple, kwargs []skylark.Tuple) (skylark.Value, error) {
	if len(args) != 1 {
		return skylark.None, errors.New("a lot of values")
	}

	key, ok := skylark.AsString(args[0])
	if !ok {
		return skylark.None, errors.New("not mathc type")
	}

	env := os.Getenv(key)
	return skylark.String(env), nil
}

func main() {
	thread := new(skylark.Thread)
	globals := make(skylark.StringDict)
	if err := skylark.ExecFile(thread, "test.sky", nil, globals); err != nil {
		log.Fatalln(err)
	}

	for key, v := range globals {
		fmt.Println(key)

		buffer := new(bytes.Buffer)
		if err := json.Compact(buffer, []byte(v.String())); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("\t", buffer.String())
	}
}
