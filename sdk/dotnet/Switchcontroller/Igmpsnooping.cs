// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller
{
    /// <summary>
    /// Configure FortiSwitch IGMP snooping global settings.
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
    ///     var trname = new Fortios.Switchcontroller.Igmpsnooping("trname", new()
    ///     {
    ///         AgingTime = 300,
    ///         FloodUnknownMulticast = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SwitchController IgmpSnooping can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/igmpsnooping:Igmpsnooping labelname SwitchControllerIgmpSnooping
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/igmpsnooping:Igmpsnooping labelname SwitchControllerIgmpSnooping
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/igmpsnooping:Igmpsnooping")]
    public partial class Igmpsnooping : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
        /// </summary>
        [Output("agingTime")]
        public Output<int> AgingTime { get; private set; } = null!;

        /// <summary>
        /// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("floodUnknownMulticast")]
        public Output<string> FloodUnknownMulticast { get; private set; } = null!;

        /// <summary>
        /// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
        /// </summary>
        [Output("queryInterval")]
        public Output<int> QueryInterval { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Igmpsnooping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Igmpsnooping(string name, IgmpsnoopingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/igmpsnooping:Igmpsnooping", name, args ?? new IgmpsnoopingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Igmpsnooping(string name, Input<string> id, IgmpsnoopingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/igmpsnooping:Igmpsnooping", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Igmpsnooping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Igmpsnooping Get(string name, Input<string> id, IgmpsnoopingState? state = null, CustomResourceOptions? options = null)
        {
            return new Igmpsnooping(name, id, state, options);
        }
    }

    public sealed class IgmpsnoopingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
        /// </summary>
        [Input("agingTime")]
        public Input<int>? AgingTime { get; set; }

        /// <summary>
        /// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("floodUnknownMulticast")]
        public Input<string>? FloodUnknownMulticast { get; set; }

        /// <summary>
        /// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
        /// </summary>
        [Input("queryInterval")]
        public Input<int>? QueryInterval { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IgmpsnoopingArgs()
        {
        }
        public static new IgmpsnoopingArgs Empty => new IgmpsnoopingArgs();
    }

    public sealed class IgmpsnoopingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
        /// </summary>
        [Input("agingTime")]
        public Input<int>? AgingTime { get; set; }

        /// <summary>
        /// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("floodUnknownMulticast")]
        public Input<string>? FloodUnknownMulticast { get; set; }

        /// <summary>
        /// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
        /// </summary>
        [Input("queryInterval")]
        public Input<int>? QueryInterval { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public IgmpsnoopingState()
        {
        }
        public static new IgmpsnoopingState Empty => new IgmpsnoopingState();
    }
}
