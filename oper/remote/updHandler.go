package remote

import (
	"encoding/json"
	"fmt"
	"github.com/promptc/pot/oper/cfg"
	"github.com/promptc/pot/oper/model"
	"github.com/promptc/pot/oper/shared"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
)

func UpdateHandler(args []string) {
	conf := cfg.Get()
	sourceResults := make(map[string]model.RepoModel)
	if len(conf.Source) == 0 {
		fmt.Println("No source found!") // TODO: warning!
		return
	}
	for _, source := range conf.Source {
		fmt.Println("- Fetching from", source)
		resp, err := http.Get(source)
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
		var repoModel model.RepoModel
		err = json.Unmarshal(body, &repoModel)
		if err != nil {
			fmt.Println("--- Error:", err)
			fmt.Println("--- Warn:", "Ignore")
			continue
		}
		sourceResults[source] = repoModel
	}
	fmt.Println("- Fetched", len(sourceResults), "source(s)")
	fmt.Println("- Updating database...")

	model.NewSqlite(shared.GetDbPath("new"))
	updateToDb("new", sourceResults)
	err := os.Remove(shared.GetDbPath("current"))
	if err != nil {
		panic(err)
	}
	err = os.Rename(shared.GetDbPath("new"), shared.GetDbPath("current"))
	if err != nil {
		panic(err)
	}
	fmt.Println("- Database updated")
}

func updateToDb(dbName string, sourceResults map[string]model.RepoModel) {
	db, err := gorm.Open(sqlite.Open(shared.GetDbPath(dbName)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	for source, repoModel := range sourceResults {
		repo := model.DbRepo{Url: source}
		err := db.Create(&repo).Error
		if err != nil {
			panic("failed to create repo")
		}
		for _, pkgModel := range repoModel {
			authBs, _ := json.Marshal(pkgModel.Author)
			pkg := model.DbRepoItem{
				RepoId:      repo.Id,
				Name:        pkgModel.Name,
				Author:      string(authBs),
				Version:     pkgModel.Version,
				Path:        pkgModel.Path,
				Description: pkgModel.Description,
				Created:     pkgModel.CreatedAt,
				LastUpdated: pkgModel.LastUpdated,
			}
			err := db.Create(&pkg).Error
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	d, _ := db.DB()
	d.Close()
}
