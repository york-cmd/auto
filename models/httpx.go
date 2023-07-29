package models

type HttpxResult struct {
	Timestamp     string   `json:"timestamp"`
	Hash          HashData `json:"hash"`
	Port          string   `json:"port"`
	URL           string   `json:"url"`
	Input         string   `json:"input"`
	Title         string   `json:"title,omitempty"` // 使用omitempty表示title字段是可选的
	Scheme        string   `json:"scheme"`
	WebServer     string   `json:"webserver"`
	ContentType   string   `json:"content_type"`
	Method        string   `json:"method"`
	Host          string   `json:"host"`
	Path          string   `json:"path"`
	FinalUrl      string   `json:"final_url,omitempty"`
	Time          string   `json:"time"`
	A             []string `json:"a"`
	Cname         []string `json:"cname"`
	Words         int      `json:"words"`
	Lines         int      `json:"lines"`
	StatusCode    int      `json:"status_code"`
	ContentLength int      `json:"content_length"`
	Failed        bool     `json:"failed"`
}

type HashData struct {
	BodyMD5       string `json:"body_md5"`
	BodyMMH3      string `json:"body_mmh3"`
	BodySHA256    string `json:"body_sha256"`
	BodySimhash   string `json:"body_simhash"`
	HeaderMD5     string `json:"header_md5"`
	HeaderMMH3    string `json:"header_mmh3"`
	HeaderSHA256  string `json:"header_sha256"`
	HeaderSimhash string `json:"header_simhash"`
}
