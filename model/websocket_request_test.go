


package model

import (
	"strings"
	"testing"
)

func TestWebSocketRequest(t *testing.T) {
	m := WebSocketRequest{Seq: 1, Action: "test"}
	json := m.ToJson()
	result := WebSocketRequestFromJson(strings.NewReader(json))

	if result == nil {
		t.Fatal("should not be nil")
	}

	badresult := WebSocketRequestFromJson(strings.NewReader("junk"))

	if badresult != nil {
		t.Fatal("should have been nil")
	}
}
