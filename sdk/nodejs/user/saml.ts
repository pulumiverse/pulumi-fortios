// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * SAML server entry configuration. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const tr3 = new fortios.user.Saml("tr3", {
 *     cert: "Fortinet_Factory",
 *     entityId: "https://1.1.1.1",
 *     idpCert: "cer11",
 *     idpEntityId: "https://1.1.1.1/acc",
 *     idpSingleLogoutUrl: "https://1.1.1.1/lo",
 *     idpSingleSignOnUrl: "https://1.1.1.1/sou",
 *     singleLogoutUrl: "https://1.1.1.1/logout",
 *     singleSignOnUrl: "https://1.1.1.1/sign",
 *     userName: "ad111",
 * });
 * ```
 *
 * ## Import
 *
 * User Saml can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/saml:Saml labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/saml:Saml labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Saml extends pulumi.CustomResource {
    /**
     * Get an existing Saml resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SamlState, opts?: pulumi.CustomResourceOptions): Saml {
        return new Saml(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/saml:Saml';

    /**
     * Returns true if the given object is an instance of Saml.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Saml {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Saml.__pulumiType;
    }

    /**
     * Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly adfsClaim!: pulumi.Output<string>;
    /**
     * URL to verify authentication.
     */
    public readonly authUrl!: pulumi.Output<string | undefined>;
    /**
     * Certificate to sign SAML messages.
     */
    public readonly cert!: pulumi.Output<string>;
    /**
     * Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
     */
    public readonly clockTolerance!: pulumi.Output<number>;
    /**
     * Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
     */
    public readonly digestMethod!: pulumi.Output<string>;
    /**
     * SP entity ID.
     */
    public readonly entityId!: pulumi.Output<string>;
    /**
     * Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
     */
    public readonly groupClaimType!: pulumi.Output<string>;
    /**
     * Group name in assertion statement.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * IDP Certificate name.
     */
    public readonly idpCert!: pulumi.Output<string>;
    /**
     * IDP entity ID.
     */
    public readonly idpEntityId!: pulumi.Output<string>;
    /**
     * IDP single logout url.
     */
    public readonly idpSingleLogoutUrl!: pulumi.Output<string>;
    /**
     * IDP single sign-on URL.
     */
    public readonly idpSingleSignOnUrl!: pulumi.Output<string>;
    /**
     * Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
     */
    public readonly limitRelaystate!: pulumi.Output<string>;
    /**
     * SAML server entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly reauth!: pulumi.Output<string>;
    /**
     * SP single logout URL.
     */
    public readonly singleLogoutUrl!: pulumi.Output<string>;
    /**
     * SP single sign-on URL.
     */
    public readonly singleSignOnUrl!: pulumi.Output<string>;
    /**
     * User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
     */
    public readonly userClaimType!: pulumi.Output<string>;
    /**
     * User name in assertion statement.
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Saml resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SamlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SamlArgs | SamlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SamlState | undefined;
            resourceInputs["adfsClaim"] = state ? state.adfsClaim : undefined;
            resourceInputs["authUrl"] = state ? state.authUrl : undefined;
            resourceInputs["cert"] = state ? state.cert : undefined;
            resourceInputs["clockTolerance"] = state ? state.clockTolerance : undefined;
            resourceInputs["digestMethod"] = state ? state.digestMethod : undefined;
            resourceInputs["entityId"] = state ? state.entityId : undefined;
            resourceInputs["groupClaimType"] = state ? state.groupClaimType : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["idpCert"] = state ? state.idpCert : undefined;
            resourceInputs["idpEntityId"] = state ? state.idpEntityId : undefined;
            resourceInputs["idpSingleLogoutUrl"] = state ? state.idpSingleLogoutUrl : undefined;
            resourceInputs["idpSingleSignOnUrl"] = state ? state.idpSingleSignOnUrl : undefined;
            resourceInputs["limitRelaystate"] = state ? state.limitRelaystate : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["reauth"] = state ? state.reauth : undefined;
            resourceInputs["singleLogoutUrl"] = state ? state.singleLogoutUrl : undefined;
            resourceInputs["singleSignOnUrl"] = state ? state.singleSignOnUrl : undefined;
            resourceInputs["userClaimType"] = state ? state.userClaimType : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SamlArgs | undefined;
            if ((!args || args.entityId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityId'");
            }
            if ((!args || args.idpCert === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idpCert'");
            }
            if ((!args || args.idpEntityId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idpEntityId'");
            }
            if ((!args || args.idpSingleSignOnUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idpSingleSignOnUrl'");
            }
            if ((!args || args.singleSignOnUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'singleSignOnUrl'");
            }
            resourceInputs["adfsClaim"] = args ? args.adfsClaim : undefined;
            resourceInputs["authUrl"] = args ? args.authUrl : undefined;
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["clockTolerance"] = args ? args.clockTolerance : undefined;
            resourceInputs["digestMethod"] = args ? args.digestMethod : undefined;
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["groupClaimType"] = args ? args.groupClaimType : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["idpCert"] = args ? args.idpCert : undefined;
            resourceInputs["idpEntityId"] = args ? args.idpEntityId : undefined;
            resourceInputs["idpSingleLogoutUrl"] = args ? args.idpSingleLogoutUrl : undefined;
            resourceInputs["idpSingleSignOnUrl"] = args ? args.idpSingleSignOnUrl : undefined;
            resourceInputs["limitRelaystate"] = args ? args.limitRelaystate : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["reauth"] = args ? args.reauth : undefined;
            resourceInputs["singleLogoutUrl"] = args ? args.singleLogoutUrl : undefined;
            resourceInputs["singleSignOnUrl"] = args ? args.singleSignOnUrl : undefined;
            resourceInputs["userClaimType"] = args ? args.userClaimType : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Saml.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Saml resources.
 */
export interface SamlState {
    /**
     * Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
     */
    adfsClaim?: pulumi.Input<string>;
    /**
     * URL to verify authentication.
     */
    authUrl?: pulumi.Input<string>;
    /**
     * Certificate to sign SAML messages.
     */
    cert?: pulumi.Input<string>;
    /**
     * Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
     */
    clockTolerance?: pulumi.Input<number>;
    /**
     * Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
     */
    digestMethod?: pulumi.Input<string>;
    /**
     * SP entity ID.
     */
    entityId?: pulumi.Input<string>;
    /**
     * Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
     */
    groupClaimType?: pulumi.Input<string>;
    /**
     * Group name in assertion statement.
     */
    groupName?: pulumi.Input<string>;
    /**
     * IDP Certificate name.
     */
    idpCert?: pulumi.Input<string>;
    /**
     * IDP entity ID.
     */
    idpEntityId?: pulumi.Input<string>;
    /**
     * IDP single logout url.
     */
    idpSingleLogoutUrl?: pulumi.Input<string>;
    /**
     * IDP single sign-on URL.
     */
    idpSingleSignOnUrl?: pulumi.Input<string>;
    /**
     * Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
     */
    limitRelaystate?: pulumi.Input<string>;
    /**
     * SAML server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `enable`, `disable`.
     */
    reauth?: pulumi.Input<string>;
    /**
     * SP single logout URL.
     */
    singleLogoutUrl?: pulumi.Input<string>;
    /**
     * SP single sign-on URL.
     */
    singleSignOnUrl?: pulumi.Input<string>;
    /**
     * User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
     */
    userClaimType?: pulumi.Input<string>;
    /**
     * User name in assertion statement.
     */
    userName?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Saml resource.
 */
export interface SamlArgs {
    /**
     * Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
     */
    adfsClaim?: pulumi.Input<string>;
    /**
     * URL to verify authentication.
     */
    authUrl?: pulumi.Input<string>;
    /**
     * Certificate to sign SAML messages.
     */
    cert?: pulumi.Input<string>;
    /**
     * Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
     */
    clockTolerance?: pulumi.Input<number>;
    /**
     * Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
     */
    digestMethod?: pulumi.Input<string>;
    /**
     * SP entity ID.
     */
    entityId: pulumi.Input<string>;
    /**
     * Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
     */
    groupClaimType?: pulumi.Input<string>;
    /**
     * Group name in assertion statement.
     */
    groupName?: pulumi.Input<string>;
    /**
     * IDP Certificate name.
     */
    idpCert: pulumi.Input<string>;
    /**
     * IDP entity ID.
     */
    idpEntityId: pulumi.Input<string>;
    /**
     * IDP single logout url.
     */
    idpSingleLogoutUrl?: pulumi.Input<string>;
    /**
     * IDP single sign-on URL.
     */
    idpSingleSignOnUrl: pulumi.Input<string>;
    /**
     * Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
     */
    limitRelaystate?: pulumi.Input<string>;
    /**
     * SAML server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `enable`, `disable`.
     */
    reauth?: pulumi.Input<string>;
    /**
     * SP single logout URL.
     */
    singleLogoutUrl?: pulumi.Input<string>;
    /**
     * SP single sign-on URL.
     */
    singleSignOnUrl: pulumi.Input<string>;
    /**
     * User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
     */
    userClaimType?: pulumi.Input<string>;
    /**
     * User name in assertion statement.
     */
    userName?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
