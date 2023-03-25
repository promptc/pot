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
