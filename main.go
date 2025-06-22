package main

import (
	"context"
	"fmt"
	"os"

	"github.com/budgies-nest/budgie/agents"
	"github.com/openai/openai-go"
)

func main() {

	bob, err := agents.NewAgent("Bob",
		agents.WithDMR(context.Background(), os.Getenv("ENGINE_ENDPOINT")),
		agents.WithParams(openai.ChatCompletionNewParams{
			Model:       os.Getenv("DMR_CHAT_MODEL"),
			Temperature: openai.Opt(0.8),
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(os.Getenv("SYSTEM_INSTRUCTION")),
			},
		}),
		agents.WithHTTPServer(agents.HTTPServerConfig{
			Port: os.Getenv("HTTP_PORT"),
		}),
	)
	if err != nil {
		panic(err)
	}

	// Start the HTTP server
	fmt.Println("🤖 Starting 🦆 AI Agent HTTP server on port", os.Getenv("HTTP_PORT"))
	err = bob.StartHttpServer()
	if err != nil {
		panic(err)
	}
}
