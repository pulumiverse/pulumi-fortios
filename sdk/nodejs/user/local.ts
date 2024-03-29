// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure local users.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname3 = new fortios.user.Ldap("trname3", {
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
 * const trname = new fortios.user.Local("trname", {
 *     authConcurrentOverride: "disable",
 *     authConcurrentValue: 0,
 *     authtimeout: 0,
 *     ldapServer: trname3.name,
 *     passwdTime: "0000-00-00 00:00:00",
 *     smsServer: "fortiguard",
 *     status: "enable",
 *     twoFactor: "disable",
 *     type: "ldap",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * User Local can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/local:Local labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/local:Local labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Local extends pulumi.CustomResource {
    /**
     * Get an existing Local resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocalState, opts?: pulumi.CustomResourceOptions): Local {
        return new Local(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/local:Local';

    /**
     * Returns true if the given object is an instance of Local.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Local {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Local.__pulumiType;
    }

    /**
     * Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
     */
    public readonly authConcurrentOverride!: pulumi.Output<string>;
    /**
     * Maximum number of concurrent logins permitted from the same user.
     */
    public readonly authConcurrentValue!: pulumi.Output<number>;
    /**
     * Time in minutes before the authentication timeout for a user is reached.
     */
    public readonly authtimeout!: pulumi.Output<number>;
    /**
     * Two-factor recipient's email address.
     */
    public readonly emailTo!: pulumi.Output<string>;
    /**
     * Two-factor recipient's FortiToken serial number.
     */
    public readonly fortitoken!: pulumi.Output<string>;
    /**
     * User ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Name of LDAP server with which the user must authenticate.
     */
    public readonly ldapServer!: pulumi.Output<string>;
    /**
     * User name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * User's password.
     */
    public readonly passwd!: pulumi.Output<string | undefined>;
    /**
     * Password policy to apply to this user, as defined in config user password-policy.
     */
    public readonly passwdPolicy!: pulumi.Output<string>;
    /**
     * Time of the last password update.
     */
    public readonly passwdTime!: pulumi.Output<string>;
    /**
     * IKEv2 Postquantum Preshared Key Identity.
     */
    public readonly ppkIdentity!: pulumi.Output<string>;
    /**
     * IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
     */
    public readonly ppkSecret!: pulumi.Output<string | undefined>;
    /**
     * Quantum Key Distribution (QKD) profile.
     */
    public readonly qkdProfile!: pulumi.Output<string>;
    /**
     * Name of RADIUS server with which the user must authenticate.
     */
    public readonly radiusServer!: pulumi.Output<string>;
    /**
     * Two-factor recipient's SMS server.
     */
    public readonly smsCustomServer!: pulumi.Output<string>;
    /**
     * Two-factor recipient's mobile phone number.
     */
    public readonly smsPhone!: pulumi.Output<string>;
    /**
     * Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
     */
    public readonly smsServer!: pulumi.Output<string>;
    /**
     * Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Name of TACACS+ server with which the user must authenticate.
     */
    public readonly tacacsServer!: pulumi.Output<string>;
    /**
     * Enable/disable two-factor authentication.
     */
    public readonly twoFactor!: pulumi.Output<string>;
    /**
     * Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
     */
    public readonly twoFactorAuthentication!: pulumi.Output<string>;
    /**
     * Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
     */
    public readonly twoFactorNotification!: pulumi.Output<string>;
    /**
     * Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
     */
    public readonly usernameCaseInsensitivity!: pulumi.Output<string>;
    /**
     * Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
     */
    public readonly usernameCaseSensitivity!: pulumi.Output<string>;
    /**
     * Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
     */
    public readonly usernameSensitivity!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
     */
    public readonly workstation!: pulumi.Output<string>;

    /**
     * Create a Local resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocalArgs | LocalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocalState | undefined;
            resourceInputs["authConcurrentOverride"] = state ? state.authConcurrentOverride : undefined;
            resourceInputs["authConcurrentValue"] = state ? state.authConcurrentValue : undefined;
            resourceInputs["authtimeout"] = state ? state.authtimeout : undefined;
            resourceInputs["emailTo"] = state ? state.emailTo : undefined;
            resourceInputs["fortitoken"] = state ? state.fortitoken : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["ldapServer"] = state ? state.ldapServer : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["passwd"] = state ? state.passwd : undefined;
            resourceInputs["passwdPolicy"] = state ? state.passwdPolicy : undefined;
            resourceInputs["passwdTime"] = state ? state.passwdTime : undefined;
            resourceInputs["ppkIdentity"] = state ? state.ppkIdentity : undefined;
            resourceInputs["ppkSecret"] = state ? state.ppkSecret : undefined;
            resourceInputs["qkdProfile"] = state ? state.qkdProfile : undefined;
            resourceInputs["radiusServer"] = state ? state.radiusServer : undefined;
            resourceInputs["smsCustomServer"] = state ? state.smsCustomServer : undefined;
            resourceInputs["smsPhone"] = state ? state.smsPhone : undefined;
            resourceInputs["smsServer"] = state ? state.smsServer : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tacacsServer"] = state ? state.tacacsServer : undefined;
            resourceInputs["twoFactor"] = state ? state.twoFactor : undefined;
            resourceInputs["twoFactorAuthentication"] = state ? state.twoFactorAuthentication : undefined;
            resourceInputs["twoFactorNotification"] = state ? state.twoFactorNotification : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["usernameCaseInsensitivity"] = state ? state.usernameCaseInsensitivity : undefined;
            resourceInputs["usernameCaseSensitivity"] = state ? state.usernameCaseSensitivity : undefined;
            resourceInputs["usernameSensitivity"] = state ? state.usernameSensitivity : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["workstation"] = state ? state.workstation : undefined;
        } else {
            const args = argsOrState as LocalArgs | undefined;
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["authConcurrentOverride"] = args ? args.authConcurrentOverride : undefined;
            resourceInputs["authConcurrentValue"] = args ? args.authConcurrentValue : undefined;
            resourceInputs["authtimeout"] = args ? args.authtimeout : undefined;
            resourceInputs["emailTo"] = args ? args.emailTo : undefined;
            resourceInputs["fortitoken"] = args ? args.fortitoken : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["ldapServer"] = args ? args.ldapServer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["passwd"] = args?.passwd ? pulumi.secret(args.passwd) : undefined;
            resourceInputs["passwdPolicy"] = args ? args.passwdPolicy : undefined;
            resourceInputs["passwdTime"] = args ? args.passwdTime : undefined;
            resourceInputs["ppkIdentity"] = args ? args.ppkIdentity : undefined;
            resourceInputs["ppkSecret"] = args?.ppkSecret ? pulumi.secret(args.ppkSecret) : undefined;
            resourceInputs["qkdProfile"] = args ? args.qkdProfile : undefined;
            resourceInputs["radiusServer"] = args ? args.radiusServer : undefined;
            resourceInputs["smsCustomServer"] = args ? args.smsCustomServer : undefined;
            resourceInputs["smsPhone"] = args ? args.smsPhone : undefined;
            resourceInputs["smsServer"] = args ? args.smsServer : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tacacsServer"] = args ? args.tacacsServer : undefined;
            resourceInputs["twoFactor"] = args ? args.twoFactor : undefined;
            resourceInputs["twoFactorAuthentication"] = args ? args.twoFactorAuthentication : undefined;
            resourceInputs["twoFactorNotification"] = args ? args.twoFactorNotification : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["usernameCaseInsensitivity"] = args ? args.usernameCaseInsensitivity : undefined;
            resourceInputs["usernameCaseSensitivity"] = args ? args.usernameCaseSensitivity : undefined;
            resourceInputs["usernameSensitivity"] = args ? args.usernameSensitivity : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["workstation"] = args ? args.workstation : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["passwd", "ppkSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Local.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Local resources.
 */
export interface LocalState {
    /**
     * Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
     */
    authConcurrentOverride?: pulumi.Input<string>;
    /**
     * Maximum number of concurrent logins permitted from the same user.
     */
    authConcurrentValue?: pulumi.Input<number>;
    /**
     * Time in minutes before the authentication timeout for a user is reached.
     */
    authtimeout?: pulumi.Input<number>;
    /**
     * Two-factor recipient's email address.
     */
    emailTo?: pulumi.Input<string>;
    /**
     * Two-factor recipient's FortiToken serial number.
     */
    fortitoken?: pulumi.Input<string>;
    /**
     * User ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Name of LDAP server with which the user must authenticate.
     */
    ldapServer?: pulumi.Input<string>;
    /**
     * User name.
     */
    name?: pulumi.Input<string>;
    /**
     * User's password.
     */
    passwd?: pulumi.Input<string>;
    /**
     * Password policy to apply to this user, as defined in config user password-policy.
     */
    passwdPolicy?: pulumi.Input<string>;
    /**
     * Time of the last password update.
     */
    passwdTime?: pulumi.Input<string>;
    /**
     * IKEv2 Postquantum Preshared Key Identity.
     */
    ppkIdentity?: pulumi.Input<string>;
    /**
     * IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
     */
    ppkSecret?: pulumi.Input<string>;
    /**
     * Quantum Key Distribution (QKD) profile.
     */
    qkdProfile?: pulumi.Input<string>;
    /**
     * Name of RADIUS server with which the user must authenticate.
     */
    radiusServer?: pulumi.Input<string>;
    /**
     * Two-factor recipient's SMS server.
     */
    smsCustomServer?: pulumi.Input<string>;
    /**
     * Two-factor recipient's mobile phone number.
     */
    smsPhone?: pulumi.Input<string>;
    /**
     * Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
     */
    smsServer?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Name of TACACS+ server with which the user must authenticate.
     */
    tacacsServer?: pulumi.Input<string>;
    /**
     * Enable/disable two-factor authentication.
     */
    twoFactor?: pulumi.Input<string>;
    /**
     * Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
     */
    twoFactorAuthentication?: pulumi.Input<string>;
    /**
     * Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
     */
    twoFactorNotification?: pulumi.Input<string>;
    /**
     * Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
     */
    type?: pulumi.Input<string>;
    /**
     * Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
     */
    usernameCaseInsensitivity?: pulumi.Input<string>;
    /**
     * Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
     */
    usernameCaseSensitivity?: pulumi.Input<string>;
    /**
     * Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
     */
    usernameSensitivity?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
     */
    workstation?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Local resource.
 */
export interface LocalArgs {
    /**
     * Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable`, `disable`.
     */
    authConcurrentOverride?: pulumi.Input<string>;
    /**
     * Maximum number of concurrent logins permitted from the same user.
     */
    authConcurrentValue?: pulumi.Input<number>;
    /**
     * Time in minutes before the authentication timeout for a user is reached.
     */
    authtimeout?: pulumi.Input<number>;
    /**
     * Two-factor recipient's email address.
     */
    emailTo?: pulumi.Input<string>;
    /**
     * Two-factor recipient's FortiToken serial number.
     */
    fortitoken?: pulumi.Input<string>;
    /**
     * User ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Name of LDAP server with which the user must authenticate.
     */
    ldapServer?: pulumi.Input<string>;
    /**
     * User name.
     */
    name?: pulumi.Input<string>;
    /**
     * User's password.
     */
    passwd?: pulumi.Input<string>;
    /**
     * Password policy to apply to this user, as defined in config user password-policy.
     */
    passwdPolicy?: pulumi.Input<string>;
    /**
     * Time of the last password update.
     */
    passwdTime?: pulumi.Input<string>;
    /**
     * IKEv2 Postquantum Preshared Key Identity.
     */
    ppkIdentity?: pulumi.Input<string>;
    /**
     * IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
     */
    ppkSecret?: pulumi.Input<string>;
    /**
     * Quantum Key Distribution (QKD) profile.
     */
    qkdProfile?: pulumi.Input<string>;
    /**
     * Name of RADIUS server with which the user must authenticate.
     */
    radiusServer?: pulumi.Input<string>;
    /**
     * Two-factor recipient's SMS server.
     */
    smsCustomServer?: pulumi.Input<string>;
    /**
     * Two-factor recipient's mobile phone number.
     */
    smsPhone?: pulumi.Input<string>;
    /**
     * Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
     */
    smsServer?: pulumi.Input<string>;
    /**
     * Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
     */
    status: pulumi.Input<string>;
    /**
     * Name of TACACS+ server with which the user must authenticate.
     */
    tacacsServer?: pulumi.Input<string>;
    /**
     * Enable/disable two-factor authentication.
     */
    twoFactor?: pulumi.Input<string>;
    /**
     * Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
     */
    twoFactorAuthentication?: pulumi.Input<string>;
    /**
     * Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
     */
    twoFactorNotification?: pulumi.Input<string>;
    /**
     * Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.
     */
    type: pulumi.Input<string>;
    /**
     * Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable`, `disable`.
     */
    usernameCaseInsensitivity?: pulumi.Input<string>;
    /**
     * Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.
     */
    usernameCaseSensitivity?: pulumi.Input<string>;
    /**
     * Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.
     */
    usernameSensitivity?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
     */
    workstation?: pulumi.Input<string>;
}
