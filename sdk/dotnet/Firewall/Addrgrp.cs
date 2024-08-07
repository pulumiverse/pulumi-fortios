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
    /// Configure IPv4 address groups.
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
    ///     var trname1 = new Fortios.Firewall.Address("trname1", new()
    ///     {
    ///         AllowRouting = "disable",
    ///         CacheTtl = 0,
    ///         Color = 0,
    ///         EndIp = "255.0.0.0",
    ///         StartIp = "12.0.0.0",
    ///         Subnet = "12.0.0.0 255.0.0.0",
    ///         Type = "ipmask",
    ///         Visibility = "enable",
    ///         Wildcard = "12.0.0.0 255.0.0.0",
    ///     });
    /// 
    ///     var trname = new Fortios.Firewall.Addrgrp("trname", new()
    ///     {
    ///         AllowRouting = "disable",
    ///         Color = 0,
    ///         Exclude = "disable",
    ///         Visibility = "enable",
    ///         Members = new[]
    ///         {
    ///             new Fortios.Firewall.Inputs.AddrgrpMemberArgs
    ///             {
    ///                 Name = trname1.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Addrgrp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/addrgrp:Addrgrp labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/addrgrp:Addrgrp labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/addrgrp:Addrgrp")]
    public partial class Addrgrp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("allowRouting")]
        public Output<string> AllowRouting { get; private set; } = null!;

        /// <summary>
        /// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Color of icon on the GUI.
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
        /// Enable/disable address exclusion. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("exclude")]
        public Output<string> Exclude { get; private set; } = null!;

        /// <summary>
        /// Address exclusion member. The structure of `exclude_member` block is documented below.
        /// </summary>
        [Output("excludeMembers")]
        public Output<ImmutableArray<Outputs.AddrgrpExcludeMember>> ExcludeMembers { get; private set; } = null!;

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fabricObject")]
        public Output<string> FabricObject { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Address objects contained within the group. The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.AddrgrpMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Address group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.AddrgrpTagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// Address group type. Valid values: `default`, `folder`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Addrgrp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Addrgrp(string name, AddrgrpArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/addrgrp:Addrgrp", name, args ?? new AddrgrpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Addrgrp(string name, Input<string> id, AddrgrpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/addrgrp:Addrgrp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Addrgrp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Addrgrp Get(string name, Input<string> id, AddrgrpState? state = null, CustomResourceOptions? options = null)
        {
            return new Addrgrp(name, id, state, options);
        }
    }

    public sealed class AddrgrpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowRouting")]
        public Input<string>? AllowRouting { get; set; }

        /// <summary>
        /// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
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
        /// Enable/disable address exclusion. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("exclude")]
        public Input<string>? Exclude { get; set; }

        [Input("excludeMembers")]
        private InputList<Inputs.AddrgrpExcludeMemberArgs>? _excludeMembers;

        /// <summary>
        /// Address exclusion member. The structure of `exclude_member` block is documented below.
        /// </summary>
        public InputList<Inputs.AddrgrpExcludeMemberArgs> ExcludeMembers
        {
            get => _excludeMembers ?? (_excludeMembers = new InputList<Inputs.AddrgrpExcludeMemberArgs>());
            set => _excludeMembers = value;
        }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("members", required: true)]
        private InputList<Inputs.AddrgrpMemberArgs>? _members;

        /// <summary>
        /// Address objects contained within the group. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.AddrgrpMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.AddrgrpMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("taggings")]
        private InputList<Inputs.AddrgrpTaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.AddrgrpTaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.AddrgrpTaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Address group type. Valid values: `default`, `folder`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

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
        /// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public AddrgrpArgs()
        {
        }
        public static new AddrgrpArgs Empty => new AddrgrpArgs();
    }

    public sealed class AddrgrpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowRouting")]
        public Input<string>? AllowRouting { get; set; }

        /// <summary>
        /// Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
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
        /// Enable/disable address exclusion. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("exclude")]
        public Input<string>? Exclude { get; set; }

        [Input("excludeMembers")]
        private InputList<Inputs.AddrgrpExcludeMemberGetArgs>? _excludeMembers;

        /// <summary>
        /// Address exclusion member. The structure of `exclude_member` block is documented below.
        /// </summary>
        public InputList<Inputs.AddrgrpExcludeMemberGetArgs> ExcludeMembers
        {
            get => _excludeMembers ?? (_excludeMembers = new InputList<Inputs.AddrgrpExcludeMemberGetArgs>());
            set => _excludeMembers = value;
        }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        [Input("members")]
        private InputList<Inputs.AddrgrpMemberGetArgs>? _members;

        /// <summary>
        /// Address objects contained within the group. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.AddrgrpMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.AddrgrpMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("taggings")]
        private InputList<Inputs.AddrgrpTaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.AddrgrpTaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.AddrgrpTaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Address group type. Valid values: `default`, `folder`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

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
        /// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public AddrgrpState()
        {
        }
        public static new AddrgrpState Empty => new AddrgrpState();
    }
}
