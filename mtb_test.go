package mtb

import (
	_ "embed"
	"encoding/json"
	"regexp"
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
	re := regexp.MustCompile("\"birthDate\":\"\\d{4}-\\d{2}-\\d{2}\"")

	expectedDate := re.FindAllStringSubmatch(string(fakeMtbData), -1)[0][0]

	mtb, _ := UnmarshalMtb(fakeMtbData)
	actual, err := json.Marshal(mtb)

	if err != nil || !strings.Contains(string(actual), expectedDate) {
		t.Errorf("Date does not match. Timezone issues?")
	}
}
