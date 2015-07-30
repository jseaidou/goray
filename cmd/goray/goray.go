package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jseaidou/goray/pkg/cmd"
)

func main() {
	cmd := cmd.NewGoRayCommand()
	log.SetLevel(log.DebugLevel)
	if err := cmd.Execute(); err != nil {

	}
	return
}
