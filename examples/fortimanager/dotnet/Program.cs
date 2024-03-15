using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Fortios = Pulumiverse.Fortios;

return await Deployment.RunAsync(() => 
{
    var fortiosProvider = new Fortios.Provider("fortios-provider", new()
    {
        FmgCabundlefile = "/path/yourCA.crt",
        FmgHostname = "192.168.88.100",
        FmgInsecure = false,
        FmgPasswd = "admin",
        FmgUsername = "APIUser",
    });

    var test1 = new Fortios.Fmg.SystemDns("test1", new()
    {
        Primary = "208.91.112.52",
        Secondary = "208.91.112.54",
    }, new CustomResourceOptions
    {
        Provider = fortiosProvider,
    });

});

