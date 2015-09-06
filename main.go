package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	"log"
	"os"
)

var unpack bool

func init() {
	flag.BoolVar(&unpack, "u", false, "Unpack")
}

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading stdin: %v", err)
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
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	var h codec.MsgpackHandle

	enc := codec.NewEncoder(os.Stdout, &h)
	err = enc.Encode(v)
	if err != nil {
		log.Fatalf("Error encoding: %v", err)
	}
}

func unpackBytes(bytes []byte) {
	var v map[string]interface{}
	var h codec.MsgpackHandle

	dec := codec.NewDecoderBytes(bytes, &h)
	err := dec.Decode(&v)
	if err != nil {
		log.Fatalf("Error decoding stream: %v", err)
	}

	bytes, err = json.Marshal(&v)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	fmt.Println(string(bytes))
}
