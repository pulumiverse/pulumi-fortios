// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch SNMP system information globally. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * SwitchController SnmpSysinfo can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/switchcontrollerSnmpsysinfo:SwitchcontrollerSnmpsysinfo labelname SwitchControllerSnmpSysinfo
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchcontrollerSnmpsysinfo:SwitchcontrollerSnmpsysinfo labelname SwitchControllerSnmpSysinfo
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchcontrollerSnmpsysinfo extends pulumi.CustomResource {
    /**
     * Get an existing SwitchcontrollerSnmpsysinfo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchcontrollerSnmpsysinfoState, opts?: pulumi.CustomResourceOptions): SwitchcontrollerSnmpsysinfo {
        return new SwitchcontrollerSnmpsysinfo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchcontrollerSnmpsysinfo:SwitchcontrollerSnmpsysinfo';

    /**
     * Returns true if the given object is an instance of SwitchcontrollerSnmpsysinfo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchcontrollerSnmpsysinfo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchcontrollerSnmpsysinfo.__pulumiType;
    }

    /**
     * Contact information.
     */
    public readonly contactInfo!: pulumi.Output<string>;
    /**
     * System description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Local SNMP engine ID string (max 24 char).
     */
    public readonly engineId!: pulumi.Output<string>;
    /**
     * System location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Enable/disable SNMP. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchcontrollerSnmpsysinfo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchcontrollerSnmpsysinfoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchcontrollerSnmpsysinfoArgs | SwitchcontrollerSnmpsysinfoState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchcontrollerSnmpsysinfoState | undefined;
            resourceInputs["contactInfo"] = state ? state.contactInfo : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["engineId"] = state ? state.engineId : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchcontrollerSnmpsysinfoArgs | undefined;
            resourceInputs["contactInfo"] = args ? args.contactInfo : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engineId"] = args ? args.engineId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchcontrollerSnmpsysinfo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchcontrollerSnmpsysinfo resources.
 */
export interface SwitchcontrollerSnmpsysinfoState {
    /**
     * Contact information.
     */
    contactInfo?: pulumi.Input<string>;
    /**
     * System description.
     */
    description?: pulumi.Input<string>;
    /**
     * Local SNMP engine ID string (max 24 char).
     */
    engineId?: pulumi.Input<string>;
    /**
     * System location.
     */
    location?: pulumi.Input<string>;
    /**
     * Enable/disable SNMP. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchcontrollerSnmpsysinfo resource.
 */
export interface SwitchcontrollerSnmpsysinfoArgs {
    /**
     * Contact information.
     */
    contactInfo?: pulumi.Input<string>;
    /**
     * System description.
     */
    description?: pulumi.Input<string>;
    /**
     * Local SNMP engine ID string (max 24 char).
     */
    engineId?: pulumi.Input<string>;
    /**
     * System location.
     */
    location?: pulumi.Input<string>;
    /**
     * Enable/disable SNMP. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
