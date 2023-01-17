/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cmdOneCmd represents the cmdOne command
var cmdOneCmd = &cobra.Command{
	Use:   "cmdOne",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmdOne called!")
		// Использовать ранее объявленный флаг в блоке init().
		number, _ := cmd.Flags().GetInt("number")
		fmt.Println("Going to use number", number)
		fmt.Printf("Square: %d\n", number*number)
	},
}

func init() {
	rootCmd.AddCommand(cmdOneCmd)
	// Определить новый флаг командной строки.
	cmdOneCmd.Flags().Int("number", 0, "A help for number")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmdOneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmdOneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
