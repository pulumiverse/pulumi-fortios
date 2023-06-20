// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname1WebproxyForwardserver = new fortios.WebproxyForwardserver("trname1WebproxyForwardserver", {
 *     addrType: "fqdn",
 *     healthcheck: "disable",
 *     ip: "0.0.0.0",
 *     monitor: "http://www.google.com",
 *     port: 1128,
 *     serverDownOption: "block",
 * });
 * const trname1WebproxyForwardservergroup = new fortios.WebproxyForwardservergroup("trname1WebproxyForwardservergroup", {
 *     affinity: "disable",
 *     groupDownOption: "block",
 *     ldbMethod: "weighted",
 *     serverLists: [{
 *         name: trname1WebproxyForwardserver.name,
 *         weight: 12,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * WebProxy ForwardServerGroup can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/webproxyForwardservergroup:WebproxyForwardservergroup labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/webproxyForwardservergroup:WebproxyForwardservergroup labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WebproxyForwardservergroup extends pulumi.CustomResource {
    /**
     * Get an existing WebproxyForwardservergroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebproxyForwardservergroupState, opts?: pulumi.CustomResourceOptions): WebproxyForwardservergroup {
        return new WebproxyForwardservergroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/webproxyForwardservergroup:WebproxyForwardservergroup';

    /**
     * Returns true if the given object is an instance of WebproxyForwardservergroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebproxyForwardservergroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebproxyForwardservergroup.__pulumiType;
    }

    /**
     * Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
     */
    public readonly affinity!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
     */
    public readonly groupDownOption!: pulumi.Output<string>;
    /**
     * Load balance method: weighted or least-session.
     */
    public readonly ldbMethod!: pulumi.Output<string>;
    /**
     * Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
     */
    public readonly serverLists!: pulumi.Output<outputs.WebproxyForwardservergroupServerList[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WebproxyForwardservergroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WebproxyForwardservergroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebproxyForwardservergroupArgs | WebproxyForwardservergroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebproxyForwardservergroupState | undefined;
            resourceInputs["affinity"] = state ? state.affinity : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["groupDownOption"] = state ? state.groupDownOption : undefined;
            resourceInputs["ldbMethod"] = state ? state.ldbMethod : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serverLists"] = state ? state.serverLists : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WebproxyForwardservergroupArgs | undefined;
            resourceInputs["affinity"] = args ? args.affinity : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["groupDownOption"] = args ? args.groupDownOption : undefined;
            resourceInputs["ldbMethod"] = args ? args.ldbMethod : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serverLists"] = args ? args.serverLists : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebproxyForwardservergroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebproxyForwardservergroup resources.
 */
export interface WebproxyForwardservergroupState {
    /**
     * Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
     */
    affinity?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
     */
    groupDownOption?: pulumi.Input<string>;
    /**
     * Load balance method: weighted or least-session.
     */
    ldbMethod?: pulumi.Input<string>;
    /**
     * Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
     */
    name?: pulumi.Input<string>;
    /**
     * Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
     */
    serverLists?: pulumi.Input<pulumi.Input<inputs.WebproxyForwardservergroupServerList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebproxyForwardservergroup resource.
 */
export interface WebproxyForwardservergroupArgs {
    /**
     * Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
     */
    affinity?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
     */
    groupDownOption?: pulumi.Input<string>;
    /**
     * Load balance method: weighted or least-session.
     */
    ldbMethod?: pulumi.Input<string>;
    /**
     * Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
     */
    name?: pulumi.Input<string>;
    /**
     * Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `serverList` block is documented below.
     */
    serverLists?: pulumi.Input<pulumi.Input<inputs.WebproxyForwardservergroupServerList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}