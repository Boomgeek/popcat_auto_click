package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	fpid, err := robotgo.FindIds("Google")
	if err == nil {
		fmt.Println(fpid)
		if len(fpid) > 0 {
			robotgo.ActivePID(fpid[0])
			fmt.Println(robotgo.GetTitle())
			playload := "BOOM"

			n := 0
			for {
				n += 1
				fmt.Println(n)
				robotgo.TypeStr(playload)
				robotgo.MicroSleep(50)
			}
		}
	}
}
