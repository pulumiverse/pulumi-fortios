// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Outputs
{

    [OutputType]
    public sealed class WtpSplitTunnelingAcl
    {
        /// <summary>
        /// Destination IP and mask for the split-tunneling subnet.
        /// </summary>
        public readonly string? DestIp;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;

        [OutputConstructor]
        private WtpSplitTunnelingAcl(
            string? destIp,

            int? id)
        {
            DestIp = destIp;
            Id = id;
        }
    }
}