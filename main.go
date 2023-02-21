package main

import (
	"fmt"
	"github.com/mstch/go-ratp/pkg/ratp"
)

func main() {
	stop := ratp.Stop{}
	stop.Request(ratp.Montgeron)
	fmt.Printf(stop.String())
}
