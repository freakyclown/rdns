package main

// rdns is a tool used to transform domains into reverse domain name notation
// (see more here - https://en.wikipedia.org/wiki/Reverse_domain_name_notation)
// it takes an input and strips out the protocol and then splits the domain by
// .'s and the reverses the field order and rebuilds it again
// for example https://www.cygenta.co.uk becomes uk.co.cygenta.www
// you can use RDNS to turn it back again, minus the protocol.
//
// Input can be via command argument or via piped stdin
//
//
// Author: FC aka Freakyclown @ Cygenta.co.uk

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var domain string

	// Check if input is provided through a pipe
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Read input from pipe
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			domain = scanner.Text()
		}
	} else {
		// Read input from command line argument
		if len(os.Args) < 2 {
			fmt.Println("Please provide a domain as input.")
			return
		}
		domain = os.Args[1]
	}

	// Remove "http://" or "https://" from the domain
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimPrefix(domain, "https://")

	// Split the domain into fields based on '.'
	fields := strings.Split(domain, ".")

	// Reverse the order of fields
	reversedFields := make([]string, len(fields))
	for i := len(fields) - 1; i >= 0; i-- {
		reversedFields[len(fields)-1-i] = fields[i]
	}

	// Rebuild the domain using '.'
	reversedDomain := strings.Join(reversedFields, ".")

	fmt.Println(reversedDomain)
}
