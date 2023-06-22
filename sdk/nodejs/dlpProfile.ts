// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure DLP profiles. Applies to FortiOS Version `>= 7.2.0`.
 *
 * ## Import
 *
 * Dlp Profile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/dlpProfile:DlpProfile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/dlpProfile:DlpProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class DlpProfile extends pulumi.CustomResource {
    /**
     * Get an existing DlpProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DlpProfileState, opts?: pulumi.CustomResourceOptions): DlpProfile {
        return new DlpProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dlpProfile:DlpProfile';

    /**
     * Returns true if the given object is an instance of DlpProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DlpProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DlpProfile.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable DLP logging. Valid values: `enable`, `disable`.
     */
    public readonly dlpLog!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
     */
    public readonly extendedLog!: pulumi.Output<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    public readonly featureSet!: pulumi.Output<string>;
    /**
     * Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
     */
    public readonly fullArchiveProto!: pulumi.Output<string>;
    /**
     * Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
     */
    public readonly nacQuarLog!: pulumi.Output<string>;
    /**
     * Name of the DLP profile.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Replacement message group used by this DLP profile.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Set up DLP rules for this profile. The structure of `rule` block is documented below.
     */
    public readonly rules!: pulumi.Output<outputs.DlpProfileRule[] | undefined>;
    /**
     * Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
     */
    public readonly summaryProto!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a DlpProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DlpProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DlpProfileArgs | DlpProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DlpProfileState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dlpLog"] = state ? state.dlpLog : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["fullArchiveProto"] = state ? state.fullArchiveProto : undefined;
            resourceInputs["nacQuarLog"] = state ? state.nacQuarLog : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["summaryProto"] = state ? state.summaryProto : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DlpProfileArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dlpLog"] = args ? args.dlpLog : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["fullArchiveProto"] = args ? args.fullArchiveProto : undefined;
            resourceInputs["nacQuarLog"] = args ? args.nacQuarLog : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["summaryProto"] = args ? args.summaryProto : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DlpProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DlpProfile resources.
 */
export interface DlpProfileState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable DLP logging. Valid values: `enable`, `disable`.
     */
    dlpLog?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
     */
    fullArchiveProto?: pulumi.Input<string>;
    /**
     * Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
     */
    nacQuarLog?: pulumi.Input<string>;
    /**
     * Name of the DLP profile.
     */
    name?: pulumi.Input<string>;
    /**
     * Replacement message group used by this DLP profile.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Set up DLP rules for this profile. The structure of `rule` block is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.DlpProfileRule>[]>;
    /**
     * Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
     */
    summaryProto?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DlpProfile resource.
 */
export interface DlpProfileArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable/disable DLP logging. Valid values: `enable`, `disable`.
     */
    dlpLog?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
     */
    fullArchiveProto?: pulumi.Input<string>;
    /**
     * Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
     */
    nacQuarLog?: pulumi.Input<string>;
    /**
     * Name of the DLP profile.
     */
    name?: pulumi.Input<string>;
    /**
     * Replacement message group used by this DLP profile.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Set up DLP rules for this profile. The structure of `rule` block is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.DlpProfileRule>[]>;
    /**
     * Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
     */
    summaryProto?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
