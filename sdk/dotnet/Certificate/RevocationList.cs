// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Certificate
{
    /// <summary>
    /// Certificate Revocation List as a PEM file.
    /// 
    /// ## Import
    /// 
    /// Certificate Crl can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:certificate/crl:Crl labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:certificate/crl:Crl labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:certificate/crl:Crl")]
    public partial class RevocationList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Certificate Revocation List as a PEM file.
        /// </summary>
        [Output("crl")]
        public Output<string> Crl { get; private set; } = null!;

        /// <summary>
        /// HTTP server URL for CRL auto-update.
        /// </summary>
        [Output("httpUrl")]
        public Output<string> HttpUrl { get; private set; } = null!;

        /// <summary>
        /// Time at which CRL was last updated.
        /// </summary>
        [Output("lastUpdated")]
        public Output<int> LastUpdated { get; private set; } = null!;

        /// <summary>
        /// LDAP server user password.
        /// </summary>
        [Output("ldapPassword")]
        public Output<string?> LdapPassword { get; private set; } = null!;

        /// <summary>
        /// LDAP server name for CRL auto-update.
        /// </summary>
        [Output("ldapServer")]
        public Output<string> LdapServer { get; private set; } = null!;

        /// <summary>
        /// LDAP server user name.
        /// </summary>
        [Output("ldapUsername")]
        public Output<string> LdapUsername { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Output("range")]
        public Output<string> Range { get; private set; } = null!;

        /// <summary>
        /// Local certificate for SCEP communication for CRL auto-update.
        /// </summary>
        [Output("scepCert")]
        public Output<string> ScepCert { get; private set; } = null!;

        /// <summary>
        /// SCEP server URL for CRL auto-update.
        /// </summary>
        [Output("scepUrl")]
        public Output<string> ScepUrl { get; private set; } = null!;

        /// <summary>
        /// Certificate source type.
        /// </summary>
        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        /// <summary>
        /// Source IP address for communications to a HTTP or SCEP CA server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        /// </summary>
        [Output("updateInterval")]
        public Output<int> UpdateInterval { get; private set; } = null!;

        /// <summary>
        /// VDOM for CRL update.
        /// </summary>
        [Output("updateVdom")]
        public Output<string> UpdateVdom { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a RevocationList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RevocationList(string name, RevocationListArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:certificate/crl:Crl", name, args ?? new RevocationListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RevocationList(string name, Input<string> id, RevocationListState? state = null, CustomResourceOptions? options = null)
            : base("fortios:certificate/crl:Crl", name, state, MakeResourceOptions(options, id))
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
                    "ldapPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RevocationList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RevocationList Get(string name, Input<string> id, RevocationListState? state = null, CustomResourceOptions? options = null)
        {
            return new RevocationList(name, id, state, options);
        }
    }

    public sealed class RevocationListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate Revocation List as a PEM file.
        /// </summary>
        [Input("crl")]
        public Input<string>? Crl { get; set; }

        /// <summary>
        /// HTTP server URL for CRL auto-update.
        /// </summary>
        [Input("httpUrl")]
        public Input<string>? HttpUrl { get; set; }

        /// <summary>
        /// Time at which CRL was last updated.
        /// </summary>
        [Input("lastUpdated")]
        public Input<int>? LastUpdated { get; set; }

        [Input("ldapPassword")]
        private Input<string>? _ldapPassword;

        /// <summary>
        /// LDAP server user password.
        /// </summary>
        public Input<string>? LdapPassword
        {
            get => _ldapPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ldapPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// LDAP server name for CRL auto-update.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// LDAP server user name.
        /// </summary>
        [Input("ldapUsername")]
        public Input<string>? LdapUsername { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Input("range")]
        public Input<string>? Range { get; set; }

        /// <summary>
        /// Local certificate for SCEP communication for CRL auto-update.
        /// </summary>
        [Input("scepCert")]
        public Input<string>? ScepCert { get; set; }

        /// <summary>
        /// SCEP server URL for CRL auto-update.
        /// </summary>
        [Input("scepUrl")]
        public Input<string>? ScepUrl { get; set; }

        /// <summary>
        /// Certificate source type.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Source IP address for communications to a HTTP or SCEP CA server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        /// </summary>
        [Input("updateInterval")]
        public Input<int>? UpdateInterval { get; set; }

        /// <summary>
        /// VDOM for CRL update.
        /// </summary>
        [Input("updateVdom")]
        public Input<string>? UpdateVdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RevocationListArgs()
        {
        }
        public static new RevocationListArgs Empty => new RevocationListArgs();
    }

    public sealed class RevocationListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate Revocation List as a PEM file.
        /// </summary>
        [Input("crl")]
        public Input<string>? Crl { get; set; }

        /// <summary>
        /// HTTP server URL for CRL auto-update.
        /// </summary>
        [Input("httpUrl")]
        public Input<string>? HttpUrl { get; set; }

        /// <summary>
        /// Time at which CRL was last updated.
        /// </summary>
        [Input("lastUpdated")]
        public Input<int>? LastUpdated { get; set; }

        [Input("ldapPassword")]
        private Input<string>? _ldapPassword;

        /// <summary>
        /// LDAP server user password.
        /// </summary>
        public Input<string>? LdapPassword
        {
            get => _ldapPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ldapPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// LDAP server name for CRL auto-update.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// LDAP server user name.
        /// </summary>
        [Input("ldapUsername")]
        public Input<string>? LdapUsername { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
        /// </summary>
        [Input("range")]
        public Input<string>? Range { get; set; }

        /// <summary>
        /// Local certificate for SCEP communication for CRL auto-update.
        /// </summary>
        [Input("scepCert")]
        public Input<string>? ScepCert { get; set; }

        /// <summary>
        /// SCEP server URL for CRL auto-update.
        /// </summary>
        [Input("scepUrl")]
        public Input<string>? ScepUrl { get; set; }

        /// <summary>
        /// Certificate source type.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Source IP address for communications to a HTTP or SCEP CA server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
        /// </summary>
        [Input("updateInterval")]
        public Input<int>? UpdateInterval { get; set; }

        /// <summary>
        /// VDOM for CRL update.
        /// </summary>
        [Input("updateVdom")]
        public Input<string>? UpdateVdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RevocationListState()
        {
        }
        public static new RevocationListState Empty => new RevocationListState();
    }
}
