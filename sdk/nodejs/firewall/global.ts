// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Global firewall settings. Applies to FortiOS Version `>= 7.2.1`.
 *
 * ## Import
 *
 * Firewall Global can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/global:Global labelname FirewallGlobal
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/global:Global labelname FirewallGlobal
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Global extends pulumi.CustomResource {
    /**
     * Get an existing Global resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalState, opts?: pulumi.CustomResourceOptions): Global {
        return new Global(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/global:Global';

    /**
     * Returns true if the given object is an instance of Global.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Global {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Global.__pulumiType;
    }

    /**
     * Persistency of banned IPs across power cycling. Valid values: `disabled`, `permanent-only`, `all`.
     */
    public readonly bannedIpPersistency!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Global resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GlobalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalArgs | GlobalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GlobalState | undefined;
            resourceInputs["bannedIpPersistency"] = state ? state.bannedIpPersistency : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as GlobalArgs | undefined;
            resourceInputs["bannedIpPersistency"] = args ? args.bannedIpPersistency : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Global.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Global resources.
 */
export interface GlobalState {
    /**
     * Persistency of banned IPs across power cycling. Valid values: `disabled`, `permanent-only`, `all`.
     */
    bannedIpPersistency?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Global resource.
 */
export interface GlobalArgs {
    /**
     * Persistency of banned IPs across power cycling. Valid values: `disabled`, `permanent-only`, `all`.
     */
    bannedIpPersistency?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
