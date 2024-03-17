using System.Collections.Generic;
using System.Linq;
using Fortios = Pulumiverse.Fortios;

return await Deployment.RunAsync(() =>
{
    var fortiosProvider = new Fortios.Provider("fortios-provider", new()
    {
        Cabundlefile = "/path/yourCA.crt",
        Hostname = "192.168.52.177",
        Insecure = false,
        Token = "q3Hs49jxts195gkd9Hjsxnjtmr6k39",
        Vdom = "vmdomtest",
    });

    var test1 = new Fortios.Networking.RouteStatic("test1", new()
    {
        Dst = "120.2.2.122/32",
        Gateway = "2.2.2.2",
        Blackhole = "disable",
        Distance = "22",
        Weight = "3",
        Priority = "3",
        Device = "lbforvdomtest",
        Comment = "Terraform test",
    }, new CustomResourceOptions
    {
        Provider = fortiosProvider,
    });

});
