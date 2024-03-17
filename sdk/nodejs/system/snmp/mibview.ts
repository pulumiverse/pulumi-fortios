// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * SNMP Access Control MIB View configuration. Applies to FortiOS Version `>= 7.2.0`.
 *
 * ## Import
 *
 * SystemSnmp MibView can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/snmp/mibview:Mibview labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/snmp/mibview:Mibview labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Mibview extends pulumi.CustomResource {
    /**
     * Get an existing Mibview resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MibviewState, opts?: pulumi.CustomResourceOptions): Mibview {
        return new Mibview(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/snmp/mibview:Mibview';

    /**
     * Returns true if the given object is an instance of Mibview.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Mibview {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Mibview.__pulumiType;
    }

    /**
     * The OID subtrees to be excluded in the view. Maximum 64 allowed.
     */
    public readonly exclude!: pulumi.Output<string>;
    /**
     * The OID subtrees to be included in the view. Maximum 16 allowed.
     */
    public readonly include!: pulumi.Output<string>;
    /**
     * MIB view name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Mibview resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MibviewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MibviewArgs | MibviewState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MibviewState | undefined;
            resourceInputs["exclude"] = state ? state.exclude : undefined;
            resourceInputs["include"] = state ? state.include : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as MibviewArgs | undefined;
            resourceInputs["exclude"] = args ? args.exclude : undefined;
            resourceInputs["include"] = args ? args.include : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Mibview.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Mibview resources.
 */
export interface MibviewState {
    /**
     * The OID subtrees to be excluded in the view. Maximum 64 allowed.
     */
    exclude?: pulumi.Input<string>;
    /**
     * The OID subtrees to be included in the view. Maximum 16 allowed.
     */
    include?: pulumi.Input<string>;
    /**
     * MIB view name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Mibview resource.
 */
export interface MibviewArgs {
    /**
     * The OID subtrees to be excluded in the view. Maximum 64 allowed.
     */
    exclude?: pulumi.Input<string>;
    /**
     * The OID subtrees to be included in the view. Maximum 16 allowed.
     */
    include?: pulumi.Input<string>;
    /**
     * MIB view name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
