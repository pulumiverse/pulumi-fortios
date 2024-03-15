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
    /// Configure IPv4 virtual IP groups.
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
    ///     var trname1 = new Fortios.Firewall.Vip("trname1", new()
    ///     {
    ///         Extintf = "any",
    ///         Extport = "0-65535",
    ///         Extip = "2.0.0.1-2.0.0.4",
    ///         Mappedips = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.VipMappedipArgs
    ///             {
    ///                 Range = "3.0.0.0-3.0.0.3",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var trname = new Fortios.Firewall.Vipgrp("trname", new()
    ///     {
    ///         Color = 0,
    ///         Interface = "any",
    ///         Members = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.VipgrpMemberArgs
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
    /// Firewall Vipgrp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/vipgrp:Vipgrp labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/vipgrp:Vipgrp labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/vipgrp:Vipgrp")]
    public partial class Vipgrp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// interface
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.VipgrpMember>> Members { get; private set; } = null!;

        /// <summary>
        /// VIP group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// Create a Vipgrp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vipgrp(string name, VipgrpArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/vipgrp:Vipgrp", name, args ?? new VipgrpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vipgrp(string name, Input<string> id, VipgrpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/vipgrp:Vipgrp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vipgrp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vipgrp Get(string name, Input<string> id, VipgrpState? state = null, CustomResourceOptions? options = null)
        {
            return new Vipgrp(name, id, state, options);
        }
    }

    public sealed class VipgrpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// interface
        /// </summary>
        [Input("interface", required: true)]
        public Input<string> Interface { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<Inputs.VipgrpMemberArgs>? _members;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.VipgrpMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.VipgrpMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// VIP group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public VipgrpArgs()
        {
        }
        public static new VipgrpArgs Empty => new VipgrpArgs();
    }

    public sealed class VipgrpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// interface
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        [Input("members")]
        private InputList<Inputs.VipgrpMemberGetArgs>? _members;

        /// <summary>
        /// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.VipgrpMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.VipgrpMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// VIP group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public VipgrpState()
        {
        }
        public static new VipgrpState Empty => new VipgrpState();
    }
}