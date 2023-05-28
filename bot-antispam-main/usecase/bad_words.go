package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TextPayload struct {
	Text string `json:"text"`
}

func SpamWords(msg string) (res bool, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("get error in making request: %v", err)
		}
	}()
	payload := TextPayload{
		Text: msg,
	}
	jsonPayload, err := json.Marshal(payload)
	url := "http://localhost:5000/evaluate_text"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	var result []int
	err = json.NewDecoder(resp.Body).Decode(&result)

	return result[0] != 0, err
}

func ShowWords(spams []string) string {
	var spam string
	for _, w := range spams {
		spam = spam + w + ", "
	}
	return spam
}
