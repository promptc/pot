package remote

import (
	"encoding/json"
	"fmt"
	"github.com/promptc/pot/oper/cfg"
	"github.com/promptc/pot/oper/model"
	"github.com/promptc/pot/oper/shared"
	"io"
	"net/http"
	"os"
)

func UpdateHandler(args []string) {
	conf := cfg.Get()
	fetchedDb := make(map[string]string, 0)
	if len(conf.Source) == 0 {
		fmt.Println("No source found!") // TODO: warning!
		return
	}
	for _, source := range conf.Source {
		fmt.Println("- Fetching info from", source+"info.json")
		resp, err := http.Get(source + "info.json")
		if err != nil {
			fmt.Println("--- Error:", err)
			fmt.Println("--- Warn:", "Ignore")
			continue
		}
		fmt.Println("--- Status:", resp.Status)
		if resp.StatusCode != http.StatusOK {
			fmt.Println("--- Warn:", "Not 200, Ignore")
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body) // FIXME: multi-decoder: gzip, deflate, br
		if err != nil {
			fmt.Println("--- Error:", err)
			fmt.Println("--- Warn:", "Ignore")
			continue
		}
		var repoInfo model.RepoInfo
		err = json.Unmarshal(body, &repoInfo)
		if err != nil {
			fmt.Println("--- Error:", err)
			fmt.Println("--- Warn:", "Ignore")
			continue
		}
		fmt.Println("--- Fetching db from", repoInfo.Db)
		resp, err = http.Get(repoInfo.Db)
		if err != nil {
			fmt.Println("--- Error:", err)
			fmt.Println("--- Warn:", "Ignore")
			continue
		}
		fmt.Println("--- Status:", resp.Status)
		if resp.StatusCode != http.StatusOK {
			fmt.Println("--- Warn:", "Not 200, Ignore")
			continue
		}
		defer resp.Body.Close()
		file, err := os.Create(shared.GetDbPath(repoInfo.Id + "-new"))
		if err != nil {
			fmt.Println("--- Error:", err)
			fmt.Println("--- Warn:", "Ignore")
			continue
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Println("--- Error:", err)
			fmt.Println("--- Warn:", "Ignore")
			continue
		}
		fetchedDb[repoInfo.Id] = repoInfo.Prompt

	}
	fmt.Println("- Fetched", len(fetchedDb), "source(s)")
	fmt.Println("- Updating database...")
	for db, _ := range fetchedDb {
		_ = os.Remove(shared.GetDbPath(db))
		err := os.Rename(shared.GetDbPath(db+"-new"), shared.GetDbPath(db))
		if err != nil {
			panic(err)
		}
	}
	updatePromptTarget(fetchedDb)
	fmt.Println("- Database updated")
}

func updatePromptTarget(fetchedDb map[string]string) {
	var db map[string]string
	dbInfoPath := shared.DbInfo.ToPath()
	if shared.FileExists(dbInfoPath) {
		jFile, err := os.Open(dbInfoPath)
		if err != nil {
			panic(err)
		}
		bs, err := io.ReadAll(jFile)
		jFile.Close()
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bs, &db)
		if err != nil {
			panic(err)
		}
	}
	if db == nil || len(db) == 0 {
		db = make(map[string]string, 0)
	}
	for id, prompt := range fetchedDb {
		db[id] = prompt
	}
	bs, err := json.Marshal(db)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(dbInfoPath, bs, 0644)
	if err != nil {
		panic(err)
	}
}
