// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    /// <summary>
    /// Configure IPv4 addresses.
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
    ///     var trname = new Fortios.Firewall.Address("trname", new()
    ///     {
    ///         AllowRouting = "disable",
    ///         AssociatedInterface = "port2",
    ///         Color = 3,
    ///         EndIp = "255.255.255.0",
    ///         StartIp = "22.1.1.0",
    ///         Subnet = "22.1.1.0 255.255.255.0",
    ///         Type = "ipmask",
    ///         Visibility = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall Address can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/address:Address labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:firewall/address:Address labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:firewall/address:Address")]
    public partial class Address : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("allowRouting")]
        public Output<string> AllowRouting { get; private set; } = null!;

        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Output("associatedInterface")]
        public Output<string?> AssociatedInterface { get; private set; } = null!;

        /// <summary>
        /// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
        /// </summary>
        [Output("cacheTtl")]
        public Output<int?> CacheTtl { get; private set; } = null!;

        /// <summary>
        /// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
        /// </summary>
        [Output("clearpassSpt")]
        public Output<string> ClearpassSpt { get; private set; } = null!;

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Output("color")]
        public Output<int?> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// IP addresses associated to a specific country.
        /// </summary>
        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Output("endIp")]
        public Output<string> EndIp { get; private set; } = null!;

        /// <summary>
        /// Last MAC address in the range.
        /// </summary>
        [Output("endMac")]
        public Output<string> EndMac { get; private set; } = null!;

        /// <summary>
        /// Endpoint group name.
        /// </summary>
        [Output("epgName")]
        public Output<string?> EpgName { get; private set; } = null!;

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fabricObject")]
        public Output<string> FabricObject { get; private set; } = null!;

        /// <summary>
        /// Match criteria filter.
        /// </summary>
        [Output("filter")]
        public Output<string?> Filter { get; private set; } = null!;

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Output("fqdn")]
        public Output<string?> Fqdn { get; private set; } = null!;

        /// <summary>
        /// FSSO group(s). The structure of `fsso_group` block is documented below.
        /// </summary>
        [Output("fssoGroups")]
        public Output<ImmutableArray<Outputs.AddressFssoGroup>> FssoGroups { get; private set; } = null!;

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Output("getAllTables")]
        public Output<string?> GetAllTables { get; private set; } = null!;

        /// <summary>
        /// Dynamic address matching hardware model.
        /// </summary>
        [Output("hwModel")]
        public Output<string?> HwModel { get; private set; } = null!;

        /// <summary>
        /// Dynamic address matching hardware vendor.
        /// </summary>
        [Output("hwVendor")]
        public Output<string?> HwVendor { get; private set; } = null!;

        /// <summary>
        /// Name of interface whose IP address is to be used.
        /// </summary>
        [Output("interface")]
        public Output<string?> Interface { get; private set; } = null!;

        /// <summary>
        /// IP address list. The structure of `list` block is documented below.
        /// </summary>
        [Output("lists")]
        public Output<ImmutableArray<Outputs.AddressList>> Lists { get; private set; } = null!;

        /// <summary>
        /// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
        /// </summary>
        [Output("macaddrs")]
        public Output<ImmutableArray<Outputs.AddressMacaddr>> Macaddrs { get; private set; } = null!;

        /// <summary>
        /// Address name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("nodeIpOnly")]
        public Output<string> NodeIpOnly { get; private set; } = null!;

        /// <summary>
        /// Object ID for NSX.
        /// </summary>
        [Output("objId")]
        public Output<string?> ObjId { get; private set; } = null!;

        /// <summary>
        /// Tag of dynamic address object.
        /// </summary>
        [Output("objTag")]
        public Output<string?> ObjTag { get; private set; } = null!;

        /// <summary>
        /// Object type. Valid values: `ip`, `mac`.
        /// </summary>
        [Output("objType")]
        public Output<string> ObjType { get; private set; } = null!;

        /// <summary>
        /// Organization domain name (Syntax: organization/domain).
        /// </summary>
        [Output("organization")]
        public Output<string?> Organization { get; private set; } = null!;

        /// <summary>
        /// Dynamic address matching operating system.
        /// </summary>
        [Output("os")]
        public Output<string?> Os { get; private set; } = null!;

        /// <summary>
        /// Policy group name.
        /// </summary>
        [Output("policyGroup")]
        public Output<string?> PolicyGroup { get; private set; } = null!;

        /// <summary>
        /// route-tag address.
        /// </summary>
        [Output("routeTag")]
        public Output<int?> RouteTag { get; private set; } = null!;

        /// <summary>
        /// SDN.
        /// </summary>
        [Output("sdn")]
        public Output<string?> Sdn { get; private set; } = null!;

        /// <summary>
        /// Type of addresses to collect. Valid values: `private`, `public`, `all`.
        /// </summary>
        [Output("sdnAddrType")]
        public Output<string> SdnAddrType { get; private set; } = null!;

        /// <summary>
        /// SDN Tag.
        /// </summary>
        [Output("sdnTag")]
        public Output<string?> SdnTag { get; private set; } = null!;

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Output("startIp")]
        public Output<string> StartIp { get; private set; } = null!;

        /// <summary>
        /// First MAC address in the range.
        /// </summary>
        [Output("startMac")]
        public Output<string> StartMac { get; private set; } = null!;

        /// <summary>
        /// Sub-type of address.
        /// </summary>
        [Output("subType")]
        public Output<string> SubType { get; private set; } = null!;

        /// <summary>
        /// IP address and subnet mask of address.
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Subnet name.
        /// </summary>
        [Output("subnetName")]
        public Output<string?> SubnetName { get; private set; } = null!;

        /// <summary>
        /// Dynamic address matching software version.
        /// </summary>
        [Output("swVersion")]
        public Output<string?> SwVersion { get; private set; } = null!;

        /// <summary>
        /// Tag detection level of dynamic address object.
        /// </summary>
        [Output("tagDetectionLevel")]
        public Output<string?> TagDetectionLevel { get; private set; } = null!;

        /// <summary>
        /// Tag type of dynamic address object.
        /// </summary>
        [Output("tagType")]
        public Output<string?> TagType { get; private set; } = null!;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        [Output("taggings")]
        public Output<ImmutableArray<Outputs.AddressTagging>> Taggings { get; private set; } = null!;

        /// <summary>
        /// Tenant.
        /// </summary>
        [Output("tenant")]
        public Output<string?> Tenant { get; private set; } = null!;

        /// <summary>
        /// Type of address.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;

        /// <summary>
        /// IP address and wildcard netmask.
        /// </summary>
        [Output("wildcard")]
        public Output<string> Wildcard { get; private set; } = null!;

        /// <summary>
        /// Fully Qualified Domain Name with wildcard characters.
        /// </summary>
        [Output("wildcardFqdn")]
        public Output<string?> WildcardFqdn { get; private set; } = null!;


        /// <summary>
        /// Create a Address resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Address(string name, AddressArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/address:Address", name, args ?? new AddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Address(string name, Input<string> id, AddressState? state = null, CustomResourceOptions? options = null)
            : base("fortios:firewall/address:Address", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Address resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Address Get(string name, Input<string> id, AddressState? state = null, CustomResourceOptions? options = null)
        {
            return new Address(name, id, state, options);
        }
    }

    public sealed class AddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowRouting")]
        public Input<string>? AllowRouting { get; set; }

        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Input("associatedInterface")]
        public Input<string>? AssociatedInterface { get; set; }

        /// <summary>
        /// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
        /// </summary>
        [Input("cacheTtl")]
        public Input<int>? CacheTtl { get; set; }

        /// <summary>
        /// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
        /// </summary>
        [Input("clearpassSpt")]
        public Input<string>? ClearpassSpt { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// IP addresses associated to a specific country.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Last MAC address in the range.
        /// </summary>
        [Input("endMac")]
        public Input<string>? EndMac { get; set; }

        /// <summary>
        /// Endpoint group name.
        /// </summary>
        [Input("epgName")]
        public Input<string>? EpgName { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Match criteria filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        [Input("fssoGroups")]
        private InputList<Inputs.AddressFssoGroupArgs>? _fssoGroups;

        /// <summary>
        /// FSSO group(s). The structure of `fsso_group` block is documented below.
        /// </summary>
        public InputList<Inputs.AddressFssoGroupArgs> FssoGroups
        {
            get => _fssoGroups ?? (_fssoGroups = new InputList<Inputs.AddressFssoGroupArgs>());
            set => _fssoGroups = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Dynamic address matching hardware model.
        /// </summary>
        [Input("hwModel")]
        public Input<string>? HwModel { get; set; }

        /// <summary>
        /// Dynamic address matching hardware vendor.
        /// </summary>
        [Input("hwVendor")]
        public Input<string>? HwVendor { get; set; }

        /// <summary>
        /// Name of interface whose IP address is to be used.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        [Input("lists")]
        private InputList<Inputs.AddressListArgs>? _lists;

        /// <summary>
        /// IP address list. The structure of `list` block is documented below.
        /// </summary>
        public InputList<Inputs.AddressListArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.AddressListArgs>());
            set => _lists = value;
        }

        [Input("macaddrs")]
        private InputList<Inputs.AddressMacaddrArgs>? _macaddrs;

        /// <summary>
        /// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.AddressMacaddrArgs> Macaddrs
        {
            get => _macaddrs ?? (_macaddrs = new InputList<Inputs.AddressMacaddrArgs>());
            set => _macaddrs = value;
        }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nodeIpOnly")]
        public Input<string>? NodeIpOnly { get; set; }

        /// <summary>
        /// Object ID for NSX.
        /// </summary>
        [Input("objId")]
        public Input<string>? ObjId { get; set; }

        /// <summary>
        /// Tag of dynamic address object.
        /// </summary>
        [Input("objTag")]
        public Input<string>? ObjTag { get; set; }

        /// <summary>
        /// Object type. Valid values: `ip`, `mac`.
        /// </summary>
        [Input("objType")]
        public Input<string>? ObjType { get; set; }

        /// <summary>
        /// Organization domain name (Syntax: organization/domain).
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// Dynamic address matching operating system.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Policy group name.
        /// </summary>
        [Input("policyGroup")]
        public Input<string>? PolicyGroup { get; set; }

        /// <summary>
        /// route-tag address.
        /// </summary>
        [Input("routeTag")]
        public Input<int>? RouteTag { get; set; }

        /// <summary>
        /// SDN.
        /// </summary>
        [Input("sdn")]
        public Input<string>? Sdn { get; set; }

        /// <summary>
        /// Type of addresses to collect. Valid values: `private`, `public`, `all`.
        /// </summary>
        [Input("sdnAddrType")]
        public Input<string>? SdnAddrType { get; set; }

        /// <summary>
        /// SDN Tag.
        /// </summary>
        [Input("sdnTag")]
        public Input<string>? SdnTag { get; set; }

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// First MAC address in the range.
        /// </summary>
        [Input("startMac")]
        public Input<string>? StartMac { get; set; }

        /// <summary>
        /// Sub-type of address.
        /// </summary>
        [Input("subType")]
        public Input<string>? SubType { get; set; }

        /// <summary>
        /// IP address and subnet mask of address.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Subnet name.
        /// </summary>
        [Input("subnetName")]
        public Input<string>? SubnetName { get; set; }

        /// <summary>
        /// Dynamic address matching software version.
        /// </summary>
        [Input("swVersion")]
        public Input<string>? SwVersion { get; set; }

        /// <summary>
        /// Tag detection level of dynamic address object.
        /// </summary>
        [Input("tagDetectionLevel")]
        public Input<string>? TagDetectionLevel { get; set; }

        /// <summary>
        /// Tag type of dynamic address object.
        /// </summary>
        [Input("tagType")]
        public Input<string>? TagType { get; set; }

        [Input("taggings")]
        private InputList<Inputs.AddressTaggingArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.AddressTaggingArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.AddressTaggingArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Tenant.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// Type of address.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// IP address and wildcard netmask.
        /// </summary>
        [Input("wildcard")]
        public Input<string>? Wildcard { get; set; }

        /// <summary>
        /// Fully Qualified Domain Name with wildcard characters.
        /// </summary>
        [Input("wildcardFqdn")]
        public Input<string>? WildcardFqdn { get; set; }

        public AddressArgs()
        {
        }
        public static new AddressArgs Empty => new AddressArgs();
    }

    public sealed class AddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowRouting")]
        public Input<string>? AllowRouting { get; set; }

        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Input("associatedInterface")]
        public Input<string>? AssociatedInterface { get; set; }

        /// <summary>
        /// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
        /// </summary>
        [Input("cacheTtl")]
        public Input<int>? CacheTtl { get; set; }

        /// <summary>
        /// SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
        /// </summary>
        [Input("clearpassSpt")]
        public Input<string>? ClearpassSpt { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// IP addresses associated to a specific country.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Last MAC address in the range.
        /// </summary>
        [Input("endMac")]
        public Input<string>? EndMac { get; set; }

        /// <summary>
        /// Endpoint group name.
        /// </summary>
        [Input("epgName")]
        public Input<string>? EpgName { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Match criteria filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        [Input("fssoGroups")]
        private InputList<Inputs.AddressFssoGroupGetArgs>? _fssoGroups;

        /// <summary>
        /// FSSO group(s). The structure of `fsso_group` block is documented below.
        /// </summary>
        public InputList<Inputs.AddressFssoGroupGetArgs> FssoGroups
        {
            get => _fssoGroups ?? (_fssoGroups = new InputList<Inputs.AddressFssoGroupGetArgs>());
            set => _fssoGroups = value;
        }

        /// <summary>
        /// Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        /// </summary>
        [Input("getAllTables")]
        public Input<string>? GetAllTables { get; set; }

        /// <summary>
        /// Dynamic address matching hardware model.
        /// </summary>
        [Input("hwModel")]
        public Input<string>? HwModel { get; set; }

        /// <summary>
        /// Dynamic address matching hardware vendor.
        /// </summary>
        [Input("hwVendor")]
        public Input<string>? HwVendor { get; set; }

        /// <summary>
        /// Name of interface whose IP address is to be used.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        [Input("lists")]
        private InputList<Inputs.AddressListGetArgs>? _lists;

        /// <summary>
        /// IP address list. The structure of `list` block is documented below.
        /// </summary>
        public InputList<Inputs.AddressListGetArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.AddressListGetArgs>());
            set => _lists = value;
        }

        [Input("macaddrs")]
        private InputList<Inputs.AddressMacaddrGetArgs>? _macaddrs;

        /// <summary>
        /// Multiple MAC address ranges. The structure of `macaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.AddressMacaddrGetArgs> Macaddrs
        {
            get => _macaddrs ?? (_macaddrs = new InputList<Inputs.AddressMacaddrGetArgs>());
            set => _macaddrs = value;
        }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nodeIpOnly")]
        public Input<string>? NodeIpOnly { get; set; }

        /// <summary>
        /// Object ID for NSX.
        /// </summary>
        [Input("objId")]
        public Input<string>? ObjId { get; set; }

        /// <summary>
        /// Tag of dynamic address object.
        /// </summary>
        [Input("objTag")]
        public Input<string>? ObjTag { get; set; }

        /// <summary>
        /// Object type. Valid values: `ip`, `mac`.
        /// </summary>
        [Input("objType")]
        public Input<string>? ObjType { get; set; }

        /// <summary>
        /// Organization domain name (Syntax: organization/domain).
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// Dynamic address matching operating system.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Policy group name.
        /// </summary>
        [Input("policyGroup")]
        public Input<string>? PolicyGroup { get; set; }

        /// <summary>
        /// route-tag address.
        /// </summary>
        [Input("routeTag")]
        public Input<int>? RouteTag { get; set; }

        /// <summary>
        /// SDN.
        /// </summary>
        [Input("sdn")]
        public Input<string>? Sdn { get; set; }

        /// <summary>
        /// Type of addresses to collect. Valid values: `private`, `public`, `all`.
        /// </summary>
        [Input("sdnAddrType")]
        public Input<string>? SdnAddrType { get; set; }

        /// <summary>
        /// SDN Tag.
        /// </summary>
        [Input("sdnTag")]
        public Input<string>? SdnTag { get; set; }

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// First MAC address in the range.
        /// </summary>
        [Input("startMac")]
        public Input<string>? StartMac { get; set; }

        /// <summary>
        /// Sub-type of address.
        /// </summary>
        [Input("subType")]
        public Input<string>? SubType { get; set; }

        /// <summary>
        /// IP address and subnet mask of address.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Subnet name.
        /// </summary>
        [Input("subnetName")]
        public Input<string>? SubnetName { get; set; }

        /// <summary>
        /// Dynamic address matching software version.
        /// </summary>
        [Input("swVersion")]
        public Input<string>? SwVersion { get; set; }

        /// <summary>
        /// Tag detection level of dynamic address object.
        /// </summary>
        [Input("tagDetectionLevel")]
        public Input<string>? TagDetectionLevel { get; set; }

        /// <summary>
        /// Tag type of dynamic address object.
        /// </summary>
        [Input("tagType")]
        public Input<string>? TagType { get; set; }

        [Input("taggings")]
        private InputList<Inputs.AddressTaggingGetArgs>? _taggings;

        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public InputList<Inputs.AddressTaggingGetArgs> Taggings
        {
            get => _taggings ?? (_taggings = new InputList<Inputs.AddressTaggingGetArgs>());
            set => _taggings = value;
        }

        /// <summary>
        /// Tenant.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        /// <summary>
        /// Type of address.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// IP address and wildcard netmask.
        /// </summary>
        [Input("wildcard")]
        public Input<string>? Wildcard { get; set; }

        /// <summary>
        /// Fully Qualified Domain Name with wildcard characters.
        /// </summary>
        [Input("wildcardFqdn")]
        public Input<string>? WildcardFqdn { get; set; }

        public AddressState()
        {
        }
        public static new AddressState Empty => new AddressState();
    }
}
