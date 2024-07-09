// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure names for shaping classes. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * Firewall TrafficClass can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Trafficclass extends pulumi.CustomResource {
    /**
     * Get an existing Trafficclass resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrafficclassState, opts?: pulumi.CustomResourceOptions): Trafficclass {
        return new Trafficclass(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/trafficclass:Trafficclass';

    /**
     * Returns true if the given object is an instance of Trafficclass.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trafficclass {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trafficclass.__pulumiType;
    }

    /**
     * Class ID to be named.
     */
    public readonly classId!: pulumi.Output<number>;
    /**
     * Define the name for this class-id.
     */
    public readonly className!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Trafficclass resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrafficclassArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrafficclassArgs | TrafficclassState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrafficclassState | undefined;
            resourceInputs["classId"] = state ? state.classId : undefined;
            resourceInputs["className"] = state ? state.className : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as TrafficclassArgs | undefined;
            if ((!args || args.classId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'classId'");
            }
            resourceInputs["classId"] = args ? args.classId : undefined;
            resourceInputs["className"] = args ? args.className : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Trafficclass.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trafficclass resources.
 */
export interface TrafficclassState {
    /**
     * Class ID to be named.
     */
    classId?: pulumi.Input<number>;
    /**
     * Define the name for this class-id.
     */
    className?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trafficclass resource.
 */
export interface TrafficclassArgs {
    /**
     * Class ID to be named.
     */
    classId: pulumi.Input<number>;
    /**
     * Define the name for this class-id.
     */
    className?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
