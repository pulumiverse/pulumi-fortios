// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure additional port mappings for Internet Services. Applies to FortiOS Version `6.2.4,6.2.6,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.4.0,7.4.1,7.4.2`.
 *
 * ## Import
 *
 * Firewall InternetServiceAppend can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/internetserviceappend:Internetserviceappend labelname FirewallInternetServiceAppend
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/internetserviceappend:Internetserviceappend labelname FirewallInternetServiceAppend
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Internetserviceappend extends pulumi.CustomResource {
    /**
     * Get an existing Internetserviceappend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InternetserviceappendState, opts?: pulumi.CustomResourceOptions): Internetserviceappend {
        return new Internetserviceappend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/internetserviceappend:Internetserviceappend';

    /**
     * Returns true if the given object is an instance of Internetserviceappend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Internetserviceappend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Internetserviceappend.__pulumiType;
    }

    /**
     * Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
     */
    public readonly addrMode!: pulumi.Output<string>;
    /**
     * Appending TCP/UDP/SCTP destination port (1 to 65535).
     */
    public readonly appendPort!: pulumi.Output<number>;
    /**
     * Matching TCP/UDP/SCTP destination port (1 to 65535).
     */
    public readonly matchPort!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Internetserviceappend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InternetserviceappendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InternetserviceappendArgs | InternetserviceappendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InternetserviceappendState | undefined;
            resourceInputs["addrMode"] = state ? state.addrMode : undefined;
            resourceInputs["appendPort"] = state ? state.appendPort : undefined;
            resourceInputs["matchPort"] = state ? state.matchPort : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as InternetserviceappendArgs | undefined;
            resourceInputs["addrMode"] = args ? args.addrMode : undefined;
            resourceInputs["appendPort"] = args ? args.appendPort : undefined;
            resourceInputs["matchPort"] = args ? args.matchPort : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Internetserviceappend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Internetserviceappend resources.
 */
export interface InternetserviceappendState {
    /**
     * Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
     */
    addrMode?: pulumi.Input<string>;
    /**
     * Appending TCP/UDP/SCTP destination port (1 to 65535).
     */
    appendPort?: pulumi.Input<number>;
    /**
     * Matching TCP/UDP/SCTP destination port (1 to 65535).
     */
    matchPort?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Internetserviceappend resource.
 */
export interface InternetserviceappendArgs {
    /**
     * Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
     */
    addrMode?: pulumi.Input<string>;
    /**
     * Appending TCP/UDP/SCTP destination port (1 to 65535).
     */
    appendPort?: pulumi.Input<number>;
    /**
     * Matching TCP/UDP/SCTP destination port (1 to 65535).
     */
    matchPort?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
