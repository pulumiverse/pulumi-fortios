// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Dlp
{
    /// <summary>
    /// Configure predefined data type used by DLP blocking. Applies to FortiOS Version `&gt;= 7.2.0`.
    /// 
    /// ## Import
    /// 
    /// Dlp DataType can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:dlp/datatype:Datatype")]
    public partial class Datatype : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Number of characters to obtain in advance for verification (1 - 255, default = 1).
        /// </summary>
        [Output("lookAhead")]
        public Output<int> LookAhead { get; private set; } = null!;

        /// <summary>
        /// Number of characters required to save for verification (1 - 255, default = 1).
        /// </summary>
        [Output("lookBack")]
        public Output<int> LookBack { get; private set; } = null!;

        /// <summary>
        /// Name of table containing the data type.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Regular expression pattern string without look around.
        /// </summary>
        [Output("pattern")]
        public Output<string> Pattern { get; private set; } = null!;

        /// <summary>
        /// Template to transform user input to a pattern using capture group from 'pattern'.
        /// </summary>
        [Output("transform")]
        public Output<string> Transform { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Regular expression pattern string used to verify the data type.
        /// </summary>
        [Output("verify")]
        public Output<string> Verify { get; private set; } = null!;

        /// <summary>
        /// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("verifyTransformedPattern")]
        public Output<string> VerifyTransformedPattern { get; private set; } = null!;


        /// <summary>
        /// Create a Datatype resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Datatype(string name, DatatypeArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:dlp/datatype:Datatype", name, args ?? new DatatypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Datatype(string name, Input<string> id, DatatypeState? state = null, CustomResourceOptions? options = null)
            : base("fortios:dlp/datatype:Datatype", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Datatype resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Datatype Get(string name, Input<string> id, DatatypeState? state = null, CustomResourceOptions? options = null)
        {
            return new Datatype(name, id, state, options);
        }
    }

    public sealed class DatatypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Number of characters to obtain in advance for verification (1 - 255, default = 1).
        /// </summary>
        [Input("lookAhead")]
        public Input<int>? LookAhead { get; set; }

        /// <summary>
        /// Number of characters required to save for verification (1 - 255, default = 1).
        /// </summary>
        [Input("lookBack")]
        public Input<int>? LookBack { get; set; }

        /// <summary>
        /// Name of table containing the data type.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Regular expression pattern string without look around.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Template to transform user input to a pattern using capture group from 'pattern'.
        /// </summary>
        [Input("transform")]
        public Input<string>? Transform { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Regular expression pattern string used to verify the data type.
        /// </summary>
        [Input("verify")]
        public Input<string>? Verify { get; set; }

        /// <summary>
        /// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("verifyTransformedPattern")]
        public Input<string>? VerifyTransformedPattern { get; set; }

        public DatatypeArgs()
        {
        }
        public static new DatatypeArgs Empty => new DatatypeArgs();
    }

    public sealed class DatatypeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional comments.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Number of characters to obtain in advance for verification (1 - 255, default = 1).
        /// </summary>
        [Input("lookAhead")]
        public Input<int>? LookAhead { get; set; }

        /// <summary>
        /// Number of characters required to save for verification (1 - 255, default = 1).
        /// </summary>
        [Input("lookBack")]
        public Input<int>? LookBack { get; set; }

        /// <summary>
        /// Name of table containing the data type.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Regular expression pattern string without look around.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// Template to transform user input to a pattern using capture group from 'pattern'.
        /// </summary>
        [Input("transform")]
        public Input<string>? Transform { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Regular expression pattern string used to verify the data type.
        /// </summary>
        [Input("verify")]
        public Input<string>? Verify { get; set; }

        /// <summary>
        /// Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("verifyTransformedPattern")]
        public Input<string>? VerifyTransformedPattern { get; set; }

        public DatatypeState()
        {
        }
        public static new DatatypeState Empty => new DatatypeState();
    }
}
