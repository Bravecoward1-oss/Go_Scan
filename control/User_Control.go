package control

import (
	"flag"
	"fmt"

	"Go_Scan/utils"
)

var (
	banner       string
	simpleBanner string
	helpBanner   string
	t            target
)

func Run() {
	bannerControl()
	t.url, t.dictionary = TerminalParameter()
	t.payloads(t.url, t.dictionary)
}

// TerminalParameter Obtain parameters from the terminal
func TerminalParameter() (string, string) {

	var (
		help       bool
		version    bool
		url        string
		dictionary string
	)

	flag.BoolVar(&help, "h", false, "help")
	flag.BoolVar(&version, "v", false, "version")
	flag.StringVar(&url, "u", "", "url")
	flag.StringVar(&dictionary, "d", "", "dictionary")

	flag.Parse()

	if help {
		fmt.Println(helpBanner)
	} else {
		fmt.Println(simpleBanner)
	}

	if version {
		cfg := utils.LoadConfig("../../Go_Scan/config.ini")
		VERSION := cfg.Section("version").Key("VERSION").String()
		fmt.Println("Version:", VERSION)
	}

	return url, dictionary
}
