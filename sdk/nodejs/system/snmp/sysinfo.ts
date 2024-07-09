// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * SNMP system info configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.snmp.Sysinfo("trname", {
 *     status: "disable",
 *     trapHighCpuThreshold: 80,
 *     trapLogFullThreshold: 90,
 *     trapLowMemoryThreshold: 80,
 * });
 * ```
 *
 * ## Import
 *
 * SystemSnmp Sysinfo can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/snmp/sysinfo:Sysinfo labelname SystemSnmpSysinfo
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/snmp/sysinfo:Sysinfo labelname SystemSnmpSysinfo
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Sysinfo extends pulumi.CustomResource {
    /**
     * Get an existing Sysinfo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SysinfoState, opts?: pulumi.CustomResourceOptions): Sysinfo {
        return new Sysinfo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/snmp/sysinfo:Sysinfo';

    /**
     * Returns true if the given object is an instance of Sysinfo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sysinfo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sysinfo.__pulumiType;
    }

    /**
     * Enable/disable allowance of appending VDOM or interface index in some RFC tables. Valid values: `enable`, `disable`.
     */
    public readonly appendIndex!: pulumi.Output<string>;
    /**
     * Contact information.
     */
    public readonly contactInfo!: pulumi.Output<string | undefined>;
    /**
     * System description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Local SNMP engineID string. On FortiOS versions 6.2.0-7.0.0: maximum 24 characters. On FortiOS versions >= 7.0.1: maximum 27 characters.
     */
    public readonly engineId!: pulumi.Output<string>;
    /**
     * Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
     */
    public readonly engineIdType!: pulumi.Output<string>;
    /**
     * System location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable SNMP. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Free memory usage when trap is sent.
     */
    public readonly trapFreeMemoryThreshold!: pulumi.Output<number>;
    /**
     * Freeable memory usage when trap is sent.
     */
    public readonly trapFreeableMemoryThreshold!: pulumi.Output<number>;
    /**
     * CPU usage when trap is sent.
     */
    public readonly trapHighCpuThreshold!: pulumi.Output<number>;
    /**
     * Log disk usage when trap is sent.
     */
    public readonly trapLogFullThreshold!: pulumi.Output<number>;
    /**
     * Memory usage when trap is sent.
     */
    public readonly trapLowMemoryThreshold!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Sysinfo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SysinfoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SysinfoArgs | SysinfoState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SysinfoState | undefined;
            resourceInputs["appendIndex"] = state ? state.appendIndex : undefined;
            resourceInputs["contactInfo"] = state ? state.contactInfo : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["engineId"] = state ? state.engineId : undefined;
            resourceInputs["engineIdType"] = state ? state.engineIdType : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trapFreeMemoryThreshold"] = state ? state.trapFreeMemoryThreshold : undefined;
            resourceInputs["trapFreeableMemoryThreshold"] = state ? state.trapFreeableMemoryThreshold : undefined;
            resourceInputs["trapHighCpuThreshold"] = state ? state.trapHighCpuThreshold : undefined;
            resourceInputs["trapLogFullThreshold"] = state ? state.trapLogFullThreshold : undefined;
            resourceInputs["trapLowMemoryThreshold"] = state ? state.trapLowMemoryThreshold : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SysinfoArgs | undefined;
            resourceInputs["appendIndex"] = args ? args.appendIndex : undefined;
            resourceInputs["contactInfo"] = args ? args.contactInfo : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engineId"] = args ? args.engineId : undefined;
            resourceInputs["engineIdType"] = args ? args.engineIdType : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["trapFreeMemoryThreshold"] = args ? args.trapFreeMemoryThreshold : undefined;
            resourceInputs["trapFreeableMemoryThreshold"] = args ? args.trapFreeableMemoryThreshold : undefined;
            resourceInputs["trapHighCpuThreshold"] = args ? args.trapHighCpuThreshold : undefined;
            resourceInputs["trapLogFullThreshold"] = args ? args.trapLogFullThreshold : undefined;
            resourceInputs["trapLowMemoryThreshold"] = args ? args.trapLowMemoryThreshold : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sysinfo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Sysinfo resources.
 */
export interface SysinfoState {
    /**
     * Enable/disable allowance of appending VDOM or interface index in some RFC tables. Valid values: `enable`, `disable`.
     */
    appendIndex?: pulumi.Input<string>;
    /**
     * Contact information.
     */
    contactInfo?: pulumi.Input<string>;
    /**
     * System description.
     */
    description?: pulumi.Input<string>;
    /**
     * Local SNMP engineID string. On FortiOS versions 6.2.0-7.0.0: maximum 24 characters. On FortiOS versions >= 7.0.1: maximum 27 characters.
     */
    engineId?: pulumi.Input<string>;
    /**
     * Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
     */
    engineIdType?: pulumi.Input<string>;
    /**
     * System location.
     */
    location?: pulumi.Input<string>;
    /**
     * Enable/disable SNMP. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Free memory usage when trap is sent.
     */
    trapFreeMemoryThreshold?: pulumi.Input<number>;
    /**
     * Freeable memory usage when trap is sent.
     */
    trapFreeableMemoryThreshold?: pulumi.Input<number>;
    /**
     * CPU usage when trap is sent.
     */
    trapHighCpuThreshold?: pulumi.Input<number>;
    /**
     * Log disk usage when trap is sent.
     */
    trapLogFullThreshold?: pulumi.Input<number>;
    /**
     * Memory usage when trap is sent.
     */
    trapLowMemoryThreshold?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Sysinfo resource.
 */
export interface SysinfoArgs {
    /**
     * Enable/disable allowance of appending VDOM or interface index in some RFC tables. Valid values: `enable`, `disable`.
     */
    appendIndex?: pulumi.Input<string>;
    /**
     * Contact information.
     */
    contactInfo?: pulumi.Input<string>;
    /**
     * System description.
     */
    description?: pulumi.Input<string>;
    /**
     * Local SNMP engineID string. On FortiOS versions 6.2.0-7.0.0: maximum 24 characters. On FortiOS versions >= 7.0.1: maximum 27 characters.
     */
    engineId?: pulumi.Input<string>;
    /**
     * Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
     */
    engineIdType?: pulumi.Input<string>;
    /**
     * System location.
     */
    location?: pulumi.Input<string>;
    /**
     * Enable/disable SNMP. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Free memory usage when trap is sent.
     */
    trapFreeMemoryThreshold?: pulumi.Input<number>;
    /**
     * Freeable memory usage when trap is sent.
     */
    trapFreeableMemoryThreshold?: pulumi.Input<number>;
    /**
     * CPU usage when trap is sent.
     */
    trapHighCpuThreshold?: pulumi.Input<number>;
    /**
     * Log disk usage when trap is sent.
     */
    trapLogFullThreshold?: pulumi.Input<number>;
    /**
     * Memory usage when trap is sent.
     */
    trapLowMemoryThreshold?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
