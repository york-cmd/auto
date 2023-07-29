package models

type Xray struct {
	CreateTime int64  `json:"create_time"`
	Detail     Detail `json:"detail"`
	Plugin     string `json:"plugin"`
	Target     Target `json:"target"`
}

type Target struct {
	URL string `json:"url"`
}
type Nuclei struct {
	Template         string   `json:"template"`
	TemplateURL      string   `json:"template-url"`
	TemplateID       string   `json:"template-id"`
	TemplatePath     string   `json:"template-path"`
	Info             Info     `json:"info"`
	Type             string   `json:"type"`
	Host             string   `json:"host"`
	MatchedAt        string   `json:"matched-at"`
	ExtractedResults []string `json:"extracted-results"`
	Request          string   `json:"request"`
	Response         string   `json:"response"`
	IP               string   `json:"ip"`
	Timestamp        string   `json:"timestamp"`
	CurlCommand      string   `json:"curl-command"`
	MatcherStatus    bool     `json:"matcher-status"`
	MatchedLine      string   `json:"matched-line"`
}
type Xscan struct {
	Desc           string `json:"desc"`
	Key            string `json:"key"`
	Line           string `json:"line"`
	Payload        string `json:"payload"`
	Position       string `json:"position"`
	Req            string `json:"req"`
	SuggestPayload string `json:"suggest-payload"`
	URL            string `json:"url"`
	XSSType        string `json:"xssType"`
}
type Snapshot struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Detail struct {
	Addr     string     `json:"addr"`
	Payload  string     `json:"payload"`
	Snapshot [][]string `json:"snapshot"`
	Extra    Extra      `json:"extra"`
}
type Extra struct {
	Links []string `json:"Links"`
	Level string   `json:"level"`
	Param struct{} `json:"param"`
}
type Info struct {
	Name      string   `json:"name"`
	Author    []string `json:"author"`
	Tags      []string `json:"tags"`
	Reference []string `json:"reference"`
	Severity  string   `json:"severity"`
	Metadata  Metadata `json:"metadata"`
}

type Metadata struct {
	MaxRequest  int    `json:"max-request"`
	ShodanQuery string `json:"shodan-query"`
	Verified    bool   `json:"verified"`
}
