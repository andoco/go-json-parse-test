package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/buger/jsonparser"
)

var (
	objectPayload = `{"payload": {"foo": "bar"}}`
	arrayPayload  = `{"payload": ["foo", "bar"]}`
)

func main() {
	parse(objectPayload)
	parse(arrayPayload)

	extract(objectPayload, "payload")
	extract(arrayPayload, "payload")
}

func parse(jsonValue string) {
	fmt.Println("\nPARSING")
	reader := strings.NewReader(jsonValue)
	dec := json.NewDecoder(reader)

	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
}

func extract(jsonValue string, field string) {
	fmt.Println("\nEXTRACTING")
	val, t, offset, err := jsonparser.Get([]byte(jsonValue), field)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n%v\n%v\n", val, t, offset)
}
