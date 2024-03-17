// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FortiManager. Applies to FortiOS Version `<= 7.0.1`.
 *
 * By design considerations, the feature is using the fortios.system.Centralmanagement resource as documented below.
 *
 * ## Example
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Centralmanagement("trname", {
 *     allowMonitor: "enable",
 *     allowPushConfiguration: "enable",
 *     allowPushFirmware: "enable",
 *     allowRemoteFirmwareUpgrade: "enable",
 *     encAlgorithm: "high",
 *     fmg: "\"192.168.52.177\"",
 *     includeDefaultServers: "enable",
 *     mode: "normal",
 *     type: "fortimanager",
 *     vdom: "root",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class Fortimanager extends pulumi.CustomResource {
    /**
     * Get an existing Fortimanager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerState, opts?: pulumi.CustomResourceOptions): Fortimanager {
        return new Fortimanager(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/fortimanager:Fortimanager';

    /**
     * Returns true if the given object is an instance of Fortimanager.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Fortimanager {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Fortimanager.__pulumiType;
    }

    public readonly centralManagement!: pulumi.Output<string>;
    public readonly centralMgmtAutoBackup!: pulumi.Output<string>;
    public readonly centralMgmtScheduleConfigRestore!: pulumi.Output<string>;
    public readonly centralMgmtScheduleScriptRestore!: pulumi.Output<string>;
    public readonly ip!: pulumi.Output<string>;
    public readonly ipsec!: pulumi.Output<string>;
    public readonly vdom!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Fortimanager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FortimanagerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerArgs | FortimanagerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerState | undefined;
            resourceInputs["centralManagement"] = state ? state.centralManagement : undefined;
            resourceInputs["centralMgmtAutoBackup"] = state ? state.centralMgmtAutoBackup : undefined;
            resourceInputs["centralMgmtScheduleConfigRestore"] = state ? state.centralMgmtScheduleConfigRestore : undefined;
            resourceInputs["centralMgmtScheduleScriptRestore"] = state ? state.centralMgmtScheduleScriptRestore : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["ipsec"] = state ? state.ipsec : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FortimanagerArgs | undefined;
            resourceInputs["centralManagement"] = args ? args.centralManagement : undefined;
            resourceInputs["centralMgmtAutoBackup"] = args ? args.centralMgmtAutoBackup : undefined;
            resourceInputs["centralMgmtScheduleConfigRestore"] = args ? args.centralMgmtScheduleConfigRestore : undefined;
            resourceInputs["centralMgmtScheduleScriptRestore"] = args ? args.centralMgmtScheduleScriptRestore : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["ipsec"] = args ? args.ipsec : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Fortimanager.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Fortimanager resources.
 */
export interface FortimanagerState {
    centralManagement?: pulumi.Input<string>;
    centralMgmtAutoBackup?: pulumi.Input<string>;
    centralMgmtScheduleConfigRestore?: pulumi.Input<string>;
    centralMgmtScheduleScriptRestore?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    ipsec?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Fortimanager resource.
 */
export interface FortimanagerArgs {
    centralManagement?: pulumi.Input<string>;
    centralMgmtAutoBackup?: pulumi.Input<string>;
    centralMgmtScheduleConfigRestore?: pulumi.Input<string>;
    centralMgmtScheduleScriptRestore?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    ipsec?: pulumi.Input<string>;
    vdom?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
