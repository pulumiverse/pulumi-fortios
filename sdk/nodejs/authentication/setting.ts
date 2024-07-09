// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure authentication setting.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.authentication.Setting("trname", {
 *     authHttps: "enable",
 *     captivePortalIp: "0.0.0.0",
 *     captivePortalIp6: "::",
 *     captivePortalPort: 7830,
 *     captivePortalSslPort: 7831,
 *     captivePortalType: "fqdn",
 * });
 * ```
 *
 * ## Import
 *
 * Authentication Setting can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:authentication/setting:Setting labelname AuthenticationSetting
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:authentication/setting:Setting labelname AuthenticationSetting
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
    public static readonly __pulumiType = 'fortios:authentication/setting:Setting';

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
     * Active authentication method (scheme name).
     */
    public readonly activeAuthScheme!: pulumi.Output<string>;
    /**
     * Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
     */
    public readonly authHttps!: pulumi.Output<string>;
    /**
     * Captive portal host name.
     */
    public readonly captivePortal!: pulumi.Output<string>;
    /**
     * IPv6 captive portal host name.
     */
    public readonly captivePortal6!: pulumi.Output<string>;
    /**
     * Captive portal IP address.
     */
    public readonly captivePortalIp!: pulumi.Output<string>;
    /**
     * Captive portal IPv6 address.
     */
    public readonly captivePortalIp6!: pulumi.Output<string>;
    /**
     * Captive portal port number (1 - 65535, default = 7830).
     */
    public readonly captivePortalPort!: pulumi.Output<number>;
    /**
     * Captive portal SSL port number (1 - 65535, default = 7831).
     */
    public readonly captivePortalSslPort!: pulumi.Output<number>;
    /**
     * Captive portal type. Valid values: `fqdn`, `ip`.
     */
    public readonly captivePortalType!: pulumi.Output<string>;
    /**
     * Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
     */
    public readonly certAuth!: pulumi.Output<string>;
    /**
     * Certificate captive portal host name.
     */
    public readonly certCaptivePortal!: pulumi.Output<string>;
    /**
     * Certificate captive portal IP address.
     */
    public readonly certCaptivePortalIp!: pulumi.Output<string>;
    /**
     * Certificate captive portal port number (1 - 65535, default = 7832).
     */
    public readonly certCaptivePortalPort!: pulumi.Output<number>;
    /**
     * Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
     */
    public readonly cookieMaxAge!: pulumi.Output<number>;
    /**
     * Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
     */
    public readonly cookieRefreshDiv!: pulumi.Output<number>;
    /**
     * Address range for the IP based device query. The structure of `devRange` block is documented below.
     */
    public readonly devRanges!: pulumi.Output<outputs.authentication.SettingDevRange[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly ipAuthCookie!: pulumi.Output<string>;
    /**
     * Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
     */
    public readonly persistentCookie!: pulumi.Output<string>;
    /**
     * Single-Sign-On authentication method (scheme name).
     */
    public readonly ssoAuthScheme!: pulumi.Output<string>;
    /**
     * Time of the last update.
     */
    public readonly updateTime!: pulumi.Output<string>;
    /**
     * CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
     */
    public readonly userCertCas!: pulumi.Output<outputs.authentication.SettingUserCertCa[] | undefined>;
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
            resourceInputs["activeAuthScheme"] = state ? state.activeAuthScheme : undefined;
            resourceInputs["authHttps"] = state ? state.authHttps : undefined;
            resourceInputs["captivePortal"] = state ? state.captivePortal : undefined;
            resourceInputs["captivePortal6"] = state ? state.captivePortal6 : undefined;
            resourceInputs["captivePortalIp"] = state ? state.captivePortalIp : undefined;
            resourceInputs["captivePortalIp6"] = state ? state.captivePortalIp6 : undefined;
            resourceInputs["captivePortalPort"] = state ? state.captivePortalPort : undefined;
            resourceInputs["captivePortalSslPort"] = state ? state.captivePortalSslPort : undefined;
            resourceInputs["captivePortalType"] = state ? state.captivePortalType : undefined;
            resourceInputs["certAuth"] = state ? state.certAuth : undefined;
            resourceInputs["certCaptivePortal"] = state ? state.certCaptivePortal : undefined;
            resourceInputs["certCaptivePortalIp"] = state ? state.certCaptivePortalIp : undefined;
            resourceInputs["certCaptivePortalPort"] = state ? state.certCaptivePortalPort : undefined;
            resourceInputs["cookieMaxAge"] = state ? state.cookieMaxAge : undefined;
            resourceInputs["cookieRefreshDiv"] = state ? state.cookieRefreshDiv : undefined;
            resourceInputs["devRanges"] = state ? state.devRanges : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["ipAuthCookie"] = state ? state.ipAuthCookie : undefined;
            resourceInputs["persistentCookie"] = state ? state.persistentCookie : undefined;
            resourceInputs["ssoAuthScheme"] = state ? state.ssoAuthScheme : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["userCertCas"] = state ? state.userCertCas : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SettingArgs | undefined;
            resourceInputs["activeAuthScheme"] = args ? args.activeAuthScheme : undefined;
            resourceInputs["authHttps"] = args ? args.authHttps : undefined;
            resourceInputs["captivePortal"] = args ? args.captivePortal : undefined;
            resourceInputs["captivePortal6"] = args ? args.captivePortal6 : undefined;
            resourceInputs["captivePortalIp"] = args ? args.captivePortalIp : undefined;
            resourceInputs["captivePortalIp6"] = args ? args.captivePortalIp6 : undefined;
            resourceInputs["captivePortalPort"] = args ? args.captivePortalPort : undefined;
            resourceInputs["captivePortalSslPort"] = args ? args.captivePortalSslPort : undefined;
            resourceInputs["captivePortalType"] = args ? args.captivePortalType : undefined;
            resourceInputs["certAuth"] = args ? args.certAuth : undefined;
            resourceInputs["certCaptivePortal"] = args ? args.certCaptivePortal : undefined;
            resourceInputs["certCaptivePortalIp"] = args ? args.certCaptivePortalIp : undefined;
            resourceInputs["certCaptivePortalPort"] = args ? args.certCaptivePortalPort : undefined;
            resourceInputs["cookieMaxAge"] = args ? args.cookieMaxAge : undefined;
            resourceInputs["cookieRefreshDiv"] = args ? args.cookieRefreshDiv : undefined;
            resourceInputs["devRanges"] = args ? args.devRanges : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["ipAuthCookie"] = args ? args.ipAuthCookie : undefined;
            resourceInputs["persistentCookie"] = args ? args.persistentCookie : undefined;
            resourceInputs["ssoAuthScheme"] = args ? args.ssoAuthScheme : undefined;
            resourceInputs["updateTime"] = args ? args.updateTime : undefined;
            resourceInputs["userCertCas"] = args ? args.userCertCas : undefined;
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
     * Active authentication method (scheme name).
     */
    activeAuthScheme?: pulumi.Input<string>;
    /**
     * Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
     */
    authHttps?: pulumi.Input<string>;
    /**
     * Captive portal host name.
     */
    captivePortal?: pulumi.Input<string>;
    /**
     * IPv6 captive portal host name.
     */
    captivePortal6?: pulumi.Input<string>;
    /**
     * Captive portal IP address.
     */
    captivePortalIp?: pulumi.Input<string>;
    /**
     * Captive portal IPv6 address.
     */
    captivePortalIp6?: pulumi.Input<string>;
    /**
     * Captive portal port number (1 - 65535, default = 7830).
     */
    captivePortalPort?: pulumi.Input<number>;
    /**
     * Captive portal SSL port number (1 - 65535, default = 7831).
     */
    captivePortalSslPort?: pulumi.Input<number>;
    /**
     * Captive portal type. Valid values: `fqdn`, `ip`.
     */
    captivePortalType?: pulumi.Input<string>;
    /**
     * Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
     */
    certAuth?: pulumi.Input<string>;
    /**
     * Certificate captive portal host name.
     */
    certCaptivePortal?: pulumi.Input<string>;
    /**
     * Certificate captive portal IP address.
     */
    certCaptivePortalIp?: pulumi.Input<string>;
    /**
     * Certificate captive portal port number (1 - 65535, default = 7832).
     */
    certCaptivePortalPort?: pulumi.Input<number>;
    /**
     * Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
     */
    cookieMaxAge?: pulumi.Input<number>;
    /**
     * Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
     */
    cookieRefreshDiv?: pulumi.Input<number>;
    /**
     * Address range for the IP based device query. The structure of `devRange` block is documented below.
     */
    devRanges?: pulumi.Input<pulumi.Input<inputs.authentication.SettingDevRange>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
     */
    ipAuthCookie?: pulumi.Input<string>;
    /**
     * Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
     */
    persistentCookie?: pulumi.Input<string>;
    /**
     * Single-Sign-On authentication method (scheme name).
     */
    ssoAuthScheme?: pulumi.Input<string>;
    /**
     * Time of the last update.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
     */
    userCertCas?: pulumi.Input<pulumi.Input<inputs.authentication.SettingUserCertCa>[]>;
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
     * Active authentication method (scheme name).
     */
    activeAuthScheme?: pulumi.Input<string>;
    /**
     * Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `enable`, `disable`.
     */
    authHttps?: pulumi.Input<string>;
    /**
     * Captive portal host name.
     */
    captivePortal?: pulumi.Input<string>;
    /**
     * IPv6 captive portal host name.
     */
    captivePortal6?: pulumi.Input<string>;
    /**
     * Captive portal IP address.
     */
    captivePortalIp?: pulumi.Input<string>;
    /**
     * Captive portal IPv6 address.
     */
    captivePortalIp6?: pulumi.Input<string>;
    /**
     * Captive portal port number (1 - 65535, default = 7830).
     */
    captivePortalPort?: pulumi.Input<number>;
    /**
     * Captive portal SSL port number (1 - 65535, default = 7831).
     */
    captivePortalSslPort?: pulumi.Input<number>;
    /**
     * Captive portal type. Valid values: `fqdn`, `ip`.
     */
    captivePortalType?: pulumi.Input<string>;
    /**
     * Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `enable`, `disable`.
     */
    certAuth?: pulumi.Input<string>;
    /**
     * Certificate captive portal host name.
     */
    certCaptivePortal?: pulumi.Input<string>;
    /**
     * Certificate captive portal IP address.
     */
    certCaptivePortalIp?: pulumi.Input<string>;
    /**
     * Certificate captive portal port number (1 - 65535, default = 7832).
     */
    certCaptivePortalPort?: pulumi.Input<number>;
    /**
     * Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
     */
    cookieMaxAge?: pulumi.Input<number>;
    /**
     * Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
     */
    cookieRefreshDiv?: pulumi.Input<number>;
    /**
     * Address range for the IP based device query. The structure of `devRange` block is documented below.
     */
    devRanges?: pulumi.Input<pulumi.Input<inputs.authentication.SettingDevRange>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `enable`, `disable`.
     */
    ipAuthCookie?: pulumi.Input<string>;
    /**
     * Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `enable`, `disable`.
     */
    persistentCookie?: pulumi.Input<string>;
    /**
     * Single-Sign-On authentication method (scheme name).
     */
    ssoAuthScheme?: pulumi.Input<string>;
    /**
     * Time of the last update.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * CA certificate used for client certificate verification. The structure of `userCertCa` block is documented below.
     */
    userCertCas?: pulumi.Input<pulumi.Input<inputs.authentication.SettingUserCertCa>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
