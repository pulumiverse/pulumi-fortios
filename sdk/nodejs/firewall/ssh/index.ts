// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { HostkeyArgs, HostkeyState } from "./hostkey";
export type Hostkey = import("./hostkey").Hostkey;
export const Hostkey: typeof import("./hostkey").Hostkey = null as any;
utilities.lazyLoad(exports, ["Hostkey"], () => require("./hostkey"));

export { LocalcaArgs, LocalcaState } from "./localca";
export type Localca = import("./localca").Localca;
export const Localca: typeof import("./localca").Localca = null as any;
utilities.lazyLoad(exports, ["Localca"], () => require("./localca"));

export { LocalkeyArgs, LocalkeyState } from "./localkey";
export type Localkey = import("./localkey").Localkey;
export const Localkey: typeof import("./localkey").Localkey = null as any;
utilities.lazyLoad(exports, ["Localkey"], () => require("./localkey"));

export { SettingArgs, SettingState } from "./setting";
export type Setting = import("./setting").Setting;
export const Setting: typeof import("./setting").Setting = null as any;
utilities.lazyLoad(exports, ["Setting"], () => require("./setting"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:firewall/ssh/hostkey:Hostkey":
                return new Hostkey(name, <any>undefined, { urn })
            case "fortios:firewall/ssh/localca:Localca":
                return new Localca(name, <any>undefined, { urn })
            case "fortios:firewall/ssh/localkey:Localkey":
                return new Localkey(name, <any>undefined, { urn })
            case "fortios:firewall/ssh/setting:Setting":
                return new Setting(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "firewall/ssh/hostkey", _module)
pulumi.runtime.registerResourceModule("fortios", "firewall/ssh/localca", _module)
pulumi.runtime.registerResourceModule("fortios", "firewall/ssh/localkey", _module)
pulumi.runtime.registerResourceModule("fortios", "firewall/ssh/setting", _module)
