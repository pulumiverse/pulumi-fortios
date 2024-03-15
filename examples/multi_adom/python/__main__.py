import pulumi
import pulumiverse_fortios as fortios

script = fortios.fmg.DevicemanagerScript("script",
    name="config-intf3",
    description="configure interface3",
    content="""config system interface
 edit port3
	 set vdom "root"
	 set ip 10.10.0.200 255.255.0.0
	 set allowaccess ping http https
	 next
 end""",
    target="device_database",
    adom="test-adom")

script_execute = fortios.fmg.DevicemanagerScriptExecute("script",
    script_name=script.name,
    target_devname="FGVM64-test",
    adom="test-adom",
    vdom="root")

_ = fortios.fmg.DevicemanagerInstallDevice("script",
    target_devname=script_execute.target_devname,
    adom="test-adom",
    vdom="root")
