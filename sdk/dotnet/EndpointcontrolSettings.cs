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
    /// Configure endpoint control settings. Applies to FortiOS Version `&lt;= 6.2.6`.
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
    ///     var trname = new Fortios.EndpointcontrolSettings("trname", new()
    ///     {
    ///         DownloadLocation = "fortiguard",
    ///         ForticlientAvdbUpdateInterval = 8,
    ///         ForticlientDeregUnsupportedClient = "enable",
    ///         ForticlientEmsRestApiCallTimeout = 5000,
    ///         ForticlientKeepaliveInterval = 60,
    ///         ForticlientOfflineGrace = "disable",
    ///         ForticlientOfflineGraceInterval = 120,
    ///         ForticlientRegKeyEnforce = "disable",
    ///         ForticlientRegTimeout = 7,
    ///         ForticlientSysUpdateInterval = 720,
    ///         ForticlientUserAvatar = "enable",
    ///         ForticlientWarningInterval = 1,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EndpointControl Settings can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/endpointcontrolSettings:EndpointcontrolSettings labelname EndpointControlSettings
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/endpointcontrolSettings:EndpointcontrolSettings labelname EndpointControlSettings
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/endpointcontrolSettings:EndpointcontrolSettings")]
    public partial class EndpointcontrolSettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Customized URL for downloading FortiClient.
        /// </summary>
        [Output("downloadCustomLink")]
        public Output<string> DownloadCustomLink { get; private set; } = null!;

        /// <summary>
        /// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
        /// </summary>
        [Output("downloadLocation")]
        public Output<string> DownloadLocation { get; private set; } = null!;

        /// <summary>
        /// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
        /// </summary>
        [Output("forticlientAvdbUpdateInterval")]
        public Output<int> ForticlientAvdbUpdateInterval { get; private set; } = null!;

        /// <summary>
        /// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("forticlientDeregUnsupportedClient")]
        public Output<string> ForticlientDeregUnsupportedClient { get; private set; } = null!;

        /// <summary>
        /// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("forticlientDisconnectUnsupportedClient")]
        public Output<string> ForticlientDisconnectUnsupportedClient { get; private set; } = null!;

        /// <summary>
        /// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        /// </summary>
        [Output("forticlientEmsRestApiCallTimeout")]
        public Output<int> ForticlientEmsRestApiCallTimeout { get; private set; } = null!;

        /// <summary>
        /// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
        /// </summary>
        [Output("forticlientKeepaliveInterval")]
        public Output<int> ForticlientKeepaliveInterval { get; private set; } = null!;

        /// <summary>
        /// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("forticlientOfflineGrace")]
        public Output<string> ForticlientOfflineGrace { get; private set; } = null!;

        /// <summary>
        /// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
        /// </summary>
        [Output("forticlientOfflineGraceInterval")]
        public Output<int> ForticlientOfflineGraceInterval { get; private set; } = null!;

        /// <summary>
        /// FortiClient registration key.
        /// </summary>
        [Output("forticlientRegKey")]
        public Output<string?> ForticlientRegKey { get; private set; } = null!;

        /// <summary>
        /// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("forticlientRegKeyEnforce")]
        public Output<string> ForticlientRegKeyEnforce { get; private set; } = null!;

        /// <summary>
        /// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
        /// </summary>
        [Output("forticlientRegTimeout")]
        public Output<int> ForticlientRegTimeout { get; private set; } = null!;

        /// <summary>
        /// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
        /// </summary>
        [Output("forticlientSysUpdateInterval")]
        public Output<int> ForticlientSysUpdateInterval { get; private set; } = null!;

        /// <summary>
        /// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("forticlientUserAvatar")]
        public Output<string> ForticlientUserAvatar { get; private set; } = null!;

        /// <summary>
        /// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
        /// </summary>
        [Output("forticlientWarningInterval")]
        public Output<int> ForticlientWarningInterval { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointcontrolSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointcontrolSettings(string name, EndpointcontrolSettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/endpointcontrolSettings:EndpointcontrolSettings", name, args ?? new EndpointcontrolSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointcontrolSettings(string name, Input<string> id, EndpointcontrolSettingsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/endpointcontrolSettings:EndpointcontrolSettings", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "forticlientRegKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EndpointcontrolSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointcontrolSettings Get(string name, Input<string> id, EndpointcontrolSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointcontrolSettings(name, id, state, options);
        }
    }

    public sealed class EndpointcontrolSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Customized URL for downloading FortiClient.
        /// </summary>
        [Input("downloadCustomLink")]
        public Input<string>? DownloadCustomLink { get; set; }

        /// <summary>
        /// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
        /// </summary>
        [Input("downloadLocation")]
        public Input<string>? DownloadLocation { get; set; }

        /// <summary>
        /// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
        /// </summary>
        [Input("forticlientAvdbUpdateInterval")]
        public Input<int>? ForticlientAvdbUpdateInterval { get; set; }

        /// <summary>
        /// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientDeregUnsupportedClient")]
        public Input<string>? ForticlientDeregUnsupportedClient { get; set; }

        /// <summary>
        /// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientDisconnectUnsupportedClient")]
        public Input<string>? ForticlientDisconnectUnsupportedClient { get; set; }

        /// <summary>
        /// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        /// </summary>
        [Input("forticlientEmsRestApiCallTimeout")]
        public Input<int>? ForticlientEmsRestApiCallTimeout { get; set; }

        /// <summary>
        /// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
        /// </summary>
        [Input("forticlientKeepaliveInterval")]
        public Input<int>? ForticlientKeepaliveInterval { get; set; }

        /// <summary>
        /// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientOfflineGrace")]
        public Input<string>? ForticlientOfflineGrace { get; set; }

        /// <summary>
        /// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
        /// </summary>
        [Input("forticlientOfflineGraceInterval")]
        public Input<int>? ForticlientOfflineGraceInterval { get; set; }

        [Input("forticlientRegKey")]
        private Input<string>? _forticlientRegKey;

        /// <summary>
        /// FortiClient registration key.
        /// </summary>
        public Input<string>? ForticlientRegKey
        {
            get => _forticlientRegKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _forticlientRegKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientRegKeyEnforce")]
        public Input<string>? ForticlientRegKeyEnforce { get; set; }

        /// <summary>
        /// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
        /// </summary>
        [Input("forticlientRegTimeout")]
        public Input<int>? ForticlientRegTimeout { get; set; }

        /// <summary>
        /// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
        /// </summary>
        [Input("forticlientSysUpdateInterval")]
        public Input<int>? ForticlientSysUpdateInterval { get; set; }

        /// <summary>
        /// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientUserAvatar")]
        public Input<string>? ForticlientUserAvatar { get; set; }

        /// <summary>
        /// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
        /// </summary>
        [Input("forticlientWarningInterval")]
        public Input<int>? ForticlientWarningInterval { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public EndpointcontrolSettingsArgs()
        {
        }
        public static new EndpointcontrolSettingsArgs Empty => new EndpointcontrolSettingsArgs();
    }

    public sealed class EndpointcontrolSettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Customized URL for downloading FortiClient.
        /// </summary>
        [Input("downloadCustomLink")]
        public Input<string>? DownloadCustomLink { get; set; }

        /// <summary>
        /// FortiClient download location (FortiGuard or custom). Valid values: `fortiguard`, `custom`.
        /// </summary>
        [Input("downloadLocation")]
        public Input<string>? DownloadLocation { get; set; }

        /// <summary>
        /// Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8).
        /// </summary>
        [Input("forticlientAvdbUpdateInterval")]
        public Input<int>? ForticlientAvdbUpdateInterval { get; set; }

        /// <summary>
        /// Enable/disable deregistering unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientDeregUnsupportedClient")]
        public Input<string>? ForticlientDeregUnsupportedClient { get; set; }

        /// <summary>
        /// Enable/disable disconnecting of unsupported FortiClient endpoints. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientDisconnectUnsupportedClient")]
        public Input<string>? ForticlientDisconnectUnsupportedClient { get; set; }

        /// <summary>
        /// FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        /// </summary>
        [Input("forticlientEmsRestApiCallTimeout")]
        public Input<int>? ForticlientEmsRestApiCallTimeout { get; set; }

        /// <summary>
        /// Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
        /// </summary>
        [Input("forticlientKeepaliveInterval")]
        public Input<int>? ForticlientKeepaliveInterval { get; set; }

        /// <summary>
        /// Enable/disable grace period for offline registered clients. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientOfflineGrace")]
        public Input<string>? ForticlientOfflineGrace { get; set; }

        /// <summary>
        /// Grace period for offline registered FortiClient (60 - 600 sec, default = 120).
        /// </summary>
        [Input("forticlientOfflineGraceInterval")]
        public Input<int>? ForticlientOfflineGraceInterval { get; set; }

        [Input("forticlientRegKey")]
        private Input<string>? _forticlientRegKey;

        /// <summary>
        /// FortiClient registration key.
        /// </summary>
        public Input<string>? ForticlientRegKey
        {
            get => _forticlientRegKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _forticlientRegKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable requiring or enforcing FortiClient registration keys. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientRegKeyEnforce")]
        public Input<string>? ForticlientRegKeyEnforce { get; set; }

        /// <summary>
        /// FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited).
        /// </summary>
        [Input("forticlientRegTimeout")]
        public Input<int>? ForticlientRegTimeout { get; set; }

        /// <summary>
        /// Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
        /// </summary>
        [Input("forticlientSysUpdateInterval")]
        public Input<int>? ForticlientSysUpdateInterval { get; set; }

        /// <summary>
        /// Enable/disable uploading FortiClient user avatars. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("forticlientUserAvatar")]
        public Input<string>? ForticlientUserAvatar { get; set; }

        /// <summary>
        /// Period of time between FortiClient portal warnings (0 - 24 hours, default = 1).
        /// </summary>
        [Input("forticlientWarningInterval")]
        public Input<int>? ForticlientWarningInterval { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public EndpointcontrolSettingsState()
        {
        }
        public static new EndpointcontrolSettingsState Empty => new EndpointcontrolSettingsState();
    }
}
