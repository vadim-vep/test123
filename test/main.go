package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")

	log.Logger = log.With().Caller().Logger()
	log.Info().Msg("hello world")
	// Output: {"time":1516134303,"level":"debug","message":"hello world"}
	url := "/Users/vadym.veprytskyi/GolandProjects/awesomeProject/test/main.go"
	line := "19"
	fmt.Printf("%v:%v", url, line)

	//test
	//test2
	//test3

}
