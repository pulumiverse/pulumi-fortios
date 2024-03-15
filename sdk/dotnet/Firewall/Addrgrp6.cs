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
    /// Configure IPv6 address groups.
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
    ///     var trname1 = new Fortios.Firewall.Address6("trname1", new()
    ///     {
    ///         CacheTtl = 0,
    ///         Color = 0,
    ///         EndIp = "::",
    ///         Host = "",
    ///         HostType = "any",
    ///         Ip6 = "fdff:ffff::/120",
    ///         StartIp = "",
    ///         Type = "ipprefix",
    ///         Visibility = "enable",
    ///     });
    /// 
    ///     var trname = new Fortios.Firewall.Addrgrp6("trname", new()
    ///     {
    ///         Color = 0,
    ///         Visibility = "enable",
    ///         Members = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.Addrgrp6MemberArgs
    ///             {
    ///                 Name = trname1.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall Addrgrp6 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/addrgrp6:Addrgrp6 labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/addrgrp6:Addrgrp6 labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/addrgrp6:Addrgrp6")]
    public partial class Addrgrp6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
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
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fabricObject")]
        public Output<string> FabricObject { get; private set; } = null!;

        /// <summary>
        /// Address objects contained within the group. The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.Addrgrp6Member>> Members { get; private set; } = null!;

        /// <summary>
        /// IPv6 address group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.Addrgrp6Tagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Addrgrp6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Addrgrp6(string name, Addrgrp6Args args, CustomResourceOptions? options = null)
            : base("fortios:firewall/addrgrp6:Addrgrp6", name, args ?? new Addrgrp6Args(), MakeResourceOptions(options, ""))
        {
        }

        private Addrgrp6(string name, Input<string> id, Addrgrp6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/addrgrp6:Addrgrp6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Addrgrp6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Addrgrp6 Get(string name, Input<string> id, Addrgrp6State? state = null, CustomResourceOptions? options = null)
        {
            return new Addrgrp6(name, id, state, options);
        }
    }

    public sealed class Addrgrp6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
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
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        [Input("members", required: true)]
        private InputList<Inputs.Addrgrp6MemberArgs>? _members;

        /// <summary>
        /// Address objects contained within the group. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.Addrgrp6MemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.Addrgrp6MemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// IPv6 address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("taggings")]
        private InputList<Inputs.Addrgrp6TaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.Addrgrp6TaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.Addrgrp6TaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public Addrgrp6Args()
        {
        }
        public static new Addrgrp6Args Empty => new Addrgrp6Args();
    }

    public sealed class Addrgrp6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
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
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        [Input("members")]
        private InputList<Inputs.Addrgrp6MemberGetArgs>? _members;

        /// <summary>
        /// Address objects contained within the group. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.Addrgrp6MemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.Addrgrp6MemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// IPv6 address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("taggings")]
        private InputList<Inputs.Addrgrp6TaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.Addrgrp6TaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.Addrgrp6TaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public Addrgrp6State()
        {
        }
        public static new Addrgrp6State Empty => new Addrgrp6State();
    }
}
