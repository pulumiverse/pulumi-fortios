// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure firewall authentication portals.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Authportal("trname", {
 *     groups: [{
 *         name: "Guest-group",
 *     }],
 *     portalAddr: "1.1.1.1",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall AuthPortal can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/authportal:Authportal labelname FirewallAuthPortal
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/authportal:Authportal labelname FirewallAuthPortal
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Authportal extends pulumi.CustomResource {
    /**
     * Get an existing Authportal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthportalState, opts?: pulumi.CustomResourceOptions): Authportal {
        return new Authportal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/authportal:Authportal';

    /**
     * Returns true if the given object is an instance of Authportal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Authportal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Authportal.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
     */
    public readonly groups!: pulumi.Output<outputs.firewall.AuthportalGroup[] | undefined>;
    /**
     * Name of the identity-based route that applies to this portal.
     */
    public readonly identityBasedRoute!: pulumi.Output<string>;
    /**
     * Address (or FQDN) of the authentication portal.
     */
    public readonly portalAddr!: pulumi.Output<string>;
    /**
     * IPv6 address (or FQDN) of authentication portal.
     */
    public readonly portalAddr6!: pulumi.Output<string>;
    /**
     * Enable/disable authentication by proxy daemon (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly proxyAuth!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Authportal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthportalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthportalArgs | AuthportalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthportalState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["identityBasedRoute"] = state ? state.identityBasedRoute : undefined;
            resourceInputs["portalAddr"] = state ? state.portalAddr : undefined;
            resourceInputs["portalAddr6"] = state ? state.portalAddr6 : undefined;
            resourceInputs["proxyAuth"] = state ? state.proxyAuth : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AuthportalArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["identityBasedRoute"] = args ? args.identityBasedRoute : undefined;
            resourceInputs["portalAddr"] = args ? args.portalAddr : undefined;
            resourceInputs["portalAddr6"] = args ? args.portalAddr6 : undefined;
            resourceInputs["proxyAuth"] = args ? args.proxyAuth : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Authportal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Authportal resources.
 */
export interface AuthportalState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.firewall.AuthportalGroup>[]>;
    /**
     * Name of the identity-based route that applies to this portal.
     */
    identityBasedRoute?: pulumi.Input<string>;
    /**
     * Address (or FQDN) of the authentication portal.
     */
    portalAddr?: pulumi.Input<string>;
    /**
     * IPv6 address (or FQDN) of authentication portal.
     */
    portalAddr6?: pulumi.Input<string>;
    /**
     * Enable/disable authentication by proxy daemon (default = disable). Valid values: `enable`, `disable`.
     */
    proxyAuth?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Authportal resource.
 */
export interface AuthportalArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.firewall.AuthportalGroup>[]>;
    /**
     * Name of the identity-based route that applies to this portal.
     */
    identityBasedRoute?: pulumi.Input<string>;
    /**
     * Address (or FQDN) of the authentication portal.
     */
    portalAddr?: pulumi.Input<string>;
    /**
     * IPv6 address (or FQDN) of authentication portal.
     */
    portalAddr6?: pulumi.Input<string>;
    /**
     * Enable/disable authentication by proxy daemon (default = disable). Valid values: `enable`, `disable`.
     */
    proxyAuth?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
