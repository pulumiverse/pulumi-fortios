// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Autoupdate
{
    /// <summary>
    /// Configure push updates. Applies to FortiOS Version `&lt;= 7.0.0`.
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
    ///     var trname = new Fortios.System.Autoupdate.Pushupdate("trname", new()
    ///     {
    ///         Address = "0.0.0.0",
    ///         Override = "disable",
    ///         Port = 9443,
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SystemAutoupdate PushUpdate can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/autoupdate/pushupdate:Pushupdate labelname SystemAutoupdatePushUpdate
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/autoupdate/pushupdate:Pushupdate labelname SystemAutoupdatePushUpdate
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/autoupdate/pushupdate:Pushupdate")]
    public partial class Pushupdate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Push update override server.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Enable/disable push update override server. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("override")]
        public Output<string> Override { get; private set; } = null!;

        /// <summary>
        /// Push update override port. (Do not overlap with other service ports)
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Enable/disable push updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Pushupdate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pushupdate(string name, PushupdateArgs args, CustomResourceOptions? options = null)
            : base("fortios:system/autoupdate/pushupdate:Pushupdate", name, args ?? new PushupdateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pushupdate(string name, Input<string> id, PushupdateState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/autoupdate/pushupdate:Pushupdate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Pushupdate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pushupdate Get(string name, Input<string> id, PushupdateState? state = null, CustomResourceOptions? options = null)
        {
            return new Pushupdate(name, id, state, options);
        }
    }

    public sealed class PushupdateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Push update override server.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Enable/disable push update override server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("override", required: true)]
        public Input<string> Override { get; set; } = null!;

        /// <summary>
        /// Push update override port. (Do not overlap with other service ports)
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Enable/disable push updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PushupdateArgs()
        {
        }
        public static new PushupdateArgs Empty => new PushupdateArgs();
    }

    public sealed class PushupdateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Push update override server.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Enable/disable push update override server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("override")]
        public Input<string>? Override { get; set; }

        /// <summary>
        /// Push update override port. (Do not overlap with other service ports)
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Enable/disable push updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PushupdateState()
        {
        }
        public static new PushupdateState Empty => new PushupdateState();
    }
}
