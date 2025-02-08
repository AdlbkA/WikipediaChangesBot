package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type WikiEvent struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	User      string `json:"user"`
	Timestamp int64  `json:"timestamp"`
	Wiki      string `json:"wiki"`
	Comment   string `json:"comment"`
	URL       string `json:"title_url"`
}

const prefix string = "!"

var (
	defaultLang = "en"
	streaming   = false
	url         = "https://stream.wikimedia.org/v2/stream/recentchange"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Bot is running!")
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if strings.HasPrefix(m.Content, prefix) {
			command := strings.TrimPrefix(m.Content, prefix)
			args := strings.Split(command, " ")

			if args[0] == "recent" {
				streaming = true
				fetchData(s, m.ChannelID)
			} else if args[0] == "setLang" {
				defaultLang = args[1]
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Set language to '%s'", defaultLang))
			}
			if args[0] == "stop" {
				streaming = false
				s.ChannelMessageSend(m.ChannelID, "Goodbye!")
			}
		}

	})

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func fetchData(s *discordgo.Session, channelID string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if !streaming {
			return
		} else {
			line := scanner.Text()
			if len(line) > 0 && line[:1] == ":" {
				continue
			}

			if len(line) > 6 && line[:5] == "data:" {
				jsonData := line[6:]

				var event WikiEvent
				err := json.Unmarshal([]byte(jsonData), &event)
				if err != nil {
					fmt.Println("Error parsing JSON:", err)
					continue
				}

				timestamp := event.Timestamp
				tzone := time.FixedZone("Almaty", 5*60*60)

				t := time.Unix(timestamp, 0).In(tzone)
				if event.Wiki == defaultLang+"wiki" {
					var message = fmt.Sprintf("Edit on '%s'\n Time: %s by %s\n Comment: %s\nLink: %s", event.Title, t.Format("2006-01-02 15:04:15"), event.User, event.Comment, event.URL)
					s.ChannelMessageSend(channelID, message)
				}
			}
		}
	}
}
