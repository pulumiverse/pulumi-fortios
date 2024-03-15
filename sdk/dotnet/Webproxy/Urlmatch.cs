// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Webproxy
{
    /// <summary>
    /// Exempt URLs from web proxy forwarding and caching.
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
    ///     var trname2 = new Fortios.Webproxy.Forwardserver("trname2", new()
    ///     {
    ///         AddrType = "fqdn",
    ///         Healthcheck = "disable",
    ///         Ip = "0.0.0.0",
    ///         Monitor = "http://www.google.com",
    ///         Port = 3128,
    ///         ServerDownOption = "block",
    ///     });
    /// 
    ///     var trname = new Fortios.Webproxy.Urlmatch("trname", new()
    ///     {
    ///         CacheExemption = "disable",
    ///         ForwardServer = trname2.Name,
    ///         Status = "enable",
    ///         UrlPattern = "/examples/servlet/*Servlet",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// WebProxy UrlMatch can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:webproxy/urlmatch:Urlmatch labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:webproxy/urlmatch:Urlmatch labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:webproxy/urlmatch:Urlmatch")]
    public partial class Urlmatch : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable exempting this URL pattern from caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("cacheExemption")]
        public Output<string> CacheExemption { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Forward server name.
        /// </summary>
        [Output("forwardServer")]
        public Output<string> ForwardServer { get; private set; } = null!;

        /// <summary>
        /// Configure a name for the URL to be exempted.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// URL pattern to be exempted from web proxy forwarding and caching.
        /// </summary>
        [Output("urlPattern")]
        public Output<string> UrlPattern { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Urlmatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Urlmatch(string name, UrlmatchArgs args, CustomResourceOptions? options = null)
            : base("fortios:webproxy/urlmatch:Urlmatch", name, args ?? new UrlmatchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Urlmatch(string name, Input<string> id, UrlmatchState? state = null, CustomResourceOptions? options = null)
            : base("fortios:webproxy/urlmatch:Urlmatch", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Urlmatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Urlmatch Get(string name, Input<string> id, UrlmatchState? state = null, CustomResourceOptions? options = null)
        {
            return new Urlmatch(name, id, state, options);
        }
    }

    public sealed class UrlmatchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable exempting this URL pattern from caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cacheExemption")]
        public Input<string>? CacheExemption { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Forward server name.
        /// </summary>
        [Input("forwardServer")]
        public Input<string>? ForwardServer { get; set; }

        /// <summary>
        /// Configure a name for the URL to be exempted.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// URL pattern to be exempted from web proxy forwarding and caching.
        /// </summary>
        [Input("urlPattern", required: true)]
        public Input<string> UrlPattern { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UrlmatchArgs()
        {
        }
        public static new UrlmatchArgs Empty => new UrlmatchArgs();
    }

    public sealed class UrlmatchState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable exempting this URL pattern from caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("cacheExemption")]
        public Input<string>? CacheExemption { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Forward server name.
        /// </summary>
        [Input("forwardServer")]
        public Input<string>? ForwardServer { get; set; }

        /// <summary>
        /// Configure a name for the URL to be exempted.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// URL pattern to be exempted from web proxy forwarding and caching.
        /// </summary>
        [Input("urlPattern")]
        public Input<string>? UrlPattern { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UrlmatchState()
        {
        }
        public static new UrlmatchState Empty => new UrlmatchState();
    }
}
