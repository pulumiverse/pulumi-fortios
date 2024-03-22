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
    /// Configure IPS VDOM parameter.
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
    ///     var trname = new Fortios.Ips.Settings("trname", new()
    ///     {
    ///         IpsPacketQuota = 0,
    ///         PacketLogHistory = 1,
    ///         PacketLogMemory = 256,
    ///         PacketLogPostAttack = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Ips Settings can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:ips/settings:Settings labelname IpsSettings
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:ips/settings:Settings labelname IpsSettings
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:ips/settings:Settings")]
    public partial class Settings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        /// </summary>
        [Output("ipsPacketQuota")]
        public Output<int> IpsPacketQuota { get; private set; } = null!;

        /// <summary>
        /// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        /// </summary>
        [Output("packetLogHistory")]
        public Output<int> PacketLogHistory { get; private set; } = null!;

        /// <summary>
        /// Maximum memory can be used by packet log (64 - 8192 kB).
        /// </summary>
        [Output("packetLogMemory")]
        public Output<int> PacketLogMemory { get; private set; } = null!;

        /// <summary>
        /// Number of packets to log after the IPS signature is detected (0 - 255).
        /// </summary>
        [Output("packetLogPostAttack")]
        public Output<int> PacketLogPostAttack { get; private set; } = null!;

        /// <summary>
        /// Enable/disable proxy-mode policy inline IPS support. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("proxyInlineIps")]
        public Output<string> ProxyInlineIps { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Settings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Settings(string name, SettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:ips/settings:Settings", name, args ?? new SettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Settings(string name, Input<string> id, SettingsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:ips/settings:Settings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Settings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Settings Get(string name, Input<string> id, SettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new Settings(name, id, state, options);
        }
    }

    public sealed class SettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        /// </summary>
        [Input("ipsPacketQuota")]
        public Input<int>? IpsPacketQuota { get; set; }

        /// <summary>
        /// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        /// </summary>
        [Input("packetLogHistory")]
        public Input<int>? PacketLogHistory { get; set; }

        /// <summary>
        /// Maximum memory can be used by packet log (64 - 8192 kB).
        /// </summary>
        [Input("packetLogMemory")]
        public Input<int>? PacketLogMemory { get; set; }

        /// <summary>
        /// Number of packets to log after the IPS signature is detected (0 - 255).
        /// </summary>
        [Input("packetLogPostAttack")]
        public Input<int>? PacketLogPostAttack { get; set; }

        /// <summary>
        /// Enable/disable proxy-mode policy inline IPS support. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("proxyInlineIps")]
        public Input<string>? ProxyInlineIps { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingsArgs()
        {
        }
        public static new SettingsArgs Empty => new SettingsArgs();
    }

    public sealed class SettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.
        /// </summary>
        [Input("ipsPacketQuota")]
        public Input<int>? IpsPacketQuota { get; set; }

        /// <summary>
        /// Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
        /// </summary>
        [Input("packetLogHistory")]
        public Input<int>? PacketLogHistory { get; set; }

        /// <summary>
        /// Maximum memory can be used by packet log (64 - 8192 kB).
        /// </summary>
        [Input("packetLogMemory")]
        public Input<int>? PacketLogMemory { get; set; }

        /// <summary>
        /// Number of packets to log after the IPS signature is detected (0 - 255).
        /// </summary>
        [Input("packetLogPostAttack")]
        public Input<int>? PacketLogPostAttack { get; set; }

        /// <summary>
        /// Enable/disable proxy-mode policy inline IPS support. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("proxyInlineIps")]
        public Input<string>? ProxyInlineIps { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingsState()
        {
        }
        public static new SettingsState Empty => new SettingsState();
    }
}
