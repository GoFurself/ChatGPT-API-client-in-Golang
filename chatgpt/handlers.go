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
		// * Most likely a configuration error, thus panic.
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+apikey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		//* We assume network:ish error, and that the request was not successful.
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// * The request was successful, but the response body could not be read.
		// * Thus Panic seems to be the only way to handle this..
		return nil, err
	}
	if resp.StatusCode != 200 {
		// * The request was successful, but the response status code was not 200.
		return nil, errors.New("We wanted a '200 OK'. However, the server returned status code: " + resp.Status + "\nAnd body:\n" + string(body) + "\n)")
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
