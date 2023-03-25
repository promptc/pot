package cfg

import (
	"encoding/json"
	"github.com/promptc/pot/oper/shared"
	"io"
	"os"
)

func createPath() {
	err := os.MkdirAll(shared.ConfigPath.ToPath(), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

var cfg *Model

func Get() *Model {
	if cfg != nil {
		return cfg
	}
	createPath()
	cfgFile := shared.ConfigPath.ToPath()
	if shared.FileExists(cfgFile) {
		jFile, err := os.Open(cfgFile)
		if err != nil {
			panic(err)
		}
		defer jFile.Close()
		bs, err := io.ReadAll(jFile)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bs, cfg)
		if err != nil {
			panic(err)
		}
	}
	cfg = defaultModel()
	Save()
	return cfg
}

func Save() {
	bs, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(shared.ConfigPath.ToPath(), bs, 0644)
	if err != nil {
		panic(err)
	}
}
