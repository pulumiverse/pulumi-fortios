// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test2 = new fortios.system.LicenseVm("test2", {fileContent: "LS0tLS1CRUdJTiBGR1QgVk0gTElDRU5TRS0tLS0tDQpRQUFBQUxXaTdCVnVkV2x3QXJZcC92S2J2Yk5zME5YNWluUW9sVldmcFoxWldJQi9pL2g4c01oR0psWWc5Vkl1DQorSlBJRis1aFphMWwyNm9yNHdiEQE3RnJDeVZnQUFBQWhxWjliWHFLK1hGN2o3dnB3WTB6QXRTaTdOMVM1ZWNxDQpWYmRRREZyYklUdnRvUWNyRU1jV0ltQzFqWWs5dmVoeGlYTG1OV0MwN25BSitYTTJFNmh2b29DMjE1YUwxK2wrDQovUHl5M0VLVnNTNjJDT2hMZHc3UndXajB3V3RqMmZiWg0KLS0tLS1FTkQgRkdUIFZNIExJQ0VOU0UtLS0tLQ0K"});
 * ```
 */
export class LicenseVm extends pulumi.CustomResource {
    /**
     * Get an existing LicenseVm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LicenseVmState, opts?: pulumi.CustomResourceOptions): LicenseVm {
        return new LicenseVm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/licenseVm:LicenseVm';

    /**
     * Returns true if the given object is an instance of LicenseVm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LicenseVm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LicenseVm.__pulumiType;
    }

    /**
     * The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
     */
    public readonly fileContent!: pulumi.Output<string>;

    /**
     * Create a LicenseVm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LicenseVmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LicenseVmArgs | LicenseVmState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LicenseVmState | undefined;
            resourceInputs["fileContent"] = state ? state.fileContent : undefined;
        } else {
            const args = argsOrState as LicenseVmArgs | undefined;
            if ((!args || args.fileContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileContent'");
            }
            resourceInputs["fileContent"] = args ? args.fileContent : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LicenseVm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LicenseVm resources.
 */
export interface LicenseVmState {
    /**
     * The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
     */
    fileContent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LicenseVm resource.
 */
export interface LicenseVmArgs {
    /**
     * The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.
     */
    fileContent: pulumi.Input<string>;
}
