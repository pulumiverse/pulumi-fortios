import * as pulumi from "@pulumi/pulumi";
import * as fortios from "@pulumiverse/fortios";

const fortiosProvider = new fortios.Provider("fortios-provider", {
    cabundlefile: "/path/yourCA.crt",
    hostname: "192.168.52.177",
    insecure: false,
    token: "q3Hs49jxts195gkd9Hjsxnjtmr6k39",
    vdom: "vmdomtest",
});
const test1 = new fortios.networking.RouteStatic("test1", {
    dst: "120.2.2.122/32",
    gateway: "2.2.2.2",
    blackhole: "disable",
    distance: "22",
    weight: "3",
    priority: "3",
    device: "lbforvdomtest",
    comment: "Terraform test",
}, {
    provider: fortiosProvider,
});
