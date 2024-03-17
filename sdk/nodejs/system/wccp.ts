// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure WCCP.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Wccp("trname", {
 *     assignmentBucketFormat: "cisco-implementation",
 *     assignmentDstaddrMask: "0.0.0.0",
 *     assignmentMethod: "HASH",
 *     assignmentSrcaddrMask: "0.0.23.65",
 *     assignmentWeight: 0,
 *     authentication: "disable",
 *     cacheEngineMethod: "GRE",
 *     cacheId: "1.1.1.1",
 *     forwardMethod: "GRE",
 *     groupAddress: "0.0.0.0",
 *     primaryHash: "dst-ip",
 *     priority: 0,
 *     protocol: 0,
 *     returnMethod: "GRE",
 *     routerId: "1.1.1.1",
 *     routerList: "\"1.0.0.0\" ",
 *     serverType: "forward",
 *     serviceId: "1",
 *     serviceType: "auto",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * System Wccp can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/wccp:Wccp labelname {{service_id}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/wccp:Wccp labelname {{service_id}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Wccp extends pulumi.CustomResource {
    /**
     * Get an existing Wccp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WccpState, opts?: pulumi.CustomResourceOptions): Wccp {
        return new Wccp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/wccp:Wccp';

    /**
     * Returns true if the given object is an instance of Wccp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Wccp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Wccp.__pulumiType;
    }

    /**
     * Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
     */
    public readonly assignmentBucketFormat!: pulumi.Output<string>;
    /**
     * Assignment destination address mask.
     */
    public readonly assignmentDstaddrMask!: pulumi.Output<string>;
    /**
     * Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
     */
    public readonly assignmentMethod!: pulumi.Output<string>;
    /**
     * Assignment source address mask.
     */
    public readonly assignmentSrcaddrMask!: pulumi.Output<string>;
    /**
     * Assignment of hash weight/ratio for the WCCP cache engine.
     */
    public readonly assignmentWeight!: pulumi.Output<number>;
    /**
     * Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
     */
    public readonly authentication!: pulumi.Output<string>;
    /**
     * Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
     */
    public readonly cacheEngineMethod!: pulumi.Output<string>;
    /**
     * IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
     */
    public readonly cacheId!: pulumi.Output<string>;
    /**
     * Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
     */
    public readonly forwardMethod!: pulumi.Output<string>;
    /**
     * IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
     */
    public readonly groupAddress!: pulumi.Output<string>;
    /**
     * Password for MD5 authentication.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Service ports.
     */
    public readonly ports!: pulumi.Output<string>;
    /**
     * Match method. Valid values: `source`, `destination`.
     */
    public readonly portsDefined!: pulumi.Output<string>;
    /**
     * Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
     */
    public readonly primaryHash!: pulumi.Output<string>;
    /**
     * Service priority.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Service protocol.
     */
    public readonly protocol!: pulumi.Output<number>;
    /**
     * Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
     */
    public readonly returnMethod!: pulumi.Output<string>;
    /**
     * IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
     */
    public readonly routerId!: pulumi.Output<string>;
    /**
     * IP addresses of one or more WCCP routers.
     */
    public readonly routerList!: pulumi.Output<string>;
    /**
     * IP addresses and netmasks for up to four cache servers.
     */
    public readonly serverList!: pulumi.Output<string>;
    /**
     * Cache server type. Valid values: `forward`, `proxy`.
     */
    public readonly serverType!: pulumi.Output<string>;
    /**
     * Service ID.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
     */
    public readonly serviceType!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Wccp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WccpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WccpArgs | WccpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WccpState | undefined;
            resourceInputs["assignmentBucketFormat"] = state ? state.assignmentBucketFormat : undefined;
            resourceInputs["assignmentDstaddrMask"] = state ? state.assignmentDstaddrMask : undefined;
            resourceInputs["assignmentMethod"] = state ? state.assignmentMethod : undefined;
            resourceInputs["assignmentSrcaddrMask"] = state ? state.assignmentSrcaddrMask : undefined;
            resourceInputs["assignmentWeight"] = state ? state.assignmentWeight : undefined;
            resourceInputs["authentication"] = state ? state.authentication : undefined;
            resourceInputs["cacheEngineMethod"] = state ? state.cacheEngineMethod : undefined;
            resourceInputs["cacheId"] = state ? state.cacheId : undefined;
            resourceInputs["forwardMethod"] = state ? state.forwardMethod : undefined;
            resourceInputs["groupAddress"] = state ? state.groupAddress : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["portsDefined"] = state ? state.portsDefined : undefined;
            resourceInputs["primaryHash"] = state ? state.primaryHash : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["returnMethod"] = state ? state.returnMethod : undefined;
            resourceInputs["routerId"] = state ? state.routerId : undefined;
            resourceInputs["routerList"] = state ? state.routerList : undefined;
            resourceInputs["serverList"] = state ? state.serverList : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceType"] = state ? state.serviceType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WccpArgs | undefined;
            resourceInputs["assignmentBucketFormat"] = args ? args.assignmentBucketFormat : undefined;
            resourceInputs["assignmentDstaddrMask"] = args ? args.assignmentDstaddrMask : undefined;
            resourceInputs["assignmentMethod"] = args ? args.assignmentMethod : undefined;
            resourceInputs["assignmentSrcaddrMask"] = args ? args.assignmentSrcaddrMask : undefined;
            resourceInputs["assignmentWeight"] = args ? args.assignmentWeight : undefined;
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["cacheEngineMethod"] = args ? args.cacheEngineMethod : undefined;
            resourceInputs["cacheId"] = args ? args.cacheId : undefined;
            resourceInputs["forwardMethod"] = args ? args.forwardMethod : undefined;
            resourceInputs["groupAddress"] = args ? args.groupAddress : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["portsDefined"] = args ? args.portsDefined : undefined;
            resourceInputs["primaryHash"] = args ? args.primaryHash : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["returnMethod"] = args ? args.returnMethod : undefined;
            resourceInputs["routerId"] = args ? args.routerId : undefined;
            resourceInputs["routerList"] = args ? args.routerList : undefined;
            resourceInputs["serverList"] = args ? args.serverList : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["serviceType"] = args ? args.serviceType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Wccp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Wccp resources.
 */
export interface WccpState {
    /**
     * Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
     */
    assignmentBucketFormat?: pulumi.Input<string>;
    /**
     * Assignment destination address mask.
     */
    assignmentDstaddrMask?: pulumi.Input<string>;
    /**
     * Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
     */
    assignmentMethod?: pulumi.Input<string>;
    /**
     * Assignment source address mask.
     */
    assignmentSrcaddrMask?: pulumi.Input<string>;
    /**
     * Assignment of hash weight/ratio for the WCCP cache engine.
     */
    assignmentWeight?: pulumi.Input<number>;
    /**
     * Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
     */
    cacheEngineMethod?: pulumi.Input<string>;
    /**
     * IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
     */
    cacheId?: pulumi.Input<string>;
    /**
     * Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
     */
    forwardMethod?: pulumi.Input<string>;
    /**
     * IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
     */
    groupAddress?: pulumi.Input<string>;
    /**
     * Password for MD5 authentication.
     */
    password?: pulumi.Input<string>;
    /**
     * Service ports.
     */
    ports?: pulumi.Input<string>;
    /**
     * Match method. Valid values: `source`, `destination`.
     */
    portsDefined?: pulumi.Input<string>;
    /**
     * Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
     */
    primaryHash?: pulumi.Input<string>;
    /**
     * Service priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Service protocol.
     */
    protocol?: pulumi.Input<number>;
    /**
     * Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
     */
    returnMethod?: pulumi.Input<string>;
    /**
     * IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
     */
    routerId?: pulumi.Input<string>;
    /**
     * IP addresses of one or more WCCP routers.
     */
    routerList?: pulumi.Input<string>;
    /**
     * IP addresses and netmasks for up to four cache servers.
     */
    serverList?: pulumi.Input<string>;
    /**
     * Cache server type. Valid values: `forward`, `proxy`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Service ID.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
     */
    serviceType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Wccp resource.
 */
export interface WccpArgs {
    /**
     * Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
     */
    assignmentBucketFormat?: pulumi.Input<string>;
    /**
     * Assignment destination address mask.
     */
    assignmentDstaddrMask?: pulumi.Input<string>;
    /**
     * Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
     */
    assignmentMethod?: pulumi.Input<string>;
    /**
     * Assignment source address mask.
     */
    assignmentSrcaddrMask?: pulumi.Input<string>;
    /**
     * Assignment of hash weight/ratio for the WCCP cache engine.
     */
    assignmentWeight?: pulumi.Input<number>;
    /**
     * Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
     */
    authentication?: pulumi.Input<string>;
    /**
     * Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
     */
    cacheEngineMethod?: pulumi.Input<string>;
    /**
     * IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
     */
    cacheId?: pulumi.Input<string>;
    /**
     * Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
     */
    forwardMethod?: pulumi.Input<string>;
    /**
     * IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
     */
    groupAddress?: pulumi.Input<string>;
    /**
     * Password for MD5 authentication.
     */
    password?: pulumi.Input<string>;
    /**
     * Service ports.
     */
    ports?: pulumi.Input<string>;
    /**
     * Match method. Valid values: `source`, `destination`.
     */
    portsDefined?: pulumi.Input<string>;
    /**
     * Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
     */
    primaryHash?: pulumi.Input<string>;
    /**
     * Service priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * Service protocol.
     */
    protocol?: pulumi.Input<number>;
    /**
     * Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
     */
    returnMethod?: pulumi.Input<string>;
    /**
     * IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
     */
    routerId?: pulumi.Input<string>;
    /**
     * IP addresses of one or more WCCP routers.
     */
    routerList?: pulumi.Input<string>;
    /**
     * IP addresses and netmasks for up to four cache servers.
     */
    serverList?: pulumi.Input<string>;
    /**
     * Cache server type. Valid values: `forward`, `proxy`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Service ID.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
     */
    serviceType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
