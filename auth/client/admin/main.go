package main

import "github.com/open-cloud-initiative/glue/auth/client/admin/cmd"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
