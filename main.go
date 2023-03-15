package main

import (
	"context"
	"fmt"

	"github.com/go-jarvis/cobrautils"
	"github.com/spf13/cobra"
	"github.com/tangx/envsubst2/pkg/envsubst2"
	"github.com/tangx/envsubst2/version"
)

func main() {
	err := root.Execute()
	if err != nil {
		panic(err)
	}
}

var tip = `Render environment variables into file. 
Same as the envsubst.

version: %s
`
var root = &cobra.Command{
	Use: "envsubst2",
	// Short: version.Version,
	Long: fmt.Sprintf(tip, version.Version),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		envsubst2.InitDefaultValues(flag)
		envsubst2.Replace(ctx, flag)
	},
}
var flag = &envsubst2.Flag{
	ForceReplace: true,
}

func init() {
	cobrautils.BindFlags(root, flag)
}
