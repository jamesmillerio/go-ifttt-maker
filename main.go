package GoIFTTTMaker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//MakerChannel assists in the sending of POST requests to the IFTTT Maker Channel.
type MakerChannel struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
}

//Send a request to the IFTTT Maker Channel.
func (maker *MakerChannel) Send(key string, event string) bool {

	api := fmt.Sprintf("https://maker.ifttt.com/trigger/%v/with/key/%v", event, key)
	body, err := json.Marshal(maker)

	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest("POST", api, buffer)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200
}
