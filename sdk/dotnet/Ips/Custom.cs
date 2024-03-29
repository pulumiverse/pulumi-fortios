// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Ips
{
    /// <summary>
    /// Configure IPS custom signature.
    /// 
    /// ## Import
    /// 
    /// Ips Custom can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:ips/custom:Custom labelname {{tag}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:ips/custom:Custom labelname {{tag}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:ips/custom:Custom")]
    public partial class Custom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// Applications to be protected. Blank for all applications.
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// Protect client or server traffic.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Enable/disable logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("log")]
        public Output<string> Log { get; private set; } = null!;

        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("logPacket")]
        public Output<string> LogPacket { get; private set; } = null!;

        /// <summary>
        /// Operating system(s) that the signature protects. Blank for all operating systems.
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// Protocol(s) that the signature scans. Blank for all protocols.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Signature ID.
        /// </summary>
        [Output("ruleId")]
        public Output<int> RuleId { get; private set; } = null!;

        /// <summary>
        /// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        /// <summary>
        /// Signature name.
        /// </summary>
        [Output("sigName")]
        public Output<string> SigName { get; private set; } = null!;

        /// <summary>
        /// Custom signature enclosed in single quotes.
        /// </summary>
        [Output("signature")]
        public Output<string> Signature { get; private set; } = null!;

        /// <summary>
        /// Enable/disable this signature. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Signature tag.
        /// </summary>
        [Output("tag")]
        public Output<string> Tag { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Custom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Custom(string name, CustomArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:ips/custom:Custom", name, args ?? new CustomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Custom(string name, Input<string> id, CustomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:ips/custom:Custom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Custom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Custom Get(string name, Input<string> id, CustomState? state = null, CustomResourceOptions? options = null)
        {
            return new Custom(name, id, state, options);
        }
    }

    public sealed class CustomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Applications to be protected. Blank for all applications.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Protect client or server traffic.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        /// <summary>
        /// Operating system(s) that the signature protects. Blank for all operating systems.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Protocol(s) that the signature scans. Blank for all protocols.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Signature ID.
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Signature name.
        /// </summary>
        [Input("sigName")]
        public Input<string>? SigName { get; set; }

        /// <summary>
        /// Custom signature enclosed in single quotes.
        /// </summary>
        [Input("signature")]
        public Input<string>? Signature { get; set; }

        /// <summary>
        /// Enable/disable this signature. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Signature tag.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CustomArgs()
        {
        }
        public static new CustomArgs Empty => new CustomArgs();
    }

    public sealed class CustomState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default action (pass or block) for this signature. Valid values: `pass`, `block`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Applications to be protected. Blank for all applications.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Protect client or server traffic.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        /// <summary>
        /// Operating system(s) that the signature protects. Blank for all operating systems.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Protocol(s) that the signature scans. Blank for all protocols.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Signature ID.
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Signature name.
        /// </summary>
        [Input("sigName")]
        public Input<string>? SigName { get; set; }

        /// <summary>
        /// Custom signature enclosed in single quotes.
        /// </summary>
        [Input("signature")]
        public Input<string>? Signature { get; set; }

        /// <summary>
        /// Enable/disable this signature. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Signature tag.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public CustomState()
        {
        }
        public static new CustomState Empty => new CustomState();
    }
}
