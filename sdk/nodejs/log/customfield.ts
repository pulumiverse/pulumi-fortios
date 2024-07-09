// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure custom log fields.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.log.Customfield("trname", {
 *     fosid: "1",
 *     value: "logteststr",
 * });
 * ```
 *
 * ## Import
 *
 * Log CustomField can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:log/customfield:Customfield labelname {{fosid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:log/customfield:Customfield labelname {{fosid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Customfield extends pulumi.CustomResource {
    /**
     * Get an existing Customfield resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomfieldState, opts?: pulumi.CustomResourceOptions): Customfield {
        return new Customfield(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:log/customfield:Customfield';

    /**
     * Returns true if the given object is an instance of Customfield.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Customfield {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Customfield.__pulumiType;
    }

    /**
     * field ID <string>.
     */
    public readonly fosid!: pulumi.Output<string>;
    /**
     * Field name (max: 15 characters).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Field value (max: 15 characters).
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Customfield resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomfieldArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomfieldArgs | CustomfieldState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomfieldState | undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as CustomfieldArgs | undefined;
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Customfield.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Customfield resources.
 */
export interface CustomfieldState {
    /**
     * field ID <string>.
     */
    fosid?: pulumi.Input<string>;
    /**
     * Field name (max: 15 characters).
     */
    name?: pulumi.Input<string>;
    /**
     * Field value (max: 15 characters).
     */
    value?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Customfield resource.
 */
export interface CustomfieldArgs {
    /**
     * field ID <string>.
     */
    fosid?: pulumi.Input<string>;
    /**
     * Field name (max: 15 characters).
     */
    name?: pulumi.Input<string>;
    /**
     * Field value (max: 15 characters).
     */
    value: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
