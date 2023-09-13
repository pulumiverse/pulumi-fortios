// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * 3G MODEM custom. Applies to FortiOS Version `7.0.4`.
 *
 * ## Import
 *
 * System3GModem Custom can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:system3gmodem/custom:Custom labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:system3gmodem/custom:Custom labelname {{fosid}}
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
    public static readonly __pulumiType = 'fortios:system3gmodem/custom:Custom';

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
     * USB interface class in hexadecimal format (00-ff).
     */
    public readonly classId!: pulumi.Output<string>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Init string in hexadecimal format (even length).
     */
    public readonly initString!: pulumi.Output<string>;
    /**
     * MODEM model name.
     */
    public readonly model!: pulumi.Output<string>;
    /**
     * USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
     */
    public readonly modeswitchString!: pulumi.Output<string>;
    /**
     * USB product ID in hexadecimal format (0000-ffff).
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * MODEM vendor name.
     */
    public readonly vendor!: pulumi.Output<string>;
    /**
     * USB vendor ID in hexadecimal format (0000-ffff).
     */
    public readonly vendorId!: pulumi.Output<string>;

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
            resourceInputs["classId"] = state ? state.classId : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["initString"] = state ? state.initString : undefined;
            resourceInputs["model"] = state ? state.model : undefined;
            resourceInputs["modeswitchString"] = state ? state.modeswitchString : undefined;
            resourceInputs["productId"] = state ? state.productId : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
            resourceInputs["vendorId"] = state ? state.vendorId : undefined;
        } else {
            const args = argsOrState as CustomArgs | undefined;
            resourceInputs["classId"] = args ? args.classId : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["initString"] = args ? args.initString : undefined;
            resourceInputs["model"] = args ? args.model : undefined;
            resourceInputs["modeswitchString"] = args ? args.modeswitchString : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
            resourceInputs["vendorId"] = args ? args.vendorId : undefined;
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
     * USB interface class in hexadecimal format (00-ff).
     */
    classId?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Init string in hexadecimal format (even length).
     */
    initString?: pulumi.Input<string>;
    /**
     * MODEM model name.
     */
    model?: pulumi.Input<string>;
    /**
     * USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
     */
    modeswitchString?: pulumi.Input<string>;
    /**
     * USB product ID in hexadecimal format (0000-ffff).
     */
    productId?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * MODEM vendor name.
     */
    vendor?: pulumi.Input<string>;
    /**
     * USB vendor ID in hexadecimal format (0000-ffff).
     */
    vendorId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Custom resource.
 */
export interface CustomArgs {
    /**
     * USB interface class in hexadecimal format (00-ff).
     */
    classId?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Init string in hexadecimal format (even length).
     */
    initString?: pulumi.Input<string>;
    /**
     * MODEM model name.
     */
    model?: pulumi.Input<string>;
    /**
     * USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
     */
    modeswitchString?: pulumi.Input<string>;
    /**
     * USB product ID in hexadecimal format (0000-ffff).
     */
    productId?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * MODEM vendor name.
     */
    vendor?: pulumi.Input<string>;
    /**
     * USB vendor ID in hexadecimal format (0000-ffff).
     */
    vendorId?: pulumi.Input<string>;
}
