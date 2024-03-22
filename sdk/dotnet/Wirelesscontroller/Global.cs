// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller
{
    /// <summary>
    /// Configure wireless controller global settings.
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
    ///     var trname = new Fortios.Wirelesscontroller.Global("trname", new()
    ///     {
    ///         ApLogServer = "disable",
    ///         ApLogServerIp = "0.0.0.0",
    ///         ApLogServerPort = 0,
    ///         ControlMessageOffload = "ebp-frame aeroscout-tag ap-list sta-list sta-cap-list stats aeroscout-mu",
    ///         DataEthernetIi = "disable",
    ///         DiscoveryMcAddr = "224.0.1.140",
    ///         FiappEthType = 5252,
    ///         ImageDownload = "enable",
    ///         IpsecBaseIp = "169.254.0.1",
    ///         LinkAggregation = "disable",
    ///         MaxClients = 0,
    ///         MaxRetransmit = 3,
    ///         MeshEthType = 8755,
    ///         RogueScanMacAdjacency = 7,
    ///         WtpShare = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// WirelessController Global can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/global:Global labelname WirelessControllerGlobal
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wirelesscontroller/global:Global labelname WirelessControllerGlobal
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/global:Global")]
    public partial class Global : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configure the number cw_acd daemons for multi-core CPU support (default = 0).
        /// </summary>
        [Output("acdProcessCount")]
        public Output<int> AcdProcessCount { get; private set; } = null!;

        /// <summary>
        /// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("apLogServer")]
        public Output<string> ApLogServer { get; private set; } = null!;

        /// <summary>
        /// IP address that APs or FortiAPs send log messages to.
        /// </summary>
        [Output("apLogServerIp")]
        public Output<string> ApLogServerIp { get; private set; } = null!;

        /// <summary>
        /// Port that APs or FortiAPs send log messages to.
        /// </summary>
        [Output("apLogServerPort")]
        public Output<int> ApLogServerPort { get; private set; } = null!;

        /// <summary>
        /// Configure CAPWAP control message data channel offload.
        /// </summary>
        [Output("controlMessageOffload")]
        public Output<string> ControlMessageOffload { get; private set; } = null!;

        /// <summary>
        /// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dataEthernetIi")]
        public Output<string> DataEthernetIi { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DFS certificate lab test mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dfsLabTest")]
        public Output<string> DfsLabTest { get; private set; } = null!;

        /// <summary>
        /// Multicast IP address for AP discovery (default = 244.0.1.140).
        /// </summary>
        [Output("discoveryMcAddr")]
        public Output<string> DiscoveryMcAddr { get; private set; } = null!;

        /// <summary>
        /// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
        /// </summary>
        [Output("fiappEthType")]
        public Output<int> FiappEthType { get; private set; } = null!;

        /// <summary>
        /// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("imageDownload")]
        public Output<string> ImageDownload { get; private set; } = null!;

        /// <summary>
        /// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
        /// </summary>
        [Output("ipsecBaseIp")]
        public Output<string> IpsecBaseIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("linkAggregation")]
        public Output<string> LinkAggregation { get; private set; } = null!;

        /// <summary>
        /// Description of the location of the wireless controller.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
        /// </summary>
        [Output("maxClients")]
        public Output<int> MaxClients { get; private set; } = null!;

        /// <summary>
        /// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
        /// </summary>
        [Output("maxRetransmit")]
        public Output<int> MaxRetransmit { get; private set; } = null!;

        /// <summary>
        /// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
        /// </summary>
        [Output("meshEthType")]
        public Output<int> MeshEthType { get; private set; } = null!;

        /// <summary>
        /// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
        /// </summary>
        [Output("nacInterval")]
        public Output<int> NacInterval { get; private set; } = null!;

        /// <summary>
        /// Name of the wireless controller.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
        /// </summary>
        [Output("rogueScanMacAdjacency")]
        public Output<int> RogueScanMacAdjacency { get; private set; } = null!;

        /// <summary>
        /// Enable/disable rolling WTP upgrade (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("rollingWtpUpgrade")]
        public Output<string> RollingWtpUpgrade { get; private set; } = null!;

        /// <summary>
        /// Minimum signal level/threshold in dBm required for the managed WTP to be included in rolling WTP upgrade (-95 to -20, default = -80).
        /// </summary>
        [Output("rollingWtpUpgradeThreshold")]
        public Output<string> RollingWtpUpgradeThreshold { get; private set; } = null!;

        /// <summary>
        /// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        /// </summary>
        [Output("tunnelMode")]
        public Output<string> TunnelMode { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Wpad daemon process count for multi-core CPU support.
        /// </summary>
        [Output("wpadProcessCount")]
        public Output<int> WpadProcessCount { get; private set; } = null!;

        /// <summary>
        /// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("wtpShare")]
        public Output<string> WtpShare { get; private set; } = null!;


        /// <summary>
        /// Create a Global resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Global(string name, GlobalArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/global:Global", name, args ?? new GlobalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Global(string name, Input<string> id, GlobalState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/global:Global", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Global resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Global Get(string name, Input<string> id, GlobalState? state = null, CustomResourceOptions? options = null)
        {
            return new Global(name, id, state, options);
        }
    }

    public sealed class GlobalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configure the number cw_acd daemons for multi-core CPU support (default = 0).
        /// </summary>
        [Input("acdProcessCount")]
        public Input<int>? AcdProcessCount { get; set; }

        /// <summary>
        /// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("apLogServer")]
        public Input<string>? ApLogServer { get; set; }

        /// <summary>
        /// IP address that APs or FortiAPs send log messages to.
        /// </summary>
        [Input("apLogServerIp")]
        public Input<string>? ApLogServerIp { get; set; }

        /// <summary>
        /// Port that APs or FortiAPs send log messages to.
        /// </summary>
        [Input("apLogServerPort")]
        public Input<int>? ApLogServerPort { get; set; }

        /// <summary>
        /// Configure CAPWAP control message data channel offload.
        /// </summary>
        [Input("controlMessageOffload")]
        public Input<string>? ControlMessageOffload { get; set; }

        /// <summary>
        /// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dataEthernetIi")]
        public Input<string>? DataEthernetIi { get; set; }

        /// <summary>
        /// Enable/disable DFS certificate lab test mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dfsLabTest")]
        public Input<string>? DfsLabTest { get; set; }

        /// <summary>
        /// Multicast IP address for AP discovery (default = 244.0.1.140).
        /// </summary>
        [Input("discoveryMcAddr")]
        public Input<string>? DiscoveryMcAddr { get; set; }

        /// <summary>
        /// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
        /// </summary>
        [Input("fiappEthType")]
        public Input<int>? FiappEthType { get; set; }

        /// <summary>
        /// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("imageDownload")]
        public Input<string>? ImageDownload { get; set; }

        /// <summary>
        /// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
        /// </summary>
        [Input("ipsecBaseIp")]
        public Input<string>? IpsecBaseIp { get; set; }

        /// <summary>
        /// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("linkAggregation")]
        public Input<string>? LinkAggregation { get; set; }

        /// <summary>
        /// Description of the location of the wireless controller.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
        /// </summary>
        [Input("maxClients")]
        public Input<int>? MaxClients { get; set; }

        /// <summary>
        /// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
        /// </summary>
        [Input("maxRetransmit")]
        public Input<int>? MaxRetransmit { get; set; }

        /// <summary>
        /// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
        /// </summary>
        [Input("meshEthType")]
        public Input<int>? MeshEthType { get; set; }

        /// <summary>
        /// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
        /// </summary>
        [Input("nacInterval")]
        public Input<int>? NacInterval { get; set; }

        /// <summary>
        /// Name of the wireless controller.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
        /// </summary>
        [Input("rogueScanMacAdjacency")]
        public Input<int>? RogueScanMacAdjacency { get; set; }

        /// <summary>
        /// Enable/disable rolling WTP upgrade (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rollingWtpUpgrade")]
        public Input<string>? RollingWtpUpgrade { get; set; }

        /// <summary>
        /// Minimum signal level/threshold in dBm required for the managed WTP to be included in rolling WTP upgrade (-95 to -20, default = -80).
        /// </summary>
        [Input("rollingWtpUpgradeThreshold")]
        public Input<string>? RollingWtpUpgradeThreshold { get; set; }

        /// <summary>
        /// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        /// </summary>
        [Input("tunnelMode")]
        public Input<string>? TunnelMode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Wpad daemon process count for multi-core CPU support.
        /// </summary>
        [Input("wpadProcessCount")]
        public Input<int>? WpadProcessCount { get; set; }

        /// <summary>
        /// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wtpShare")]
        public Input<string>? WtpShare { get; set; }

        public GlobalArgs()
        {
        }
        public static new GlobalArgs Empty => new GlobalArgs();
    }

    public sealed class GlobalState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configure the number cw_acd daemons for multi-core CPU support (default = 0).
        /// </summary>
        [Input("acdProcessCount")]
        public Input<int>? AcdProcessCount { get; set; }

        /// <summary>
        /// Enable/disable configuring APs or FortiAPs to send log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("apLogServer")]
        public Input<string>? ApLogServer { get; set; }

        /// <summary>
        /// IP address that APs or FortiAPs send log messages to.
        /// </summary>
        [Input("apLogServerIp")]
        public Input<string>? ApLogServerIp { get; set; }

        /// <summary>
        /// Port that APs or FortiAPs send log messages to.
        /// </summary>
        [Input("apLogServerPort")]
        public Input<int>? ApLogServerPort { get; set; }

        /// <summary>
        /// Configure CAPWAP control message data channel offload.
        /// </summary>
        [Input("controlMessageOffload")]
        public Input<string>? ControlMessageOffload { get; set; }

        /// <summary>
        /// Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dataEthernetIi")]
        public Input<string>? DataEthernetIi { get; set; }

        /// <summary>
        /// Enable/disable DFS certificate lab test mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dfsLabTest")]
        public Input<string>? DfsLabTest { get; set; }

        /// <summary>
        /// Multicast IP address for AP discovery (default = 244.0.1.140).
        /// </summary>
        [Input("discoveryMcAddr")]
        public Input<string>? DiscoveryMcAddr { get; set; }

        /// <summary>
        /// Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
        /// </summary>
        [Input("fiappEthType")]
        public Input<int>? FiappEthType { get; set; }

        /// <summary>
        /// Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("imageDownload")]
        public Input<string>? ImageDownload { get; set; }

        /// <summary>
        /// Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
        /// </summary>
        [Input("ipsecBaseIp")]
        public Input<string>? IpsecBaseIp { get; set; }

        /// <summary>
        /// Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("linkAggregation")]
        public Input<string>? LinkAggregation { get; set; }

        /// <summary>
        /// Description of the location of the wireless controller.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
        /// </summary>
        [Input("maxClients")]
        public Input<int>? MaxClients { get; set; }

        /// <summary>
        /// Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
        /// </summary>
        [Input("maxRetransmit")]
        public Input<int>? MaxRetransmit { get; set; }

        /// <summary>
        /// Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
        /// </summary>
        [Input("meshEthType")]
        public Input<int>? MeshEthType { get; set; }

        /// <summary>
        /// Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
        /// </summary>
        [Input("nacInterval")]
        public Input<int>? NacInterval { get; set; }

        /// <summary>
        /// Name of the wireless controller.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
        /// </summary>
        [Input("rogueScanMacAdjacency")]
        public Input<int>? RogueScanMacAdjacency { get; set; }

        /// <summary>
        /// Enable/disable rolling WTP upgrade (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rollingWtpUpgrade")]
        public Input<string>? RollingWtpUpgrade { get; set; }

        /// <summary>
        /// Minimum signal level/threshold in dBm required for the managed WTP to be included in rolling WTP upgrade (-95 to -20, default = -80).
        /// </summary>
        [Input("rollingWtpUpgradeThreshold")]
        public Input<string>? RollingWtpUpgradeThreshold { get; set; }

        /// <summary>
        /// Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
        /// </summary>
        [Input("tunnelMode")]
        public Input<string>? TunnelMode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Wpad daemon process count for multi-core CPU support.
        /// </summary>
        [Input("wpadProcessCount")]
        public Input<int>? WpadProcessCount { get; set; }

        /// <summary>
        /// Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("wtpShare")]
        public Input<string>? WtpShare { get; set; }

        public GlobalState()
        {
        }
        public static new GlobalState Empty => new GlobalState();
    }
}
