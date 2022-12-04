package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"tipo-server/app/models"
)

func FetchCheckTypo(input string) (result *string, err error) {
	url := os.Getenv("TIPO_URL") + "/api/v1/checkTypo"
	method := "POST"

	payload := strings.NewReader(`{
		"text": "` + input + `"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Printf("FetchCheckTypo::http.NewRequest::%v", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer xxx")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("FetchCheckTypo::client.Do::%v", err)
		return nil, err
	}

	defer res.Body.Close()

	var data models.TipoReturn
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Printf("FetchCheckTypo::json.NewDecoder::%v", err)
		return nil, err
	}

	if !data.IsSuccess {
		fmt.Printf("FetchCheckTypo::data::%v", data.Message)
		return nil, errors.New(data.Message)
	}

	return &data.Result, nil
}
