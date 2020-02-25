package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/influxdata/influxdb/kit/cli"
	"golang.org/x/net/idna"
)

var flags struct {
	decode bool
}

func toUnicode(domain string) string {
	s, _ := idna.ToUnicode(domain)
	return s
}

func toASCII(domain string) string {
	s, _ := idna.ToASCII(domain)
	return s
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if flags.decode {
			fmt.Printf("%v\n", toUnicode(domain))
		} else {
			fmt.Printf("%v\n", toASCII(domain))
		}
	}

	return nil
}

func main() {
	cmd := cli.NewCommand(&cli.Program{
		Run:  run,
		Name: "puny",
		Opts: []cli.Opt{
			{
				DestP:   &flags.decode,
				Flag:    "decode",
				Default: false,
				Desc:    "Decode from Punycode",
			},
		},
	})
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
