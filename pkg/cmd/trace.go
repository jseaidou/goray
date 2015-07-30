package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCmdTracer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trace",
		Short: "Invokes a simple ray tracer",
		Long:  "Invokes a simple ray tracer with limited capabilities",
		Run: func(cmd *cobra.Command, args []string) {
			err := trace()
			if err != nil {
				log.Fatalf("%v", err)
			}

		},
	}
	return cmd
}

func trace() error {
	log.Debug("tracing")
	return nil
}
