package main

import (
	"NotifyMe/notifymewebhook"
	"time"
	"fmt"
)

func main() {

	go func() {
        for {
            token, bear, err := notifymewebhook.GetToken()
            if err != nil {
                fmt.Printf("Erro: %v\n", err)
                return 
            }

            fmt.Println("Token gerado:", token)
            fmt.Println("Bear:", bear)

            time.Sleep(3600 * time.Second)
        }
    }()

    select {}
}
