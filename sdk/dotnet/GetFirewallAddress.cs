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
    public static class GetFirewallAddress
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall address
        /// </summary>
        public static Task<GetFirewallAddressResult> InvokeAsync(GetFirewallAddressArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallAddressResult>("fortios:index/getFirewallAddress:getFirewallAddress", args ?? new GetFirewallAddressArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall address
        /// </summary>
        public static Output<GetFirewallAddressResult> Invoke(GetFirewallAddressInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallAddressResult>("fortios:index/getFirewallAddress:getFirewallAddress", args ?? new GetFirewallAddressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallAddressArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall address.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallAddressArgs()
        {
        }
        public static new GetFirewallAddressArgs Empty => new GetFirewallAddressArgs();
    }

    public sealed class GetFirewallAddressInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall address.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallAddressInvokeArgs()
        {
        }
        public static new GetFirewallAddressInvokeArgs Empty => new GetFirewallAddressInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallAddressResult
    {
        /// <summary>
        /// Enable/disable use of this address in the static route configuration.
        /// </summary>
        public readonly string AllowRouting;
        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        public readonly string AssociatedInterface;
        /// <summary>
        /// Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
        /// </summary>
        public readonly int CacheTtl;
        /// <summary>
        /// SPT (System Posture Token) value.
        /// </summary>
        public readonly string ClearpassSpt;
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        public readonly int Color;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// IP addresses associated to a specific country.
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        public readonly string EndIp;
        /// <summary>
        /// Last MAC address in the range.
        /// </summary>
        public readonly string EndMac;
        /// <summary>
        /// Endpoint group name.
        /// </summary>
        public readonly string EpgName;
        /// <summary>
        /// Security Fabric global object setting.
        /// </summary>
        public readonly string FabricObject;
        /// <summary>
        /// Match criteria filter.
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// FSSO group(s). The structure of `fsso_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallAddressFssoGroupResult> FssoGroups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of interface whose IP address is to be used.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// IP address list. The structure of `list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallAddressListResult> Lists;
        /// <summary>
        /// MAC address ranges &lt;start&gt;[-&lt;end&gt;] separated by space.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallAddressMacaddrResult> Macaddrs;
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Enable/disable collection of node addresses only in Kubernetes.
        /// </summary>
        public readonly string NodeIpOnly;
        /// <summary>
        /// Object ID for NSX.
        /// </summary>
        public readonly string ObjId;
        /// <summary>
        /// Tag of dynamic address object.
        /// </summary>
        public readonly string ObjTag;
        /// <summary>
        /// Object type.
        /// </summary>
        public readonly string ObjType;
        /// <summary>
        /// Organization domain name (Syntax: organization/domain).
        /// </summary>
        public readonly string Organization;
        /// <summary>
        /// Policy group name.
        /// </summary>
        public readonly string PolicyGroup;
        /// <summary>
        /// SDN.
        /// </summary>
        public readonly string Sdn;
        /// <summary>
        /// Type of addresses to collect.
        /// </summary>
        public readonly string SdnAddrType;
        /// <summary>
        /// SDN Tag.
        /// </summary>
        public readonly string SdnTag;
        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        public readonly string StartIp;
        /// <summary>
        /// First MAC address in the range.
        /// </summary>
        public readonly string StartMac;
        /// <summary>
        /// Sub-type of address.
        /// </summary>
        public readonly string SubType;
        /// <summary>
        /// IP address and subnet mask of address.
        /// </summary>
        public readonly string Subnet;
        /// <summary>
        /// Subnet name.
        /// </summary>
        public readonly string SubnetName;
        /// <summary>
        /// Tag detection level of dynamic address object.
        /// </summary>
        public readonly string TagDetectionLevel;
        /// <summary>
        /// Tag type of dynamic address object.
        /// </summary>
        public readonly string TagType;
        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallAddressTaggingResult> Taggings;
        /// <summary>
        /// Tenant.
        /// </summary>
        public readonly string Tenant;
        /// <summary>
        /// Type of address.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable address visibility in the GUI.
        /// </summary>
        public readonly string Visibility;
        /// <summary>
        /// IP address and wildcard netmask.
        /// </summary>
        public readonly string Wildcard;
        /// <summary>
        /// Fully Qualified Domain Name with wildcard characters.
        /// </summary>
        public readonly string WildcardFqdn;

        [OutputConstructor]
        private GetFirewallAddressResult(
            string allowRouting,

            string associatedInterface,

            int cacheTtl,

            string clearpassSpt,

            int color,

            string comment,

            string country,

            string endIp,

            string endMac,

            string epgName,

            string fabricObject,

            string filter,

            string fqdn,

            ImmutableArray<Outputs.GetFirewallAddressFssoGroupResult> fssoGroups,

            string id,

            string @interface,

            ImmutableArray<Outputs.GetFirewallAddressListResult> lists,

            ImmutableArray<Outputs.GetFirewallAddressMacaddrResult> macaddrs,

            string name,

            string nodeIpOnly,

            string objId,

            string objTag,

            string objType,

            string organization,

            string policyGroup,

            string sdn,

            string sdnAddrType,

            string sdnTag,

            string startIp,

            string startMac,

            string subType,

            string subnet,

            string subnetName,

            string tagDetectionLevel,

            string tagType,

            ImmutableArray<Outputs.GetFirewallAddressTaggingResult> taggings,

            string tenant,

            string type,

            string uuid,

            string? vdomparam,

            string visibility,

            string wildcard,

            string wildcardFqdn)
        {
            AllowRouting = allowRouting;
            AssociatedInterface = associatedInterface;
            CacheTtl = cacheTtl;
            ClearpassSpt = clearpassSpt;
            Color = color;
            Comment = comment;
            Country = country;
            EndIp = endIp;
            EndMac = endMac;
            EpgName = epgName;
            FabricObject = fabricObject;
            Filter = filter;
            Fqdn = fqdn;
            FssoGroups = fssoGroups;
            Id = id;
            Interface = @interface;
            Lists = lists;
            Macaddrs = macaddrs;
            Name = name;
            NodeIpOnly = nodeIpOnly;
            ObjId = objId;
            ObjTag = objTag;
            ObjType = objType;
            Organization = organization;
            PolicyGroup = policyGroup;
            Sdn = sdn;
            SdnAddrType = sdnAddrType;
            SdnTag = sdnTag;
            StartIp = startIp;
            StartMac = startMac;
            SubType = subType;
            Subnet = subnet;
            SubnetName = subnetName;
            TagDetectionLevel = tagDetectionLevel;
            TagType = tagType;
            Taggings = taggings;
            Tenant = tenant;
            Type = type;
            Uuid = uuid;
            Vdomparam = vdomparam;
            Visibility = visibility;
            Wildcard = wildcard;
            WildcardFqdn = wildcardFqdn;
        }
    }
}