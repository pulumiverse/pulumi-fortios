using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Fortios = Pulumiverse.Fortios;

return await Deployment.RunAsync(() =>
{
    var config = new Config();

    var fortiosCabundlefile = config.Require("fortios:cabundlefile");
    var fortiosHostname = config.Require("fortios:hostname");
    var fortiosInsecure = config.Require("fortios:insecure");
    var fortiosToken = config.Require("fortios:token");

    var fortiosProvider = new Fortios.Provider("fortios-provider", new()
    {
        Cabundlefile = fortiosCabundlefile,
        Hostname = fortiosHostname,
        Insecure = fortiosInsecure,
        Token = fortiosToken,
    });

    var test1 = new Fortios.Networking.RouteStatic("test1", new()
    {
        Dst = "110.2.2.122/32",
        Gateway = "2.2.2.2",
        Device = "device",
    }, new CustomResourceOptions
    {
        Provider = fortiosProvider,
    });

});
