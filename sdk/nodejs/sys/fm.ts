// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure FM. Applies to FortiOS Version `<= 7.0.1`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.sys.Fm("trname", {
 *     autoBackup: "disable",
 *     ip: "0.0.0.0",
 *     ipsec: "disable",
 *     scheduledConfigRestore: "disable",
 *     status: "disable",
 *     vdom: "root",
 * });
 * ```
 *
 * ## Import
 *
 * System Fm can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:sys/fm:Fm labelname SystemFm
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:sys/fm:Fm labelname SystemFm
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Fm extends pulumi.CustomResource {
    /**
     * Get an existing Fm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FmState, opts?: pulumi.CustomResourceOptions): Fm {
        return new Fm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/fm:Fm';

    /**
     * Returns true if the given object is an instance of Fm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Fm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Fm.__pulumiType;
    }

    /**
     * Enable/disable automatic backup. Valid values: `enable`, `disable`.
     */
    public readonly autoBackup!: pulumi.Output<string>;
    /**
     * ID.
     */
    public readonly fosid!: pulumi.Output<string>;
    /**
     * IP address.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Enable/disable IPsec. Valid values: `enable`, `disable`.
     */
    public readonly ipsec!: pulumi.Output<string>;
    /**
     * Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
     */
    public readonly scheduledConfigRestore!: pulumi.Output<string>;
    /**
     * Enable/disable FM. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * VDOM.
     */
    public readonly vdom!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Fm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FmArgs | FmState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FmState | undefined;
            resourceInputs["autoBackup"] = state ? state.autoBackup : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["ipsec"] = state ? state.ipsec : undefined;
            resourceInputs["scheduledConfigRestore"] = state ? state.scheduledConfigRestore : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FmArgs | undefined;
            resourceInputs["autoBackup"] = args ? args.autoBackup : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["ipsec"] = args ? args.ipsec : undefined;
            resourceInputs["scheduledConfigRestore"] = args ? args.scheduledConfigRestore : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Fm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Fm resources.
 */
export interface FmState {
    /**
     * Enable/disable automatic backup. Valid values: `enable`, `disable`.
     */
    autoBackup?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * Enable/disable IPsec. Valid values: `enable`, `disable`.
     */
    ipsec?: pulumi.Input<string>;
    /**
     * Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
     */
    scheduledConfigRestore?: pulumi.Input<string>;
    /**
     * Enable/disable FM. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * VDOM.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Fm resource.
 */
export interface FmArgs {
    /**
     * Enable/disable automatic backup. Valid values: `enable`, `disable`.
     */
    autoBackup?: pulumi.Input<string>;
    /**
     * ID.
     */
    fosid?: pulumi.Input<string>;
    /**
     * IP address.
     */
    ip?: pulumi.Input<string>;
    /**
     * Enable/disable IPsec. Valid values: `enable`, `disable`.
     */
    ipsec?: pulumi.Input<string>;
    /**
     * Enable/disable scheduled configuration restore. Valid values: `enable`, `disable`.
     */
    scheduledConfigRestore?: pulumi.Input<string>;
    /**
     * Enable/disable FM. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * VDOM.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
