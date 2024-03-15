// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure user defined IPv6 local-in policies.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Localinpolicy6("trname", {
 *     action: "accept",
 *     dstaddrs: [{
 *         name: "all",
 *     }],
 *     intf: "port4",
 *     policyid: 1,
 *     schedule: "always",
 *     services: [{
 *         name: "ALL",
 *     }],
 *     srcaddrs: [{
 *         name: "all",
 *     }],
 *     status: "enable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Firewall LocalInPolicy6 can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/localinpolicy6:Localinpolicy6 labelname {{policyid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/localinpolicy6:Localinpolicy6 labelname {{policyid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Localinpolicy6 extends pulumi.CustomResource {
    /**
     * Get an existing Localinpolicy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Localinpolicy6State, opts?: pulumi.CustomResourceOptions): Localinpolicy6 {
        return new Localinpolicy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/localinpolicy6:Localinpolicy6';

    /**
     * Returns true if the given object is an instance of Localinpolicy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Localinpolicy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Localinpolicy6.__pulumiType;
    }

    /**
     * Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
     */
    public readonly dstaddrNegate!: pulumi.Output<string>;
    /**
     * Destination address object from available options. The structure of `dstaddr` block is documented below.
     */
    public readonly dstaddrs!: pulumi.Output<outputs.firewall.Localinpolicy6Dstaddr[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Incoming interface name from available options.
     */
    public readonly intf!: pulumi.Output<string>;
    /**
     * User defined local in policy ID.
     */
    public readonly policyid!: pulumi.Output<number>;
    /**
     * Schedule object from available options.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
     */
    public readonly serviceNegate!: pulumi.Output<string>;
    /**
     * Service object from available options. Separate names with a space. The structure of `service` block is documented below.
     */
    public readonly services!: pulumi.Output<outputs.firewall.Localinpolicy6Service[]>;
    /**
     * When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
     */
    public readonly srcaddrNegate!: pulumi.Output<string>;
    /**
     * Source address object from available options. The structure of `srcaddr` block is documented below.
     */
    public readonly srcaddrs!: pulumi.Output<outputs.firewall.Localinpolicy6Srcaddr[]>;
    /**
     * Enable/disable this local-in policy. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Localinpolicy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Localinpolicy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Localinpolicy6Args | Localinpolicy6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Localinpolicy6State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dstaddrNegate"] = state ? state.dstaddrNegate : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["intf"] = state ? state.intf : undefined;
            resourceInputs["policyid"] = state ? state.policyid : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["serviceNegate"] = state ? state.serviceNegate : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["srcaddrNegate"] = state ? state.srcaddrNegate : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Localinpolicy6Args | undefined;
            if ((!args || args.dstaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstaddrs'");
            }
            if ((!args || args.intf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'intf'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            if ((!args || args.services === undefined) && !opts.urn) {
                throw new Error("Missing required property 'services'");
            }
            if ((!args || args.srcaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcaddrs'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dstaddrNegate"] = args ? args.dstaddrNegate : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["intf"] = args ? args.intf : undefined;
            resourceInputs["policyid"] = args ? args.policyid : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["serviceNegate"] = args ? args.serviceNegate : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["srcaddrNegate"] = args ? args.srcaddrNegate : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Localinpolicy6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Localinpolicy6 resources.
 */
export interface Localinpolicy6State {
    /**
     * Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
     */
    dstaddrNegate?: pulumi.Input<string>;
    /**
     * Destination address object from available options. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.firewall.Localinpolicy6Dstaddr>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Incoming interface name from available options.
     */
    intf?: pulumi.Input<string>;
    /**
     * User defined local in policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Schedule object from available options.
     */
    schedule?: pulumi.Input<string>;
    /**
     * When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
     */
    serviceNegate?: pulumi.Input<string>;
    /**
     * Service object from available options. Separate names with a space. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.firewall.Localinpolicy6Service>[]>;
    /**
     * When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
     */
    srcaddrNegate?: pulumi.Input<string>;
    /**
     * Source address object from available options. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.firewall.Localinpolicy6Srcaddr>[]>;
    /**
     * Enable/disable this local-in policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Localinpolicy6 resource.
 */
export interface Localinpolicy6Args {
    /**
     * Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
     */
    dstaddrNegate?: pulumi.Input<string>;
    /**
     * Destination address object from available options. The structure of `dstaddr` block is documented below.
     */
    dstaddrs: pulumi.Input<pulumi.Input<inputs.firewall.Localinpolicy6Dstaddr>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Incoming interface name from available options.
     */
    intf: pulumi.Input<string>;
    /**
     * User defined local in policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Schedule object from available options.
     */
    schedule: pulumi.Input<string>;
    /**
     * When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
     */
    serviceNegate?: pulumi.Input<string>;
    /**
     * Service object from available options. Separate names with a space. The structure of `service` block is documented below.
     */
    services: pulumi.Input<pulumi.Input<inputs.firewall.Localinpolicy6Service>[]>;
    /**
     * When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
     */
    srcaddrNegate?: pulumi.Input<string>;
    /**
     * Source address object from available options. The structure of `srcaddr` block is documented below.
     */
    srcaddrs: pulumi.Input<pulumi.Input<inputs.firewall.Localinpolicy6Srcaddr>[]>;
    /**
     * Enable/disable this local-in policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}