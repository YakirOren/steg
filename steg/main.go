package main

import (
	steg "steg/steg/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var cmdAdd = &cobra.Command{
		Use:   "append [srcFileName] [dstFileName]",
		Short: "Append a file to the end of another file.",
		Long:  ``,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			steg.AppendToFile(args[0], steg.ReadData(args[1]))
		},
	}

	var cmdNew = &cobra.Command{
		Use:   "new [srcFileName]",
		Short: "Create a file that you can send on discord",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			baseDir := steg.CopyTemplate(args[0])

			for _, arg := range args {
				steg.AppendToFile(baseDir, steg.ReadData(arg))

			}
		},
	}

	var cmdSplit = &cobra.Command{
		Use:   "split [srcFileName]",
		Short: "Split a combined file",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			steg.Split(args[0])
		},
	}

	var cmdPrint = &cobra.Command{
		Use:   "print [srcFileName]",
		Short: "Print file content",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			steg.PrintFileContent(args[0])
		},
	}

	//cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "steg"}
	rootCmd.AddCommand(cmdNew, cmdAdd, cmdSplit, cmdPrint)
	rootCmd.Execute()
}
