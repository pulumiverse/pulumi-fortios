// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * FortiExtender extender profile configuration. Applies to FortiOS Version `>= 7.2.1`.
 *
 * ## Import
 *
 * ExtensionController ExtenderProfile can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:extensioncontroller/extenderprofile:Extenderprofile labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:extensioncontroller/extenderprofile:Extenderprofile labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Extenderprofile extends pulumi.CustomResource {
    /**
     * Get an existing Extenderprofile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtenderprofileState, opts?: pulumi.CustomResourceOptions): Extenderprofile {
        return new Extenderprofile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:extensioncontroller/extenderprofile:Extenderprofile';

    /**
     * Returns true if the given object is an instance of Extenderprofile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Extenderprofile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Extenderprofile.__pulumiType;
    }

    /**
     * Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
     */
    public readonly allowaccess!: pulumi.Output<string>;
    /**
     * FortiExtender LAN extension bandwidth limit (Mbps).
     */
    public readonly bandwidthLimit!: pulumi.Output<number>;
    /**
     * FortiExtender cellular configuration. The structure of `cellular` block is documented below.
     */
    public readonly cellular!: pulumi.Output<outputs.extensioncontroller.ExtenderprofileCellular>;
    /**
     * Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
     */
    public readonly enforceBandwidth!: pulumi.Output<string>;
    /**
     * Extension option. Valid values: `wan-extension`, `lan-extension`.
     */
    public readonly extension!: pulumi.Output<string>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
     */
    public readonly lanExtension!: pulumi.Output<outputs.extensioncontroller.ExtenderprofileLanExtension>;
    /**
     * Set the managed extender's administrator password.
     */
    public readonly loginPassword!: pulumi.Output<string | undefined>;
    /**
     * Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    public readonly loginPasswordChange!: pulumi.Output<string>;
    /**
     * Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
     */
    public readonly model!: pulumi.Output<string>;
    /**
     * FortiExtender profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Extenderprofile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExtenderprofileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtenderprofileArgs | ExtenderprofileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtenderprofileState | undefined;
            resourceInputs["allowaccess"] = state ? state.allowaccess : undefined;
            resourceInputs["bandwidthLimit"] = state ? state.bandwidthLimit : undefined;
            resourceInputs["cellular"] = state ? state.cellular : undefined;
            resourceInputs["enforceBandwidth"] = state ? state.enforceBandwidth : undefined;
            resourceInputs["extension"] = state ? state.extension : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["lanExtension"] = state ? state.lanExtension : undefined;
            resourceInputs["loginPassword"] = state ? state.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = state ? state.loginPasswordChange : undefined;
            resourceInputs["model"] = state ? state.model : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ExtenderprofileArgs | undefined;
            resourceInputs["allowaccess"] = args ? args.allowaccess : undefined;
            resourceInputs["bandwidthLimit"] = args ? args.bandwidthLimit : undefined;
            resourceInputs["cellular"] = args ? args.cellular : undefined;
            resourceInputs["enforceBandwidth"] = args ? args.enforceBandwidth : undefined;
            resourceInputs["extension"] = args ? args.extension : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["lanExtension"] = args ? args.lanExtension : undefined;
            resourceInputs["loginPassword"] = args ? args.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = args ? args.loginPasswordChange : undefined;
            resourceInputs["model"] = args ? args.model : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Extenderprofile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Extenderprofile resources.
 */
export interface ExtenderprofileState {
    /**
     * Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * FortiExtender LAN extension bandwidth limit (Mbps).
     */
    bandwidthLimit?: pulumi.Input<number>;
    /**
     * FortiExtender cellular configuration. The structure of `cellular` block is documented below.
     */
    cellular?: pulumi.Input<inputs.extensioncontroller.ExtenderprofileCellular>;
    /**
     * Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
     */
    enforceBandwidth?: pulumi.Input<string>;
    /**
     * Extension option. Valid values: `wan-extension`, `lan-extension`.
     */
    extension?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
     */
    lanExtension?: pulumi.Input<inputs.extensioncontroller.ExtenderprofileLanExtension>;
    /**
     * Set the managed extender's administrator password.
     */
    loginPassword?: pulumi.Input<string>;
    /**
     * Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    loginPasswordChange?: pulumi.Input<string>;
    /**
     * Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
     */
    model?: pulumi.Input<string>;
    /**
     * FortiExtender profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Extenderprofile resource.
 */
export interface ExtenderprofileArgs {
    /**
     * Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * FortiExtender LAN extension bandwidth limit (Mbps).
     */
    bandwidthLimit?: pulumi.Input<number>;
    /**
     * FortiExtender cellular configuration. The structure of `cellular` block is documented below.
     */
    cellular?: pulumi.Input<inputs.extensioncontroller.ExtenderprofileCellular>;
    /**
     * Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
     */
    enforceBandwidth?: pulumi.Input<string>;
    /**
     * Extension option. Valid values: `wan-extension`, `lan-extension`.
     */
    extension?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * FortiExtender lan extension configuration. The structure of `lanExtension` block is documented below.
     */
    lanExtension?: pulumi.Input<inputs.extensioncontroller.ExtenderprofileLanExtension>;
    /**
     * Set the managed extender's administrator password.
     */
    loginPassword?: pulumi.Input<string>;
    /**
     * Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    loginPasswordChange?: pulumi.Input<string>;
    /**
     * Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
     */
    model?: pulumi.Input<string>;
    /**
     * FortiExtender profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
