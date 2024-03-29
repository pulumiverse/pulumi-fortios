// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure predefined data type used by DLP blocking. Applies to FortiOS Version `>= 7.2.0`.
 *
 * ## Import
 *
 * Dlp DataType can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Datatype extends pulumi.CustomResource {
    /**
     * Get an existing Datatype resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatatypeState, opts?: pulumi.CustomResourceOptions): Datatype {
        return new Datatype(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:dlp/datatype:Datatype';

    /**
     * Returns true if the given object is an instance of Datatype.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Datatype {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Datatype.__pulumiType;
    }

    /**
     * Optional comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Number of characters to obtain in advance for verification (1 - 255, default = 1).
     */
    public readonly lookAhead!: pulumi.Output<number>;
    /**
     * Number of characters required to save for verification (1 - 255, default = 1).
     */
    public readonly lookBack!: pulumi.Output<number>;
    /**
     * Number of characters behind for match-around (1 - 4096, default = 1).
     */
    public readonly matchAhead!: pulumi.Output<number>;
    /**
     * Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
     */
    public readonly matchAround!: pulumi.Output<string>;
    /**
     * Number of characters in front for match-around (1 - 4096, default = 1).
     */
    public readonly matchBack!: pulumi.Output<number>;
    /**
     * Name of table containing the data type.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Regular expression pattern string without look around.
     */
    public readonly pattern!: pulumi.Output<string>;
    /**
     * Template to transform user input to a pattern using capture group from 'pattern'.
     */
    public readonly transform!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Regular expression pattern string used to verify the data type.
     */
    public readonly verify!: pulumi.Output<string>;
    /**
     * Extra regular expression pattern string used to verify the data type.
     */
    public readonly verify2!: pulumi.Output<string>;
    /**
     * Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
     */
    public readonly verifyTransformedPattern!: pulumi.Output<string>;

    /**
     * Create a Datatype resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatatypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatatypeArgs | DatatypeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatatypeState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["lookAhead"] = state ? state.lookAhead : undefined;
            resourceInputs["lookBack"] = state ? state.lookBack : undefined;
            resourceInputs["matchAhead"] = state ? state.matchAhead : undefined;
            resourceInputs["matchAround"] = state ? state.matchAround : undefined;
            resourceInputs["matchBack"] = state ? state.matchBack : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pattern"] = state ? state.pattern : undefined;
            resourceInputs["transform"] = state ? state.transform : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["verify"] = state ? state.verify : undefined;
            resourceInputs["verify2"] = state ? state.verify2 : undefined;
            resourceInputs["verifyTransformedPattern"] = state ? state.verifyTransformedPattern : undefined;
        } else {
            const args = argsOrState as DatatypeArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["lookAhead"] = args ? args.lookAhead : undefined;
            resourceInputs["lookBack"] = args ? args.lookBack : undefined;
            resourceInputs["matchAhead"] = args ? args.matchAhead : undefined;
            resourceInputs["matchAround"] = args ? args.matchAround : undefined;
            resourceInputs["matchBack"] = args ? args.matchBack : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pattern"] = args ? args.pattern : undefined;
            resourceInputs["transform"] = args ? args.transform : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["verify"] = args ? args.verify : undefined;
            resourceInputs["verify2"] = args ? args.verify2 : undefined;
            resourceInputs["verifyTransformedPattern"] = args ? args.verifyTransformedPattern : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Datatype.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Datatype resources.
 */
export interface DatatypeState {
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Number of characters to obtain in advance for verification (1 - 255, default = 1).
     */
    lookAhead?: pulumi.Input<number>;
    /**
     * Number of characters required to save for verification (1 - 255, default = 1).
     */
    lookBack?: pulumi.Input<number>;
    /**
     * Number of characters behind for match-around (1 - 4096, default = 1).
     */
    matchAhead?: pulumi.Input<number>;
    /**
     * Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
     */
    matchAround?: pulumi.Input<string>;
    /**
     * Number of characters in front for match-around (1 - 4096, default = 1).
     */
    matchBack?: pulumi.Input<number>;
    /**
     * Name of table containing the data type.
     */
    name?: pulumi.Input<string>;
    /**
     * Regular expression pattern string without look around.
     */
    pattern?: pulumi.Input<string>;
    /**
     * Template to transform user input to a pattern using capture group from 'pattern'.
     */
    transform?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Regular expression pattern string used to verify the data type.
     */
    verify?: pulumi.Input<string>;
    /**
     * Extra regular expression pattern string used to verify the data type.
     */
    verify2?: pulumi.Input<string>;
    /**
     * Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
     */
    verifyTransformedPattern?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Datatype resource.
 */
export interface DatatypeArgs {
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Number of characters to obtain in advance for verification (1 - 255, default = 1).
     */
    lookAhead?: pulumi.Input<number>;
    /**
     * Number of characters required to save for verification (1 - 255, default = 1).
     */
    lookBack?: pulumi.Input<number>;
    /**
     * Number of characters behind for match-around (1 - 4096, default = 1).
     */
    matchAhead?: pulumi.Input<number>;
    /**
     * Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
     */
    matchAround?: pulumi.Input<string>;
    /**
     * Number of characters in front for match-around (1 - 4096, default = 1).
     */
    matchBack?: pulumi.Input<number>;
    /**
     * Name of table containing the data type.
     */
    name?: pulumi.Input<string>;
    /**
     * Regular expression pattern string without look around.
     */
    pattern?: pulumi.Input<string>;
    /**
     * Template to transform user input to a pattern using capture group from 'pattern'.
     */
    transform?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Regular expression pattern string used to verify the data type.
     */
    verify?: pulumi.Input<string>;
    /**
     * Extra regular expression pattern string used to verify the data type.
     */
    verify2?: pulumi.Input<string>;
    /**
     * Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
     */
    verifyTransformedPattern?: pulumi.Input<string>;
}
