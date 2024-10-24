package cmd

import (
	identityvalidate "github.com/Method-Security/identityvalidate/generated/go"
	portal "github.com/Method-Security/identityvalidate/internal/portal"
	"github.com/spf13/cobra"
)

func (a *IdentityValidate) InitPortalCommand() {
	portalCmd := &cobra.Command{
		Use:   "portal",
		Short: "Trigger portal security controls",
		Long:  `Trigger portal security controls`,
	}

	azureCmd := &cobra.Command{
		Use:   "azure",
		Short: "Azure triggers",
		Long:  `Azure triggers`,
	}

	owaCmd := &cobra.Command{
		Use:   "owa",
		Short: "Azure OWA portal trigger",
		Long:  `Azure OWA portal trigger`,
		Run: func(cmd *cobra.Command, args []string) {
			// Run Configs
			attempts, err := cmd.Flags().GetInt("attempts")
			if err != nil {
				a.OutputSignal.AddError(err)
				return
			}
			username, err := cmd.Flags().GetString("username")
			if err != nil {
				a.OutputSignal.AddError(err)
				return
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				a.OutputSignal.AddError(err)
				return
			}
			agentHeader, err := cmd.Flags().GetString("agentheader")
			if err != nil {
				a.OutputSignal.AddError(err)
				return
			}
			interval, err := cmd.Flags().GetInt("interval")
			if err != nil {
				a.OutputSignal.AddError(err)
				return
			}
			timeout, err := cmd.Flags().GetInt("timeout")
			if err != nil {
				a.OutputSignal.AddError(err)
				return
			}
			config, err := newPortalConfig(
				identityvalidate.PortalTypeAzure,
				identityvalidate.ModuleNameOwaLogin,
				attempts,
				username,
				password,
				agentHeader,
				interval,
				timeout,
			)
			if err != nil {
				a.OutputSignal.AddError(err)
				return
			}

			engine := portal.NewEngine(config)
			report, err := engine.Launch(cmd.Context())
			if err != nil {
				a.OutputSignal.AddError(err)
			}
			a.OutputSignal.Content = report
		},
	}

	owaCmd.Flags().Int("attempts", 1, "Number of retry attempts")
	owaCmd.Flags().String("username", "", "Login portal username")
	owaCmd.Flags().String("password", "", "Login portal password")
	owaCmd.Flags().String("agentheader", "", "The agent header set in the request")
	owaCmd.Flags().Int("interval", 0, "Trigger sleep intervals for making multiple attempts (Seconds)")
	owaCmd.Flags().Int("timeout", 10, "Timeout limit (Seconds)")

	_ = owaCmd.MarkFlagRequired("username")

	portalCmd.AddCommand(azureCmd)
	azureCmd.AddCommand(owaCmd)

	a.RootCmd.AddCommand(portalCmd)
}

func newPortalConfig(portalType identityvalidate.PortalType, moduleName identityvalidate.ModuleName, attempts int, username string, password string, agentHeader string, interval int, timeout int) (*identityvalidate.PortalConfig, error) {
	config := &identityvalidate.PortalConfig{
		PortalType:  portalType,
		ModuleName:  moduleName,
		Attempts:    attempts,
		Username:    username,
		Password:    password,
		AgentHeader: agentHeader,
		Interval:    interval,
		Timeout:     timeout,
	}

	return config, nil
}
