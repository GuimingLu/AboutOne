


package model

import (
	"strings"
	"testing"
)

func TestSystemJson(t *testing.T) {
	system := System{Name: "test", Value: NewId()}
	json := system.ToJson()
	result := SystemFromJson(strings.NewReader(json))

	if result.Name != "test" {
		t.Fatal("Ids do not match")
	}
}
