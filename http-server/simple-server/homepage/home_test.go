package homepage

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "good",
			in:             httptest.NewRequest("GET", "/", strings.NewReader("{\n  \"id\": 999,\n  \"value\": \"content\"\n}")),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   "not right",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handlers := NewHandlers(nil)
			handlers.Home(test.out, test.in)
			if test.out.Code != http.StatusOK {
				t.Logf("expected : %d, but : %d \n", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			if body := test.out.Body.String(); body != test.expectedBody {
				t.Logf("expected : %s, but : %s \n", test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
