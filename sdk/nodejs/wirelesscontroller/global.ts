// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure wireless controller global settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.wirelesscontroller.Global("trname", {
 *     apLogServer: "disable",
 *     apLogServerIp: "0.0.0.0",
 *     apLogServerPort: 0,
 *     controlMessageOffload: "ebp-frame aeroscout-tag ap-list sta-list sta-cap-list stats aeroscout-mu",
 *     dataEthernetIi: "disable",
 *     discoveryMcAddr: "224.0.1.140",
 *     fiappEthType: 5252,
 *     imageDownload: "enable",
 *     ipsecBaseIp: "169.254.0.1",
 *     linkAggregation: "disable",
 *     maxClients: 0,
 *     maxRetransmit: 3,
 *     meshEthType: 8755,
 *     rogueScanMacAdjacency: 7,
 *     wtpShare: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * WirelessController Global can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:wirelesscontroller/global:Global labelname WirelessControllerGlobal
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:wirelesscontroller/global:Global labelname WirelessControllerGlobal
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Global extends pulumi.CustomResource {
    /**
     * Get an existing Global resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalState, opts?: pulumi.CustomResourceOptions): Global {
        return new Global(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wirelesscontroller/global:Global';

    /**
     * Returns true if the given object is an instance of Global.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Global {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Global.__pulumiType;
    }

    /**
     * Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly apLogServer!: pulumi.Output<string>;
    /**
     * IP address that APs or FortiAPs send log messages to.
     */
    public readonly apLogServerIp!: pulumi.Output<string>;
    /**
     * Port that APs or FortiAPs send log messages to.
     */
    public readonly apLogServerPort!: pulumi.Output<number>;
    /**
     * Configure CAPWAP control message data channel offload.
     */
    public readonly controlMessageOffload!: pulumi.Output<string>;
    /**
     * Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly dataEthernetIi!: pulumi.Output<string>;
    /**
     * Multicast IP address for AP discovery (default = 244.0.1.140).
     */
    public readonly discoveryMcAddr!: pulumi.Output<string>;
    /**
     * Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
     */
    public readonly fiappEthType!: pulumi.Output<number>;
    /**
     * Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
     */
    public readonly imageDownload!: pulumi.Output<string>;
    /**
     * Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
     */
    public readonly ipsecBaseIp!: pulumi.Output<string>;
    /**
     * Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly linkAggregation!: pulumi.Output<string>;
    /**
     * Description of the location of the wireless controller.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
     */
    public readonly maxClients!: pulumi.Output<number>;
    /**
     * Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
     */
    public readonly maxRetransmit!: pulumi.Output<number>;
    /**
     * Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
     */
    public readonly meshEthType!: pulumi.Output<number>;
    /**
     * Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
     */
    public readonly nacInterval!: pulumi.Output<number>;
    /**
     * Name of the wireless controller.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
     */
    public readonly rogueScanMacAdjacency!: pulumi.Output<number>;
    /**
     * Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
     */
    public readonly tunnelMode!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
     */
    public readonly wtpShare!: pulumi.Output<string>;

    /**
     * Create a Global resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GlobalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalArgs | GlobalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GlobalState | undefined;
            resourceInputs["apLogServer"] = state ? state.apLogServer : undefined;
            resourceInputs["apLogServerIp"] = state ? state.apLogServerIp : undefined;
            resourceInputs["apLogServerPort"] = state ? state.apLogServerPort : undefined;
            resourceInputs["controlMessageOffload"] = state ? state.controlMessageOffload : undefined;
            resourceInputs["dataEthernetIi"] = state ? state.dataEthernetIi : undefined;
            resourceInputs["discoveryMcAddr"] = state ? state.discoveryMcAddr : undefined;
            resourceInputs["fiappEthType"] = state ? state.fiappEthType : undefined;
            resourceInputs["imageDownload"] = state ? state.imageDownload : undefined;
            resourceInputs["ipsecBaseIp"] = state ? state.ipsecBaseIp : undefined;
            resourceInputs["linkAggregation"] = state ? state.linkAggregation : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["maxClients"] = state ? state.maxClients : undefined;
            resourceInputs["maxRetransmit"] = state ? state.maxRetransmit : undefined;
            resourceInputs["meshEthType"] = state ? state.meshEthType : undefined;
            resourceInputs["nacInterval"] = state ? state.nacInterval : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rogueScanMacAdjacency"] = state ? state.rogueScanMacAdjacency : undefined;
            resourceInputs["tunnelMode"] = state ? state.tunnelMode : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wtpShare"] = state ? state.wtpShare : undefined;
        } else {
            const args = argsOrState as GlobalArgs | undefined;
            resourceInputs["apLogServer"] = args ? args.apLogServer : undefined;
            resourceInputs["apLogServerIp"] = args ? args.apLogServerIp : undefined;
            resourceInputs["apLogServerPort"] = args ? args.apLogServerPort : undefined;
            resourceInputs["controlMessageOffload"] = args ? args.controlMessageOffload : undefined;
            resourceInputs["dataEthernetIi"] = args ? args.dataEthernetIi : undefined;
            resourceInputs["discoveryMcAddr"] = args ? args.discoveryMcAddr : undefined;
            resourceInputs["fiappEthType"] = args ? args.fiappEthType : undefined;
            resourceInputs["imageDownload"] = args ? args.imageDownload : undefined;
            resourceInputs["ipsecBaseIp"] = args ? args.ipsecBaseIp : undefined;
            resourceInputs["linkAggregation"] = args ? args.linkAggregation : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["maxClients"] = args ? args.maxClients : undefined;
            resourceInputs["maxRetransmit"] = args ? args.maxRetransmit : undefined;
            resourceInputs["meshEthType"] = args ? args.meshEthType : undefined;
            resourceInputs["nacInterval"] = args ? args.nacInterval : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rogueScanMacAdjacency"] = args ? args.rogueScanMacAdjacency : undefined;
            resourceInputs["tunnelMode"] = args ? args.tunnelMode : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wtpShare"] = args ? args.wtpShare : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Global.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Global resources.
 */
export interface GlobalState {
    /**
     * Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
     */
    apLogServer?: pulumi.Input<string>;
    /**
     * IP address that APs or FortiAPs send log messages to.
     */
    apLogServerIp?: pulumi.Input<string>;
    /**
     * Port that APs or FortiAPs send log messages to.
     */
    apLogServerPort?: pulumi.Input<number>;
    /**
     * Configure CAPWAP control message data channel offload.
     */
    controlMessageOffload?: pulumi.Input<string>;
    /**
     * Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
     */
    dataEthernetIi?: pulumi.Input<string>;
    /**
     * Multicast IP address for AP discovery (default = 244.0.1.140).
     */
    discoveryMcAddr?: pulumi.Input<string>;
    /**
     * Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
     */
    fiappEthType?: pulumi.Input<number>;
    /**
     * Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
     */
    imageDownload?: pulumi.Input<string>;
    /**
     * Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
     */
    ipsecBaseIp?: pulumi.Input<string>;
    /**
     * Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
     */
    linkAggregation?: pulumi.Input<string>;
    /**
     * Description of the location of the wireless controller.
     */
    location?: pulumi.Input<string>;
    /**
     * Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
     */
    maxClients?: pulumi.Input<number>;
    /**
     * Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
     */
    maxRetransmit?: pulumi.Input<number>;
    /**
     * Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
     */
    meshEthType?: pulumi.Input<number>;
    /**
     * Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
     */
    nacInterval?: pulumi.Input<number>;
    /**
     * Name of the wireless controller.
     */
    name?: pulumi.Input<string>;
    /**
     * Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
     */
    rogueScanMacAdjacency?: pulumi.Input<number>;
    /**
     * Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
     */
    tunnelMode?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
     */
    wtpShare?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Global resource.
 */
export interface GlobalArgs {
    /**
     * Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
     */
    apLogServer?: pulumi.Input<string>;
    /**
     * IP address that APs or FortiAPs send log messages to.
     */
    apLogServerIp?: pulumi.Input<string>;
    /**
     * Port that APs or FortiAPs send log messages to.
     */
    apLogServerPort?: pulumi.Input<number>;
    /**
     * Configure CAPWAP control message data channel offload.
     */
    controlMessageOffload?: pulumi.Input<string>;
    /**
     * Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
     */
    dataEthernetIi?: pulumi.Input<string>;
    /**
     * Multicast IP address for AP discovery (default = 244.0.1.140).
     */
    discoveryMcAddr?: pulumi.Input<string>;
    /**
     * Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
     */
    fiappEthType?: pulumi.Input<number>;
    /**
     * Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
     */
    imageDownload?: pulumi.Input<string>;
    /**
     * Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
     */
    ipsecBaseIp?: pulumi.Input<string>;
    /**
     * Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
     */
    linkAggregation?: pulumi.Input<string>;
    /**
     * Description of the location of the wireless controller.
     */
    location?: pulumi.Input<string>;
    /**
     * Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
     */
    maxClients?: pulumi.Input<number>;
    /**
     * Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
     */
    maxRetransmit?: pulumi.Input<number>;
    /**
     * Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
     */
    meshEthType?: pulumi.Input<number>;
    /**
     * Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
     */
    nacInterval?: pulumi.Input<number>;
    /**
     * Name of the wireless controller.
     */
    name?: pulumi.Input<string>;
    /**
     * Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
     */
    rogueScanMacAdjacency?: pulumi.Input<number>;
    /**
     * Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
     */
    tunnelMode?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
     */
    wtpShare?: pulumi.Input<string>;
}
