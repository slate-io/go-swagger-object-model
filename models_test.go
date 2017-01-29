package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSwaggerObjectModel(t *testing.T) {
	b, err := ioutil.ReadFile("./fixtures/swagger.json")
	if err != nil {
		t.Error(err)
	}
	var root = new(RootDocument)
	if err := json.Unmarshal(b, root); err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\r\n", root)
}
