// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure AP local configuration profiles. Applies to FortiOS Version `>= 6.4.0`.
 *
 * ## Import
 *
 * WirelessController ApcfgProfile can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:wirelesscontroller/apcfgprofile:Apcfgprofile labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:wirelesscontroller/apcfgprofile:Apcfgprofile labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Apcfgprofile extends pulumi.CustomResource {
    /**
     * Get an existing Apcfgprofile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApcfgprofileState, opts?: pulumi.CustomResourceOptions): Apcfgprofile {
        return new Apcfgprofile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wirelesscontroller/apcfgprofile:Apcfgprofile';

    /**
     * Returns true if the given object is an instance of Apcfgprofile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Apcfgprofile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Apcfgprofile.__pulumiType;
    }

    /**
     * IP address of the validation controller that AP must be able to join after applying AP local configuration.
     */
    public readonly acIp!: pulumi.Output<string>;
    /**
     * Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
     */
    public readonly acPort!: pulumi.Output<number>;
    /**
     * Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
     */
    public readonly acTimer!: pulumi.Output<number>;
    /**
     * Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
     */
    public readonly acType!: pulumi.Output<string>;
    /**
     * FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
     */
    public readonly apFamily!: pulumi.Output<string>;
    /**
     * AP local configuration command list. The structure of `commandList` block is documented below.
     */
    public readonly commandLists!: pulumi.Output<outputs.wirelesscontroller.ApcfgprofileCommandList[] | undefined>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * AP local configuration profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Apcfgprofile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApcfgprofileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApcfgprofileArgs | ApcfgprofileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApcfgprofileState | undefined;
            resourceInputs["acIp"] = state ? state.acIp : undefined;
            resourceInputs["acPort"] = state ? state.acPort : undefined;
            resourceInputs["acTimer"] = state ? state.acTimer : undefined;
            resourceInputs["acType"] = state ? state.acType : undefined;
            resourceInputs["apFamily"] = state ? state.apFamily : undefined;
            resourceInputs["commandLists"] = state ? state.commandLists : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ApcfgprofileArgs | undefined;
            resourceInputs["acIp"] = args ? args.acIp : undefined;
            resourceInputs["acPort"] = args ? args.acPort : undefined;
            resourceInputs["acTimer"] = args ? args.acTimer : undefined;
            resourceInputs["acType"] = args ? args.acType : undefined;
            resourceInputs["apFamily"] = args ? args.apFamily : undefined;
            resourceInputs["commandLists"] = args ? args.commandLists : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Apcfgprofile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Apcfgprofile resources.
 */
export interface ApcfgprofileState {
    /**
     * IP address of the validation controller that AP must be able to join after applying AP local configuration.
     */
    acIp?: pulumi.Input<string>;
    /**
     * Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
     */
    acPort?: pulumi.Input<number>;
    /**
     * Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
     */
    acTimer?: pulumi.Input<number>;
    /**
     * Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
     */
    acType?: pulumi.Input<string>;
    /**
     * FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
     */
    apFamily?: pulumi.Input<string>;
    /**
     * AP local configuration command list. The structure of `commandList` block is documented below.
     */
    commandLists?: pulumi.Input<pulumi.Input<inputs.wirelesscontroller.ApcfgprofileCommandList>[]>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * AP local configuration profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Apcfgprofile resource.
 */
export interface ApcfgprofileArgs {
    /**
     * IP address of the validation controller that AP must be able to join after applying AP local configuration.
     */
    acIp?: pulumi.Input<string>;
    /**
     * Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
     */
    acPort?: pulumi.Input<number>;
    /**
     * Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
     */
    acTimer?: pulumi.Input<number>;
    /**
     * Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
     */
    acType?: pulumi.Input<string>;
    /**
     * FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
     */
    apFamily?: pulumi.Input<string>;
    /**
     * AP local configuration command list. The structure of `commandList` block is documented below.
     */
    commandLists?: pulumi.Input<pulumi.Input<inputs.wirelesscontroller.ApcfgprofileCommandList>[]>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * AP local configuration profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
