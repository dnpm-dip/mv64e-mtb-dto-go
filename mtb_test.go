package mtb

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"
)

//go:embed tests/mv64e-mtb-fake-patient.json
var fakeMtbData []byte

func TestShouldDeserializeJson(t *testing.T) {
	_, err := UnmarshalMtb(fakeMtbData)

	if err != nil {
		t.Errorf("Cannot deserialize MTB file")
	}
}

func TestShouldKeepTimezone(t *testing.T) {
	mtb, _ := UnmarshalMtb(fakeMtbData)
	actual, err := json.Marshal(mtb)
	if err != nil || !strings.Contains(string(actual), "\"birthDate\":\"1985-05-19\"") {
		t.Errorf("Date does not match. Timezone issues?")
	}
}
