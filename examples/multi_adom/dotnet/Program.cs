using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Fortios = Pulumiverse.Fortios;

return await Deployment.RunAsync(() =>
{
    var script = new Fortios.Fmg.DevicemanagerScript("script", new()
    {
        Name = "config-intf3",
        Description = "configure interface3",
        Content = @"config system interface
 edit port3
	 set vdom ""root""
	 set ip 10.10.0.200 255.255.0.0
	 set allowaccess ping http https
	 next
 end",
        Target = "device_database",
        Adom = "test-adom",
    });

    var scriptExecute = new Fortios.Fmg.DevicemanagerScriptExecute("script", new()
    {
        ScriptName = script.Name,
        TargetDevname = "FGVM64-test",
        Adom = "test-adom",
        Vdom = "root",
    });

    var installDevice = new Fortios.Fmg.DevicemanagerInstallDevice("script", new()
    {
        TargetDevname = scriptExecute.TargetDevname,
        Adom = "test-adom",
        Vdom = "root",
    });

});
