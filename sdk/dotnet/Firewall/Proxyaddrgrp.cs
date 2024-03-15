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
    /// Web proxy address group configuration.
    /// 
    /// ## Import
    /// 
    /// Firewall ProxyAddrgrp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/proxyaddrgrp:Proxyaddrgrp labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/proxyaddrgrp:Proxyaddrgrp labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/proxyaddrgrp:Proxyaddrgrp")]
    public partial class Proxyaddrgrp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Members of address group. The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.ProxyaddrgrpMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Address group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.ProxyaddrgrpTagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// Source or destination address group type. Valid values: `src`, `dst`.
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
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Proxyaddrgrp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Proxyaddrgrp(string name, ProxyaddrgrpArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/proxyaddrgrp:Proxyaddrgrp", name, args ?? new ProxyaddrgrpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Proxyaddrgrp(string name, Input<string> id, ProxyaddrgrpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/proxyaddrgrp:Proxyaddrgrp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Proxyaddrgrp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Proxyaddrgrp Get(string name, Input<string> id, ProxyaddrgrpState? state = null, CustomResourceOptions? options = null)
        {
            return new Proxyaddrgrp(name, id, state, options);
        }
    }

    public sealed class ProxyaddrgrpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("members", required: true)]
        private InputList<Inputs.ProxyaddrgrpMemberArgs>? _members;

        /// <summary>
        /// Members of address group. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddrgrpMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.ProxyaddrgrpMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("taggings")]
        private InputList<Inputs.ProxyaddrgrpTaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddrgrpTaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.ProxyaddrgrpTaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Source or destination address group type. Valid values: `src`, `dst`.
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
        /// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ProxyaddrgrpArgs()
        {
        }
        public static new ProxyaddrgrpArgs Empty => new ProxyaddrgrpArgs();
    }

    public sealed class ProxyaddrgrpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("members")]
        private InputList<Inputs.ProxyaddrgrpMemberGetArgs>? _members;

        /// <summary>
        /// Members of address group. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddrgrpMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.ProxyaddrgrpMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Address group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("taggings")]
        private InputList<Inputs.ProxyaddrgrpTaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.ProxyaddrgrpTaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.ProxyaddrgrpTaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Source or destination address group type. Valid values: `src`, `dst`.
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
        /// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ProxyaddrgrpState()
        {
        }
        public static new ProxyaddrgrpState Empty => new ProxyaddrgrpState();
    }
}