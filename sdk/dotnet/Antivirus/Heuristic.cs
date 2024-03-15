// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Antivirus
{
    /// <summary>
    /// Configure global heuristic options. Applies to FortiOS Version `&lt;= 7.0.0`.
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
    ///     var trname = new Fortios.Antivirus.Heuristic("trname", new()
    ///     {
    ///         Mode = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Antivirus Heuristic can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:antivirus/heuristic:Heuristic labelname AntivirusHeuristic
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:antivirus/heuristic:Heuristic labelname AntivirusHeuristic
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:antivirus/heuristic:Heuristic")]
    public partial class Heuristic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Heuristic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Heuristic(string name, HeuristicArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:antivirus/heuristic:Heuristic", name, args ?? new HeuristicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Heuristic(string name, Input<string> id, HeuristicState? state = null, CustomResourceOptions? options = null)
            : base("fortios:antivirus/heuristic:Heuristic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Heuristic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Heuristic Get(string name, Input<string> id, HeuristicState? state = null, CustomResourceOptions? options = null)
        {
            return new Heuristic(name, id, state, options);
        }
    }

    public sealed class HeuristicArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public HeuristicArgs()
        {
        }
        public static new HeuristicArgs Empty => new HeuristicArgs();
    }

    public sealed class HeuristicState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public HeuristicState()
        {
        }
        public static new HeuristicState Empty => new HeuristicState();
    }
}