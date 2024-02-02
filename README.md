# DiscordGo Bot for Faceit CS:GO Stats

This project is a Discord bot developed in Golang using the DiscordGo library. It's designed to provide asynchronous message handling and fetch statistics from Faceit CS:GO accounts directly within Discord. Users can interact with the bot using various commands, the most notable being the ability to retrieve player stats from Faceit.

## Features

- **Asynchronous Message Handling:** The bot processes messages asynchronously, ensuring efficient communication without blocking.
- **Faceit CS:GO Stats:** Users can fetch their Faceit CS:GO account stats by providing their account details.
- **Help Command:** A `!help` command is available to list all the commands that the bot supports, making it easy for new users to get started.

## Current Limitations

- **Parsing Faceit API:** Due to challenges with the Golang default HTTP library in parsing the Faceit API, the current implementation uses cURL to fetch data. This is a temporary workaround, and future updates aim to integrate a more seamless approach.
- **Direct Messages Only:** Currently, the bot responds exclusively in direct messages. This limitation is on the roadmap for improvement.

## Prerequisites

Before running the project, ensure you have the following installed:
- Go (1.x or later)
- cURL (for the Faceit API data fetching)

## Setup and Running

1. **Clone the Repository:**
    ```
    git clone <repository-url>
    cd <repository-name>
    ```

2. **Environment Configuration:**
   The project uses a `test.env` file to manage environment variables, including the bot token and any other necessary API keys. Ensure you create this file in the root directory and populate it with your tokens and keys.

3. **Install Dependencies:**
    ```
    go mod tidy
    ```

4. **Run the Bot:**
    ```
    go run .
    ```

## Inviting the Bot to Your Server

To invite the bot to your Discord server, use the following link:

[Invite Bot](https://discord.com/api/oauth2/authorize?client_id=1202436890985373768&permissions=8&scope=bot)

Ensure you have the necessary permissions on the Discord server to add bots.

## Contributions

Feedback and contributions are welcome. If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request.

# DiscordBotGolang
