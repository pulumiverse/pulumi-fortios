// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports modifying system admin setting for FortiManager.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.SystemAdmin("test1", {
 *     httpPort: 80,
 *     httpsPort: 443,
 *     idleTimeout: 20,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SystemAdmin extends pulumi.CustomResource {
    /**
     * Get an existing SystemAdmin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemAdminState, opts?: pulumi.CustomResourceOptions): SystemAdmin {
        return new SystemAdmin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/systemAdmin:SystemAdmin';

    /**
     * Returns true if the given object is an instance of SystemAdmin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemAdmin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemAdmin.__pulumiType;
    }

    /**
     * Http port.
     */
    public readonly httpPort!: pulumi.Output<number | undefined>;
    /**
     * Https port.
     */
    public readonly httpsPort!: pulumi.Output<number | undefined>;
    /**
     * Idle Timeout(1-480 minute).
     */
    public readonly idleTimeout!: pulumi.Output<number | undefined>;

    /**
     * Create a SystemAdmin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemAdminArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemAdminArgs | SystemAdminState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemAdminState | undefined;
            resourceInputs["httpPort"] = state ? state.httpPort : undefined;
            resourceInputs["httpsPort"] = state ? state.httpsPort : undefined;
            resourceInputs["idleTimeout"] = state ? state.idleTimeout : undefined;
        } else {
            const args = argsOrState as SystemAdminArgs | undefined;
            resourceInputs["httpPort"] = args ? args.httpPort : undefined;
            resourceInputs["httpsPort"] = args ? args.httpsPort : undefined;
            resourceInputs["idleTimeout"] = args ? args.idleTimeout : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemAdmin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemAdmin resources.
 */
export interface SystemAdminState {
    /**
     * Http port.
     */
    httpPort?: pulumi.Input<number>;
    /**
     * Https port.
     */
    httpsPort?: pulumi.Input<number>;
    /**
     * Idle Timeout(1-480 minute).
     */
    idleTimeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SystemAdmin resource.
 */
export interface SystemAdminArgs {
    /**
     * Http port.
     */
    httpPort?: pulumi.Input<number>;
    /**
     * Https port.
     */
    httpsPort?: pulumi.Input<number>;
    /**
     * Idle Timeout(1-480 minute).
     */
    idleTimeout?: pulumi.Input<number>;
}
