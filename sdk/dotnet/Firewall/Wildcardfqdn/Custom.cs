// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Wildcardfqdn
{
    /// <summary>
    /// Config global/VDOM Wildcard FQDN address.
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
    ///     var trname = new Fortios.Firewall.Wildcardfqdn.Custom("trname", new()
    ///     {
    ///         Color = 0,
    ///         Visibility = "enable",
    ///         WildcardFqdn = "*.go.google.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallWildcardFqdn Custom can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/wildcardfqdn/custom:Custom labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/wildcardfqdn/custom:Custom labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/wildcardfqdn/custom:Custom")]
    public partial class Custom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// GUI icon color.
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Address name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// Enable/disable address visibility. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;

        /// <summary>
        /// Wildcard FQDN.
        /// </summary>
        [Output("wildcardFqdn")]
        public Output<string> WildcardFqdn { get; private set; } = null!;


        /// <summary>
        /// Create a Custom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Custom(string name, CustomArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/wildcardfqdn/custom:Custom", name, args ?? new CustomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Custom(string name, Input<string> id, CustomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/wildcardfqdn/custom:Custom", name, state, MakeResourceOptions(options, id))
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
        /// GUI icon color.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// Enable/disable address visibility. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// Wildcard FQDN.
        /// </summary>
        [Input("wildcardFqdn")]
        public Input<string>? WildcardFqdn { get; set; }

        public CustomArgs()
        {
        }
        public static new CustomArgs Empty => new CustomArgs();
    }

    public sealed class CustomState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// GUI icon color.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// Enable/disable address visibility. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// Wildcard FQDN.
        /// </summary>
        [Input("wildcardFqdn")]
        public Input<string>? WildcardFqdn { get; set; }

        public CustomState()
        {
        }
        public static new CustomState Empty => new CustomState();
    }
}
