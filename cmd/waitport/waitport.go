package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/miklosn/waitport/pkg/waitport"
)

func init() {
	log.SetFlags(0)
}

func main() {

	timeout := flag.Duration("t", 30*time.Second, "timeout in seconds")
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Println("No port provided")
		os.Exit(2)
	}

	port, err := strconv.ParseUint(flag.Arg(0), 10, 16)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	err = waitport.WaitPort(uint16(port), *timeout)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return
}
