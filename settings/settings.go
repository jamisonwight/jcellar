package settings

type Config struct {
	APIKey           string
	PathRoot         string
	InitQueryString  bool
	InitLocationPath bool
	InitView         string
	InitData         []InitData
	AssertViewState  bool
	TraverseHistory  bool
	AddLocales       []string
	SetLocale        string
	AddMessages      []CustomMessage
	AddTemplates     []CustomTemplate
	DisableCache     bool
	Stats            Stats
	ViewAliases      []ViewAlias
}

type Stats struct {
	FBG bool
	GTM bool
}

type ViewAlias struct {
	Name  string
	Alias string
}

type InitData struct {
	Name string
	Data string
}

type CustomMessage struct {
	View    string
	Name    string
	Message string
}

type CustomTemplate struct {
	View string
	Name string
	URL  string
}

func GetSettingsDefault() map[string]any {
	return map[string]any{}
}
