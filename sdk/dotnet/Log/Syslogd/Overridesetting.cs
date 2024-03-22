// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log.Syslogd
{
    /// <summary>
    /// Override settings for remote syslog server.
    /// 
    /// ## Import
    /// 
    /// LogSyslogd OverrideSetting can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/syslogd/overridesetting:Overridesetting labelname LogSyslogdOverrideSetting
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/syslogd/overridesetting:Overridesetting labelname LogSyslogdOverrideSetting
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:log/syslogd/overridesetting:Overridesetting")]
    public partial class Overridesetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Certificate used to communicate with Syslog server.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Custom field name for CEF format logging. The structure of `custom_field_name` block is documented below.
        /// </summary>
        [Output("customFieldNames")]
        public Output<ImmutableArray<Outputs.OverridesettingCustomFieldName>> CustomFieldNames { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
        /// </summary>
        [Output("encAlgorithm")]
        public Output<string> EncAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        /// </summary>
        [Output("facility")]
        public Output<string> Facility { get; private set; } = null!;

        /// <summary>
        /// Log format.
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

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
        /// Syslog maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Output("maxLogRate")]
        public Output<int> MaxLogRate { get; private set; } = null!;

        /// <summary>
        /// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("override")]
        public Output<string> Override { get; private set; } = null!;

        /// <summary>
        /// Server listen port.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

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
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Output("sslMinProtoVersion")]
        public Output<string> SslMinProtoVersion { get; private set; } = null!;

        /// <summary>
        /// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Hidden setting index of Syslog.
        /// </summary>
        [Output("syslogType")]
        public Output<int> SyslogType { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Overridesetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Overridesetting(string name, OverridesettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:log/syslogd/overridesetting:Overridesetting", name, args ?? new OverridesettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Overridesetting(string name, Input<string> id, OverridesettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/syslogd/overridesetting:Overridesetting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Overridesetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Overridesetting Get(string name, Input<string> id, OverridesettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Overridesetting(name, id, state, options);
        }
    }

    public sealed class OverridesettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate used to communicate with Syslog server.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("customFieldNames")]
        private InputList<Inputs.OverridesettingCustomFieldNameArgs>? _customFieldNames;

        /// <summary>
        /// Custom field name for CEF format logging. The structure of `custom_field_name` block is documented below.
        /// </summary>
        public InputList<Inputs.OverridesettingCustomFieldNameArgs> CustomFieldNames
        {
            get => _customFieldNames ?? (_customFieldNames = new InputList<Inputs.OverridesettingCustomFieldNameArgs>());
            set => _customFieldNames = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
        /// </summary>
        [Input("encAlgorithm")]
        public Input<string>? EncAlgorithm { get; set; }

        /// <summary>
        /// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Log format.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

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
        /// Syslog maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Input("maxLogRate")]
        public Input<int>? MaxLogRate { get; set; }

        /// <summary>
        /// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("override")]
        public Input<string>? Override { get; set; }

        /// <summary>
        /// Server listen port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

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
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Hidden setting index of Syslog.
        /// </summary>
        [Input("syslogType")]
        public Input<int>? SyslogType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OverridesettingArgs()
        {
        }
        public static new OverridesettingArgs Empty => new OverridesettingArgs();
    }

    public sealed class OverridesettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate used to communicate with Syslog server.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("customFieldNames")]
        private InputList<Inputs.OverridesettingCustomFieldNameGetArgs>? _customFieldNames;

        /// <summary>
        /// Custom field name for CEF format logging. The structure of `custom_field_name` block is documented below.
        /// </summary>
        public InputList<Inputs.OverridesettingCustomFieldNameGetArgs> CustomFieldNames
        {
            get => _customFieldNames ?? (_customFieldNames = new InputList<Inputs.OverridesettingCustomFieldNameGetArgs>());
            set => _customFieldNames = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
        /// </summary>
        [Input("encAlgorithm")]
        public Input<string>? EncAlgorithm { get; set; }

        /// <summary>
        /// Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Log format.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

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
        /// Syslog maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Input("maxLogRate")]
        public Input<int>? MaxLogRate { get; set; }

        /// <summary>
        /// Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Enable/disable override syslog settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("override")]
        public Input<string>? Override { get; set; }

        /// <summary>
        /// Server listen port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

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
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Hidden setting index of Syslog.
        /// </summary>
        [Input("syslogType")]
        public Input<int>? SyslogType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OverridesettingState()
        {
        }
        public static new OverridesettingState Empty => new OverridesettingState();
    }
}
