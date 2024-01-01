package chatgpt

// * Contains implementations for interfaces defined in types.go * //

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func HTTPRequestHandler(url string, apikey string, requestModel []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestModel))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apikey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Request failed with status code: " + resp.Status + "\nAnd body:\n" + string(body) + "\n)")
	}

	return body, nil
}

type JsonMarshalHandler struct{}

func (jt *JsonMarshalHandler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func (jt *JsonMarshalHandler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
