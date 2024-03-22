// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure explicit FTP proxy settings.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.ftpproxy.Explicit("trname", {
 *     incomingIp: "0.0.0.0",
 *     secDefaultAction: "deny",
 *     status: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * FtpProxy Explicit can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:ftpproxy/explicit:Explicit labelname FtpProxyExplicit
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:ftpproxy/explicit:Explicit labelname FtpProxyExplicit
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Explicit extends pulumi.CustomResource {
    /**
     * Get an existing Explicit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExplicitState, opts?: pulumi.CustomResourceOptions): Explicit {
        return new Explicit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:ftpproxy/explicit:Explicit';

    /**
     * Returns true if the given object is an instance of Explicit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Explicit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Explicit.__pulumiType;
    }

    /**
     * Accept incoming FTP requests from this IP address. An interface must have this IP address.
     */
    public readonly incomingIp!: pulumi.Output<string>;
    /**
     * Accept incoming FTP requests on one or more ports.
     */
    public readonly incomingPort!: pulumi.Output<string>;
    /**
     * Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
     */
    public readonly outgoingIp!: pulumi.Output<string>;
    /**
     * Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
     */
    public readonly secDefaultAction!: pulumi.Output<string>;
    /**
     * Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
     */
    public readonly serverDataMode!: pulumi.Output<string>;
    /**
     * Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
     */
    public readonly ssl!: pulumi.Output<string>;
    /**
     * Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
     */
    public readonly sslAlgorithm!: pulumi.Output<string>;
    /**
     * Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
     */
    public readonly sslCert!: pulumi.Output<string>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    public readonly sslDhBits!: pulumi.Output<string>;
    /**
     * Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Explicit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExplicitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExplicitArgs | ExplicitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExplicitState | undefined;
            resourceInputs["incomingIp"] = state ? state.incomingIp : undefined;
            resourceInputs["incomingPort"] = state ? state.incomingPort : undefined;
            resourceInputs["outgoingIp"] = state ? state.outgoingIp : undefined;
            resourceInputs["secDefaultAction"] = state ? state.secDefaultAction : undefined;
            resourceInputs["serverDataMode"] = state ? state.serverDataMode : undefined;
            resourceInputs["ssl"] = state ? state.ssl : undefined;
            resourceInputs["sslAlgorithm"] = state ? state.sslAlgorithm : undefined;
            resourceInputs["sslCert"] = state ? state.sslCert : undefined;
            resourceInputs["sslDhBits"] = state ? state.sslDhBits : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ExplicitArgs | undefined;
            resourceInputs["incomingIp"] = args ? args.incomingIp : undefined;
            resourceInputs["incomingPort"] = args ? args.incomingPort : undefined;
            resourceInputs["outgoingIp"] = args ? args.outgoingIp : undefined;
            resourceInputs["secDefaultAction"] = args ? args.secDefaultAction : undefined;
            resourceInputs["serverDataMode"] = args ? args.serverDataMode : undefined;
            resourceInputs["ssl"] = args ? args.ssl : undefined;
            resourceInputs["sslAlgorithm"] = args ? args.sslAlgorithm : undefined;
            resourceInputs["sslCert"] = args ? args.sslCert : undefined;
            resourceInputs["sslDhBits"] = args ? args.sslDhBits : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Explicit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Explicit resources.
 */
export interface ExplicitState {
    /**
     * Accept incoming FTP requests from this IP address. An interface must have this IP address.
     */
    incomingIp?: pulumi.Input<string>;
    /**
     * Accept incoming FTP requests on one or more ports.
     */
    incomingPort?: pulumi.Input<string>;
    /**
     * Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
     */
    outgoingIp?: pulumi.Input<string>;
    /**
     * Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
     */
    secDefaultAction?: pulumi.Input<string>;
    /**
     * Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
     */
    serverDataMode?: pulumi.Input<string>;
    /**
     * Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
     */
    ssl?: pulumi.Input<string>;
    /**
     * Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
     */
    sslAlgorithm?: pulumi.Input<string>;
    /**
     * Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
     */
    sslCert?: pulumi.Input<string>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    sslDhBits?: pulumi.Input<string>;
    /**
     * Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Explicit resource.
 */
export interface ExplicitArgs {
    /**
     * Accept incoming FTP requests from this IP address. An interface must have this IP address.
     */
    incomingIp?: pulumi.Input<string>;
    /**
     * Accept incoming FTP requests on one or more ports.
     */
    incomingPort?: pulumi.Input<string>;
    /**
     * Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
     */
    outgoingIp?: pulumi.Input<string>;
    /**
     * Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
     */
    secDefaultAction?: pulumi.Input<string>;
    /**
     * Determine mode of data session on FTP server side. Valid values: `client`, `passive`.
     */
    serverDataMode?: pulumi.Input<string>;
    /**
     * Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
     */
    ssl?: pulumi.Input<string>;
    /**
     * Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
     */
    sslAlgorithm?: pulumi.Input<string>;
    /**
     * Name of certificate for SSL connections to this server. On FortiOS versions 6.2.4-7.4.0: default = "Fortinet_CA_SSL". On FortiOS versions 7.4.1: default = "Fortinet_SSL".
     */
    sslCert?: pulumi.Input<string>;
    /**
     * Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
     */
    sslDhBits?: pulumi.Input<string>;
    /**
     * Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
