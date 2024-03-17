import * as pulumi from "@pulumi/pulumi";
import * as fortios from "@pulumiverse/fortios";

const config = new pulumi.Config();
const fortiosCabundlefile = config.require("fortios:cabundlefile");
const fortiosHostname = config.require("fortios:hostname");
const fortiosInsecure = config.require("fortios:insecure");
const fortiosToken = config.require("fortios:token");

const fortiosProvider = new fortios.Provider("fortios-provider", {
    cabundlefile: fortiosCabundlefile,
    hostname: fortiosHostname,
    insecure: fortiosInsecure,
    token: fortiosToken,
});

const test1 = new fortios.networking.RouteStatic("test1", {
    dst: "110.2.2.122/32",
    gateway: "2.2.2.2",
    device: "device",
}, {
    provider: fortiosProvider,
});
