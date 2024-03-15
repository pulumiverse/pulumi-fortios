package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/networking"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		fortiosCabundlefile := cfg.Require("fortios:cabundlefile")
		fortiosHostname := cfg.Require("fortios:hostname")
		fortiosInsecure := cfg.Require("fortios:insecure")
		fortiosToken := cfg.Require("fortios:token")

		fortiosProvider, err := fortios.NewProvider(ctx, "fortios-provider", &fortios.ProviderArgs{
			Cabundlefile: pulumi.String(fortiosCabundlefile),
			Hostname:     pulumi.String(fortiosHostname),
			Insecure:     pulumi.String(fortiosInsecure),
			Token:        pulumi.String(fortiosToken),
		})
		if err != nil {
			return err
		}
		_, err = networking.NewRouteStatic(ctx, "test1", &networking.RouteStaticArgs{
			Dst:     pulumi.String("110.2.2.122/32"),
			Gateway: pulumi.String("2.2.2.2"),
			Device:  pulumi.String("device"),
		}, pulumi.Provider(fortiosProvider))
		if err != nil {
			return err
		}
		return nil
	})
}
