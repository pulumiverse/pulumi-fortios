// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to configure virtual IP groups of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.FirewallVipgrp`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const v11 = new fortios.FirewallObjectVipgroup("v11", {
 *     comments: "comments",
 *     "interface": "port3",
 *     members: [
 *         "vip1",
 *         "vip3",
 *     ],
 * });
 * ```
 */
export class FirewallObjectVipgroup extends pulumi.CustomResource {
    /**
     * Get an existing FirewallObjectVipgroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallObjectVipgroupState, opts?: pulumi.CustomResourceOptions): FirewallObjectVipgroup {
        return new FirewallObjectVipgroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallObjectVipgroup:FirewallObjectVipgroup';

    /**
     * Returns true if the given object is an instance of FirewallObjectVipgroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallObjectVipgroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallObjectVipgroup.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Interface name.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Member VIP objects of the group.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * VIP group name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a FirewallObjectVipgroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallObjectVipgroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallObjectVipgroupArgs | FirewallObjectVipgroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallObjectVipgroupState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as FirewallObjectVipgroupArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallObjectVipgroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallObjectVipgroup resources.
 */
export interface FirewallObjectVipgroupState {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Member VIP objects of the group.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VIP group name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallObjectVipgroup resource.
 */
export interface FirewallObjectVipgroupArgs {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Member VIP objects of the group.
     */
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VIP group name.
     */
    name?: pulumi.Input<string>;
}
