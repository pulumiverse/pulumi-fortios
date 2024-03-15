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
    /// Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
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
    ///     var trname = new Fortios.System.Geoipoverride("trname", new()
    ///     {
    ///         Description = "TEST for country",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System GeoipOverride can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/geoipoverride:Geoipoverride labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/geoipoverride:Geoipoverride labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/geoipoverride:Geoipoverride")]
    public partial class Geoipoverride : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Two character Country ID code.
        /// </summary>
        [Output("countryId")]
        public Output<string> CountryId { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        /// </summary>
        [Output("ip6Ranges")]
        public Output<ImmutableArray<Outputs.GeoipoverrideIp6Range>> Ip6Ranges { get; private set; } = null!;

        /// <summary>
        /// Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        /// </summary>
        [Output("ipRanges")]
        public Output<ImmutableArray<Outputs.GeoipoverrideIpRange>> IpRanges { get; private set; } = null!;

        /// <summary>
        /// Location name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Geoipoverride resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Geoipoverride(string name, GeoipoverrideArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/geoipoverride:Geoipoverride", name, args ?? new GeoipoverrideArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Geoipoverride(string name, Input<string> id, GeoipoverrideState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/geoipoverride:Geoipoverride", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Geoipoverride resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Geoipoverride Get(string name, Input<string> id, GeoipoverrideState? state = null, CustomResourceOptions? options = null)
        {
            return new Geoipoverride(name, id, state, options);
        }
    }

    public sealed class GeoipoverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Two character Country ID code.
        /// </summary>
        [Input("countryId")]
        public Input<string>? CountryId { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("ip6Ranges")]
        private InputList<Inputs.GeoipoverrideIp6RangeArgs>? _ip6Ranges;

        /// <summary>
        /// Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        /// </summary>
        public InputList<Inputs.GeoipoverrideIp6RangeArgs> Ip6Ranges
        {
            get => _ip6Ranges ?? (_ip6Ranges = new InputList<Inputs.GeoipoverrideIp6RangeArgs>());
            set => _ip6Ranges = value;
        }

        [Input("ipRanges")]
        private InputList<Inputs.GeoipoverrideIpRangeArgs>? _ipRanges;

        /// <summary>
        /// Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        /// </summary>
        public InputList<Inputs.GeoipoverrideIpRangeArgs> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<Inputs.GeoipoverrideIpRangeArgs>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// Location name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GeoipoverrideArgs()
        {
        }
        public static new GeoipoverrideArgs Empty => new GeoipoverrideArgs();
    }

    public sealed class GeoipoverrideState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Two character Country ID code.
        /// </summary>
        [Input("countryId")]
        public Input<string>? CountryId { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("ip6Ranges")]
        private InputList<Inputs.GeoipoverrideIp6RangeGetArgs>? _ip6Ranges;

        /// <summary>
        /// Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
        /// </summary>
        public InputList<Inputs.GeoipoverrideIp6RangeGetArgs> Ip6Ranges
        {
            get => _ip6Ranges ?? (_ip6Ranges = new InputList<Inputs.GeoipoverrideIp6RangeGetArgs>());
            set => _ip6Ranges = value;
        }

        [Input("ipRanges")]
        private InputList<Inputs.GeoipoverrideIpRangeGetArgs>? _ipRanges;

        /// <summary>
        /// Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
        /// </summary>
        public InputList<Inputs.GeoipoverrideIpRangeGetArgs> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<Inputs.GeoipoverrideIpRangeGetArgs>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// Location name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GeoipoverrideState()
        {
        }
        public static new GeoipoverrideState Empty => new GeoipoverrideState();
    }
}