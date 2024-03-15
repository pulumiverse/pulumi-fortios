package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/fmg"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		script, err := fmg.NewDevicemanagerScript(ctx, "script", &fmg.DevicemanagerScriptArgs{
			Name:        pulumi.String("config-intf3"),
			Description: pulumi.String("configure interface3"),
			Content: pulumi.String(`config system interface
 edit port3
	 set vdom "root"
	 set ip 10.10.0.200 255.255.0.0
	 set allowaccess ping http https
	 next
 end`),
			Target: pulumi.String("device_database"),
			Adom:   pulumi.String("test-adom"),
		})
		if err != nil {
			return err
		}
		scriptExecute, err := fmg.NewDevicemanagerScriptExecute(ctx, "script", &fmg.DevicemanagerScriptExecuteArgs{
			ScriptName:    script.Name,
			TargetDevname: pulumi.String("FGVM64-test"),
			Adom:          pulumi.String("test-adom"),
			Vdom:          pulumi.String("root"),
		})
		if err != nil {
			return err
		}
		_, err = fmg.NewDevicemanagerInstallDevice(ctx, "script", &fmg.DevicemanagerInstallDeviceArgs{
			TargetDevname: scriptExecute.TargetDevname,
			Adom:          pulumi.String("test-adom"),
			Vdom:          pulumi.String("root"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
