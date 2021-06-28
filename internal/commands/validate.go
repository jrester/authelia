package commands

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/authelia/authelia/internal/configuration"
	"github.com/authelia/authelia/internal/logging"
)

func newValidateConfigCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "validate-config [yaml]",
		Short: "Check a configuration against the internal configuration validation mechanisms",
		Args:  cobra.MinimumNArgs(1),
		Run:   cmdValidateConfigRun,
	}

	return cmd
}

func cmdValidateConfigRun(_ *cobra.Command, args []string) {
	logger := logging.Logger()

	configPath := args[0]
	if _, err := os.Stat(configPath); err != nil {
		logger.Fatalf("Error Loading Configuration: %v\n", err)
	}

	provider := configuration.NewProvider()

	errs := provider.LoadSources(configuration.NewYAMLFileSource(configPath))
	if len(errs) != 0 {
		logger.Error("Error loading configuration sources:")

		for _, err := range errs {
			logger.Errorf("  %+v", err)
		}

		logger.Fatalf("Can't continue due to the errors loading the configuration sources")
	}

	warns, errs := provider.Unmarshal()

	if len(warns) != 0 {
		logger.Warn("Warnings occurred while loading the configuration:")

		for _, warn := range warns {
			logger.Warnf("  %+v", warn)
		}
	}

	if len(errs) != 0 {
		logger.Error("Errors occurred while loading the configuration:")

		for _, err := range errs {
			logger.Errorf("  %+v", err)
		}

		logger.Fatal("Can't continue due to errors")
	}

	log.Println("Configuration parsed successfully without errors.")
}
