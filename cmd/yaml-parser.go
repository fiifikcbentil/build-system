package cmd

import (
 "fmt"
 "github.com/spf13/cobra"
 "github.com/fiifikcbentil/build-system/internal/yaml-parser/yamlparser"
)

var parserCmd = &cobra.Command{
 Use:   "parse",
 Short: "Parse <filepath>",
 Long:  "Parses the specified yaml file.",
 Args:  cobra.ExactArgs(1),
 Run: func(cmd *cobra.Command, args []string) {
  path := args[0]
  fmt.Printf("parsing the file: %s", path)
  tasks := yamlparser.parse()
  fmt.Printf("Task added: %s\n", description)
 },
}
