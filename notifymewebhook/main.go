package main

import (
	"NotifyMe/notifymewebhook/structs"
	"NotifyMe/notifymewebhook/token"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/newalbumavailable/:artistid", start)
	route.Run(":8080")
}

func start(c *gin.Context) {
	token, bear := token.GetToken()
	artistid, err := c.Params.Get("artistid")

	if !err {
		c.AbortWithStatus(404)
	}

	verifyNewAlbum(token, bear, artistid)
	c.String(http.StatusOK, "")
}

// http://localhost:8080/newalbumavailable/0Riv2KnFcLZA3JSVryRg4y
// https://developer.spotify.com/documentation/web-api/reference/get-an-artists-albums
func verifyNewAlbum(token string, bear string, artistid string) {
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

	fmt.Println(data)
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
