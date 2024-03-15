// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Configure IPsec manual keys.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.vpn.ipsec.Manualkey("trname", {
 *     authentication: "md5",
 *     authkey: "EE32CA121ECD772A-ECACAABA212345EC",
 *     enckey: "-",
 *     encryption: "null",
 *     "interface": "port4",
 *     localGw: "0.0.0.0",
 *     localspi: "0x100",
 *     remoteGw: "1.1.1.1",
 *     remotespi: "0x100",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * VpnIpsec Manualkey can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/ipsec/manualkey:Manualkey labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/ipsec/manualkey:Manualkey labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Manualkey extends pulumi.CustomResource {
    /**
     * Get an existing Manualkey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManualkeyState, opts?: pulumi.CustomResourceOptions): Manualkey {
        return new Manualkey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/ipsec/manualkey:Manualkey';

    /**
     * Returns true if the given object is an instance of Manualkey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Manualkey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Manualkey.__pulumiType;
    }

    /**
     * Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
     */
    public readonly authentication!: pulumi.Output<string>;
    /**
     * Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
     */
    public readonly authkey!: pulumi.Output<string>;
    /**
     * Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
     */
    public readonly enckey!: pulumi.Output<string>;
    /**
     * Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
     */
    public readonly encryption!: pulumi.Output<string>;
    /**
     * Name of the physical, aggregate, or VLAN interface.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Local gateway.
     */
    public readonly localGw!: pulumi.Output<string>;
    /**
     * Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
     */
    public readonly localspi!: pulumi.Output<string>;
    /**
     * IPsec tunnel name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable NPU offloading. Valid values: `enable`, `disable`.
     */
    public readonly npuOffload!: pulumi.Output<string>;
    /**
     * Peer gateway.
     */
    public readonly remoteGw!: pulumi.Output<string>;
    /**
     * Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
     */
    public readonly remotespi!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Manualkey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManualkeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManualkeyArgs | ManualkeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManualkeyState | undefined;
            resourceInputs["authentication"] = state ? state.authentication : undefined;
            resourceInputs["authkey"] = state ? state.authkey : undefined;
            resourceInputs["enckey"] = state ? state.enckey : undefined;
            resourceInputs["encryption"] = state ? state.encryption : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["localGw"] = state ? state.localGw : undefined;
            resourceInputs["localspi"] = state ? state.localspi : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["npuOffload"] = state ? state.npuOffload : undefined;
            resourceInputs["remoteGw"] = state ? state.remoteGw : undefined;
            resourceInputs["remotespi"] = state ? state.remotespi : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ManualkeyArgs | undefined;
            if ((!args || args.authentication === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authentication'");
            }
            if ((!args || args.encryption === undefined) && !opts.urn) {
                throw new Error("Missing required property 'encryption'");
            }
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.remoteGw === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteGw'");
            }
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["authkey"] = args?.authkey ? pulumi.secret(args.authkey) : undefined;
            resourceInputs["enckey"] = args?.enckey ? pulumi.secret(args.enckey) : undefined;
            resourceInputs["encryption"] = args ? args.encryption : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["localGw"] = args ? args.localGw : undefined;
            resourceInputs["localspi"] = args ? args.localspi : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["npuOffload"] = args ? args.npuOffload : undefined;
            resourceInputs["remoteGw"] = args ? args.remoteGw : undefined;
            resourceInputs["remotespi"] = args ? args.remotespi : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["authkey", "enckey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Manualkey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Manualkey resources.
 */
export interface ManualkeyState {
    /**
     * Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
     */
    authkey?: pulumi.Input<string>;
    /**
     * Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
     */
    enckey?: pulumi.Input<string>;
    /**
     * Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
     */
    encryption?: pulumi.Input<string>;
    /**
     * Name of the physical, aggregate, or VLAN interface.
     */
    interface?: pulumi.Input<string>;
    /**
     * Local gateway.
     */
    localGw?: pulumi.Input<string>;
    /**
     * Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
     */
    localspi?: pulumi.Input<string>;
    /**
     * IPsec tunnel name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable NPU offloading. Valid values: `enable`, `disable`.
     */
    npuOffload?: pulumi.Input<string>;
    /**
     * Peer gateway.
     */
    remoteGw?: pulumi.Input<string>;
    /**
     * Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
     */
    remotespi?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Manualkey resource.
 */
export interface ManualkeyArgs {
    /**
     * Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
     */
    authentication: pulumi.Input<string>;
    /**
     * Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
     */
    authkey?: pulumi.Input<string>;
    /**
     * Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
     */
    enckey?: pulumi.Input<string>;
    /**
     * Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.
     */
    encryption: pulumi.Input<string>;
    /**
     * Name of the physical, aggregate, or VLAN interface.
     */
    interface: pulumi.Input<string>;
    /**
     * Local gateway.
     */
    localGw?: pulumi.Input<string>;
    /**
     * Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
     */
    localspi?: pulumi.Input<string>;
    /**
     * IPsec tunnel name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable NPU offloading. Valid values: `enable`, `disable`.
     */
    npuOffload?: pulumi.Input<string>;
    /**
     * Peer gateway.
     */
    remoteGw: pulumi.Input<string>;
    /**
     * Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
     */
    remotespi?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}