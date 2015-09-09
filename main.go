package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	"log"
	"os"
)

var unpack bool
var filename string

func init() {
	flag.BoolVar(&unpack, "u", false, "Unpack")
	flag.StringVar(&filename, "f", "", "Read from this file (else, stdin)")
}

func main() {
	flag.Parse()

	var bytes []byte
	var err error

	if filename == "" {
		bytes, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("error reading stdin: %v", err)
		}
	} else {
		bytes, err = ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("error reading %s: %v", filename, err)
		}
	}

	if unpack {
		unpackBytes(bytes)
	} else {
		packBytes(bytes)
	}
}

func packBytes(bytes []byte) {
	var v interface{}

	err := json.Unmarshal(bytes, &v)
	if err != nil {
		log.Fatalf("error unmarshaling json: %v", err)
	}

	var h codec.MsgpackHandle

	buffered := bufio.NewWriter(os.Stdout)
	defer buffered.Flush()

	enc := codec.NewEncoder(buffered, &h)
	err = enc.Encode(v)
	if err != nil {
		log.Fatalf("error encoding: %v", err)
	}
}

func unpackBytes(bytes []byte) {
	var v map[string]interface{}
	var h codec.MsgpackHandle

	dec := codec.NewDecoderBytes(bytes, &h)
	err := dec.Decode(&v)
	if err != nil {
		log.Fatalf("error decoding bytes: %v", err)
	}

	bytes, err = json.Marshal(&v)
	if err != nil {
		log.Fatalf("error decoding json: %v", err)
	}

	fmt.Println(string(bytes))
}
