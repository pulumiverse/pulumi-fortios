// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller
{
    /// <summary>
    /// Configure system-wide switch controller settings.
    /// 
    /// ## Import
    /// 
    /// SwitchController System can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/system:System")]
    public partial class System : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        /// </summary>
        [Output("dataSyncInterval")]
        public Output<int> DataSyncInterval { get; private set; } = null!;

        /// <summary>
        /// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        /// </summary>
        [Output("dynamicPeriodicInterval")]
        public Output<int> DynamicPeriodicInterval { get; private set; } = null!;

        /// <summary>
        /// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        /// </summary>
        [Output("iotHoldoff")]
        public Output<int> IotHoldoff { get; private set; } = null!;

        /// <summary>
        /// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        /// </summary>
        [Output("iotMacIdle")]
        public Output<int> IotMacIdle { get; private set; } = null!;

        /// <summary>
        /// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        /// </summary>
        [Output("iotScanInterval")]
        public Output<int> IotScanInterval { get; private set; } = null!;

        /// <summary>
        /// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        /// </summary>
        [Output("iotWeightThreshold")]
        public Output<int> IotWeightThreshold { get; private set; } = null!;

        /// <summary>
        /// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        /// </summary>
        [Output("nacPeriodicInterval")]
        public Output<int> NacPeriodicInterval { get; private set; } = null!;

        /// <summary>
        /// Maximum number of parallel processes (1 - 300, default = 1).
        /// </summary>
        [Output("parallelProcess")]
        public Output<int> ParallelProcess { get; private set; } = null!;

        /// <summary>
        /// Enable/disable parallel process override. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("parallelProcessOverride")]
        public Output<string> ParallelProcessOverride { get; private set; } = null!;

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
        /// Create a System resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public System(string name, SystemArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/system:System", name, args ?? new SystemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private System(string name, Input<string> id, SystemState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/system:System", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing System resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static System Get(string name, Input<string> id, SystemState? state = null, CustomResourceOptions? options = null)
        {
            return new System(name, id, state, options);
        }
    }

    public sealed class SystemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        /// </summary>
        [Input("dataSyncInterval")]
        public Input<int>? DataSyncInterval { get; set; }

        /// <summary>
        /// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        /// </summary>
        [Input("dynamicPeriodicInterval")]
        public Input<int>? DynamicPeriodicInterval { get; set; }

        /// <summary>
        /// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        /// </summary>
        [Input("iotHoldoff")]
        public Input<int>? IotHoldoff { get; set; }

        /// <summary>
        /// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        /// </summary>
        [Input("iotMacIdle")]
        public Input<int>? IotMacIdle { get; set; }

        /// <summary>
        /// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        /// </summary>
        [Input("iotScanInterval")]
        public Input<int>? IotScanInterval { get; set; }

        /// <summary>
        /// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        /// </summary>
        [Input("iotWeightThreshold")]
        public Input<int>? IotWeightThreshold { get; set; }

        /// <summary>
        /// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        /// </summary>
        [Input("nacPeriodicInterval")]
        public Input<int>? NacPeriodicInterval { get; set; }

        /// <summary>
        /// Maximum number of parallel processes (1 - 300, default = 1).
        /// </summary>
        [Input("parallelProcess")]
        public Input<int>? ParallelProcess { get; set; }

        /// <summary>
        /// Enable/disable parallel process override. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("parallelProcessOverride")]
        public Input<string>? ParallelProcessOverride { get; set; }

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

        public SystemArgs()
        {
        }
        public static new SystemArgs Empty => new SystemArgs();
    }

    public sealed class SystemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
        /// </summary>
        [Input("dataSyncInterval")]
        public Input<int>? DataSyncInterval { get; set; }

        /// <summary>
        /// Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
        /// </summary>
        [Input("dynamicPeriodicInterval")]
        public Input<int>? DynamicPeriodicInterval { get; set; }

        /// <summary>
        /// MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
        /// </summary>
        [Input("iotHoldoff")]
        public Input<int>? IotHoldoff { get; set; }

        /// <summary>
        /// MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
        /// </summary>
        [Input("iotMacIdle")]
        public Input<int>? IotMacIdle { get; set; }

        /// <summary>
        /// IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
        /// </summary>
        [Input("iotScanInterval")]
        public Input<int>? IotScanInterval { get; set; }

        /// <summary>
        /// MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
        /// </summary>
        [Input("iotWeightThreshold")]
        public Input<int>? IotWeightThreshold { get; set; }

        /// <summary>
        /// Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
        /// </summary>
        [Input("nacPeriodicInterval")]
        public Input<int>? NacPeriodicInterval { get; set; }

        /// <summary>
        /// Maximum number of parallel processes (1 - 300, default = 1).
        /// </summary>
        [Input("parallelProcess")]
        public Input<int>? ParallelProcess { get; set; }

        /// <summary>
        /// Enable/disable parallel process override. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("parallelProcessOverride")]
        public Input<string>? ParallelProcessOverride { get; set; }

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

        public SystemState()
        {
        }
        public static new SystemState Empty => new SystemState();
    }
}
