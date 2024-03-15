// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// Configure device groups. Applies to FortiOS Version `&lt;= 6.2.0`.
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
    ///     var trnames12 = new Fortios.User.Device("trnames12", new()
    ///     {
    ///         Alias = "user_devices2",
    ///         Category = "amazon-device",
    ///         Mac = "08:00:20:0a:1c:1d",
    ///         Type = "unknown",
    ///     });
    /// 
    ///     var trname = new Fortios.User.Devicegroup("trname", new()
    ///     {
    ///         Members = new[]
    ///         {
    ///             new Fortios.User.Inputs.DevicegroupMemberArgs
    ///             {
    ///                 Name = trnames12.Alias,
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
    /// User DeviceGroup can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/devicegroup:Devicegroup labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/devicegroup:Devicegroup labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/devicegroup:Devicegroup")]
    public partial class Devicegroup : global::Pulumi.CustomResource
    {
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
        /// Device group member. The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.DevicegroupMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Device group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.DevicegroupTagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Devicegroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Devicegroup(string name, DevicegroupArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:user/devicegroup:Devicegroup", name, args ?? new DevicegroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Devicegroup(string name, Input<string> id, DevicegroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/devicegroup:Devicegroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Devicegroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Devicegroup Get(string name, Input<string> id, DevicegroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Devicegroup(name, id, state, options);
        }
    }

    public sealed class DevicegroupArgs : global::Pulumi.ResourceArgs
    {
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

        [Input("members")]
        private InputList<Inputs.DevicegroupMemberArgs>? _members;

        /// <summary>
        /// Device group member. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.DevicegroupMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.DevicegroupMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Device group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("taggings")]
        private InputList<Inputs.DevicegroupTaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.DevicegroupTaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.DevicegroupTaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DevicegroupArgs()
        {
        }
        public static new DevicegroupArgs Empty => new DevicegroupArgs();
    }

    public sealed class DevicegroupState : global::Pulumi.ResourceArgs
    {
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

        [Input("members")]
        private InputList<Inputs.DevicegroupMemberGetArgs>? _members;

        /// <summary>
        /// Device group member. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.DevicegroupMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.DevicegroupMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Device group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("taggings")]
        private InputList<Inputs.DevicegroupTaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.DevicegroupTaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.DevicegroupTaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DevicegroupState()
        {
        }
        public static new DevicegroupState Empty => new DevicegroupState();
    }
}
