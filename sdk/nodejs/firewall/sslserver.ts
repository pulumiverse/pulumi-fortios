// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure SSL servers.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Sslserver("trname", {
 *     addHeaderXForwardedProto: "enable",
 *     ip: "1.1.1.1",
 *     mappedPort: 2234,
 *     port: 32321,
 *     sslAlgorithm: "high",
 *     sslCert: "Fortinet_CA_SSL",
 *     sslClientRenegotiation: "allow",
 *     sslDhBits: "2048",
 *     sslMaxVersion: "tls-1.2",
 *     sslMinVersion: "tls-1.1",
 *     sslMode: "half",
 *     sslSendEmptyFrags: "enable",
 *     urlRewrite: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Firewall SslServer can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/sslserver:Sslserver labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/sslserver:Sslserver labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Sslserver extends pulumi.CustomResource {
    /**
     * Get an existing Sslserver resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SslserverState, opts?: pulumi.CustomResourceOptions): Sslserver {
        return new Sslserver(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/sslserver:Sslserver';

    /**
     * Returns true if the given object is an instance of Sslserver.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sslserver {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sslserver.__pulumiType;
    }

    /**
     * Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
     */
    public readonly addHeaderXForwardedProto!: pulumi.Output<string>;
    /**
     * IPv4 address of the SSL server.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Mapped server service port (1 - 65535, default = 80).
     */
    public readonly mappedPort!: pulumi.Output<number>;
    /**
     * Server name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Server service port (1 - 65535, default = 443).
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
     */
    public readonly sslAlgorithm!: pulumi.Output<string>;
    /**
     * Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
     */
    public readonly sslCert!: pulumi.Output<string>;
    /**
     * Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
     */
    public readonly sslClientRenegotiation!: pulumi.Output<string>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    public readonly sslDhBits!: pulumi.Output<string>;
    /**
     * Highest SSL/TLS version to negotiate.
     */
    public readonly sslMaxVersion!: pulumi.Output<string>;
    /**
     * Lowest SSL/TLS version to negotiate.
     */
    public readonly sslMinVersion!: pulumi.Output<string>;
    /**
     * SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
     */
    public readonly sslMode!: pulumi.Output<string>;
    /**
     * Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
     */
    public readonly sslSendEmptyFrags!: pulumi.Output<string>;
    /**
     * Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
     */
    public readonly urlRewrite!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Sslserver resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SslserverArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SslserverArgs | SslserverState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SslserverState | undefined;
            resourceInputs["addHeaderXForwardedProto"] = state ? state.addHeaderXForwardedProto : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["mappedPort"] = state ? state.mappedPort : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["sslAlgorithm"] = state ? state.sslAlgorithm : undefined;
            resourceInputs["sslCert"] = state ? state.sslCert : undefined;
            resourceInputs["sslClientRenegotiation"] = state ? state.sslClientRenegotiation : undefined;
            resourceInputs["sslDhBits"] = state ? state.sslDhBits : undefined;
            resourceInputs["sslMaxVersion"] = state ? state.sslMaxVersion : undefined;
            resourceInputs["sslMinVersion"] = state ? state.sslMinVersion : undefined;
            resourceInputs["sslMode"] = state ? state.sslMode : undefined;
            resourceInputs["sslSendEmptyFrags"] = state ? state.sslSendEmptyFrags : undefined;
            resourceInputs["urlRewrite"] = state ? state.urlRewrite : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SslserverArgs | undefined;
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.sslCert === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslCert'");
            }
            resourceInputs["addHeaderXForwardedProto"] = args ? args.addHeaderXForwardedProto : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["mappedPort"] = args ? args.mappedPort : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["sslAlgorithm"] = args ? args.sslAlgorithm : undefined;
            resourceInputs["sslCert"] = args ? args.sslCert : undefined;
            resourceInputs["sslClientRenegotiation"] = args ? args.sslClientRenegotiation : undefined;
            resourceInputs["sslDhBits"] = args ? args.sslDhBits : undefined;
            resourceInputs["sslMaxVersion"] = args ? args.sslMaxVersion : undefined;
            resourceInputs["sslMinVersion"] = args ? args.sslMinVersion : undefined;
            resourceInputs["sslMode"] = args ? args.sslMode : undefined;
            resourceInputs["sslSendEmptyFrags"] = args ? args.sslSendEmptyFrags : undefined;
            resourceInputs["urlRewrite"] = args ? args.urlRewrite : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sslserver.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Sslserver resources.
 */
export interface SslserverState {
    /**
     * Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
     */
    addHeaderXForwardedProto?: pulumi.Input<string>;
    /**
     * IPv4 address of the SSL server.
     */
    ip?: pulumi.Input<string>;
    /**
     * Mapped server service port (1 - 65535, default = 80).
     */
    mappedPort?: pulumi.Input<number>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * Server service port (1 - 65535, default = 443).
     */
    port?: pulumi.Input<number>;
    /**
     * Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
     */
    sslAlgorithm?: pulumi.Input<string>;
    /**
     * Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
     */
    sslCert?: pulumi.Input<string>;
    /**
     * Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
     */
    sslClientRenegotiation?: pulumi.Input<string>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    sslDhBits?: pulumi.Input<string>;
    /**
     * Highest SSL/TLS version to negotiate.
     */
    sslMaxVersion?: pulumi.Input<string>;
    /**
     * Lowest SSL/TLS version to negotiate.
     */
    sslMinVersion?: pulumi.Input<string>;
    /**
     * SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
     */
    sslMode?: pulumi.Input<string>;
    /**
     * Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
     */
    sslSendEmptyFrags?: pulumi.Input<string>;
    /**
     * Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
     */
    urlRewrite?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Sslserver resource.
 */
export interface SslserverArgs {
    /**
     * Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
     */
    addHeaderXForwardedProto?: pulumi.Input<string>;
    /**
     * IPv4 address of the SSL server.
     */
    ip: pulumi.Input<string>;
    /**
     * Mapped server service port (1 - 65535, default = 80).
     */
    mappedPort?: pulumi.Input<number>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * Server service port (1 - 65535, default = 443).
     */
    port: pulumi.Input<number>;
    /**
     * Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
     */
    sslAlgorithm?: pulumi.Input<string>;
    /**
     * Name of certificate for SSL connections to this server. On FortiOS versions 6.2.0-7.2.6: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.0-7.4.1: default = "Fortinet_SSL".
     */
    sslCert: pulumi.Input<string>;
    /**
     * Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
     */
    sslClientRenegotiation?: pulumi.Input<string>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    sslDhBits?: pulumi.Input<string>;
    /**
     * Highest SSL/TLS version to negotiate.
     */
    sslMaxVersion?: pulumi.Input<string>;
    /**
     * Lowest SSL/TLS version to negotiate.
     */
    sslMinVersion?: pulumi.Input<string>;
    /**
     * SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
     */
    sslMode?: pulumi.Input<string>;
    /**
     * Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
     */
    sslSendEmptyFrags?: pulumi.Input<string>;
    /**
     * Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
     */
    urlRewrite?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
