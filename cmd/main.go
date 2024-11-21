package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
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

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"Info message - Ping",
		slog.Group("request", slog.String("method", r.Method), slog.String("url", r.URL.Path), slog.String("ip", r.RemoteAddr)),
	)

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