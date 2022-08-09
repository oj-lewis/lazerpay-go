package lazerpay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)


func ToJSON(data interface{}) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return json
}

func GetRequest( url string, secretKey string) ( map[string]interface{},  error) { 
		client := &http.Client{
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return  nil, err
		}

		req.Header.Set("Authorization", "Bearer " + secretKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		var body map[string]interface{}
		x, err:= io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(x, &body)

		return body, nil


}


func PostRequest(payload interface{}, url string, secretKey string) (map[string]interface{}, error) {
		data := ToJSON(payload)
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			return  nil, err
		}

		req.Header.Add("Authorization", "Bearer " + secretKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		var body map[string]interface{}
		x, err:= io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(x, &body)

		return body, nil

}