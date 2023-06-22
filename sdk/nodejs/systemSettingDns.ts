// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to configure DNS of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.SystemDns`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.SystemSettingDns("test1", {
 *     dnsOverTls: "disable",
 *     primary: "208.91.112.53",
 *     secondary: "208.91.112.22",
 * });
 * ```
 */
export class SystemSettingDns extends pulumi.CustomResource {
    /**
     * Get an existing SystemSettingDns resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemSettingDnsState, opts?: pulumi.CustomResourceOptions): SystemSettingDns {
        return new SystemSettingDns(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemSettingDns:SystemSettingDns';

    /**
     * Returns true if the given object is an instance of SystemSettingDns.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemSettingDns {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemSettingDns.__pulumiType;
    }

    /**
     * Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
     */
    public readonly dnsOverTls!: pulumi.Output<string>;
    /**
     * Primary DNS server IP address.
     */
    public readonly primary!: pulumi.Output<string>;
    /**
     * Secondary DNS server IP address.
     */
    public readonly secondary!: pulumi.Output<string>;

    /**
     * Create a SystemSettingDns resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemSettingDnsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemSettingDnsArgs | SystemSettingDnsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemSettingDnsState | undefined;
            resourceInputs["dnsOverTls"] = state ? state.dnsOverTls : undefined;
            resourceInputs["primary"] = state ? state.primary : undefined;
            resourceInputs["secondary"] = state ? state.secondary : undefined;
        } else {
            const args = argsOrState as SystemSettingDnsArgs | undefined;
            resourceInputs["dnsOverTls"] = args ? args.dnsOverTls : undefined;
            resourceInputs["primary"] = args ? args.primary : undefined;
            resourceInputs["secondary"] = args ? args.secondary : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemSettingDns.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemSettingDns resources.
 */
export interface SystemSettingDnsState {
    /**
     * Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
     */
    dnsOverTls?: pulumi.Input<string>;
    /**
     * Primary DNS server IP address.
     */
    primary?: pulumi.Input<string>;
    /**
     * Secondary DNS server IP address.
     */
    secondary?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemSettingDns resource.
 */
export interface SystemSettingDnsArgs {
    /**
     * Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
     */
    dnsOverTls?: pulumi.Input<string>;
    /**
     * Primary DNS server IP address.
     */
    primary?: pulumi.Input<string>;
    /**
     * Secondary DNS server IP address.
     */
    secondary?: pulumi.Input<string>;
}
