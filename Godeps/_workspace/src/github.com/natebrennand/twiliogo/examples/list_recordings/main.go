package main

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/recording"
)

func main() {
	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Recordings.List(recording.ListFilter{
		CallSid: "CAa612188279206d89f1d75e3690536011",
	})
	if err != nil {
		fmt.Println("Error getting call list: ", err.Error())
	}

	fmt.Printf("%#v\n", resp)
	for _, m := range *resp.Recordings {
		fmt.Printf("%#v\n", m)
	}
}
