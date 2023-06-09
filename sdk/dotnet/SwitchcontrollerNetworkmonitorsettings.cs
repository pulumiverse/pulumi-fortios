// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure network monitor settings.
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
    ///     var trname = new Fortios.SwitchcontrollerNetworkmonitorsettings("trname", new()
    ///     {
    ///         NetworkMonitoring = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SwitchController NetworkMonitorSettings can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerNetworkmonitorsettings:SwitchcontrollerNetworkmonitorsettings labelname SwitchControllerNetworkMonitorSettings
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerNetworkmonitorsettings:SwitchcontrollerNetworkmonitorsettings labelname SwitchControllerNetworkMonitorSettings
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerNetworkmonitorsettings:SwitchcontrollerNetworkmonitorsettings")]
    public partial class SwitchcontrollerNetworkmonitorsettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("networkMonitoring")]
        public Output<string> NetworkMonitoring { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerNetworkmonitorsettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerNetworkmonitorsettings(string name, SwitchcontrollerNetworkmonitorsettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerNetworkmonitorsettings:SwitchcontrollerNetworkmonitorsettings", name, args ?? new SwitchcontrollerNetworkmonitorsettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerNetworkmonitorsettings(string name, Input<string> id, SwitchcontrollerNetworkmonitorsettingsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerNetworkmonitorsettings:SwitchcontrollerNetworkmonitorsettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchcontrollerNetworkmonitorsettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerNetworkmonitorsettings Get(string name, Input<string> id, SwitchcontrollerNetworkmonitorsettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerNetworkmonitorsettings(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerNetworkmonitorsettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("networkMonitoring")]
        public Input<string>? NetworkMonitoring { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerNetworkmonitorsettingsArgs()
        {
        }
        public static new SwitchcontrollerNetworkmonitorsettingsArgs Empty => new SwitchcontrollerNetworkmonitorsettingsArgs();
    }

    public sealed class SwitchcontrollerNetworkmonitorsettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("networkMonitoring")]
        public Input<string>? NetworkMonitoring { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerNetworkmonitorsettingsState()
        {
        }
        public static new SwitchcontrollerNetworkmonitorsettingsState Empty => new SwitchcontrollerNetworkmonitorsettingsState();
    }
}
