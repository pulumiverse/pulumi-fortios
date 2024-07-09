// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure custom application signatures.
 *
 * ## Import
 *
 * Application Custom can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:application/custom:Custom labelname {{tag}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:application/custom:Custom labelname {{tag}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Custom extends pulumi.CustomResource {
    /**
     * Get an existing Custom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomState, opts?: pulumi.CustomResourceOptions): Custom {
        return new Custom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:application/custom:Custom';

    /**
     * Returns true if the given object is an instance of Custom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Custom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Custom.__pulumiType;
    }

    /**
     * Custom application signature behavior.
     */
    public readonly behavior!: pulumi.Output<string>;
    /**
     * Custom application category ID (use ? to view available options).
     */
    public readonly category!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * Custom application category ID (use ? to view available options).
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Name of this custom application signature.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Custom application signature protocol.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The text that makes up the actual custom application signature.
     */
    public readonly signature!: pulumi.Output<string>;
    /**
     * Signature tag.
     */
    public readonly tag!: pulumi.Output<string>;
    /**
     * Custom application signature technology.
     */
    public readonly technology!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Custom application signature vendor.
     */
    public readonly vendor!: pulumi.Output<string>;

    /**
     * Create a Custom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomArgs | CustomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomState | undefined;
            resourceInputs["behavior"] = state ? state.behavior : undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["signature"] = state ? state.signature : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
            resourceInputs["technology"] = state ? state.technology : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
        } else {
            const args = argsOrState as CustomArgs | undefined;
            if ((!args || args.category === undefined) && !opts.urn) {
                throw new Error("Missing required property 'category'");
            }
            resourceInputs["behavior"] = args ? args.behavior : undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["signature"] = args ? args.signature : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["technology"] = args ? args.technology : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Custom.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Custom resources.
 */
export interface CustomState {
    /**
     * Custom application signature behavior.
     */
    behavior?: pulumi.Input<string>;
    /**
     * Custom application category ID (use ? to view available options).
     */
    category?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Custom application category ID (use ? to view available options).
     */
    fosid?: pulumi.Input<number>;
    /**
     * Name of this custom application signature.
     */
    name?: pulumi.Input<string>;
    /**
     * Custom application signature protocol.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The text that makes up the actual custom application signature.
     */
    signature?: pulumi.Input<string>;
    /**
     * Signature tag.
     */
    tag?: pulumi.Input<string>;
    /**
     * Custom application signature technology.
     */
    technology?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Custom application signature vendor.
     */
    vendor?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Custom resource.
 */
export interface CustomArgs {
    /**
     * Custom application signature behavior.
     */
    behavior?: pulumi.Input<string>;
    /**
     * Custom application category ID (use ? to view available options).
     */
    category: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Custom application category ID (use ? to view available options).
     */
    fosid?: pulumi.Input<number>;
    /**
     * Name of this custom application signature.
     */
    name?: pulumi.Input<string>;
    /**
     * Custom application signature protocol.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The text that makes up the actual custom application signature.
     */
    signature?: pulumi.Input<string>;
    /**
     * Signature tag.
     */
    tag?: pulumi.Input<string>;
    /**
     * Custom application signature technology.
     */
    technology?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Custom application signature vendor.
     */
    vendor?: pulumi.Input<string>;
}
