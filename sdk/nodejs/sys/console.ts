// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure console.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.sys.Console("trname", {
 *     baudrate: "9600",
 *     login: "enable",
 *     mode: "line",
 *     output: "more",
 * });
 * ```
 *
 * ## Import
 *
 * System Console can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:sys/console:Console labelname SystemConsole
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:sys/console:Console labelname SystemConsole
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Console extends pulumi.CustomResource {
    /**
     * Get an existing Console resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConsoleState, opts?: pulumi.CustomResourceOptions): Console {
        return new Console(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/console:Console';

    /**
     * Returns true if the given object is an instance of Console.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Console {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Console.__pulumiType;
    }

    /**
     * Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
     */
    public readonly baudrate!: pulumi.Output<string>;
    /**
     * Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
     */
    public readonly fortiexplorer!: pulumi.Output<string>;
    /**
     * Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
     */
    public readonly login!: pulumi.Output<string>;
    /**
     * Console mode. Valid values: `batch`, `line`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Console output mode. Valid values: `standard`, `more`.
     */
    public readonly output!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Console resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConsoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConsoleArgs | ConsoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConsoleState | undefined;
            resourceInputs["baudrate"] = state ? state.baudrate : undefined;
            resourceInputs["fortiexplorer"] = state ? state.fortiexplorer : undefined;
            resourceInputs["login"] = state ? state.login : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["output"] = state ? state.output : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ConsoleArgs | undefined;
            resourceInputs["baudrate"] = args ? args.baudrate : undefined;
            resourceInputs["fortiexplorer"] = args ? args.fortiexplorer : undefined;
            resourceInputs["login"] = args ? args.login : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["output"] = args ? args.output : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Console.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Console resources.
 */
export interface ConsoleState {
    /**
     * Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
     */
    baudrate?: pulumi.Input<string>;
    /**
     * Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
     */
    fortiexplorer?: pulumi.Input<string>;
    /**
     * Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
     */
    login?: pulumi.Input<string>;
    /**
     * Console mode. Valid values: `batch`, `line`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Console output mode. Valid values: `standard`, `more`.
     */
    output?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Console resource.
 */
export interface ConsoleArgs {
    /**
     * Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
     */
    baudrate?: pulumi.Input<string>;
    /**
     * Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
     */
    fortiexplorer?: pulumi.Input<string>;
    /**
     * Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
     */
    login?: pulumi.Input<string>;
    /**
     * Console mode. Valid values: `batch`, `line`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Console output mode. Valid values: `standard`, `more`.
     */
    output?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
