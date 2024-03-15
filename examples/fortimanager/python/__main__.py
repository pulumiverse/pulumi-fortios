import pulumiverse_fortios as fortios

fortios_provider = fortios.Provider("fortios-provider",
    fmg_cabundlefile="/path/yourCA.crt",
    fmg_hostname="192.168.88.100",
    fmg_insecure=False,
    fmg_passwd="admin",
    fmg_username="APIUser")

test1 = fortios.fmg.SystemDns("test1",
    primary="208.91.112.52",
    secondary="208.91.112.54",
    opts=pulumi.ResourceOptions(provider=fortios_provider))
