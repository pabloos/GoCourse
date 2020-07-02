package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pabloos/http/greet"
)

func Test_Greet(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name   string
		Greet  greet.Greet
		Wanted string
	}{
		{
			Name: "Green case",
			Greet: greet.Greet{
				Name:     "John Doe",
				Location: "NY",
			},
			Wanted: "Hello John Doe, from NY\n",
		},
		{
			Name: "void case",
			Greet: greet.Greet{
				Name:     "",
				Location: "",
			},
			Wanted: "Tell us what is your name and where do you come from!\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(tc.Greet)
			if err != nil {
				t.Fatalf("Problem while encoding JSON: " + err.Error())
			}

			req, err := http.NewRequest("POST", "localhost:8080/greet", buf)
			if err != nil {
				t.Fatalf("Couldn't create request: %v", err)
			}

			rec := httptest.NewRecorder()

			greetHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			msgBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Couldn't read response: %v", err)
			}

			msg := string(msgBytes)
			if msg != tc.Wanted {
				t.Errorf("%s failed: wanted %s, get %s", tc.Name, tc.Wanted, msg)
			}
		})
	}
}
