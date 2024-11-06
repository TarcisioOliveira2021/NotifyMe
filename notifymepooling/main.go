package main

import (
	"NotifyMe/notifymepooling/structs"
	"NotifyMe/notifymepooling/token"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
)

var previousTotalAlbums int64

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(".env file not found", err.Error())
	}

	artistid := os.Getenv("ARTIST_ID")

	go func() {
		for {
			token, bear := token.GetToken()
			newAlbum := verifyNewAlbum(token, bear, artistid)

			// if newAlbum.TotalTracks > 0{
			// 	notifyMyApp(newAlbum)
			// }

			notifyMyApp(newAlbum)
			time.Sleep(24 * time.Hour)
		}
	}()
	select {}
}

func notifyMyApp(data structs.Item) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(".env file not found", err.Error())
	}
	apiURL := os.Getenv("API_URL")
	jObj, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("JSON ENVIADO:", string(jObj))

	request, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jObj))
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	client.Do(request)
}

func verifyNewAlbum(token string, bear string, artistid string) structs.Item {
	client := http.Client{}
	request := createRequest(token, bear, artistid)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Println(fmt.Errorf("Request failed: %w" + response.Status))
	}

	data, err := convertJSON(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// if data.Total == previousTotalAlbums || previousTotalAlbums == 0 {
	// 	return structs.Item{}
	// }

	previousTotalAlbums = data.Total
	return data.Items[0]
}

func createRequest(token string, bear string, artistid string) *http.Request {
	url := "https://api.spotify.com/v1/artists/" + artistid + "/albums?market=BR"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", bear+" "+token)
	return request
}

func convertJSON(body io.ReadCloser) (structs.Response, error) {
	data := structs.Response{}

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
