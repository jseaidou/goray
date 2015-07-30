package cmd

import "github.com/spf13/cobra"

func NewGoRayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goray",
		Short: "A simple ray tracer",
		Long:  "A simple ray tracer developed in Go. This Ray tracer is multi threaded and supports different types of ray tracing techniques and is real time",
	}
	cmd.AddCommand(NewCmdTracer())
	return cmd
}
