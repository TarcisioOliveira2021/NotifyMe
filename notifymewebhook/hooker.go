package notifymewebhook

import (
	"NotifyMe/notifymewebhook/structs"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)


func GetToken() (string, string, error) {
	request, err := createRequest()
	if err != nil {
		panic(err)
	}

	response, err := requestApiTokenSpotify(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", "", fmt.Errorf("Request failed:" + response.Status)
	}

	data, err := convertJSON(response.Body)
	if err != nil {
		return "","", err	}

	return data.AccessToken, data.TokenType, nil
}

func convertJSON(body io.ReadCloser) (structs.Resp, error) {
	data := structs.Resp{}

	read, err := io.ReadAll(body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(read, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func getClientVariables() (string, string, error) {
	// gotoenv starta junto ao main.
	err := godotenv.Load("../notifymewebhook/.env")
	if err != nil {
		return "", "", fmt.Errorf(".env file not found. %w", err)
	}

	client_id := os.Getenv("CLIENT_ID")
	client_secret := os.Getenv("CLIENT_SECRET")

	return client_id, client_secret, nil
}

func createRequest() (*http.Request, error) {
	client_id, client_secret, err := getClientVariables()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	bodyData := url.Values{}
	bodyData.Set("grant_type", "client_credentials")
	encodedBodyData := bodyData.Encode()
	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(encodedBodyData))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	request.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(client_id+":"+client_secret)))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return request, nil
}

func requestApiTokenSpotify(request *http.Request) (*http.Response, error) {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return response, nil
}