// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure virtual wire pairs.
 *
 * ## Import
 *
 * System VirtualWirePair can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/virtualwirepair:Virtualwirepair labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/virtualwirepair:Virtualwirepair labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Virtualwirepair extends pulumi.CustomResource {
    /**
     * Get an existing Virtualwirepair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualwirepairState, opts?: pulumi.CustomResourceOptions): Virtualwirepair {
        return new Virtualwirepair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/virtualwirepair:Virtualwirepair';

    /**
     * Returns true if the given object is an instance of Virtualwirepair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Virtualwirepair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Virtualwirepair.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.system.VirtualwirepairMember[]>;
    /**
     * Virtual-wire-pair name. Must be a unique interface name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Set VLAN filters.
     */
    public readonly vlanFilter!: pulumi.Output<string>;
    /**
     * Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
     */
    public readonly wildcardVlan!: pulumi.Output<string>;

    /**
     * Create a Virtualwirepair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualwirepairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualwirepairArgs | VirtualwirepairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualwirepairState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlanFilter"] = state ? state.vlanFilter : undefined;
            resourceInputs["wildcardVlan"] = state ? state.wildcardVlan : undefined;
        } else {
            const args = argsOrState as VirtualwirepairArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlanFilter"] = args ? args.vlanFilter : undefined;
            resourceInputs["wildcardVlan"] = args ? args.wildcardVlan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Virtualwirepair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Virtualwirepair resources.
 */
export interface VirtualwirepairState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.system.VirtualwirepairMember>[]>;
    /**
     * Virtual-wire-pair name. Must be a unique interface name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Set VLAN filters.
     */
    vlanFilter?: pulumi.Input<string>;
    /**
     * Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
     */
    wildcardVlan?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Virtualwirepair resource.
 */
export interface VirtualwirepairArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
     */
    members: pulumi.Input<pulumi.Input<inputs.system.VirtualwirepairMember>[]>;
    /**
     * Virtual-wire-pair name. Must be a unique interface name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Set VLAN filters.
     */
    vlanFilter?: pulumi.Input<string>;
    /**
     * Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
     */
    wildcardVlan?: pulumi.Input<string>;
}
