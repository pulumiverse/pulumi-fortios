// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * OCSP server configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.vpn.certificate.Ocspserver("trname", {
 *     cert: "ACCVRAIZ1",
 *     sourceIp: "0.0.0.0",
 *     unavailAction: "revoke",
 *     url: "www.tetserv.com",
 * });
 * ```
 *
 * ## Import
 *
 * VpnCertificate OcspServer can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/certificate/ocspserver:Ocspserver labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/certificate/ocspserver:Ocspserver labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ocspserver extends pulumi.CustomResource {
    /**
     * Get an existing Ocspserver resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OcspserverState, opts?: pulumi.CustomResourceOptions): Ocspserver {
        return new Ocspserver(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/certificate/ocspserver:Ocspserver';

    /**
     * Returns true if the given object is an instance of Ocspserver.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ocspserver {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ocspserver.__pulumiType;
    }

    /**
     * OCSP server certificate.
     */
    public readonly cert!: pulumi.Output<string>;
    /**
     * OCSP server entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Secondary OCSP server certificate.
     */
    public readonly secondaryCert!: pulumi.Output<string>;
    /**
     * Secondary OCSP server URL.
     */
    public readonly secondaryUrl!: pulumi.Output<string>;
    /**
     * Source IP address for communications to the OCSP server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
     */
    public readonly unavailAction!: pulumi.Output<string>;
    /**
     * OCSP server URL.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ocspserver resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OcspserverArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OcspserverArgs | OcspserverState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OcspserverState | undefined;
            resourceInputs["cert"] = state ? state.cert : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secondaryCert"] = state ? state.secondaryCert : undefined;
            resourceInputs["secondaryUrl"] = state ? state.secondaryUrl : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["unavailAction"] = state ? state.unavailAction : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as OcspserverArgs | undefined;
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["secondaryCert"] = args ? args.secondaryCert : undefined;
            resourceInputs["secondaryUrl"] = args ? args.secondaryUrl : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["unavailAction"] = args ? args.unavailAction : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ocspserver.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ocspserver resources.
 */
export interface OcspserverState {
    /**
     * OCSP server certificate.
     */
    cert?: pulumi.Input<string>;
    /**
     * OCSP server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Secondary OCSP server certificate.
     */
    secondaryCert?: pulumi.Input<string>;
    /**
     * Secondary OCSP server URL.
     */
    secondaryUrl?: pulumi.Input<string>;
    /**
     * Source IP address for communications to the OCSP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
     */
    unavailAction?: pulumi.Input<string>;
    /**
     * OCSP server URL.
     */
    url?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ocspserver resource.
 */
export interface OcspserverArgs {
    /**
     * OCSP server certificate.
     */
    cert?: pulumi.Input<string>;
    /**
     * OCSP server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Secondary OCSP server certificate.
     */
    secondaryCert?: pulumi.Input<string>;
    /**
     * Secondary OCSP server URL.
     */
    secondaryUrl?: pulumi.Input<string>;
    /**
     * Source IP address for communications to the OCSP server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
     */
    unavailAction?: pulumi.Input<string>;
    /**
     * OCSP server URL.
     */
    url?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
