// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Configure IP to MAC binding settings.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.ipmacbinding.Setting("trname", {
 *     bindthroughfw: "disable",
 *     bindtofw: "disable",
 *     undefinedhost: "block",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * FirewallIpmacbinding Setting can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/ipmacbinding/setting:Setting labelname FirewallIpmacbindingSetting
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/ipmacbinding/setting:Setting labelname FirewallIpmacbindingSetting
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Setting extends pulumi.CustomResource {
    /**
     * Get an existing Setting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SettingState, opts?: pulumi.CustomResourceOptions): Setting {
        return new Setting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/ipmacbinding/setting:Setting';

    /**
     * Returns true if the given object is an instance of Setting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Setting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Setting.__pulumiType;
    }

    /**
     * Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
     */
    public readonly bindthroughfw!: pulumi.Output<string>;
    /**
     * Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
     */
    public readonly bindtofw!: pulumi.Output<string>;
    /**
     * Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
     */
    public readonly undefinedhost!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Setting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SettingArgs | SettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SettingState | undefined;
            resourceInputs["bindthroughfw"] = state ? state.bindthroughfw : undefined;
            resourceInputs["bindtofw"] = state ? state.bindtofw : undefined;
            resourceInputs["undefinedhost"] = state ? state.undefinedhost : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SettingArgs | undefined;
            resourceInputs["bindthroughfw"] = args ? args.bindthroughfw : undefined;
            resourceInputs["bindtofw"] = args ? args.bindtofw : undefined;
            resourceInputs["undefinedhost"] = args ? args.undefinedhost : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Setting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Setting resources.
 */
export interface SettingState {
    /**
     * Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
     */
    bindthroughfw?: pulumi.Input<string>;
    /**
     * Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
     */
    bindtofw?: pulumi.Input<string>;
    /**
     * Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
     */
    undefinedhost?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Setting resource.
 */
export interface SettingArgs {
    /**
     * Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
     */
    bindthroughfw?: pulumi.Input<string>;
    /**
     * Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
     */
    bindtofw?: pulumi.Input<string>;
    /**
     * Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
     */
    undefinedhost?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
