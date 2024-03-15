import * as pulumi from "@pulumi/pulumi";
import * as fortios from "@pulumiverse/fortios";

const script = new fortios.fmg.DevicemanagerScript("script", {
    name: "config-intf3",
    description: "configure interface3",
    content: `config system interface
 edit port3
	 set vdom "root"
	 set ip 10.10.0.200 255.255.0.0
	 set allowaccess ping http https
	 next
 end`,
    target: "device_database",
    adom: "test-adom",
});

const scriptExecute = new fortios.fmg.DevicemanagerScriptExecute("script", {
    scriptName: script.name,
    targetDevname: "FGVM64-test",
    adom: "test-adom",
    vdom: "root",
});

_ = new fortios.fmg.DevicemanagerInstallDevice("script", {
    targetDevname: scriptExecute.targetDevname,
    adom: "test-adom",
    vdom: "root",
});
