// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetNeighborArgs, GetNeighborResult, GetNeighborOutputArgs } from "./getNeighbor";
export const getNeighbor: typeof import("./getNeighbor").getNeighbor = null as any;
export const getNeighborOutput: typeof import("./getNeighbor").getNeighborOutput = null as any;
utilities.lazyLoad(exports, ["getNeighbor","getNeighborOutput"], () => require("./getNeighbor"));

export { GetNeighborlistArgs, GetNeighborlistResult, GetNeighborlistOutputArgs } from "./getNeighborlist";
export const getNeighborlist: typeof import("./getNeighborlist").getNeighborlist = null as any;
export const getNeighborlistOutput: typeof import("./getNeighborlist").getNeighborlistOutput = null as any;
utilities.lazyLoad(exports, ["getNeighborlist","getNeighborlistOutput"], () => require("./getNeighborlist"));

export { NeighborArgs, NeighborState } from "./neighbor";
export type Neighbor = import("./neighbor").Neighbor;
export const Neighbor: typeof import("./neighbor").Neighbor = null as any;
utilities.lazyLoad(exports, ["Neighbor"], () => require("./neighbor"));

export { NetworkArgs, NetworkState } from "./network";
export type Network = import("./network").Network;
export const Network: typeof import("./network").Network = null as any;
utilities.lazyLoad(exports, ["Network"], () => require("./network"));

export { Network6Args, Network6State } from "./network6";
export type Network6 = import("./network6").Network6;
export const Network6: typeof import("./network6").Network6 = null as any;
utilities.lazyLoad(exports, ["Network6"], () => require("./network6"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:router/bgp/neighbor:Neighbor":
                return new Neighbor(name, <any>undefined, { urn })
            case "fortios:router/bgp/network6:Network6":
                return new Network6(name, <any>undefined, { urn })
            case "fortios:router/bgp/network:Network":
                return new Network(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "router/bgp/neighbor", _module)
pulumi.runtime.registerResourceModule("fortios", "router/bgp/network", _module)
pulumi.runtime.registerResourceModule("fortios", "router/bgp/network6", _module)