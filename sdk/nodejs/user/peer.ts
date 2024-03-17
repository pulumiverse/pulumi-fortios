// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure peer users.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname1 = new fortios.user.Peer("trname1", {
 *     ca: "EC-ACC",
 *     cnType: "string",
 *     ldapMode: "password",
 *     mandatoryCaVerify: "enable",
 *     twoFactor: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * User Peer can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:user/peer:Peer labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:user/peer:Peer labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Peer extends pulumi.CustomResource {
    /**
     * Get an existing Peer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PeerState, opts?: pulumi.CustomResourceOptions): Peer {
        return new Peer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/peer:Peer';

    /**
     * Returns true if the given object is an instance of Peer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Peer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Peer.__pulumiType;
    }

    /**
     * Name of the CA certificate as returned by the execute vpn certificate ca list command.
     */
    public readonly ca!: pulumi.Output<string>;
    /**
     * Peer certificate common name.
     */
    public readonly cn!: pulumi.Output<string>;
    /**
     * Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
     */
    public readonly cnType!: pulumi.Output<string>;
    /**
     * Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
     */
    public readonly ldapMode!: pulumi.Output<string>;
    /**
     * Password for LDAP server bind.
     */
    public readonly ldapPassword!: pulumi.Output<string | undefined>;
    /**
     * Name of an LDAP server defined under the user ldap command. Performs client access rights check.
     */
    public readonly ldapServer!: pulumi.Output<string>;
    /**
     * Username for LDAP server bind.
     */
    public readonly ldapUsername!: pulumi.Output<string>;
    /**
     * Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
     */
    public readonly mandatoryCaVerify!: pulumi.Output<string>;
    /**
     * Peer name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Online Certificate Status Protocol (OCSP) server for certificate retrieval.
     */
    public readonly ocspOverrideServer!: pulumi.Output<string>;
    /**
     * Peer's password used for two-factor authentication.
     */
    public readonly passwd!: pulumi.Output<string | undefined>;
    /**
     * Peer certificate name constraints.
     */
    public readonly subject!: pulumi.Output<string>;
    /**
     * Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
     */
    public readonly twoFactor!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Peer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PeerArgs | PeerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PeerState | undefined;
            resourceInputs["ca"] = state ? state.ca : undefined;
            resourceInputs["cn"] = state ? state.cn : undefined;
            resourceInputs["cnType"] = state ? state.cnType : undefined;
            resourceInputs["ldapMode"] = state ? state.ldapMode : undefined;
            resourceInputs["ldapPassword"] = state ? state.ldapPassword : undefined;
            resourceInputs["ldapServer"] = state ? state.ldapServer : undefined;
            resourceInputs["ldapUsername"] = state ? state.ldapUsername : undefined;
            resourceInputs["mandatoryCaVerify"] = state ? state.mandatoryCaVerify : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ocspOverrideServer"] = state ? state.ocspOverrideServer : undefined;
            resourceInputs["passwd"] = state ? state.passwd : undefined;
            resourceInputs["subject"] = state ? state.subject : undefined;
            resourceInputs["twoFactor"] = state ? state.twoFactor : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as PeerArgs | undefined;
            resourceInputs["ca"] = args ? args.ca : undefined;
            resourceInputs["cn"] = args ? args.cn : undefined;
            resourceInputs["cnType"] = args ? args.cnType : undefined;
            resourceInputs["ldapMode"] = args ? args.ldapMode : undefined;
            resourceInputs["ldapPassword"] = args?.ldapPassword ? pulumi.secret(args.ldapPassword) : undefined;
            resourceInputs["ldapServer"] = args ? args.ldapServer : undefined;
            resourceInputs["ldapUsername"] = args ? args.ldapUsername : undefined;
            resourceInputs["mandatoryCaVerify"] = args ? args.mandatoryCaVerify : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ocspOverrideServer"] = args ? args.ocspOverrideServer : undefined;
            resourceInputs["passwd"] = args?.passwd ? pulumi.secret(args.passwd) : undefined;
            resourceInputs["subject"] = args ? args.subject : undefined;
            resourceInputs["twoFactor"] = args ? args.twoFactor : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["ldapPassword", "passwd"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Peer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Peer resources.
 */
export interface PeerState {
    /**
     * Name of the CA certificate as returned by the execute vpn certificate ca list command.
     */
    ca?: pulumi.Input<string>;
    /**
     * Peer certificate common name.
     */
    cn?: pulumi.Input<string>;
    /**
     * Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
     */
    cnType?: pulumi.Input<string>;
    /**
     * Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
     */
    ldapMode?: pulumi.Input<string>;
    /**
     * Password for LDAP server bind.
     */
    ldapPassword?: pulumi.Input<string>;
    /**
     * Name of an LDAP server defined under the user ldap command. Performs client access rights check.
     */
    ldapServer?: pulumi.Input<string>;
    /**
     * Username for LDAP server bind.
     */
    ldapUsername?: pulumi.Input<string>;
    /**
     * Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
     */
    mandatoryCaVerify?: pulumi.Input<string>;
    /**
     * Peer name.
     */
    name?: pulumi.Input<string>;
    /**
     * Online Certificate Status Protocol (OCSP) server for certificate retrieval.
     */
    ocspOverrideServer?: pulumi.Input<string>;
    /**
     * Peer's password used for two-factor authentication.
     */
    passwd?: pulumi.Input<string>;
    /**
     * Peer certificate name constraints.
     */
    subject?: pulumi.Input<string>;
    /**
     * Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
     */
    twoFactor?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Peer resource.
 */
export interface PeerArgs {
    /**
     * Name of the CA certificate as returned by the execute vpn certificate ca list command.
     */
    ca?: pulumi.Input<string>;
    /**
     * Peer certificate common name.
     */
    cn?: pulumi.Input<string>;
    /**
     * Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
     */
    cnType?: pulumi.Input<string>;
    /**
     * Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
     */
    ldapMode?: pulumi.Input<string>;
    /**
     * Password for LDAP server bind.
     */
    ldapPassword?: pulumi.Input<string>;
    /**
     * Name of an LDAP server defined under the user ldap command. Performs client access rights check.
     */
    ldapServer?: pulumi.Input<string>;
    /**
     * Username for LDAP server bind.
     */
    ldapUsername?: pulumi.Input<string>;
    /**
     * Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
     */
    mandatoryCaVerify?: pulumi.Input<string>;
    /**
     * Peer name.
     */
    name?: pulumi.Input<string>;
    /**
     * Online Certificate Status Protocol (OCSP) server for certificate retrieval.
     */
    ocspOverrideServer?: pulumi.Input<string>;
    /**
     * Peer's password used for two-factor authentication.
     */
    passwd?: pulumi.Input<string>;
    /**
     * Peer certificate name constraints.
     */
    subject?: pulumi.Input<string>;
    /**
     * Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
     */
    twoFactor?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
