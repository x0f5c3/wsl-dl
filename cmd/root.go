package cmd

import (
	"os"
	"os/signal"
	"runtime"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/x0f5c3/wsl-dl/internal"
)

var workers = runtime.NumCPU()

var rootCmd = &cobra.Command{
	Use:     "wsl-dl [output directory]",
	Args:    cobra.MaximumNArgs(1),
	Short:   "This cli tool will download WSL distros as zip files and unpack them to allow for custom installations",
	Version: "v0.0.3", // <---VERSION---> Updating this version, will also create a new GitHub release.
	// Uncomment the following lines if your bare application has an action associated with it:
	RunE: runE,
}

func runE(_ *cobra.Command, args []string) error {
	var outDir string
	if len(args) == 0 {
		outDir = "."
	} else {
		outDir = args[0]
	}
	sel, err := internal.AskForDistro()
	if err != nil {
		return err
	}
	return internal.DownloadDistro(sel, outDir, workers)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		cobra.CheckErr(pcli.CheckForUpdates())
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(pcli.CheckForUpdates())
		os.Exit(1)
	}

	cobra.CheckErr(pcli.CheckForUpdates())
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "", false, "disables update checks")
	rootCmd.PersistentFlags().IntVarP(&workers, "workers", "w", workers, "number of workers to use for downloading")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	cobra.CheckErr(pcli.SetRepo("x0f5c3/wsl-dl"))
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
