// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Logfortianalyzer3
{
    /// <summary>
    /// Override FortiAnalyzer settings.
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
    ///     var trname = new Fortios.Logfortianalyzer3.Overridesetting("trname", new()
    ///     {
    ///         __changeIp = 0,
    ///         ConnTimeout = 10,
    ///         EncAlgorithm = "high",
    ///         FazType = 6,
    ///         HmacAlgorithm = "sha256",
    ///         IpsArchive = "enable",
    ///         MonitorFailureRetryPeriod = 5,
    ///         MonitorKeepalivePeriod = 5,
    ///         Override = "disable",
    ///         Reliable = "disable",
    ///         SslMinProtoVersion = "default",
    ///         Status = "disable",
    ///         UploadInterval = "daily",
    ///         UploadOption = "5-minute",
    ///         UploadTime = "00:59",
    ///         UseManagementVdom = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LogFortianalyzer3 OverrideSetting can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:logfortianalyzer3/overridesetting:Overridesetting labelname LogFortianalyzer3OverrideSetting
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:logfortianalyzer3/overridesetting:Overridesetting labelname LogFortianalyzer3OverrideSetting
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:logfortianalyzer3/overridesetting:Overridesetting")]
    public partial class Overridesetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Hidden attribute.
        /// </summary>
        [Output("__changeIp")]
        public Output<int> __changeIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("accessConfig")]
        public Output<string> AccessConfig { get; private set; } = null!;

        /// <summary>
        /// Certificate used to communicate with FortiAnalyzer.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("certificateVerification")]
        public Output<string> CertificateVerification { get; private set; } = null!;

        /// <summary>
        /// FortiAnalyzer connection time-out in seconds (for status and log buffer).
        /// </summary>
        [Output("connTimeout")]
        public Output<int> ConnTimeout { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
        /// </summary>
        [Output("encAlgorithm")]
        public Output<string> EncAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Hidden setting index of FortiAnalyzer.
        /// </summary>
        [Output("fazType")]
        public Output<int> FazType { get; private set; } = null!;

        /// <summary>
        /// FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
        /// </summary>
        [Output("hmacAlgorithm")]
        public Output<string> HmacAlgorithm { get; private set; } = null!;

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
        /// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsArchive")]
        public Output<string> IpsArchive { get; private set; } = null!;

        /// <summary>
        /// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Output("maxLogRate")]
        public Output<int> MaxLogRate { get; private set; } = null!;

        /// <summary>
        /// Hidden management name of FortiAnalyzer.
        /// </summary>
        [Output("mgmtName")]
        public Output<string> MgmtName { get; private set; } = null!;

        /// <summary>
        /// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
        /// </summary>
        [Output("monitorFailureRetryPeriod")]
        public Output<int> MonitorFailureRetryPeriod { get; private set; } = null!;

        /// <summary>
        /// Time between OFTP keepalives in seconds (for status and log buffer).
        /// </summary>
        [Output("monitorKeepalivePeriod")]
        public Output<int> MonitorKeepalivePeriod { get; private set; } = null!;

        /// <summary>
        /// Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("override")]
        public Output<string> Override { get; private set; } = null!;

        /// <summary>
        /// Preshared-key used for auto-authorization on FortiAnalyzer.
        /// </summary>
        [Output("presharedKey")]
        public Output<string> PresharedKey { get; private set; } = null!;

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        /// <summary>
        /// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("reliable")]
        public Output<string> Reliable { get; private set; } = null!;

        /// <summary>
        /// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
        /// </summary>
        [Output("serials")]
        public Output<ImmutableArray<Outputs.OverridesettingSerial>> Serials { get; private set; } = null!;

        /// <summary>
        /// The remote FortiAnalyzer.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Output("sslMinProtoVersion")]
        public Output<string> SslMinProtoVersion { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Day of week (month) to upload logs.
        /// </summary>
        [Output("uploadDay")]
        public Output<string> UploadDay { get; private set; } = null!;

        /// <summary>
        /// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
        /// </summary>
        [Output("uploadInterval")]
        public Output<string> UploadInterval { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        /// </summary>
        [Output("uploadOption")]
        public Output<string> UploadOption { get; private set; } = null!;

        /// <summary>
        /// Time to upload logs (hh:mm).
        /// </summary>
        [Output("uploadTime")]
        public Output<string> UploadTime { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("useManagementVdom")]
        public Output<string> UseManagementVdom { get; private set; } = null!;

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
            : base("fortios:logfortianalyzer3/overridesetting:Overridesetting", name, args ?? new OverridesettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Overridesetting(string name, Input<string> id, OverridesettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:logfortianalyzer3/overridesetting:Overridesetting", name, state, MakeResourceOptions(options, id))
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
        /// Hidden attribute.
        /// </summary>
        [Input("__changeIp")]
        public Input<int>? __changeIp { get; set; }

        /// <summary>
        /// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessConfig")]
        public Input<string>? AccessConfig { get; set; }

        /// <summary>
        /// Certificate used to communicate with FortiAnalyzer.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("certificateVerification")]
        public Input<string>? CertificateVerification { get; set; }

        /// <summary>
        /// FortiAnalyzer connection time-out in seconds (for status and log buffer).
        /// </summary>
        [Input("connTimeout")]
        public Input<int>? ConnTimeout { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
        /// </summary>
        [Input("encAlgorithm")]
        public Input<string>? EncAlgorithm { get; set; }

        /// <summary>
        /// Hidden setting index of FortiAnalyzer.
        /// </summary>
        [Input("fazType")]
        public Input<int>? FazType { get; set; }

        /// <summary>
        /// FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
        /// </summary>
        [Input("hmacAlgorithm")]
        public Input<string>? HmacAlgorithm { get; set; }

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
        /// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsArchive")]
        public Input<string>? IpsArchive { get; set; }

        /// <summary>
        /// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Input("maxLogRate")]
        public Input<int>? MaxLogRate { get; set; }

        /// <summary>
        /// Hidden management name of FortiAnalyzer.
        /// </summary>
        [Input("mgmtName")]
        public Input<string>? MgmtName { get; set; }

        /// <summary>
        /// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
        /// </summary>
        [Input("monitorFailureRetryPeriod")]
        public Input<int>? MonitorFailureRetryPeriod { get; set; }

        /// <summary>
        /// Time between OFTP keepalives in seconds (for status and log buffer).
        /// </summary>
        [Input("monitorKeepalivePeriod")]
        public Input<int>? MonitorKeepalivePeriod { get; set; }

        /// <summary>
        /// Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("override")]
        public Input<string>? Override { get; set; }

        /// <summary>
        /// Preshared-key used for auto-authorization on FortiAnalyzer.
        /// </summary>
        [Input("presharedKey")]
        public Input<string>? PresharedKey { get; set; }

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("reliable")]
        public Input<string>? Reliable { get; set; }

        [Input("serials")]
        private InputList<Inputs.OverridesettingSerialArgs>? _serials;

        /// <summary>
        /// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
        /// </summary>
        public InputList<Inputs.OverridesettingSerialArgs> Serials
        {
            get => _serials ?? (_serials = new InputList<Inputs.OverridesettingSerialArgs>());
            set => _serials = value;
        }

        /// <summary>
        /// The remote FortiAnalyzer.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Day of week (month) to upload logs.
        /// </summary>
        [Input("uploadDay")]
        public Input<string>? UploadDay { get; set; }

        /// <summary>
        /// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
        /// </summary>
        [Input("uploadInterval")]
        public Input<string>? UploadInterval { get; set; }

        /// <summary>
        /// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        /// </summary>
        [Input("uploadOption")]
        public Input<string>? UploadOption { get; set; }

        /// <summary>
        /// Time to upload logs (hh:mm).
        /// </summary>
        [Input("uploadTime")]
        public Input<string>? UploadTime { get; set; }

        /// <summary>
        /// Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("useManagementVdom")]
        public Input<string>? UseManagementVdom { get; set; }

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
        /// Hidden attribute.
        /// </summary>
        [Input("__changeIp")]
        public Input<int>? __changeIp { get; set; }

        /// <summary>
        /// Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("accessConfig")]
        public Input<string>? AccessConfig { get; set; }

        /// <summary>
        /// Certificate used to communicate with FortiAnalyzer.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("certificateVerification")]
        public Input<string>? CertificateVerification { get; set; }

        /// <summary>
        /// FortiAnalyzer connection time-out in seconds (for status and log buffer).
        /// </summary>
        [Input("connTimeout")]
        public Input<int>? ConnTimeout { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
        /// </summary>
        [Input("encAlgorithm")]
        public Input<string>? EncAlgorithm { get; set; }

        /// <summary>
        /// Hidden setting index of FortiAnalyzer.
        /// </summary>
        [Input("fazType")]
        public Input<int>? FazType { get; set; }

        /// <summary>
        /// FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
        /// </summary>
        [Input("hmacAlgorithm")]
        public Input<string>? HmacAlgorithm { get; set; }

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
        /// Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsArchive")]
        public Input<string>? IpsArchive { get; set; }

        /// <summary>
        /// FortiAnalyzer maximum log rate in MBps (0 = unlimited).
        /// </summary>
        [Input("maxLogRate")]
        public Input<int>? MaxLogRate { get; set; }

        /// <summary>
        /// Hidden management name of FortiAnalyzer.
        /// </summary>
        [Input("mgmtName")]
        public Input<string>? MgmtName { get; set; }

        /// <summary>
        /// Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
        /// </summary>
        [Input("monitorFailureRetryPeriod")]
        public Input<int>? MonitorFailureRetryPeriod { get; set; }

        /// <summary>
        /// Time between OFTP keepalives in seconds (for status and log buffer).
        /// </summary>
        [Input("monitorKeepalivePeriod")]
        public Input<int>? MonitorKeepalivePeriod { get; set; }

        /// <summary>
        /// Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("override")]
        public Input<string>? Override { get; set; }

        /// <summary>
        /// Preshared-key used for auto-authorization on FortiAnalyzer.
        /// </summary>
        [Input("presharedKey")]
        public Input<string>? PresharedKey { get; set; }

        /// <summary>
        /// Set log transmission priority. Valid values: `default`, `low`.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        /// <summary>
        /// Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("reliable")]
        public Input<string>? Reliable { get; set; }

        [Input("serials")]
        private InputList<Inputs.OverridesettingSerialGetArgs>? _serials;

        /// <summary>
        /// Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
        /// </summary>
        public InputList<Inputs.OverridesettingSerialGetArgs> Serials
        {
            get => _serials ?? (_serials = new InputList<Inputs.OverridesettingSerialGetArgs>());
            set => _serials = value;
        }

        /// <summary>
        /// The remote FortiAnalyzer.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        [Input("sslMinProtoVersion")]
        public Input<string>? SslMinProtoVersion { get; set; }

        /// <summary>
        /// Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Day of week (month) to upload logs.
        /// </summary>
        [Input("uploadDay")]
        public Input<string>? UploadDay { get; set; }

        /// <summary>
        /// Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
        /// </summary>
        [Input("uploadInterval")]
        public Input<string>? UploadInterval { get; set; }

        /// <summary>
        /// Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
        /// </summary>
        [Input("uploadOption")]
        public Input<string>? UploadOption { get; set; }

        /// <summary>
        /// Time to upload logs (hh:mm).
        /// </summary>
        [Input("uploadTime")]
        public Input<string>? UploadTime { get; set; }

        /// <summary>
        /// Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("useManagementVdom")]
        public Input<string>? UseManagementVdom { get; set; }

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
