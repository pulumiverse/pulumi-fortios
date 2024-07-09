// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../../types/input";
import * as outputs from "../../../types/output";
import * as utilities from "../../../utilities";

/**
 * SSL-VPN host check software.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.vpn.ssl.web.Hostchecksoftware("trname", {
 *     osType: "windows",
 *     type: "fw",
 * });
 * ```
 *
 * ## Import
 *
 * VpnSslWeb HostCheckSoftware can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/ssl/web/hostchecksoftware:Hostchecksoftware labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/ssl/web/hostchecksoftware:Hostchecksoftware labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Hostchecksoftware extends pulumi.CustomResource {
    /**
     * Get an existing Hostchecksoftware resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostchecksoftwareState, opts?: pulumi.CustomResourceOptions): Hostchecksoftware {
        return new Hostchecksoftware(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/ssl/web/hostchecksoftware:Hostchecksoftware';

    /**
     * Returns true if the given object is an instance of Hostchecksoftware.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Hostchecksoftware {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Hostchecksoftware.__pulumiType;
    }

    /**
     * Check item list. The structure of `checkItemList` block is documented below.
     */
    public readonly checkItemLists!: pulumi.Output<outputs.vpn.ssl.web.HostchecksoftwareCheckItemList[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Globally unique ID.
     */
    public readonly guid!: pulumi.Output<string>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * OS type. Valid values: `windows`, `macos`.
     */
    public readonly osType!: pulumi.Output<string>;
    /**
     * Type. Valid values: `av`, `fw`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Version.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Hostchecksoftware resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: HostchecksoftwareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostchecksoftwareArgs | HostchecksoftwareState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostchecksoftwareState | undefined;
            resourceInputs["checkItemLists"] = state ? state.checkItemLists : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["guid"] = state ? state.guid : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["osType"] = state ? state.osType : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as HostchecksoftwareArgs | undefined;
            resourceInputs["checkItemLists"] = args ? args.checkItemLists : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["guid"] = args ? args.guid : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["osType"] = args ? args.osType : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Hostchecksoftware.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Hostchecksoftware resources.
 */
export interface HostchecksoftwareState {
    /**
     * Check item list. The structure of `checkItemList` block is documented below.
     */
    checkItemLists?: pulumi.Input<pulumi.Input<inputs.vpn.ssl.web.HostchecksoftwareCheckItemList>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Globally unique ID.
     */
    guid?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * OS type. Valid values: `windows`, `macos`.
     */
    osType?: pulumi.Input<string>;
    /**
     * Type. Valid values: `av`, `fw`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Version.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Hostchecksoftware resource.
 */
export interface HostchecksoftwareArgs {
    /**
     * Check item list. The structure of `checkItemList` block is documented below.
     */
    checkItemLists?: pulumi.Input<pulumi.Input<inputs.vpn.ssl.web.HostchecksoftwareCheckItemList>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Globally unique ID.
     */
    guid?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * OS type. Valid values: `windows`, `macos`.
     */
    osType?: pulumi.Input<string>;
    /**
     * Type. Valid values: `av`, `fw`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Version.
     */
    version?: pulumi.Input<string>;
}
