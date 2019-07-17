package suggestions

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"strings"
	"time"
)

const placeHolderText = "+"

// NameExtensions struct defines parsed extensions structure
type NameExtensions struct {
	Extensions []string `json:"extensions"`
}

func handleErrors(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// LoadNameExtensions parses/loads predefined name extensions
func LoadNameExtensions(r io.Reader) (e NameExtensions) {
	// Alternative #1
	d := json.NewDecoder(r)
	if err := d.Decode(&e); err != nil {
		handleErrors(err)
	}

	// Alternative #2
	// b, _ := ioutil.ReadAll(r)
	// json.Unmarshal(b, &e)

	return e
}

// Suggestions returns and random extension suggestion
func Suggestions(s string, e NameExtensions) []string {
	rand.Seed(time.Now().UTC().UnixNano())

	results := 10
	var r []string
	for i := 0; i < results; i++ {
		ext := e.Extensions[rand.Intn(len(e.Extensions))]
		r = append(r, strings.Replace(ext, placeHolderText, s, -1))
	}

	return r
}
