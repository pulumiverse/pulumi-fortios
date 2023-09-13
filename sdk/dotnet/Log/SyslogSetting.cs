// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log
{
    /// <summary>
    /// Provides a resource to configure logging to remote Syslog logging servers.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.logsyslogd.Setting`, we recommend that you use the new resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test2 = new Fortios.Log.SyslogSetting("test2", new()
    ///     {
    ///         Facility = "local7",
    ///         Format = "csv",
    ///         Mode = "udp",
    ///         Port = "514",
    ///         Server = "2.2.2.2",
    ///         SourceIp = "10.2.2.199",
    ///         Status = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:log/syslogSetting:SyslogSetting")]
    public partial class SyslogSetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Remote syslog facility.
        /// </summary>
        [Output("facility")]
        public Output<string> Facility { get; private set; } = null!;

        /// <summary>
        /// Log format.
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// Remote syslog logging over UDP/Reliable TCP.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Server listen port.
        /// </summary>
        [Output("port")]
        public Output<string> Port { get; private set; } = null!;

        /// <summary>
        /// Address of remote syslog server.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Source IP address of syslog.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable remote syslog logging.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SyslogSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyslogSetting(string name, SyslogSettingArgs args, CustomResourceOptions? options = null)
            : base("fortios:log/syslogSetting:SyslogSetting", name, args ?? new SyslogSettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyslogSetting(string name, Input<string> id, SyslogSettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/syslogSetting:SyslogSetting", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SyslogSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyslogSetting Get(string name, Input<string> id, SyslogSettingState? state = null, CustomResourceOptions? options = null)
        {
            return new SyslogSetting(name, id, state, options);
        }
    }

    public sealed class SyslogSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Remote syslog facility.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Log format.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Remote syslog logging over UDP/Reliable TCP.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Server listen port.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Address of remote syslog server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Source IP address of syslog.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable remote syslog logging.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public SyslogSettingArgs()
        {
        }
        public static new SyslogSettingArgs Empty => new SyslogSettingArgs();
    }

    public sealed class SyslogSettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Remote syslog facility.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Log format.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Remote syslog logging over UDP/Reliable TCP.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Server listen port.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Address of remote syslog server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Source IP address of syslog.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable remote syslog logging.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SyslogSettingState()
        {
        }
        public static new SyslogSettingState Empty => new SyslogSettingState();
    }
}
