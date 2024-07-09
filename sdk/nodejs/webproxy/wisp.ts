// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure Wireless Internet service provider (WISP) servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.webproxy.Wisp("trname", {
 *     maxConnections: 64,
 *     outgoingIp: "0.0.0.0",
 *     serverIp: "1.1.1.1",
 *     serverPort: 15868,
 *     timeout: 5,
 * });
 * ```
 *
 * ## Import
 *
 * WebProxy Wisp can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:webproxy/wisp:Wisp labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:webproxy/wisp:Wisp labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Wisp extends pulumi.CustomResource {
    /**
     * Get an existing Wisp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WispState, opts?: pulumi.CustomResourceOptions): Wisp {
        return new Wisp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:webproxy/wisp:Wisp';

    /**
     * Returns true if the given object is an instance of Wisp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Wisp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Wisp.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Maximum number of web proxy WISP connections (4 - 4096, default = 64).
     */
    public readonly maxConnections!: pulumi.Output<number>;
    /**
     * Server name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * WISP outgoing IP address.
     */
    public readonly outgoingIp!: pulumi.Output<string>;
    /**
     * WISP server IP address.
     */
    public readonly serverIp!: pulumi.Output<string>;
    /**
     * WISP server port (1 - 65535, default = 15868).
     */
    public readonly serverPort!: pulumi.Output<number>;
    /**
     * Period of time before WISP requests time out (1 - 15 sec, default = 5).
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Wisp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WispArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WispArgs | WispState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WispState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["maxConnections"] = state ? state.maxConnections : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outgoingIp"] = state ? state.outgoingIp : undefined;
            resourceInputs["serverIp"] = state ? state.serverIp : undefined;
            resourceInputs["serverPort"] = state ? state.serverPort : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WispArgs | undefined;
            if ((!args || args.serverIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverIp'");
            }
            if ((!args || args.serverPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverPort'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["maxConnections"] = args ? args.maxConnections : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outgoingIp"] = args ? args.outgoingIp : undefined;
            resourceInputs["serverIp"] = args ? args.serverIp : undefined;
            resourceInputs["serverPort"] = args ? args.serverPort : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Wisp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Wisp resources.
 */
export interface WispState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Maximum number of web proxy WISP connections (4 - 4096, default = 64).
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * WISP outgoing IP address.
     */
    outgoingIp?: pulumi.Input<string>;
    /**
     * WISP server IP address.
     */
    serverIp?: pulumi.Input<string>;
    /**
     * WISP server port (1 - 65535, default = 15868).
     */
    serverPort?: pulumi.Input<number>;
    /**
     * Period of time before WISP requests time out (1 - 15 sec, default = 5).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Wisp resource.
 */
export interface WispArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Maximum number of web proxy WISP connections (4 - 4096, default = 64).
     */
    maxConnections?: pulumi.Input<number>;
    /**
     * Server name.
     */
    name?: pulumi.Input<string>;
    /**
     * WISP outgoing IP address.
     */
    outgoingIp?: pulumi.Input<string>;
    /**
     * WISP server IP address.
     */
    serverIp: pulumi.Input<string>;
    /**
     * WISP server port (1 - 65535, default = 15868).
     */
    serverPort: pulumi.Input<number>;
    /**
     * Period of time before WISP requests time out (1 - 15 sec, default = 5).
     */
    timeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
