// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User
{
    /// <summary>
    /// Configure domain controller entries.
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
    ///     var trname1 = new Fortios.User.Ldap("trname1", new()
    ///     {
    ///         AccountKeyFilter = "(&amp;(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))",
    ///         AccountKeyProcessing = "same",
    ///         Cnid = "cn",
    ///         Dn = "EIWNCIEW",
    ///         GroupMemberCheck = "user-attr",
    ///         GroupObjectFilter = "(&amp;(objectcategory=group)(member=*))",
    ///         MemberAttr = "memberOf",
    ///         PasswordExpiryWarning = "disable",
    ///         PasswordRenewal = "disable",
    ///         Port = 389,
    ///         Secure = "disable",
    ///         Server = "1.1.1.1",
    ///         ServerIdentityCheck = "disable",
    ///         SourceIp = "0.0.0.0",
    ///         SslMinProtoVersion = "default",
    ///         Type = "simple",
    ///     });
    /// 
    ///     var trname = new Fortios.User.Domaincontroller("trname", new()
    ///     {
    ///         DomainName = "s.com",
    ///         IpAddress = "1.1.1.1",
    ///         LdapServer = trname1.Name,
    ///         Port = 445,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// User DomainController can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/domaincontroller:Domaincontroller labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:user/domaincontroller:Domaincontroller labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:user/domaincontroller:Domaincontroller")]
    public partial class Domaincontroller : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
        /// </summary>
        [Output("adMode")]
        public Output<string> AdMode { get; private set; } = null!;

        /// <summary>
        /// AD LDS distinguished name.
        /// </summary>
        [Output("adldsDn")]
        public Output<string> AdldsDn { get; private set; } = null!;

        /// <summary>
        /// AD LDS IPv6 address.
        /// </summary>
        [Output("adldsIp6")]
        public Output<string> AdldsIp6 { get; private set; } = null!;

        /// <summary>
        /// AD LDS IPv4 address.
        /// </summary>
        [Output("adldsIpAddress")]
        public Output<string> AdldsIpAddress { get; private set; } = null!;

        /// <summary>
        /// Port number of AD LDS service (default = 389).
        /// </summary>
        [Output("adldsPort")]
        public Output<int> AdldsPort { get; private set; } = null!;

        /// <summary>
        /// Enable/disable detection of a configuration change in the Active Directory server. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("changeDetection")]
        public Output<string> ChangeDetection { get; private set; } = null!;

        /// <summary>
        /// Minutes to detect a configuration change in the Active Directory server (5 - 10080 minutes (7 days), default = 60).
        /// </summary>
        [Output("changeDetectionPeriod")]
        public Output<int> ChangeDetectionPeriod { get; private set; } = null!;

        /// <summary>
        /// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("dnsSrvLookup")]
        public Output<string> DnsSrvLookup { get; private set; } = null!;

        /// <summary>
        /// Domain DNS name.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// extra servers. The structure of `extra_server` block is documented below.
        /// </summary>
        [Output("extraServers")]
        public Output<ImmutableArray<Outputs.DomaincontrollerExtraServer>> ExtraServers { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Hostname of the server to connect to.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Domain controller IPv6 address.
        /// </summary>
        [Output("ip6")]
        public Output<string> Ip6 { get; private set; } = null!;

        /// <summary>
        /// Domain controller IP address.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// LDAP server name.
        /// </summary>
        [Output("ldapServer")]
        public Output<string> LdapServer { get; private set; } = null!;

        /// <summary>
        /// Domain controller entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password for specified username.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Port to be used for communication with the domain controller (default = 445).
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
        /// </summary>
        [Output("replicationPort")]
        public Output<int> ReplicationPort { get; private set; } = null!;

        /// <summary>
        /// FortiGate IPv6 address to be used for communication with the domain controller.
        /// </summary>
        [Output("sourceIp6")]
        public Output<string> SourceIp6 { get; private set; } = null!;

        /// <summary>
        /// FortiGate IPv4 address to be used for communication with the domain controller.
        /// </summary>
        [Output("sourceIpAddress")]
        public Output<string> SourceIpAddress { get; private set; } = null!;

        /// <summary>
        /// Source port to be used for communication with the domain controller.
        /// </summary>
        [Output("sourcePort")]
        public Output<int> SourcePort { get; private set; } = null!;

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
        public Domaincontroller(string name, DomaincontrollerArgs args, CustomResourceOptions? options = null)
            : base("fortios:user/domaincontroller:Domaincontroller", name, args ?? new DomaincontrollerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domaincontroller(string name, Input<string> id, DomaincontrollerState? state = null, CustomResourceOptions? options = null)
            : base("fortios:user/domaincontroller:Domaincontroller", name, state, MakeResourceOptions(options, id))
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
        /// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
        /// </summary>
        [Input("adMode")]
        public Input<string>? AdMode { get; set; }

        /// <summary>
        /// AD LDS distinguished name.
        /// </summary>
        [Input("adldsDn")]
        public Input<string>? AdldsDn { get; set; }

        /// <summary>
        /// AD LDS IPv6 address.
        /// </summary>
        [Input("adldsIp6")]
        public Input<string>? AdldsIp6 { get; set; }

        /// <summary>
        /// AD LDS IPv4 address.
        /// </summary>
        [Input("adldsIpAddress")]
        public Input<string>? AdldsIpAddress { get; set; }

        /// <summary>
        /// Port number of AD LDS service (default = 389).
        /// </summary>
        [Input("adldsPort")]
        public Input<int>? AdldsPort { get; set; }

        /// <summary>
        /// Enable/disable detection of a configuration change in the Active Directory server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("changeDetection")]
        public Input<string>? ChangeDetection { get; set; }

        /// <summary>
        /// Minutes to detect a configuration change in the Active Directory server (5 - 10080 minutes (7 days), default = 60).
        /// </summary>
        [Input("changeDetectionPeriod")]
        public Input<int>? ChangeDetectionPeriod { get; set; }

        /// <summary>
        /// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dnsSrvLookup")]
        public Input<string>? DnsSrvLookup { get; set; }

        /// <summary>
        /// Domain DNS name.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("extraServers")]
        private InputList<Inputs.DomaincontrollerExtraServerArgs>? _extraServers;

        /// <summary>
        /// extra servers. The structure of `extra_server` block is documented below.
        /// </summary>
        public InputList<Inputs.DomaincontrollerExtraServerArgs> ExtraServers
        {
            get => _extraServers ?? (_extraServers = new InputList<Inputs.DomaincontrollerExtraServerArgs>());
            set => _extraServers = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Hostname of the server to connect to.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Domain controller IPv6 address.
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        /// <summary>
        /// Domain controller IP address.
        /// </summary>
        [Input("ipAddress", required: true)]
        public Input<string> IpAddress { get; set; } = null!;

        /// <summary>
        /// LDAP server name.
        /// </summary>
        [Input("ldapServer", required: true)]
        public Input<string> LdapServer { get; set; } = null!;

        /// <summary>
        /// Domain controller entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password for specified username.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Port to be used for communication with the domain controller (default = 445).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
        /// </summary>
        [Input("replicationPort")]
        public Input<int>? ReplicationPort { get; set; }

        /// <summary>
        /// FortiGate IPv6 address to be used for communication with the domain controller.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// FortiGate IPv4 address to be used for communication with the domain controller.
        /// </summary>
        [Input("sourceIpAddress")]
        public Input<string>? SourceIpAddress { get; set; }

        /// <summary>
        /// Source port to be used for communication with the domain controller.
        /// </summary>
        [Input("sourcePort")]
        public Input<int>? SourcePort { get; set; }

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
        /// Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
        /// </summary>
        [Input("adMode")]
        public Input<string>? AdMode { get; set; }

        /// <summary>
        /// AD LDS distinguished name.
        /// </summary>
        [Input("adldsDn")]
        public Input<string>? AdldsDn { get; set; }

        /// <summary>
        /// AD LDS IPv6 address.
        /// </summary>
        [Input("adldsIp6")]
        public Input<string>? AdldsIp6 { get; set; }

        /// <summary>
        /// AD LDS IPv4 address.
        /// </summary>
        [Input("adldsIpAddress")]
        public Input<string>? AdldsIpAddress { get; set; }

        /// <summary>
        /// Port number of AD LDS service (default = 389).
        /// </summary>
        [Input("adldsPort")]
        public Input<int>? AdldsPort { get; set; }

        /// <summary>
        /// Enable/disable detection of a configuration change in the Active Directory server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("changeDetection")]
        public Input<string>? ChangeDetection { get; set; }

        /// <summary>
        /// Minutes to detect a configuration change in the Active Directory server (5 - 10080 minutes (7 days), default = 60).
        /// </summary>
        [Input("changeDetectionPeriod")]
        public Input<int>? ChangeDetectionPeriod { get; set; }

        /// <summary>
        /// Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dnsSrvLookup")]
        public Input<string>? DnsSrvLookup { get; set; }

        /// <summary>
        /// Domain DNS name.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        [Input("extraServers")]
        private InputList<Inputs.DomaincontrollerExtraServerGetArgs>? _extraServers;

        /// <summary>
        /// extra servers. The structure of `extra_server` block is documented below.
        /// </summary>
        public InputList<Inputs.DomaincontrollerExtraServerGetArgs> ExtraServers
        {
            get => _extraServers ?? (_extraServers = new InputList<Inputs.DomaincontrollerExtraServerGetArgs>());
            set => _extraServers = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Hostname of the server to connect to.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Domain controller IPv6 address.
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        /// <summary>
        /// Domain controller IP address.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// LDAP server name.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// Domain controller entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password for specified username.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Port to be used for communication with the domain controller (default = 445).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
        /// </summary>
        [Input("replicationPort")]
        public Input<int>? ReplicationPort { get; set; }

        /// <summary>
        /// FortiGate IPv6 address to be used for communication with the domain controller.
        /// </summary>
        [Input("sourceIp6")]
        public Input<string>? SourceIp6 { get; set; }

        /// <summary>
        /// FortiGate IPv4 address to be used for communication with the domain controller.
        /// </summary>
        [Input("sourceIpAddress")]
        public Input<string>? SourceIpAddress { get; set; }

        /// <summary>
        /// Source port to be used for communication with the domain controller.
        /// </summary>
        [Input("sourcePort")]
        public Input<int>? SourcePort { get; set; }

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
