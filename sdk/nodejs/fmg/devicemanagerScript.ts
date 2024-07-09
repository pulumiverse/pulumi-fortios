// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports Create/Read/Update/Delete devicemanager script for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.DevicemanagerScript("test1", {
 *     content: `config system interface 
 *  edit port3 
 * \x09 set vdom "root"
 * \x09 set ip 10.7.0.200 255.255.0.0 
 * \x09 set allowaccess ping http https
 * \x09 next 
 *  end
 * `,
 *     description: "description",
 *     target: "remote_device",
 * });
 * ```
 */
export class DevicemanagerScript extends pulumi.CustomResource {
    /**
     * Get an existing DevicemanagerScript resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DevicemanagerScriptState, opts?: pulumi.CustomResourceOptions): DevicemanagerScript {
        return new DevicemanagerScript(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/devicemanagerScript:DevicemanagerScript';

    /**
     * Returns true if the given object is an instance of DevicemanagerScript.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DevicemanagerScript {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DevicemanagerScript.__pulumiType;
    }

    /**
     * ADOM name. default is 'root'.
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Script content, only cli script is supported now
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Script name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
     */
    public readonly target!: pulumi.Output<string | undefined>;

    /**
     * Create a DevicemanagerScript resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DevicemanagerScriptArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DevicemanagerScriptArgs | DevicemanagerScriptState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DevicemanagerScriptState | undefined;
            resourceInputs["adom"] = state ? state.adom : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as DevicemanagerScriptArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            resourceInputs["adom"] = args ? args.adom : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DevicemanagerScript.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DevicemanagerScript resources.
 */
export interface DevicemanagerScriptState {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Script content, only cli script is supported now
     */
    content?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Script name.
     */
    name?: pulumi.Input<string>;
    /**
     * Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
     */
    target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DevicemanagerScript resource.
 */
export interface DevicemanagerScriptArgs {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Script content, only cli script is supported now
     */
    content: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Script name.
     */
    name?: pulumi.Input<string>;
    /**
     * Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
     */
    target?: pulumi.Input<string>;
}
