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
    /// Configure system PTP information.
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
    ///     var trname = new Fortios.SystemPtp("trname", new()
    ///     {
    ///         DelayMechanism = "E2E",
    ///         Interface = "port3",
    ///         Mode = "multicast",
    ///         RequestInterval = 1,
    ///         Status = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System Ptp can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemPtp:SystemPtp labelname SystemPtp
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemPtp:SystemPtp labelname SystemPtp
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemPtp:SystemPtp")]
    public partial class SystemPtp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        /// </summary>
        [Output("delayMechanism")]
        public Output<string> DelayMechanism { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// PTP slave will reply through this interface.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        /// </summary>
        [Output("requestInterval")]
        public Output<int> RequestInterval { get; private set; } = null!;

        /// <summary>
        /// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        /// </summary>
        [Output("serverInterfaces")]
        public Output<ImmutableArray<Outputs.SystemPtpServerInterface>> ServerInterfaces { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("serverMode")]
        public Output<string> ServerMode { get; private set; } = null!;

        /// <summary>
        /// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemPtp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemPtp(string name, SystemPtpArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/systemPtp:SystemPtp", name, args ?? new SystemPtpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemPtp(string name, Input<string> id, SystemPtpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemPtp:SystemPtp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemPtp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemPtp Get(string name, Input<string> id, SystemPtpState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemPtp(name, id, state, options);
        }
    }

    public sealed class SystemPtpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        /// </summary>
        [Input("delayMechanism")]
        public Input<string>? DelayMechanism { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// PTP slave will reply through this interface.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        /// </summary>
        [Input("requestInterval")]
        public Input<int>? RequestInterval { get; set; }

        [Input("serverInterfaces")]
        private InputList<Inputs.SystemPtpServerInterfaceArgs>? _serverInterfaces;

        /// <summary>
        /// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemPtpServerInterfaceArgs> ServerInterfaces
        {
            get => _serverInterfaces ?? (_serverInterfaces = new InputList<Inputs.SystemPtpServerInterfaceArgs>());
            set => _serverInterfaces = value;
        }

        /// <summary>
        /// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("serverMode")]
        public Input<string>? ServerMode { get; set; }

        /// <summary>
        /// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemPtpArgs()
        {
        }
        public static new SystemPtpArgs Empty => new SystemPtpArgs();
    }

    public sealed class SystemPtpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
        /// </summary>
        [Input("delayMechanism")]
        public Input<string>? DelayMechanism { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// PTP slave will reply through this interface.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
        /// </summary>
        [Input("requestInterval")]
        public Input<int>? RequestInterval { get; set; }

        [Input("serverInterfaces")]
        private InputList<Inputs.SystemPtpServerInterfaceGetArgs>? _serverInterfaces;

        /// <summary>
        /// FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemPtpServerInterfaceGetArgs> ServerInterfaces
        {
            get => _serverInterfaces ?? (_serverInterfaces = new InputList<Inputs.SystemPtpServerInterfaceGetArgs>());
            set => _serverInterfaces = value;
        }

        /// <summary>
        /// Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("serverMode")]
        public Input<string>? ServerMode { get; set; }

        /// <summary>
        /// Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemPtpState()
        {
        }
        public static new SystemPtpState Empty => new SystemPtpState();
    }
}
