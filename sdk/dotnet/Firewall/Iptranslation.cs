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
    /// Configure firewall IP-translation.
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
    ///     var trname = new Fortios.Firewall.Iptranslation("trname", new()
    ///     {
    ///         Endip = "2.2.2.2",
    ///         MapStartip = "0.0.0.0",
    ///         Startip = "1.1.1.1",
    ///         Transid = 1,
    ///         Type = "SCTP",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall IpTranslation can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/iptranslation:Iptranslation labelname {{transid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/iptranslation:Iptranslation labelname {{transid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/iptranslation:Iptranslation")]
    public partial class Iptranslation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Output("endip")]
        public Output<string> Endip { get; private set; } = null!;

        /// <summary>
        /// Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Output("mapStartip")]
        public Output<string> MapStartip { get; private set; } = null!;

        /// <summary>
        /// First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Output("startip")]
        public Output<string> Startip { get; private set; } = null!;

        /// <summary>
        /// IP translation ID.
        /// </summary>
        [Output("transid")]
        public Output<int> Transid { get; private set; } = null!;

        /// <summary>
        /// IP translation type (option: SCTP). Valid values: `SCTP`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Iptranslation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Iptranslation(string name, IptranslationArgs args, CustomResourceOptions? options = null)
            : base("fortios:firewall/iptranslation:Iptranslation", name, args ?? new IptranslationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Iptranslation(string name, Input<string> id, IptranslationState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/iptranslation:Iptranslation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Iptranslation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Iptranslation Get(string name, Input<string> id, IptranslationState? state = null, CustomResourceOptions? options = null)
        {
            return new Iptranslation(name, id, state, options);
        }
    }

    public sealed class IptranslationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Input("endip", required: true)]
        public Input<string> Endip { get; set; } = null!;

        /// <summary>
        /// Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Input("mapStartip", required: true)]
        public Input<string> MapStartip { get; set; } = null!;

        /// <summary>
        /// First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Input("startip", required: true)]
        public Input<string> Startip { get; set; } = null!;

        /// <summary>
        /// IP translation ID.
        /// </summary>
        [Input("transid")]
        public Input<int>? Transid { get; set; }

        /// <summary>
        /// IP translation type (option: SCTP). Valid values: `SCTP`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IptranslationArgs()
        {
        }
        public static new IptranslationArgs Empty => new IptranslationArgs();
    }

    public sealed class IptranslationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Input("endip")]
        public Input<string>? Endip { get; set; }

        /// <summary>
        /// Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Input("mapStartip")]
        public Input<string>? MapStartip { get; set; }

        /// <summary>
        /// First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
        /// </summary>
        [Input("startip")]
        public Input<string>? Startip { get; set; }

        /// <summary>
        /// IP translation ID.
        /// </summary>
        [Input("transid")]
        public Input<int>? Transid { get; set; }

        /// <summary>
        /// IP translation type (option: SCTP). Valid values: `SCTP`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IptranslationState()
        {
        }
        public static new IptranslationState Empty => new IptranslationState();
    }
}