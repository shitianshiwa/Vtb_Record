package main

import (
	. "github.com/fzxiao233/Vtb_Record/plugins"
	"github.com/fzxiao233/Vtb_Record/plugins/monitor"
	. "github.com/fzxiao233/Vtb_Record/utils"
	"log"
)

func arrangeTask() {
	var ch chan int
	for _, module := range Config.Module {
		if module.Enable {
			for _, usersConfig := range module.Users {
				log.Printf("%s|%s is up", module.Name, usersConfig.Name)
				go StartMonitor(monitor.CreateVideoMonitor(module.Name), usersConfig)
			}
		}
	}
	<-ch
}
func main() {
	arrangeTask()
}
