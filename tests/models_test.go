package tests

import (
	"jumia/models"
	"testing"
)

func TestRegexValidation(t *testing.T) {

	tests := []struct {
		regex, phone  string
		valueExpected bool
	}{
		{"\\(237\\) ?[2368]\\d{7,8}$", "(237) 699209115", true},
		{"\\(251\\) ?[1-59]\\d{8}$", "(251) 911203317", true},
		{"\\(212\\) ?[5-9]\\d{8}$", "(212) 6007989253", false},
		{"\\(258\\) ?[28]\\d{7,8}$", "(258) 847651504", true},
		{"\\(256\\) ?\\d{9}$", "(256) 7503O6263", false},
	}

	for _, teste := range tests {
		value := models.RegexValidation(teste.regex, teste.phone)

		if value != teste.valueExpected {
			t.Errorf("the phone number %v is not valid", teste.phone)
		}
	}

}
