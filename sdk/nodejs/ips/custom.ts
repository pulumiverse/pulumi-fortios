// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure IPS custom signature.
 *
 * ## Import
 *
 * Ips Custom can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:ips/custom:Custom labelname {{tag}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:ips/custom:Custom labelname {{tag}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
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
    public static readonly __pulumiType = 'fortios:ips/custom:Custom';

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
     * Default action (pass or block) for this signature. Valid values: `pass`, `block`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Applications to be protected. Blank for all applications.
     */
    public readonly application!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * Protect client or server traffic.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Enable/disable logging. Valid values: `disable`, `enable`.
     */
    public readonly log!: pulumi.Output<string>;
    /**
     * Enable/disable packet logging. Valid values: `disable`, `enable`.
     */
    public readonly logPacket!: pulumi.Output<string>;
    /**
     * Operating system(s) that the signature protects. Blank for all operating systems.
     */
    public readonly os!: pulumi.Output<string>;
    /**
     * Protocol(s) that the signature scans. Blank for all protocols.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Signature ID.
     */
    public readonly ruleId!: pulumi.Output<number>;
    /**
     * Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Signature name.
     */
    public readonly sigName!: pulumi.Output<string>;
    /**
     * Custom signature enclosed in single quotes.
     */
    public readonly signature!: pulumi.Output<string>;
    /**
     * Enable/disable this signature. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Signature tag.
     */
    public readonly tag!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Custom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CustomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomArgs | CustomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["application"] = state ? state.application : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["log"] = state ? state.log : undefined;
            resourceInputs["logPacket"] = state ? state.logPacket : undefined;
            resourceInputs["os"] = state ? state.os : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["sigName"] = state ? state.sigName : undefined;
            resourceInputs["signature"] = state ? state.signature : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as CustomArgs | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["application"] = args ? args.application : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["log"] = args ? args.log : undefined;
            resourceInputs["logPacket"] = args ? args.logPacket : undefined;
            resourceInputs["os"] = args ? args.os : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["sigName"] = args ? args.sigName : undefined;
            resourceInputs["signature"] = args ? args.signature : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
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
     * Default action (pass or block) for this signature. Valid values: `pass`, `block`.
     */
    action?: pulumi.Input<string>;
    /**
     * Applications to be protected. Blank for all applications.
     */
    application?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Protect client or server traffic.
     */
    location?: pulumi.Input<string>;
    /**
     * Enable/disable logging. Valid values: `disable`, `enable`.
     */
    log?: pulumi.Input<string>;
    /**
     * Enable/disable packet logging. Valid values: `disable`, `enable`.
     */
    logPacket?: pulumi.Input<string>;
    /**
     * Operating system(s) that the signature protects. Blank for all operating systems.
     */
    os?: pulumi.Input<string>;
    /**
     * Protocol(s) that the signature scans. Blank for all protocols.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Signature ID.
     */
    ruleId?: pulumi.Input<number>;
    /**
     * Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
     */
    severity?: pulumi.Input<string>;
    /**
     * Signature name.
     */
    sigName?: pulumi.Input<string>;
    /**
     * Custom signature enclosed in single quotes.
     */
    signature?: pulumi.Input<string>;
    /**
     * Enable/disable this signature. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Signature tag.
     */
    tag?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Custom resource.
 */
export interface CustomArgs {
    /**
     * Default action (pass or block) for this signature. Valid values: `pass`, `block`.
     */
    action?: pulumi.Input<string>;
    /**
     * Applications to be protected. Blank for all applications.
     */
    application?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Protect client or server traffic.
     */
    location?: pulumi.Input<string>;
    /**
     * Enable/disable logging. Valid values: `disable`, `enable`.
     */
    log?: pulumi.Input<string>;
    /**
     * Enable/disable packet logging. Valid values: `disable`, `enable`.
     */
    logPacket?: pulumi.Input<string>;
    /**
     * Operating system(s) that the signature protects. Blank for all operating systems.
     */
    os?: pulumi.Input<string>;
    /**
     * Protocol(s) that the signature scans. Blank for all protocols.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Signature ID.
     */
    ruleId?: pulumi.Input<number>;
    /**
     * Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
     */
    severity?: pulumi.Input<string>;
    /**
     * Signature name.
     */
    sigName?: pulumi.Input<string>;
    /**
     * Custom signature enclosed in single quotes.
     */
    signature?: pulumi.Input<string>;
    /**
     * Enable/disable this signature. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Signature tag.
     */
    tag?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
