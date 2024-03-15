import * as fortios from "@pulumiverse/fortios";

const fortiosProvider = new fortios.Provider("fortios-provider", {
    fmgCabundlefile: "/path/yourCA.crt",
    fmgHostname: "192.168.88.100",
    fmgInsecure: false,
    fmgPasswd: "admin",
    fmgUsername: "APIUser",
});

const test1 = new fortios.fmg.SystemDns("test1", {
    primary: "208.91.112.52",
    secondary: "208.91.112.54",
}, {
    provider: fortiosProvider,
});
