// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * PTP interface-policy configuration. Applies to FortiOS Version `>= 7.4.1`.
 *
 * ## Import
 *
 * SwitchControllerPtp InterfacePolicy can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/ptp/interfacepolicy:Interfacepolicy labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:switchcontroller/ptp/interfacepolicy:Interfacepolicy labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Interfacepolicy extends pulumi.CustomResource {
    /**
     * Get an existing Interfacepolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InterfacepolicyState, opts?: pulumi.CustomResourceOptions): Interfacepolicy {
        return new Interfacepolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/ptp/interfacepolicy:Interfacepolicy';

    /**
     * Returns true if the given object is an instance of Interfacepolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Interfacepolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Interfacepolicy.__pulumiType;
    }

    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * PTP VLAN.
     */
    public readonly vlan!: pulumi.Output<string>;
    /**
     * Configure PTP VLAN priority (0 - 7, default = 4).
     */
    public readonly vlanPri!: pulumi.Output<number>;

    /**
     * Create a Interfacepolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InterfacepolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InterfacepolicyArgs | InterfacepolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InterfacepolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
            resourceInputs["vlanPri"] = state ? state.vlanPri : undefined;
        } else {
            const args = argsOrState as InterfacepolicyArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
            resourceInputs["vlanPri"] = args ? args.vlanPri : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Interfacepolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Interfacepolicy resources.
 */
export interface InterfacepolicyState {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * PTP VLAN.
     */
    vlan?: pulumi.Input<string>;
    /**
     * Configure PTP VLAN priority (0 - 7, default = 4).
     */
    vlanPri?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Interfacepolicy resource.
 */
export interface InterfacepolicyArgs {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * PTP VLAN.
     */
    vlan?: pulumi.Input<string>;
    /**
     * Configure PTP VLAN priority (0 - 7, default = 4).
     */
    vlanPri?: pulumi.Input<number>;
}
