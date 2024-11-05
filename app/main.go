package main

import (
	"NotifyMe/notifymepooling/structs"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.POST("/notification", generateEmailNotification)
	route.Run()
}

func generateEmailNotification(c *gin.Context) {
	//Leitura do body
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer c.Request.Body.Close()
	data := structs.Item{}

	//Transformação em json
	err = json.Unmarshal(requestBody, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Geração do email
	message := gomail.NewMessage()
	message.SetHeader("From", "notifymeBot@sandbox.com")
	message.SetHeader("To", "tarcisio.zark.veloso@gmail.com")
	message.SetHeader("Subject", "New album from "+data.Artists[0].Name+" avaialble in Spotify !")

	albumImgUrl := data.Images[1].URL //0=> 640x640 1=>300x300 2=>64x64
	albumName := data.Name
	albumUrl := data.ExternalUrl.Spotify
	tracksNumber := strconv.FormatInt(data.TotalTracks, 10)

	listItems := ""
	for _, artist := range data.Artists {
		listItems += fmt.Sprintf(`<p>• <a href="%s"> <strong>%s</strong> </a> - Type: %s</p>`, artist.ExternalUrls.Spotify, artist.Name, strings.ToUpper(artist.Type))
	}

	htmlString := fmt.Sprintf(`
  				<div class="">
    				<img src="%s" alt="albumImage300x300">
    				<div class="">
      					<h2><a href="%s">%s</a></h2>
      					
						<p class="text">Number of tracks: %s</p>
						%s
    				</div>
  				</div>
	`, albumImgUrl, albumUrl, albumName, tracksNumber, listItems)

	message.SetBody("text/html", htmlString)

	sender := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, "6bf7de91759ec9", "cf8c990571cc51")
	sender.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	//Envio
	err = sender.DialAndSend(message)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Email, sended!")
}
