package model

import "time"

type RepoModel []PromptInfo

type PromptInfo struct {
	Name        string
	Path        string
	Author      []string
	Version     string
	Description string
	CreatedAt   time.Time
	LastUpdated time.Time
}

type RepoInfo struct {
	Prompt string `json:"prompt"` // https://promptc.dev/prompts
	Db     string `json:"db"`     // https://raw.githubusercontent.com/promptc/repository/main/promptc.dev.db
	Id     string `json:"id"`     // promptc.dev
}
