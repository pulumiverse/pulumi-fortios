// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FortiSandbox.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Fortisandbox("trname", {
 *     encAlgorithm: "default",
 *     sslMinProtoVersion: "default",
 *     status: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * System Fortisandbox can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/fortisandbox:Fortisandbox labelname SystemFortisandbox
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/fortisandbox:Fortisandbox labelname SystemFortisandbox
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Fortisandbox extends pulumi.CustomResource {
    /**
     * Get an existing Fortisandbox resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortisandboxState, opts?: pulumi.CustomResourceOptions): Fortisandbox {
        return new Fortisandbox(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/fortisandbox:Fortisandbox';

    /**
     * Returns true if the given object is an instance of Fortisandbox.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Fortisandbox {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Fortisandbox.__pulumiType;
    }

    /**
     * Notifier email address.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
     */
    public readonly encAlgorithm!: pulumi.Output<string>;
    /**
     * Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
     */
    public readonly forticloud!: pulumi.Output<string>;
    /**
     * Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
     */
    public readonly inlineScan!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * IPv4 or IPv6 address of the remote FortiSandbox.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Source IP address for communications to FortiSandbox.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Fortisandbox resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FortisandboxArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortisandboxArgs | FortisandboxState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortisandboxState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            resourceInputs["forticloud"] = state ? state.forticloud : undefined;
            resourceInputs["inlineScan"] = state ? state.inlineScan : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FortisandboxArgs | undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            resourceInputs["forticloud"] = args ? args.forticloud : undefined;
            resourceInputs["inlineScan"] = args ? args.inlineScan : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Fortisandbox.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Fortisandbox resources.
 */
export interface FortisandboxState {
    /**
     * Notifier email address.
     */
    email?: pulumi.Input<string>;
    /**
     * Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
     */
    forticloud?: pulumi.Input<string>;
    /**
     * Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
     */
    inlineScan?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * IPv4 or IPv6 address of the remote FortiSandbox.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IP address for communications to FortiSandbox.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Fortisandbox resource.
 */
export interface FortisandboxArgs {
    /**
     * Notifier email address.
     */
    email?: pulumi.Input<string>;
    /**
     * Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
     */
    forticloud?: pulumi.Input<string>;
    /**
     * Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
     */
    inlineScan?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * IPv4 or IPv6 address of the remote FortiSandbox.
     */
    server?: pulumi.Input<string>;
    /**
     * Source IP address for communications to FortiSandbox.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
