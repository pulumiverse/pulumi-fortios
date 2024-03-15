// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Global PTP settings. Applies to FortiOS Version `>= 6.4.2`.
 *
 * ## Import
 *
 * SwitchControllerPtp Settings can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/ptp/settings:Settings labelname SwitchControllerPtpSettings
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/ptp/settings:Settings labelname SwitchControllerPtpSettings
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
    public static readonly __pulumiType = 'fortios:switchcontroller/ptp/settings:Settings';

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
     * Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
     */
    public readonly mode!: pulumi.Output<string>;
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
    constructor(name: string, args?: SettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SettingsArgs | SettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SettingsState | undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SettingsArgs | undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
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
     * Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
     */
    mode?: pulumi.Input<string>;
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
     * Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
