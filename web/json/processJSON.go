package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON has become the most popular medium used by processes
// and systems to exchange data and information
// Note that only the members of a structure that begin with an
// uppercase letter will be in the JSON output because members
// that begin with a lowercase letter are considered private
type record struct {
	Name    string
	Surname string
	Tel     []telephone
}

type telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename string, key interface{}) {
	out, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// creates an encoder that writes to the created file
	jsonEncoder := json.NewEncoder(out)
	err = jsonEncoder.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// close the opened file to free memory space and avoid memory leak
	out.Close()
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	jsonDecoder := json.NewDecoder(in)
	err = jsonDecoder.Decode(key)
	if err != nil {
		return err
	}

	in.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please a provide a filename to write to..")
		os.Exit(1)
	}

	filename := arguments[1]

	newRecord := record{
		Name:    "Olumide",
		Surname: "Ogundele",
		Tel: []telephone{
			telephone{Mobile: true, Number: "1234-567"},
			telephone{Mobile: true, Number: "1234-abcd"},
			telephone{Mobile: false, Number: "abcc-567"},
		},
	}

	saveToJSON(filename, newRecord)

	if len(arguments) == 2 {
		fmt.Println("Please a provide a file to read from..")
		os.Exit(1)
	}

	var myRecord record

	err := loadFromJSON(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myRecord)
	}

	// using marshalling(encode) convert Go data type to JSON and
	// unmarshalling(decode) convert JSON to Go data type, more like it writes
	// to the data type
	recordJSON, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(recordJSON))

	var unRec record
	err = json.Unmarshal(recordJSON, &unRec)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(unRec)
}
