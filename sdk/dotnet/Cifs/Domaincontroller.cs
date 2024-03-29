// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Cifs
{
    /// <summary>
    /// Define known domain controller servers. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1`.
    /// 
    /// ## Import
    /// 
    /// Cifs DomainController can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:cifs/domaincontroller:Domaincontroller labelname {{server_name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:cifs/domaincontroller:Domaincontroller labelname {{server_name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:cifs/domaincontroller:Domaincontroller")]
    public partial class Domaincontroller : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// IPv4 server address.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// IPv6 server address.
        /// </summary>
        [Output("ip6")]
        public Output<string> Ip6 { get; private set; } = null!;

        /// <summary>
        /// Password for specified username.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Port number of service. Port number 0 indicates automatic discovery.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Name of the server to connect to.
        /// </summary>
        [Output("serverName")]
        public Output<string> ServerName { get; private set; } = null!;

        /// <summary>
        /// User name to sign in with. Must have proper permissions for service.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Domaincontroller resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domaincontroller(string name, DomaincontrollerArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:cifs/domaincontroller:Domaincontroller", name, args ?? new DomaincontrollerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domaincontroller(string name, Input<string> id, DomaincontrollerState? state = null, CustomResourceOptions? options = null)
            : base("fortios:cifs/domaincontroller:Domaincontroller", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Domaincontroller resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domaincontroller Get(string name, Input<string> id, DomaincontrollerState? state = null, CustomResourceOptions? options = null)
        {
            return new Domaincontroller(name, id, state, options);
        }
    }

    public sealed class DomaincontrollerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// IPv4 server address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// IPv6 server address.
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for specified username.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number of service. Port number 0 indicates automatic discovery.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Name of the server to connect to.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// User name to sign in with. Must have proper permissions for service.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DomaincontrollerArgs()
        {
        }
        public static new DomaincontrollerArgs Empty => new DomaincontrollerArgs();
    }

    public sealed class DomaincontrollerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// IPv4 server address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// IPv6 server address.
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for specified username.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number of service. Port number 0 indicates automatic discovery.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Name of the server to connect to.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// User name to sign in with. Must have proper permissions for service.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public DomaincontrollerState()
        {
        }
        public static new DomaincontrollerState Empty => new DomaincontrollerState();
    }
}
