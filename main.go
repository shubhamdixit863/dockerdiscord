package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/shubhamdixit863/discordgo/pkg"
)

var (
	// "OTIzODMzMjA2MzYzNTM3NDI5.YcVwuA.QOyVM2afeV-nGQPGGNsGlQa1ShA"
	Token string
)

func init() {
	godotenv.Load()
	Token = os.Getenv("token")
	fmt.Println(Token)
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(pkg.MessageCreate)

	chron := gocron.NewScheduler(time.Local)
	seconds, err := strconv.ParseUint(os.Getenv("seconds"), 10, 32)
	fmt.Println(seconds)
	if err != nil {
		panic("Seconds Should be a Number")
	}
	chron.Every(10).Seconds().Do(func() {
		fmt.Println("hii")

		pkg.DeleteChannelForCategory(dg, os.Getenv("category1"), os.Getenv("category2"))
	})

	chron.StartAsync()

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
