// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * CA certificate.
 *
 * ## Import
 *
 * VpnCertificate Ca can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/certificate/ca:Ca labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/certificate/ca:Ca labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ca extends pulumi.CustomResource {
    /**
     * Get an existing Ca resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CaState, opts?: pulumi.CustomResourceOptions): Ca {
        return new Ca(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/certificate/ca:Ca';

    /**
     * Returns true if the given object is an instance of Ca.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ca {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ca.__pulumiType;
    }

    /**
     * Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
     */
    public readonly autoUpdateDays!: pulumi.Output<number>;
    /**
     * Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
     */
    public readonly autoUpdateDaysWarning!: pulumi.Output<number>;
    /**
     * CA certificate as a PEM file.
     */
    public readonly ca!: pulumi.Output<string>;
    /**
     * CA identifier of the SCEP server.
     */
    public readonly caIdentifier!: pulumi.Output<string>;
    /**
     * URL of the EST server.
     */
    public readonly estUrl!: pulumi.Output<string>;
    /**
     * Enable/disable synchronization of CA across Security Fabric. Valid values: `disable`, `enable`.
     */
    public readonly fabricCa!: pulumi.Output<string>;
    /**
     * Time at which CA was last updated.
     */
    public readonly lastUpdated!: pulumi.Output<number>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
     */
    public readonly obsolete!: pulumi.Output<string>;
    /**
     * Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
     */
    public readonly range!: pulumi.Output<string>;
    /**
     * URL of the SCEP server.
     */
    public readonly scepUrl!: pulumi.Output<string>;
    /**
     * CA certificate source type.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
     */
    public readonly sslInspectionTrusted!: pulumi.Output<string>;
    /**
     * Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
     */
    public readonly trusted!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

    /**
     * Create a Ca resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CaArgs | CaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CaState | undefined;
            resourceInputs["autoUpdateDays"] = state ? state.autoUpdateDays : undefined;
            resourceInputs["autoUpdateDaysWarning"] = state ? state.autoUpdateDaysWarning : undefined;
            resourceInputs["ca"] = state ? state.ca : undefined;
            resourceInputs["caIdentifier"] = state ? state.caIdentifier : undefined;
            resourceInputs["estUrl"] = state ? state.estUrl : undefined;
            resourceInputs["fabricCa"] = state ? state.fabricCa : undefined;
            resourceInputs["lastUpdated"] = state ? state.lastUpdated : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["obsolete"] = state ? state.obsolete : undefined;
            resourceInputs["range"] = state ? state.range : undefined;
            resourceInputs["scepUrl"] = state ? state.scepUrl : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sslInspectionTrusted"] = state ? state.sslInspectionTrusted : undefined;
            resourceInputs["trusted"] = state ? state.trusted : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as CaArgs | undefined;
            if ((!args || args.ca === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ca'");
            }
            resourceInputs["autoUpdateDays"] = args ? args.autoUpdateDays : undefined;
            resourceInputs["autoUpdateDaysWarning"] = args ? args.autoUpdateDaysWarning : undefined;
            resourceInputs["ca"] = args?.ca ? pulumi.secret(args.ca) : undefined;
            resourceInputs["caIdentifier"] = args ? args.caIdentifier : undefined;
            resourceInputs["estUrl"] = args ? args.estUrl : undefined;
            resourceInputs["fabricCa"] = args ? args.fabricCa : undefined;
            resourceInputs["lastUpdated"] = args ? args.lastUpdated : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["obsolete"] = args ? args.obsolete : undefined;
            resourceInputs["range"] = args ? args.range : undefined;
            resourceInputs["scepUrl"] = args ? args.scepUrl : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sslInspectionTrusted"] = args ? args.sslInspectionTrusted : undefined;
            resourceInputs["trusted"] = args ? args.trusted : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["ca"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Ca.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ca resources.
 */
export interface CaState {
    /**
     * Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
     */
    autoUpdateDays?: pulumi.Input<number>;
    /**
     * Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
     */
    autoUpdateDaysWarning?: pulumi.Input<number>;
    /**
     * CA certificate as a PEM file.
     */
    ca?: pulumi.Input<string>;
    /**
     * CA identifier of the SCEP server.
     */
    caIdentifier?: pulumi.Input<string>;
    /**
     * URL of the EST server.
     */
    estUrl?: pulumi.Input<string>;
    /**
     * Enable/disable synchronization of CA across Security Fabric. Valid values: `disable`, `enable`.
     */
    fabricCa?: pulumi.Input<string>;
    /**
     * Time at which CA was last updated.
     */
    lastUpdated?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
     */
    obsolete?: pulumi.Input<string>;
    /**
     * Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
     */
    range?: pulumi.Input<string>;
    /**
     * URL of the SCEP server.
     */
    scepUrl?: pulumi.Input<string>;
    /**
     * CA certificate source type.
     */
    source?: pulumi.Input<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
     */
    sslInspectionTrusted?: pulumi.Input<string>;
    /**
     * Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
     */
    trusted?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ca resource.
 */
export interface CaArgs {
    /**
     * Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
     */
    autoUpdateDays?: pulumi.Input<number>;
    /**
     * Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
     */
    autoUpdateDaysWarning?: pulumi.Input<number>;
    /**
     * CA certificate as a PEM file.
     */
    ca: pulumi.Input<string>;
    /**
     * CA identifier of the SCEP server.
     */
    caIdentifier?: pulumi.Input<string>;
    /**
     * URL of the EST server.
     */
    estUrl?: pulumi.Input<string>;
    /**
     * Enable/disable synchronization of CA across Security Fabric. Valid values: `disable`, `enable`.
     */
    fabricCa?: pulumi.Input<string>;
    /**
     * Time at which CA was last updated.
     */
    lastUpdated?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
     */
    obsolete?: pulumi.Input<string>;
    /**
     * Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
     */
    range?: pulumi.Input<string>;
    /**
     * URL of the SCEP server.
     */
    scepUrl?: pulumi.Input<string>;
    /**
     * CA certificate source type.
     */
    source?: pulumi.Input<string>;
    /**
     * Source IP address for communications to the SCEP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
     */
    sslInspectionTrusted?: pulumi.Input<string>;
    /**
     * Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
     */
    trusted?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
