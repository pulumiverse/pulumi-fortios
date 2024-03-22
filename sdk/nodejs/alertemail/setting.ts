// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure alert email settings.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.alertemail.Setting("trname", {
 *     adminLoginLogs: "disable",
 *     alertInterval: 2,
 *     amcInterfaceBypassMode: "disable",
 *     antivirusLogs: "disable",
 *     configurationChangesLogs: "disable",
 *     criticalInterval: 3,
 *     debugInterval: 60,
 *     emailInterval: 5,
 *     emergencyInterval: 1,
 *     errorInterval: 5,
 *     fdsLicenseExpiringDays: 15,
 *     informationInterval: 30,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Alertemail Setting can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:alertemail/setting:Setting labelname AlertemailSetting
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:alertemail/setting:Setting labelname AlertemailSetting
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
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
    public static readonly __pulumiType = 'fortios:alertemail/setting:Setting';

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
     * Enable/disable administrator login/logout logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly adminLoginLogs!: pulumi.Output<string>;
    /**
     * Alert alert interval in minutes.
     */
    public readonly alertInterval!: pulumi.Output<number>;
    /**
     * Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly amcInterfaceBypassMode!: pulumi.Output<string>;
    /**
     * Enable/disable antivirus logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly antivirusLogs!: pulumi.Output<string>;
    /**
     * Enable/disable configuration change logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly configurationChangesLogs!: pulumi.Output<string>;
    /**
     * Critical alert interval in minutes.
     */
    public readonly criticalInterval!: pulumi.Output<number>;
    /**
     * Debug alert interval in minutes.
     */
    public readonly debugInterval!: pulumi.Output<number>;
    /**
     * Interval between sending alert emails (1 - 99999 min, default = 5).
     */
    public readonly emailInterval!: pulumi.Output<number>;
    /**
     * Emergency alert interval in minutes.
     */
    public readonly emergencyInterval!: pulumi.Output<number>;
    /**
     * Error alert interval in minutes.
     */
    public readonly errorInterval!: pulumi.Output<number>;
    /**
     * Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days). On FortiOS versions 6.2.0-7.2.0: default = 100. On FortiOS versions 7.2.1-7.2.6: default = 15.
     */
    public readonly fdsLicenseExpiringDays!: pulumi.Output<number>;
    /**
     * Enable/disable FortiGuard license expiration warnings in alert email. Valid values: `enable`, `disable`.
     */
    public readonly fdsLicenseExpiringWarning!: pulumi.Output<string>;
    /**
     * Enable/disable FortiGuard update logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly fdsUpdateLogs!: pulumi.Output<string>;
    /**
     * How to filter log messages that are sent to alert emails. Valid values: `category`, `threshold`.
     */
    public readonly filterMode!: pulumi.Output<string>;
    /**
     * Enable/disable FIPS and Common Criteria error logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly fipsCcErrors!: pulumi.Output<string>;
    /**
     * Enable/disable firewall authentication failure logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly firewallAuthenticationFailureLogs!: pulumi.Output<string>;
    /**
     * Enable/disable FortiCloud log quota warnings in alert email. Valid values: `enable`, `disable`.
     */
    public readonly fortiguardLogQuotaWarning!: pulumi.Output<string>;
    /**
     * Enable/disable logging of FSSO collector agent disconnect. Valid values: `enable`, `disable`.
     */
    public readonly fssoDisconnectLogs!: pulumi.Output<string>;
    /**
     * Enable/disable HA logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly haLogs!: pulumi.Output<string>;
    /**
     * Information alert interval in minutes.
     */
    public readonly informationInterval!: pulumi.Output<number>;
    /**
     * Enable/disable IPS logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly ipsLogs!: pulumi.Output<string>;
    /**
     * Enable/disable IPsec error logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly ipsecErrorsLogs!: pulumi.Output<string>;
    /**
     * Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).
     */
    public readonly localDiskUsage!: pulumi.Output<number>;
    /**
     * Enable/disable disk usage warnings in alert email. Valid values: `enable`, `disable`.
     */
    public readonly logDiskUsageWarning!: pulumi.Output<string>;
    /**
     * Email address to send alert email to (usually a system administrator). On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    public readonly mailto1!: pulumi.Output<string>;
    /**
     * Optional second email address to send alert email to. On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    public readonly mailto2!: pulumi.Output<string>;
    /**
     * Optional third email address to send alert email to. On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    public readonly mailto3!: pulumi.Output<string>;
    /**
     * Notification alert interval in minutes.
     */
    public readonly notificationInterval!: pulumi.Output<number>;
    /**
     * Enable/disable PPP error logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly pppErrorsLogs!: pulumi.Output<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Enable/disable SSH logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly sshLogs!: pulumi.Output<string>;
    /**
     * Enable/disable SSL-VPN authentication error logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly sslvpnAuthenticationErrorsLogs!: pulumi.Output<string>;
    /**
     * Name that appears in the From: field of alert emails. On FortiOS versions 6.2.0-6.4.0: max. 36 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable violation traffic logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly violationTrafficLogs!: pulumi.Output<string>;
    /**
     * Warning alert interval in minutes.
     */
    public readonly warningInterval!: pulumi.Output<number>;
    /**
     * Enable/disable web filter logs in alert email. Valid values: `enable`, `disable`.
     */
    public readonly webfilterLogs!: pulumi.Output<string>;

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
            resourceInputs["adminLoginLogs"] = state ? state.adminLoginLogs : undefined;
            resourceInputs["alertInterval"] = state ? state.alertInterval : undefined;
            resourceInputs["amcInterfaceBypassMode"] = state ? state.amcInterfaceBypassMode : undefined;
            resourceInputs["antivirusLogs"] = state ? state.antivirusLogs : undefined;
            resourceInputs["configurationChangesLogs"] = state ? state.configurationChangesLogs : undefined;
            resourceInputs["criticalInterval"] = state ? state.criticalInterval : undefined;
            resourceInputs["debugInterval"] = state ? state.debugInterval : undefined;
            resourceInputs["emailInterval"] = state ? state.emailInterval : undefined;
            resourceInputs["emergencyInterval"] = state ? state.emergencyInterval : undefined;
            resourceInputs["errorInterval"] = state ? state.errorInterval : undefined;
            resourceInputs["fdsLicenseExpiringDays"] = state ? state.fdsLicenseExpiringDays : undefined;
            resourceInputs["fdsLicenseExpiringWarning"] = state ? state.fdsLicenseExpiringWarning : undefined;
            resourceInputs["fdsUpdateLogs"] = state ? state.fdsUpdateLogs : undefined;
            resourceInputs["filterMode"] = state ? state.filterMode : undefined;
            resourceInputs["fipsCcErrors"] = state ? state.fipsCcErrors : undefined;
            resourceInputs["firewallAuthenticationFailureLogs"] = state ? state.firewallAuthenticationFailureLogs : undefined;
            resourceInputs["fortiguardLogQuotaWarning"] = state ? state.fortiguardLogQuotaWarning : undefined;
            resourceInputs["fssoDisconnectLogs"] = state ? state.fssoDisconnectLogs : undefined;
            resourceInputs["haLogs"] = state ? state.haLogs : undefined;
            resourceInputs["informationInterval"] = state ? state.informationInterval : undefined;
            resourceInputs["ipsLogs"] = state ? state.ipsLogs : undefined;
            resourceInputs["ipsecErrorsLogs"] = state ? state.ipsecErrorsLogs : undefined;
            resourceInputs["localDiskUsage"] = state ? state.localDiskUsage : undefined;
            resourceInputs["logDiskUsageWarning"] = state ? state.logDiskUsageWarning : undefined;
            resourceInputs["mailto1"] = state ? state.mailto1 : undefined;
            resourceInputs["mailto2"] = state ? state.mailto2 : undefined;
            resourceInputs["mailto3"] = state ? state.mailto3 : undefined;
            resourceInputs["notificationInterval"] = state ? state.notificationInterval : undefined;
            resourceInputs["pppErrorsLogs"] = state ? state.pppErrorsLogs : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["sshLogs"] = state ? state.sshLogs : undefined;
            resourceInputs["sslvpnAuthenticationErrorsLogs"] = state ? state.sslvpnAuthenticationErrorsLogs : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["violationTrafficLogs"] = state ? state.violationTrafficLogs : undefined;
            resourceInputs["warningInterval"] = state ? state.warningInterval : undefined;
            resourceInputs["webfilterLogs"] = state ? state.webfilterLogs : undefined;
        } else {
            const args = argsOrState as SettingArgs | undefined;
            resourceInputs["adminLoginLogs"] = args ? args.adminLoginLogs : undefined;
            resourceInputs["alertInterval"] = args ? args.alertInterval : undefined;
            resourceInputs["amcInterfaceBypassMode"] = args ? args.amcInterfaceBypassMode : undefined;
            resourceInputs["antivirusLogs"] = args ? args.antivirusLogs : undefined;
            resourceInputs["configurationChangesLogs"] = args ? args.configurationChangesLogs : undefined;
            resourceInputs["criticalInterval"] = args ? args.criticalInterval : undefined;
            resourceInputs["debugInterval"] = args ? args.debugInterval : undefined;
            resourceInputs["emailInterval"] = args ? args.emailInterval : undefined;
            resourceInputs["emergencyInterval"] = args ? args.emergencyInterval : undefined;
            resourceInputs["errorInterval"] = args ? args.errorInterval : undefined;
            resourceInputs["fdsLicenseExpiringDays"] = args ? args.fdsLicenseExpiringDays : undefined;
            resourceInputs["fdsLicenseExpiringWarning"] = args ? args.fdsLicenseExpiringWarning : undefined;
            resourceInputs["fdsUpdateLogs"] = args ? args.fdsUpdateLogs : undefined;
            resourceInputs["filterMode"] = args ? args.filterMode : undefined;
            resourceInputs["fipsCcErrors"] = args ? args.fipsCcErrors : undefined;
            resourceInputs["firewallAuthenticationFailureLogs"] = args ? args.firewallAuthenticationFailureLogs : undefined;
            resourceInputs["fortiguardLogQuotaWarning"] = args ? args.fortiguardLogQuotaWarning : undefined;
            resourceInputs["fssoDisconnectLogs"] = args ? args.fssoDisconnectLogs : undefined;
            resourceInputs["haLogs"] = args ? args.haLogs : undefined;
            resourceInputs["informationInterval"] = args ? args.informationInterval : undefined;
            resourceInputs["ipsLogs"] = args ? args.ipsLogs : undefined;
            resourceInputs["ipsecErrorsLogs"] = args ? args.ipsecErrorsLogs : undefined;
            resourceInputs["localDiskUsage"] = args ? args.localDiskUsage : undefined;
            resourceInputs["logDiskUsageWarning"] = args ? args.logDiskUsageWarning : undefined;
            resourceInputs["mailto1"] = args ? args.mailto1 : undefined;
            resourceInputs["mailto2"] = args ? args.mailto2 : undefined;
            resourceInputs["mailto3"] = args ? args.mailto3 : undefined;
            resourceInputs["notificationInterval"] = args ? args.notificationInterval : undefined;
            resourceInputs["pppErrorsLogs"] = args ? args.pppErrorsLogs : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["sshLogs"] = args ? args.sshLogs : undefined;
            resourceInputs["sslvpnAuthenticationErrorsLogs"] = args ? args.sslvpnAuthenticationErrorsLogs : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["violationTrafficLogs"] = args ? args.violationTrafficLogs : undefined;
            resourceInputs["warningInterval"] = args ? args.warningInterval : undefined;
            resourceInputs["webfilterLogs"] = args ? args.webfilterLogs : undefined;
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
     * Enable/disable administrator login/logout logs in alert email. Valid values: `enable`, `disable`.
     */
    adminLoginLogs?: pulumi.Input<string>;
    /**
     * Alert alert interval in minutes.
     */
    alertInterval?: pulumi.Input<number>;
    /**
     * Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email. Valid values: `enable`, `disable`.
     */
    amcInterfaceBypassMode?: pulumi.Input<string>;
    /**
     * Enable/disable antivirus logs in alert email. Valid values: `enable`, `disable`.
     */
    antivirusLogs?: pulumi.Input<string>;
    /**
     * Enable/disable configuration change logs in alert email. Valid values: `enable`, `disable`.
     */
    configurationChangesLogs?: pulumi.Input<string>;
    /**
     * Critical alert interval in minutes.
     */
    criticalInterval?: pulumi.Input<number>;
    /**
     * Debug alert interval in minutes.
     */
    debugInterval?: pulumi.Input<number>;
    /**
     * Interval between sending alert emails (1 - 99999 min, default = 5).
     */
    emailInterval?: pulumi.Input<number>;
    /**
     * Emergency alert interval in minutes.
     */
    emergencyInterval?: pulumi.Input<number>;
    /**
     * Error alert interval in minutes.
     */
    errorInterval?: pulumi.Input<number>;
    /**
     * Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days). On FortiOS versions 6.2.0-7.2.0: default = 100. On FortiOS versions 7.2.1-7.2.6: default = 15.
     */
    fdsLicenseExpiringDays?: pulumi.Input<number>;
    /**
     * Enable/disable FortiGuard license expiration warnings in alert email. Valid values: `enable`, `disable`.
     */
    fdsLicenseExpiringWarning?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGuard update logs in alert email. Valid values: `enable`, `disable`.
     */
    fdsUpdateLogs?: pulumi.Input<string>;
    /**
     * How to filter log messages that are sent to alert emails. Valid values: `category`, `threshold`.
     */
    filterMode?: pulumi.Input<string>;
    /**
     * Enable/disable FIPS and Common Criteria error logs in alert email. Valid values: `enable`, `disable`.
     */
    fipsCcErrors?: pulumi.Input<string>;
    /**
     * Enable/disable firewall authentication failure logs in alert email. Valid values: `enable`, `disable`.
     */
    firewallAuthenticationFailureLogs?: pulumi.Input<string>;
    /**
     * Enable/disable FortiCloud log quota warnings in alert email. Valid values: `enable`, `disable`.
     */
    fortiguardLogQuotaWarning?: pulumi.Input<string>;
    /**
     * Enable/disable logging of FSSO collector agent disconnect. Valid values: `enable`, `disable`.
     */
    fssoDisconnectLogs?: pulumi.Input<string>;
    /**
     * Enable/disable HA logs in alert email. Valid values: `enable`, `disable`.
     */
    haLogs?: pulumi.Input<string>;
    /**
     * Information alert interval in minutes.
     */
    informationInterval?: pulumi.Input<number>;
    /**
     * Enable/disable IPS logs in alert email. Valid values: `enable`, `disable`.
     */
    ipsLogs?: pulumi.Input<string>;
    /**
     * Enable/disable IPsec error logs in alert email. Valid values: `enable`, `disable`.
     */
    ipsecErrorsLogs?: pulumi.Input<string>;
    /**
     * Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).
     */
    localDiskUsage?: pulumi.Input<number>;
    /**
     * Enable/disable disk usage warnings in alert email. Valid values: `enable`, `disable`.
     */
    logDiskUsageWarning?: pulumi.Input<string>;
    /**
     * Email address to send alert email to (usually a system administrator). On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    mailto1?: pulumi.Input<string>;
    /**
     * Optional second email address to send alert email to. On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    mailto2?: pulumi.Input<string>;
    /**
     * Optional third email address to send alert email to. On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    mailto3?: pulumi.Input<string>;
    /**
     * Notification alert interval in minutes.
     */
    notificationInterval?: pulumi.Input<number>;
    /**
     * Enable/disable PPP error logs in alert email. Valid values: `enable`, `disable`.
     */
    pppErrorsLogs?: pulumi.Input<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable SSH logs in alert email. Valid values: `enable`, `disable`.
     */
    sshLogs?: pulumi.Input<string>;
    /**
     * Enable/disable SSL-VPN authentication error logs in alert email. Valid values: `enable`, `disable`.
     */
    sslvpnAuthenticationErrorsLogs?: pulumi.Input<string>;
    /**
     * Name that appears in the From: field of alert emails. On FortiOS versions 6.2.0-6.4.0: max. 36 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable violation traffic logs in alert email. Valid values: `enable`, `disable`.
     */
    violationTrafficLogs?: pulumi.Input<string>;
    /**
     * Warning alert interval in minutes.
     */
    warningInterval?: pulumi.Input<number>;
    /**
     * Enable/disable web filter logs in alert email. Valid values: `enable`, `disable`.
     */
    webfilterLogs?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Setting resource.
 */
export interface SettingArgs {
    /**
     * Enable/disable administrator login/logout logs in alert email. Valid values: `enable`, `disable`.
     */
    adminLoginLogs?: pulumi.Input<string>;
    /**
     * Alert alert interval in minutes.
     */
    alertInterval?: pulumi.Input<number>;
    /**
     * Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email. Valid values: `enable`, `disable`.
     */
    amcInterfaceBypassMode?: pulumi.Input<string>;
    /**
     * Enable/disable antivirus logs in alert email. Valid values: `enable`, `disable`.
     */
    antivirusLogs?: pulumi.Input<string>;
    /**
     * Enable/disable configuration change logs in alert email. Valid values: `enable`, `disable`.
     */
    configurationChangesLogs?: pulumi.Input<string>;
    /**
     * Critical alert interval in minutes.
     */
    criticalInterval?: pulumi.Input<number>;
    /**
     * Debug alert interval in minutes.
     */
    debugInterval?: pulumi.Input<number>;
    /**
     * Interval between sending alert emails (1 - 99999 min, default = 5).
     */
    emailInterval?: pulumi.Input<number>;
    /**
     * Emergency alert interval in minutes.
     */
    emergencyInterval?: pulumi.Input<number>;
    /**
     * Error alert interval in minutes.
     */
    errorInterval?: pulumi.Input<number>;
    /**
     * Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days). On FortiOS versions 6.2.0-7.2.0: default = 100. On FortiOS versions 7.2.1-7.2.6: default = 15.
     */
    fdsLicenseExpiringDays?: pulumi.Input<number>;
    /**
     * Enable/disable FortiGuard license expiration warnings in alert email. Valid values: `enable`, `disable`.
     */
    fdsLicenseExpiringWarning?: pulumi.Input<string>;
    /**
     * Enable/disable FortiGuard update logs in alert email. Valid values: `enable`, `disable`.
     */
    fdsUpdateLogs?: pulumi.Input<string>;
    /**
     * How to filter log messages that are sent to alert emails. Valid values: `category`, `threshold`.
     */
    filterMode?: pulumi.Input<string>;
    /**
     * Enable/disable FIPS and Common Criteria error logs in alert email. Valid values: `enable`, `disable`.
     */
    fipsCcErrors?: pulumi.Input<string>;
    /**
     * Enable/disable firewall authentication failure logs in alert email. Valid values: `enable`, `disable`.
     */
    firewallAuthenticationFailureLogs?: pulumi.Input<string>;
    /**
     * Enable/disable FortiCloud log quota warnings in alert email. Valid values: `enable`, `disable`.
     */
    fortiguardLogQuotaWarning?: pulumi.Input<string>;
    /**
     * Enable/disable logging of FSSO collector agent disconnect. Valid values: `enable`, `disable`.
     */
    fssoDisconnectLogs?: pulumi.Input<string>;
    /**
     * Enable/disable HA logs in alert email. Valid values: `enable`, `disable`.
     */
    haLogs?: pulumi.Input<string>;
    /**
     * Information alert interval in minutes.
     */
    informationInterval?: pulumi.Input<number>;
    /**
     * Enable/disable IPS logs in alert email. Valid values: `enable`, `disable`.
     */
    ipsLogs?: pulumi.Input<string>;
    /**
     * Enable/disable IPsec error logs in alert email. Valid values: `enable`, `disable`.
     */
    ipsecErrorsLogs?: pulumi.Input<string>;
    /**
     * Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).
     */
    localDiskUsage?: pulumi.Input<number>;
    /**
     * Enable/disable disk usage warnings in alert email. Valid values: `enable`, `disable`.
     */
    logDiskUsageWarning?: pulumi.Input<string>;
    /**
     * Email address to send alert email to (usually a system administrator). On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    mailto1?: pulumi.Input<string>;
    /**
     * Optional second email address to send alert email to. On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    mailto2?: pulumi.Input<string>;
    /**
     * Optional third email address to send alert email to. On FortiOS versions 6.2.0-6.4.0: max. 64 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    mailto3?: pulumi.Input<string>;
    /**
     * Notification alert interval in minutes.
     */
    notificationInterval?: pulumi.Input<number>;
    /**
     * Enable/disable PPP error logs in alert email. Valid values: `enable`, `disable`.
     */
    pppErrorsLogs?: pulumi.Input<string>;
    /**
     * Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable SSH logs in alert email. Valid values: `enable`, `disable`.
     */
    sshLogs?: pulumi.Input<string>;
    /**
     * Enable/disable SSL-VPN authentication error logs in alert email. Valid values: `enable`, `disable`.
     */
    sslvpnAuthenticationErrorsLogs?: pulumi.Input<string>;
    /**
     * Name that appears in the From: field of alert emails. On FortiOS versions 6.2.0-6.4.0: max. 36 characters. On FortiOS versions >= 6.4.1: max. 63 characters.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable violation traffic logs in alert email. Valid values: `enable`, `disable`.
     */
    violationTrafficLogs?: pulumi.Input<string>;
    /**
     * Warning alert interval in minutes.
     */
    warningInterval?: pulumi.Input<number>;
    /**
     * Enable/disable web filter logs in alert email. Valid values: `enable`, `disable`.
     */
    webfilterLogs?: pulumi.Input<string>;
}
