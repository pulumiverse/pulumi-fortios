// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Resource to sort firewall security policies by policyid or policy name, in ascending or descending order.
 *
 * ## Example Usage
 *
 * ### Example1
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test = new fortios.firewall.SecurityPolicysort("test", {
 *     sortby: "policyid",
 *     sortdirection: "descending",
 * });
 * export const policylistAfterApply = test.statePolicyLists;
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SecurityPolicysort extends pulumi.CustomResource {
    /**
     * Get an existing SecurityPolicysort resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityPolicysortState, opts?: pulumi.CustomResourceOptions): SecurityPolicysort {
        return new SecurityPolicysort(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/securityPolicysort:SecurityPolicysort';

    /**
     * Returns true if the given object is an instance of SecurityPolicysort.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityPolicysort {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityPolicysort.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
     */
    public readonly forceRecreate!: pulumi.Output<string | undefined>;
    /**
     * Sort security policies by the value, it currently supports "policyid" and "name".
     */
    public readonly sortby!: pulumi.Output<string>;
    /**
     * Sort dirction, supports "ascending" and "descending".
     */
    public readonly sortdirection!: pulumi.Output<string>;
    public /*out*/ readonly statePolicyLists!: pulumi.Output<outputs.firewall.SecurityPolicysortStatePolicyList[]>;
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SecurityPolicysort resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityPolicysortArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityPolicysortArgs | SecurityPolicysortState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityPolicysortState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["forceRecreate"] = state ? state.forceRecreate : undefined;
            resourceInputs["sortby"] = state ? state.sortby : undefined;
            resourceInputs["sortdirection"] = state ? state.sortdirection : undefined;
            resourceInputs["statePolicyLists"] = state ? state.statePolicyLists : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SecurityPolicysortArgs | undefined;
            if ((!args || args.sortby === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sortby'");
            }
            if ((!args || args.sortdirection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sortdirection'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["forceRecreate"] = args ? args.forceRecreate : undefined;
            resourceInputs["sortby"] = args ? args.sortby : undefined;
            resourceInputs["sortdirection"] = args ? args.sortdirection : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["statePolicyLists"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityPolicysort.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityPolicysort resources.
 */
export interface SecurityPolicysortState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
     */
    forceRecreate?: pulumi.Input<string>;
    /**
     * Sort security policies by the value, it currently supports "policyid" and "name".
     */
    sortby?: pulumi.Input<string>;
    /**
     * Sort dirction, supports "ascending" and "descending".
     */
    sortdirection?: pulumi.Input<string>;
    statePolicyLists?: pulumi.Input<pulumi.Input<inputs.firewall.SecurityPolicysortStatePolicyList>[]>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityPolicysort resource.
 */
export interface SecurityPolicysortArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
     */
    forceRecreate?: pulumi.Input<string>;
    /**
     * Sort security policies by the value, it currently supports "policyid" and "name".
     */
    sortby: pulumi.Input<string>;
    /**
     * Sort dirction, supports "ascending" and "descending".
     */
    sortdirection: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
