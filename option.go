package phantomjs

type Option struct {
	Headers        map[string]string
	Timeout        int
	UseGzip        bool
	AllowRedirects bool
	load_images    bool
	Time           float64
	JsScriptResult string
	FetcherJsPath  string
}
