


package model

import (
	"strings"
	"testing"
)

func TestCompliance(t *testing.T) {
	o := Compliance{Desc: "test", CreateAt: GetMillis()}
	json := o.ToJson()
	result := ComplianceFromJson(strings.NewReader(json))

	if o.Desc != result.Desc {
		t.Fatal("JobName do not match")
	}
}
