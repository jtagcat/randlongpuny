package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/net/idna"
)

func main() {
	app := &cli.App{
		Name:        "randlongpuny",
		Description: "Generates a long random punycode from charset. Likely inefficient. Increadably dumb.",
		Usage:       "randlongpuny <charset>...",
		Action: func(c *cli.Context) error {
			args := c.Args()
			if !args.Present() {
				return fmt.Errorf("no charset specified")
			}
			charset := []rune(strings.Join(args.Slice(), ""))

			var domain []rune
			r := rand.New(rand.NewSource(time.Now().UnixNano())) // multiple runs should give different results
			for {
				// rand char
				r := r.Intn(len(charset))
				domain = append(domain, charset[r])

				if full, err := idna.ToASCII(string(domain)); err != nil {
					break
				} else {
					if len(full) > 63 {
						break
					}
				}
			}

			lastValid := string(domain[:len(domain)-1])
			fmt.Println(lastValid)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
