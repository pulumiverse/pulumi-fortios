// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Configure FortiGuard Web Filter administrative overrides.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.filter.web.Override("trname", {
 *     expires: "2021/03/16 19:18:25",
 *     fosid: 1,
 *     ip: "69.101.119.0",
 *     ip6: "4565:7700::",
 *     newProfile: "monitor-all",
 *     oldProfile: "default",
 *     scope: "user",
 *     status: "disable",
 *     user: "Eew",
 * });
 * ```
 *
 * ## Import
 *
 * Webfilter Override can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:filter/web/override:Override labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:filter/web/override:Override labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Override extends pulumi.CustomResource {
    /**
     * Get an existing Override resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OverrideState, opts?: pulumi.CustomResourceOptions): Override {
        return new Override(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:filter/web/override:Override';

    /**
     * Returns true if the given object is an instance of Override.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Override {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Override.__pulumiType;
    }

    /**
     * Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
     */
    public readonly expires!: pulumi.Output<string>;
    /**
     * Override rule ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Initiating user of override (read-only setting).
     */
    public readonly initiator!: pulumi.Output<string>;
    /**
     * IPv4 address which the override applies.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * IPv6 address which the override applies.
     */
    public readonly ip6!: pulumi.Output<string>;
    /**
     * Name of the new web filter profile used by the override.
     */
    public readonly newProfile!: pulumi.Output<string>;
    /**
     * Name of the web filter profile which the override applies.
     */
    public readonly oldProfile!: pulumi.Output<string>;
    /**
     * Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * Enable/disable override rule. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Name of the user which the override applies.
     */
    public readonly user!: pulumi.Output<string>;
    /**
     * Specify the user group for which the override applies.
     */
    public readonly userGroup!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Override resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OverrideArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OverrideArgs | OverrideState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OverrideState | undefined;
            resourceInputs["expires"] = state ? state.expires : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["initiator"] = state ? state.initiator : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["ip6"] = state ? state.ip6 : undefined;
            resourceInputs["newProfile"] = state ? state.newProfile : undefined;
            resourceInputs["oldProfile"] = state ? state.oldProfile : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["userGroup"] = state ? state.userGroup : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as OverrideArgs | undefined;
            if ((!args || args.expires === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expires'");
            }
            if ((!args || args.newProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'newProfile'");
            }
            if ((!args || args.oldProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oldProfile'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["expires"] = args ? args.expires : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["initiator"] = args ? args.initiator : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["ip6"] = args ? args.ip6 : undefined;
            resourceInputs["newProfile"] = args ? args.newProfile : undefined;
            resourceInputs["oldProfile"] = args ? args.oldProfile : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["userGroup"] = args ? args.userGroup : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Override.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Override resources.
 */
export interface OverrideState {
    /**
     * Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
     */
    expires?: pulumi.Input<string>;
    /**
     * Override rule ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Initiating user of override (read-only setting).
     */
    initiator?: pulumi.Input<string>;
    /**
     * IPv4 address which the override applies.
     */
    ip?: pulumi.Input<string>;
    /**
     * IPv6 address which the override applies.
     */
    ip6?: pulumi.Input<string>;
    /**
     * Name of the new web filter profile used by the override.
     */
    newProfile?: pulumi.Input<string>;
    /**
     * Name of the web filter profile which the override applies.
     */
    oldProfile?: pulumi.Input<string>;
    /**
     * Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
     */
    scope?: pulumi.Input<string>;
    /**
     * Enable/disable override rule. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Name of the user which the override applies.
     */
    user?: pulumi.Input<string>;
    /**
     * Specify the user group for which the override applies.
     */
    userGroup?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Override resource.
 */
export interface OverrideArgs {
    /**
     * Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
     */
    expires: pulumi.Input<string>;
    /**
     * Override rule ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Initiating user of override (read-only setting).
     */
    initiator?: pulumi.Input<string>;
    /**
     * IPv4 address which the override applies.
     */
    ip?: pulumi.Input<string>;
    /**
     * IPv6 address which the override applies.
     */
    ip6?: pulumi.Input<string>;
    /**
     * Name of the new web filter profile used by the override.
     */
    newProfile: pulumi.Input<string>;
    /**
     * Name of the web filter profile which the override applies.
     */
    oldProfile: pulumi.Input<string>;
    /**
     * Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
     */
    scope?: pulumi.Input<string>;
    /**
     * Enable/disable override rule. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Name of the user which the override applies.
     */
    user: pulumi.Input<string>;
    /**
     * Specify the user group for which the override applies.
     */
    userGroup?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
