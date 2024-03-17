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
    /// Configure interrupt affinity.
    /// 
    /// ## Import
    /// 
    /// System AffinityInterrupt can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/affinityinterrupt:Affinityinterrupt labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/affinityinterrupt:Affinityinterrupt labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/affinityinterrupt:Affinityinterrupt")]
    public partial class Affinityinterrupt : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        /// </summary>
        [Output("affinityCpumask")]
        public Output<string> AffinityCpumask { get; private set; } = null!;

        /// <summary>
        /// ID of the interrupt affinity setting.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Interrupt name.
        /// </summary>
        [Output("interrupt")]
        public Output<string> Interrupt { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Affinityinterrupt resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Affinityinterrupt(string name, AffinityinterruptArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/affinityinterrupt:Affinityinterrupt", name, args ?? new AffinityinterruptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Affinityinterrupt(string name, Input<string> id, AffinityinterruptState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/affinityinterrupt:Affinityinterrupt", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Affinityinterrupt resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Affinityinterrupt Get(string name, Input<string> id, AffinityinterruptState? state = null, CustomResourceOptions? options = null)
        {
            return new Affinityinterrupt(name, id, state, options);
        }
    }

    public sealed class AffinityinterruptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        /// </summary>
        [Input("affinityCpumask", required: true)]
        public Input<string> AffinityCpumask { get; set; } = null!;

        /// <summary>
        /// ID of the interrupt affinity setting.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Interrupt name.
        /// </summary>
        [Input("interrupt", required: true)]
        public Input<string> Interrupt { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AffinityinterruptArgs()
        {
        }
        public static new AffinityinterruptArgs Empty => new AffinityinterruptArgs();
    }

    public sealed class AffinityinterruptState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
        /// </summary>
        [Input("affinityCpumask")]
        public Input<string>? AffinityCpumask { get; set; }

        /// <summary>
        /// ID of the interrupt affinity setting.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Interrupt name.
        /// </summary>
        [Input("interrupt")]
        public Input<string>? Interrupt { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AffinityinterruptState()
        {
        }
        public static new AffinityinterruptState Empty => new AffinityinterruptState();
    }
}
