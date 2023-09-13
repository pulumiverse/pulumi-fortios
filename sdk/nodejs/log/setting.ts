// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure general log settings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.log.Setting("trname", {
 *     briefTrafficFormat: "disable",
 *     daemonLog: "disable",
 *     expolicyImplicitLog: "disable",
 *     fazOverride: "disable",
 *     fwpolicy6ImplicitLog: "disable",
 *     fwpolicyImplicitLog: "disable",
 *     localInAllow: "disable",
 *     localInDenyBroadcast: "disable",
 *     localInDenyUnicast: "disable",
 *     localOut: "disable",
 *     logInvalidPacket: "disable",
 *     logPolicyComment: "disable",
 *     logPolicyName: "disable",
 *     logUserInUpper: "disable",
 *     neighborEvent: "disable",
 *     resolveIp: "disable",
 *     resolvePort: "enable",
 *     syslogOverride: "disable",
 *     userAnonymize: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Log Setting can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:log/setting:Setting labelname LogSetting
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:log/setting:Setting labelname LogSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Setting extends pulumi.CustomResource {
    /**
     * Get an existing Setting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SettingState, opts?: pulumi.CustomResourceOptions): Setting {
        return new Setting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:log/setting:Setting';

    /**
     * Returns true if the given object is an instance of Setting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Setting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Setting.__pulumiType;
    }

    /**
     * User name anonymization hash salt.
     */
    public readonly anonymizationHash!: pulumi.Output<string>;
    /**
     * Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly briefTrafficFormat!: pulumi.Output<string>;
    /**
     * Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
     */
    public readonly customLogFields!: pulumi.Output<outputs.log.SettingCustomLogField[] | undefined>;
    /**
     * Enable/disable daemon logging. Valid values: `enable`, `disable`.
     */
    public readonly daemonLog!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
     */
    public readonly expolicyImplicitLog!: pulumi.Output<string>;
    /**
     * Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
     */
    public readonly fazOverride!: pulumi.Output<string>;
    /**
     * Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
     */
    public readonly fwpolicy6ImplicitLog!: pulumi.Output<string>;
    /**
     * Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
     */
    public readonly fwpolicyImplicitLog!: pulumi.Output<string>;
    /**
     * Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
     */
    public readonly localInAllow!: pulumi.Output<string>;
    /**
     * Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
     */
    public readonly localInDenyBroadcast!: pulumi.Output<string>;
    /**
     * Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
     */
    public readonly localInDenyUnicast!: pulumi.Output<string>;
    /**
     * Enable/disable local-out logging. Valid values: `enable`, `disable`.
     */
    public readonly localOut!: pulumi.Output<string>;
    /**
     * Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
     */
    public readonly localOutIocDetection!: pulumi.Output<string>;
    /**
     * Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly logInvalidPacket!: pulumi.Output<string>;
    /**
     * Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
     */
    public readonly logPolicyComment!: pulumi.Output<string>;
    /**
     * Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
     */
    public readonly logPolicyName!: pulumi.Output<string>;
    /**
     * Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
     */
    public readonly logUserInUpper!: pulumi.Output<string>;
    /**
     * Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
     */
    public readonly neighborEvent!: pulumi.Output<string>;
    /**
     * Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
     */
    public readonly resolveIp!: pulumi.Output<string>;
    /**
     * Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
     */
    public readonly resolvePort!: pulumi.Output<string>;
    /**
     * Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
     */
    public readonly restApiGet!: pulumi.Output<string>;
    /**
     * Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
     */
    public readonly restApiSet!: pulumi.Output<string>;
    /**
     * Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
     */
    public readonly syslogOverride!: pulumi.Output<string>;
    /**
     * Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
     */
    public readonly userAnonymize!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Setting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SettingArgs | SettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SettingState | undefined;
            resourceInputs["anonymizationHash"] = state ? state.anonymizationHash : undefined;
            resourceInputs["briefTrafficFormat"] = state ? state.briefTrafficFormat : undefined;
            resourceInputs["customLogFields"] = state ? state.customLogFields : undefined;
            resourceInputs["daemonLog"] = state ? state.daemonLog : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["expolicyImplicitLog"] = state ? state.expolicyImplicitLog : undefined;
            resourceInputs["fazOverride"] = state ? state.fazOverride : undefined;
            resourceInputs["fwpolicy6ImplicitLog"] = state ? state.fwpolicy6ImplicitLog : undefined;
            resourceInputs["fwpolicyImplicitLog"] = state ? state.fwpolicyImplicitLog : undefined;
            resourceInputs["localInAllow"] = state ? state.localInAllow : undefined;
            resourceInputs["localInDenyBroadcast"] = state ? state.localInDenyBroadcast : undefined;
            resourceInputs["localInDenyUnicast"] = state ? state.localInDenyUnicast : undefined;
            resourceInputs["localOut"] = state ? state.localOut : undefined;
            resourceInputs["localOutIocDetection"] = state ? state.localOutIocDetection : undefined;
            resourceInputs["logInvalidPacket"] = state ? state.logInvalidPacket : undefined;
            resourceInputs["logPolicyComment"] = state ? state.logPolicyComment : undefined;
            resourceInputs["logPolicyName"] = state ? state.logPolicyName : undefined;
            resourceInputs["logUserInUpper"] = state ? state.logUserInUpper : undefined;
            resourceInputs["neighborEvent"] = state ? state.neighborEvent : undefined;
            resourceInputs["resolveIp"] = state ? state.resolveIp : undefined;
            resourceInputs["resolvePort"] = state ? state.resolvePort : undefined;
            resourceInputs["restApiGet"] = state ? state.restApiGet : undefined;
            resourceInputs["restApiSet"] = state ? state.restApiSet : undefined;
            resourceInputs["syslogOverride"] = state ? state.syslogOverride : undefined;
            resourceInputs["userAnonymize"] = state ? state.userAnonymize : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SettingArgs | undefined;
            resourceInputs["anonymizationHash"] = args ? args.anonymizationHash : undefined;
            resourceInputs["briefTrafficFormat"] = args ? args.briefTrafficFormat : undefined;
            resourceInputs["customLogFields"] = args ? args.customLogFields : undefined;
            resourceInputs["daemonLog"] = args ? args.daemonLog : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["expolicyImplicitLog"] = args ? args.expolicyImplicitLog : undefined;
            resourceInputs["fazOverride"] = args ? args.fazOverride : undefined;
            resourceInputs["fwpolicy6ImplicitLog"] = args ? args.fwpolicy6ImplicitLog : undefined;
            resourceInputs["fwpolicyImplicitLog"] = args ? args.fwpolicyImplicitLog : undefined;
            resourceInputs["localInAllow"] = args ? args.localInAllow : undefined;
            resourceInputs["localInDenyBroadcast"] = args ? args.localInDenyBroadcast : undefined;
            resourceInputs["localInDenyUnicast"] = args ? args.localInDenyUnicast : undefined;
            resourceInputs["localOut"] = args ? args.localOut : undefined;
            resourceInputs["localOutIocDetection"] = args ? args.localOutIocDetection : undefined;
            resourceInputs["logInvalidPacket"] = args ? args.logInvalidPacket : undefined;
            resourceInputs["logPolicyComment"] = args ? args.logPolicyComment : undefined;
            resourceInputs["logPolicyName"] = args ? args.logPolicyName : undefined;
            resourceInputs["logUserInUpper"] = args ? args.logUserInUpper : undefined;
            resourceInputs["neighborEvent"] = args ? args.neighborEvent : undefined;
            resourceInputs["resolveIp"] = args ? args.resolveIp : undefined;
            resourceInputs["resolvePort"] = args ? args.resolvePort : undefined;
            resourceInputs["restApiGet"] = args ? args.restApiGet : undefined;
            resourceInputs["restApiSet"] = args ? args.restApiSet : undefined;
            resourceInputs["syslogOverride"] = args ? args.syslogOverride : undefined;
            resourceInputs["userAnonymize"] = args ? args.userAnonymize : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Setting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Setting resources.
 */
export interface SettingState {
    /**
     * User name anonymization hash salt.
     */
    anonymizationHash?: pulumi.Input<string>;
    /**
     * Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
     */
    briefTrafficFormat?: pulumi.Input<string>;
    /**
     * Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
     */
    customLogFields?: pulumi.Input<pulumi.Input<inputs.log.SettingCustomLogField>[]>;
    /**
     * Enable/disable daemon logging. Valid values: `enable`, `disable`.
     */
    daemonLog?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
     */
    expolicyImplicitLog?: pulumi.Input<string>;
    /**
     * Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
     */
    fazOverride?: pulumi.Input<string>;
    /**
     * Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
     */
    fwpolicy6ImplicitLog?: pulumi.Input<string>;
    /**
     * Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
     */
    fwpolicyImplicitLog?: pulumi.Input<string>;
    /**
     * Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
     */
    localInAllow?: pulumi.Input<string>;
    /**
     * Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
     */
    localInDenyBroadcast?: pulumi.Input<string>;
    /**
     * Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
     */
    localInDenyUnicast?: pulumi.Input<string>;
    /**
     * Enable/disable local-out logging. Valid values: `enable`, `disable`.
     */
    localOut?: pulumi.Input<string>;
    /**
     * Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
     */
    localOutIocDetection?: pulumi.Input<string>;
    /**
     * Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
     */
    logInvalidPacket?: pulumi.Input<string>;
    /**
     * Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
     */
    logPolicyComment?: pulumi.Input<string>;
    /**
     * Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
     */
    logPolicyName?: pulumi.Input<string>;
    /**
     * Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
     */
    logUserInUpper?: pulumi.Input<string>;
    /**
     * Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
     */
    neighborEvent?: pulumi.Input<string>;
    /**
     * Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
     */
    resolveIp?: pulumi.Input<string>;
    /**
     * Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
     */
    resolvePort?: pulumi.Input<string>;
    /**
     * Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
     */
    restApiGet?: pulumi.Input<string>;
    /**
     * Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
     */
    restApiSet?: pulumi.Input<string>;
    /**
     * Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
     */
    syslogOverride?: pulumi.Input<string>;
    /**
     * Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
     */
    userAnonymize?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Setting resource.
 */
export interface SettingArgs {
    /**
     * User name anonymization hash salt.
     */
    anonymizationHash?: pulumi.Input<string>;
    /**
     * Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
     */
    briefTrafficFormat?: pulumi.Input<string>;
    /**
     * Custom fields to append to all log messages. The structure of `customLogFields` block is documented below.
     */
    customLogFields?: pulumi.Input<pulumi.Input<inputs.log.SettingCustomLogField>[]>;
    /**
     * Enable/disable daemon logging. Valid values: `enable`, `disable`.
     */
    daemonLog?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
     */
    expolicyImplicitLog?: pulumi.Input<string>;
    /**
     * Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
     */
    fazOverride?: pulumi.Input<string>;
    /**
     * Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
     */
    fwpolicy6ImplicitLog?: pulumi.Input<string>;
    /**
     * Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
     */
    fwpolicyImplicitLog?: pulumi.Input<string>;
    /**
     * Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
     */
    localInAllow?: pulumi.Input<string>;
    /**
     * Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
     */
    localInDenyBroadcast?: pulumi.Input<string>;
    /**
     * Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
     */
    localInDenyUnicast?: pulumi.Input<string>;
    /**
     * Enable/disable local-out logging. Valid values: `enable`, `disable`.
     */
    localOut?: pulumi.Input<string>;
    /**
     * Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `enable`, `disable`.
     */
    localOutIocDetection?: pulumi.Input<string>;
    /**
     * Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
     */
    logInvalidPacket?: pulumi.Input<string>;
    /**
     * Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
     */
    logPolicyComment?: pulumi.Input<string>;
    /**
     * Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
     */
    logPolicyName?: pulumi.Input<string>;
    /**
     * Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
     */
    logUserInUpper?: pulumi.Input<string>;
    /**
     * Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
     */
    neighborEvent?: pulumi.Input<string>;
    /**
     * Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
     */
    resolveIp?: pulumi.Input<string>;
    /**
     * Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
     */
    resolvePort?: pulumi.Input<string>;
    /**
     * Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
     */
    restApiGet?: pulumi.Input<string>;
    /**
     * Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
     */
    restApiSet?: pulumi.Input<string>;
    /**
     * Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
     */
    syslogOverride?: pulumi.Input<string>;
    /**
     * Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
     */
    userAnonymize?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
