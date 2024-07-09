// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Local keys and certificates.
 *
 * ## Import
 *
 * VpnCertificate Local can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/certificate/local:Local labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/certificate/local:Local labelname {{name}}
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
    public static readonly __pulumiType = 'fortios:vpn/certificate/local:Local';

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
     * The URL for the ACME CA server (Let's Encrypt is the default provider).
     */
    public readonly acmeCaUrl!: pulumi.Output<string>;
    /**
     * A valid domain that resolves to this Fortigate.
     */
    public readonly acmeDomain!: pulumi.Output<string>;
    /**
     * Contact email address that is required by some CAs like LetsEncrypt.
     */
    public readonly acmeEmail!: pulumi.Output<string>;
    /**
     * Beginning of the renewal window (in days before certificate expiration, 30 by default).
     */
    public readonly acmeRenewWindow!: pulumi.Output<number>;
    /**
     * Length of the RSA private key of the generated cert (Minimum 2048 bits).
     */
    public readonly acmeRsaKeySize!: pulumi.Output<number>;
    /**
     * Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
     */
    public readonly autoRegenerateDays!: pulumi.Output<number>;
    /**
     * Number of days to wait before an expiry warning message is generated (0 = disabled).
     */
    public readonly autoRegenerateDaysWarning!: pulumi.Output<number>;
    /**
     * CA identifier of the CA server for signing via SCEP.
     */
    public readonly caIdentifier!: pulumi.Output<string>;
    /**
     * PEM format certificate.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Path location inside CMP server.
     */
    public readonly cmpPath!: pulumi.Output<string>;
    /**
     * CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
     */
    public readonly cmpRegenerationMethod!: pulumi.Output<string>;
    /**
     * Address and port for CMP server (format = address:port).
     */
    public readonly cmpServer!: pulumi.Output<string>;
    /**
     * CMP server certificate.
     */
    public readonly cmpServerCert!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string>;
    /**
     * Certificate Signing Request.
     */
    public readonly csr!: pulumi.Output<string>;
    /**
     * Certificate enrollment protocol.
     */
    public readonly enrollProtocol!: pulumi.Output<string>;
    /**
     * CA identifier of the CA server for signing via EST.
     */
    public readonly estCaId!: pulumi.Output<string>;
    /**
     * Certificate used to authenticate this FortiGate to EST server.
     */
    public readonly estClientCert!: pulumi.Output<string>;
    /**
     * HTTP Authentication password for signing via EST.
     */
    public readonly estHttpPassword!: pulumi.Output<string>;
    /**
     * HTTP Authentication username for signing via EST.
     */
    public readonly estHttpUsername!: pulumi.Output<string>;
    /**
     * Address and port for EST server (e.g. https://example.com:1234).
     */
    public readonly estServer!: pulumi.Output<string>;
    /**
     * EST server's certificate must be verifiable by this certificate to be authenticated.
     */
    public readonly estServerCert!: pulumi.Output<string>;
    /**
     * EST SRP authentication password.
     */
    public readonly estSrpPassword!: pulumi.Output<string>;
    /**
     * EST SRP authentication username.
     */
    public readonly estSrpUsername!: pulumi.Output<string>;
    /**
     * Local ID the FortiGate uses for authentication as a VPN client.
     */
    public readonly ikeLocalid!: pulumi.Output<string>;
    /**
     * IKE local ID type. Valid values: `asn1dn`, `fqdn`.
     */
    public readonly ikeLocalidType!: pulumi.Output<string>;
    /**
     * Time at which certificate was last updated.
     */
    public readonly lastUpdated!: pulumi.Output<number>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
     */
    public readonly nameEncoding!: pulumi.Output<string>;
    /**
     * Password as a PEM file.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * PEM format key, encrypted with a password.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly privateKeyRetain!: pulumi.Output<string>;
    /**
     * Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    public readonly range!: pulumi.Output<string>;
    /**
     * SCEP server challenge password for auto-regeneration.
     */
    public readonly scepPassword!: pulumi.Output<string | undefined>;
    /**
     * SCEP server URL.
     */
    public readonly scepUrl!: pulumi.Output<string>;
    /**
     * Certificate source type.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Certificate Signing Request State.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Local resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LocalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocalArgs | LocalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocalState | undefined;
            resourceInputs["acmeCaUrl"] = state ? state.acmeCaUrl : undefined;
            resourceInputs["acmeDomain"] = state ? state.acmeDomain : undefined;
            resourceInputs["acmeEmail"] = state ? state.acmeEmail : undefined;
            resourceInputs["acmeRenewWindow"] = state ? state.acmeRenewWindow : undefined;
            resourceInputs["acmeRsaKeySize"] = state ? state.acmeRsaKeySize : undefined;
            resourceInputs["autoRegenerateDays"] = state ? state.autoRegenerateDays : undefined;
            resourceInputs["autoRegenerateDaysWarning"] = state ? state.autoRegenerateDaysWarning : undefined;
            resourceInputs["caIdentifier"] = state ? state.caIdentifier : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["cmpPath"] = state ? state.cmpPath : undefined;
            resourceInputs["cmpRegenerationMethod"] = state ? state.cmpRegenerationMethod : undefined;
            resourceInputs["cmpServer"] = state ? state.cmpServer : undefined;
            resourceInputs["cmpServerCert"] = state ? state.cmpServerCert : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["csr"] = state ? state.csr : undefined;
            resourceInputs["enrollProtocol"] = state ? state.enrollProtocol : undefined;
            resourceInputs["estCaId"] = state ? state.estCaId : undefined;
            resourceInputs["estClientCert"] = state ? state.estClientCert : undefined;
            resourceInputs["estHttpPassword"] = state ? state.estHttpPassword : undefined;
            resourceInputs["estHttpUsername"] = state ? state.estHttpUsername : undefined;
            resourceInputs["estServer"] = state ? state.estServer : undefined;
            resourceInputs["estServerCert"] = state ? state.estServerCert : undefined;
            resourceInputs["estSrpPassword"] = state ? state.estSrpPassword : undefined;
            resourceInputs["estSrpUsername"] = state ? state.estSrpUsername : undefined;
            resourceInputs["ikeLocalid"] = state ? state.ikeLocalid : undefined;
            resourceInputs["ikeLocalidType"] = state ? state.ikeLocalidType : undefined;
            resourceInputs["lastUpdated"] = state ? state.lastUpdated : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nameEncoding"] = state ? state.nameEncoding : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["privateKeyRetain"] = state ? state.privateKeyRetain : undefined;
            resourceInputs["range"] = state ? state.range : undefined;
            resourceInputs["scepPassword"] = state ? state.scepPassword : undefined;
            resourceInputs["scepUrl"] = state ? state.scepUrl : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LocalArgs | undefined;
            resourceInputs["acmeCaUrl"] = args ? args.acmeCaUrl : undefined;
            resourceInputs["acmeDomain"] = args ? args.acmeDomain : undefined;
            resourceInputs["acmeEmail"] = args ? args.acmeEmail : undefined;
            resourceInputs["acmeRenewWindow"] = args ? args.acmeRenewWindow : undefined;
            resourceInputs["acmeRsaKeySize"] = args ? args.acmeRsaKeySize : undefined;
            resourceInputs["autoRegenerateDays"] = args ? args.autoRegenerateDays : undefined;
            resourceInputs["autoRegenerateDaysWarning"] = args ? args.autoRegenerateDaysWarning : undefined;
            resourceInputs["caIdentifier"] = args ? args.caIdentifier : undefined;
            resourceInputs["certificate"] = args?.certificate ? pulumi.secret(args.certificate) : undefined;
            resourceInputs["cmpPath"] = args ? args.cmpPath : undefined;
            resourceInputs["cmpRegenerationMethod"] = args ? args.cmpRegenerationMethod : undefined;
            resourceInputs["cmpServer"] = args ? args.cmpServer : undefined;
            resourceInputs["cmpServerCert"] = args ? args.cmpServerCert : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["csr"] = args ? args.csr : undefined;
            resourceInputs["enrollProtocol"] = args ? args.enrollProtocol : undefined;
            resourceInputs["estCaId"] = args ? args.estCaId : undefined;
            resourceInputs["estClientCert"] = args ? args.estClientCert : undefined;
            resourceInputs["estHttpPassword"] = args ? args.estHttpPassword : undefined;
            resourceInputs["estHttpUsername"] = args ? args.estHttpUsername : undefined;
            resourceInputs["estServer"] = args ? args.estServer : undefined;
            resourceInputs["estServerCert"] = args ? args.estServerCert : undefined;
            resourceInputs["estSrpPassword"] = args ? args.estSrpPassword : undefined;
            resourceInputs["estSrpUsername"] = args ? args.estSrpUsername : undefined;
            resourceInputs["ikeLocalid"] = args ? args.ikeLocalid : undefined;
            resourceInputs["ikeLocalidType"] = args ? args.ikeLocalidType : undefined;
            resourceInputs["lastUpdated"] = args ? args.lastUpdated : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nameEncoding"] = args ? args.nameEncoding : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["privateKey"] = args?.privateKey ? pulumi.secret(args.privateKey) : undefined;
            resourceInputs["privateKeyRetain"] = args ? args.privateKeyRetain : undefined;
            resourceInputs["range"] = args ? args.range : undefined;
            resourceInputs["scepPassword"] = args?.scepPassword ? pulumi.secret(args.scepPassword) : undefined;
            resourceInputs["scepUrl"] = args ? args.scepUrl : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["certificate", "password", "privateKey", "scepPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Local.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Local resources.
 */
export interface LocalState {
    /**
     * The URL for the ACME CA server (Let's Encrypt is the default provider).
     */
    acmeCaUrl?: pulumi.Input<string>;
    /**
     * A valid domain that resolves to this Fortigate.
     */
    acmeDomain?: pulumi.Input<string>;
    /**
     * Contact email address that is required by some CAs like LetsEncrypt.
     */
    acmeEmail?: pulumi.Input<string>;
    /**
     * Beginning of the renewal window (in days before certificate expiration, 30 by default).
     */
    acmeRenewWindow?: pulumi.Input<number>;
    /**
     * Length of the RSA private key of the generated cert (Minimum 2048 bits).
     */
    acmeRsaKeySize?: pulumi.Input<number>;
    /**
     * Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
     */
    autoRegenerateDays?: pulumi.Input<number>;
    /**
     * Number of days to wait before an expiry warning message is generated (0 = disabled).
     */
    autoRegenerateDaysWarning?: pulumi.Input<number>;
    /**
     * CA identifier of the CA server for signing via SCEP.
     */
    caIdentifier?: pulumi.Input<string>;
    /**
     * PEM format certificate.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Path location inside CMP server.
     */
    cmpPath?: pulumi.Input<string>;
    /**
     * CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
     */
    cmpRegenerationMethod?: pulumi.Input<string>;
    /**
     * Address and port for CMP server (format = address:port).
     */
    cmpServer?: pulumi.Input<string>;
    /**
     * CMP server certificate.
     */
    cmpServerCert?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Certificate Signing Request.
     */
    csr?: pulumi.Input<string>;
    /**
     * Certificate enrollment protocol.
     */
    enrollProtocol?: pulumi.Input<string>;
    /**
     * CA identifier of the CA server for signing via EST.
     */
    estCaId?: pulumi.Input<string>;
    /**
     * Certificate used to authenticate this FortiGate to EST server.
     */
    estClientCert?: pulumi.Input<string>;
    /**
     * HTTP Authentication password for signing via EST.
     */
    estHttpPassword?: pulumi.Input<string>;
    /**
     * HTTP Authentication username for signing via EST.
     */
    estHttpUsername?: pulumi.Input<string>;
    /**
     * Address and port for EST server (e.g. https://example.com:1234).
     */
    estServer?: pulumi.Input<string>;
    /**
     * EST server's certificate must be verifiable by this certificate to be authenticated.
     */
    estServerCert?: pulumi.Input<string>;
    /**
     * EST SRP authentication password.
     */
    estSrpPassword?: pulumi.Input<string>;
    /**
     * EST SRP authentication username.
     */
    estSrpUsername?: pulumi.Input<string>;
    /**
     * Local ID the FortiGate uses for authentication as a VPN client.
     */
    ikeLocalid?: pulumi.Input<string>;
    /**
     * IKE local ID type. Valid values: `asn1dn`, `fqdn`.
     */
    ikeLocalidType?: pulumi.Input<string>;
    /**
     * Time at which certificate was last updated.
     */
    lastUpdated?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
     */
    nameEncoding?: pulumi.Input<string>;
    /**
     * Password as a PEM file.
     */
    password?: pulumi.Input<string>;
    /**
     * PEM format key, encrypted with a password.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `enable`, `disable`.
     */
    privateKeyRetain?: pulumi.Input<string>;
    /**
     * Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    range?: pulumi.Input<string>;
    /**
     * SCEP server challenge password for auto-regeneration.
     */
    scepPassword?: pulumi.Input<string>;
    /**
     * SCEP server URL.
     */
    scepUrl?: pulumi.Input<string>;
    /**
     * Certificate source type.
     */
    source?: pulumi.Input<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Certificate Signing Request State.
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Local resource.
 */
export interface LocalArgs {
    /**
     * The URL for the ACME CA server (Let's Encrypt is the default provider).
     */
    acmeCaUrl?: pulumi.Input<string>;
    /**
     * A valid domain that resolves to this Fortigate.
     */
    acmeDomain?: pulumi.Input<string>;
    /**
     * Contact email address that is required by some CAs like LetsEncrypt.
     */
    acmeEmail?: pulumi.Input<string>;
    /**
     * Beginning of the renewal window (in days before certificate expiration, 30 by default).
     */
    acmeRenewWindow?: pulumi.Input<number>;
    /**
     * Length of the RSA private key of the generated cert (Minimum 2048 bits).
     */
    acmeRsaKeySize?: pulumi.Input<number>;
    /**
     * Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
     */
    autoRegenerateDays?: pulumi.Input<number>;
    /**
     * Number of days to wait before an expiry warning message is generated (0 = disabled).
     */
    autoRegenerateDaysWarning?: pulumi.Input<number>;
    /**
     * CA identifier of the CA server for signing via SCEP.
     */
    caIdentifier?: pulumi.Input<string>;
    /**
     * PEM format certificate.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Path location inside CMP server.
     */
    cmpPath?: pulumi.Input<string>;
    /**
     * CMP auto-regeneration method. Valid values: `keyupate`, `renewal`.
     */
    cmpRegenerationMethod?: pulumi.Input<string>;
    /**
     * Address and port for CMP server (format = address:port).
     */
    cmpServer?: pulumi.Input<string>;
    /**
     * CMP server certificate.
     */
    cmpServerCert?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Certificate Signing Request.
     */
    csr?: pulumi.Input<string>;
    /**
     * Certificate enrollment protocol.
     */
    enrollProtocol?: pulumi.Input<string>;
    /**
     * CA identifier of the CA server for signing via EST.
     */
    estCaId?: pulumi.Input<string>;
    /**
     * Certificate used to authenticate this FortiGate to EST server.
     */
    estClientCert?: pulumi.Input<string>;
    /**
     * HTTP Authentication password for signing via EST.
     */
    estHttpPassword?: pulumi.Input<string>;
    /**
     * HTTP Authentication username for signing via EST.
     */
    estHttpUsername?: pulumi.Input<string>;
    /**
     * Address and port for EST server (e.g. https://example.com:1234).
     */
    estServer?: pulumi.Input<string>;
    /**
     * EST server's certificate must be verifiable by this certificate to be authenticated.
     */
    estServerCert?: pulumi.Input<string>;
    /**
     * EST SRP authentication password.
     */
    estSrpPassword?: pulumi.Input<string>;
    /**
     * EST SRP authentication username.
     */
    estSrpUsername?: pulumi.Input<string>;
    /**
     * Local ID the FortiGate uses for authentication as a VPN client.
     */
    ikeLocalid?: pulumi.Input<string>;
    /**
     * IKE local ID type. Valid values: `asn1dn`, `fqdn`.
     */
    ikeLocalidType?: pulumi.Input<string>;
    /**
     * Time at which certificate was last updated.
     */
    lastUpdated?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Name encoding method for auto-regeneration. Valid values: `printable`, `utf8`.
     */
    nameEncoding?: pulumi.Input<string>;
    /**
     * Password as a PEM file.
     */
    password?: pulumi.Input<string>;
    /**
     * PEM format key, encrypted with a password.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `enable`, `disable`.
     */
    privateKeyRetain?: pulumi.Input<string>;
    /**
     * Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
     */
    range?: pulumi.Input<string>;
    /**
     * SCEP server challenge password for auto-regeneration.
     */
    scepPassword?: pulumi.Input<string>;
    /**
     * SCEP server URL.
     */
    scepUrl?: pulumi.Input<string>;
    /**
     * Certificate source type.
     */
    source?: pulumi.Input<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Certificate Signing Request State.
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
