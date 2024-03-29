// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FortiSwitch sFlow.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.switchcontroller.Sflow("trname", {
 *     collectorIp: "0.0.0.0",
 *     collectorPort: 6343,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * SwitchController Sflow can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/sflow:Sflow labelname SwitchControllerSflow
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/sflow:Sflow labelname SwitchControllerSflow
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Sflow extends pulumi.CustomResource {
    /**
     * Get an existing Sflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SflowState, opts?: pulumi.CustomResourceOptions): Sflow {
        return new Sflow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/sflow:Sflow';

    /**
     * Returns true if the given object is an instance of Sflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sflow.__pulumiType;
    }

    /**
     * Collector IP.
     */
    public readonly collectorIp!: pulumi.Output<string>;
    /**
     * SFlow collector port (0 - 65535).
     */
    public readonly collectorPort!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Sflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SflowArgs | SflowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SflowState | undefined;
            resourceInputs["collectorIp"] = state ? state.collectorIp : undefined;
            resourceInputs["collectorPort"] = state ? state.collectorPort : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SflowArgs | undefined;
            if ((!args || args.collectorIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'collectorIp'");
            }
            resourceInputs["collectorIp"] = args ? args.collectorIp : undefined;
            resourceInputs["collectorPort"] = args ? args.collectorPort : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Sflow resources.
 */
export interface SflowState {
    /**
     * Collector IP.
     */
    collectorIp?: pulumi.Input<string>;
    /**
     * SFlow collector port (0 - 65535).
     */
    collectorPort?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Sflow resource.
 */
export interface SflowArgs {
    /**
     * Collector IP.
     */
    collectorIp: pulumi.Input<string>;
    /**
     * SFlow collector port (0 - 65535).
     */
    collectorPort?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
