// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure WAN optimization settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.wanopt.Settings("trname", {
 *     autoDetectAlgorithm: "simple",
 *     hostId: "default-id",
 *     tunnelSslAlgorithm: "high",
 * });
 * ```
 *
 * ## Import
 *
 * Wanopt Settings can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:wanopt/settings:Settings labelname WanoptSettings
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:wanopt/settings:Settings labelname WanoptSettings
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Settings extends pulumi.CustomResource {
    /**
     * Get an existing Settings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SettingsState, opts?: pulumi.CustomResourceOptions): Settings {
        return new Settings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wanopt/settings:Settings';

    /**
     * Returns true if the given object is an instance of Settings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Settings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Settings.__pulumiType;
    }

    /**
     * Auto detection algorithms used in tunnel negotiations. Valid values: `simple`, `diff-req-resp`.
     */
    public readonly autoDetectAlgorithm!: pulumi.Output<string>;
    /**
     * Local host ID (must also be entered in the remote FortiGate's peer list).
     */
    public readonly hostId!: pulumi.Output<string>;
    /**
     * WANOpt tunnel optimization option. Valid values: `memory-usage`, `balanced`, `throughput`.
     */
    public readonly tunnelOptimization!: pulumi.Output<string>;
    /**
     * Relative strength of encryption algorithms accepted during tunnel negotiation. Valid values: `high`, `medium`, `low`.
     */
    public readonly tunnelSslAlgorithm!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Settings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SettingsArgs | SettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SettingsState | undefined;
            resourceInputs["autoDetectAlgorithm"] = state ? state.autoDetectAlgorithm : undefined;
            resourceInputs["hostId"] = state ? state.hostId : undefined;
            resourceInputs["tunnelOptimization"] = state ? state.tunnelOptimization : undefined;
            resourceInputs["tunnelSslAlgorithm"] = state ? state.tunnelSslAlgorithm : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SettingsArgs | undefined;
            if ((!args || args.hostId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostId'");
            }
            resourceInputs["autoDetectAlgorithm"] = args ? args.autoDetectAlgorithm : undefined;
            resourceInputs["hostId"] = args ? args.hostId : undefined;
            resourceInputs["tunnelOptimization"] = args ? args.tunnelOptimization : undefined;
            resourceInputs["tunnelSslAlgorithm"] = args ? args.tunnelSslAlgorithm : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Settings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Settings resources.
 */
export interface SettingsState {
    /**
     * Auto detection algorithms used in tunnel negotiations. Valid values: `simple`, `diff-req-resp`.
     */
    autoDetectAlgorithm?: pulumi.Input<string>;
    /**
     * Local host ID (must also be entered in the remote FortiGate's peer list).
     */
    hostId?: pulumi.Input<string>;
    /**
     * WANOpt tunnel optimization option. Valid values: `memory-usage`, `balanced`, `throughput`.
     */
    tunnelOptimization?: pulumi.Input<string>;
    /**
     * Relative strength of encryption algorithms accepted during tunnel negotiation. Valid values: `high`, `medium`, `low`.
     */
    tunnelSslAlgorithm?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Settings resource.
 */
export interface SettingsArgs {
    /**
     * Auto detection algorithms used in tunnel negotiations. Valid values: `simple`, `diff-req-resp`.
     */
    autoDetectAlgorithm?: pulumi.Input<string>;
    /**
     * Local host ID (must also be entered in the remote FortiGate's peer list).
     */
    hostId: pulumi.Input<string>;
    /**
     * WANOpt tunnel optimization option. Valid values: `memory-usage`, `balanced`, `throughput`.
     */
    tunnelOptimization?: pulumi.Input<string>;
    /**
     * Relative strength of encryption algorithms accepted during tunnel negotiation. Valid values: `high`, `medium`, `low`.
     */
    tunnelSslAlgorithm?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
