// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure LDAP server entries.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.user.Ldap("trname", {
 *     accountKeyFilter: "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))",
 *     accountKeyProcessing: "same",
 *     cnid: "cn",
 *     dn: "EIWNCIEW",
 *     groupMemberCheck: "user-attr",
 *     groupObjectFilter: "(&(objectcategory=group)(member=*))",
 *     memberAttr: "memberOf",
 *     passwordExpiryWarning: "disable",
 *     passwordRenewal: "disable",
 *     port: 389,
 *     secure: "disable",
 *     server: "1.1.1.1",
 *     serverIdentityCheck: "disable",
 *     sourceIp: "0.0.0.0",
 *     sslMinProtoVersion: "default",
 *     type: "simple",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * User Ldap can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/ldap:Ldap labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/ldap:Ldap labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ldap extends pulumi.CustomResource {
    /**
     * Get an existing Ldap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LdapState, opts?: pulumi.CustomResourceOptions): Ldap {
        return new Ldap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/ldap:Ldap';

    /**
     * Returns true if the given object is an instance of Ldap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ldap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ldap.__pulumiType;
    }

    /**
     * Account key filter, using the UPN as the search filter.
     */
    public readonly accountKeyFilter!: pulumi.Output<string>;
    /**
     * Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
     */
    public readonly accountKeyProcessing!: pulumi.Output<string>;
    /**
     * Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
     */
    public readonly antiphish!: pulumi.Output<string>;
    /**
     * CA certificate name.
     */
    public readonly caCert!: pulumi.Output<string>;
    /**
     * Client certificate name.
     */
    public readonly clientCert!: pulumi.Output<string>;
    /**
     * Enable/disable using client certificate for TLS authentication. Valid values: `enable`, `disable`.
     */
    public readonly clientCertAuth!: pulumi.Output<string>;
    /**
     * Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
     */
    public readonly cnid!: pulumi.Output<string>;
    /**
     * Distinguished name used to look up entries on the LDAP server.
     */
    public readonly dn!: pulumi.Output<string>;
    /**
     * Filter used for group matching.
     */
    public readonly groupFilter!: pulumi.Output<string>;
    /**
     * Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
     */
    public readonly groupMemberCheck!: pulumi.Output<string>;
    /**
     * Filter used for group searching.
     */
    public readonly groupObjectFilter!: pulumi.Output<string>;
    /**
     * Search base used for group searching.
     */
    public readonly groupSearchBase!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Name of attribute from which to get group membership.
     */
    public readonly memberAttr!: pulumi.Output<string>;
    /**
     * LDAP server entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
     */
    public readonly obtainUserInfo!: pulumi.Output<string>;
    /**
     * Password for initial binding.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Name of attribute to get password hash.
     */
    public readonly passwordAttr!: pulumi.Output<string>;
    /**
     * Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
     */
    public readonly passwordExpiryWarning!: pulumi.Output<string>;
    /**
     * Enable/disable online password renewal. Valid values: `enable`, `disable`.
     */
    public readonly passwordRenewal!: pulumi.Output<string>;
    /**
     * Port to be used for communication with the LDAP server (default = 389).
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Search type. Valid values: `recursive`.
     */
    public readonly searchType!: pulumi.Output<string>;
    /**
     * Secondary LDAP server CN domain name or IP.
     */
    public readonly secondaryServer!: pulumi.Output<string>;
    /**
     * Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
     */
    public readonly secure!: pulumi.Output<string>;
    /**
     * LDAP server CN domain name or IP.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
     */
    public readonly serverIdentityCheck!: pulumi.Output<string>;
    /**
     * Source IP for communications to LDAP server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Source port to be used for communication with the LDAP server.
     */
    public readonly sourcePort!: pulumi.Output<number>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * Tertiary LDAP server CN domain name or IP.
     */
    public readonly tertiaryServer!: pulumi.Output<string>;
    /**
     * Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
     */
    public readonly twoFactor!: pulumi.Output<string>;
    /**
     * Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
     */
    public readonly twoFactorAuthentication!: pulumi.Output<string>;
    /**
     * Filter used to synchronize users to FortiToken Cloud.
     */
    public readonly twoFactorFilter!: pulumi.Output<string>;
    /**
     * Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
     */
    public readonly twoFactorNotification!: pulumi.Output<string>;
    /**
     * Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * MS Exchange server from which to fetch user information.
     */
    public readonly userInfoExchangeServer!: pulumi.Output<string>;
    /**
     * Username (full DN) for initial binding.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ldap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LdapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LdapArgs | LdapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LdapState | undefined;
            resourceInputs["accountKeyFilter"] = state ? state.accountKeyFilter : undefined;
            resourceInputs["accountKeyProcessing"] = state ? state.accountKeyProcessing : undefined;
            resourceInputs["antiphish"] = state ? state.antiphish : undefined;
            resourceInputs["caCert"] = state ? state.caCert : undefined;
            resourceInputs["clientCert"] = state ? state.clientCert : undefined;
            resourceInputs["clientCertAuth"] = state ? state.clientCertAuth : undefined;
            resourceInputs["cnid"] = state ? state.cnid : undefined;
            resourceInputs["dn"] = state ? state.dn : undefined;
            resourceInputs["groupFilter"] = state ? state.groupFilter : undefined;
            resourceInputs["groupMemberCheck"] = state ? state.groupMemberCheck : undefined;
            resourceInputs["groupObjectFilter"] = state ? state.groupObjectFilter : undefined;
            resourceInputs["groupSearchBase"] = state ? state.groupSearchBase : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["memberAttr"] = state ? state.memberAttr : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["obtainUserInfo"] = state ? state.obtainUserInfo : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["passwordAttr"] = state ? state.passwordAttr : undefined;
            resourceInputs["passwordExpiryWarning"] = state ? state.passwordExpiryWarning : undefined;
            resourceInputs["passwordRenewal"] = state ? state.passwordRenewal : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["searchType"] = state ? state.searchType : undefined;
            resourceInputs["secondaryServer"] = state ? state.secondaryServer : undefined;
            resourceInputs["secure"] = state ? state.secure : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["serverIdentityCheck"] = state ? state.serverIdentityCheck : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sourcePort"] = state ? state.sourcePort : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["tertiaryServer"] = state ? state.tertiaryServer : undefined;
            resourceInputs["twoFactor"] = state ? state.twoFactor : undefined;
            resourceInputs["twoFactorAuthentication"] = state ? state.twoFactorAuthentication : undefined;
            resourceInputs["twoFactorFilter"] = state ? state.twoFactorFilter : undefined;
            resourceInputs["twoFactorNotification"] = state ? state.twoFactorNotification : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userInfoExchangeServer"] = state ? state.userInfoExchangeServer : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LdapArgs | undefined;
            if ((!args || args.dn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dn'");
            }
            if ((!args || args.server === undefined) && !opts.urn) {
                throw new Error("Missing required property 'server'");
            }
            resourceInputs["accountKeyFilter"] = args ? args.accountKeyFilter : undefined;
            resourceInputs["accountKeyProcessing"] = args ? args.accountKeyProcessing : undefined;
            resourceInputs["antiphish"] = args ? args.antiphish : undefined;
            resourceInputs["caCert"] = args ? args.caCert : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["clientCertAuth"] = args ? args.clientCertAuth : undefined;
            resourceInputs["cnid"] = args ? args.cnid : undefined;
            resourceInputs["dn"] = args ? args.dn : undefined;
            resourceInputs["groupFilter"] = args ? args.groupFilter : undefined;
            resourceInputs["groupMemberCheck"] = args ? args.groupMemberCheck : undefined;
            resourceInputs["groupObjectFilter"] = args ? args.groupObjectFilter : undefined;
            resourceInputs["groupSearchBase"] = args ? args.groupSearchBase : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["memberAttr"] = args ? args.memberAttr : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["obtainUserInfo"] = args ? args.obtainUserInfo : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["passwordAttr"] = args ? args.passwordAttr : undefined;
            resourceInputs["passwordExpiryWarning"] = args ? args.passwordExpiryWarning : undefined;
            resourceInputs["passwordRenewal"] = args ? args.passwordRenewal : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["searchType"] = args ? args.searchType : undefined;
            resourceInputs["secondaryServer"] = args ? args.secondaryServer : undefined;
            resourceInputs["secure"] = args ? args.secure : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["serverIdentityCheck"] = args ? args.serverIdentityCheck : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sourcePort"] = args ? args.sourcePort : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["tertiaryServer"] = args ? args.tertiaryServer : undefined;
            resourceInputs["twoFactor"] = args ? args.twoFactor : undefined;
            resourceInputs["twoFactorAuthentication"] = args ? args.twoFactorAuthentication : undefined;
            resourceInputs["twoFactorFilter"] = args ? args.twoFactorFilter : undefined;
            resourceInputs["twoFactorNotification"] = args ? args.twoFactorNotification : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userInfoExchangeServer"] = args ? args.userInfoExchangeServer : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Ldap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ldap resources.
 */
export interface LdapState {
    /**
     * Account key filter, using the UPN as the search filter.
     */
    accountKeyFilter?: pulumi.Input<string>;
    /**
     * Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
     */
    accountKeyProcessing?: pulumi.Input<string>;
    /**
     * Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
     */
    antiphish?: pulumi.Input<string>;
    /**
     * CA certificate name.
     */
    caCert?: pulumi.Input<string>;
    /**
     * Client certificate name.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Enable/disable using client certificate for TLS authentication. Valid values: `enable`, `disable`.
     */
    clientCertAuth?: pulumi.Input<string>;
    /**
     * Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
     */
    cnid?: pulumi.Input<string>;
    /**
     * Distinguished name used to look up entries on the LDAP server.
     */
    dn?: pulumi.Input<string>;
    /**
     * Filter used for group matching.
     */
    groupFilter?: pulumi.Input<string>;
    /**
     * Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
     */
    groupMemberCheck?: pulumi.Input<string>;
    /**
     * Filter used for group searching.
     */
    groupObjectFilter?: pulumi.Input<string>;
    /**
     * Search base used for group searching.
     */
    groupSearchBase?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Name of attribute from which to get group membership.
     */
    memberAttr?: pulumi.Input<string>;
    /**
     * LDAP server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
     */
    obtainUserInfo?: pulumi.Input<string>;
    /**
     * Password for initial binding.
     */
    password?: pulumi.Input<string>;
    /**
     * Name of attribute to get password hash.
     */
    passwordAttr?: pulumi.Input<string>;
    /**
     * Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
     */
    passwordExpiryWarning?: pulumi.Input<string>;
    /**
     * Enable/disable online password renewal. Valid values: `enable`, `disable`.
     */
    passwordRenewal?: pulumi.Input<string>;
    /**
     * Port to be used for communication with the LDAP server (default = 389).
     */
    port?: pulumi.Input<number>;
    /**
     * Search type. Valid values: `recursive`.
     */
    searchType?: pulumi.Input<string>;
    /**
     * Secondary LDAP server CN domain name or IP.
     */
    secondaryServer?: pulumi.Input<string>;
    /**
     * Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
     */
    secure?: pulumi.Input<string>;
    /**
     * LDAP server CN domain name or IP.
     */
    server?: pulumi.Input<string>;
    /**
     * Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
     */
    serverIdentityCheck?: pulumi.Input<string>;
    /**
     * Source IP for communications to LDAP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Source port to be used for communication with the LDAP server.
     */
    sourcePort?: pulumi.Input<number>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Tertiary LDAP server CN domain name or IP.
     */
    tertiaryServer?: pulumi.Input<string>;
    /**
     * Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
     */
    twoFactor?: pulumi.Input<string>;
    /**
     * Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
     */
    twoFactorAuthentication?: pulumi.Input<string>;
    /**
     * Filter used to synchronize users to FortiToken Cloud.
     */
    twoFactorFilter?: pulumi.Input<string>;
    /**
     * Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
     */
    twoFactorNotification?: pulumi.Input<string>;
    /**
     * Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
     */
    type?: pulumi.Input<string>;
    /**
     * MS Exchange server from which to fetch user information.
     */
    userInfoExchangeServer?: pulumi.Input<string>;
    /**
     * Username (full DN) for initial binding.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ldap resource.
 */
export interface LdapArgs {
    /**
     * Account key filter, using the UPN as the search filter.
     */
    accountKeyFilter?: pulumi.Input<string>;
    /**
     * Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
     */
    accountKeyProcessing?: pulumi.Input<string>;
    /**
     * Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
     */
    antiphish?: pulumi.Input<string>;
    /**
     * CA certificate name.
     */
    caCert?: pulumi.Input<string>;
    /**
     * Client certificate name.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Enable/disable using client certificate for TLS authentication. Valid values: `enable`, `disable`.
     */
    clientCertAuth?: pulumi.Input<string>;
    /**
     * Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
     */
    cnid?: pulumi.Input<string>;
    /**
     * Distinguished name used to look up entries on the LDAP server.
     */
    dn: pulumi.Input<string>;
    /**
     * Filter used for group matching.
     */
    groupFilter?: pulumi.Input<string>;
    /**
     * Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
     */
    groupMemberCheck?: pulumi.Input<string>;
    /**
     * Filter used for group searching.
     */
    groupObjectFilter?: pulumi.Input<string>;
    /**
     * Search base used for group searching.
     */
    groupSearchBase?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Name of attribute from which to get group membership.
     */
    memberAttr?: pulumi.Input<string>;
    /**
     * LDAP server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
     */
    obtainUserInfo?: pulumi.Input<string>;
    /**
     * Password for initial binding.
     */
    password?: pulumi.Input<string>;
    /**
     * Name of attribute to get password hash.
     */
    passwordAttr?: pulumi.Input<string>;
    /**
     * Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
     */
    passwordExpiryWarning?: pulumi.Input<string>;
    /**
     * Enable/disable online password renewal. Valid values: `enable`, `disable`.
     */
    passwordRenewal?: pulumi.Input<string>;
    /**
     * Port to be used for communication with the LDAP server (default = 389).
     */
    port?: pulumi.Input<number>;
    /**
     * Search type. Valid values: `recursive`.
     */
    searchType?: pulumi.Input<string>;
    /**
     * Secondary LDAP server CN domain name or IP.
     */
    secondaryServer?: pulumi.Input<string>;
    /**
     * Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
     */
    secure?: pulumi.Input<string>;
    /**
     * LDAP server CN domain name or IP.
     */
    server: pulumi.Input<string>;
    /**
     * Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
     */
    serverIdentityCheck?: pulumi.Input<string>;
    /**
     * Source IP for communications to LDAP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Source port to be used for communication with the LDAP server.
     */
    sourcePort?: pulumi.Input<number>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Tertiary LDAP server CN domain name or IP.
     */
    tertiaryServer?: pulumi.Input<string>;
    /**
     * Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
     */
    twoFactor?: pulumi.Input<string>;
    /**
     * Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
     */
    twoFactorAuthentication?: pulumi.Input<string>;
    /**
     * Filter used to synchronize users to FortiToken Cloud.
     */
    twoFactorFilter?: pulumi.Input<string>;
    /**
     * Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
     */
    twoFactorNotification?: pulumi.Input<string>;
    /**
     * Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
     */
    type?: pulumi.Input<string>;
    /**
     * MS Exchange server from which to fetch user information.
     */
    userInfoExchangeServer?: pulumi.Input<string>;
    /**
     * Username (full DN) for initial binding.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
