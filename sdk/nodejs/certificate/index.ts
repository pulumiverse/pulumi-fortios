// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CaArgs, CaState } from "./ca";
export type Ca = import("./ca").Ca;
export const Ca: typeof import("./ca").Ca = null as any;
utilities.lazyLoad(exports, ["Ca"], () => require("./ca"));

export { CrlArgs, CrlState } from "./crl";
export type Crl = import("./crl").Crl;
export const Crl: typeof import("./crl").Crl = null as any;
utilities.lazyLoad(exports, ["Crl"], () => require("./crl"));

export { LocalArgs, LocalState } from "./local";
export type Local = import("./local").Local;
export const Local: typeof import("./local").Local = null as any;
utilities.lazyLoad(exports, ["Local"], () => require("./local"));

export { RemoteArgs, RemoteState } from "./remote";
export type Remote = import("./remote").Remote;
export const Remote: typeof import("./remote").Remote = null as any;
utilities.lazyLoad(exports, ["Remote"], () => require("./remote"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:certificate/ca:Ca":
                return new Ca(name, <any>undefined, { urn })
            case "fortios:certificate/crl:Crl":
                return new Crl(name, <any>undefined, { urn })
            case "fortios:certificate/local:Local":
                return new Local(name, <any>undefined, { urn })
            case "fortios:certificate/remote:Remote":
                return new Remote(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "certificate/ca", _module)
pulumi.runtime.registerResourceModule("fortios", "certificate/crl", _module)
pulumi.runtime.registerResourceModule("fortios", "certificate/local", _module)
pulumi.runtime.registerResourceModule("fortios", "certificate/remote", _module)
