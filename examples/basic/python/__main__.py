import pulumi
import pulumiverse_fortios as fortios

config = pulumi.Config()

fortios_cabundlefile = config.require("fortios:cabundlefile")
fortios_hostname = config.require("fortios:hostname")
fortios_insecure = config.require("fortios:insecure")
fortios_token = config.require("fortios:token")

fortios_provider = fortios.Provider("fortios-provider",
    cabundlefile=fortios_cabundlefile,
    hostname=fortios_hostname,
    insecure=fortios_insecure,
    token=fortios_token)

test1 = fortios.networking.RouteStatic("test1",
    dst="110.2.2.122/32",
    gateway="2.2.2.2",
    device="device",
    opts=pulumi.ResourceOptions(provider=fortios_provider))
