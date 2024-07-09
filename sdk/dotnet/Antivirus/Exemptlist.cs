// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Antivirus
{
    /// <summary>
    /// Configure a list of hashes to be exempt from AV scanning. Applies to FortiOS Version `&gt;= 7.2.4`.
    /// 
    /// ## Import
    /// 
    /// Antivirus ExemptList can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:antivirus/exemptlist:Exemptlist labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:antivirus/exemptlist:Exemptlist labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:antivirus/exemptlist:Exemptlist")]
    public partial class Exemptlist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Hash value to be matched.
        /// </summary>
        [Output("hash")]
        public Output<string> Hash { get; private set; } = null!;

        /// <summary>
        /// Hash type. Valid values: `md5`, `sha1`, `sha256`.
        /// </summary>
        [Output("hashType")]
        public Output<string> HashType { get; private set; } = null!;

        /// <summary>
        /// Table entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable table entry. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Exemptlist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Exemptlist(string name, ExemptlistArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:antivirus/exemptlist:Exemptlist", name, args ?? new ExemptlistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Exemptlist(string name, Input<string> id, ExemptlistState? state = null, CustomResourceOptions? options = null)
            : base("fortios:antivirus/exemptlist:Exemptlist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Exemptlist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Exemptlist Get(string name, Input<string> id, ExemptlistState? state = null, CustomResourceOptions? options = null)
        {
            return new Exemptlist(name, id, state, options);
        }
    }

    public sealed class ExemptlistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Hash value to be matched.
        /// </summary>
        [Input("hash")]
        public Input<string>? Hash { get; set; }

        /// <summary>
        /// Hash type. Valid values: `md5`, `sha1`, `sha256`.
        /// </summary>
        [Input("hashType")]
        public Input<string>? HashType { get; set; }

        /// <summary>
        /// Table entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable table entry. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ExemptlistArgs()
        {
        }
        public static new ExemptlistArgs Empty => new ExemptlistArgs();
    }

    public sealed class ExemptlistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Hash value to be matched.
        /// </summary>
        [Input("hash")]
        public Input<string>? Hash { get; set; }

        /// <summary>
        /// Hash type. Valid values: `md5`, `sha1`, `sha256`.
        /// </summary>
        [Input("hashType")]
        public Input<string>? HashType { get; set; }

        /// <summary>
        /// Table entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable table entry. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ExemptlistState()
        {
        }
        public static new ExemptlistState Empty => new ExemptlistState();
    }
}
