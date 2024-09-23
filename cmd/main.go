package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"portfolio-discord-bot/bot"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/formComplete", portfolioContact)


	fmt.Println("Listening on port 3005, logs")
	http.ListenAndServe(":3005", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func portfolioContact(w http.ResponseWriter, r *http.Request) {
	var embed bot.Embed

	err := json.NewDecoder(r.Body).Decode(&embed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("uhhhh"))
	
	w.WriteHeader(http.StatusOK)


	bot.BotToken = os.Getenv("DISCORD_TOKEN")
	bot.Run(embed);
	
}