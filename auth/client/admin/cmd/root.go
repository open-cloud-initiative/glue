package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(UserCmd)
	RootCmd.PersistentFlags().StringVarP(&adminCmdConfig.Server, "server", "s", "http://localhost:8080", "Address of the authentication server")
}

type AdminCmdConfig struct {
	Server string
}

var adminCmdConfig = &AdminCmdConfig{}

var RootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin client for the authentication service",
	Long:  `This is the admin client for managing the authentication service.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Placeholder for admin client functionality
	},
}
