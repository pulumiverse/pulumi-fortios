// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.sys.Zone("trname", {intrazone: "allow"});
 * ```
 *
 * ## Import
 *
 * System Zone can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:sys/zone:Zone labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:sys/zone:Zone labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneState, opts?: pulumi.CustomResourceOptions): Zone {
        return new Zone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/zone:Zone';

    /**
     * Returns true if the given object is an instance of Zone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Zone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Zone.__pulumiType;
    }

    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.sys.ZoneInterface[] | undefined>;
    /**
     * Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
     */
    public readonly intrazone!: pulumi.Output<string>;
    /**
     * Zone name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    public readonly taggings!: pulumi.Output<outputs.sys.ZoneTagging[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneArgs | ZoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["intrazone"] = state ? state.intrazone : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["taggings"] = state ? state.taggings : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ZoneArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["intrazone"] = args ? args.intrazone : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["taggings"] = args ? args.taggings : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Zone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Zone resources.
 */
export interface ZoneState {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.sys.ZoneInterface>[]>;
    /**
     * Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
     */
    intrazone?: pulumi.Input<string>;
    /**
     * Zone name.
     */
    name?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.sys.ZoneTagging>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.sys.ZoneInterface>[]>;
    /**
     * Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.
     */
    intrazone?: pulumi.Input<string>;
    /**
     * Zone name.
     */
    name?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.sys.ZoneTagging>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
