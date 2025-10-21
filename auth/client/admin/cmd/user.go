package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	UserCmd.AddCommand(GetUserCmd)
	UserCmd.Flags().StringVarP(&getUserCmdConfig.Username, "username", "u", "", "Username of the user to manage")
}

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users in the authentication service",
	Long:  `This command allows administrators to manage users in the authentication service.`,
}

type GetUserCmdConfig struct {
	Username string
}

var getUserCmdConfig = &GetUserCmdConfig{}

var GetUserCmd = &cobra.Command{
	Use:   "get",
	Short: "Get user details",
	Long:  `Retrieve details of a specific user by username.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		opts := []grpc.DialOption{}

		conn, err := grpc.NewClient(adminCmdConfig.Server, opts...)
		if err != nil {

		}
		defer conn.Close()
	},
}
