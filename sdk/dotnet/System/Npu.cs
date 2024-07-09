// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    /// <summary>
    /// Configure NPU attributes. Applies to FortiOS Version `7.0.4`.
    /// 
    /// ## Import
    /// 
    /// System Npu can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/npu:Npu labelname SystemNpu
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/npu:Npu labelname SystemNpu
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/npu:Npu")]
    public partial class Npu : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("capwapOffload")]
        public Output<string> CapwapOffload { get; private set; } = null!;

        /// <summary>
        /// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
        /// </summary>
        [Output("dedicatedManagementAffinity")]
        public Output<string> DedicatedManagementAffinity { get; private set; } = null!;

        /// <summary>
        /// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dedicatedManagementCpu")]
        public Output<string> DedicatedManagementCpu { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
        /// </summary>
        [Output("fastpath")]
        public Output<string> Fastpath { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
        /// </summary>
        [Output("ipsecDecSubengineMask")]
        public Output<string> IpsecDecSubengineMask { get; private set; } = null!;

        /// <summary>
        /// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
        /// </summary>
        [Output("ipsecEncSubengineMask")]
        public Output<string> IpsecEncSubengineMask { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsecInboundCache")]
        public Output<string> IpsecInboundCache { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("ipsecMtuOverride")]
        public Output<string> IpsecMtuOverride { get; private set; } = null!;

        /// <summary>
        /// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ipsecOverVlink")]
        public Output<string> IpsecOverVlink { get; private set; } = null!;

        /// <summary>
        /// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
        /// </summary>
        [Output("mcastSessionAccounting")]
        public Output<string> McastSessionAccounting { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("np6CpsOptimizationMode")]
        public Output<string> Np6CpsOptimizationMode { get; private set; } = null!;

        /// <summary>
        /// Configure NPU priority protocol. The structure of `priority_protocol` block is documented below.
        /// </summary>
        [Output("priorityProtocol")]
        public Output<Outputs.NpuPriorityProtocol> PriorityProtocol { get; private set; } = null!;

        /// <summary>
        /// Enable/disable rdp offload. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("rdpOffload")]
        public Output<string> RdpOffload { get; private set; } = null!;

        /// <summary>
        /// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("sessionDeniedOffload")]
        public Output<string> SessionDeniedOffload { get; private set; } = null!;

        /// <summary>
        /// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("sseBackpressure")]
        public Output<string> SseBackpressure { get; private set; } = null!;

        /// <summary>
        /// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("stripClearTextPadding")]
        public Output<string> StripClearTextPadding { get; private set; } = null!;

        /// <summary>
        /// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("stripEspPadding")]
        public Output<string> StripEspPadding { get; private set; } = null!;

        /// <summary>
        /// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
        /// </summary>
        [Output("swNpBandwidth")]
        public Output<string> SwNpBandwidth { get; private set; } = null!;

        /// <summary>
        /// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("uespOffload")]
        public Output<string> UespOffload { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Npu resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Npu(string name, NpuArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/npu:Npu", name, args ?? new NpuArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Npu(string name, Input<string> id, NpuState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/npu:Npu", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Npu resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Npu Get(string name, Input<string> id, NpuState? state = null, CustomResourceOptions? options = null)
        {
            return new Npu(name, id, state, options);
        }
    }

    public sealed class NpuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capwapOffload")]
        public Input<string>? CapwapOffload { get; set; }

        /// <summary>
        /// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
        /// </summary>
        [Input("dedicatedManagementAffinity")]
        public Input<string>? DedicatedManagementAffinity { get; set; }

        /// <summary>
        /// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dedicatedManagementCpu")]
        public Input<string>? DedicatedManagementCpu { get; set; }

        /// <summary>
        /// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("fastpath")]
        public Input<string>? Fastpath { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
        /// </summary>
        [Input("ipsecDecSubengineMask")]
        public Input<string>? IpsecDecSubengineMask { get; set; }

        /// <summary>
        /// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
        /// </summary>
        [Input("ipsecEncSubengineMask")]
        public Input<string>? IpsecEncSubengineMask { get; set; }

        /// <summary>
        /// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsecInboundCache")]
        public Input<string>? IpsecInboundCache { get; set; }

        /// <summary>
        /// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("ipsecMtuOverride")]
        public Input<string>? IpsecMtuOverride { get; set; }

        /// <summary>
        /// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsecOverVlink")]
        public Input<string>? IpsecOverVlink { get; set; }

        /// <summary>
        /// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
        /// </summary>
        [Input("mcastSessionAccounting")]
        public Input<string>? McastSessionAccounting { get; set; }

        /// <summary>
        /// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("np6CpsOptimizationMode")]
        public Input<string>? Np6CpsOptimizationMode { get; set; }

        /// <summary>
        /// Configure NPU priority protocol. The structure of `priority_protocol` block is documented below.
        /// </summary>
        [Input("priorityProtocol")]
        public Input<Inputs.NpuPriorityProtocolArgs>? PriorityProtocol { get; set; }

        /// <summary>
        /// Enable/disable rdp offload. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rdpOffload")]
        public Input<string>? RdpOffload { get; set; }

        /// <summary>
        /// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sessionDeniedOffload")]
        public Input<string>? SessionDeniedOffload { get; set; }

        /// <summary>
        /// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sseBackpressure")]
        public Input<string>? SseBackpressure { get; set; }

        /// <summary>
        /// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("stripClearTextPadding")]
        public Input<string>? StripClearTextPadding { get; set; }

        /// <summary>
        /// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("stripEspPadding")]
        public Input<string>? StripEspPadding { get; set; }

        /// <summary>
        /// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
        /// </summary>
        [Input("swNpBandwidth")]
        public Input<string>? SwNpBandwidth { get; set; }

        /// <summary>
        /// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("uespOffload")]
        public Input<string>? UespOffload { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NpuArgs()
        {
        }
        public static new NpuArgs Empty => new NpuArgs();
    }

    public sealed class NpuState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capwapOffload")]
        public Input<string>? CapwapOffload { get; set; }

        /// <summary>
        /// Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
        /// </summary>
        [Input("dedicatedManagementAffinity")]
        public Input<string>? DedicatedManagementAffinity { get; set; }

        /// <summary>
        /// Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dedicatedManagementCpu")]
        public Input<string>? DedicatedManagementCpu { get; set; }

        /// <summary>
        /// Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("fastpath")]
        public Input<string>? Fastpath { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
        /// </summary>
        [Input("ipsecDecSubengineMask")]
        public Input<string>? IpsecDecSubengineMask { get; set; }

        /// <summary>
        /// IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
        /// </summary>
        [Input("ipsecEncSubengineMask")]
        public Input<string>? IpsecEncSubengineMask { get; set; }

        /// <summary>
        /// Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsecInboundCache")]
        public Input<string>? IpsecInboundCache { get; set; }

        /// <summary>
        /// Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("ipsecMtuOverride")]
        public Input<string>? IpsecMtuOverride { get; set; }

        /// <summary>
        /// Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ipsecOverVlink")]
        public Input<string>? IpsecOverVlink { get; set; }

        /// <summary>
        /// Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
        /// </summary>
        [Input("mcastSessionAccounting")]
        public Input<string>? McastSessionAccounting { get; set; }

        /// <summary>
        /// Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("np6CpsOptimizationMode")]
        public Input<string>? Np6CpsOptimizationMode { get; set; }

        /// <summary>
        /// Configure NPU priority protocol. The structure of `priority_protocol` block is documented below.
        /// </summary>
        [Input("priorityProtocol")]
        public Input<Inputs.NpuPriorityProtocolGetArgs>? PriorityProtocol { get; set; }

        /// <summary>
        /// Enable/disable rdp offload. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("rdpOffload")]
        public Input<string>? RdpOffload { get; set; }

        /// <summary>
        /// Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sessionDeniedOffload")]
        public Input<string>? SessionDeniedOffload { get; set; }

        /// <summary>
        /// Enable/disable sse backpressure. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sseBackpressure")]
        public Input<string>? SseBackpressure { get; set; }

        /// <summary>
        /// Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("stripClearTextPadding")]
        public Input<string>? StripClearTextPadding { get; set; }

        /// <summary>
        /// Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("stripEspPadding")]
        public Input<string>? StripEspPadding { get; set; }

        /// <summary>
        /// Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
        /// </summary>
        [Input("swNpBandwidth")]
        public Input<string>? SwNpBandwidth { get; set; }

        /// <summary>
        /// Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("uespOffload")]
        public Input<string>? UespOffload { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NpuState()
        {
        }
        public static new NpuState Empty => new NpuState();
    }
}
