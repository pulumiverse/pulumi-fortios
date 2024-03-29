// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure router settings.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.router.Setting("trname", {hostname: "s1"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Router Setting can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:router/setting:Setting labelname RouterSetting
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:router/setting:Setting labelname RouterSetting
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Setting extends pulumi.CustomResource {
    /**
     * Get an existing Setting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SettingState, opts?: pulumi.CustomResourceOptions): Setting {
        return new Setting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:router/setting:Setting';

    /**
     * Returns true if the given object is an instance of Setting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Setting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Setting.__pulumiType;
    }

    /**
     * bgp_debug_flags
     */
    public readonly bgpDebugFlags!: pulumi.Output<string>;
    /**
     * Hostname for this virtual domain router.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * igmp_debug_flags
     */
    public readonly igmpDebugFlags!: pulumi.Output<string>;
    /**
     * imi_debug_flags
     */
    public readonly imiDebugFlags!: pulumi.Output<string>;
    /**
     * isis_debug_flags
     */
    public readonly isisDebugFlags!: pulumi.Output<string>;
    /**
     * ospf6_debug_events_flags
     */
    public readonly ospf6DebugEventsFlags!: pulumi.Output<string>;
    /**
     * ospf6_debug_ifsm_flags
     */
    public readonly ospf6DebugIfsmFlags!: pulumi.Output<string>;
    /**
     * ospf6_debug_lsa_flags
     */
    public readonly ospf6DebugLsaFlags!: pulumi.Output<string>;
    /**
     * ospf6_debug_nfsm_flags
     */
    public readonly ospf6DebugNfsmFlags!: pulumi.Output<string>;
    /**
     * ospf6_debug_nsm_flags
     */
    public readonly ospf6DebugNsmFlags!: pulumi.Output<string>;
    /**
     * ospf6_debug_packet_flags
     */
    public readonly ospf6DebugPacketFlags!: pulumi.Output<string>;
    /**
     * ospf6_debug_route_flags
     */
    public readonly ospf6DebugRouteFlags!: pulumi.Output<string>;
    /**
     * ospf_debug_events_flags
     */
    public readonly ospfDebugEventsFlags!: pulumi.Output<string>;
    /**
     * ospf_debug_ifsm_flags
     */
    public readonly ospfDebugIfsmFlags!: pulumi.Output<string>;
    /**
     * ospf_debug_lsa_flags
     */
    public readonly ospfDebugLsaFlags!: pulumi.Output<string>;
    /**
     * ospf_debug_nfsm_flags
     */
    public readonly ospfDebugNfsmFlags!: pulumi.Output<string>;
    /**
     * ospf_debug_nsm_flags
     */
    public readonly ospfDebugNsmFlags!: pulumi.Output<string>;
    /**
     * ospf_debug_packet_flags
     */
    public readonly ospfDebugPacketFlags!: pulumi.Output<string>;
    /**
     * ospf_debug_route_flags
     */
    public readonly ospfDebugRouteFlags!: pulumi.Output<string>;
    /**
     * pimdm_debug_flags
     */
    public readonly pimdmDebugFlags!: pulumi.Output<string>;
    /**
     * pimsm_debug_joinprune_flags
     */
    public readonly pimsmDebugJoinpruneFlags!: pulumi.Output<string>;
    /**
     * pimsm_debug_simple_flags
     */
    public readonly pimsmDebugSimpleFlags!: pulumi.Output<string>;
    /**
     * pimsm_debug_timer_flags
     */
    public readonly pimsmDebugTimerFlags!: pulumi.Output<string>;
    /**
     * rip_debug_flags
     */
    public readonly ripDebugFlags!: pulumi.Output<string>;
    /**
     * ripng_debug_flags
     */
    public readonly ripngDebugFlags!: pulumi.Output<string>;
    /**
     * Prefix-list as filter for showing routes.
     */
    public readonly showFilter!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Setting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SettingArgs | SettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SettingState | undefined;
            resourceInputs["bgpDebugFlags"] = state ? state.bgpDebugFlags : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["igmpDebugFlags"] = state ? state.igmpDebugFlags : undefined;
            resourceInputs["imiDebugFlags"] = state ? state.imiDebugFlags : undefined;
            resourceInputs["isisDebugFlags"] = state ? state.isisDebugFlags : undefined;
            resourceInputs["ospf6DebugEventsFlags"] = state ? state.ospf6DebugEventsFlags : undefined;
            resourceInputs["ospf6DebugIfsmFlags"] = state ? state.ospf6DebugIfsmFlags : undefined;
            resourceInputs["ospf6DebugLsaFlags"] = state ? state.ospf6DebugLsaFlags : undefined;
            resourceInputs["ospf6DebugNfsmFlags"] = state ? state.ospf6DebugNfsmFlags : undefined;
            resourceInputs["ospf6DebugNsmFlags"] = state ? state.ospf6DebugNsmFlags : undefined;
            resourceInputs["ospf6DebugPacketFlags"] = state ? state.ospf6DebugPacketFlags : undefined;
            resourceInputs["ospf6DebugRouteFlags"] = state ? state.ospf6DebugRouteFlags : undefined;
            resourceInputs["ospfDebugEventsFlags"] = state ? state.ospfDebugEventsFlags : undefined;
            resourceInputs["ospfDebugIfsmFlags"] = state ? state.ospfDebugIfsmFlags : undefined;
            resourceInputs["ospfDebugLsaFlags"] = state ? state.ospfDebugLsaFlags : undefined;
            resourceInputs["ospfDebugNfsmFlags"] = state ? state.ospfDebugNfsmFlags : undefined;
            resourceInputs["ospfDebugNsmFlags"] = state ? state.ospfDebugNsmFlags : undefined;
            resourceInputs["ospfDebugPacketFlags"] = state ? state.ospfDebugPacketFlags : undefined;
            resourceInputs["ospfDebugRouteFlags"] = state ? state.ospfDebugRouteFlags : undefined;
            resourceInputs["pimdmDebugFlags"] = state ? state.pimdmDebugFlags : undefined;
            resourceInputs["pimsmDebugJoinpruneFlags"] = state ? state.pimsmDebugJoinpruneFlags : undefined;
            resourceInputs["pimsmDebugSimpleFlags"] = state ? state.pimsmDebugSimpleFlags : undefined;
            resourceInputs["pimsmDebugTimerFlags"] = state ? state.pimsmDebugTimerFlags : undefined;
            resourceInputs["ripDebugFlags"] = state ? state.ripDebugFlags : undefined;
            resourceInputs["ripngDebugFlags"] = state ? state.ripngDebugFlags : undefined;
            resourceInputs["showFilter"] = state ? state.showFilter : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SettingArgs | undefined;
            resourceInputs["bgpDebugFlags"] = args ? args.bgpDebugFlags : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["igmpDebugFlags"] = args ? args.igmpDebugFlags : undefined;
            resourceInputs["imiDebugFlags"] = args ? args.imiDebugFlags : undefined;
            resourceInputs["isisDebugFlags"] = args ? args.isisDebugFlags : undefined;
            resourceInputs["ospf6DebugEventsFlags"] = args ? args.ospf6DebugEventsFlags : undefined;
            resourceInputs["ospf6DebugIfsmFlags"] = args ? args.ospf6DebugIfsmFlags : undefined;
            resourceInputs["ospf6DebugLsaFlags"] = args ? args.ospf6DebugLsaFlags : undefined;
            resourceInputs["ospf6DebugNfsmFlags"] = args ? args.ospf6DebugNfsmFlags : undefined;
            resourceInputs["ospf6DebugNsmFlags"] = args ? args.ospf6DebugNsmFlags : undefined;
            resourceInputs["ospf6DebugPacketFlags"] = args ? args.ospf6DebugPacketFlags : undefined;
            resourceInputs["ospf6DebugRouteFlags"] = args ? args.ospf6DebugRouteFlags : undefined;
            resourceInputs["ospfDebugEventsFlags"] = args ? args.ospfDebugEventsFlags : undefined;
            resourceInputs["ospfDebugIfsmFlags"] = args ? args.ospfDebugIfsmFlags : undefined;
            resourceInputs["ospfDebugLsaFlags"] = args ? args.ospfDebugLsaFlags : undefined;
            resourceInputs["ospfDebugNfsmFlags"] = args ? args.ospfDebugNfsmFlags : undefined;
            resourceInputs["ospfDebugNsmFlags"] = args ? args.ospfDebugNsmFlags : undefined;
            resourceInputs["ospfDebugPacketFlags"] = args ? args.ospfDebugPacketFlags : undefined;
            resourceInputs["ospfDebugRouteFlags"] = args ? args.ospfDebugRouteFlags : undefined;
            resourceInputs["pimdmDebugFlags"] = args ? args.pimdmDebugFlags : undefined;
            resourceInputs["pimsmDebugJoinpruneFlags"] = args ? args.pimsmDebugJoinpruneFlags : undefined;
            resourceInputs["pimsmDebugSimpleFlags"] = args ? args.pimsmDebugSimpleFlags : undefined;
            resourceInputs["pimsmDebugTimerFlags"] = args ? args.pimsmDebugTimerFlags : undefined;
            resourceInputs["ripDebugFlags"] = args ? args.ripDebugFlags : undefined;
            resourceInputs["ripngDebugFlags"] = args ? args.ripngDebugFlags : undefined;
            resourceInputs["showFilter"] = args ? args.showFilter : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Setting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Setting resources.
 */
export interface SettingState {
    /**
     * bgp_debug_flags
     */
    bgpDebugFlags?: pulumi.Input<string>;
    /**
     * Hostname for this virtual domain router.
     */
    hostname?: pulumi.Input<string>;
    /**
     * igmp_debug_flags
     */
    igmpDebugFlags?: pulumi.Input<string>;
    /**
     * imi_debug_flags
     */
    imiDebugFlags?: pulumi.Input<string>;
    /**
     * isis_debug_flags
     */
    isisDebugFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_events_flags
     */
    ospf6DebugEventsFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_ifsm_flags
     */
    ospf6DebugIfsmFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_lsa_flags
     */
    ospf6DebugLsaFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_nfsm_flags
     */
    ospf6DebugNfsmFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_nsm_flags
     */
    ospf6DebugNsmFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_packet_flags
     */
    ospf6DebugPacketFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_route_flags
     */
    ospf6DebugRouteFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_events_flags
     */
    ospfDebugEventsFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_ifsm_flags
     */
    ospfDebugIfsmFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_lsa_flags
     */
    ospfDebugLsaFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_nfsm_flags
     */
    ospfDebugNfsmFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_nsm_flags
     */
    ospfDebugNsmFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_packet_flags
     */
    ospfDebugPacketFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_route_flags
     */
    ospfDebugRouteFlags?: pulumi.Input<string>;
    /**
     * pimdm_debug_flags
     */
    pimdmDebugFlags?: pulumi.Input<string>;
    /**
     * pimsm_debug_joinprune_flags
     */
    pimsmDebugJoinpruneFlags?: pulumi.Input<string>;
    /**
     * pimsm_debug_simple_flags
     */
    pimsmDebugSimpleFlags?: pulumi.Input<string>;
    /**
     * pimsm_debug_timer_flags
     */
    pimsmDebugTimerFlags?: pulumi.Input<string>;
    /**
     * rip_debug_flags
     */
    ripDebugFlags?: pulumi.Input<string>;
    /**
     * ripng_debug_flags
     */
    ripngDebugFlags?: pulumi.Input<string>;
    /**
     * Prefix-list as filter for showing routes.
     */
    showFilter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Setting resource.
 */
export interface SettingArgs {
    /**
     * bgp_debug_flags
     */
    bgpDebugFlags?: pulumi.Input<string>;
    /**
     * Hostname for this virtual domain router.
     */
    hostname?: pulumi.Input<string>;
    /**
     * igmp_debug_flags
     */
    igmpDebugFlags?: pulumi.Input<string>;
    /**
     * imi_debug_flags
     */
    imiDebugFlags?: pulumi.Input<string>;
    /**
     * isis_debug_flags
     */
    isisDebugFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_events_flags
     */
    ospf6DebugEventsFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_ifsm_flags
     */
    ospf6DebugIfsmFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_lsa_flags
     */
    ospf6DebugLsaFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_nfsm_flags
     */
    ospf6DebugNfsmFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_nsm_flags
     */
    ospf6DebugNsmFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_packet_flags
     */
    ospf6DebugPacketFlags?: pulumi.Input<string>;
    /**
     * ospf6_debug_route_flags
     */
    ospf6DebugRouteFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_events_flags
     */
    ospfDebugEventsFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_ifsm_flags
     */
    ospfDebugIfsmFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_lsa_flags
     */
    ospfDebugLsaFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_nfsm_flags
     */
    ospfDebugNfsmFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_nsm_flags
     */
    ospfDebugNsmFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_packet_flags
     */
    ospfDebugPacketFlags?: pulumi.Input<string>;
    /**
     * ospf_debug_route_flags
     */
    ospfDebugRouteFlags?: pulumi.Input<string>;
    /**
     * pimdm_debug_flags
     */
    pimdmDebugFlags?: pulumi.Input<string>;
    /**
     * pimsm_debug_joinprune_flags
     */
    pimsmDebugJoinpruneFlags?: pulumi.Input<string>;
    /**
     * pimsm_debug_simple_flags
     */
    pimsmDebugSimpleFlags?: pulumi.Input<string>;
    /**
     * pimsm_debug_timer_flags
     */
    pimsmDebugTimerFlags?: pulumi.Input<string>;
    /**
     * rip_debug_flags
     */
    ripDebugFlags?: pulumi.Input<string>;
    /**
     * ripng_debug_flags
     */
    ripngDebugFlags?: pulumi.Input<string>;
    /**
     * Prefix-list as filter for showing routes.
     */
    showFilter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
