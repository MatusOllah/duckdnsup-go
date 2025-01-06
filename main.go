package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func main() {
	domains := pflag.StringSliceP("domain", "d", []string{}, "The domain(s) to be updated")
	ip := pflag.StringP("ip", "i", "", "The IP address, auto-detect when empty")
	token := pflag.StringP("token", "t", "", "The token")
	verbose := pflag.BoolP("verbose", "v", false, "Print verbose information")
	pflag.Parse()

	if len(*domains) == 0 {
		fmt.Fprintln(os.Stderr, "Error: domains not specified")
		os.Exit(1)
	}

	if *token == "" {
		fmt.Fprintln(os.Stderr, "Error: token not specified")
		os.Exit(1)
	}

	url := fmt.Sprintf("https://www.duckdns.org/update?domains=%s&token=%s&ip=%s&verbose=%t", strings.Join(*domains, ","), *token, *ip, *verbose)

	if *verbose {
		fmt.Printf("Making HTTP request: %s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}

	if bytes.HasPrefix(body, []byte("OK")) {
		fmt.Println("Update successful.")
	} else if bytes.HasPrefix(body, []byte("KO")) {
		fmt.Println("Update failed.")
	} else {
		fmt.Println("Unexpected response.")
	}

	if *verbose {
		fmt.Printf("Response from Duck DNS:\n%s\n", string(body))
	}
}
