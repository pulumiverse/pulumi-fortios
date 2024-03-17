// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class Ospf6AreaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication mode. Valid values: `none`, `ah`, `esp`.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Summary default cost of stub or NSSA area.
        /// </summary>
        [Input("defaultCost")]
        public Input<int>? DefaultCost { get; set; }

        /// <summary>
        /// Area entry IP address.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
        /// </summary>
        [Input("ipsecAuthAlg")]
        public Input<string>? IpsecAuthAlg { get; set; }

        /// <summary>
        /// Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
        /// </summary>
        [Input("ipsecEncAlg")]
        public Input<string>? IpsecEncAlg { get; set; }

        [Input("ipsecKeys")]
        private InputList<Inputs.Ospf6AreaIpsecKeyGetArgs>? _ipsecKeys;

        /// <summary>
        /// IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6AreaIpsecKeyGetArgs> IpsecKeys
        {
            get => _ipsecKeys ?? (_ipsecKeys = new InputList<Inputs.Ospf6AreaIpsecKeyGetArgs>());
            set => _ipsecKeys = value;
        }

        /// <summary>
        /// Key roll-over interval.
        /// </summary>
        [Input("keyRolloverInterval")]
        public Input<int>? KeyRolloverInterval { get; set; }

        /// <summary>
        /// Enable/disable originate type 7 default into NSSA area. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nssaDefaultInformationOriginate")]
        public Input<string>? NssaDefaultInformationOriginate { get; set; }

        /// <summary>
        /// OSPFv3 default metric.
        /// </summary>
        [Input("nssaDefaultInformationOriginateMetric")]
        public Input<int>? NssaDefaultInformationOriginateMetric { get; set; }

        /// <summary>
        /// OSPFv3 metric type for default routes. Valid values: `1`, `2`.
        /// </summary>
        [Input("nssaDefaultInformationOriginateMetricType")]
        public Input<string>? NssaDefaultInformationOriginateMetricType { get; set; }

        /// <summary>
        /// Enable/disable redistribute into NSSA area. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nssaRedistribution")]
        public Input<string>? NssaRedistribution { get; set; }

        /// <summary>
        /// NSSA translator role type. Valid values: `candidate`, `never`, `always`.
        /// </summary>
        [Input("nssaTranslatorRole")]
        public Input<string>? NssaTranslatorRole { get; set; }

        [Input("ranges")]
        private InputList<Inputs.Ospf6AreaRangeGetArgs>? _ranges;

        /// <summary>
        /// OSPF6 area range configuration. The structure of `range` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6AreaRangeGetArgs> Ranges
        {
            get => _ranges ?? (_ranges = new InputList<Inputs.Ospf6AreaRangeGetArgs>());
            set => _ranges = value;
        }

        /// <summary>
        /// Stub summary setting. Valid values: `no-summary`, `summary`.
        /// </summary>
        [Input("stubType")]
        public Input<string>? StubType { get; set; }

        /// <summary>
        /// Area type setting. Valid values: `regular`, `nssa`, `stub`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("virtualLinks")]
        private InputList<Inputs.Ospf6AreaVirtualLinkGetArgs>? _virtualLinks;

        /// <summary>
        /// OSPF6 virtual link configuration. The structure of `virtual_link` block is documented below.
        /// </summary>
        public InputList<Inputs.Ospf6AreaVirtualLinkGetArgs> VirtualLinks
        {
            get => _virtualLinks ?? (_virtualLinks = new InputList<Inputs.Ospf6AreaVirtualLinkGetArgs>());
            set => _virtualLinks = value;
        }

        public Ospf6AreaGetArgs()
        {
        }
        public static new Ospf6AreaGetArgs Empty => new Ospf6AreaGetArgs();
    }
}
