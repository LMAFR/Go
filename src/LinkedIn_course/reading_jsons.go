package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// In this script, we are going to read the name of several countries stored in the archive lista.json
// In order to do it, we are going to define  the structure of the JSON:
type Pais struct {
	NOMBRE    string
	POBLACION float32
}

// The JSON has more than 2 variables defined (per country), but we won't have to take them all.
// Now, we define a second structure (Paises) that will be a JSON in which each element has the
// structure of the type "Pais" (similar to de archive we want to open, but with less variables):
type Paises struct {
	Paises []Pais `json:"paises"`
}

func main() {

	data, err := os.Open("lista.json")

	// If we get an error, we print it:
	if err != nil {
		fmt.Print(err)
	}

	// ReadAll will read all the information stored in the JSON archive.
	byteValue, _ := ioutil.ReadAll(data)

	// Let's define a variable paises in which we will use a structure similar to the structure of the JSON we have read:
	var paises Paises

	// We take all the information of the JSON (stored in byteValue)
	// and we store it as a JSON in the variable "paises" (whose type is Paises).
	// Any possible output would be because of the presence of errors (so we call it "err").
	err = json.Unmarshal(byteValue, &paises)

	// If we get an error, we print it:
	if err != nil {
		fmt.Println("error: ", err)
	}

	// For each element in our variable "paises", in the structure regarding "Paises", we are going to iterate to show each country name:
	for i := 0; i < len(paises.Paises); i++ {

		println("* " + paises.Paises[i].NOMBRE)

	}

}
