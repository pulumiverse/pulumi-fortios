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
    /// Configure FortiSwitch sFlow.
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
    ///     var trname = new Fortios.Switchcontroller.Sflow("trname", new()
    ///     {
    ///         CollectorIp = "0.0.0.0",
    ///         CollectorPort = 6343,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SwitchController Sflow can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/sflow:Sflow labelname SwitchControllerSflow
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/sflow:Sflow labelname SwitchControllerSflow
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/sflow:Sflow")]
    public partial class Sflow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Collector IP.
        /// </summary>
        [Output("collectorIp")]
        public Output<string> CollectorIp { get; private set; } = null!;

        /// <summary>
        /// SFlow collector port (0 - 65535).
        /// </summary>
        [Output("collectorPort")]
        public Output<int> CollectorPort { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Sflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sflow(string name, SflowArgs args, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/sflow:Sflow", name, args ?? new SflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sflow(string name, Input<string> id, SflowState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/sflow:Sflow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Sflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sflow Get(string name, Input<string> id, SflowState? state = null, CustomResourceOptions? options = null)
        {
            return new Sflow(name, id, state, options);
        }
    }

    public sealed class SflowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Collector IP.
        /// </summary>
        [Input("collectorIp", required: true)]
        public Input<string> CollectorIp { get; set; } = null!;

        /// <summary>
        /// SFlow collector port (0 - 65535).
        /// </summary>
        [Input("collectorPort")]
        public Input<int>? CollectorPort { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SflowArgs()
        {
        }
        public static new SflowArgs Empty => new SflowArgs();
    }

    public sealed class SflowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Collector IP.
        /// </summary>
        [Input("collectorIp")]
        public Input<string>? CollectorIp { get; set; }

        /// <summary>
        /// SFlow collector port (0 - 65535).
        /// </summary>
        [Input("collectorPort")]
        public Input<int>? CollectorPort { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SflowState()
        {
        }
        public static new SflowState Empty => new SflowState();
    }
}
