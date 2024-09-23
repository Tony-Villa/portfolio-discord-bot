package bot

import (
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

type Embed struct {
	Title string
	Description string
	Email string
}

var BotToken string

func checkNilErr(e error) {
	if e != nil {
	 log.Fatal("Error message")
	}
 }

 func Run(e Embed) {
	botChannel := os.Getenv("BOT_CHANNEL_ID")

	color := 0xFF7A5C

	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	discord.Open()

	embed := &discordgo.MessageEmbed{
    Author:      &discordgo.MessageEmbedAuthor{
			Name: e.Email,
		},
    Color:       color, // Green
    Description: e.Description,
    Fields: []*discordgo.MessageEmbedField{},
    Image: &discordgo.MessageEmbedImage{},
    Thumbnail: &discordgo.MessageEmbedThumbnail{},
    Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
    Title:     e.Title,
	}

	discord.ChannelMessageSendEmbed(botChannel, embed)

	// discord.ChannelMessage(botChannel, "howdy")

	discord.Close() // close session, after function termination


 }