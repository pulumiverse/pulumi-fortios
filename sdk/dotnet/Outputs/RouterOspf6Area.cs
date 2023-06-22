// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class RouterOspf6Area
    {
        /// <summary>
        /// Authentication mode. Valid values: `none`, `ah`, `esp`.
        /// </summary>
        public readonly string? Authentication;
        /// <summary>
        /// Summary default cost of stub or NSSA area.
        /// </summary>
        public readonly int? DefaultCost;
        /// <summary>
        /// Area entry IP address.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        /// </summary>
        public readonly string? IpsecAuthAlg;
        /// <summary>
        /// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
        /// </summary>
        public readonly string? IpsecEncAlg;
        /// <summary>
        /// IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterOspf6AreaIpsecKey> IpsecKeys;
        /// <summary>
        /// Key roll-over interval.
        /// </summary>
        public readonly int? KeyRolloverInterval;
        /// <summary>
        /// Enable/disable originate type 7 default into NSSA area. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? NssaDefaultInformationOriginate;
        /// <summary>
        /// OSPFv3 default metric.
        /// </summary>
        public readonly int? NssaDefaultInformationOriginateMetric;
        /// <summary>
        /// OSPFv3 metric type for default routes. Valid values: `1`, `2`.
        /// </summary>
        public readonly string? NssaDefaultInformationOriginateMetricType;
        /// <summary>
        /// Enable/disable redistribute into NSSA area. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? NssaRedistribution;
        /// <summary>
        /// NSSA translator role type. Valid values: `candidate`, `never`, `always`.
        /// </summary>
        public readonly string? NssaTranslatorRole;
        /// <summary>
        /// OSPF6 area range configuration. The structure of `range` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterOspf6AreaRange> Ranges;
        /// <summary>
        /// Stub summary setting. Valid values: `no-summary`, `summary`.
        /// </summary>
        public readonly string? StubType;
        /// <summary>
        /// Area type setting. Valid values: `regular`, `nssa`, `stub`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// OSPF6 virtual link configuration. The structure of `virtual_link` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterOspf6AreaVirtualLink> VirtualLinks;

        [OutputConstructor]
        private RouterOspf6Area(
            string? authentication,

            int? defaultCost,

            string? id,

            string? ipsecAuthAlg,

            string? ipsecEncAlg,

            ImmutableArray<Outputs.RouterOspf6AreaIpsecKey> ipsecKeys,

            int? keyRolloverInterval,

            string? nssaDefaultInformationOriginate,

            int? nssaDefaultInformationOriginateMetric,

            string? nssaDefaultInformationOriginateMetricType,

            string? nssaRedistribution,

            string? nssaTranslatorRole,

            ImmutableArray<Outputs.RouterOspf6AreaRange> ranges,

            string? stubType,

            string? type,

            ImmutableArray<Outputs.RouterOspf6AreaVirtualLink> virtualLinks)
        {
            Authentication = authentication;
            DefaultCost = defaultCost;
            Id = id;
            IpsecAuthAlg = ipsecAuthAlg;
            IpsecEncAlg = ipsecEncAlg;
            IpsecKeys = ipsecKeys;
            KeyRolloverInterval = keyRolloverInterval;
            NssaDefaultInformationOriginate = nssaDefaultInformationOriginate;
            NssaDefaultInformationOriginateMetric = nssaDefaultInformationOriginateMetric;
            NssaDefaultInformationOriginateMetricType = nssaDefaultInformationOriginateMetricType;
            NssaRedistribution = nssaRedistribution;
            NssaTranslatorRole = nssaTranslatorRole;
            Ranges = ranges;
            StubType = stubType;
            Type = type;
            VirtualLinks = virtualLinks;
        }
    }
}
