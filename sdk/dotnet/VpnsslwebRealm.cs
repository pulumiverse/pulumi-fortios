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
    /// Realm.
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
    ///     var trname = new Fortios.VpnsslwebRealm("trname", new()
    ///     {
    ///         LoginPage = "1.htm",
    ///         MaxConcurrentUser = 33,
    ///         UrlPath = "1",
    ///         VirtualHost = "2.2.2.2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VpnSslWeb Realm can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/vpnsslwebRealm:VpnsslwebRealm labelname {{url_path}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/vpnsslwebRealm:VpnsslwebRealm labelname {{url_path}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/vpnsslwebRealm:VpnsslwebRealm")]
    public partial class VpnsslwebRealm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Replacement HTML for SSL-VPN login page.
        /// </summary>
        [Output("loginPage")]
        public Output<string?> LoginPage { get; private set; } = null!;

        /// <summary>
        /// Maximum concurrent users (0 - 65535, 0 means unlimited).
        /// </summary>
        [Output("maxConcurrentUser")]
        public Output<int> MaxConcurrentUser { get; private set; } = null!;

        /// <summary>
        /// IP address used as a NAS-IP to communicate with the RADIUS server.
        /// </summary>
        [Output("nasIp")]
        public Output<string> NasIp { get; private set; } = null!;

        /// <summary>
        /// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
        /// </summary>
        [Output("radiusPort")]
        public Output<int> RadiusPort { get; private set; } = null!;

        /// <summary>
        /// RADIUS server associated with realm.
        /// </summary>
        [Output("radiusServer")]
        public Output<string> RadiusServer { get; private set; } = null!;

        /// <summary>
        /// URL path to access SSL-VPN login page.
        /// </summary>
        [Output("urlPath")]
        public Output<string> UrlPath { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Virtual host name for realm.
        /// </summary>
        [Output("virtualHost")]
        public Output<string?> VirtualHost { get; private set; } = null!;

        /// <summary>
        /// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("virtualHostOnly")]
        public Output<string> VirtualHostOnly { get; private set; } = null!;

        /// <summary>
        /// Name of the server certificate to used for this realm.
        /// </summary>
        [Output("virtualHostServerCert")]
        public Output<string> VirtualHostServerCert { get; private set; } = null!;


        /// <summary>
        /// Create a VpnsslwebRealm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnsslwebRealm(string name, VpnsslwebRealmArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/vpnsslwebRealm:VpnsslwebRealm", name, args ?? new VpnsslwebRealmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnsslwebRealm(string name, Input<string> id, VpnsslwebRealmState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/vpnsslwebRealm:VpnsslwebRealm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpnsslwebRealm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnsslwebRealm Get(string name, Input<string> id, VpnsslwebRealmState? state = null, CustomResourceOptions? options = null)
        {
            return new VpnsslwebRealm(name, id, state, options);
        }
    }

    public sealed class VpnsslwebRealmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Replacement HTML for SSL-VPN login page.
        /// </summary>
        [Input("loginPage")]
        public Input<string>? LoginPage { get; set; }

        /// <summary>
        /// Maximum concurrent users (0 - 65535, 0 means unlimited).
        /// </summary>
        [Input("maxConcurrentUser")]
        public Input<int>? MaxConcurrentUser { get; set; }

        /// <summary>
        /// IP address used as a NAS-IP to communicate with the RADIUS server.
        /// </summary>
        [Input("nasIp")]
        public Input<string>? NasIp { get; set; }

        /// <summary>
        /// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
        /// </summary>
        [Input("radiusPort")]
        public Input<int>? RadiusPort { get; set; }

        /// <summary>
        /// RADIUS server associated with realm.
        /// </summary>
        [Input("radiusServer")]
        public Input<string>? RadiusServer { get; set; }

        /// <summary>
        /// URL path to access SSL-VPN login page.
        /// </summary>
        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Virtual host name for realm.
        /// </summary>
        [Input("virtualHost")]
        public Input<string>? VirtualHost { get; set; }

        /// <summary>
        /// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("virtualHostOnly")]
        public Input<string>? VirtualHostOnly { get; set; }

        /// <summary>
        /// Name of the server certificate to used for this realm.
        /// </summary>
        [Input("virtualHostServerCert")]
        public Input<string>? VirtualHostServerCert { get; set; }

        public VpnsslwebRealmArgs()
        {
        }
        public static new VpnsslwebRealmArgs Empty => new VpnsslwebRealmArgs();
    }

    public sealed class VpnsslwebRealmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Replacement HTML for SSL-VPN login page.
        /// </summary>
        [Input("loginPage")]
        public Input<string>? LoginPage { get; set; }

        /// <summary>
        /// Maximum concurrent users (0 - 65535, 0 means unlimited).
        /// </summary>
        [Input("maxConcurrentUser")]
        public Input<int>? MaxConcurrentUser { get; set; }

        /// <summary>
        /// IP address used as a NAS-IP to communicate with the RADIUS server.
        /// </summary>
        [Input("nasIp")]
        public Input<string>? NasIp { get; set; }

        /// <summary>
        /// RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
        /// </summary>
        [Input("radiusPort")]
        public Input<int>? RadiusPort { get; set; }

        /// <summary>
        /// RADIUS server associated with realm.
        /// </summary>
        [Input("radiusServer")]
        public Input<string>? RadiusServer { get; set; }

        /// <summary>
        /// URL path to access SSL-VPN login page.
        /// </summary>
        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Virtual host name for realm.
        /// </summary>
        [Input("virtualHost")]
        public Input<string>? VirtualHost { get; set; }

        /// <summary>
        /// Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("virtualHostOnly")]
        public Input<string>? VirtualHostOnly { get; set; }

        /// <summary>
        /// Name of the server certificate to used for this realm.
        /// </summary>
        [Input("virtualHostServerCert")]
        public Input<string>? VirtualHostServerCert { get; set; }

        public VpnsslwebRealmState()
        {
        }
        public static new VpnsslwebRealmState Empty => new VpnsslwebRealmState();
    }
}
