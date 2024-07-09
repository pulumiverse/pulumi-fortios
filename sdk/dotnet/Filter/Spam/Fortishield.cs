// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Spam
{
    /// <summary>
    /// Configure FortiGuard - AntiSpam. Applies to FortiOS Version `&lt;= 6.2.0`.
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
    ///     var trname = new Fortios.Filter.Spam.Fortishield("trname", new()
    ///     {
    ///         SpamSubmitForce = "enable",
    ///         SpamSubmitSrv = "www.nospammer.net",
    ///         SpamSubmitTxt2htm = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Spamfilter Fortishield can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/spam/fortishield:Fortishield labelname SpamfilterFortishield
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/spam/fortishield:Fortishield labelname SpamfilterFortishield
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:filter/spam/fortishield:Fortishield")]
    public partial class Fortishield : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("spamSubmitForce")]
        public Output<string> SpamSubmitForce { get; private set; } = null!;

        /// <summary>
        /// Hostname of the spam submission server.
        /// </summary>
        [Output("spamSubmitSrv")]
        public Output<string> SpamSubmitSrv { get; private set; } = null!;

        /// <summary>
        /// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("spamSubmitTxt2htm")]
        public Output<string> SpamSubmitTxt2htm { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Fortishield resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fortishield(string name, FortishieldArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:filter/spam/fortishield:Fortishield", name, args ?? new FortishieldArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fortishield(string name, Input<string> id, FortishieldState? state = null, CustomResourceOptions? options = null)
            : base("fortios:filter/spam/fortishield:Fortishield", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fortishield resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fortishield Get(string name, Input<string> id, FortishieldState? state = null, CustomResourceOptions? options = null)
        {
            return new Fortishield(name, id, state, options);
        }
    }

    public sealed class FortishieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamSubmitForce")]
        public Input<string>? SpamSubmitForce { get; set; }

        /// <summary>
        /// Hostname of the spam submission server.
        /// </summary>
        [Input("spamSubmitSrv")]
        public Input<string>? SpamSubmitSrv { get; set; }

        /// <summary>
        /// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamSubmitTxt2htm")]
        public Input<string>? SpamSubmitTxt2htm { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FortishieldArgs()
        {
        }
        public static new FortishieldArgs Empty => new FortishieldArgs();
    }

    public sealed class FortishieldState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamSubmitForce")]
        public Input<string>? SpamSubmitForce { get; set; }

        /// <summary>
        /// Hostname of the spam submission server.
        /// </summary>
        [Input("spamSubmitSrv")]
        public Input<string>? SpamSubmitSrv { get; set; }

        /// <summary>
        /// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("spamSubmitTxt2htm")]
        public Input<string>? SpamSubmitTxt2htm { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FortishieldState()
        {
        }
        public static new FortishieldState Empty => new FortishieldState();
    }
}
