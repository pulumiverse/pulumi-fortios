// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure port policy to be applied on the managed FortiSwitch ports through NAC device. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,7.0.0`.
 *
 * ## Import
 *
 * SwitchController PortPolicy can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/portpolicy:Portpolicy labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/portpolicy:Portpolicy labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Portpolicy extends pulumi.CustomResource {
    /**
     * Get an existing Portpolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortpolicyState, opts?: pulumi.CustomResourceOptions): Portpolicy {
        return new Portpolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/portpolicy:Portpolicy';

    /**
     * Returns true if the given object is an instance of Portpolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Portpolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Portpolicy.__pulumiType;
    }

    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
     */
    public readonly bouncePortLink!: pulumi.Output<string>;
    /**
     * Description for the port policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * FortiLink interface for which this port policy belongs to.
     */
    public readonly fortilink!: pulumi.Output<string>;
    /**
     * LLDP profile to be applied when using this port-policy.
     */
    public readonly lldpProfile!: pulumi.Output<string>;
    /**
     * 802.1x security policy to be applied when using this port-policy.
     */
    public readonly n8021x!: pulumi.Output<string>;
    /**
     * Port policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * QoS policy to be applied when using this port-policy.
     */
    public readonly qosPolicy!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * VLAN policy to be applied when using this port-policy.
     */
    public readonly vlanPolicy!: pulumi.Output<string>;

    /**
     * Create a Portpolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PortpolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortpolicyArgs | PortpolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PortpolicyState | undefined;
            resourceInputs["bouncePortLink"] = state ? state.bouncePortLink : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["fortilink"] = state ? state.fortilink : undefined;
            resourceInputs["lldpProfile"] = state ? state.lldpProfile : undefined;
            resourceInputs["n8021x"] = state ? state.n8021x : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["qosPolicy"] = state ? state.qosPolicy : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlanPolicy"] = state ? state.vlanPolicy : undefined;
        } else {
            const args = argsOrState as PortpolicyArgs | undefined;
            resourceInputs["bouncePortLink"] = args ? args.bouncePortLink : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["fortilink"] = args ? args.fortilink : undefined;
            resourceInputs["lldpProfile"] = args ? args.lldpProfile : undefined;
            resourceInputs["n8021x"] = args ? args.n8021x : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["qosPolicy"] = args ? args.qosPolicy : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlanPolicy"] = args ? args.vlanPolicy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Portpolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Portpolicy resources.
 */
export interface PortpolicyState {
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
     */
    bouncePortLink?: pulumi.Input<string>;
    /**
     * Description for the port policy.
     */
    description?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this port policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * LLDP profile to be applied when using this port-policy.
     */
    lldpProfile?: pulumi.Input<string>;
    /**
     * 802.1x security policy to be applied when using this port-policy.
     */
    n8021x?: pulumi.Input<string>;
    /**
     * Port policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * QoS policy to be applied when using this port-policy.
     */
    qosPolicy?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN policy to be applied when using this port-policy.
     */
    vlanPolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Portpolicy resource.
 */
export interface PortpolicyArgs {
    /**
     * Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
     */
    bouncePortLink?: pulumi.Input<string>;
    /**
     * Description for the port policy.
     */
    description?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this port policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * LLDP profile to be applied when using this port-policy.
     */
    lldpProfile?: pulumi.Input<string>;
    /**
     * 802.1x security policy to be applied when using this port-policy.
     */
    n8021x?: pulumi.Input<string>;
    /**
     * Port policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * QoS policy to be applied when using this port-policy.
     */
    qosPolicy?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN policy to be applied when using this port-policy.
     */
    vlanPolicy?: pulumi.Input<string>;
}
