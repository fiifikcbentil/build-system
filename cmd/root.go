package cmd

import (
 "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
 Use:   "fbuild",
 Short: "Artifact-Based Build System",
 Long:  "fbuild-system is a simple CLI tool to build applications and manage dependencies.",
}

func Execute() {
 cobra.CheckErr(rootCmd.Execute())
}

func init() {
 rootCmd.AddCommand(parserCmd)
}