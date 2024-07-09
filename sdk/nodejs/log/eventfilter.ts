// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure log event filters.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.log.Eventfilter("trname", {
 *     complianceCheck: "enable",
 *     endpoint: "enable",
 *     event: "enable",
 *     ha: "enable",
 *     router: "enable",
 *     securityRating: "enable",
 *     system: "enable",
 *     user: "enable",
 *     vpn: "enable",
 *     wanOpt: "enable",
 *     wirelessActivity: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Log Eventfilter can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:log/eventfilter:Eventfilter labelname LogEventfilter
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:log/eventfilter:Eventfilter labelname LogEventfilter
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Eventfilter extends pulumi.CustomResource {
    /**
     * Get an existing Eventfilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventfilterState, opts?: pulumi.CustomResourceOptions): Eventfilter {
        return new Eventfilter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:log/eventfilter:Eventfilter';

    /**
     * Returns true if the given object is an instance of Eventfilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Eventfilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Eventfilter.__pulumiType;
    }

    /**
     * Enable/disable CIFS logging. Valid values: `enable`, `disable`.
     */
    public readonly cifs!: pulumi.Output<string>;
    /**
     * Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
     */
    public readonly complianceCheck!: pulumi.Output<string>;
    /**
     * Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
     */
    public readonly connector!: pulumi.Output<string>;
    /**
     * Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * Enable/disable event logging. Valid values: `enable`, `disable`.
     */
    public readonly event!: pulumi.Output<string>;
    /**
     * Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
     */
    public readonly fortiextender!: pulumi.Output<string>;
    /**
     * Enable/disable ha event logging. Valid values: `enable`, `disable`.
     */
    public readonly ha!: pulumi.Output<string>;
    /**
     * Enable/disable REST API logging. Valid values: `enable`, `disable`.
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * Enable/disable router event logging. Valid values: `enable`, `disable`.
     */
    public readonly router!: pulumi.Output<string>;
    /**
     * Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
     */
    public readonly sdwan!: pulumi.Output<string>;
    /**
     * Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
     */
    public readonly securityRating!: pulumi.Output<string>;
    /**
     * Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
     */
    public readonly switchController!: pulumi.Output<string>;
    /**
     * Enable/disable system event logging. Valid values: `enable`, `disable`.
     */
    public readonly system!: pulumi.Output<string>;
    /**
     * Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
     */
    public readonly user!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Enable/disable VPN event logging. Valid values: `enable`, `disable`.
     */
    public readonly vpn!: pulumi.Output<string>;
    /**
     * Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
     */
    public readonly wanOpt!: pulumi.Output<string>;
    /**
     * Enable/disable web proxy event logging. Valid values: `enable`, `disable`.
     */
    public readonly webproxy!: pulumi.Output<string>;
    /**
     * Enable/disable wireless event logging. Valid values: `enable`, `disable`.
     */
    public readonly wirelessActivity!: pulumi.Output<string>;

    /**
     * Create a Eventfilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EventfilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventfilterArgs | EventfilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventfilterState | undefined;
            resourceInputs["cifs"] = state ? state.cifs : undefined;
            resourceInputs["complianceCheck"] = state ? state.complianceCheck : undefined;
            resourceInputs["connector"] = state ? state.connector : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["event"] = state ? state.event : undefined;
            resourceInputs["fortiextender"] = state ? state.fortiextender : undefined;
            resourceInputs["ha"] = state ? state.ha : undefined;
            resourceInputs["restApi"] = state ? state.restApi : undefined;
            resourceInputs["router"] = state ? state.router : undefined;
            resourceInputs["sdwan"] = state ? state.sdwan : undefined;
            resourceInputs["securityRating"] = state ? state.securityRating : undefined;
            resourceInputs["switchController"] = state ? state.switchController : undefined;
            resourceInputs["system"] = state ? state.system : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vpn"] = state ? state.vpn : undefined;
            resourceInputs["wanOpt"] = state ? state.wanOpt : undefined;
            resourceInputs["webproxy"] = state ? state.webproxy : undefined;
            resourceInputs["wirelessActivity"] = state ? state.wirelessActivity : undefined;
        } else {
            const args = argsOrState as EventfilterArgs | undefined;
            resourceInputs["cifs"] = args ? args.cifs : undefined;
            resourceInputs["complianceCheck"] = args ? args.complianceCheck : undefined;
            resourceInputs["connector"] = args ? args.connector : undefined;
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["event"] = args ? args.event : undefined;
            resourceInputs["fortiextender"] = args ? args.fortiextender : undefined;
            resourceInputs["ha"] = args ? args.ha : undefined;
            resourceInputs["restApi"] = args ? args.restApi : undefined;
            resourceInputs["router"] = args ? args.router : undefined;
            resourceInputs["sdwan"] = args ? args.sdwan : undefined;
            resourceInputs["securityRating"] = args ? args.securityRating : undefined;
            resourceInputs["switchController"] = args ? args.switchController : undefined;
            resourceInputs["system"] = args ? args.system : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vpn"] = args ? args.vpn : undefined;
            resourceInputs["wanOpt"] = args ? args.wanOpt : undefined;
            resourceInputs["webproxy"] = args ? args.webproxy : undefined;
            resourceInputs["wirelessActivity"] = args ? args.wirelessActivity : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Eventfilter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Eventfilter resources.
 */
export interface EventfilterState {
    /**
     * Enable/disable CIFS logging. Valid values: `enable`, `disable`.
     */
    cifs?: pulumi.Input<string>;
    /**
     * Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
     */
    complianceCheck?: pulumi.Input<string>;
    /**
     * Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
     */
    connector?: pulumi.Input<string>;
    /**
     * Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Enable/disable event logging. Valid values: `enable`, `disable`.
     */
    event?: pulumi.Input<string>;
    /**
     * Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
     */
    fortiextender?: pulumi.Input<string>;
    /**
     * Enable/disable ha event logging. Valid values: `enable`, `disable`.
     */
    ha?: pulumi.Input<string>;
    /**
     * Enable/disable REST API logging. Valid values: `enable`, `disable`.
     */
    restApi?: pulumi.Input<string>;
    /**
     * Enable/disable router event logging. Valid values: `enable`, `disable`.
     */
    router?: pulumi.Input<string>;
    /**
     * Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
     */
    sdwan?: pulumi.Input<string>;
    /**
     * Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
     */
    securityRating?: pulumi.Input<string>;
    /**
     * Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
     */
    switchController?: pulumi.Input<string>;
    /**
     * Enable/disable system event logging. Valid values: `enable`, `disable`.
     */
    system?: pulumi.Input<string>;
    /**
     * Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
     */
    user?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable VPN event logging. Valid values: `enable`, `disable`.
     */
    vpn?: pulumi.Input<string>;
    /**
     * Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
     */
    wanOpt?: pulumi.Input<string>;
    /**
     * Enable/disable web proxy event logging. Valid values: `enable`, `disable`.
     */
    webproxy?: pulumi.Input<string>;
    /**
     * Enable/disable wireless event logging. Valid values: `enable`, `disable`.
     */
    wirelessActivity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Eventfilter resource.
 */
export interface EventfilterArgs {
    /**
     * Enable/disable CIFS logging. Valid values: `enable`, `disable`.
     */
    cifs?: pulumi.Input<string>;
    /**
     * Enable/disable PCI DSS compliance check logging. Valid values: `enable`, `disable`.
     */
    complianceCheck?: pulumi.Input<string>;
    /**
     * Enable/disable SDN connector logging. Valid values: `enable`, `disable`.
     */
    connector?: pulumi.Input<string>;
    /**
     * Enable/disable endpoint event logging. Valid values: `enable`, `disable`.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Enable/disable event logging. Valid values: `enable`, `disable`.
     */
    event?: pulumi.Input<string>;
    /**
     * Enable/disable FortiExtender logging. Valid values: `enable`, `disable`.
     */
    fortiextender?: pulumi.Input<string>;
    /**
     * Enable/disable ha event logging. Valid values: `enable`, `disable`.
     */
    ha?: pulumi.Input<string>;
    /**
     * Enable/disable REST API logging. Valid values: `enable`, `disable`.
     */
    restApi?: pulumi.Input<string>;
    /**
     * Enable/disable router event logging. Valid values: `enable`, `disable`.
     */
    router?: pulumi.Input<string>;
    /**
     * Enable/disable SD-WAN logging. Valid values: `enable`, `disable`.
     */
    sdwan?: pulumi.Input<string>;
    /**
     * Enable/disable Security Rating result logging. Valid values: `enable`, `disable`.
     */
    securityRating?: pulumi.Input<string>;
    /**
     * Enable/disable Switch-Controller logging. Valid values: `enable`, `disable`.
     */
    switchController?: pulumi.Input<string>;
    /**
     * Enable/disable system event logging. Valid values: `enable`, `disable`.
     */
    system?: pulumi.Input<string>;
    /**
     * Enable/disable user authentication event logging. Valid values: `enable`, `disable`.
     */
    user?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable VPN event logging. Valid values: `enable`, `disable`.
     */
    vpn?: pulumi.Input<string>;
    /**
     * Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
     */
    wanOpt?: pulumi.Input<string>;
    /**
     * Enable/disable web proxy event logging. Valid values: `enable`, `disable`.
     */
    webproxy?: pulumi.Input<string>;
    /**
     * Enable/disable wireless event logging. Valid values: `enable`, `disable`.
     */
    wirelessActivity?: pulumi.Input<string>;
}
