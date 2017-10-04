package main

import (
	"fmt"
	"log"

	"github.com/google/skylark"
	"encoding/json"
	"bytes"
)

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
