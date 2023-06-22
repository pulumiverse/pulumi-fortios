// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure Wireless Termination Points (WTP) system log server profile. Applies to FortiOS Version `&gt;= 7.0.2`.
    /// 
    /// ## Import
    /// 
    /// WirelessController SyslogProfile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile")]
    public partial class WirelesscontrollerSyslogprofile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        /// </summary>
        [Output("logLevel")]
        public Output<string> LogLevel { get; private set; } = null!;

        /// <summary>
        /// WTP system log server profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        /// </summary>
        [Output("serverAddrType")]
        public Output<string> ServerAddrType { get; private set; } = null!;

        /// <summary>
        /// FQDN of syslog server that FortiAP units send log messages to.
        /// </summary>
        [Output("serverFqdn")]
        public Output<string> ServerFqdn { get; private set; } = null!;

        /// <summary>
        /// IP address of syslog server that FortiAP units send log messages to.
        /// </summary>
        [Output("serverIp")]
        public Output<string> ServerIp { get; private set; } = null!;

        /// <summary>
        /// Port number of syslog server that FortiAP units send log messages to (default = 514).
        /// </summary>
        [Output("serverPort")]
        public Output<int> ServerPort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("serverStatus")]
        public Output<string> ServerStatus { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a WirelesscontrollerSyslogprofile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WirelesscontrollerSyslogprofile(string name, WirelesscontrollerSyslogprofileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile", name, args ?? new WirelesscontrollerSyslogprofileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WirelesscontrollerSyslogprofile(string name, Input<string> id, WirelesscontrollerSyslogprofileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/wirelesscontrollerSyslogprofile:WirelesscontrollerSyslogprofile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WirelesscontrollerSyslogprofile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WirelesscontrollerSyslogprofile Get(string name, Input<string> id, WirelesscontrollerSyslogprofileState? state = null, CustomResourceOptions? options = null)
        {
            return new WirelesscontrollerSyslogprofile(name, id, state, options);
        }
    }

    public sealed class WirelesscontrollerSyslogprofileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// WTP system log server profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        /// </summary>
        [Input("serverAddrType")]
        public Input<string>? ServerAddrType { get; set; }

        /// <summary>
        /// FQDN of syslog server that FortiAP units send log messages to.
        /// </summary>
        [Input("serverFqdn")]
        public Input<string>? ServerFqdn { get; set; }

        /// <summary>
        /// IP address of syslog server that FortiAP units send log messages to.
        /// </summary>
        [Input("serverIp")]
        public Input<string>? ServerIp { get; set; }

        /// <summary>
        /// Port number of syslog server that FortiAP units send log messages to (default = 514).
        /// </summary>
        [Input("serverPort")]
        public Input<int>? ServerPort { get; set; }

        /// <summary>
        /// Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("serverStatus")]
        public Input<string>? ServerStatus { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WirelesscontrollerSyslogprofileArgs()
        {
        }
        public static new WirelesscontrollerSyslogprofileArgs Empty => new WirelesscontrollerSyslogprofileArgs();
    }

    public sealed class WirelesscontrollerSyslogprofileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Lowest level of log messages that FortiAP units send to this server (default = information) Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debugging`.
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// WTP system log server profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Syslog server address type (default = IP) Valid values: `fqdn`, `ip`.
        /// </summary>
        [Input("serverAddrType")]
        public Input<string>? ServerAddrType { get; set; }

        /// <summary>
        /// FQDN of syslog server that FortiAP units send log messages to.
        /// </summary>
        [Input("serverFqdn")]
        public Input<string>? ServerFqdn { get; set; }

        /// <summary>
        /// IP address of syslog server that FortiAP units send log messages to.
        /// </summary>
        [Input("serverIp")]
        public Input<string>? ServerIp { get; set; }

        /// <summary>
        /// Port number of syslog server that FortiAP units send log messages to (default = 514).
        /// </summary>
        [Input("serverPort")]
        public Input<int>? ServerPort { get; set; }

        /// <summary>
        /// Enable/disable FortiAP units to send log messages to a syslog server (default = enable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("serverStatus")]
        public Input<string>? ServerStatus { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WirelesscontrollerSyslogprofileState()
        {
        }
        public static new WirelesscontrollerSyslogprofileState Empty => new WirelesscontrollerSyslogprofileState();
    }
}
