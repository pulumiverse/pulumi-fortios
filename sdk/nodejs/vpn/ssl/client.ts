// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * client Applies to FortiOS Version `>= 7.0.1`.
 *
 * ## Import
 *
 * VpnSsl Client can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/ssl/client:Client labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/ssl/client:Client labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Client extends pulumi.CustomResource {
    /**
     * Get an existing Client resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientState, opts?: pulumi.CustomResourceOptions): Client {
        return new Client(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/ssl/client:Client';

    /**
     * Returns true if the given object is an instance of Client.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Client {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Client.__pulumiType;
    }

    /**
     * Certificate to offer to SSL-VPN server if it requests one.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Traffic class ID.
     */
    public readonly classId!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Distance for routes added by SSL-VPN (1 - 255).
     */
    public readonly distance!: pulumi.Output<number>;
    /**
     * SSL interface to send/receive traffic over.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IPv4 subnets that the client is protecting.
     */
    public readonly ipv4Subnets!: pulumi.Output<string>;
    /**
     * IPv6 subnets that the client is protecting.
     */
    public readonly ipv6Subnets!: pulumi.Output<string>;
    /**
     * SSL-VPN tunnel name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Authenticate peer's certificate with the peer/peergrp.
     */
    public readonly peer!: pulumi.Output<string>;
    /**
     * SSL-VPN server port.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Priority for routes added by SSL-VPN (0 - 4294967295).
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
     */
    public readonly psk!: pulumi.Output<string | undefined>;
    /**
     * Realm name configured on SSL-VPN server.
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * IPv4, IPv6 or DNS address of the SSL-VPN server.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Username to offer to the peer to authenticate the client.
     */
    public readonly user!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Client resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientArgs | ClientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientState | undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["classId"] = state ? state.classId : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["distance"] = state ? state.distance : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ipv4Subnets"] = state ? state.ipv4Subnets : undefined;
            resourceInputs["ipv6Subnets"] = state ? state.ipv6Subnets : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["peer"] = state ? state.peer : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["psk"] = state ? state.psk : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ClientArgs | undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["classId"] = args ? args.classId : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["distance"] = args ? args.distance : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ipv4Subnets"] = args ? args.ipv4Subnets : undefined;
            resourceInputs["ipv6Subnets"] = args ? args.ipv6Subnets : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["peer"] = args ? args.peer : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["psk"] = args ? args.psk : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Client.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Client resources.
 */
export interface ClientState {
    /**
     * Certificate to offer to SSL-VPN server if it requests one.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Traffic class ID.
     */
    classId?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Distance for routes added by SSL-VPN (1 - 255).
     */
    distance?: pulumi.Input<number>;
    /**
     * SSL interface to send/receive traffic over.
     */
    interface?: pulumi.Input<string>;
    /**
     * IPv4 subnets that the client is protecting.
     */
    ipv4Subnets?: pulumi.Input<string>;
    /**
     * IPv6 subnets that the client is protecting.
     */
    ipv6Subnets?: pulumi.Input<string>;
    /**
     * SSL-VPN tunnel name.
     */
    name?: pulumi.Input<string>;
    /**
     * Authenticate peer's certificate with the peer/peergrp.
     */
    peer?: pulumi.Input<string>;
    /**
     * SSL-VPN server port.
     */
    port?: pulumi.Input<number>;
    /**
     * Priority for routes added by SSL-VPN (0 - 4294967295).
     */
    priority?: pulumi.Input<number>;
    /**
     * Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
     */
    psk?: pulumi.Input<string>;
    /**
     * Realm name configured on SSL-VPN server.
     */
    realm?: pulumi.Input<string>;
    /**
     * IPv4, IPv6 or DNS address of the SSL-VPN server.
     */
    server?: pulumi.Input<string>;
    /**
     * IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Username to offer to the peer to authenticate the client.
     */
    user?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Client resource.
 */
export interface ClientArgs {
    /**
     * Certificate to offer to SSL-VPN server if it requests one.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Traffic class ID.
     */
    classId?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Distance for routes added by SSL-VPN (1 - 255).
     */
    distance?: pulumi.Input<number>;
    /**
     * SSL interface to send/receive traffic over.
     */
    interface?: pulumi.Input<string>;
    /**
     * IPv4 subnets that the client is protecting.
     */
    ipv4Subnets?: pulumi.Input<string>;
    /**
     * IPv6 subnets that the client is protecting.
     */
    ipv6Subnets?: pulumi.Input<string>;
    /**
     * SSL-VPN tunnel name.
     */
    name?: pulumi.Input<string>;
    /**
     * Authenticate peer's certificate with the peer/peergrp.
     */
    peer?: pulumi.Input<string>;
    /**
     * SSL-VPN server port.
     */
    port?: pulumi.Input<number>;
    /**
     * Priority for routes added by SSL-VPN (0 - 4294967295).
     */
    priority?: pulumi.Input<number>;
    /**
     * Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
     */
    psk?: pulumi.Input<string>;
    /**
     * Realm name configured on SSL-VPN server.
     */
    realm?: pulumi.Input<string>;
    /**
     * IPv4, IPv6 or DNS address of the SSL-VPN server.
     */
    server?: pulumi.Input<string>;
    /**
     * IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Username to offer to the peer to authenticate the client.
     */
    user?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
