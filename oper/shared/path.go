package shared

import (
	"path"
	"runtime"
)

func getSysConfigBase() string {
	if runtime.GOOS == "windows" {
		return "%APPDATA%"
	}
	return "$HOME/.config"
}

func GetUserFolder() string {
	return path.Join(getSysConfigBase(), "prompt", "pot")
}

type PotPath string

const (
	ConfigPath PotPath = "config.json"
)

func (p PotPath) ToPath() string {
	return path.Join(GetUserFolder(), string(p))
}

func GetPath(potPath PotPath) string {
	return path.Join(GetUserFolder(), string(potPath))
}
