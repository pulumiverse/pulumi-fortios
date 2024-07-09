// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Configure Web filter profiles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.filter.web.Profile("trname", {
 *     extendedLog: "disable",
 *     ftgdWf: {
 *         exemptQuota: "17",
 *         filters: [
 *             {
 *                 action: "warning",
 *                 category: 2,
 *                 id: 1,
 *                 log: "enable",
 *                 warnDuration: "5m",
 *                 warningDurationType: "timeout",
 *                 warningPrompt: "per-category",
 *             },
 *             {
 *                 action: "warning",
 *                 category: 7,
 *                 id: 2,
 *                 log: "enable",
 *                 warnDuration: "5m",
 *                 warningDurationType: "timeout",
 *                 warningPrompt: "per-category",
 *             },
 *         ],
 *         maxQuotaTimeout: 300,
 *         rateCrlUrls: "enable",
 *         rateCssUrls: "enable",
 *         rateImageUrls: "enable",
 *         rateJavascriptUrls: "enable",
 *     },
 *     httpsReplacemsg: "enable",
 *     inspectionMode: "flow-based",
 *     logAllUrl: "disable",
 *     override: {
 *         ovrdCookie: "deny",
 *         ovrdDur: "15m",
 *         ovrdDurMode: "constant",
 *         ovrdScope: "user",
 *         profileAttribute: "Login-LAT-Service",
 *         profileType: "list",
 *     },
 *     postAction: "normal",
 *     web: {
 *         blacklist: "disable",
 *         bwordTable: 0,
 *         bwordThreshold: 10,
 *         contentHeaderList: 0,
 *         logSearch: "disable",
 *         urlfilterTable: 0,
 *         youtubeRestrict: "none",
 *     },
 *     webContentLog: "enable",
 *     webExtendedAllActionLog: "disable",
 *     webFilterActivexLog: "enable",
 *     webFilterAppletLog: "enable",
 *     webFilterCommandBlockLog: "enable",
 *     webFilterCookieLog: "enable",
 *     webFilterCookieRemovalLog: "enable",
 *     webFilterJsLog: "enable",
 *     webFilterJscriptLog: "enable",
 *     webFilterRefererLog: "enable",
 *     webFilterUnknownLog: "enable",
 *     webFilterVbsLog: "enable",
 *     webFtgdErrLog: "enable",
 *     webFtgdQuotaUsage: "enable",
 *     webInvalidDomainLog: "enable",
 *     webUrlLog: "enable",
 *     wisp: "disable",
 *     wispAlgorithm: "auto-learning",
 *     youtubeChannelStatus: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Webfilter Profile can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:filter/web/profile:Profile labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:filter/web/profile:Profile labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Profile extends pulumi.CustomResource {
    /**
     * Get an existing Profile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileState, opts?: pulumi.CustomResourceOptions): Profile {
        return new Profile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:filter/web/profile:Profile';

    /**
     * Returns true if the given object is an instance of Profile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Profile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Profile.__pulumiType;
    }

    /**
     * AntiPhishing profile. The structure of `antiphish` block is documented below.
     */
    public readonly antiphish!: pulumi.Output<outputs.filter.web.ProfileAntiphish>;
    /**
     * Optional comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
     */
    public readonly extendedLog!: pulumi.Output<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    public readonly featureSet!: pulumi.Output<string>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    public readonly fileFilter!: pulumi.Output<outputs.filter.web.ProfileFileFilter>;
    /**
     * FortiGuard Web Filter settings. The structure of `ftgdWf` block is documented below.
     */
    public readonly ftgdWf!: pulumi.Output<outputs.filter.web.ProfileFtgdWf>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
     */
    public readonly httpsReplacemsg!: pulumi.Output<string>;
    /**
     * Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
     */
    public readonly inspectionMode!: pulumi.Output<string>;
    /**
     * Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
     */
    public readonly logAllUrl!: pulumi.Output<string>;
    /**
     * Profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Options.
     */
    public readonly options!: pulumi.Output<string>;
    /**
     * Web Filter override settings. The structure of `override` block is documented below.
     */
    public readonly override!: pulumi.Output<outputs.filter.web.ProfileOverride>;
    /**
     * Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
     */
    public readonly ovrdPerm!: pulumi.Output<string>;
    /**
     * Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
     */
    public readonly postAction!: pulumi.Output<string>;
    /**
     * Replacement message group.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Web content filtering settings. The structure of `web` block is documented below.
     */
    public readonly web!: pulumi.Output<outputs.filter.web.ProfileWeb>;
    /**
     * Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
     */
    public readonly webAntiphishingLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
     */
    public readonly webContentLog!: pulumi.Output<string>;
    /**
     * Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
     */
    public readonly webExtendedAllActionLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
     */
    public readonly webFilterActivexLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging Java applets. Valid values: `enable`, `disable`.
     */
    public readonly webFilterAppletLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
     */
    public readonly webFilterCommandBlockLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
     */
    public readonly webFilterCookieLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
     */
    public readonly webFilterCookieRemovalLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
     */
    public readonly webFilterJsLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging JScripts. Valid values: `enable`, `disable`.
     */
    public readonly webFilterJscriptLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging referrers. Valid values: `enable`, `disable`.
     */
    public readonly webFilterRefererLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
     */
    public readonly webFilterUnknownLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
     */
    public readonly webFilterVbsLog!: pulumi.Output<string>;
    /**
     * Log encoding in flow mode. Valid values: `utf-8`, `punycode`.
     */
    public readonly webFlowLogEncoding!: pulumi.Output<string>;
    /**
     * Enable/disable logging rating errors. Valid values: `enable`, `disable`.
     */
    public readonly webFtgdErrLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
     */
    public readonly webFtgdQuotaUsage!: pulumi.Output<string>;
    /**
     * Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
     */
    public readonly webInvalidDomainLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
     */
    public readonly webUrlLog!: pulumi.Output<string>;
    /**
     * Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
     */
    public readonly wisp!: pulumi.Output<string>;
    /**
     * WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
     */
    public readonly wispAlgorithm!: pulumi.Output<string>;
    /**
     * WISP servers. The structure of `wispServers` block is documented below.
     */
    public readonly wispServers!: pulumi.Output<outputs.filter.web.ProfileWispServer[] | undefined>;
    /**
     * YouTube channel filter. The structure of `youtubeChannelFilter` block is documented below.
     */
    public readonly youtubeChannelFilters!: pulumi.Output<outputs.filter.web.ProfileYoutubeChannelFilter[] | undefined>;
    /**
     * YouTube channel filter status.
     */
    public readonly youtubeChannelStatus!: pulumi.Output<string>;

    /**
     * Create a Profile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileArgs | ProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileState | undefined;
            resourceInputs["antiphish"] = state ? state.antiphish : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["fileFilter"] = state ? state.fileFilter : undefined;
            resourceInputs["ftgdWf"] = state ? state.ftgdWf : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["httpsReplacemsg"] = state ? state.httpsReplacemsg : undefined;
            resourceInputs["inspectionMode"] = state ? state.inspectionMode : undefined;
            resourceInputs["logAllUrl"] = state ? state.logAllUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["override"] = state ? state.override : undefined;
            resourceInputs["ovrdPerm"] = state ? state.ovrdPerm : undefined;
            resourceInputs["postAction"] = state ? state.postAction : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["web"] = state ? state.web : undefined;
            resourceInputs["webAntiphishingLog"] = state ? state.webAntiphishingLog : undefined;
            resourceInputs["webContentLog"] = state ? state.webContentLog : undefined;
            resourceInputs["webExtendedAllActionLog"] = state ? state.webExtendedAllActionLog : undefined;
            resourceInputs["webFilterActivexLog"] = state ? state.webFilterActivexLog : undefined;
            resourceInputs["webFilterAppletLog"] = state ? state.webFilterAppletLog : undefined;
            resourceInputs["webFilterCommandBlockLog"] = state ? state.webFilterCommandBlockLog : undefined;
            resourceInputs["webFilterCookieLog"] = state ? state.webFilterCookieLog : undefined;
            resourceInputs["webFilterCookieRemovalLog"] = state ? state.webFilterCookieRemovalLog : undefined;
            resourceInputs["webFilterJsLog"] = state ? state.webFilterJsLog : undefined;
            resourceInputs["webFilterJscriptLog"] = state ? state.webFilterJscriptLog : undefined;
            resourceInputs["webFilterRefererLog"] = state ? state.webFilterRefererLog : undefined;
            resourceInputs["webFilterUnknownLog"] = state ? state.webFilterUnknownLog : undefined;
            resourceInputs["webFilterVbsLog"] = state ? state.webFilterVbsLog : undefined;
            resourceInputs["webFlowLogEncoding"] = state ? state.webFlowLogEncoding : undefined;
            resourceInputs["webFtgdErrLog"] = state ? state.webFtgdErrLog : undefined;
            resourceInputs["webFtgdQuotaUsage"] = state ? state.webFtgdQuotaUsage : undefined;
            resourceInputs["webInvalidDomainLog"] = state ? state.webInvalidDomainLog : undefined;
            resourceInputs["webUrlLog"] = state ? state.webUrlLog : undefined;
            resourceInputs["wisp"] = state ? state.wisp : undefined;
            resourceInputs["wispAlgorithm"] = state ? state.wispAlgorithm : undefined;
            resourceInputs["wispServers"] = state ? state.wispServers : undefined;
            resourceInputs["youtubeChannelFilters"] = state ? state.youtubeChannelFilters : undefined;
            resourceInputs["youtubeChannelStatus"] = state ? state.youtubeChannelStatus : undefined;
        } else {
            const args = argsOrState as ProfileArgs | undefined;
            resourceInputs["antiphish"] = args ? args.antiphish : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["fileFilter"] = args ? args.fileFilter : undefined;
            resourceInputs["ftgdWf"] = args ? args.ftgdWf : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["httpsReplacemsg"] = args ? args.httpsReplacemsg : undefined;
            resourceInputs["inspectionMode"] = args ? args.inspectionMode : undefined;
            resourceInputs["logAllUrl"] = args ? args.logAllUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["override"] = args ? args.override : undefined;
            resourceInputs["ovrdPerm"] = args ? args.ovrdPerm : undefined;
            resourceInputs["postAction"] = args ? args.postAction : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["web"] = args ? args.web : undefined;
            resourceInputs["webAntiphishingLog"] = args ? args.webAntiphishingLog : undefined;
            resourceInputs["webContentLog"] = args ? args.webContentLog : undefined;
            resourceInputs["webExtendedAllActionLog"] = args ? args.webExtendedAllActionLog : undefined;
            resourceInputs["webFilterActivexLog"] = args ? args.webFilterActivexLog : undefined;
            resourceInputs["webFilterAppletLog"] = args ? args.webFilterAppletLog : undefined;
            resourceInputs["webFilterCommandBlockLog"] = args ? args.webFilterCommandBlockLog : undefined;
            resourceInputs["webFilterCookieLog"] = args ? args.webFilterCookieLog : undefined;
            resourceInputs["webFilterCookieRemovalLog"] = args ? args.webFilterCookieRemovalLog : undefined;
            resourceInputs["webFilterJsLog"] = args ? args.webFilterJsLog : undefined;
            resourceInputs["webFilterJscriptLog"] = args ? args.webFilterJscriptLog : undefined;
            resourceInputs["webFilterRefererLog"] = args ? args.webFilterRefererLog : undefined;
            resourceInputs["webFilterUnknownLog"] = args ? args.webFilterUnknownLog : undefined;
            resourceInputs["webFilterVbsLog"] = args ? args.webFilterVbsLog : undefined;
            resourceInputs["webFlowLogEncoding"] = args ? args.webFlowLogEncoding : undefined;
            resourceInputs["webFtgdErrLog"] = args ? args.webFtgdErrLog : undefined;
            resourceInputs["webFtgdQuotaUsage"] = args ? args.webFtgdQuotaUsage : undefined;
            resourceInputs["webInvalidDomainLog"] = args ? args.webInvalidDomainLog : undefined;
            resourceInputs["webUrlLog"] = args ? args.webUrlLog : undefined;
            resourceInputs["wisp"] = args ? args.wisp : undefined;
            resourceInputs["wispAlgorithm"] = args ? args.wispAlgorithm : undefined;
            resourceInputs["wispServers"] = args ? args.wispServers : undefined;
            resourceInputs["youtubeChannelFilters"] = args ? args.youtubeChannelFilters : undefined;
            resourceInputs["youtubeChannelStatus"] = args ? args.youtubeChannelStatus : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Profile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Profile resources.
 */
export interface ProfileState {
    /**
     * AntiPhishing profile. The structure of `antiphish` block is documented below.
     */
    antiphish?: pulumi.Input<inputs.filter.web.ProfileAntiphish>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    fileFilter?: pulumi.Input<inputs.filter.web.ProfileFileFilter>;
    /**
     * FortiGuard Web Filter settings. The structure of `ftgdWf` block is documented below.
     */
    ftgdWf?: pulumi.Input<inputs.filter.web.ProfileFtgdWf>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
     */
    httpsReplacemsg?: pulumi.Input<string>;
    /**
     * Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
     */
    inspectionMode?: pulumi.Input<string>;
    /**
     * Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
     */
    logAllUrl?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Options.
     */
    options?: pulumi.Input<string>;
    /**
     * Web Filter override settings. The structure of `override` block is documented below.
     */
    override?: pulumi.Input<inputs.filter.web.ProfileOverride>;
    /**
     * Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
     */
    ovrdPerm?: pulumi.Input<string>;
    /**
     * Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
     */
    postAction?: pulumi.Input<string>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Web content filtering settings. The structure of `web` block is documented below.
     */
    web?: pulumi.Input<inputs.filter.web.ProfileWeb>;
    /**
     * Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
     */
    webAntiphishingLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
     */
    webContentLog?: pulumi.Input<string>;
    /**
     * Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
     */
    webExtendedAllActionLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
     */
    webFilterActivexLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging Java applets. Valid values: `enable`, `disable`.
     */
    webFilterAppletLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
     */
    webFilterCommandBlockLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
     */
    webFilterCookieLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
     */
    webFilterCookieRemovalLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
     */
    webFilterJsLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging JScripts. Valid values: `enable`, `disable`.
     */
    webFilterJscriptLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging referrers. Valid values: `enable`, `disable`.
     */
    webFilterRefererLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
     */
    webFilterUnknownLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
     */
    webFilterVbsLog?: pulumi.Input<string>;
    /**
     * Log encoding in flow mode. Valid values: `utf-8`, `punycode`.
     */
    webFlowLogEncoding?: pulumi.Input<string>;
    /**
     * Enable/disable logging rating errors. Valid values: `enable`, `disable`.
     */
    webFtgdErrLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
     */
    webFtgdQuotaUsage?: pulumi.Input<string>;
    /**
     * Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
     */
    webInvalidDomainLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
     */
    webUrlLog?: pulumi.Input<string>;
    /**
     * Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
     */
    wisp?: pulumi.Input<string>;
    /**
     * WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
     */
    wispAlgorithm?: pulumi.Input<string>;
    /**
     * WISP servers. The structure of `wispServers` block is documented below.
     */
    wispServers?: pulumi.Input<pulumi.Input<inputs.filter.web.ProfileWispServer>[]>;
    /**
     * YouTube channel filter. The structure of `youtubeChannelFilter` block is documented below.
     */
    youtubeChannelFilters?: pulumi.Input<pulumi.Input<inputs.filter.web.ProfileYoutubeChannelFilter>[]>;
    /**
     * YouTube channel filter status.
     */
    youtubeChannelStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Profile resource.
 */
export interface ProfileArgs {
    /**
     * AntiPhishing profile. The structure of `antiphish` block is documented below.
     */
    antiphish?: pulumi.Input<inputs.filter.web.ProfileAntiphish>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * File filter. The structure of `fileFilter` block is documented below.
     */
    fileFilter?: pulumi.Input<inputs.filter.web.ProfileFileFilter>;
    /**
     * FortiGuard Web Filter settings. The structure of `ftgdWf` block is documented below.
     */
    ftgdWf?: pulumi.Input<inputs.filter.web.ProfileFtgdWf>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
     */
    httpsReplacemsg?: pulumi.Input<string>;
    /**
     * Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
     */
    inspectionMode?: pulumi.Input<string>;
    /**
     * Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
     */
    logAllUrl?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Options.
     */
    options?: pulumi.Input<string>;
    /**
     * Web Filter override settings. The structure of `override` block is documented below.
     */
    override?: pulumi.Input<inputs.filter.web.ProfileOverride>;
    /**
     * Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
     */
    ovrdPerm?: pulumi.Input<string>;
    /**
     * Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
     */
    postAction?: pulumi.Input<string>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Web content filtering settings. The structure of `web` block is documented below.
     */
    web?: pulumi.Input<inputs.filter.web.ProfileWeb>;
    /**
     * Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
     */
    webAntiphishingLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
     */
    webContentLog?: pulumi.Input<string>;
    /**
     * Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
     */
    webExtendedAllActionLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
     */
    webFilterActivexLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging Java applets. Valid values: `enable`, `disable`.
     */
    webFilterAppletLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
     */
    webFilterCommandBlockLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
     */
    webFilterCookieLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
     */
    webFilterCookieRemovalLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
     */
    webFilterJsLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging JScripts. Valid values: `enable`, `disable`.
     */
    webFilterJscriptLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging referrers. Valid values: `enable`, `disable`.
     */
    webFilterRefererLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
     */
    webFilterUnknownLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
     */
    webFilterVbsLog?: pulumi.Input<string>;
    /**
     * Log encoding in flow mode. Valid values: `utf-8`, `punycode`.
     */
    webFlowLogEncoding?: pulumi.Input<string>;
    /**
     * Enable/disable logging rating errors. Valid values: `enable`, `disable`.
     */
    webFtgdErrLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
     */
    webFtgdQuotaUsage?: pulumi.Input<string>;
    /**
     * Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
     */
    webInvalidDomainLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
     */
    webUrlLog?: pulumi.Input<string>;
    /**
     * Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
     */
    wisp?: pulumi.Input<string>;
    /**
     * WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
     */
    wispAlgorithm?: pulumi.Input<string>;
    /**
     * WISP servers. The structure of `wispServers` block is documented below.
     */
    wispServers?: pulumi.Input<pulumi.Input<inputs.filter.web.ProfileWispServer>[]>;
    /**
     * YouTube channel filter. The structure of `youtubeChannelFilter` block is documented below.
     */
    youtubeChannelFilters?: pulumi.Input<pulumi.Input<inputs.filter.web.ProfileYoutubeChannelFilter>[]>;
    /**
     * YouTube channel filter status.
     */
    youtubeChannelStatus?: pulumi.Input<string>;
}
