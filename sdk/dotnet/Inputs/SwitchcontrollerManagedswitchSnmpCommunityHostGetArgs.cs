// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SwitchcontrollerManagedswitchSnmpCommunityHostGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IPv4 address of the SNMP manager (host).
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public SwitchcontrollerManagedswitchSnmpCommunityHostGetArgs()
        {
        }
        public static new SwitchcontrollerManagedswitchSnmpCommunityHostGetArgs Empty => new SwitchcontrollerManagedswitchSnmpCommunityHostGetArgs();
    }
}
