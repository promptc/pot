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
	var fetchedDb []string
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
		file, err := os.Create(shared.GetDbPath(repoInfo.UniqueName + "-new"))
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
		fetchedDb = append(fetchedDb, repoInfo.UniqueName)

	}
	fmt.Println("- Fetched", len(fetchedDb), "source(s)")
	fmt.Println("- Updating database...")
	for _, db := range fetchedDb {
		_ = os.Remove(shared.GetDbPath(db))
		err := os.Rename(shared.GetDbPath(db+"-new"), shared.GetDbPath(db))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("- Database updated")
}
