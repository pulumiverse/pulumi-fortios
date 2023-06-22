// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Settings for local disk logging.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.LogdiskSetting("trname", {
 *     diskfull: "overwrite",
 *     dlpArchiveQuota: 0,
 *     fullFinalWarningThreshold: 95,
 *     fullFirstWarningThreshold: 75,
 *     fullSecondWarningThreshold: 90,
 *     ipsArchive: "enable",
 *     logQuota: 0,
 *     maxLogFileSize: 20,
 *     maxPolicyPacketCaptureSize: 100,
 *     maximumLogAge: 7,
 *     reportQuota: 0,
 *     rollDay: "sunday",
 *     rollSchedule: "daily",
 *     rollTime: "00:00",
 *     sourceIp: "0.0.0.0",
 *     status: "enable",
 *     upload: "disable",
 *     uploadDeleteFiles: "enable",
 *     uploadDestination: "ftp-server",
 *     uploadSslConn: "default",
 *     uploadip: "0.0.0.0",
 *     uploadport: 21,
 *     uploadsched: "disable",
 *     uploadtime: "00:00",
 *     uploadtype: "traffic event virus webfilter IPS spamfilter dlp-archive anomaly voip dlp app-ctrl waf netscan gtp dns",
 * });
 * ```
 *
 * ## Import
 *
 * LogDisk Setting can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/logdiskSetting:LogdiskSetting labelname LogDiskSetting
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/logdiskSetting:LogdiskSetting labelname LogDiskSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class LogdiskSetting extends pulumi.CustomResource {
    /**
     * Get an existing LogdiskSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogdiskSettingState, opts?: pulumi.CustomResourceOptions): LogdiskSetting {
        return new LogdiskSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/logdiskSetting:LogdiskSetting';

    /**
     * Returns true if the given object is an instance of LogdiskSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogdiskSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogdiskSetting.__pulumiType;
    }

    /**
     * Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
     */
    public readonly diskfull!: pulumi.Output<string>;
    /**
     * DLP archive quota (MB).
     */
    public readonly dlpArchiveQuota!: pulumi.Output<number>;
    /**
     * Log full final warning threshold as a percent (3 - 100, default = 95).
     */
    public readonly fullFinalWarningThreshold!: pulumi.Output<number>;
    /**
     * Log full first warning threshold as a percent (1 - 98, default = 75).
     */
    public readonly fullFirstWarningThreshold!: pulumi.Output<number>;
    /**
     * Log full second warning threshold as a percent (2 - 99, default = 90).
     */
    public readonly fullSecondWarningThreshold!: pulumi.Output<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    /**
     * Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
     */
    public readonly ipsArchive!: pulumi.Output<string>;
    /**
     * Disk log quota (MB).
     */
    public readonly logQuota!: pulumi.Output<number>;
    /**
     * Maximum log file size before rolling (1 - 100 Mbytes).
     */
    public readonly maxLogFileSize!: pulumi.Output<number>;
    /**
     * Maximum size of policy sniffer in MB (0 means unlimited).
     */
    public readonly maxPolicyPacketCaptureSize!: pulumi.Output<number>;
    /**
     * Delete log files older than (days).
     */
    public readonly maximumLogAge!: pulumi.Output<number>;
    /**
     * Report quota (MB).
     */
    public readonly reportQuota!: pulumi.Output<number>;
    /**
     * Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    public readonly rollDay!: pulumi.Output<string>;
    /**
     * Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
     */
    public readonly rollSchedule!: pulumi.Output<string>;
    /**
     * Time of day to roll the log file (hh:mm).
     */
    public readonly rollTime!: pulumi.Output<string>;
    /**
     * Source IP address to use for uploading disk log files.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Enable/disable local disk logging. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
     */
    public readonly upload!: pulumi.Output<string>;
    /**
     * Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
     */
    public readonly uploadDeleteFiles!: pulumi.Output<string>;
    /**
     * The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
     */
    public readonly uploadDestination!: pulumi.Output<string>;
    /**
     * Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
     */
    public readonly uploadSslConn!: pulumi.Output<string>;
    /**
     * The remote directory on the FTP server to upload log files to.
     */
    public readonly uploaddir!: pulumi.Output<string>;
    /**
     * IP address of the FTP server to upload log files to.
     */
    public readonly uploadip!: pulumi.Output<string>;
    /**
     * Password required to log into the FTP server to upload disk log files.
     */
    public readonly uploadpass!: pulumi.Output<string | undefined>;
    /**
     * TCP port to use for communicating with the FTP server (default = 21).
     */
    public readonly uploadport!: pulumi.Output<number>;
    /**
     * Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
     */
    public readonly uploadsched!: pulumi.Output<string>;
    /**
     * Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
     */
    public readonly uploadtime!: pulumi.Output<string>;
    /**
     * Types of log files to upload. Separate multiple entries with a space.
     */
    public readonly uploadtype!: pulumi.Output<string>;
    /**
     * Username required to log into the FTP server to upload disk log files.
     */
    public readonly uploaduser!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a LogdiskSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogdiskSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogdiskSettingArgs | LogdiskSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogdiskSettingState | undefined;
            resourceInputs["diskfull"] = state ? state.diskfull : undefined;
            resourceInputs["dlpArchiveQuota"] = state ? state.dlpArchiveQuota : undefined;
            resourceInputs["fullFinalWarningThreshold"] = state ? state.fullFinalWarningThreshold : undefined;
            resourceInputs["fullFirstWarningThreshold"] = state ? state.fullFirstWarningThreshold : undefined;
            resourceInputs["fullSecondWarningThreshold"] = state ? state.fullSecondWarningThreshold : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["ipsArchive"] = state ? state.ipsArchive : undefined;
            resourceInputs["logQuota"] = state ? state.logQuota : undefined;
            resourceInputs["maxLogFileSize"] = state ? state.maxLogFileSize : undefined;
            resourceInputs["maxPolicyPacketCaptureSize"] = state ? state.maxPolicyPacketCaptureSize : undefined;
            resourceInputs["maximumLogAge"] = state ? state.maximumLogAge : undefined;
            resourceInputs["reportQuota"] = state ? state.reportQuota : undefined;
            resourceInputs["rollDay"] = state ? state.rollDay : undefined;
            resourceInputs["rollSchedule"] = state ? state.rollSchedule : undefined;
            resourceInputs["rollTime"] = state ? state.rollTime : undefined;
            resourceInputs["sourceIp"] = state ? state.sourceIp : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["upload"] = state ? state.upload : undefined;
            resourceInputs["uploadDeleteFiles"] = state ? state.uploadDeleteFiles : undefined;
            resourceInputs["uploadDestination"] = state ? state.uploadDestination : undefined;
            resourceInputs["uploadSslConn"] = state ? state.uploadSslConn : undefined;
            resourceInputs["uploaddir"] = state ? state.uploaddir : undefined;
            resourceInputs["uploadip"] = state ? state.uploadip : undefined;
            resourceInputs["uploadpass"] = state ? state.uploadpass : undefined;
            resourceInputs["uploadport"] = state ? state.uploadport : undefined;
            resourceInputs["uploadsched"] = state ? state.uploadsched : undefined;
            resourceInputs["uploadtime"] = state ? state.uploadtime : undefined;
            resourceInputs["uploadtype"] = state ? state.uploadtype : undefined;
            resourceInputs["uploaduser"] = state ? state.uploaduser : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LogdiskSettingArgs | undefined;
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["diskfull"] = args ? args.diskfull : undefined;
            resourceInputs["dlpArchiveQuota"] = args ? args.dlpArchiveQuota : undefined;
            resourceInputs["fullFinalWarningThreshold"] = args ? args.fullFinalWarningThreshold : undefined;
            resourceInputs["fullFirstWarningThreshold"] = args ? args.fullFirstWarningThreshold : undefined;
            resourceInputs["fullSecondWarningThreshold"] = args ? args.fullSecondWarningThreshold : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["ipsArchive"] = args ? args.ipsArchive : undefined;
            resourceInputs["logQuota"] = args ? args.logQuota : undefined;
            resourceInputs["maxLogFileSize"] = args ? args.maxLogFileSize : undefined;
            resourceInputs["maxPolicyPacketCaptureSize"] = args ? args.maxPolicyPacketCaptureSize : undefined;
            resourceInputs["maximumLogAge"] = args ? args.maximumLogAge : undefined;
            resourceInputs["reportQuota"] = args ? args.reportQuota : undefined;
            resourceInputs["rollDay"] = args ? args.rollDay : undefined;
            resourceInputs["rollSchedule"] = args ? args.rollSchedule : undefined;
            resourceInputs["rollTime"] = args ? args.rollTime : undefined;
            resourceInputs["sourceIp"] = args ? args.sourceIp : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["upload"] = args ? args.upload : undefined;
            resourceInputs["uploadDeleteFiles"] = args ? args.uploadDeleteFiles : undefined;
            resourceInputs["uploadDestination"] = args ? args.uploadDestination : undefined;
            resourceInputs["uploadSslConn"] = args ? args.uploadSslConn : undefined;
            resourceInputs["uploaddir"] = args ? args.uploaddir : undefined;
            resourceInputs["uploadip"] = args ? args.uploadip : undefined;
            resourceInputs["uploadpass"] = args?.uploadpass ? pulumi.secret(args.uploadpass) : undefined;
            resourceInputs["uploadport"] = args ? args.uploadport : undefined;
            resourceInputs["uploadsched"] = args ? args.uploadsched : undefined;
            resourceInputs["uploadtime"] = args ? args.uploadtime : undefined;
            resourceInputs["uploadtype"] = args ? args.uploadtype : undefined;
            resourceInputs["uploaduser"] = args ? args.uploaduser : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["uploadpass"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(LogdiskSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogdiskSetting resources.
 */
export interface LogdiskSettingState {
    /**
     * Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
     */
    diskfull?: pulumi.Input<string>;
    /**
     * DLP archive quota (MB).
     */
    dlpArchiveQuota?: pulumi.Input<number>;
    /**
     * Log full final warning threshold as a percent (3 - 100, default = 95).
     */
    fullFinalWarningThreshold?: pulumi.Input<number>;
    /**
     * Log full first warning threshold as a percent (1 - 98, default = 75).
     */
    fullFirstWarningThreshold?: pulumi.Input<number>;
    /**
     * Log full second warning threshold as a percent (2 - 99, default = 90).
     */
    fullSecondWarningThreshold?: pulumi.Input<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
     */
    ipsArchive?: pulumi.Input<string>;
    /**
     * Disk log quota (MB).
     */
    logQuota?: pulumi.Input<number>;
    /**
     * Maximum log file size before rolling (1 - 100 Mbytes).
     */
    maxLogFileSize?: pulumi.Input<number>;
    /**
     * Maximum size of policy sniffer in MB (0 means unlimited).
     */
    maxPolicyPacketCaptureSize?: pulumi.Input<number>;
    /**
     * Delete log files older than (days).
     */
    maximumLogAge?: pulumi.Input<number>;
    /**
     * Report quota (MB).
     */
    reportQuota?: pulumi.Input<number>;
    /**
     * Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    rollDay?: pulumi.Input<string>;
    /**
     * Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
     */
    rollSchedule?: pulumi.Input<string>;
    /**
     * Time of day to roll the log file (hh:mm).
     */
    rollTime?: pulumi.Input<string>;
    /**
     * Source IP address to use for uploading disk log files.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable local disk logging. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
     */
    upload?: pulumi.Input<string>;
    /**
     * Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
     */
    uploadDeleteFiles?: pulumi.Input<string>;
    /**
     * The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
     */
    uploadDestination?: pulumi.Input<string>;
    /**
     * Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
     */
    uploadSslConn?: pulumi.Input<string>;
    /**
     * The remote directory on the FTP server to upload log files to.
     */
    uploaddir?: pulumi.Input<string>;
    /**
     * IP address of the FTP server to upload log files to.
     */
    uploadip?: pulumi.Input<string>;
    /**
     * Password required to log into the FTP server to upload disk log files.
     */
    uploadpass?: pulumi.Input<string>;
    /**
     * TCP port to use for communicating with the FTP server (default = 21).
     */
    uploadport?: pulumi.Input<number>;
    /**
     * Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
     */
    uploadsched?: pulumi.Input<string>;
    /**
     * Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
     */
    uploadtime?: pulumi.Input<string>;
    /**
     * Types of log files to upload. Separate multiple entries with a space.
     */
    uploadtype?: pulumi.Input<string>;
    /**
     * Username required to log into the FTP server to upload disk log files.
     */
    uploaduser?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LogdiskSetting resource.
 */
export interface LogdiskSettingArgs {
    /**
     * Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
     */
    diskfull?: pulumi.Input<string>;
    /**
     * DLP archive quota (MB).
     */
    dlpArchiveQuota?: pulumi.Input<number>;
    /**
     * Log full final warning threshold as a percent (3 - 100, default = 95).
     */
    fullFinalWarningThreshold?: pulumi.Input<number>;
    /**
     * Log full first warning threshold as a percent (1 - 98, default = 75).
     */
    fullFirstWarningThreshold?: pulumi.Input<number>;
    /**
     * Log full second warning threshold as a percent (2 - 99, default = 90).
     */
    fullSecondWarningThreshold?: pulumi.Input<number>;
    /**
     * Specify outgoing interface to reach server.
     */
    interface?: pulumi.Input<string>;
    /**
     * Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
     */
    interfaceSelectMethod?: pulumi.Input<string>;
    /**
     * Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
     */
    ipsArchive?: pulumi.Input<string>;
    /**
     * Disk log quota (MB).
     */
    logQuota?: pulumi.Input<number>;
    /**
     * Maximum log file size before rolling (1 - 100 Mbytes).
     */
    maxLogFileSize?: pulumi.Input<number>;
    /**
     * Maximum size of policy sniffer in MB (0 means unlimited).
     */
    maxPolicyPacketCaptureSize?: pulumi.Input<number>;
    /**
     * Delete log files older than (days).
     */
    maximumLogAge?: pulumi.Input<number>;
    /**
     * Report quota (MB).
     */
    reportQuota?: pulumi.Input<number>;
    /**
     * Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    rollDay?: pulumi.Input<string>;
    /**
     * Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
     */
    rollSchedule?: pulumi.Input<string>;
    /**
     * Time of day to roll the log file (hh:mm).
     */
    rollTime?: pulumi.Input<string>;
    /**
     * Source IP address to use for uploading disk log files.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable local disk logging. Valid values: `enable`, `disable`.
     */
    status: pulumi.Input<string>;
    /**
     * Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
     */
    upload?: pulumi.Input<string>;
    /**
     * Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
     */
    uploadDeleteFiles?: pulumi.Input<string>;
    /**
     * The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
     */
    uploadDestination?: pulumi.Input<string>;
    /**
     * Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
     */
    uploadSslConn?: pulumi.Input<string>;
    /**
     * The remote directory on the FTP server to upload log files to.
     */
    uploaddir?: pulumi.Input<string>;
    /**
     * IP address of the FTP server to upload log files to.
     */
    uploadip?: pulumi.Input<string>;
    /**
     * Password required to log into the FTP server to upload disk log files.
     */
    uploadpass?: pulumi.Input<string>;
    /**
     * TCP port to use for communicating with the FTP server (default = 21).
     */
    uploadport?: pulumi.Input<number>;
    /**
     * Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
     */
    uploadsched?: pulumi.Input<string>;
    /**
     * Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
     */
    uploadtime?: pulumi.Input<string>;
    /**
     * Types of log files to upload. Separate multiple entries with a space.
     */
    uploadtype?: pulumi.Input<string>;
    /**
     * Username required to log into the FTP server to upload disk log files.
     */
    uploaduser?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
