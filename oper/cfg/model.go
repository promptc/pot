package cfg

type Model struct {
	Source []string `json:"source"`
}

func defaultModel() *Model {
	return &Model{
		Source: []string{"https://raw.githubusercontent.com/promptc/repository/main/db.json"},
	}
}
