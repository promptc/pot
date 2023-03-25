package shared

import (
	"os"
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

func GetDbPath(dbname string) string {
	return path.Join(GetUserFolder(), "db", dbname+".db")
}

func GetDbFolder() string {
	return path.Join(GetUserFolder(), "db")
}

func InitPath() {
	err := os.MkdirAll(GetUserFolder(), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	err = os.MkdirAll(GetDbFolder(), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}
