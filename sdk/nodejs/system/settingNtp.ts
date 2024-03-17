// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to configure Network Time Protocol (NTP) servers of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.system.Ntp`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test2 = new fortios.system.SettingNtp("test2", {
 *     ntpservers: [
 *         "1.1.1.1",
 *         "3.3.3.3",
 *     ],
 *     ntpsync: "disable",
 *     type: "custom",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SettingNtp extends pulumi.CustomResource {
    /**
     * Get an existing SettingNtp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SettingNtpState, opts?: pulumi.CustomResourceOptions): SettingNtp {
        return new SettingNtp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/settingNtp:SettingNtp';

    /**
     * Returns true if the given object is an instance of SettingNtp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SettingNtp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SettingNtp.__pulumiType;
    }

    /**
     * Configure the FortiGate to connect to any available third-party NTP server.
     */
    public readonly ntpservers!: pulumi.Output<string[]>;
    /**
     * Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
     */
    public readonly ntpsync!: pulumi.Output<string>;
    /**
     * Use the FortiGuard NTP server or any other available NTP Server.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a SettingNtp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SettingNtpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SettingNtpArgs | SettingNtpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SettingNtpState | undefined;
            resourceInputs["ntpservers"] = state ? state.ntpservers : undefined;
            resourceInputs["ntpsync"] = state ? state.ntpsync : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SettingNtpArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["ntpservers"] = args ? args.ntpservers : undefined;
            resourceInputs["ntpsync"] = args ? args.ntpsync : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SettingNtp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SettingNtp resources.
 */
export interface SettingNtpState {
    /**
     * Configure the FortiGate to connect to any available third-party NTP server.
     */
    ntpservers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
     */
    ntpsync?: pulumi.Input<string>;
    /**
     * Use the FortiGuard NTP server or any other available NTP Server.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SettingNtp resource.
 */
export interface SettingNtpArgs {
    /**
     * Configure the FortiGate to connect to any available third-party NTP server.
     */
    ntpservers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
     */
    ntpsync?: pulumi.Input<string>;
    /**
     * Use the FortiGuard NTP server or any other available NTP Server.
     */
    type: pulumi.Input<string>;
}
