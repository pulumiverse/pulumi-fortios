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
    /// Configure WAN optimization profiles.
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
    ///     var trname = new Fortios.WanoptProfile("trname", new()
    ///     {
    ///         Cifs = new Fortios.Inputs.WanoptProfileCifsArgs
    ///         {
    ///             ByteCaching = "enable",
    ///             LogTraffic = "enable",
    ///             Port = 445,
    ///             PreferChunking = "fix",
    ///             SecureTunnel = "disable",
    ///             Status = "disable",
    ///             TunnelSharing = "private",
    ///         },
    ///         Comments = "test",
    ///         Ftp = new Fortios.Inputs.WanoptProfileFtpArgs
    ///         {
    ///             ByteCaching = "enable",
    ///             LogTraffic = "enable",
    ///             Port = 21,
    ///             PreferChunking = "fix",
    ///             SecureTunnel = "disable",
    ///             Status = "disable",
    ///             TunnelSharing = "private",
    ///         },
    ///         Http = new Fortios.Inputs.WanoptProfileHttpArgs
    ///         {
    ///             ByteCaching = "enable",
    ///             LogTraffic = "enable",
    ///             Port = 80,
    ///             PreferChunking = "fix",
    ///             SecureTunnel = "disable",
    ///             Ssl = "disable",
    ///             SslPort = 443,
    ///             Status = "disable",
    ///             TunnelNonHttp = "disable",
    ///             TunnelSharing = "private",
    ///             UnknownHttpVersion = "tunnel",
    ///         },
    ///         Mapi = new Fortios.Inputs.WanoptProfileMapiArgs
    ///         {
    ///             ByteCaching = "enable",
    ///             LogTraffic = "enable",
    ///             Port = 135,
    ///             SecureTunnel = "disable",
    ///             Status = "disable",
    ///             TunnelSharing = "private",
    ///         },
    ///         Tcp = new Fortios.Inputs.WanoptProfileTcpArgs
    ///         {
    ///             ByteCaching = "disable",
    ///             ByteCachingOpt = "mem-only",
    ///             LogTraffic = "enable",
    ///             Port = "1-65535",
    ///             SecureTunnel = "disable",
    ///             Ssl = "disable",
    ///             SslPort = 443,
    ///             Status = "disable",
    ///             TunnelSharing = "private",
    ///         },
    ///         Transparent = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Wanopt Profile can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/wanoptProfile:WanoptProfile")]
    public partial class WanoptProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        /// </summary>
        [Output("authGroup")]
        public Output<string> AuthGroup { get; private set; } = null!;

        /// <summary>
        /// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        /// </summary>
        [Output("cifs")]
        public Output<Outputs.WanoptProfileCifs> Cifs { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        /// </summary>
        [Output("ftp")]
        public Output<Outputs.WanoptProfileFtp> Ftp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        /// </summary>
        [Output("http")]
        public Output<Outputs.WanoptProfileHttp> Http { get; private set; } = null!;

        /// <summary>
        /// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        /// </summary>
        [Output("mapi")]
        public Output<Outputs.WanoptProfileMapi> Mapi { get; private set; } = null!;

        /// <summary>
        /// Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        /// </summary>
        [Output("tcp")]
        public Output<Outputs.WanoptProfileTcp> Tcp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable transparent mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("transparent")]
        public Output<string> Transparent { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a WanoptProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WanoptProfile(string name, WanoptProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/wanoptProfile:WanoptProfile", name, args ?? new WanoptProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WanoptProfile(string name, Input<string> id, WanoptProfileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/wanoptProfile:WanoptProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WanoptProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WanoptProfile Get(string name, Input<string> id, WanoptProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new WanoptProfile(name, id, state, options);
        }
    }

    public sealed class WanoptProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        /// </summary>
        [Input("authGroup")]
        public Input<string>? AuthGroup { get; set; }

        /// <summary>
        /// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        /// </summary>
        [Input("cifs")]
        public Input<Inputs.WanoptProfileCifsArgs>? Cifs { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        /// </summary>
        [Input("ftp")]
        public Input<Inputs.WanoptProfileFtpArgs>? Ftp { get; set; }

        /// <summary>
        /// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        /// </summary>
        [Input("http")]
        public Input<Inputs.WanoptProfileHttpArgs>? Http { get; set; }

        /// <summary>
        /// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        /// </summary>
        [Input("mapi")]
        public Input<Inputs.WanoptProfileMapiArgs>? Mapi { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        /// </summary>
        [Input("tcp")]
        public Input<Inputs.WanoptProfileTcpArgs>? Tcp { get; set; }

        /// <summary>
        /// Enable/disable transparent mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("transparent")]
        public Input<string>? Transparent { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WanoptProfileArgs()
        {
        }
        public static new WanoptProfileArgs Empty => new WanoptProfileArgs();
    }

    public sealed class WanoptProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
        /// </summary>
        [Input("authGroup")]
        public Input<string>? AuthGroup { get; set; }

        /// <summary>
        /// Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
        /// </summary>
        [Input("cifs")]
        public Input<Inputs.WanoptProfileCifsGetArgs>? Cifs { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
        /// </summary>
        [Input("ftp")]
        public Input<Inputs.WanoptProfileFtpGetArgs>? Ftp { get; set; }

        /// <summary>
        /// Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
        /// </summary>
        [Input("http")]
        public Input<Inputs.WanoptProfileHttpGetArgs>? Http { get; set; }

        /// <summary>
        /// Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
        /// </summary>
        [Input("mapi")]
        public Input<Inputs.WanoptProfileMapiGetArgs>? Mapi { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
        /// </summary>
        [Input("tcp")]
        public Input<Inputs.WanoptProfileTcpGetArgs>? Tcp { get; set; }

        /// <summary>
        /// Enable/disable transparent mode. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("transparent")]
        public Input<string>? Transparent { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public WanoptProfileState()
        {
        }
        public static new WanoptProfileState Empty => new WanoptProfileState();
    }
}