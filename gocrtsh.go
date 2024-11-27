package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// Define command line flags
	help := flag.Bool("h", false, "Display help message")
	flag.BoolVar(help, "help", false, "Display help message")

	flag.Parse()

	if *help || len(flag.Args()) == 0 {
		printHelp()
		return
	}

	target := flag.Arg(0)
	target = strings.TrimSpace(target)

	url := fmt.Sprintf("https://crt.sh/?q=%s&output=json", target)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var json_data []map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&json_data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	//parse domains
	var domains []string
	for _, value := range json_data {
		var domainString = strings.TrimSpace(string(value["name_value"].(string)))
		//prefix to remove
		prefix := target + "\n"
		domainString = strings.Replace(domainString, prefix, "", -1)
		domains = append(domains, domainString)
	}
	domains = removeDuplicates(domains)

	//print domains
	for _, domain := range removeDuplicates(domains) {
		fmt.Println(domain)
	}
}

func removeDuplicates(input []string) []string {
	var output []string
	seen := make(map[string]bool)

	for _, s := range input {
		exists := seen[s]
		if !exists {
			seen[s] = true
			output = append(output, s)
		}
	}

	return output
}

func printHelp() {
	fmt.Println("Usage: crt.sh <DOMAIN>")
	flag.PrintDefaults()
}
