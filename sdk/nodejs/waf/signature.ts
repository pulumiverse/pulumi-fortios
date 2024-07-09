// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Hidden table for datasource.
 *
 * ## Import
 *
 * Waf Signature can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:waf/signature:Signature labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:waf/signature:Signature labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Signature extends pulumi.CustomResource {
    /**
     * Get an existing Signature resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SignatureState, opts?: pulumi.CustomResourceOptions): Signature {
        return new Signature(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:waf/signature:Signature';

    /**
     * Returns true if the given object is an instance of Signature.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Signature {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Signature.__pulumiType;
    }

    /**
     * Signature description.
     */
    public readonly desc!: pulumi.Output<string>;
    /**
     * Signature ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Signature resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SignatureArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SignatureArgs | SignatureState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SignatureState | undefined;
            resourceInputs["desc"] = state ? state.desc : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SignatureArgs | undefined;
            resourceInputs["desc"] = args ? args.desc : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Signature.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Signature resources.
 */
export interface SignatureState {
    /**
     * Signature description.
     */
    desc?: pulumi.Input<string>;
    /**
     * Signature ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Signature resource.
 */
export interface SignatureArgs {
    /**
     * Signature description.
     */
    desc?: pulumi.Input<string>;
    /**
     * Signature ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
