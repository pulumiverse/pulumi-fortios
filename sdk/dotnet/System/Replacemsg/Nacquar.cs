// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Replacemsg
{
    /// <summary>
    /// Replacement messages.
    /// 
    /// ## Import
    /// 
    /// SystemReplacemsg NacQuar can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/replacemsg/nacquar:Nacquar labelname {{msg_type}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/replacemsg/nacquar:Nacquar labelname {{msg_type}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/replacemsg/nacquar:Nacquar")]
    public partial class Nacquar : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Message string.
        /// </summary>
        [Output("buffer")]
        public Output<string?> Buffer { get; private set; } = null!;

        /// <summary>
        /// Format flag.
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// Header flag. Valid values: `none`, `http`, `8bit`.
        /// </summary>
        [Output("header")]
        public Output<string> Header { get; private set; } = null!;

        /// <summary>
        /// Message type.
        /// </summary>
        [Output("msgType")]
        public Output<string> MsgType { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Nacquar resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Nacquar(string name, NacquarArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/replacemsg/nacquar:Nacquar", name, args ?? new NacquarArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Nacquar(string name, Input<string> id, NacquarState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/replacemsg/nacquar:Nacquar", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Nacquar resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Nacquar Get(string name, Input<string> id, NacquarState? state = null, CustomResourceOptions? options = null)
        {
            return new Nacquar(name, id, state, options);
        }
    }

    public sealed class NacquarArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Message string.
        /// </summary>
        [Input("buffer")]
        public Input<string>? Buffer { get; set; }

        /// <summary>
        /// Format flag.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Header flag. Valid values: `none`, `http`, `8bit`.
        /// </summary>
        [Input("header")]
        public Input<string>? Header { get; set; }

        /// <summary>
        /// Message type.
        /// </summary>
        [Input("msgType", required: true)]
        public Input<string> MsgType { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NacquarArgs()
        {
        }
        public static new NacquarArgs Empty => new NacquarArgs();
    }

    public sealed class NacquarState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Message string.
        /// </summary>
        [Input("buffer")]
        public Input<string>? Buffer { get; set; }

        /// <summary>
        /// Format flag.
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// Header flag. Valid values: `none`, `http`, `8bit`.
        /// </summary>
        [Input("header")]
        public Input<string>? Header { get; set; }

        /// <summary>
        /// Message type.
        /// </summary>
        [Input("msgType")]
        public Input<string>? MsgType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public NacquarState()
        {
        }
        public static new NacquarState Empty => new NacquarState();
    }
}
