// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to add a VDOM license for FortiOS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test2 = new fortios.system.LicenseVdom("test2", {license: "license"});
 * ```
 */
export class LicenseVdom extends pulumi.CustomResource {
    /**
     * Get an existing LicenseVdom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LicenseVdomState, opts?: pulumi.CustomResourceOptions): LicenseVdom {
        return new LicenseVdom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/licenseVdom:LicenseVdom';

    /**
     * Returns true if the given object is an instance of LicenseVdom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LicenseVdom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LicenseVdom.__pulumiType;
    }

    /**
     * Registration code.
     */
    public readonly license!: pulumi.Output<string>;

    /**
     * Create a LicenseVdom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LicenseVdomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LicenseVdomArgs | LicenseVdomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LicenseVdomState | undefined;
            resourceInputs["license"] = state ? state.license : undefined;
        } else {
            const args = argsOrState as LicenseVdomArgs | undefined;
            if ((!args || args.license === undefined) && !opts.urn) {
                throw new Error("Missing required property 'license'");
            }
            resourceInputs["license"] = args ? args.license : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LicenseVdom.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LicenseVdom resources.
 */
export interface LicenseVdomState {
    /**
     * Registration code.
     */
    license?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LicenseVdom resource.
 */
export interface LicenseVdomArgs {
    /**
     * Registration code.
     */
    license: pulumi.Input<string>;
}
