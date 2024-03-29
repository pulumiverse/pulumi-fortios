// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure IPS rules.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * // import first and then modify
 * const trname = new fortios.ips.Rule("trname", {
 *     action: "block",
 *     application: "All",
 *     date: 1462435200,
 *     group: "backdoor",
 *     location: "server",
 *     log: "enable",
 *     logPacket: "disable",
 *     os: "All",
 *     rev: 6637,
 *     ruleId: 40473,
 *     service: "UDP, DNS",
 *     severity: "critical",
 *     status: "enable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Ips Rule can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:ips/rule:Rule labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:ips/rule:Rule labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:ips/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * Action. Valid values: `pass`, `block`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Vulnerable applications.
     */
    public readonly application!: pulumi.Output<string>;
    /**
     * Date.
     */
    public readonly date!: pulumi.Output<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Group.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * Vulnerable location.
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
     * Meta data. The structure of `metadata` block is documented below.
     */
    public readonly metadatas!: pulumi.Output<outputs.ips.RuleMetadata[] | undefined>;
    /**
     * Rule name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Vulnerable operation systems.
     */
    public readonly os!: pulumi.Output<string>;
    /**
     * Revision.
     */
    public readonly rev!: pulumi.Output<number>;
    /**
     * Rule ID.
     */
    public readonly ruleId!: pulumi.Output<number>;
    /**
     * Vulnerable service.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Severity.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Enable/disable status. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["application"] = state ? state.application : undefined;
            resourceInputs["date"] = state ? state.date : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["log"] = state ? state.log : undefined;
            resourceInputs["logPacket"] = state ? state.logPacket : undefined;
            resourceInputs["metadatas"] = state ? state.metadatas : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["os"] = state ? state.os : undefined;
            resourceInputs["rev"] = state ? state.rev : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["application"] = args ? args.application : undefined;
            resourceInputs["date"] = args ? args.date : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["log"] = args ? args.log : undefined;
            resourceInputs["logPacket"] = args ? args.logPacket : undefined;
            resourceInputs["metadatas"] = args ? args.metadatas : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["os"] = args ? args.os : undefined;
            resourceInputs["rev"] = args ? args.rev : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * Action. Valid values: `pass`, `block`.
     */
    action?: pulumi.Input<string>;
    /**
     * Vulnerable applications.
     */
    application?: pulumi.Input<string>;
    /**
     * Date.
     */
    date?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Group.
     */
    group?: pulumi.Input<string>;
    /**
     * Vulnerable location.
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
     * Meta data. The structure of `metadata` block is documented below.
     */
    metadatas?: pulumi.Input<pulumi.Input<inputs.ips.RuleMetadata>[]>;
    /**
     * Rule name.
     */
    name?: pulumi.Input<string>;
    /**
     * Vulnerable operation systems.
     */
    os?: pulumi.Input<string>;
    /**
     * Revision.
     */
    rev?: pulumi.Input<number>;
    /**
     * Rule ID.
     */
    ruleId?: pulumi.Input<number>;
    /**
     * Vulnerable service.
     */
    service?: pulumi.Input<string>;
    /**
     * Severity.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable status. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * Action. Valid values: `pass`, `block`.
     */
    action?: pulumi.Input<string>;
    /**
     * Vulnerable applications.
     */
    application?: pulumi.Input<string>;
    /**
     * Date.
     */
    date?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Group.
     */
    group?: pulumi.Input<string>;
    /**
     * Vulnerable location.
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
     * Meta data. The structure of `metadata` block is documented below.
     */
    metadatas?: pulumi.Input<pulumi.Input<inputs.ips.RuleMetadata>[]>;
    /**
     * Rule name.
     */
    name?: pulumi.Input<string>;
    /**
     * Vulnerable operation systems.
     */
    os?: pulumi.Input<string>;
    /**
     * Revision.
     */
    rev?: pulumi.Input<number>;
    /**
     * Rule ID.
     */
    ruleId?: pulumi.Input<number>;
    /**
     * Vulnerable service.
     */
    service?: pulumi.Input<string>;
    /**
     * Severity.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable status. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
