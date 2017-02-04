package models

import (
	"encoding/json"
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
	if root.Definitions["Pet"].Properties["category"].GetReference() != "Category" {
		t.Error()
	}
}
