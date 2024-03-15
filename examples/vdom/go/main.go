package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/networking"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fortiosProvider, err := fortios.NewProvider(ctx, "fortios-provider", &fortios.ProviderArgs{
			Cabundlefile: pulumi.String("/path/yourCA.crt"),
			Hostname:     pulumi.String("192.168.52.177"),
			Insecure:     pulumi.Bool(false),
			Token:        pulumi.String("q3Hs49jxts195gkd9Hjsxnjtmr6k39"),
			Vdom:         pulumi.String("vmdomtest"),
		})
		if err != nil {
			return err
		}
		_, err = networking.NewRouteStatic(ctx, "test1", &networking.RouteStaticArgs{
			Dst:       pulumi.String("120.2.2.122/32"),
			Gateway:   pulumi.String("2.2.2.2"),
			Blackhole: pulumi.String("disable"),
			Distance:  pulumi.String("22"),
			Weight:    pulumi.String("3"),
			Priority:  pulumi.String("3"),
			Device:    pulumi.String("lbforvdomtest"),
			Comment:   pulumi.String("Terraform test"),
		}, pulumi.Provider(fortiosProvider))
		if err != nil {
			return err
		}
		return nil
	})
}
