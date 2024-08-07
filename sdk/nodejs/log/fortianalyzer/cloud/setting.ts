// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../../types/input";
import * as outputs from "../../../types/output";
import * as utilities from "../../../utilities";

/**
 * Global FortiAnalyzer Cloud settings. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * LogFortianalyzerCloud Setting can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:log/fortianalyzer/cloud/setting:Setting labelname LogFortianalyzerCloudSetting
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:log/fortianalyzer/cloud/setting:Setting labelname LogFortianalyzerCloudSetting
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
    public static readonly __pulumiType = 'fortios:log/fortianalyzer/cloud/setting:Setting';

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
     * Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
     */
    public readonly accessConfig!: pulumi.Output<string>;
    /**
     * Certificate used to communicate with FortiAnalyzer.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
     */
    public readonly certificateVerification!: pulumi.Output<string>;
    /**
     * FortiAnalyzer connection time-out in seconds (for status and log buffer).
     */
    public readonly connTimeout!: pulumi.Output<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Configure the level of SSL protection for secure communication with FortiAnalyzer. Valid values: `high-medium`, `high`, `low`.
     */
    public readonly encAlgorithm!: pulumi.Output<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm.
     */
    public readonly hmacAlgorithm!: pulumi.Output<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
     */
    public readonly ipsArchive!: pulumi.Output<string>;
    /**
     * FortiAnalyzer maximum log rate in MBps (0 = unlimited).
     */
    public readonly maxLogRate!: pulumi.Output<number>;
    /**
     * Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
     */
    public readonly monitorFailureRetryPeriod!: pulumi.Output<number>;
    /**
     * Time between OFTP keepalives in seconds (for status and log buffer).
     */
    public readonly monitorKeepalivePeriod!: pulumi.Output<number>;
    /**
     * Preshared-key used for auto-authorization on FortiAnalyzer.
     */
    public readonly presharedKey!: pulumi.Output<string>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    public readonly priority!: pulumi.Output<string>;
    /**
     * Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
     */
    public readonly serials!: pulumi.Output<outputs.log.fortianalyzer.cloud.SettingSerial[] | undefined>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Day of week (month) to upload logs.
     */
    public readonly uploadDay!: pulumi.Output<string>;
    /**
     * Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
     */
    public readonly uploadInterval!: pulumi.Output<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
     */
    public readonly uploadOption!: pulumi.Output<string>;
    /**
     * Time to upload logs (hh:mm).
     */
    public readonly uploadTime!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;

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
            resourceInputs["accessConfig"] = state ? state.accessConfig : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["certificateVerification"] = state ? state.certificateVerification : undefined;
            resourceInputs["connTimeout"] = state ? state.connTimeout : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = state ? state.encAlgorithm : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["hmacAlgorithm"] = state ? state.hmacAlgorithm : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["ipsArchive"] = state ? state.ipsArchive : undefined;
            resourceInputs["maxLogRate"] = state ? state.maxLogRate : undefined;
            resourceInputs["monitorFailureRetryPeriod"] = state ? state.monitorFailureRetryPeriod : undefined;
            resourceInputs["monitorKeepalivePeriod"] = state ? state.monitorKeepalivePeriod : undefined;
            resourceInputs["presharedKey"] = state ? state.presharedKey : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["serials"] = state ? state.serials : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uploadDay"] = state ? state.uploadDay : undefined;
            resourceInputs["uploadInterval"] = state ? state.uploadInterval : undefined;
            resourceInputs["uploadOption"] = state ? state.uploadOption : undefined;
            resourceInputs["uploadTime"] = state ? state.uploadTime : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SettingArgs | undefined;
            resourceInputs["accessConfig"] = args ? args.accessConfig : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["certificateVerification"] = args ? args.certificateVerification : undefined;
            resourceInputs["connTimeout"] = args ? args.connTimeout : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["encAlgorithm"] = args ? args.encAlgorithm : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["hmacAlgorithm"] = args ? args.hmacAlgorithm : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["ipsArchive"] = args ? args.ipsArchive : undefined;
            resourceInputs["maxLogRate"] = args ? args.maxLogRate : undefined;
            resourceInputs["monitorFailureRetryPeriod"] = args ? args.monitorFailureRetryPeriod : undefined;
            resourceInputs["monitorKeepalivePeriod"] = args ? args.monitorKeepalivePeriod : undefined;
            resourceInputs["presharedKey"] = args ? args.presharedKey : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["serials"] = args ? args.serials : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["uploadDay"] = args ? args.uploadDay : undefined;
            resourceInputs["uploadInterval"] = args ? args.uploadInterval : undefined;
            resourceInputs["uploadOption"] = args ? args.uploadOption : undefined;
            resourceInputs["uploadTime"] = args ? args.uploadTime : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
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
     * Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
     */
    accessConfig?: pulumi.Input<string>;
    /**
     * Certificate used to communicate with FortiAnalyzer.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
     */
    certificateVerification?: pulumi.Input<string>;
    /**
     * FortiAnalyzer connection time-out in seconds (for status and log buffer).
     */
    connTimeout?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure the level of SSL protection for secure communication with FortiAnalyzer. Valid values: `high-medium`, `high`, `low`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm.
     */
    hmacAlgorithm?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
     */
    ipsArchive?: pulumi.Input<string>;
    /**
     * FortiAnalyzer maximum log rate in MBps (0 = unlimited).
     */
    maxLogRate?: pulumi.Input<number>;
    /**
     * Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
     */
    monitorFailureRetryPeriod?: pulumi.Input<number>;
    /**
     * Time between OFTP keepalives in seconds (for status and log buffer).
     */
    monitorKeepalivePeriod?: pulumi.Input<number>;
    /**
     * Preshared-key used for auto-authorization on FortiAnalyzer.
     */
    presharedKey?: pulumi.Input<string>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
     */
    serials?: pulumi.Input<pulumi.Input<inputs.log.fortianalyzer.cloud.SettingSerial>[]>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Day of week (month) to upload logs.
     */
    uploadDay?: pulumi.Input<string>;
    /**
     * Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
     */
    uploadInterval?: pulumi.Input<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
     */
    uploadOption?: pulumi.Input<string>;
    /**
     * Time to upload logs (hh:mm).
     */
    uploadTime?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Setting resource.
 */
export interface SettingArgs {
    /**
     * Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
     */
    accessConfig?: pulumi.Input<string>;
    /**
     * Certificate used to communicate with FortiAnalyzer.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
     */
    certificateVerification?: pulumi.Input<string>;
    /**
     * FortiAnalyzer connection time-out in seconds (for status and log buffer).
     */
    connTimeout?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure the level of SSL protection for secure communication with FortiAnalyzer. Valid values: `high-medium`, `high`, `low`.
     */
    encAlgorithm?: pulumi.Input<string>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * FortiAnalyzer IPsec tunnel HMAC algorithm.
     */
    hmacAlgorithm?: pulumi.Input<string>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
     */
    ipsArchive?: pulumi.Input<string>;
    /**
     * FortiAnalyzer maximum log rate in MBps (0 = unlimited).
     */
    maxLogRate?: pulumi.Input<number>;
    /**
     * Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
     */
    monitorFailureRetryPeriod?: pulumi.Input<number>;
    /**
     * Time between OFTP keepalives in seconds (for status and log buffer).
     */
    monitorKeepalivePeriod?: pulumi.Input<number>;
    /**
     * Preshared-key used for auto-authorization on FortiAnalyzer.
     */
    presharedKey?: pulumi.Input<string>;
    /**
     * Set log transmission priority. Valid values: `default`, `low`.
     */
    priority?: pulumi.Input<string>;
    /**
     * Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
     */
    serials?: pulumi.Input<pulumi.Input<inputs.log.fortianalyzer.cloud.SettingSerial>[]>;
    /**
     * Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Day of week (month) to upload logs.
     */
    uploadDay?: pulumi.Input<string>;
    /**
     * Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
     */
    uploadInterval?: pulumi.Input<string>;
    /**
     * Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
     */
    uploadOption?: pulumi.Input<string>;
    /**
     * Time to upload logs (hh:mm).
     */
    uploadTime?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
