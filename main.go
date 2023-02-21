package main

import (
	"context"

	"github.com/go-jarvis/cobrautils"
	"github.com/spf13/cobra"
	"github.com/tangx/envsubst2/pkg/envsubst2"
)

func main() {
	err := root.Execute()
	if err != nil {
		panic(err)
	}
}

var root = &cobra.Command{
	Use: "envsubst2",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		envsubst2.Replace(ctx, flag)
	},
}
var flag = &envsubst2.Flag{
	ForceUpdate: true,
}

func init() {
	cobrautils.BindFlags(root, flag)
}
