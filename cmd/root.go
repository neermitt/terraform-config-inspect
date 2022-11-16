package cmd

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var Version = "0.0.1"

var format string

var rootCmd = &cobra.Command{
	Use:   "terraform-config-inspect",
	Short: "Inspect Terraform configuration files",
	Long: `terraform-config-inspect is a tool for inspecting Terraform configuration files.
			It can be used to extract the configuration in a machine-readable format, or to
			validate the configuration for correctness.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "."
		if len(args) != 0 {
			dir = args[0]
		}
		module, _ := tfconfig.LoadModule(dir)

		switch format {
		case "json":
			return showModuleJSON(module)
		case "yaml":
			return showModuleYAML(module)
		default:
			return showModuleMarkdown(module)
		}
	},
	Version: Version,
}

func init() {
	rootCmd.Flags().StringVar(&format, "format", "", "Output format (json or yaml)")
}

func Execute() error {
	return rootCmd.Execute()
}

func showModuleJSON(module *tfconfig.Module) (err error) {
	bw := bufio.NewWriter(os.Stdout)
	defer func() {
		flushErr := bw.Flush()
		if err == nil {
			err = flushErr
		}
	}()
	encoder := json.NewEncoder(bw)
	encoder.SetIndent("", "  ")

	return encoder.Encode(module)
}

func showModuleYAML(module *tfconfig.Module) (err error) {
	bw := bufio.NewWriter(os.Stdout)
	defer func() {
		flushErr := bw.Flush()
		if err == nil {
			err = flushErr
		}
	}()

	encoder := yaml.NewEncoder(bw)

	defer func() {
		flushErr := encoder.Close()
		if err == nil {
			err = flushErr
		}
	}()

	return encoder.Encode(module)
}

func showModuleMarkdown(module *tfconfig.Module) error {
	return tfconfig.RenderMarkdown(os.Stdout, module)
}
