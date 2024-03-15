// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports adding/deleting online FortiGate to/from FortiManager
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.DevicemanagerDevice("test1", {
 *     deviceName: "FGVM64-test",
 *     ipaddr: "192.168.88.101",
 *     password: "",
 *     userid: "admin",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class DevicemanagerDevice extends pulumi.CustomResource {
    /**
     * Get an existing DevicemanagerDevice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DevicemanagerDeviceState, opts?: pulumi.CustomResourceOptions): DevicemanagerDevice {
        return new DevicemanagerDevice(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/devicemanagerDevice:DevicemanagerDevice';

    /**
     * Returns true if the given object is an instance of DevicemanagerDevice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DevicemanagerDevice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DevicemanagerDevice.__pulumiType;
    }

    /**
     * Name or ID of the ADOM where the command is to be executed on.
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Fortigate's device name.
     */
    public readonly deviceName!: pulumi.Output<string>;
    /**
     * Fortigate's ipaddress.
     */
    public readonly ipaddr!: pulumi.Output<string>;
    /**
     * Password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * User name.
     */
    public readonly userid!: pulumi.Output<string>;

    /**
     * Create a DevicemanagerDevice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DevicemanagerDeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DevicemanagerDeviceArgs | DevicemanagerDeviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DevicemanagerDeviceState | undefined;
            resourceInputs["adom"] = state ? state.adom : undefined;
            resourceInputs["deviceName"] = state ? state.deviceName : undefined;
            resourceInputs["ipaddr"] = state ? state.ipaddr : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["userid"] = state ? state.userid : undefined;
        } else {
            const args = argsOrState as DevicemanagerDeviceArgs | undefined;
            if ((!args || args.deviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceName'");
            }
            if ((!args || args.ipaddr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipaddr'");
            }
            if ((!args || args.userid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userid'");
            }
            resourceInputs["adom"] = args ? args.adom : undefined;
            resourceInputs["deviceName"] = args ? args.deviceName : undefined;
            resourceInputs["ipaddr"] = args ? args.ipaddr : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["userid"] = args ? args.userid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DevicemanagerDevice.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DevicemanagerDevice resources.
 */
export interface DevicemanagerDeviceState {
    /**
     * Name or ID of the ADOM where the command is to be executed on.
     */
    adom?: pulumi.Input<string>;
    /**
     * Fortigate's device name.
     */
    deviceName?: pulumi.Input<string>;
    /**
     * Fortigate's ipaddress.
     */
    ipaddr?: pulumi.Input<string>;
    /**
     * Password.
     */
    password?: pulumi.Input<string>;
    /**
     * User name.
     */
    userid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DevicemanagerDevice resource.
 */
export interface DevicemanagerDeviceArgs {
    /**
     * Name or ID of the ADOM where the command is to be executed on.
     */
    adom?: pulumi.Input<string>;
    /**
     * Fortigate's device name.
     */
    deviceName: pulumi.Input<string>;
    /**
     * Fortigate's ipaddress.
     */
    ipaddr: pulumi.Input<string>;
    /**
     * Password.
     */
    password?: pulumi.Input<string>;
    /**
     * User name.
     */
    userid: pulumi.Input<string>;
}
