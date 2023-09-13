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
    /// Configure FortiAP regions (for floor plans and maps).
    /// 
    /// ## Import
    /// 
    /// WirelessController Region can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:wirelesscontroller/region:Region labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:wirelesscontroller/region:Region labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wirelesscontroller/region:Region")]
    public partial class Region : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comments.
        /// </summary>
        [Output("comments")]
        public Output<string> Comments { get; private set; } = null!;

        /// <summary>
        /// Region image grayscale. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("grayscale")]
        public Output<string> Grayscale { get; private set; } = null!;

        /// <summary>
        /// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
        /// </summary>
        [Output("imageType")]
        public Output<string> ImageType { get; private set; } = null!;

        /// <summary>
        /// FortiAP region name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Region image opacity (0 - 100).
        /// </summary>
        [Output("opacity")]
        public Output<int> Opacity { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Region resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Region(string name, RegionArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/region:Region", name, args ?? new RegionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Region(string name, Input<string> id, RegionState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wirelesscontroller/region:Region", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Region resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Region Get(string name, Input<string> id, RegionState? state = null, CustomResourceOptions? options = null)
        {
            return new Region(name, id, state, options);
        }
    }

    public sealed class RegionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Region image grayscale. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("grayscale")]
        public Input<string>? Grayscale { get; set; }

        /// <summary>
        /// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
        /// </summary>
        [Input("imageType")]
        public Input<string>? ImageType { get; set; }

        /// <summary>
        /// FortiAP region name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Region image opacity (0 - 100).
        /// </summary>
        [Input("opacity")]
        public Input<int>? Opacity { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RegionArgs()
        {
        }
        public static new RegionArgs Empty => new RegionArgs();
    }

    public sealed class RegionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comments.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Region image grayscale. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("grayscale")]
        public Input<string>? Grayscale { get; set; }

        /// <summary>
        /// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
        /// </summary>
        [Input("imageType")]
        public Input<string>? ImageType { get; set; }

        /// <summary>
        /// FortiAP region name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Region image opacity (0 - 100).
        /// </summary>
        [Input("opacity")]
        public Input<int>? Opacity { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RegionState()
        {
        }
        public static new RegionState Empty => new RegionState();
    }
}
