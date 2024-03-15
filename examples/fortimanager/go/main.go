package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fortiosProvider, err := fortios.NewProvider(ctx, "fortios-provider", &fortios.ProviderArgs{
			FmgCabundlefile: pulumi.String("/path/yourCA.crt"),
			FmgHostname:     pulumi.String("192.168.88.100"),
			FmgInsecure:     pulumi.Bool(false),
			FmgPasswd:       pulumi.String("admin"),
			FmgUsername:     pulumi.String("APIUser"),
		})
		if err != nil {
			return err
		}
		_, err = fmg.NewSystemDns(ctx, "test1", &fmg.SystemDnsArgs{
			Primary:   pulumi.String("208.91.112.52"),
			Secondary: pulumi.String("208.91.112.54"),
		}, pulumi.Provider(fortiosProvider))
		if err != nil {
			return err
		}
		return nil
	})
}
