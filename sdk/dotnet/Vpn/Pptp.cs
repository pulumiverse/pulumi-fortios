// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn
{
    /// <summary>
    /// Configure PPTP.
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
    ///     var trname = new Fortios.Vpn.Pptp("trname", new()
    ///     {
    ///         Eip = "1.1.1.22",
    ///         IpMode = "range",
    ///         LocalIp = "0.0.0.0",
    ///         Sip = "1.1.1.1",
    ///         Status = "enable",
    ///         Usrgrp = "Guest-group",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Vpn Pptp can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/pptp:Pptp labelname VpnPptp
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:vpn/pptp:Pptp labelname VpnPptp
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:vpn/pptp:Pptp")]
    public partial class Pptp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// End IP.
        /// </summary>
        [Output("eip")]
        public Output<string> Eip { get; private set; } = null!;

        /// <summary>
        /// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        /// </summary>
        [Output("ipMode")]
        public Output<string> IpMode { get; private set; } = null!;

        /// <summary>
        /// Local IP to be used for peer's remote IP.
        /// </summary>
        [Output("localIp")]
        public Output<string> LocalIp { get; private set; } = null!;

        /// <summary>
        /// Start IP.
        /// </summary>
        [Output("sip")]
        public Output<string> Sip { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// User group.
        /// </summary>
        [Output("usrgrp")]
        public Output<string> Usrgrp { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Pptp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pptp(string name, PptpArgs args, CustomResourceOptions? options = null)
            : base("fortios:vpn/pptp:Pptp", name, args ?? new PptpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pptp(string name, Input<string> id, PptpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:vpn/pptp:Pptp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Pptp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pptp Get(string name, Input<string> id, PptpState? state = null, CustomResourceOptions? options = null)
        {
            return new Pptp(name, id, state, options);
        }
    }

    public sealed class PptpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End IP.
        /// </summary>
        [Input("eip")]
        public Input<string>? Eip { get; set; }

        /// <summary>
        /// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        /// </summary>
        [Input("ipMode")]
        public Input<string>? IpMode { get; set; }

        /// <summary>
        /// Local IP to be used for peer's remote IP.
        /// </summary>
        [Input("localIp")]
        public Input<string>? LocalIp { get; set; }

        /// <summary>
        /// Start IP.
        /// </summary>
        [Input("sip")]
        public Input<string>? Sip { get; set; }

        /// <summary>
        /// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// User group.
        /// </summary>
        [Input("usrgrp")]
        public Input<string>? Usrgrp { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PptpArgs()
        {
        }
        public static new PptpArgs Empty => new PptpArgs();
    }

    public sealed class PptpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End IP.
        /// </summary>
        [Input("eip")]
        public Input<string>? Eip { get; set; }

        /// <summary>
        /// IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
        /// </summary>
        [Input("ipMode")]
        public Input<string>? IpMode { get; set; }

        /// <summary>
        /// Local IP to be used for peer's remote IP.
        /// </summary>
        [Input("localIp")]
        public Input<string>? LocalIp { get; set; }

        /// <summary>
        /// Start IP.
        /// </summary>
        [Input("sip")]
        public Input<string>? Sip { get; set; }

        /// <summary>
        /// Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// User group.
        /// </summary>
        [Input("usrgrp")]
        public Input<string>? Usrgrp { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PptpState()
        {
        }
        public static new PptpState Empty => new PptpState();
    }
}
