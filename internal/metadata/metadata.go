package metadata

var (
	appName    string = "gyaml"
	appDesc    string = "Golang YAML Tool"
	authorName string = "Takumi Takahashi"
)

func AppName() string {
	return appName
}

func AppDesc() string {
	return appDesc
}

func AuthorName() string {
	return authorName
}
