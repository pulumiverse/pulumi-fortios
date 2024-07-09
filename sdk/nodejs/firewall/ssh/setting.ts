// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * SSH proxy settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.ssh.Setting("trname", {
 *     caname: "Fortinet_SSH_CA",
 *     hostTrustedChecking: "enable",
 *     hostkeyDsa1024: "Fortinet_SSH_DSA1024",
 *     hostkeyEcdsa256: "Fortinet_SSH_ECDSA256",
 *     hostkeyEcdsa384: "Fortinet_SSH_ECDSA384",
 *     hostkeyEcdsa521: "Fortinet_SSH_ECDSA521",
 *     hostkeyEd25519: "Fortinet_SSH_ED25519",
 *     hostkeyRsa2048: "Fortinet_SSH_RSA2048",
 *     untrustedCaname: "Fortinet_SSH_CA_Untrusted",
 * });
 * ```
 *
 * ## Import
 *
 * FirewallSsh Setting can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/ssh/setting:Setting labelname FirewallSshSetting
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/ssh/setting:Setting labelname FirewallSshSetting
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
    public static readonly __pulumiType = 'fortios:firewall/ssh/setting:Setting';

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
     * CA certificate used by SSH Inspection.
     */
    public readonly caname!: pulumi.Output<string>;
    /**
     * Enable/disable host trusted checking. Valid values: `enable`, `disable`.
     */
    public readonly hostTrustedChecking!: pulumi.Output<string>;
    /**
     * DSA certificate used by SSH proxy.
     */
    public readonly hostkeyDsa1024!: pulumi.Output<string>;
    /**
     * ECDSA nid256 certificate used by SSH proxy.
     */
    public readonly hostkeyEcdsa256!: pulumi.Output<string>;
    /**
     * ECDSA nid384 certificate used by SSH proxy.
     */
    public readonly hostkeyEcdsa384!: pulumi.Output<string>;
    /**
     * ECDSA nid384 certificate used by SSH proxy.
     */
    public readonly hostkeyEcdsa521!: pulumi.Output<string>;
    /**
     * ED25519 hostkey used by SSH proxy.
     */
    public readonly hostkeyEd25519!: pulumi.Output<string>;
    /**
     * RSA certificate used by SSH proxy.
     */
    public readonly hostkeyRsa2048!: pulumi.Output<string>;
    /**
     * Untrusted CA certificate used by SSH Inspection.
     */
    public readonly untrustedCaname!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

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
            resourceInputs["caname"] = state ? state.caname : undefined;
            resourceInputs["hostTrustedChecking"] = state ? state.hostTrustedChecking : undefined;
            resourceInputs["hostkeyDsa1024"] = state ? state.hostkeyDsa1024 : undefined;
            resourceInputs["hostkeyEcdsa256"] = state ? state.hostkeyEcdsa256 : undefined;
            resourceInputs["hostkeyEcdsa384"] = state ? state.hostkeyEcdsa384 : undefined;
            resourceInputs["hostkeyEcdsa521"] = state ? state.hostkeyEcdsa521 : undefined;
            resourceInputs["hostkeyEd25519"] = state ? state.hostkeyEd25519 : undefined;
            resourceInputs["hostkeyRsa2048"] = state ? state.hostkeyRsa2048 : undefined;
            resourceInputs["untrustedCaname"] = state ? state.untrustedCaname : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SettingArgs | undefined;
            resourceInputs["caname"] = args ? args.caname : undefined;
            resourceInputs["hostTrustedChecking"] = args ? args.hostTrustedChecking : undefined;
            resourceInputs["hostkeyDsa1024"] = args ? args.hostkeyDsa1024 : undefined;
            resourceInputs["hostkeyEcdsa256"] = args ? args.hostkeyEcdsa256 : undefined;
            resourceInputs["hostkeyEcdsa384"] = args ? args.hostkeyEcdsa384 : undefined;
            resourceInputs["hostkeyEcdsa521"] = args ? args.hostkeyEcdsa521 : undefined;
            resourceInputs["hostkeyEd25519"] = args ? args.hostkeyEd25519 : undefined;
            resourceInputs["hostkeyRsa2048"] = args ? args.hostkeyRsa2048 : undefined;
            resourceInputs["untrustedCaname"] = args ? args.untrustedCaname : undefined;
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
     * CA certificate used by SSH Inspection.
     */
    caname?: pulumi.Input<string>;
    /**
     * Enable/disable host trusted checking. Valid values: `enable`, `disable`.
     */
    hostTrustedChecking?: pulumi.Input<string>;
    /**
     * DSA certificate used by SSH proxy.
     */
    hostkeyDsa1024?: pulumi.Input<string>;
    /**
     * ECDSA nid256 certificate used by SSH proxy.
     */
    hostkeyEcdsa256?: pulumi.Input<string>;
    /**
     * ECDSA nid384 certificate used by SSH proxy.
     */
    hostkeyEcdsa384?: pulumi.Input<string>;
    /**
     * ECDSA nid384 certificate used by SSH proxy.
     */
    hostkeyEcdsa521?: pulumi.Input<string>;
    /**
     * ED25519 hostkey used by SSH proxy.
     */
    hostkeyEd25519?: pulumi.Input<string>;
    /**
     * RSA certificate used by SSH proxy.
     */
    hostkeyRsa2048?: pulumi.Input<string>;
    /**
     * Untrusted CA certificate used by SSH Inspection.
     */
    untrustedCaname?: pulumi.Input<string>;
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
     * CA certificate used by SSH Inspection.
     */
    caname?: pulumi.Input<string>;
    /**
     * Enable/disable host trusted checking. Valid values: `enable`, `disable`.
     */
    hostTrustedChecking?: pulumi.Input<string>;
    /**
     * DSA certificate used by SSH proxy.
     */
    hostkeyDsa1024?: pulumi.Input<string>;
    /**
     * ECDSA nid256 certificate used by SSH proxy.
     */
    hostkeyEcdsa256?: pulumi.Input<string>;
    /**
     * ECDSA nid384 certificate used by SSH proxy.
     */
    hostkeyEcdsa384?: pulumi.Input<string>;
    /**
     * ECDSA nid384 certificate used by SSH proxy.
     */
    hostkeyEcdsa521?: pulumi.Input<string>;
    /**
     * ED25519 hostkey used by SSH proxy.
     */
    hostkeyEd25519?: pulumi.Input<string>;
    /**
     * RSA certificate used by SSH proxy.
     */
    hostkeyRsa2048?: pulumi.Input<string>;
    /**
     * Untrusted CA certificate used by SSH Inspection.
     */
    untrustedCaname?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
