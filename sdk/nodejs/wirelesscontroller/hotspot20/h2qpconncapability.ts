// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Configure connection capability.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.wirelesscontroller.hotspot20.H2qpconncapability("trname", {
 *     espPort: "unknown",
 *     ftpPort: "unknown",
 *     httpPort: "unknown",
 *     icmpPort: "closed",
 *     ikev2Port: "unknown",
 *     ikev2XxPort: "unknown",
 *     pptpVpnPort: "unknown",
 *     sshPort: "unknown",
 *     tlsPort: "unknown",
 *     voipTcpPort: "unknown",
 *     voipUdpPort: "unknown",
 * });
 * ```
 *
 * ## Import
 *
 * WirelessControllerHotspot20 H2QpConnCapability can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpconncapability:H2qpconncapability labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:wirelesscontroller/hotspot20/h2qpconncapability:H2qpconncapability labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class H2qpconncapability extends pulumi.CustomResource {
    /**
     * Get an existing H2qpconncapability resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: H2qpconncapabilityState, opts?: pulumi.CustomResourceOptions): H2qpconncapability {
        return new H2qpconncapability(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wirelesscontroller/hotspot20/h2qpconncapability:H2qpconncapability';

    /**
     * Returns true if the given object is an instance of H2qpconncapability.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is H2qpconncapability {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === H2qpconncapability.__pulumiType;
    }

    /**
     * Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly espPort!: pulumi.Output<string>;
    /**
     * Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly ftpPort!: pulumi.Output<string>;
    /**
     * Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly httpPort!: pulumi.Output<string>;
    /**
     * Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly icmpPort!: pulumi.Output<string>;
    /**
     * Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly ikev2Port!: pulumi.Output<string>;
    /**
     * Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly ikev2XxPort!: pulumi.Output<string>;
    /**
     * Connection capability name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly pptpVpnPort!: pulumi.Output<string>;
    /**
     * Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly sshPort!: pulumi.Output<string>;
    /**
     * Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly tlsPort!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly voipTcpPort!: pulumi.Output<string>;
    /**
     * Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    public readonly voipUdpPort!: pulumi.Output<string>;

    /**
     * Create a H2qpconncapability resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: H2qpconncapabilityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: H2qpconncapabilityArgs | H2qpconncapabilityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as H2qpconncapabilityState | undefined;
            resourceInputs["espPort"] = state ? state.espPort : undefined;
            resourceInputs["ftpPort"] = state ? state.ftpPort : undefined;
            resourceInputs["httpPort"] = state ? state.httpPort : undefined;
            resourceInputs["icmpPort"] = state ? state.icmpPort : undefined;
            resourceInputs["ikev2Port"] = state ? state.ikev2Port : undefined;
            resourceInputs["ikev2XxPort"] = state ? state.ikev2XxPort : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pptpVpnPort"] = state ? state.pptpVpnPort : undefined;
            resourceInputs["sshPort"] = state ? state.sshPort : undefined;
            resourceInputs["tlsPort"] = state ? state.tlsPort : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["voipTcpPort"] = state ? state.voipTcpPort : undefined;
            resourceInputs["voipUdpPort"] = state ? state.voipUdpPort : undefined;
        } else {
            const args = argsOrState as H2qpconncapabilityArgs | undefined;
            resourceInputs["espPort"] = args ? args.espPort : undefined;
            resourceInputs["ftpPort"] = args ? args.ftpPort : undefined;
            resourceInputs["httpPort"] = args ? args.httpPort : undefined;
            resourceInputs["icmpPort"] = args ? args.icmpPort : undefined;
            resourceInputs["ikev2Port"] = args ? args.ikev2Port : undefined;
            resourceInputs["ikev2XxPort"] = args ? args.ikev2XxPort : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pptpVpnPort"] = args ? args.pptpVpnPort : undefined;
            resourceInputs["sshPort"] = args ? args.sshPort : undefined;
            resourceInputs["tlsPort"] = args ? args.tlsPort : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["voipTcpPort"] = args ? args.voipTcpPort : undefined;
            resourceInputs["voipUdpPort"] = args ? args.voipUdpPort : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(H2qpconncapability.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering H2qpconncapability resources.
 */
export interface H2qpconncapabilityState {
    /**
     * Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
     */
    espPort?: pulumi.Input<string>;
    /**
     * Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    ftpPort?: pulumi.Input<string>;
    /**
     * Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    httpPort?: pulumi.Input<string>;
    /**
     * Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    icmpPort?: pulumi.Input<string>;
    /**
     * Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
     */
    ikev2Port?: pulumi.Input<string>;
    /**
     * Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
     */
    ikev2XxPort?: pulumi.Input<string>;
    /**
     * Connection capability name.
     */
    name?: pulumi.Input<string>;
    /**
     * Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
     */
    pptpVpnPort?: pulumi.Input<string>;
    /**
     * Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
     */
    sshPort?: pulumi.Input<string>;
    /**
     * Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
     */
    tlsPort?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    voipTcpPort?: pulumi.Input<string>;
    /**
     * Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    voipUdpPort?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a H2qpconncapability resource.
 */
export interface H2qpconncapabilityArgs {
    /**
     * Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
     */
    espPort?: pulumi.Input<string>;
    /**
     * Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    ftpPort?: pulumi.Input<string>;
    /**
     * Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    httpPort?: pulumi.Input<string>;
    /**
     * Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    icmpPort?: pulumi.Input<string>;
    /**
     * Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
     */
    ikev2Port?: pulumi.Input<string>;
    /**
     * Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
     */
    ikev2XxPort?: pulumi.Input<string>;
    /**
     * Connection capability name.
     */
    name?: pulumi.Input<string>;
    /**
     * Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
     */
    pptpVpnPort?: pulumi.Input<string>;
    /**
     * Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
     */
    sshPort?: pulumi.Input<string>;
    /**
     * Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
     */
    tlsPort?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    voipTcpPort?: pulumi.Input<string>;
    /**
     * Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
     */
    voipUdpPort?: pulumi.Input<string>;
}
