// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Dlp
{
    /// <summary>
    /// Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// Dlp Sensitivity can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:dlp/sensitivity:Sensitivity labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:dlp/sensitivity:Sensitivity labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:dlp/sensitivity:Sensitivity")]
    public partial class Sensitivity : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DLP Sensitivity Levels.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Sensitivity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sensitivity(string name, SensitivityArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:dlp/sensitivity:Sensitivity", name, args ?? new SensitivityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sensitivity(string name, Input<string> id, SensitivityState? state = null, CustomResourceOptions? options = null)
            : base("fortios:dlp/sensitivity:Sensitivity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Sensitivity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sensitivity Get(string name, Input<string> id, SensitivityState? state = null, CustomResourceOptions? options = null)
        {
            return new Sensitivity(name, id, state, options);
        }
    }

    public sealed class SensitivityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DLP Sensitivity Levels.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SensitivityArgs()
        {
        }
        public static new SensitivityArgs Empty => new SensitivityArgs();
    }

    public sealed class SensitivityState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DLP Sensitivity Levels.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SensitivityState()
        {
        }
        public static new SensitivityState Empty => new SensitivityState();
    }
}
