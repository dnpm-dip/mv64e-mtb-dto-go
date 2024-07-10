package mtb

import (
	_ "embed"
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
