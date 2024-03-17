import pulumiverse_fortios as fortios

fortios_provider = fortios.Provider("fortios-provider",
    cabundlefile="/path/yourCA.crt",
    hostname="192.168.52.177",
    insecure=False,
    token="q3Hs49jxts195gkd9Hjsxnjtmr6k39",
    vdom="vmdomtest")

test1 = fortios.networking.RouteStatic("test1",
    dst="120.2.2.122/32",
    gateway="2.2.2.2",
    blackhole="disable",
    distance="22",
    weight="3",
    priority="3",
    device="lbforvdomtest",
    comment="Terraform test",
    opts=pulumi.ResourceOptions(provider=fortios_provider))
