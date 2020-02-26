package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc"
)

func main() {

	log.SetFlags(log.Lshortfile)

	twClient := twitch.NewClient("CodeAfterDark", os.Getenv("TWITCH_OAUTH_SECRET"))

	twClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})

	twClient.Join("CodeAfterDark")

	go func() {
		err := twClient.Connect()
		if err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(os.Stdin)

	for {

		in, err := r.ReadString([]byte("\n")[0])
		if err != nil {
			log.Print(err)
			return
		}

		log.Printf("Read %d characters", len(in))

		twClient.Say("CodeAfterDark", in)

	}

}
