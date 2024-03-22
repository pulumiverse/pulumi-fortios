// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ExemptlistArgs, ExemptlistState } from "./exemptlist";
export type Exemptlist = import("./exemptlist").Exemptlist;
export const Exemptlist: typeof import("./exemptlist").Exemptlist = null as any;
utilities.lazyLoad(exports, ["Exemptlist"], () => require("./exemptlist"));

export { HeuristicArgs, HeuristicState } from "./heuristic";
export type Heuristic = import("./heuristic").Heuristic;
export const Heuristic: typeof import("./heuristic").Heuristic = null as any;
utilities.lazyLoad(exports, ["Heuristic"], () => require("./heuristic"));

export { ProfileArgs, ProfileState } from "./profile";
export type Profile = import("./profile").Profile;
export const Profile: typeof import("./profile").Profile = null as any;
utilities.lazyLoad(exports, ["Profile"], () => require("./profile"));

export { QuarantineArgs, QuarantineState } from "./quarantine";
export type Quarantine = import("./quarantine").Quarantine;
export const Quarantine: typeof import("./quarantine").Quarantine = null as any;
utilities.lazyLoad(exports, ["Quarantine"], () => require("./quarantine"));

export { SettingsArgs, SettingsState } from "./settings";
export type Settings = import("./settings").Settings;
export const Settings: typeof import("./settings").Settings = null as any;
utilities.lazyLoad(exports, ["Settings"], () => require("./settings"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:antivirus/exemptlist:Exemptlist":
                return new Exemptlist(name, <any>undefined, { urn })
            case "fortios:antivirus/heuristic:Heuristic":
                return new Heuristic(name, <any>undefined, { urn })
            case "fortios:antivirus/profile:Profile":
                return new Profile(name, <any>undefined, { urn })
            case "fortios:antivirus/quarantine:Quarantine":
                return new Quarantine(name, <any>undefined, { urn })
            case "fortios:antivirus/settings:Settings":
                return new Settings(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "antivirus/exemptlist", _module)
pulumi.runtime.registerResourceModule("fortios", "antivirus/heuristic", _module)
pulumi.runtime.registerResourceModule("fortios", "antivirus/profile", _module)
pulumi.runtime.registerResourceModule("fortios", "antivirus/quarantine", _module)
pulumi.runtime.registerResourceModule("fortios", "antivirus/settings", _module)
