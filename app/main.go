package main

import (
	"NotifyMe/notifymepooling/structs"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	route := gin.Default()
	route.POST("/notification", generateEmailNotification)
	route.Run()
}

func generateEmailNotification(c *gin.Context) {

	emailTo, emailHost, emailPort, emailUsername, emailPassword, emailFrom := loadEnviromentVariables()
	data := deserializeRequestJSON(c)
	emailMessage := createEmailMessage(data, emailTo, emailFrom)
	emailPortConv := covertStringToInt(emailPort)
	emailSender := createEmailSender(emailHost, emailPortConv, emailUsername, emailPassword)

	sendEmail(emailMessage, emailSender)
}

func deserializeRequestJSON(c *gin.Context) structs.Item {
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer c.Request.Body.Close()
	data := structs.Item{}

	err = json.Unmarshal(requestBody, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return data
}

func createEmailMessage(data structs.Item, emailTo string, emailFrom string) *gomail.Message {

	message := gomail.NewMessage()
	message.SetHeader("To", emailTo)
	message.SetHeader("From", emailFrom)
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

	return message
}

func createEmailSender(emailHost string, emailPort int, emailUsername string, emailPassword string) *gomail.Dialer {
	sender := gomail.NewDialer(emailHost, emailPort, emailUsername, emailPassword)
	return sender
}

func loadEnviromentVariables() (string, string, string, string, string, string) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(".env file not found", err.Error())
	}

	emailTo := os.Getenv("EMAIL_TO")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailHost := os.Getenv("EMAIL_HOST")
	emailPort := os.Getenv("EMAIL_PORT")
	emailUsername := os.Getenv("EMAIL_USERNAME")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	return emailTo, emailHost, emailPort, emailUsername, emailPassword, emailFrom
}

func covertStringToInt(port string) int {
	portConv, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	return portConv
}

func sendEmail(message *gomail.Message, dialer *gomail.Dialer) {
	err := dialer.DialAndSend(message)
	if err != nil {
		panic(err)
	}
}
