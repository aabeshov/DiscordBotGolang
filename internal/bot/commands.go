package bot

import (
	"DiscordBot/internal/api"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"strings"
)

func (b *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by the bot itself to avoid a loop
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message starts with "!"
	if strings.HasPrefix(m.Content, "!") {
		// Handle the command
		go b.processCommand(s, m)
	}
}

func (b *Bot) processCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	commandParts := strings.Fields(m.Content)
	command := commandParts[0]
	command = strings.TrimPrefix(command, "!")
	switch command {
	// TODO More commands
	case "help":
		b.sendHelpMessage(s, m.ChannelID)
	case "stats":
		if len(commandParts) < 2 {
			s.ChannelMessageSend(m.ChannelID, "Please provide a user ID. Usage: !stats <user_id>")
			return
		}
		nickname := commandParts[1]
		fmt.Println(nickname)
		b.fetchAndSendUserStats(s, m.ChannelID, nickname)

	default:
		s.ChannelMessageSend(m.ChannelID, "Unknown command.Use !help to get all available commands.")
	}
}

func (b *Bot) sendHelpMessage(s *discordgo.Session, channelID string) {
	msg := "Available commands:\n"
	msg += "!help - Show this message\n"
	msg += "!stats - Show player statistics\n"
	// TODO More commands
	s.ChannelMessageSend(channelID, msg)
}

func (b *Bot) fetchAndSendUserStats(s *discordgo.Session, channelID, nickname string) {
	stats, err := faceit_handle_request(nickname)
	if err != nil {
		s.ChannelMessageSend(channelID, "12Failed to parse JSON response: "+err.Error())
		return
	}

	s.ChannelMessageSend(channelID, stats)

}

func faceit_handle_request(nickname string) (string, error) {

	err := godotenv.Load("test.env")
	if err != nil {
		log.Fatal("Error loading test.env file")
	}
	apiKey := os.Getenv("FACEIT_API_KEY")

	curlCmd := fmt.Sprintf(`curl --location 'https://open.faceit.com/data/v4/players?nickname=%s&game=csgo' \
	--header 'Authorization:Bearer %s' \`, nickname, apiKey)
	//fmt.Println(curlCmd)

	cmd := exec.Command("bash", "-c", curlCmd)

	output, _ := cmd.CombinedOutput()
	outputStr := string(output)
	// Parsing (kostil')
	start := strings.Index(outputStr, "{")
	end := strings.LastIndex(outputStr, "}")
	//fmt.Println(start, end)
	if start != -1 && end != -1 {
		jsonObject := outputStr[start : end+1]
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(jsonObject), &data); err != nil {
			return "Error parsing JSON", nil
		}

		player_id := data["player_id"]
		curlCmd = fmt.Sprintf(`curl --location 'https://open.faceit.com/data/v4/players/%s/stats/csgo' \
		--header 'Authorization:Bearer %s'`, player_id, apiKey)
		cmd = exec.Command("bash", "-c", curlCmd)

		output, _ = cmd.CombinedOutput()
		outputStr = string(output)
		// Parsing (kostil')
		start = strings.Index(outputStr, "{")
		end = strings.LastIndex(outputStr, "}")
		if start != -1 && end != -1 {
			fmt.Println(1)
			jsonObject = outputStr[start : end+1]
			var playerStats api.PlayerStatsResponse
			if err := json.Unmarshal([]byte(jsonObject), &playerStats); err != nil {
				return "Error parsing JSON", nil
			}
			respond := fmt.Sprintf("You won %s", playerStats.Lifetime.Wins)
			return respond, nil
		}

	} else {
		return "JSON object not found in the response", nil
	}
	return "Nothing was found", nil

}
