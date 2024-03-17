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
    /// Configure IPv6 neighbor cache table.
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
    ///     var trname = new Fortios.System.Ipv6neighborcache("trname", new()
    ///     {
    ///         Fosid = 1,
    ///         Interface = "port2",
    ///         Ipv6 = "fe80::b11a:5ae3:198:ba1c",
    ///         Mac = "00:00:00:00:00:00",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System Ipv6NeighborCache can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/ipv6neighborcache:Ipv6neighborcache labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/ipv6neighborcache:Ipv6neighborcache labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/ipv6neighborcache:Ipv6neighborcache")]
    public partial class Ipv6neighborcache : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unique integer ID of the entry.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Output("ipv6")]
        public Output<string> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// MAC address (format: xx:xx:xx:xx:xx:xx).
        /// </summary>
        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv6neighborcache resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv6neighborcache(string name, Ipv6neighborcacheArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/ipv6neighborcache:Ipv6neighborcache", name, args ?? new Ipv6neighborcacheArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv6neighborcache(string name, Input<string> id, Ipv6neighborcacheState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/ipv6neighborcache:Ipv6neighborcache", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv6neighborcache resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv6neighborcache Get(string name, Input<string> id, Ipv6neighborcacheState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv6neighborcache(name, id, state, options);
        }
    }

    public sealed class Ipv6neighborcacheArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique integer ID of the entry.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        /// <summary>
        /// IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Input("ipv6", required: true)]
        public Input<string> Ipv6 { get; set; } = null!;

        /// <summary>
        /// MAC address (format: xx:xx:xx:xx:xx:xx).
        /// </summary>
        [Input("mac", required: true)]
        public Input<string> Mac { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ipv6neighborcacheArgs()
        {
        }
        public static new Ipv6neighborcacheArgs Empty => new Ipv6neighborcacheArgs();
    }

    public sealed class Ipv6neighborcacheState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique integer ID of the entry.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Select the associated interface name from available options.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// MAC address (format: xx:xx:xx:xx:xx:xx).
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Ipv6neighborcacheState()
        {
        }
        public static new Ipv6neighborcacheState Empty => new Ipv6neighborcacheState();
    }
}
