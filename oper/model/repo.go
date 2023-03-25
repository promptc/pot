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
	Url        string `json:"url"`         // https://promptc.dev
	Db         string `json:"db"`          // https://raw.githubusercontent.com/promptc/repository/main/promptc.dev.db
	UniqueName string `json:"unique_name"` // promptc.dev
}
