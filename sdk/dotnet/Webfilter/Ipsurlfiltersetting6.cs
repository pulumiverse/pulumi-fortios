// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webfilter
{
    /// <summary>
    /// Configure IPS URL filter settings for IPv6.
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
    ///     var trname = new Fortios.Webfilter.Ipsurlfiltersetting6("trname", new()
    ///     {
    ///         Distance = 1,
    ///         Gateway6 = "::",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Webfilter IpsUrlfilterSetting6 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/ipsurlfiltersetting6:Ipsurlfiltersetting6 labelname WebfilterIpsUrlfilterSetting6
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:webfilter/ipsurlfiltersetting6:Ipsurlfiltersetting6 labelname WebfilterIpsUrlfilterSetting6
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webfilter/ipsurlfiltersetting6:Ipsurlfiltersetting6")]
    public partial class Ipsurlfiltersetting6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Interface for this route.
        /// </summary>
        [Output("device")]
        public Output<string> Device { get; private set; } = null!;

        /// <summary>
        /// Administrative distance (1 - 255) for this route.
        /// </summary>
        [Output("distance")]
        public Output<int> Distance { get; private set; } = null!;

        /// <summary>
        /// Gateway IPv6 address for this route.
        /// </summary>
        [Output("gateway6")]
        public Output<string> Gateway6 { get; private set; } = null!;

        /// <summary>
        /// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
        /// </summary>
        [Output("geoFilter")]
        public Output<string?> GeoFilter { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ipsurlfiltersetting6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipsurlfiltersetting6(string name, Ipsurlfiltersetting6Args? args = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/ipsurlfiltersetting6:Ipsurlfiltersetting6", name, args ?? new Ipsurlfiltersetting6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Ipsurlfiltersetting6(string name, Input<string> id, Ipsurlfiltersetting6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:webfilter/ipsurlfiltersetting6:Ipsurlfiltersetting6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipsurlfiltersetting6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipsurlfiltersetting6 Get(string name, Input<string> id, Ipsurlfiltersetting6State? state = null, CustomResourceOptions? options = null)
        {
            return new Ipsurlfiltersetting6(name, id, state, options);
        }
    }

    public sealed class Ipsurlfiltersetting6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface for this route.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance (1 - 255) for this route.
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Gateway IPv6 address for this route.
        /// </summary>
        [Input("gateway6")]
        public Input<string>? Gateway6 { get; set; }

        /// <summary>
        /// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
        /// </summary>
        [Input("geoFilter")]
        public Input<string>? GeoFilter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ipsurlfiltersetting6Args()
        {
        }
        public static new Ipsurlfiltersetting6Args Empty => new Ipsurlfiltersetting6Args();
    }

    public sealed class Ipsurlfiltersetting6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface for this route.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance (1 - 255) for this route.
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Gateway IPv6 address for this route.
        /// </summary>
        [Input("gateway6")]
        public Input<string>? Gateway6 { get; set; }

        /// <summary>
        /// Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
        /// </summary>
        [Input("geoFilter")]
        public Input<string>? GeoFilter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ipsurlfiltersetting6State()
        {
        }
        public static new Ipsurlfiltersetting6State Empty => new Ipsurlfiltersetting6State();
    }
}
