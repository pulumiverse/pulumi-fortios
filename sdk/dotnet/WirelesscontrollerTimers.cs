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
    /// Configure CAPWAP timers.
    /// 
    /// ## Import
    /// 
    /// WirelessController Timers can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wirelesscontrollerTimers:WirelesscontrollerTimers labelname WirelessControllerTimers
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wirelesscontrollerTimers:WirelesscontrollerTimers labelname WirelessControllerTimers
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/wirelesscontrollerTimers:WirelesscontrollerTimers")]
    public partial class WirelesscontrollerTimers : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time after which a client is considered failed in RADIUS authentication and times out (5 - 30 sec, default = 5).
        /// </summary>
        [Output("authTimeout")]
        public Output<int> AuthTimeout { get; private set; } = null!;

        /// <summary>
        /// Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
        /// </summary>
        [Output("bleScanReportIntv")]
        public Output<int> BleScanReportIntv { get; private set; } = null!;

        /// <summary>
        /// Time after which a client is considered idle and disconnected from the home controller (2 - 3600 sec, default = 20, 0 for no timeout).
        /// </summary>
        [Output("clientIdleRehomeTimeout")]
        public Output<int> ClientIdleRehomeTimeout { get; private set; } = null!;

        /// <summary>
        /// Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
        /// </summary>
        [Output("clientIdleTimeout")]
        public Output<int> ClientIdleTimeout { get; private set; } = null!;

        /// <summary>
        /// Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
        /// </summary>
        [Output("darrpDay")]
        public Output<string> DarrpDay { get; private set; } = null!;

        /// <summary>
        /// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
        /// </summary>
        [Output("darrpOptimize")]
        public Output<int> DarrpOptimize { get; private set; } = null!;

        /// <summary>
        /// Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrp_time` block is documented below.
        /// </summary>
        [Output("darrpTimes")]
        public Output<ImmutableArray<Outputs.WirelesscontrollerTimersDarrpTime>> DarrpTimes { get; private set; } = null!;

        /// <summary>
        /// Time between discovery requests (2 - 180 sec, default = 5).
        /// </summary>
        [Output("discoveryInterval")]
        public Output<int> DiscoveryInterval { get; private set; } = null!;

        /// <summary>
        /// Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
        /// </summary>
        [Output("drmaInterval")]
        public Output<int> DrmaInterval { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
        /// </summary>
        [Output("echoInterval")]
        public Output<int> EchoInterval { get; private set; } = null!;

        /// <summary>
        /// Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
        /// </summary>
        [Output("fakeApLog")]
        public Output<int> FakeApLog { get; private set; } = null!;

        /// <summary>
        /// Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
        /// </summary>
        [Output("ipsecIntfCleanup")]
        public Output<int> IpsecIntfCleanup { get; private set; } = null!;

        /// <summary>
        /// Time between running radio reports (1 - 255 sec, default = 15).
        /// </summary>
        [Output("radioStatsInterval")]
        public Output<int> RadioStatsInterval { get; private set; } = null!;

        /// <summary>
        /// Time period in minutes to keep rogue AP after it is gone (default = 0).
        /// </summary>
        [Output("rogueApCleanup")]
        public Output<int> RogueApCleanup { get; private set; } = null!;

        /// <summary>
        /// Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
        /// </summary>
        [Output("rogueApLog")]
        public Output<int> RogueApLog { get; private set; } = null!;

        /// <summary>
        /// Time between running station capability reports (1 - 255 sec, default = 30).
        /// </summary>
        [Output("staCapabilityInterval")]
        public Output<int> StaCapabilityInterval { get; private set; } = null!;

        /// <summary>
        /// Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
        /// </summary>
        [Output("staLocateTimer")]
        public Output<int> StaLocateTimer { get; private set; } = null!;

        /// <summary>
        /// Time between running client (station) reports (1 - 255 sec, default = 1).
        /// </summary>
        [Output("staStatsInterval")]
        public Output<int> StaStatsInterval { get; private set; } = null!;

        /// <summary>
        /// Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
        /// </summary>
        [Output("vapStatsInterval")]
        public Output<int> VapStatsInterval { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a WirelesscontrollerTimers resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WirelesscontrollerTimers(string name, WirelesscontrollerTimersArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/wirelesscontrollerTimers:WirelesscontrollerTimers", name, args ?? new WirelesscontrollerTimersArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WirelesscontrollerTimers(string name, Input<string> id, WirelesscontrollerTimersState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/wirelesscontrollerTimers:WirelesscontrollerTimers", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WirelesscontrollerTimers resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WirelesscontrollerTimers Get(string name, Input<string> id, WirelesscontrollerTimersState? state = null, CustomResourceOptions? options = null)
        {
            return new WirelesscontrollerTimers(name, id, state, options);
        }
    }

    public sealed class WirelesscontrollerTimersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time after which a client is considered failed in RADIUS authentication and times out (5 - 30 sec, default = 5).
        /// </summary>
        [Input("authTimeout")]
        public Input<int>? AuthTimeout { get; set; }

        /// <summary>
        /// Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
        /// </summary>
        [Input("bleScanReportIntv")]
        public Input<int>? BleScanReportIntv { get; set; }

        /// <summary>
        /// Time after which a client is considered idle and disconnected from the home controller (2 - 3600 sec, default = 20, 0 for no timeout).
        /// </summary>
        [Input("clientIdleRehomeTimeout")]
        public Input<int>? ClientIdleRehomeTimeout { get; set; }

        /// <summary>
        /// Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
        /// </summary>
        [Input("clientIdleTimeout")]
        public Input<int>? ClientIdleTimeout { get; set; }

        /// <summary>
        /// Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
        /// </summary>
        [Input("darrpDay")]
        public Input<string>? DarrpDay { get; set; }

        /// <summary>
        /// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
        /// </summary>
        [Input("darrpOptimize")]
        public Input<int>? DarrpOptimize { get; set; }

        [Input("darrpTimes")]
        private InputList<Inputs.WirelesscontrollerTimersDarrpTimeArgs>? _darrpTimes;

        /// <summary>
        /// Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrp_time` block is documented below.
        /// </summary>
        public InputList<Inputs.WirelesscontrollerTimersDarrpTimeArgs> DarrpTimes
        {
            get => _darrpTimes ?? (_darrpTimes = new InputList<Inputs.WirelesscontrollerTimersDarrpTimeArgs>());
            set => _darrpTimes = value;
        }

        /// <summary>
        /// Time between discovery requests (2 - 180 sec, default = 5).
        /// </summary>
        [Input("discoveryInterval")]
        public Input<int>? DiscoveryInterval { get; set; }

        /// <summary>
        /// Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
        /// </summary>
        [Input("drmaInterval")]
        public Input<int>? DrmaInterval { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
        /// </summary>
        [Input("echoInterval")]
        public Input<int>? EchoInterval { get; set; }

        /// <summary>
        /// Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
        /// </summary>
        [Input("fakeApLog")]
        public Input<int>? FakeApLog { get; set; }

        /// <summary>
        /// Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
        /// </summary>
        [Input("ipsecIntfCleanup")]
        public Input<int>? IpsecIntfCleanup { get; set; }

        /// <summary>
        /// Time between running radio reports (1 - 255 sec, default = 15).
        /// </summary>
        [Input("radioStatsInterval")]
        public Input<int>? RadioStatsInterval { get; set; }

        /// <summary>
        /// Time period in minutes to keep rogue AP after it is gone (default = 0).
        /// </summary>
        [Input("rogueApCleanup")]
        public Input<int>? RogueApCleanup { get; set; }

        /// <summary>
        /// Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
        /// </summary>
        [Input("rogueApLog")]
        public Input<int>? RogueApLog { get; set; }

        /// <summary>
        /// Time between running station capability reports (1 - 255 sec, default = 30).
        /// </summary>
        [Input("staCapabilityInterval")]
        public Input<int>? StaCapabilityInterval { get; set; }

        /// <summary>
        /// Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
        /// </summary>
        [Input("staLocateTimer")]
        public Input<int>? StaLocateTimer { get; set; }

        /// <summary>
        /// Time between running client (station) reports (1 - 255 sec, default = 1).
        /// </summary>
        [Input("staStatsInterval")]
        public Input<int>? StaStatsInterval { get; set; }

        /// <summary>
        /// Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
        /// </summary>
        [Input("vapStatsInterval")]
        public Input<int>? VapStatsInterval { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WirelesscontrollerTimersArgs()
        {
        }
        public static new WirelesscontrollerTimersArgs Empty => new WirelesscontrollerTimersArgs();
    }

    public sealed class WirelesscontrollerTimersState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time after which a client is considered failed in RADIUS authentication and times out (5 - 30 sec, default = 5).
        /// </summary>
        [Input("authTimeout")]
        public Input<int>? AuthTimeout { get; set; }

        /// <summary>
        /// Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
        /// </summary>
        [Input("bleScanReportIntv")]
        public Input<int>? BleScanReportIntv { get; set; }

        /// <summary>
        /// Time after which a client is considered idle and disconnected from the home controller (2 - 3600 sec, default = 20, 0 for no timeout).
        /// </summary>
        [Input("clientIdleRehomeTimeout")]
        public Input<int>? ClientIdleRehomeTimeout { get; set; }

        /// <summary>
        /// Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
        /// </summary>
        [Input("clientIdleTimeout")]
        public Input<int>? ClientIdleTimeout { get; set; }

        /// <summary>
        /// Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
        /// </summary>
        [Input("darrpDay")]
        public Input<string>? DarrpDay { get; set; }

        /// <summary>
        /// Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
        /// </summary>
        [Input("darrpOptimize")]
        public Input<int>? DarrpOptimize { get; set; }

        [Input("darrpTimes")]
        private InputList<Inputs.WirelesscontrollerTimersDarrpTimeGetArgs>? _darrpTimes;

        /// <summary>
        /// Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrp_time` block is documented below.
        /// </summary>
        public InputList<Inputs.WirelesscontrollerTimersDarrpTimeGetArgs> DarrpTimes
        {
            get => _darrpTimes ?? (_darrpTimes = new InputList<Inputs.WirelesscontrollerTimersDarrpTimeGetArgs>());
            set => _darrpTimes = value;
        }

        /// <summary>
        /// Time between discovery requests (2 - 180 sec, default = 5).
        /// </summary>
        [Input("discoveryInterval")]
        public Input<int>? DiscoveryInterval { get; set; }

        /// <summary>
        /// Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
        /// </summary>
        [Input("drmaInterval")]
        public Input<int>? DrmaInterval { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
        /// </summary>
        [Input("echoInterval")]
        public Input<int>? EchoInterval { get; set; }

        /// <summary>
        /// Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
        /// </summary>
        [Input("fakeApLog")]
        public Input<int>? FakeApLog { get; set; }

        /// <summary>
        /// Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
        /// </summary>
        [Input("ipsecIntfCleanup")]
        public Input<int>? IpsecIntfCleanup { get; set; }

        /// <summary>
        /// Time between running radio reports (1 - 255 sec, default = 15).
        /// </summary>
        [Input("radioStatsInterval")]
        public Input<int>? RadioStatsInterval { get; set; }

        /// <summary>
        /// Time period in minutes to keep rogue AP after it is gone (default = 0).
        /// </summary>
        [Input("rogueApCleanup")]
        public Input<int>? RogueApCleanup { get; set; }

        /// <summary>
        /// Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
        /// </summary>
        [Input("rogueApLog")]
        public Input<int>? RogueApLog { get; set; }

        /// <summary>
        /// Time between running station capability reports (1 - 255 sec, default = 30).
        /// </summary>
        [Input("staCapabilityInterval")]
        public Input<int>? StaCapabilityInterval { get; set; }

        /// <summary>
        /// Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
        /// </summary>
        [Input("staLocateTimer")]
        public Input<int>? StaLocateTimer { get; set; }

        /// <summary>
        /// Time between running client (station) reports (1 - 255 sec, default = 1).
        /// </summary>
        [Input("staStatsInterval")]
        public Input<int>? StaStatsInterval { get; set; }

        /// <summary>
        /// Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
        /// </summary>
        [Input("vapStatsInterval")]
        public Input<int>? VapStatsInterval { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WirelesscontrollerTimersState()
        {
        }
        public static new WirelesscontrollerTimersState Empty => new WirelesscontrollerTimersState();
    }
}
