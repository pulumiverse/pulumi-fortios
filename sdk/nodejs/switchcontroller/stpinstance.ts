// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure FortiSwitch multiple spanning tree protocol (MSTP) instances. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * SwitchController StpInstance can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/stpinstance:Stpinstance labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/stpinstance:Stpinstance labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Stpinstance extends pulumi.CustomResource {
    /**
     * Get an existing Stpinstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StpinstanceState, opts?: pulumi.CustomResourceOptions): Stpinstance {
        return new Stpinstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/stpinstance:Stpinstance';

    /**
     * Returns true if the given object is an instance of Stpinstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stpinstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stpinstance.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Instance ID.
     */
    public readonly fosid!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
     */
    public readonly vlanRanges!: pulumi.Output<outputs.switchcontroller.StpinstanceVlanRange[] | undefined>;

    /**
     * Create a Stpinstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StpinstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StpinstanceArgs | StpinstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StpinstanceState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlanRanges"] = state ? state.vlanRanges : undefined;
        } else {
            const args = argsOrState as StpinstanceArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlanRanges"] = args ? args.vlanRanges : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stpinstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stpinstance resources.
 */
export interface StpinstanceState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
     */
    vlanRanges?: pulumi.Input<pulumi.Input<inputs.switchcontroller.StpinstanceVlanRange>[]>;
}

/**
 * The set of arguments for constructing a Stpinstance resource.
 */
export interface StpinstanceArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Configure VLAN range for STP instance. The structure of `vlanRange` block is documented below.
     */
    vlanRanges?: pulumi.Input<pulumi.Input<inputs.switchcontroller.StpinstanceVlanRange>[]>;
}
