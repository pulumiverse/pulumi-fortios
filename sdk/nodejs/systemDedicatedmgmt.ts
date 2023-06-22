// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure dedicated management.
 *
 * ## Import
 *
 * System DedicatedMgmt can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt labelname SystemDedicatedMgmt
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt labelname SystemDedicatedMgmt
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemDedicatedmgmt extends pulumi.CustomResource {
    /**
     * Get an existing SystemDedicatedmgmt resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemDedicatedmgmtState, opts?: pulumi.CustomResourceOptions): SystemDedicatedmgmt {
        return new SystemDedicatedmgmt(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt';

    /**
     * Returns true if the given object is an instance of SystemDedicatedmgmt.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemDedicatedmgmt {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemDedicatedmgmt.__pulumiType;
    }

    /**
     * Default gateway for dedicated management interface.
     */
    public readonly defaultGateway!: pulumi.Output<string>;
    /**
     * DHCP end IP for dedicated management.
     */
    public readonly dhcpEndIp!: pulumi.Output<string>;
    /**
     * DHCP netmask.
     */
    public readonly dhcpNetmask!: pulumi.Output<string>;
    /**
     * Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
     */
    public readonly dhcpServer!: pulumi.Output<string>;
    /**
     * DHCP start IP for dedicated management.
     */
    public readonly dhcpStartIp!: pulumi.Output<string>;
    /**
     * Dedicated management interface.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Enable/disable dedicated management. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemDedicatedmgmt resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemDedicatedmgmtArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemDedicatedmgmtArgs | SystemDedicatedmgmtState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemDedicatedmgmtState | undefined;
            resourceInputs["defaultGateway"] = state ? state.defaultGateway : undefined;
            resourceInputs["dhcpEndIp"] = state ? state.dhcpEndIp : undefined;
            resourceInputs["dhcpNetmask"] = state ? state.dhcpNetmask : undefined;
            resourceInputs["dhcpServer"] = state ? state.dhcpServer : undefined;
            resourceInputs["dhcpStartIp"] = state ? state.dhcpStartIp : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemDedicatedmgmtArgs | undefined;
            resourceInputs["defaultGateway"] = args ? args.defaultGateway : undefined;
            resourceInputs["dhcpEndIp"] = args ? args.dhcpEndIp : undefined;
            resourceInputs["dhcpNetmask"] = args ? args.dhcpNetmask : undefined;
            resourceInputs["dhcpServer"] = args ? args.dhcpServer : undefined;
            resourceInputs["dhcpStartIp"] = args ? args.dhcpStartIp : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemDedicatedmgmt.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemDedicatedmgmt resources.
 */
export interface SystemDedicatedmgmtState {
    /**
     * Default gateway for dedicated management interface.
     */
    defaultGateway?: pulumi.Input<string>;
    /**
     * DHCP end IP for dedicated management.
     */
    dhcpEndIp?: pulumi.Input<string>;
    /**
     * DHCP netmask.
     */
    dhcpNetmask?: pulumi.Input<string>;
    /**
     * Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
     */
    dhcpServer?: pulumi.Input<string>;
    /**
     * DHCP start IP for dedicated management.
     */
    dhcpStartIp?: pulumi.Input<string>;
    /**
     * Dedicated management interface.
     */
    interface?: pulumi.Input<string>;
    /**
     * Enable/disable dedicated management. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemDedicatedmgmt resource.
 */
export interface SystemDedicatedmgmtArgs {
    /**
     * Default gateway for dedicated management interface.
     */
    defaultGateway?: pulumi.Input<string>;
    /**
     * DHCP end IP for dedicated management.
     */
    dhcpEndIp?: pulumi.Input<string>;
    /**
     * DHCP netmask.
     */
    dhcpNetmask?: pulumi.Input<string>;
    /**
     * Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
     */
    dhcpServer?: pulumi.Input<string>;
    /**
     * DHCP start IP for dedicated management.
     */
    dhcpStartIp?: pulumi.Input<string>;
    /**
     * Dedicated management interface.
     */
    interface?: pulumi.Input<string>;
    /**
     * Enable/disable dedicated management. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
