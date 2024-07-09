// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FortiToken Mobile push services.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Ftmpush("trname", {
 *     serverIp: "0.0.0.0",
 *     serverPort: 4433,
 *     status: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * System FtmPush can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/ftmpush:Ftmpush labelname SystemFtmPush
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/ftmpush:Ftmpush labelname SystemFtmPush
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ftmpush extends pulumi.CustomResource {
    /**
     * Get an existing Ftmpush resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FtmpushState, opts?: pulumi.CustomResourceOptions): Ftmpush {
        return new Ftmpush(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/ftmpush:Ftmpush';

    /**
     * Returns true if the given object is an instance of Ftmpush.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ftmpush {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ftmpush.__pulumiType;
    }

    /**
     * Enable/disable communication to the proxy server in FortiGuard configuration. Valid values: `enable`, `disable`.
     */
    public readonly proxy!: pulumi.Output<string>;
    /**
     * IPv4 address or domain name of FortiToken Mobile push services server.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Name of the server certificate to be used for SSL. On FortiOS versions 6.4.0-7.4.3: default = Fortinet_Factory.
     */
    public readonly serverCert!: pulumi.Output<string>;
    /**
     * IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
     */
    public readonly serverIp!: pulumi.Output<string>;
    /**
     * Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
     */
    public readonly serverPort!: pulumi.Output<number>;
    /**
     * Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Ftmpush resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FtmpushArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FtmpushArgs | FtmpushState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FtmpushState | undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["serverCert"] = state ? state.serverCert : undefined;
            resourceInputs["serverIp"] = state ? state.serverIp : undefined;
            resourceInputs["serverPort"] = state ? state.serverPort : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FtmpushArgs | undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["serverCert"] = args ? args.serverCert : undefined;
            resourceInputs["serverIp"] = args ? args.serverIp : undefined;
            resourceInputs["serverPort"] = args ? args.serverPort : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ftmpush.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ftmpush resources.
 */
export interface FtmpushState {
    /**
     * Enable/disable communication to the proxy server in FortiGuard configuration. Valid values: `enable`, `disable`.
     */
    proxy?: pulumi.Input<string>;
    /**
     * IPv4 address or domain name of FortiToken Mobile push services server.
     */
    server?: pulumi.Input<string>;
    /**
     * Name of the server certificate to be used for SSL. On FortiOS versions 6.4.0-7.4.3: default = Fortinet_Factory.
     */
    serverCert?: pulumi.Input<string>;
    /**
     * IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
     */
    serverIp?: pulumi.Input<string>;
    /**
     * Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
     */
    serverPort?: pulumi.Input<number>;
    /**
     * Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ftmpush resource.
 */
export interface FtmpushArgs {
    /**
     * Enable/disable communication to the proxy server in FortiGuard configuration. Valid values: `enable`, `disable`.
     */
    proxy?: pulumi.Input<string>;
    /**
     * IPv4 address or domain name of FortiToken Mobile push services server.
     */
    server?: pulumi.Input<string>;
    /**
     * Name of the server certificate to be used for SSL. On FortiOS versions 6.4.0-7.4.3: default = Fortinet_Factory.
     */
    serverCert?: pulumi.Input<string>;
    /**
     * IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
     */
    serverIp?: pulumi.Input<string>;
    /**
     * Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
     */
    serverPort?: pulumi.Input<number>;
    /**
     * Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
