// Package cmd implements the CobraCLI commands for the identityvalidation CLI. Subcommands for the CLI should all live within
// this package. Logic should be delegated to internal packages and functions to keep the CLI commands clean and
// focused on CLI I/O.
package cmd

import (
	"errors"
	"strings"
	"time"

	"github.com/Method-Security/identityvalidation/internal/config"
	"github.com/Method-Security/pkg/signal"
	"github.com/Method-Security/pkg/writer"
	"github.com/palantir/pkg/datetime"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"

	// Import wlog-zap for its side effects, initializing the zap logger
	_ "github.com/palantir/witchcraft-go-logging/wlog-zap"
	"github.com/spf13/cobra"
)

// IdentityValidation is the main struct that holds the root command and all subcommands that are used throughout execution
// of the CLI. It is also responsible for holding the Output configuration, and Output signal
// for use by subcommands. The output signal is used to write the output of the command to the desired output format
// after the execution of the invoked commands Run function.
type IdentityValidation struct {
	Version      string
	RootFlags    config.RootFlags
	OutputConfig writer.OutputConfig
	OutputSignal signal.Signal
	RootCmd      *cobra.Command
}

// NewIdentityValidation returns a new IdentityValidation struct with the provided version string. The IdentityValidation struct is used to
// initialize the root command and all subcommands that are used throughout execution of the CLI.
// We pass the version command in here from the main.go file, where we set the version string during the build process.
func NewIdentityValidation(version string) *IdentityValidation {
	identityValidation := IdentityValidation{
		Version: version,
		RootFlags: config.RootFlags{
			Quiet:   false,
			Verbose: false,
		},
		OutputConfig: writer.NewOutputConfig(nil, writer.NewFormat(writer.SIGNAL)),
		OutputSignal: signal.NewSignal(nil, datetime.DateTime(time.Now()), nil, 0, nil),
	}
	return &identityValidation
}

// Helper function to set up common configurations
func (a *IdentityValidation) setupCommonConfig(cmd *cobra.Command, outputFormat string, outputFile string) error {
	var err error

	cmd.SetContext(svc1log.WithLogger(cmd.Context(), config.InitializeLogging(cmd, &a.RootFlags)))
	format, err := validateOutputFormat(outputFormat)
	if err != nil {
		return err
	}
	var outputFilePointer *string
	if outputFile != "" {
		outputFilePointer = &outputFile
	} else {
		outputFilePointer = nil
	}
	a.OutputConfig = writer.NewOutputConfig(outputFilePointer, format)

	return nil
}

// InitRootCommand initializes the root command for the identityvalidation CLI. This command is used to set the global flags
// that are used by all subcommands, such as the output format, and output file. It also initializes the
// version command that prints the version of the CLI.
// Critically, this sets the PersistentPreRunE and PersistentPostRunE functions that are inherited by most subcommands.
// The PersistentPostRunE function is used to write the output of the command to the desired output format after the execution of the invoked
// command's Run function.
func (a *IdentityValidation) InitRootCommand() {
	var outputFormat string
	var outputFile string
	a.RootCmd = &cobra.Command{
		Use:          "identityvalidation",
		Short:        "Validation of Security Controls pertianing to Identity",
		Long:         "Validation of Security Controls pertianing to Identity",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return a.setupCommonConfig(cmd, outputFormat, outputFile)
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			completedAt := datetime.DateTime(time.Now())
			a.OutputSignal.CompletedAt = &completedAt
			return writer.Write(
				a.OutputSignal.Content,
				a.OutputConfig,
				a.OutputSignal.StartedAt,
				a.OutputSignal.CompletedAt,
				a.OutputSignal.Status,
				a.OutputSignal.ErrorMessage,
			)
		},
	}

	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Quiet, "quiet", "q", false, "Suppress output")
	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Verbose, "verbose", "v", false, "Verbose output")
	a.RootCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "f", "", "Path to output file. If blank, will output to STDOUT")
	a.RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "signal", "Output format (signal, json, yaml). Default value is signal")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of identityvalidation",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(a.Version)
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}

	a.RootCmd.AddCommand(versionCmd)
}

// A utility function to validate that the provided output format is one of the supported formats: json, yaml, signal.
func validateOutputFormat(output string) (writer.Format, error) {
	var format writer.FormatValue
	switch strings.ToLower(output) {
	case "json":
		format = writer.JSON
	case "yaml":
		return writer.Format{}, errors.New("yaml output format is not supported for identityvalidation")
	case "signal":
		format = writer.SIGNAL
	default:
		return writer.Format{}, errors.New("invalid output format. Valid formats are: json, yaml, signal")
	}
	return writer.NewFormat(format), nil
}
