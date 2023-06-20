// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure the client with its MAC address. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,6.4.10,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6`.
 *
 * ## Import
 *
 * WirelessController Address can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelesscontrollerAddress extends pulumi.CustomResource {
    /**
     * Get an existing WirelesscontrollerAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelesscontrollerAddressState, opts?: pulumi.CustomResourceOptions): WirelesscontrollerAddress {
        return new WirelesscontrollerAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress';

    /**
     * Returns true if the given object is an instance of WirelesscontrollerAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelesscontrollerAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelesscontrollerAddress.__pulumiType;
    }

    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<string>;
    /**
     * MAC address.
     */
    public readonly mac!: pulumi.Output<string>;
    /**
     * Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelesscontrollerAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelesscontrollerAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelesscontrollerAddressArgs | WirelesscontrollerAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelesscontrollerAddressState | undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelesscontrollerAddressArgs | undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["mac"] = args ? args.mac : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelesscontrollerAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelesscontrollerAddress resources.
 */
export interface WirelesscontrollerAddressState {
    /**
     * ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * MAC address.
     */
    mac?: pulumi.Input<string>;
    /**
     * Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
     */
    policy?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelesscontrollerAddress resource.
 */
export interface WirelesscontrollerAddressArgs {
    /**
     * ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * MAC address.
     */
    mac?: pulumi.Input<string>;
    /**
     * Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
     */
    policy?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}