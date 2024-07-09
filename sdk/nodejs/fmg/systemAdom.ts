// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports Create/Read/Update/Delete system adom for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.SystemAdom("test1", {
 *     actionWhenConflictsOccurDuringPolicyCheck: "Continue",
 *     autoPushPolicyPackagesWhenDeviceBackOnline: "Enable",
 *     centralManagementFortiap: true,
 *     centralManagementSdwan: false,
 *     centralManagementVpn: false,
 *     mode: "Normal",
 *     performPolicyCheckBeforeEveryInstall: true,
 *     status: 1,
 *     type: "FortiCarrier",
 * });
 * ```
 */
export class SystemAdom extends pulumi.CustomResource {
    /**
     * Get an existing SystemAdom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemAdomState, opts?: pulumi.CustomResourceOptions): SystemAdom {
        return new SystemAdom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/systemAdom:SystemAdom';

    /**
     * Returns true if the given object is an instance of SystemAdom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemAdom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemAdom.__pulumiType;
    }

    /**
     * True or False.
     */
    public readonly actionWhenConflictsOccurDuringPolicyCheck!: pulumi.Output<string | undefined>;
    /**
     * True or False.
     */
    public readonly autoPushPolicyPackagesWhenDeviceBackOnline!: pulumi.Output<string | undefined>;
    /**
     * True or False.
     */
    public readonly centralManagementFortiap!: pulumi.Output<boolean | undefined>;
    /**
     * True or False.
     */
    public readonly centralManagementSdwan!: pulumi.Output<boolean | undefined>;
    /**
     * True or False.
     */
    public readonly centralManagementVpn!: pulumi.Output<boolean | undefined>;
    /**
     * Adom mode: Normal or Backup.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * Administrative Domain name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * True or False.
     */
    public readonly performPolicyCheckBeforeEveryInstall!: pulumi.Output<boolean | undefined>;
    /**
     * Adom status. 0 means off and 1 means on.
     */
    public readonly status!: pulumi.Output<number | undefined>;
    /**
     * Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemAdom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemAdomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemAdomArgs | SystemAdomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemAdomState | undefined;
            resourceInputs["actionWhenConflictsOccurDuringPolicyCheck"] = state ? state.actionWhenConflictsOccurDuringPolicyCheck : undefined;
            resourceInputs["autoPushPolicyPackagesWhenDeviceBackOnline"] = state ? state.autoPushPolicyPackagesWhenDeviceBackOnline : undefined;
            resourceInputs["centralManagementFortiap"] = state ? state.centralManagementFortiap : undefined;
            resourceInputs["centralManagementSdwan"] = state ? state.centralManagementSdwan : undefined;
            resourceInputs["centralManagementVpn"] = state ? state.centralManagementVpn : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["performPolicyCheckBeforeEveryInstall"] = state ? state.performPolicyCheckBeforeEveryInstall : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SystemAdomArgs | undefined;
            resourceInputs["actionWhenConflictsOccurDuringPolicyCheck"] = args ? args.actionWhenConflictsOccurDuringPolicyCheck : undefined;
            resourceInputs["autoPushPolicyPackagesWhenDeviceBackOnline"] = args ? args.autoPushPolicyPackagesWhenDeviceBackOnline : undefined;
            resourceInputs["centralManagementFortiap"] = args ? args.centralManagementFortiap : undefined;
            resourceInputs["centralManagementSdwan"] = args ? args.centralManagementSdwan : undefined;
            resourceInputs["centralManagementVpn"] = args ? args.centralManagementVpn : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["performPolicyCheckBeforeEveryInstall"] = args ? args.performPolicyCheckBeforeEveryInstall : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemAdom.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemAdom resources.
 */
export interface SystemAdomState {
    /**
     * True or False.
     */
    actionWhenConflictsOccurDuringPolicyCheck?: pulumi.Input<string>;
    /**
     * True or False.
     */
    autoPushPolicyPackagesWhenDeviceBackOnline?: pulumi.Input<string>;
    /**
     * True or False.
     */
    centralManagementFortiap?: pulumi.Input<boolean>;
    /**
     * True or False.
     */
    centralManagementSdwan?: pulumi.Input<boolean>;
    /**
     * True or False.
     */
    centralManagementVpn?: pulumi.Input<boolean>;
    /**
     * Adom mode: Normal or Backup.
     */
    mode?: pulumi.Input<string>;
    /**
     * Administrative Domain name.
     */
    name?: pulumi.Input<string>;
    /**
     * True or False.
     */
    performPolicyCheckBeforeEveryInstall?: pulumi.Input<boolean>;
    /**
     * Adom status. 0 means off and 1 means on.
     */
    status?: pulumi.Input<number>;
    /**
     * Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemAdom resource.
 */
export interface SystemAdomArgs {
    /**
     * True or False.
     */
    actionWhenConflictsOccurDuringPolicyCheck?: pulumi.Input<string>;
    /**
     * True or False.
     */
    autoPushPolicyPackagesWhenDeviceBackOnline?: pulumi.Input<string>;
    /**
     * True or False.
     */
    centralManagementFortiap?: pulumi.Input<boolean>;
    /**
     * True or False.
     */
    centralManagementSdwan?: pulumi.Input<boolean>;
    /**
     * True or False.
     */
    centralManagementVpn?: pulumi.Input<boolean>;
    /**
     * Adom mode: Normal or Backup.
     */
    mode?: pulumi.Input<string>;
    /**
     * Administrative Domain name.
     */
    name?: pulumi.Input<string>;
    /**
     * True or False.
     */
    performPolicyCheckBeforeEveryInstall?: pulumi.Input<boolean>;
    /**
     * Adom status. 0 means off and 1 means on.
     */
    status?: pulumi.Input<number>;
    /**
     * Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
     */
    type?: pulumi.Input<string>;
}
