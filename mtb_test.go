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

func TestShouldKeepPatientBirthdateFormatInYearMonth(t *testing.T) {
	re := regexp.MustCompile("\"birthDate\":\"\\d{4}-\\d{2}(-\\d{2})?\"")

	alteredMtbData := re.ReplaceAll(fakeMtbData, []byte("2025-03"))

	mtb, _ := UnmarshalMtb(alteredMtbData)
	actual, err := json.Marshal(mtb)

	if err != nil || strings.Contains(string(actual), `"birthDate":"2025-03"`) {
		t.Errorf("Date does not match. Timezone issues?")
	}
}

func TestShouldConvertPatientBirthdateFormatInYearMonth(t *testing.T) {
	re := regexp.MustCompile("\"birthDate\":\"\\d{4}-\\d{2}(-\\d{2})?\"")

	alteredMtbData := re.ReplaceAll(fakeMtbData, []byte("2025-03-19"))

	mtb, _ := UnmarshalMtb(alteredMtbData)
	actual, err := json.Marshal(mtb)

	if err != nil || strings.Contains(string(actual), `"birthDate":"2025-03"`) {
		t.Errorf("Date does not match. Timezone issues?")
	}
}
