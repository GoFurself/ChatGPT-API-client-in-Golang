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
		// * Most likely a configuration error; however we let the caller handle this.
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apikey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	//	Timeout: time.Second * 15,
	//}

	resp, err := (client).Do(req)
	if err != nil {
		//* We assume network:ish error, and that the request was not successful.
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// * The request was successful, but the response body could not be read; however we let the caller handle this.
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
