// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure auto script.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const auto2 = new fortios.system.Autoscript("auto2", {
 *     interval: 1,
 *     outputSize: 10,
 *     repeat: 1,
 *     script: `config firewall address
 *     edit "111"
 *         set color 3
 *         set subnet 1.1.1.1 255.255.255.255
 *     next
 * end
 *
 * `,
 *     start: "auto",
 * });
 * ```
 *
 * ## Import
 *
 * System AutoScript can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/autoscript:Autoscript labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/autoscript:Autoscript labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Autoscript extends pulumi.CustomResource {
    /**
     * Get an existing Autoscript resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoscriptState, opts?: pulumi.CustomResourceOptions): Autoscript {
        return new Autoscript(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/autoscript:Autoscript';

    /**
     * Returns true if the given object is an instance of Autoscript.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Autoscript {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Autoscript.__pulumiType;
    }

    /**
     * Repeat interval in seconds.
     */
    public readonly interval!: pulumi.Output<number>;
    /**
     * Auto script name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of megabytes to limit script output to (10 - 1024, default = 10).
     */
    public readonly outputSize!: pulumi.Output<number>;
    /**
     * Number of times to repeat this script (0 = infinite).
     */
    public readonly repeat!: pulumi.Output<number>;
    /**
     * List of FortiOS CLI commands to repeat.
     */
    public readonly script!: pulumi.Output<string | undefined>;
    /**
     * Script starting mode. Valid values: `manual`, `auto`.
     */
    public readonly start!: pulumi.Output<string>;
    /**
     * Maximum running time for this script in seconds (0 = no timeout).
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Autoscript resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AutoscriptArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoscriptArgs | AutoscriptState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoscriptState | undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outputSize"] = state ? state.outputSize : undefined;
            resourceInputs["repeat"] = state ? state.repeat : undefined;
            resourceInputs["script"] = state ? state.script : undefined;
            resourceInputs["start"] = state ? state.start : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AutoscriptArgs | undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputSize"] = args ? args.outputSize : undefined;
            resourceInputs["repeat"] = args ? args.repeat : undefined;
            resourceInputs["script"] = args ? args.script : undefined;
            resourceInputs["start"] = args ? args.start : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Autoscript.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Autoscript resources.
 */
export interface AutoscriptState {
    /**
     * Repeat interval in seconds.
     */
    interval?: pulumi.Input<number>;
    /**
     * Auto script name.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of megabytes to limit script output to (10 - 1024, default = 10).
     */
    outputSize?: pulumi.Input<number>;
    /**
     * Number of times to repeat this script (0 = infinite).
     */
    repeat?: pulumi.Input<number>;
    /**
     * List of FortiOS CLI commands to repeat.
     */
    script?: pulumi.Input<string>;
    /**
     * Script starting mode. Valid values: `manual`, `auto`.
     */
    start?: pulumi.Input<string>;
    /**
     * Maximum running time for this script in seconds (0 = no timeout).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Autoscript resource.
 */
export interface AutoscriptArgs {
    /**
     * Repeat interval in seconds.
     */
    interval?: pulumi.Input<number>;
    /**
     * Auto script name.
     */
    name?: pulumi.Input<string>;
    /**
     * Number of megabytes to limit script output to (10 - 1024, default = 10).
     */
    outputSize?: pulumi.Input<number>;
    /**
     * Number of times to repeat this script (0 = infinite).
     */
    repeat?: pulumi.Input<number>;
    /**
     * List of FortiOS CLI commands to repeat.
     */
    script?: pulumi.Input<string>;
    /**
     * Script starting mode. Valid values: `manual`, `auto`.
     */
    start?: pulumi.Input<string>;
    /**
     * Maximum running time for this script in seconds (0 = no timeout).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
