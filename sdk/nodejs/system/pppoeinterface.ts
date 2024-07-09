// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure the PPPoE interfaces.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Pppoeinterface("trname", {
 *     authType: "auto",
 *     device: "port3",
 *     dialOnDemand: "disable",
 *     discRetryTimeout: 1,
 *     idleTimeout: 0,
 *     ipunnumbered: "0.0.0.0",
 *     ipv6: "disable",
 *     lcpEchoInterval: 5,
 *     lcpMaxEchoFails: 3,
 *     padtRetryTimeout: 1,
 *     pppoeUnnumberedNegotiate: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * System PppoeInterface can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/pppoeinterface:Pppoeinterface labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/pppoeinterface:Pppoeinterface labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Pppoeinterface extends pulumi.CustomResource {
    /**
     * Get an existing Pppoeinterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PppoeinterfaceState, opts?: pulumi.CustomResourceOptions): Pppoeinterface {
        return new Pppoeinterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/pppoeinterface:Pppoeinterface';

    /**
     * Returns true if the given object is an instance of Pppoeinterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pppoeinterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pppoeinterface.__pulumiType;
    }

    /**
     * PPPoE AC name.
     */
    public readonly acName!: pulumi.Output<string>;
    /**
     * PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
     */
    public readonly authType!: pulumi.Output<string>;
    /**
     * Name for the physical interface.
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
     */
    public readonly dialOnDemand!: pulumi.Output<string>;
    /**
     * PPPoE discovery init timeout value in (0-4294967295 sec).
     */
    public readonly discRetryTimeout!: pulumi.Output<number>;
    /**
     * PPPoE auto disconnect after idle timeout (0-4294967295 sec).
     */
    public readonly idleTimeout!: pulumi.Output<number>;
    /**
     * PPPoE unnumbered IP.
     */
    public readonly ipunnumbered!: pulumi.Output<string>;
    /**
     * Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
     */
    public readonly ipv6!: pulumi.Output<string>;
    /**
     * PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
     */
    public readonly lcpEchoInterval!: pulumi.Output<number>;
    /**
     * Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).
     */
    public readonly lcpMaxEchoFails!: pulumi.Output<number>;
    /**
     * Name of the PPPoE interface.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * PPPoE terminate timeout value in (0-4294967295 sec).
     */
    public readonly padtRetryTimeout!: pulumi.Output<number>;
    /**
     * Enter the password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
     */
    public readonly pppoeUnnumberedNegotiate!: pulumi.Output<string>;
    /**
     * PPPoE service name.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * User name.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Pppoeinterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PppoeinterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PppoeinterfaceArgs | PppoeinterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PppoeinterfaceState | undefined;
            resourceInputs["acName"] = state ? state.acName : undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["dialOnDemand"] = state ? state.dialOnDemand : undefined;
            resourceInputs["discRetryTimeout"] = state ? state.discRetryTimeout : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            resourceInputs["ipunnumbered"] = state ? state.ipunnumbered : undefined;
            resourceInputs["ipv6"] = state ? state.ipv6 : undefined;
            resourceInputs["lcpEchoInterval"] = state ? state.lcpEchoInterval : undefined;
            resourceInputs["lcpMaxEchoFails"] = state ? state.lcpMaxEchoFails : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["padtRetryTimeout"] = state ? state.padtRetryTimeout : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["pppoeUnnumberedNegotiate"] = state ? state.pppoeUnnumberedNegotiate : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as PppoeinterfaceArgs | undefined;
            if ((!args || args.device === undefined) && !opts.urn) {
                throw new Error("Missing required property 'device'");
            }
            resourceInputs["acName"] = args ? args.acName : undefined;
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["dialOnDemand"] = args ? args.dialOnDemand : undefined;
            resourceInputs["discRetryTimeout"] = args ? args.discRetryTimeout : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            resourceInputs["ipunnumbered"] = args ? args.ipunnumbered : undefined;
            resourceInputs["ipv6"] = args ? args.ipv6 : undefined;
            resourceInputs["lcpEchoInterval"] = args ? args.lcpEchoInterval : undefined;
            resourceInputs["lcpMaxEchoFails"] = args ? args.lcpMaxEchoFails : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["padtRetryTimeout"] = args ? args.padtRetryTimeout : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["pppoeUnnumberedNegotiate"] = args ? args.pppoeUnnumberedNegotiate : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Pppoeinterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pppoeinterface resources.
 */
export interface PppoeinterfaceState {
    /**
     * PPPoE AC name.
     */
    acName?: pulumi.Input<string>;
    /**
     * PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
     */
    authType?: pulumi.Input<string>;
    /**
     * Name for the physical interface.
     */
    device?: pulumi.Input<string>;
    /**
     * Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
     */
    dialOnDemand?: pulumi.Input<string>;
    /**
     * PPPoE discovery init timeout value in (0-4294967295 sec).
     */
    discRetryTimeout?: pulumi.Input<number>;
    /**
     * PPPoE auto disconnect after idle timeout (0-4294967295 sec).
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * PPPoE unnumbered IP.
     */
    ipunnumbered?: pulumi.Input<string>;
    /**
     * Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
     */
    ipv6?: pulumi.Input<string>;
    /**
     * PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
     */
    lcpEchoInterval?: pulumi.Input<number>;
    /**
     * Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).
     */
    lcpMaxEchoFails?: pulumi.Input<number>;
    /**
     * Name of the PPPoE interface.
     */
    name?: pulumi.Input<string>;
    /**
     * PPPoE terminate timeout value in (0-4294967295 sec).
     */
    padtRetryTimeout?: pulumi.Input<number>;
    /**
     * Enter the password.
     */
    password?: pulumi.Input<string>;
    /**
     * Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
     */
    pppoeUnnumberedNegotiate?: pulumi.Input<string>;
    /**
     * PPPoE service name.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * User name.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pppoeinterface resource.
 */
export interface PppoeinterfaceArgs {
    /**
     * PPPoE AC name.
     */
    acName?: pulumi.Input<string>;
    /**
     * PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
     */
    authType?: pulumi.Input<string>;
    /**
     * Name for the physical interface.
     */
    device: pulumi.Input<string>;
    /**
     * Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. Valid values: `enable`, `disable`.
     */
    dialOnDemand?: pulumi.Input<string>;
    /**
     * PPPoE discovery init timeout value in (0-4294967295 sec).
     */
    discRetryTimeout?: pulumi.Input<number>;
    /**
     * PPPoE auto disconnect after idle timeout (0-4294967295 sec).
     */
    idleTimeout?: pulumi.Input<number>;
    /**
     * PPPoE unnumbered IP.
     */
    ipunnumbered?: pulumi.Input<string>;
    /**
     * Enable/disable IPv6 Control Protocol (IPv6CP). Valid values: `enable`, `disable`.
     */
    ipv6?: pulumi.Input<string>;
    /**
     * PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
     */
    lcpEchoInterval?: pulumi.Input<number>;
    /**
     * Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).
     */
    lcpMaxEchoFails?: pulumi.Input<number>;
    /**
     * Name of the PPPoE interface.
     */
    name?: pulumi.Input<string>;
    /**
     * PPPoE terminate timeout value in (0-4294967295 sec).
     */
    padtRetryTimeout?: pulumi.Input<number>;
    /**
     * Enter the password.
     */
    password?: pulumi.Input<string>;
    /**
     * Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
     */
    pppoeUnnumberedNegotiate?: pulumi.Input<string>;
    /**
     * PPPoE service name.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * User name.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
