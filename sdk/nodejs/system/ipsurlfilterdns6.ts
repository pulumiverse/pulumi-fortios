// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure IPS URL filter IPv6 DNS servers.
 *
 * ## Import
 *
 * System IpsUrlfilterDns6 can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/ipsurlfilterdns6:Ipsurlfilterdns6 labelname {{address6}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/ipsurlfilterdns6:Ipsurlfilterdns6 labelname {{address6}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ipsurlfilterdns6 extends pulumi.CustomResource {
    /**
     * Get an existing Ipsurlfilterdns6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ipsurlfilterdns6State, opts?: pulumi.CustomResourceOptions): Ipsurlfilterdns6 {
        return new Ipsurlfilterdns6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/ipsurlfilterdns6:Ipsurlfilterdns6';

    /**
     * Returns true if the given object is an instance of Ipsurlfilterdns6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipsurlfilterdns6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipsurlfilterdns6.__pulumiType;
    }

    /**
     * IPv6 address of DNS server.
     */
    public readonly address6!: pulumi.Output<string>;
    /**
     * Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ipsurlfilterdns6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Ipsurlfilterdns6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ipsurlfilterdns6Args | Ipsurlfilterdns6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ipsurlfilterdns6State | undefined;
            resourceInputs["address6"] = state ? state.address6 : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Ipsurlfilterdns6Args | undefined;
            resourceInputs["address6"] = args ? args.address6 : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipsurlfilterdns6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipsurlfilterdns6 resources.
 */
export interface Ipsurlfilterdns6State {
    /**
     * IPv6 address of DNS server.
     */
    address6?: pulumi.Input<string>;
    /**
     * Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipsurlfilterdns6 resource.
 */
export interface Ipsurlfilterdns6Args {
    /**
     * IPv6 address of DNS server.
     */
    address6?: pulumi.Input<string>;
    /**
     * Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
