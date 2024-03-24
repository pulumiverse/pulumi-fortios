// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class SecuritypolicySort extends pulumi.CustomResource {
    /**
     * Get an existing SecuritypolicySort resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecuritypolicySortState, opts?: pulumi.CustomResourceOptions): SecuritypolicySort {
        return new SecuritypolicySort(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/securitypolicySort:SecuritypolicySort';

    /**
     * Returns true if the given object is an instance of SecuritypolicySort.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecuritypolicySort {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecuritypolicySort.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly forceRecreate!: pulumi.Output<string>;
    public readonly manualOrders!: pulumi.Output<string[]>;
    public readonly sortby!: pulumi.Output<string>;
    public readonly sortdirection!: pulumi.Output<string>;
    public /*out*/ readonly statePolicyLists!: pulumi.Output<outputs.firewall.SecuritypolicySortStatePolicyList[]>;
    public readonly status!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SecuritypolicySort resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecuritypolicySortArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecuritypolicySortArgs | SecuritypolicySortState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecuritypolicySortState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["forceRecreate"] = state ? state.forceRecreate : undefined;
            resourceInputs["manualOrders"] = state ? state.manualOrders : undefined;
            resourceInputs["sortby"] = state ? state.sortby : undefined;
            resourceInputs["sortdirection"] = state ? state.sortdirection : undefined;
            resourceInputs["statePolicyLists"] = state ? state.statePolicyLists : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SecuritypolicySortArgs | undefined;
            if ((!args || args.sortby === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sortby'");
            }
            if ((!args || args.sortdirection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sortdirection'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["forceRecreate"] = args ? args.forceRecreate : undefined;
            resourceInputs["manualOrders"] = args ? args.manualOrders : undefined;
            resourceInputs["sortby"] = args ? args.sortby : undefined;
            resourceInputs["sortdirection"] = args ? args.sortdirection : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["statePolicyLists"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecuritypolicySort.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecuritypolicySort resources.
 */
export interface SecuritypolicySortState {
    comment?: pulumi.Input<string>;
    forceRecreate?: pulumi.Input<string>;
    manualOrders?: pulumi.Input<pulumi.Input<string>[]>;
    sortby?: pulumi.Input<string>;
    sortdirection?: pulumi.Input<string>;
    statePolicyLists?: pulumi.Input<pulumi.Input<inputs.firewall.SecuritypolicySortStatePolicyList>[]>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecuritypolicySort resource.
 */
export interface SecuritypolicySortArgs {
    comment?: pulumi.Input<string>;
    forceRecreate?: pulumi.Input<string>;
    manualOrders?: pulumi.Input<pulumi.Input<string>[]>;
    sortby: pulumi.Input<string>;
    sortdirection: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}