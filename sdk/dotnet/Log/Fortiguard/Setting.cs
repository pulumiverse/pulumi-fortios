// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log.Fortiguard
{
    /// <summary>
    /// Configure logging to FortiCloud.
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
    ///     var trname = new Fortios.Log.Fortiguard.Setting("trname", new()
    ///     {
    ///         EncAlgorithm = "high",
    ///         SourceIp = "0.0.0.0",
    ///         SslMinProtoVersion = "default",
    ///         Status = "disable",
    ///         UploadInterval = "daily",
    ///         UploadOption = "5-minute",
    ///         UploadTime = "00:00",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LogFortiguard Setting can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/fortiguard/setting:Setting labelname LogFortiguardSetting
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/fortiguard/setting:Setting labelname LogFortiguardSetting
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:log/fortiguard/setting:Setting")]
    public partial class Setting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("accessConfig")]
        public Output<string> AccessConfig { get; private set; } = null!;

        /// <summary>
        /// FortiGate Cloud connection timeout in seconds.
        /// </summary>
        [Output("connTimeout")]
        public Output<int> ConnTimeout { get; private set; } = null!;

        /// <summary>
        /// Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
        /// </summary>
        [Output("encAlgorithm")]
        public Output<string> EncAlgorithm { get; private set; } = null!;

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
        /// FortiCloud maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Output("maxLogRate")]
        public Output<int> MaxLogRate { get; private set; } = null!;

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        /// <summary>
        /// Source IP address used to connect FortiCloud.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Output("sslMinProtoVersion")]
        public Output<string> SslMinProtoVersion { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Day of week to roll logs.
        /// </summary>
        [Output("uploadDay")]
        public Output<string> UploadDay { get; private set; } = null!;

        /// <summary>
        /// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        /// </summary>
        [Output("uploadInterval")]
        public Output<string> UploadInterval { get; private set; } = null!;

        /// <summary>
        /// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        /// </summary>
        [Output("uploadOption")]
        public Output<string> UploadOption { get; private set; } = null!;

        /// <summary>
        /// Time of day to roll logs (hh:mm).
        /// </summary>
        [Output("uploadTime")]
        public Output<string> UploadTime { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Setting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Setting(string name, SettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:log/fortiguard/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/fortiguard/setting:Setting", name, state, MakeResourceOptions(options, id))
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
        /// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessConfig")]
        public Input<string>? AccessConfig { get; set; }

        /// <summary>
        /// FortiGate Cloud connection timeout in seconds.
        /// </summary>
        [Input("connTimeout")]
        public Input<int>? ConnTimeout { get; set; }

        /// <summary>
        /// Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
        /// </summary>
        [Input("encAlgorithm")]
        public Input<string>? EncAlgorithm { get; set; }

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
        /// FortiCloud maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Input("maxLogRate")]
        public Input<int>? MaxLogRate { get; set; }

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Source IP address used to connect FortiCloud.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Day of week to roll logs.
        /// </summary>
        [Input("uploadDay")]
        public Input<string>? UploadDay { get; set; }

        /// <summary>
        /// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        /// </summary>
        [Input("uploadInterval")]
        public Input<string>? UploadInterval { get; set; }

        /// <summary>
        /// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        /// </summary>
        [Input("uploadOption")]
        public Input<string>? UploadOption { get; set; }

        /// <summary>
        /// Time of day to roll logs (hh:mm).
        /// </summary>
        [Input("uploadTime")]
        public Input<string>? UploadTime { get; set; }

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
        /// Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessConfig")]
        public Input<string>? AccessConfig { get; set; }

        /// <summary>
        /// FortiGate Cloud connection timeout in seconds.
        /// </summary>
        [Input("connTimeout")]
        public Input<int>? ConnTimeout { get; set; }

        /// <summary>
        /// Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
        /// </summary>
        [Input("encAlgorithm")]
        public Input<string>? EncAlgorithm { get; set; }

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
        /// FortiCloud maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Input("maxLogRate")]
        public Input<int>? MaxLogRate { get; set; }

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Source IP address used to connect FortiCloud.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Day of week to roll logs.
        /// </summary>
        [Input("uploadDay")]
        public Input<string>? UploadDay { get; set; }

        /// <summary>
        /// Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
        /// </summary>
        [Input("uploadInterval")]
        public Input<string>? UploadInterval { get; set; }

        /// <summary>
        /// Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        /// </summary>
        [Input("uploadOption")]
        public Input<string>? UploadOption { get; set; }

        /// <summary>
        /// Time of day to roll logs (hh:mm).
        /// </summary>
        [Input("uploadTime")]
        public Input<string>? UploadTime { get; set; }

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
