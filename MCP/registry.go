package main

import (
	"github.com/slack-ai-plugin/mcp-server/config"
	"github.com/slack-ai-plugin/mcp-server/models"
	tools_ai_alpha_search_messages "github.com/slack-ai-plugin/mcp-server/tools/ai_alpha_search_messages"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_ai_alpha_search_messages.CreateAi_alpha_search_messagesTool(cfg),
	}
}
