// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure ICAP servers.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.icap.Server("trname", {
 *     ip6Address: "::",
 *     ipAddress: "1.1.1.1",
 *     ipVersion: "4",
 *     maxConnections: 100,
 *     port: 22,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Icap Server can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:icap/server:Server labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:icap/server:Server labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerState, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:icap/server:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * Address type of the remote ICAP server: IPv4, IPv6 or FQDN. Valid values: `ip4`, `ip6`, `fqdn`.
     */
    public readonly addrType!: pulumi.Output<string>;
    /**
     * ICAP remote server Fully Qualified Domain Name (FQDN).
     */
    public readonly fqdn!: pulumi.Output<string>;
    /**
     * Enable/disable ICAP remote server health checking. Attempts to connect to the remote ICAP server to verify that the server is operating normally. Valid values: `disable`, `enable`.
     */
    public readonly healthcheck!: pulumi.Output<string>;
    /**
     * ICAP Service name to use for health checks.
     */
    public readonly healthcheckService!: pulumi.Output<string>;
    /**
     * IPv6 address of the ICAP server.
     */
    public readonly ip6Address!: pulumi.Output<string>;
    /**
     * IPv4 address of the ICAP server.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * IP version. Valid values: `4`, `6`.
     */
    public readonly ipVersion!: pulumi.Output<string>;
    /**
     * Maximum number of concurrent connections to ICAP server. Must not be less than wad-worker-count.
     */
    public readonly maxConnections!: pulumi.Output<number>;
    /**
     * Server name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ICAP server port.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
     */
    public readonly secure!: pulumi.Output<string>;
    /**
     * CA certificate name.
     */
    public readonly sslCert!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs | ServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerState | undefined;
            resourceInputs["addrType"] = state ? state.addrType : undefined;
            resourceInputs["fqdn"] = state ? state.fqdn : undefined;
            resourceInputs["healthcheck"] = state ? state.healthcheck : undefined;
            resourceInputs["healthcheckService"] = state ? state.healthcheckService : undefined;
            resourceInputs["ip6Address"] = state ? state.ip6Address : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["maxConnections"] = state ? state.maxConnections : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["secure"] = state ? state.secure : undefined;
            resourceInputs["sslCert"] = state ? state.sslCert : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            resourceInputs["addrType"] = args ? args.addrType : undefined;
            resourceInputs["fqdn"] = args ? args.fqdn : undefined;
            resourceInputs["healthcheck"] = args ? args.healthcheck : undefined;
            resourceInputs["healthcheckService"] = args ? args.healthcheckService : undefined;
            resourceInputs["ip6Address"] = args ? args.ip6Address : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["maxConnections"] = args ? args.maxConnections : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["secure"] = args ? args.secure : undefined;
            resourceInputs["sslCert"] = args ? args.sslCert : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Server.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * Address type of the remote ICAP server: IPv4, IPv6 or FQDN. Valid values: `ip4`, `ip6`, `fqdn`.
     */
    addrType?: pulumi.Input<string>;
    /**
     * ICAP remote server Fully Qualified Domain Name (FQDN).
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Enable/disable ICAP remote server health checking. Attempts to connect to the remote ICAP server to verify that the server is operating normally. Valid values: `disable`, `enable`.
     */
    healthcheck?: pulumi.Input<string>;
    /**
     * ICAP Service name to use for health checks.
     */
    healthcheckService?: pulumi.Input<string>;
    /**
     * IPv6 address of the ICAP server.
     */
    ip6Address?: pulumi.Input<string>;
    /**
     * IPv4 address of the ICAP server.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * IP version. Valid values: `4`, `6`.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * Maximum number of concurrent connections to ICAP server. Must not be less than wad-worker-count.
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * ICAP server port.
     */
    port?: pulumi.Input<number>;
    /**
     * Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
     */
    secure?: pulumi.Input<string>;
    /**
     * CA certificate name.
     */
    sslCert?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * Address type of the remote ICAP server: IPv4, IPv6 or FQDN. Valid values: `ip4`, `ip6`, `fqdn`.
     */
    addrType?: pulumi.Input<string>;
    /**
     * ICAP remote server Fully Qualified Domain Name (FQDN).
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Enable/disable ICAP remote server health checking. Attempts to connect to the remote ICAP server to verify that the server is operating normally. Valid values: `disable`, `enable`.
     */
    healthcheck?: pulumi.Input<string>;
    /**
     * ICAP Service name to use for health checks.
     */
    healthcheckService?: pulumi.Input<string>;
    /**
     * IPv6 address of the ICAP server.
     */
    ip6Address?: pulumi.Input<string>;
    /**
     * IPv4 address of the ICAP server.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * IP version. Valid values: `4`, `6`.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * Maximum number of concurrent connections to ICAP server. Must not be less than wad-worker-count.
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * ICAP server port.
     */
    port?: pulumi.Input<number>;
    /**
     * Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
     */
    secure?: pulumi.Input<string>;
    /**
     * CA certificate name.
     */
    sslCert?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
