// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure quarantine support.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.user.Quarantine("trname", {quarantine: "enable"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * User Quarantine can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/quarantine:Quarantine labelname UserQuarantine
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/quarantine:Quarantine labelname UserQuarantine
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Quarantine extends pulumi.CustomResource {
    /**
     * Get an existing Quarantine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QuarantineState, opts?: pulumi.CustomResourceOptions): Quarantine {
        return new Quarantine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/quarantine:Quarantine';

    /**
     * Returns true if the given object is an instance of Quarantine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Quarantine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Quarantine.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Firewall address group which includes all quarantine MAC address.
     */
    public readonly firewallGroups!: pulumi.Output<string>;
    /**
     * Enable/disable quarantine. Valid values: `enable`, `disable`.
     */
    public readonly quarantine!: pulumi.Output<string>;
    /**
     * Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
     */
    public readonly targets!: pulumi.Output<outputs.user.QuarantineTarget[] | undefined>;
    /**
     * Traffic policy for quarantined MACs.
     */
    public readonly trafficPolicy!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Quarantine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: QuarantineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QuarantineArgs | QuarantineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QuarantineState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["firewallGroups"] = state ? state.firewallGroups : undefined;
            resourceInputs["quarantine"] = state ? state.quarantine : undefined;
            resourceInputs["targets"] = state ? state.targets : undefined;
            resourceInputs["trafficPolicy"] = state ? state.trafficPolicy : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as QuarantineArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["firewallGroups"] = args ? args.firewallGroups : undefined;
            resourceInputs["quarantine"] = args ? args.quarantine : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["trafficPolicy"] = args ? args.trafficPolicy : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Quarantine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Quarantine resources.
 */
export interface QuarantineState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Firewall address group which includes all quarantine MAC address.
     */
    firewallGroups?: pulumi.Input<string>;
    /**
     * Enable/disable quarantine. Valid values: `enable`, `disable`.
     */
    quarantine?: pulumi.Input<string>;
    /**
     * Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.user.QuarantineTarget>[]>;
    /**
     * Traffic policy for quarantined MACs.
     */
    trafficPolicy?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Quarantine resource.
 */
export interface QuarantineArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Firewall address group which includes all quarantine MAC address.
     */
    firewallGroups?: pulumi.Input<string>;
    /**
     * Enable/disable quarantine. Valid values: `enable`, `disable`.
     */
    quarantine?: pulumi.Input<string>;
    /**
     * Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.user.QuarantineTarget>[]>;
    /**
     * Traffic policy for quarantined MACs.
     */
    trafficPolicy?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
