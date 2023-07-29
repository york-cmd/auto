package models

type Config struct {
	Vulnscan      Vulnscan  `yaml:"vulnscan"`
	Goroutine     Goroutine `yaml:"goroutine"`
	Subdomain     Subdomain `yaml:"subdomain"`
	Portscan      Portscan  `yaml:"portscan"`
	Other         Other     `yaml:"other"`
	ServerSendKey string    `yaml:"serverSendKey"`
}

type Vulnscan struct {
	XscanSingle    string `yaml:"XscanSingle"`
	XscanMultiple  string `yaml:"XscanMultiple"`
	NucleiSingle   string `yaml:"NucleiSingle"`
	NucleiMultiple string `yaml:"NucleiMultiple"`
	Crawlergo      string `yaml:"Crawlergo"`
	XrayListen     string `yaml:"XrayListen"`
}

type Goroutine struct {
	Cdn       int `yaml:"Cdn"`
	Nmap      int `yaml:"Nmap"`
	Crawlergo int `yaml:"Crawlergo"`
}

type Subdomain struct {
	KsubVerify string `yaml:"KsubVerify"`
	OneForAll  string `yaml:"OneForAll"`
	KsubEnum   string `yaml:"KsubEnum"`
	Subfinder  string `yaml:"Subfinder"`
	Alterx     string `yaml:"Alterx"`
}

type Portscan struct {
	Masscan string `yaml:"Masscan"`
	Nmap    string `yaml:"Nmap"`
}

type Other struct {
	Wafw00f    string `yaml:"Wafw00f"`
	Httpx      string `yaml:"Httpx"`
	TideFinger string `yaml:"TideFinger"`
}
