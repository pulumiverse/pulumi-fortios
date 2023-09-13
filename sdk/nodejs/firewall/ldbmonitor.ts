// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure server load balancing health monitors.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Ldbmonitor("trname", {
 *     httpMaxRedirects: 0,
 *     interval: 10,
 *     port: 0,
 *     retry: 3,
 *     timeout: 2,
 *     type: "ping",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall LdbMonitor can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:firewall/ldbmonitor:Ldbmonitor labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ldbmonitor extends pulumi.CustomResource {
    /**
     * Get an existing Ldbmonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LdbmonitorState, opts?: pulumi.CustomResourceOptions): Ldbmonitor {
        return new Ldbmonitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/ldbmonitor:Ldbmonitor';

    /**
     * Returns true if the given object is an instance of Ldbmonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ldbmonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ldbmonitor.__pulumiType;
    }

    /**
     * Response IP expected from DNS server.
     */
    public readonly dnsMatchIp!: pulumi.Output<string>;
    /**
     * Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
     */
    public readonly dnsProtocol!: pulumi.Output<string>;
    /**
     * Fully qualified domain name to resolve for the DNS probe.
     */
    public readonly dnsRequestDomain!: pulumi.Output<string>;
    /**
     * URL used to send a GET request to check the health of an HTTP server.
     */
    public readonly httpGet!: pulumi.Output<string>;
    /**
     * String to match the value expected in response to an HTTP-GET request.
     */
    public readonly httpMatch!: pulumi.Output<string>;
    /**
     * The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
     */
    public readonly httpMaxRedirects!: pulumi.Output<number>;
    /**
     * Time between health checks (5 - 65635 sec, default = 10).
     */
    public readonly interval!: pulumi.Output<number>;
    /**
     * Monitor name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Number health check attempts before the server is considered down (1 - 255, default = 3).
     */
    public readonly retry!: pulumi.Output<number>;
    /**
     * Source IP for ldb-monitor.
     */
    public readonly srcIp!: pulumi.Output<string>;
    /**
     * Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ldbmonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LdbmonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LdbmonitorArgs | LdbmonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LdbmonitorState | undefined;
            resourceInputs["dnsMatchIp"] = state ? state.dnsMatchIp : undefined;
            resourceInputs["dnsProtocol"] = state ? state.dnsProtocol : undefined;
            resourceInputs["dnsRequestDomain"] = state ? state.dnsRequestDomain : undefined;
            resourceInputs["httpGet"] = state ? state.httpGet : undefined;
            resourceInputs["httpMatch"] = state ? state.httpMatch : undefined;
            resourceInputs["httpMaxRedirects"] = state ? state.httpMaxRedirects : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["retry"] = state ? state.retry : undefined;
            resourceInputs["srcIp"] = state ? state.srcIp : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LdbmonitorArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["dnsMatchIp"] = args ? args.dnsMatchIp : undefined;
            resourceInputs["dnsProtocol"] = args ? args.dnsProtocol : undefined;
            resourceInputs["dnsRequestDomain"] = args ? args.dnsRequestDomain : undefined;
            resourceInputs["httpGet"] = args ? args.httpGet : undefined;
            resourceInputs["httpMatch"] = args ? args.httpMatch : undefined;
            resourceInputs["httpMaxRedirects"] = args ? args.httpMaxRedirects : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["retry"] = args ? args.retry : undefined;
            resourceInputs["srcIp"] = args ? args.srcIp : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ldbmonitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ldbmonitor resources.
 */
export interface LdbmonitorState {
    /**
     * Response IP expected from DNS server.
     */
    dnsMatchIp?: pulumi.Input<string>;
    /**
     * Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
     */
    dnsProtocol?: pulumi.Input<string>;
    /**
     * Fully qualified domain name to resolve for the DNS probe.
     */
    dnsRequestDomain?: pulumi.Input<string>;
    /**
     * URL used to send a GET request to check the health of an HTTP server.
     */
    httpGet?: pulumi.Input<string>;
    /**
     * String to match the value expected in response to an HTTP-GET request.
     */
    httpMatch?: pulumi.Input<string>;
    /**
     * The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
     */
    httpMaxRedirects?: pulumi.Input<number>;
    /**
     * Time between health checks (5 - 65635 sec, default = 10).
     */
    interval?: pulumi.Input<number>;
    /**
     * Monitor name.
     */
    name?: pulumi.Input<string>;
    /**
     * Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
     */
    port?: pulumi.Input<number>;
    /**
     * Number health check attempts before the server is considered down (1 - 255, default = 3).
     */
    retry?: pulumi.Input<number>;
    /**
     * Source IP for ldb-monitor.
     */
    srcIp?: pulumi.Input<string>;
    /**
     * Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ldbmonitor resource.
 */
export interface LdbmonitorArgs {
    /**
     * Response IP expected from DNS server.
     */
    dnsMatchIp?: pulumi.Input<string>;
    /**
     * Select the protocol used by the DNS health check monitor to check the health of the server (UDP | TCP). Valid values: `udp`, `tcp`.
     */
    dnsProtocol?: pulumi.Input<string>;
    /**
     * Fully qualified domain name to resolve for the DNS probe.
     */
    dnsRequestDomain?: pulumi.Input<string>;
    /**
     * URL used to send a GET request to check the health of an HTTP server.
     */
    httpGet?: pulumi.Input<string>;
    /**
     * String to match the value expected in response to an HTTP-GET request.
     */
    httpMatch?: pulumi.Input<string>;
    /**
     * The maximum number of HTTP redirects to be allowed (0 - 5, default = 0).
     */
    httpMaxRedirects?: pulumi.Input<number>;
    /**
     * Time between health checks (5 - 65635 sec, default = 10).
     */
    interval?: pulumi.Input<number>;
    /**
     * Monitor name.
     */
    name?: pulumi.Input<string>;
    /**
     * Service port used to perform the health check. If 0, health check monitor inherits port configured for the server (0 - 65635, default = 0).
     */
    port?: pulumi.Input<number>;
    /**
     * Number health check attempts before the server is considered down (1 - 255, default = 3).
     */
    retry?: pulumi.Input<number>;
    /**
     * Source IP for ldb-monitor.
     */
    srcIp?: pulumi.Input<string>;
    /**
     * Time to wait to receive response to a health check from a server. Reaching the timeout means the health check failed (1 - 255 sec, default = 2).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Select the Monitor type used by the health check monitor to check the health of the server (PING | TCP | HTTP).
     */
    type: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
