// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web
{
    /// <summary>
    /// Configure FortiGuard Web Filter administrative overrides.
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
    ///     var trname = new Fortios.Filter.Web.Override("trname", new()
    ///     {
    ///         Expires = "2021/03/16 19:18:25",
    ///         Fosid = 1,
    ///         Ip = "69.101.119.0",
    ///         Ip6 = "4565:7700::",
    ///         NewProfile = "monitor-all",
    ///         OldProfile = "default",
    ///         Scope = "user",
    ///         Status = "disable",
    ///         User = "Eew",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Webfilter Override can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/web/override:Override labelname {{fosid}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/web/override:Override labelname {{fosid}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:filter/web/override:Override")]
    public partial class Override : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        /// </summary>
        [Output("expires")]
        public Output<string> Expires { get; private set; } = null!;

        /// <summary>
        /// Override rule ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Initiating user of override (read-only setting).
        /// </summary>
        [Output("initiator")]
        public Output<string> Initiator { get; private set; } = null!;

        /// <summary>
        /// IPv4 address which the override applies.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// IPv6 address which the override applies.
        /// </summary>
        [Output("ip6")]
        public Output<string> Ip6 { get; private set; } = null!;

        /// <summary>
        /// Name of the new web filter profile used by the override.
        /// </summary>
        [Output("newProfile")]
        public Output<string> NewProfile { get; private set; } = null!;

        /// <summary>
        /// Name of the web filter profile which the override applies.
        /// </summary>
        [Output("oldProfile")]
        public Output<string> OldProfile { get; private set; } = null!;

        /// <summary>
        /// Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// Enable/disable override rule. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Name of the user which the override applies.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// Specify the user group for which the override applies.
        /// </summary>
        [Output("userGroup")]
        public Output<string> UserGroup { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Override resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Override(string name, OverrideArgs args, CustomResourceOptions? options = null)
            : base("fortios:filter/web/override:Override", name, args ?? new OverrideArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Override(string name, Input<string> id, OverrideState? state = null, CustomResourceOptions? options = null)
            : base("fortios:filter/web/override:Override", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Override resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Override Get(string name, Input<string> id, OverrideState? state = null, CustomResourceOptions? options = null)
        {
            return new Override(name, id, state, options);
        }
    }

    public sealed class OverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        /// </summary>
        [Input("expires", required: true)]
        public Input<string> Expires { get; set; } = null!;

        /// <summary>
        /// Override rule ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Initiating user of override (read-only setting).
        /// </summary>
        [Input("initiator")]
        public Input<string>? Initiator { get; set; }

        /// <summary>
        /// IPv4 address which the override applies.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// IPv6 address which the override applies.
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        /// <summary>
        /// Name of the new web filter profile used by the override.
        /// </summary>
        [Input("newProfile", required: true)]
        public Input<string> NewProfile { get; set; } = null!;

        /// <summary>
        /// Name of the web filter profile which the override applies.
        /// </summary>
        [Input("oldProfile", required: true)]
        public Input<string> OldProfile { get; set; } = null!;

        /// <summary>
        /// Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Enable/disable override rule. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Name of the user which the override applies.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        /// <summary>
        /// Specify the user group for which the override applies.
        /// </summary>
        [Input("userGroup")]
        public Input<string>? UserGroup { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OverrideArgs()
        {
        }
        public static new OverrideArgs Empty => new OverrideArgs();
    }

    public sealed class OverrideState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
        /// </summary>
        [Input("expires")]
        public Input<string>? Expires { get; set; }

        /// <summary>
        /// Override rule ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Initiating user of override (read-only setting).
        /// </summary>
        [Input("initiator")]
        public Input<string>? Initiator { get; set; }

        /// <summary>
        /// IPv4 address which the override applies.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// IPv6 address which the override applies.
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        /// <summary>
        /// Name of the new web filter profile used by the override.
        /// </summary>
        [Input("newProfile")]
        public Input<string>? NewProfile { get; set; }

        /// <summary>
        /// Name of the web filter profile which the override applies.
        /// </summary>
        [Input("oldProfile")]
        public Input<string>? OldProfile { get; set; }

        /// <summary>
        /// Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Enable/disable override rule. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Name of the user which the override applies.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Specify the user group for which the override applies.
        /// </summary>
        [Input("userGroup")]
        public Input<string>? UserGroup { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public OverrideState()
        {
        }
        public static new OverrideState Empty => new OverrideState();
    }
}
