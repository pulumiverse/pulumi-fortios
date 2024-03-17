// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class ManagedswitchSnmpCommunityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `ent-conf-change`.
        /// </summary>
        [Input("events")]
        public Input<string>? Events { get; set; }

        [Input("hosts")]
        private InputList<Inputs.ManagedswitchSnmpCommunityHostArgs>? _hosts;

        /// <summary>
        /// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
        /// </summary>
        public InputList<Inputs.ManagedswitchSnmpCommunityHostArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.ManagedswitchSnmpCommunityHostArgs>());
            set => _hosts = value;
        }

        /// <summary>
        /// SNMP community ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// SNMP community name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// SNMP v1 query port (default = 161).
        /// </summary>
        [Input("queryV1Port")]
        public Input<int>? QueryV1Port { get; set; }

        /// <summary>
        /// Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("queryV1Status")]
        public Input<string>? QueryV1Status { get; set; }

        /// <summary>
        /// SNMP v2c query port (default = 161).
        /// </summary>
        [Input("queryV2cPort")]
        public Input<int>? QueryV2cPort { get; set; }

        /// <summary>
        /// Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("queryV2cStatus")]
        public Input<string>? QueryV2cStatus { get; set; }

        /// <summary>
        /// Enable/disable this SNMP community. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// SNMP v2c trap local port (default = 162).
        /// </summary>
        [Input("trapV1Lport")]
        public Input<int>? TrapV1Lport { get; set; }

        /// <summary>
        /// SNMP v2c trap remote port (default = 162).
        /// </summary>
        [Input("trapV1Rport")]
        public Input<int>? TrapV1Rport { get; set; }

        /// <summary>
        /// Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("trapV1Status")]
        public Input<string>? TrapV1Status { get; set; }

        /// <summary>
        /// SNMP v2c trap local port (default = 162).
        /// </summary>
        [Input("trapV2cLport")]
        public Input<int>? TrapV2cLport { get; set; }

        /// <summary>
        /// SNMP v2c trap remote port (default = 162).
        /// </summary>
        [Input("trapV2cRport")]
        public Input<int>? TrapV2cRport { get; set; }

        /// <summary>
        /// Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("trapV2cStatus")]
        public Input<string>? TrapV2cStatus { get; set; }

        public ManagedswitchSnmpCommunityArgs()
        {
        }
        public static new ManagedswitchSnmpCommunityArgs Empty => new ManagedswitchSnmpCommunityArgs();
    }
}
