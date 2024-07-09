// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure Authentication Rules.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.authentication.Rule("trname", {
 *     ipBased: "enable",
 *     protocol: "ftp",
 *     status: "enable",
 *     transactionBased: "disable",
 *     webAuthCookie: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Authentication Rule can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:authentication/rule:Rule labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:authentication/rule:Rule labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:authentication/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * Select an active authentication method.
     */
    public readonly activeAuthMethod!: pulumi.Output<string>;
    /**
     * Enable/disable to use device certificate as authentication cookie (default = enable). Valid values: `enable`, `disable`.
     */
    public readonly certAuthCookie!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Depth to allow CORS access (default = 3).
     */
    public readonly corsDepth!: pulumi.Output<number>;
    /**
     * Enable/disable allowance of CORS access (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly corsStateful!: pulumi.Output<string>;
    /**
     * Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
     */
    public readonly dstaddr6s!: pulumi.Output<outputs.authentication.RuleDstaddr6[] | undefined>;
    /**
     * Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
     */
    public readonly dstaddrs!: pulumi.Output<outputs.authentication.RuleDstaddr[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
     */
    public readonly ipBased!: pulumi.Output<string>;
    /**
     * Authentication rule name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Authentication is required for the selected protocol (default = http). Valid values: `http`, `ftp`, `socks`, `ssh`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
     */
    public readonly srcaddr6s!: pulumi.Output<outputs.authentication.RuleSrcaddr6[] | undefined>;
    /**
     * Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
     */
    public readonly srcaddrs!: pulumi.Output<outputs.authentication.RuleSrcaddr[] | undefined>;
    /**
     * Incoming (ingress) interface. The structure of `srcintf` block is documented below.
     */
    public readonly srcintfs!: pulumi.Output<outputs.authentication.RuleSrcintf[] | undefined>;
    /**
     * Select a single-sign on (SSO) authentication method.
     */
    public readonly ssoAuthMethod!: pulumi.Output<string>;
    /**
     * Enable/disable this authentication rule. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly transactionBased!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly webAuthCookie!: pulumi.Output<string>;
    /**
     * Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
     */
    public readonly webPortal!: pulumi.Output<string>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleState | undefined;
            resourceInputs["activeAuthMethod"] = state ? state.activeAuthMethod : undefined;
            resourceInputs["certAuthCookie"] = state ? state.certAuthCookie : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["corsDepth"] = state ? state.corsDepth : undefined;
            resourceInputs["corsStateful"] = state ? state.corsStateful : undefined;
            resourceInputs["dstaddr6s"] = state ? state.dstaddr6s : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["ipBased"] = state ? state.ipBased : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["srcaddr6s"] = state ? state.srcaddr6s : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["srcintfs"] = state ? state.srcintfs : undefined;
            resourceInputs["ssoAuthMethod"] = state ? state.ssoAuthMethod : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["transactionBased"] = state ? state.transactionBased : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["webAuthCookie"] = state ? state.webAuthCookie : undefined;
            resourceInputs["webPortal"] = state ? state.webPortal : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            resourceInputs["activeAuthMethod"] = args ? args.activeAuthMethod : undefined;
            resourceInputs["certAuthCookie"] = args ? args.certAuthCookie : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["corsDepth"] = args ? args.corsDepth : undefined;
            resourceInputs["corsStateful"] = args ? args.corsStateful : undefined;
            resourceInputs["dstaddr6s"] = args ? args.dstaddr6s : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["ipBased"] = args ? args.ipBased : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["srcaddr6s"] = args ? args.srcaddr6s : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["srcintfs"] = args ? args.srcintfs : undefined;
            resourceInputs["ssoAuthMethod"] = args ? args.ssoAuthMethod : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["transactionBased"] = args ? args.transactionBased : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["webAuthCookie"] = args ? args.webAuthCookie : undefined;
            resourceInputs["webPortal"] = args ? args.webPortal : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * Select an active authentication method.
     */
    activeAuthMethod?: pulumi.Input<string>;
    /**
     * Enable/disable to use device certificate as authentication cookie (default = enable). Valid values: `enable`, `disable`.
     */
    certAuthCookie?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Depth to allow CORS access (default = 3).
     */
    corsDepth?: pulumi.Input<number>;
    /**
     * Enable/disable allowance of CORS access (default = disable). Valid values: `enable`, `disable`.
     */
    corsStateful?: pulumi.Input<string>;
    /**
     * Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
     */
    dstaddr6s?: pulumi.Input<pulumi.Input<inputs.authentication.RuleDstaddr6>[]>;
    /**
     * Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.authentication.RuleDstaddr>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
     */
    ipBased?: pulumi.Input<string>;
    /**
     * Authentication rule name.
     */
    name?: pulumi.Input<string>;
    /**
     * Authentication is required for the selected protocol (default = http). Valid values: `http`, `ftp`, `socks`, `ssh`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
     */
    srcaddr6s?: pulumi.Input<pulumi.Input<inputs.authentication.RuleSrcaddr6>[]>;
    /**
     * Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.authentication.RuleSrcaddr>[]>;
    /**
     * Incoming (ingress) interface. The structure of `srcintf` block is documented below.
     */
    srcintfs?: pulumi.Input<pulumi.Input<inputs.authentication.RuleSrcintf>[]>;
    /**
     * Select a single-sign on (SSO) authentication method.
     */
    ssoAuthMethod?: pulumi.Input<string>;
    /**
     * Enable/disable this authentication rule. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
     */
    transactionBased?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
     */
    webAuthCookie?: pulumi.Input<string>;
    /**
     * Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
     */
    webPortal?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * Select an active authentication method.
     */
    activeAuthMethod?: pulumi.Input<string>;
    /**
     * Enable/disable to use device certificate as authentication cookie (default = enable). Valid values: `enable`, `disable`.
     */
    certAuthCookie?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Depth to allow CORS access (default = 3).
     */
    corsDepth?: pulumi.Input<number>;
    /**
     * Enable/disable allowance of CORS access (default = disable). Valid values: `enable`, `disable`.
     */
    corsStateful?: pulumi.Input<string>;
    /**
     * Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
     */
    dstaddr6s?: pulumi.Input<pulumi.Input<inputs.authentication.RuleDstaddr6>[]>;
    /**
     * Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.authentication.RuleDstaddr>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
     */
    ipBased?: pulumi.Input<string>;
    /**
     * Authentication rule name.
     */
    name?: pulumi.Input<string>;
    /**
     * Authentication is required for the selected protocol (default = http). Valid values: `http`, `ftp`, `socks`, `ssh`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
     */
    srcaddr6s?: pulumi.Input<pulumi.Input<inputs.authentication.RuleSrcaddr6>[]>;
    /**
     * Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.authentication.RuleSrcaddr>[]>;
    /**
     * Incoming (ingress) interface. The structure of `srcintf` block is documented below.
     */
    srcintfs?: pulumi.Input<pulumi.Input<inputs.authentication.RuleSrcintf>[]>;
    /**
     * Select a single-sign on (SSO) authentication method.
     */
    ssoAuthMethod?: pulumi.Input<string>;
    /**
     * Enable/disable this authentication rule. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
     */
    transactionBased?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
     */
    webAuthCookie?: pulumi.Input<string>;
    /**
     * Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
     */
    webPortal?: pulumi.Input<string>;
}
