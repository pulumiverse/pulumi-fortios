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
    /// Configure FortiSwitch SNMP trap threshold values globally. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// SwitchController SnmpTrapThreshold can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold labelname SwitchControllerSnmpTrapThreshold
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold labelname SwitchControllerSnmpTrapThreshold
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold")]
    public partial class Snmptrapthreshold : global::Pulumi.CustomResource
    {
        /// <summary>
        /// CPU usage when trap is sent.
        /// </summary>
        [Output("trapHighCpuThreshold")]
        public Output<int> TrapHighCpuThreshold { get; private set; } = null!;

        /// <summary>
        /// Log disk usage when trap is sent.
        /// </summary>
        [Output("trapLogFullThreshold")]
        public Output<int> TrapLogFullThreshold { get; private set; } = null!;

        /// <summary>
        /// Memory usage when trap is sent.
        /// </summary>
        [Output("trapLowMemoryThreshold")]
        public Output<int> TrapLowMemoryThreshold { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Snmptrapthreshold resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snmptrapthreshold(string name, SnmptrapthresholdArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold", name, args ?? new SnmptrapthresholdArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Snmptrapthreshold(string name, Input<string> id, SnmptrapthresholdState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/snmptrapthreshold:Snmptrapthreshold", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snmptrapthreshold resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snmptrapthreshold Get(string name, Input<string> id, SnmptrapthresholdState? state = null, CustomResourceOptions? options = null)
        {
            return new Snmptrapthreshold(name, id, state, options);
        }
    }

    public sealed class SnmptrapthresholdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CPU usage when trap is sent.
        /// </summary>
        [Input("trapHighCpuThreshold")]
        public Input<int>? TrapHighCpuThreshold { get; set; }

        /// <summary>
        /// Log disk usage when trap is sent.
        /// </summary>
        [Input("trapLogFullThreshold")]
        public Input<int>? TrapLogFullThreshold { get; set; }

        /// <summary>
        /// Memory usage when trap is sent.
        /// </summary>
        [Input("trapLowMemoryThreshold")]
        public Input<int>? TrapLowMemoryThreshold { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SnmptrapthresholdArgs()
        {
        }
        public static new SnmptrapthresholdArgs Empty => new SnmptrapthresholdArgs();
    }

    public sealed class SnmptrapthresholdState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CPU usage when trap is sent.
        /// </summary>
        [Input("trapHighCpuThreshold")]
        public Input<int>? TrapHighCpuThreshold { get; set; }

        /// <summary>
        /// Log disk usage when trap is sent.
        /// </summary>
        [Input("trapLogFullThreshold")]
        public Input<int>? TrapLogFullThreshold { get; set; }

        /// <summary>
        /// Memory usage when trap is sent.
        /// </summary>
        [Input("trapLowMemoryThreshold")]
        public Input<int>? TrapLowMemoryThreshold { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SnmptrapthresholdState()
        {
        }
        public static new SnmptrapthresholdState Empty => new SnmptrapthresholdState();
    }
}
