// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * IP blacklist vendor. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * Firewall InternetServiceIpblVendor can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Internetserviceipblvendor extends pulumi.CustomResource {
    /**
     * Get an existing Internetserviceipblvendor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InternetserviceipblvendorState, opts?: pulumi.CustomResourceOptions): Internetserviceipblvendor {
        return new Internetserviceipblvendor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/internetserviceipblvendor:Internetserviceipblvendor';

    /**
     * Returns true if the given object is an instance of Internetserviceipblvendor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Internetserviceipblvendor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Internetserviceipblvendor.__pulumiType;
    }

    /**
     * IP blacklist vendor ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * IP blacklist vendor name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Internetserviceipblvendor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InternetserviceipblvendorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InternetserviceipblvendorArgs | InternetserviceipblvendorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InternetserviceipblvendorState | undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as InternetserviceipblvendorArgs | undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Internetserviceipblvendor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Internetserviceipblvendor resources.
 */
export interface InternetserviceipblvendorState {
    /**
     * IP blacklist vendor ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * IP blacklist vendor name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Internetserviceipblvendor resource.
 */
export interface InternetserviceipblvendorArgs {
    /**
     * IP blacklist vendor ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * IP blacklist vendor name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
