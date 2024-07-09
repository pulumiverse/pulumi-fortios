// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure multicast addresses.
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
    ///     var trname = new Fortios.Firewall.Multicastaddress("trname", new()
    ///     {
    ///         Color = 0,
    ///         EndIp = "224.0.0.22",
    ///         StartIp = "224.0.0.11",
    ///         Subnet = "224.0.0.11 224.0.0.22",
    ///         Type = "multicastrange",
    ///         Visibility = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall MulticastAddress can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/multicastaddress:Multicastaddress labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/multicastaddress:Multicastaddress labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/multicastaddress:Multicastaddress")]
    public partial class Multicastaddress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
        /// </summary>
        [Output("associatedInterface")]
        public Output<string> AssociatedInterface { get; private set; } = null!;

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address.
        /// </summary>
        [Output("endIp")]
        public Output<string> EndIp { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Multicast address name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address.
        /// </summary>
        [Output("startIp")]
        public Output<string> StartIp { get; private set; } = null!;

        /// <summary>
        /// Broadcast address and subnet.
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.MulticastaddressTagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable visibility of the multicast address on the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Multicastaddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Multicastaddress(string name, MulticastaddressArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/multicastaddress:Multicastaddress", name, args ?? new MulticastaddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Multicastaddress(string name, Input<string> id, MulticastaddressState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/multicastaddress:Multicastaddress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Multicastaddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Multicastaddress Get(string name, Input<string> id, MulticastaddressState? state = null, CustomResourceOptions? options = null)
        {
            return new Multicastaddress(name, id, state, options);
        }
    }

    public sealed class MulticastaddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
        /// </summary>
        [Input("associatedInterface")]
        public Input<string>? AssociatedInterface { get; set; }

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Multicast address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// Broadcast address and subnet.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("taggings")]
        private InputList<Inputs.MulticastaddressTaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.MulticastaddressTaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.MulticastaddressTaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable visibility of the multicast address on the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public MulticastaddressArgs()
        {
        }
        public static new MulticastaddressArgs Empty => new MulticastaddressArgs();
    }

    public sealed class MulticastaddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
        /// </summary>
        [Input("associatedInterface")]
        public Input<string>? AssociatedInterface { get; set; }

        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Multicast address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// Broadcast address and subnet.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        [Input("taggings")]
        private InputList<Inputs.MulticastaddressTaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.MulticastaddressTaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.MulticastaddressTaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable visibility of the multicast address on the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public MulticastaddressState()
        {
        }
        public static new MulticastaddressState Empty => new MulticastaddressState();
    }
}
