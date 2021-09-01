package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	linturl "github.com/JitenPalaparthi/urllinter/pkg/lint"
)

func main() {
	// get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	var pathFlag = flag.String("path", wd, "path to be provided")                                                       // default is current working directory
	var configPathFlag = flag.String("config", ".urllintconfig.yaml", "path for the configuration file to be provided") // default config is the config.json file that is there in the urllint path
	var showSumary = flag.Bool("summary", false, "to get summary pass summary=true;to off either dont pass or summary=false")
	var detailedSummary = flag.String("details", "Fail", "detailed summary can be Fail,Pass")
	flag.Parse()
	llint, err := linturl.New(*configPathFlag)
	if err != nil {
		log.Fatal(err)
	}
	err = llint.Init(*pathFlag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The following is the path that lint is working on: ", *pathFlag)

	isFatal := llint.LintAll()

	if *showSumary {
		llint.ShowSummary()
		fmt.Println()
	}
	switch *detailedSummary {
	case "Fail", "fail", "FAIL":
		llint.ShowFailSummary()
	case "Pass", "pass", "PASS":
		llint.ShowPassSummary()
	}

	if isFatal {
		os.Exit(1)
	}

}
