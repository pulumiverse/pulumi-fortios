// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Replacement messages.
 *
 * ## Import
 *
 * SystemReplacemsg Utm can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/replacemsg/utm:Utm labelname {{msg_type}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/replacemsg/utm:Utm labelname {{msg_type}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Utm extends pulumi.CustomResource {
    /**
     * Get an existing Utm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UtmState, opts?: pulumi.CustomResourceOptions): Utm {
        return new Utm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/replacemsg/utm:Utm';

    /**
     * Returns true if the given object is an instance of Utm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Utm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Utm.__pulumiType;
    }

    /**
     * Message string.
     */
    public readonly buffer!: pulumi.Output<string | undefined>;
    /**
     * Format flag.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Header flag. Valid values: `none`, `http`, `8bit`.
     */
    public readonly header!: pulumi.Output<string>;
    /**
     * Message type.
     */
    public readonly msgType!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Utm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UtmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UtmArgs | UtmState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UtmState | undefined;
            resourceInputs["buffer"] = state ? state.buffer : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["header"] = state ? state.header : undefined;
            resourceInputs["msgType"] = state ? state.msgType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UtmArgs | undefined;
            if ((!args || args.msgType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'msgType'");
            }
            resourceInputs["buffer"] = args ? args.buffer : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["header"] = args ? args.header : undefined;
            resourceInputs["msgType"] = args ? args.msgType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Utm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Utm resources.
 */
export interface UtmState {
    /**
     * Message string.
     */
    buffer?: pulumi.Input<string>;
    /**
     * Format flag.
     */
    format?: pulumi.Input<string>;
    /**
     * Header flag. Valid values: `none`, `http`, `8bit`.
     */
    header?: pulumi.Input<string>;
    /**
     * Message type.
     */
    msgType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Utm resource.
 */
export interface UtmArgs {
    /**
     * Message string.
     */
    buffer?: pulumi.Input<string>;
    /**
     * Format flag.
     */
    format?: pulumi.Input<string>;
    /**
     * Header flag. Valid values: `none`, `http`, `8bit`.
     */
    header?: pulumi.Input<string>;
    /**
     * Message type.
     */
    msgType: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
