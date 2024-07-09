// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Report setting configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.report.Setting("trname", {
 *     fortiview: "enable",
 *     pdfReport: "enable",
 *     reportSource: "forward-traffic",
 *     topN: 1000,
 *     webBrowsingThreshold: 3,
 * });
 * ```
 *
 * ## Import
 *
 * Report Setting can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:report/setting:Setting labelname ReportSetting
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:report/setting:Setting labelname ReportSetting
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Setting extends pulumi.CustomResource {
    /**
     * Get an existing Setting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SettingState, opts?: pulumi.CustomResourceOptions): Setting {
        return new Setting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:report/setting:Setting';

    /**
     * Returns true if the given object is an instance of Setting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Setting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Setting.__pulumiType;
    }

    /**
     * Enable/disable historical FortiView. Valid values: `enable`, `disable`.
     */
    public readonly fortiview!: pulumi.Output<string>;
    /**
     * Enable/disable PDF report. Valid values: `enable`, `disable`.
     */
    public readonly pdfReport!: pulumi.Output<string>;
    /**
     * Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
     */
    public readonly reportSource!: pulumi.Output<string>;
    /**
     * Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
     */
    public readonly topN!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Web browsing time calculation threshold (3 - 15 min).
     */
    public readonly webBrowsingThreshold!: pulumi.Output<number>;

    /**
     * Create a Setting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SettingArgs | SettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SettingState | undefined;
            resourceInputs["fortiview"] = state ? state.fortiview : undefined;
            resourceInputs["pdfReport"] = state ? state.pdfReport : undefined;
            resourceInputs["reportSource"] = state ? state.reportSource : undefined;
            resourceInputs["topN"] = state ? state.topN : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["webBrowsingThreshold"] = state ? state.webBrowsingThreshold : undefined;
        } else {
            const args = argsOrState as SettingArgs | undefined;
            resourceInputs["fortiview"] = args ? args.fortiview : undefined;
            resourceInputs["pdfReport"] = args ? args.pdfReport : undefined;
            resourceInputs["reportSource"] = args ? args.reportSource : undefined;
            resourceInputs["topN"] = args ? args.topN : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["webBrowsingThreshold"] = args ? args.webBrowsingThreshold : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Setting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Setting resources.
 */
export interface SettingState {
    /**
     * Enable/disable historical FortiView. Valid values: `enable`, `disable`.
     */
    fortiview?: pulumi.Input<string>;
    /**
     * Enable/disable PDF report. Valid values: `enable`, `disable`.
     */
    pdfReport?: pulumi.Input<string>;
    /**
     * Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
     */
    reportSource?: pulumi.Input<string>;
    /**
     * Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
     */
    topN?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Web browsing time calculation threshold (3 - 15 min).
     */
    webBrowsingThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Setting resource.
 */
export interface SettingArgs {
    /**
     * Enable/disable historical FortiView. Valid values: `enable`, `disable`.
     */
    fortiview?: pulumi.Input<string>;
    /**
     * Enable/disable PDF report. Valid values: `enable`, `disable`.
     */
    pdfReport?: pulumi.Input<string>;
    /**
     * Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
     */
    reportSource?: pulumi.Input<string>;
    /**
     * Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
     */
    topN?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Web browsing time calculation threshold (3 - 15 min).
     */
    webBrowsingThreshold?: pulumi.Input<number>;
}
