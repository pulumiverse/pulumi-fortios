// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Global settings for SAML authentication. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.SystemSaml("trname", {
 *     defaultLoginPage: "normal",
 *     defaultProfile: "admin_no_access",
 *     life: 30,
 *     role: "service-provider",
 *     status: "disable",
 *     tolerance: 5,
 * });
 * ```
 *
 * ## Import
 *
 * System Saml can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemSaml:SystemSaml labelname SystemSaml
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemSaml:SystemSaml labelname SystemSaml
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemSaml extends pulumi.CustomResource {
    /**
     * Get an existing SystemSaml resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSamlState, opts?: pulumi.CustomResourceOptions): SystemSaml {
        return new SystemSaml(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSaml:SystemSaml';

    /**
     * Returns true if the given object is an instance of SystemSaml.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSaml {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSaml.__pulumiType;
    }

    /**
     * IdP Binding protocol. Valid values: `post`, `redirect`.
     */
    public readonly bindingProtocol!: pulumi.Output<string>;
    /**
     * Certificate to sign SAML messages.
     */
    public readonly cert!: pulumi.Output<string>;
    /**
     * Choose default login page. Valid values: `normal`, `sso`.
     */
    public readonly defaultLoginPage!: pulumi.Output<string>;
    /**
     * Default profile for new SSO admin.
     */
    public readonly defaultProfile!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * SP entity ID.
     */
    public readonly entityId!: pulumi.Output<string>;
    /**
     * IDP certificate name.
     */
    public readonly idpCert!: pulumi.Output<string>;
    /**
     * IDP entity ID.
     */
    public readonly idpEntityId!: pulumi.Output<string>;
    /**
     * IDP single logout URL.
     */
    public readonly idpSingleLogoutUrl!: pulumi.Output<string>;
    /**
     * IDP single sign-on URL.
     */
    public readonly idpSingleSignOnUrl!: pulumi.Output<string>;
    /**
     * Length of the range of time when the assertion is valid (in minutes).
     */
    public readonly life!: pulumi.Output<number>;
    /**
     * SP portal URL.
     */
    public readonly portalUrl!: pulumi.Output<string>;
    /**
     * SAML role. Valid values: `identity-provider`, `service-provider`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Server address.
     */
    public readonly serverAddress!: pulumi.Output<string>;
    /**
     * Authorized service providers. The structure of `serviceProviders` block is documented below.
     */
    public readonly serviceProviders!: pulumi.Output<outputs.SystemSamlServiceProvider[] | undefined>;
    /**
     * SP single logout URL.
     */
    public readonly singleLogoutUrl!: pulumi.Output<string>;
    /**
     * SP single sign-on URL.
     */
    public readonly singleSignOnUrl!: pulumi.Output<string>;
    /**
     * Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Tolerance to the range of time when the assertion is valid (in minutes).
     */
    public readonly tolerance!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemSaml resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemSamlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSamlArgs | SystemSamlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSamlState | undefined;
            resourceInputs["bindingProtocol"] = state ? state.bindingProtocol : undefined;
            resourceInputs["cert"] = state ? state.cert : undefined;
            resourceInputs["defaultLoginPage"] = state ? state.defaultLoginPage : undefined;
            resourceInputs["defaultProfile"] = state ? state.defaultProfile : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entityId"] = state ? state.entityId : undefined;
            resourceInputs["idpCert"] = state ? state.idpCert : undefined;
            resourceInputs["idpEntityId"] = state ? state.idpEntityId : undefined;
            resourceInputs["idpSingleLogoutUrl"] = state ? state.idpSingleLogoutUrl : undefined;
            resourceInputs["idpSingleSignOnUrl"] = state ? state.idpSingleSignOnUrl : undefined;
            resourceInputs["life"] = state ? state.life : undefined;
            resourceInputs["portalUrl"] = state ? state.portalUrl : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["serverAddress"] = state ? state.serverAddress : undefined;
            resourceInputs["serviceProviders"] = state ? state.serviceProviders : undefined;
            resourceInputs["singleLogoutUrl"] = state ? state.singleLogoutUrl : undefined;
            resourceInputs["singleSignOnUrl"] = state ? state.singleSignOnUrl : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tolerance"] = state ? state.tolerance : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemSamlArgs | undefined;
            resourceInputs["bindingProtocol"] = args ? args.bindingProtocol : undefined;
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["defaultLoginPage"] = args ? args.defaultLoginPage : undefined;
            resourceInputs["defaultProfile"] = args ? args.defaultProfile : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["idpCert"] = args ? args.idpCert : undefined;
            resourceInputs["idpEntityId"] = args ? args.idpEntityId : undefined;
            resourceInputs["idpSingleLogoutUrl"] = args ? args.idpSingleLogoutUrl : undefined;
            resourceInputs["idpSingleSignOnUrl"] = args ? args.idpSingleSignOnUrl : undefined;
            resourceInputs["life"] = args ? args.life : undefined;
            resourceInputs["portalUrl"] = args ? args.portalUrl : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["serverAddress"] = args ? args.serverAddress : undefined;
            resourceInputs["serviceProviders"] = args ? args.serviceProviders : undefined;
            resourceInputs["singleLogoutUrl"] = args ? args.singleLogoutUrl : undefined;
            resourceInputs["singleSignOnUrl"] = args ? args.singleSignOnUrl : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tolerance"] = args ? args.tolerance : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemSaml.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSaml resources.
 */
export interface SystemSamlState {
    /**
     * IdP Binding protocol. Valid values: `post`, `redirect`.
     */
    bindingProtocol?: pulumi.Input<string>;
    /**
     * Certificate to sign SAML messages.
     */
    cert?: pulumi.Input<string>;
    /**
     * Choose default login page. Valid values: `normal`, `sso`.
     */
    defaultLoginPage?: pulumi.Input<string>;
    /**
     * Default profile for new SSO admin.
     */
    defaultProfile?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * SP entity ID.
     */
    entityId?: pulumi.Input<string>;
    /**
     * IDP certificate name.
     */
    idpCert?: pulumi.Input<string>;
    /**
     * IDP entity ID.
     */
    idpEntityId?: pulumi.Input<string>;
    /**
     * IDP single logout URL.
     */
    idpSingleLogoutUrl?: pulumi.Input<string>;
    /**
     * IDP single sign-on URL.
     */
    idpSingleSignOnUrl?: pulumi.Input<string>;
    /**
     * Length of the range of time when the assertion is valid (in minutes).
     */
    life?: pulumi.Input<number>;
    /**
     * SP portal URL.
     */
    portalUrl?: pulumi.Input<string>;
    /**
     * SAML role. Valid values: `identity-provider`, `service-provider`.
     */
    role?: pulumi.Input<string>;
    /**
     * Server address.
     */
    serverAddress?: pulumi.Input<string>;
    /**
     * Authorized service providers. The structure of `serviceProviders` block is documented below.
     */
    serviceProviders?: pulumi.Input<pulumi.Input<inputs.SystemSamlServiceProvider>[]>;
    /**
     * SP single logout URL.
     */
    singleLogoutUrl?: pulumi.Input<string>;
    /**
     * SP single sign-on URL.
     */
    singleSignOnUrl?: pulumi.Input<string>;
    /**
     * Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Tolerance to the range of time when the assertion is valid (in minutes).
     */
    tolerance?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemSaml resource.
 */
export interface SystemSamlArgs {
    /**
     * IdP Binding protocol. Valid values: `post`, `redirect`.
     */
    bindingProtocol?: pulumi.Input<string>;
    /**
     * Certificate to sign SAML messages.
     */
    cert?: pulumi.Input<string>;
    /**
     * Choose default login page. Valid values: `normal`, `sso`.
     */
    defaultLoginPage?: pulumi.Input<string>;
    /**
     * Default profile for new SSO admin.
     */
    defaultProfile?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * SP entity ID.
     */
    entityId?: pulumi.Input<string>;
    /**
     * IDP certificate name.
     */
    idpCert?: pulumi.Input<string>;
    /**
     * IDP entity ID.
     */
    idpEntityId?: pulumi.Input<string>;
    /**
     * IDP single logout URL.
     */
    idpSingleLogoutUrl?: pulumi.Input<string>;
    /**
     * IDP single sign-on URL.
     */
    idpSingleSignOnUrl?: pulumi.Input<string>;
    /**
     * Length of the range of time when the assertion is valid (in minutes).
     */
    life?: pulumi.Input<number>;
    /**
     * SP portal URL.
     */
    portalUrl?: pulumi.Input<string>;
    /**
     * SAML role. Valid values: `identity-provider`, `service-provider`.
     */
    role?: pulumi.Input<string>;
    /**
     * Server address.
     */
    serverAddress?: pulumi.Input<string>;
    /**
     * Authorized service providers. The structure of `serviceProviders` block is documented below.
     */
    serviceProviders?: pulumi.Input<pulumi.Input<inputs.SystemSamlServiceProvider>[]>;
    /**
     * SP single logout URL.
     */
    singleLogoutUrl?: pulumi.Input<string>;
    /**
     * SP single sign-on URL.
     */
    singleSignOnUrl?: pulumi.Input<string>;
    /**
     * Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Tolerance to the range of time when the assertion is valid (in minutes).
     */
    tolerance?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}