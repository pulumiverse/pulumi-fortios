// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log.Disk
{
    /// <summary>
    /// Settings for local disk logging.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Log.Disk.Setting("trname", new()
    ///     {
    ///         Diskfull = "overwrite",
    ///         DlpArchiveQuota = 0,
    ///         FullFinalWarningThreshold = 95,
    ///         FullFirstWarningThreshold = 75,
    ///         FullSecondWarningThreshold = 90,
    ///         IpsArchive = "enable",
    ///         LogQuota = 0,
    ///         MaxLogFileSize = 20,
    ///         MaxPolicyPacketCaptureSize = 100,
    ///         MaximumLogAge = 7,
    ///         ReportQuota = 0,
    ///         RollDay = "sunday",
    ///         RollSchedule = "daily",
    ///         RollTime = "00:00",
    ///         SourceIp = "0.0.0.0",
    ///         Status = "enable",
    ///         Upload = "disable",
    ///         UploadDeleteFiles = "enable",
    ///         UploadDestination = "ftp-server",
    ///         UploadSslConn = "default",
    ///         Uploadip = "0.0.0.0",
    ///         Uploadport = 21,
    ///         Uploadsched = "disable",
    ///         Uploadtime = "00:00",
    ///         Uploadtype = "traffic event virus webfilter IPS spamfilter dlp-archive anomaly voip dlp app-ctrl waf netscan gtp dns",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// LogDisk Setting can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/disk/setting:Setting labelname LogDiskSetting
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/disk/setting:Setting labelname LogDiskSetting
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:log/disk/setting:Setting")]
    public partial class Setting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
        /// </summary>
        [Output("diskfull")]
        public Output<string> Diskfull { get; private set; } = null!;

        /// <summary>
        /// DLP archive quota (MB).
        /// </summary>
        [Output("dlpArchiveQuota")]
        public Output<int> DlpArchiveQuota { get; private set; } = null!;

        /// <summary>
        /// Log full final warning threshold as a percent (3 - 100, default = 95).
        /// </summary>
        [Output("fullFinalWarningThreshold")]
        public Output<int> FullFinalWarningThreshold { get; private set; } = null!;

        /// <summary>
        /// Log full first warning threshold as a percent (1 - 98, default = 75).
        /// </summary>
        [Output("fullFirstWarningThreshold")]
        public Output<int> FullFirstWarningThreshold { get; private set; } = null!;

        /// <summary>
        /// Log full second warning threshold as a percent (2 - 99, default = 90).
        /// </summary>
        [Output("fullSecondWarningThreshold")]
        public Output<int> FullSecondWarningThreshold { get; private set; } = null!;

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsArchive")]
        public Output<string> IpsArchive { get; private set; } = null!;

        /// <summary>
        /// Disk log quota (MB).
        /// </summary>
        [Output("logQuota")]
        public Output<int> LogQuota { get; private set; } = null!;

        /// <summary>
        /// Maximum log file size before rolling (1 - 100 Mbytes).
        /// </summary>
        [Output("maxLogFileSize")]
        public Output<int> MaxLogFileSize { get; private set; } = null!;

        /// <summary>
        /// Maximum size of policy sniffer in MB (0 means unlimited).
        /// </summary>
        [Output("maxPolicyPacketCaptureSize")]
        public Output<int> MaxPolicyPacketCaptureSize { get; private set; } = null!;

        /// <summary>
        /// Delete log files older than (days).
        /// </summary>
        [Output("maximumLogAge")]
        public Output<int> MaximumLogAge { get; private set; } = null!;

        /// <summary>
        /// Report quota (MB).
        /// </summary>
        [Output("reportQuota")]
        public Output<int> ReportQuota { get; private set; } = null!;

        /// <summary>
        /// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
        /// </summary>
        [Output("rollDay")]
        public Output<string> RollDay { get; private set; } = null!;

        /// <summary>
        /// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
        /// </summary>
        [Output("rollSchedule")]
        public Output<string> RollSchedule { get; private set; } = null!;

        /// <summary>
        /// Time of day to roll the log file (hh:mm).
        /// </summary>
        [Output("rollTime")]
        public Output<string> RollTime { get; private set; } = null!;

        /// <summary>
        /// Source IP address to use for uploading disk log files.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable local disk logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("upload")]
        public Output<string> Upload { get; private set; } = null!;

        /// <summary>
        /// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("uploadDeleteFiles")]
        public Output<string> UploadDeleteFiles { get; private set; } = null!;

        /// <summary>
        /// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
        /// </summary>
        [Output("uploadDestination")]
        public Output<string> UploadDestination { get; private set; } = null!;

        /// <summary>
        /// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
        /// </summary>
        [Output("uploadSslConn")]
        public Output<string> UploadSslConn { get; private set; } = null!;

        /// <summary>
        /// The remote directory on the FTP server to upload log files to.
        /// </summary>
        [Output("uploaddir")]
        public Output<string> Uploaddir { get; private set; } = null!;

        /// <summary>
        /// IP address of the FTP server to upload log files to.
        /// </summary>
        [Output("uploadip")]
        public Output<string> Uploadip { get; private set; } = null!;

        /// <summary>
        /// Password required to log into the FTP server to upload disk log files.
        /// </summary>
        [Output("uploadpass")]
        public Output<string?> Uploadpass { get; private set; } = null!;

        /// <summary>
        /// TCP port to use for communicating with the FTP server (default = 21).
        /// </summary>
        [Output("uploadport")]
        public Output<int> Uploadport { get; private set; } = null!;

        /// <summary>
        /// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
        /// </summary>
        [Output("uploadsched")]
        public Output<string> Uploadsched { get; private set; } = null!;

        /// <summary>
        /// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
        /// </summary>
        [Output("uploadtime")]
        public Output<string> Uploadtime { get; private set; } = null!;

        /// <summary>
        /// Types of log files to upload. Separate multiple entries with a space.
        /// </summary>
        [Output("uploadtype")]
        public Output<string> Uploadtype { get; private set; } = null!;

        /// <summary>
        /// Username required to log into the FTP server to upload disk log files.
        /// </summary>
        [Output("uploaduser")]
        public Output<string> Uploaduser { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Setting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Setting(string name, SettingArgs args, CustomResourceOptions? options = null)
            : base("fortios:log/disk/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/disk/setting:Setting", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "uploadpass",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Setting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Setting Get(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Setting(name, id, state, options);
        }
    }

    public sealed class SettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
        /// </summary>
        [Input("diskfull")]
        public Input<string>? Diskfull { get; set; }

        /// <summary>
        /// DLP archive quota (MB).
        /// </summary>
        [Input("dlpArchiveQuota")]
        public Input<int>? DlpArchiveQuota { get; set; }

        /// <summary>
        /// Log full final warning threshold as a percent (3 - 100, default = 95).
        /// </summary>
        [Input("fullFinalWarningThreshold")]
        public Input<int>? FullFinalWarningThreshold { get; set; }

        /// <summary>
        /// Log full first warning threshold as a percent (1 - 98, default = 75).
        /// </summary>
        [Input("fullFirstWarningThreshold")]
        public Input<int>? FullFirstWarningThreshold { get; set; }

        /// <summary>
        /// Log full second warning threshold as a percent (2 - 99, default = 90).
        /// </summary>
        [Input("fullSecondWarningThreshold")]
        public Input<int>? FullSecondWarningThreshold { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsArchive")]
        public Input<string>? IpsArchive { get; set; }

        /// <summary>
        /// Disk log quota (MB).
        /// </summary>
        [Input("logQuota")]
        public Input<int>? LogQuota { get; set; }

        /// <summary>
        /// Maximum log file size before rolling (1 - 100 Mbytes).
        /// </summary>
        [Input("maxLogFileSize")]
        public Input<int>? MaxLogFileSize { get; set; }

        /// <summary>
        /// Maximum size of policy sniffer in MB (0 means unlimited).
        /// </summary>
        [Input("maxPolicyPacketCaptureSize")]
        public Input<int>? MaxPolicyPacketCaptureSize { get; set; }

        /// <summary>
        /// Delete log files older than (days).
        /// </summary>
        [Input("maximumLogAge")]
        public Input<int>? MaximumLogAge { get; set; }

        /// <summary>
        /// Report quota (MB).
        /// </summary>
        [Input("reportQuota")]
        public Input<int>? ReportQuota { get; set; }

        /// <summary>
        /// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
        /// </summary>
        [Input("rollDay")]
        public Input<string>? RollDay { get; set; }

        /// <summary>
        /// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
        /// </summary>
        [Input("rollSchedule")]
        public Input<string>? RollSchedule { get; set; }

        /// <summary>
        /// Time of day to roll the log file (hh:mm).
        /// </summary>
        [Input("rollTime")]
        public Input<string>? RollTime { get; set; }

        /// <summary>
        /// Source IP address to use for uploading disk log files.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable local disk logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("upload")]
        public Input<string>? Upload { get; set; }

        /// <summary>
        /// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("uploadDeleteFiles")]
        public Input<string>? UploadDeleteFiles { get; set; }

        /// <summary>
        /// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
        /// </summary>
        [Input("uploadDestination")]
        public Input<string>? UploadDestination { get; set; }

        /// <summary>
        /// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
        /// </summary>
        [Input("uploadSslConn")]
        public Input<string>? UploadSslConn { get; set; }

        /// <summary>
        /// The remote directory on the FTP server to upload log files to.
        /// </summary>
        [Input("uploaddir")]
        public Input<string>? Uploaddir { get; set; }

        /// <summary>
        /// IP address of the FTP server to upload log files to.
        /// </summary>
        [Input("uploadip")]
        public Input<string>? Uploadip { get; set; }

        [Input("uploadpass")]
        private Input<string>? _uploadpass;

        /// <summary>
        /// Password required to log into the FTP server to upload disk log files.
        /// </summary>
        public Input<string>? Uploadpass
        {
            get => _uploadpass;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _uploadpass = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// TCP port to use for communicating with the FTP server (default = 21).
        /// </summary>
        [Input("uploadport")]
        public Input<int>? Uploadport { get; set; }

        /// <summary>
        /// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("uploadsched")]
        public Input<string>? Uploadsched { get; set; }

        /// <summary>
        /// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
        /// </summary>
        [Input("uploadtime")]
        public Input<string>? Uploadtime { get; set; }

        /// <summary>
        /// Types of log files to upload. Separate multiple entries with a space.
        /// </summary>
        [Input("uploadtype")]
        public Input<string>? Uploadtype { get; set; }

        /// <summary>
        /// Username required to log into the FTP server to upload disk log files.
        /// </summary>
        [Input("uploaduser")]
        public Input<string>? Uploaduser { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingArgs()
        {
        }
        public static new SettingArgs Empty => new SettingArgs();
    }

    public sealed class SettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
        /// </summary>
        [Input("diskfull")]
        public Input<string>? Diskfull { get; set; }

        /// <summary>
        /// DLP archive quota (MB).
        /// </summary>
        [Input("dlpArchiveQuota")]
        public Input<int>? DlpArchiveQuota { get; set; }

        /// <summary>
        /// Log full final warning threshold as a percent (3 - 100, default = 95).
        /// </summary>
        [Input("fullFinalWarningThreshold")]
        public Input<int>? FullFinalWarningThreshold { get; set; }

        /// <summary>
        /// Log full first warning threshold as a percent (1 - 98, default = 75).
        /// </summary>
        [Input("fullFirstWarningThreshold")]
        public Input<int>? FullFirstWarningThreshold { get; set; }

        /// <summary>
        /// Log full second warning threshold as a percent (2 - 99, default = 90).
        /// </summary>
        [Input("fullSecondWarningThreshold")]
        public Input<int>? FullSecondWarningThreshold { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsArchive")]
        public Input<string>? IpsArchive { get; set; }

        /// <summary>
        /// Disk log quota (MB).
        /// </summary>
        [Input("logQuota")]
        public Input<int>? LogQuota { get; set; }

        /// <summary>
        /// Maximum log file size before rolling (1 - 100 Mbytes).
        /// </summary>
        [Input("maxLogFileSize")]
        public Input<int>? MaxLogFileSize { get; set; }

        /// <summary>
        /// Maximum size of policy sniffer in MB (0 means unlimited).
        /// </summary>
        [Input("maxPolicyPacketCaptureSize")]
        public Input<int>? MaxPolicyPacketCaptureSize { get; set; }

        /// <summary>
        /// Delete log files older than (days).
        /// </summary>
        [Input("maximumLogAge")]
        public Input<int>? MaximumLogAge { get; set; }

        /// <summary>
        /// Report quota (MB).
        /// </summary>
        [Input("reportQuota")]
        public Input<int>? ReportQuota { get; set; }

        /// <summary>
        /// Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
        /// </summary>
        [Input("rollDay")]
        public Input<string>? RollDay { get; set; }

        /// <summary>
        /// Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
        /// </summary>
        [Input("rollSchedule")]
        public Input<string>? RollSchedule { get; set; }

        /// <summary>
        /// Time of day to roll the log file (hh:mm).
        /// </summary>
        [Input("rollTime")]
        public Input<string>? RollTime { get; set; }

        /// <summary>
        /// Source IP address to use for uploading disk log files.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable local disk logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("upload")]
        public Input<string>? Upload { get; set; }

        /// <summary>
        /// Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("uploadDeleteFiles")]
        public Input<string>? UploadDeleteFiles { get; set; }

        /// <summary>
        /// The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
        /// </summary>
        [Input("uploadDestination")]
        public Input<string>? UploadDestination { get; set; }

        /// <summary>
        /// Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
        /// </summary>
        [Input("uploadSslConn")]
        public Input<string>? UploadSslConn { get; set; }

        /// <summary>
        /// The remote directory on the FTP server to upload log files to.
        /// </summary>
        [Input("uploaddir")]
        public Input<string>? Uploaddir { get; set; }

        /// <summary>
        /// IP address of the FTP server to upload log files to.
        /// </summary>
        [Input("uploadip")]
        public Input<string>? Uploadip { get; set; }

        [Input("uploadpass")]
        private Input<string>? _uploadpass;

        /// <summary>
        /// Password required to log into the FTP server to upload disk log files.
        /// </summary>
        public Input<string>? Uploadpass
        {
            get => _uploadpass;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _uploadpass = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// TCP port to use for communicating with the FTP server (default = 21).
        /// </summary>
        [Input("uploadport")]
        public Input<int>? Uploadport { get; set; }

        /// <summary>
        /// Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("uploadsched")]
        public Input<string>? Uploadsched { get; set; }

        /// <summary>
        /// Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
        /// </summary>
        [Input("uploadtime")]
        public Input<string>? Uploadtime { get; set; }

        /// <summary>
        /// Types of log files to upload. Separate multiple entries with a space.
        /// </summary>
        [Input("uploadtype")]
        public Input<string>? Uploadtype { get; set; }

        /// <summary>
        /// Username required to log into the FTP server to upload disk log files.
        /// </summary>
        [Input("uploaduser")]
        public Input<string>? Uploaduser { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingState()
        {
        }
        public static new SettingState Empty => new SettingState();
    }
}
