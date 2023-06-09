// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Coordinate federated upgrades within the Security Fabric. Applies to FortiOS Version `>= 7.0.0`.
 *
 * ## Import
 *
 * System FederatedUpgrade can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemFederatedupgrade:SystemFederatedupgrade labelname SystemFederatedUpgrade
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemFederatedupgrade:SystemFederatedupgrade labelname SystemFederatedUpgrade
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemFederatedupgrade extends pulumi.CustomResource {
    /**
     * Get an existing SystemFederatedupgrade resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemFederatedupgradeState, opts?: pulumi.CustomResourceOptions): SystemFederatedupgrade {
        return new SystemFederatedupgrade(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemFederatedupgrade:SystemFederatedupgrade';

    /**
     * Returns true if the given object is an instance of SystemFederatedupgrade.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemFederatedupgrade {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemFederatedupgrade.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Serial number of the node to include.
     */
    public readonly failureDevice!: pulumi.Output<string>;
    /**
     * Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
     */
    public readonly failureReason!: pulumi.Output<string>;
    /**
     * The index of the next image to upgrade to.
     */
    public readonly nextPathIndex!: pulumi.Output<number>;
    /**
     * Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
     */
    public readonly nodeLists!: pulumi.Output<outputs.SystemFederatedupgradeNodeList[] | undefined>;
    /**
     * Current status of the upgrade.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Unique identifier for this upgrade.
     */
    public readonly upgradeId!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemFederatedupgrade resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemFederatedupgradeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemFederatedupgradeArgs | SystemFederatedupgradeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemFederatedupgradeState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["failureDevice"] = state ? state.failureDevice : undefined;
            resourceInputs["failureReason"] = state ? state.failureReason : undefined;
            resourceInputs["nextPathIndex"] = state ? state.nextPathIndex : undefined;
            resourceInputs["nodeLists"] = state ? state.nodeLists : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["upgradeId"] = state ? state.upgradeId : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemFederatedupgradeArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["failureDevice"] = args ? args.failureDevice : undefined;
            resourceInputs["failureReason"] = args ? args.failureReason : undefined;
            resourceInputs["nextPathIndex"] = args ? args.nextPathIndex : undefined;
            resourceInputs["nodeLists"] = args ? args.nodeLists : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["upgradeId"] = args ? args.upgradeId : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemFederatedupgrade.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemFederatedupgrade resources.
 */
export interface SystemFederatedupgradeState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Serial number of the node to include.
     */
    failureDevice?: pulumi.Input<string>;
    /**
     * Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
     */
    failureReason?: pulumi.Input<string>;
    /**
     * The index of the next image to upgrade to.
     */
    nextPathIndex?: pulumi.Input<number>;
    /**
     * Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
     */
    nodeLists?: pulumi.Input<pulumi.Input<inputs.SystemFederatedupgradeNodeList>[]>;
    /**
     * Current status of the upgrade.
     */
    status?: pulumi.Input<string>;
    /**
     * Unique identifier for this upgrade.
     */
    upgradeId?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemFederatedupgrade resource.
 */
export interface SystemFederatedupgradeArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Serial number of the node to include.
     */
    failureDevice?: pulumi.Input<string>;
    /**
     * Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
     */
    failureReason?: pulumi.Input<string>;
    /**
     * The index of the next image to upgrade to.
     */
    nextPathIndex?: pulumi.Input<number>;
    /**
     * Nodes which will be included in the upgrade. The structure of `nodeList` block is documented below.
     */
    nodeLists?: pulumi.Input<pulumi.Input<inputs.SystemFederatedupgradeNodeList>[]>;
    /**
     * Current status of the upgrade.
     */
    status?: pulumi.Input<string>;
    /**
     * Unique identifier for this upgrade.
     */
    upgradeId?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
