// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure integrated NAC settings for FortiSwitch. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,7.0.0`.
 *
 * ## Import
 *
 * SwitchController NacSettings can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/nacsettings:Nacsettings labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/nacsettings:Nacsettings labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Nacsettings extends pulumi.CustomResource {
    /**
     * Get an existing Nacsettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NacsettingsState, opts?: pulumi.CustomResourceOptions): Nacsettings {
        return new Nacsettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/nacsettings:Nacsettings';

    /**
     * Returns true if the given object is an instance of Nacsettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Nacsettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Nacsettings.__pulumiType;
    }

    /**
     * Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
     */
    public readonly autoAuth!: pulumi.Output<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    public readonly bounceNacPort!: pulumi.Output<string>;
    /**
     * Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
     */
    public readonly inactiveTimer!: pulumi.Output<number>;
    /**
     * Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
     */
    public readonly linkDownFlush!: pulumi.Output<string>;
    /**
     * Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * NAC settings name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Default NAC Onboarding VLAN when NAC devices are discovered.
     */
    public readonly onboardingVlan!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Nacsettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NacsettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NacsettingsArgs | NacsettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NacsettingsState | undefined;
            resourceInputs["autoAuth"] = state ? state.autoAuth : undefined;
            resourceInputs["bounceNacPort"] = state ? state.bounceNacPort : undefined;
            resourceInputs["inactiveTimer"] = state ? state.inactiveTimer : undefined;
            resourceInputs["linkDownFlush"] = state ? state.linkDownFlush : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["onboardingVlan"] = state ? state.onboardingVlan : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as NacsettingsArgs | undefined;
            resourceInputs["autoAuth"] = args ? args.autoAuth : undefined;
            resourceInputs["bounceNacPort"] = args ? args.bounceNacPort : undefined;
            resourceInputs["inactiveTimer"] = args ? args.inactiveTimer : undefined;
            resourceInputs["linkDownFlush"] = args ? args.linkDownFlush : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["onboardingVlan"] = args ? args.onboardingVlan : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Nacsettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Nacsettings resources.
 */
export interface NacsettingsState {
    /**
     * Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
     */
    autoAuth?: pulumi.Input<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    bounceNacPort?: pulumi.Input<string>;
    /**
     * Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
     */
    inactiveTimer?: pulumi.Input<number>;
    /**
     * Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
     */
    linkDownFlush?: pulumi.Input<string>;
    /**
     * Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
     */
    mode?: pulumi.Input<string>;
    /**
     * NAC settings name.
     */
    name?: pulumi.Input<string>;
    /**
     * Default NAC Onboarding VLAN when NAC devices are discovered.
     */
    onboardingVlan?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Nacsettings resource.
 */
export interface NacsettingsArgs {
    /**
     * Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
     */
    autoAuth?: pulumi.Input<string>;
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
     */
    bounceNacPort?: pulumi.Input<string>;
    /**
     * Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
     */
    inactiveTimer?: pulumi.Input<number>;
    /**
     * Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
     */
    linkDownFlush?: pulumi.Input<string>;
    /**
     * Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
     */
    mode?: pulumi.Input<string>;
    /**
     * NAC settings name.
     */
    name?: pulumi.Input<string>;
    /**
     * Default NAC Onboarding VLAN when NAC devices are discovered.
     */
    onboardingVlan?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
