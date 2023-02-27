package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Jokes struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func Botops(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!random" {
		req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Accept", "text/plain")
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		s.ChannelMessageSend(m.ChannelID, string(body))
	}

}

func Main() {
	godotenv.Load(".env")
	sess, err := discordgo.New("Bot " + "MTA3OTA1MTYyMjMyMTIzMzk3MA.GR9QUX.jx4mTWF9GS3dEdAS_0zacca6Bu4yrTBvhNnHZU")
	if err != nil {
		panic(err)
	}

	sess.AddHandler(Botops)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open()
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	fmt.Println("bot is running")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
