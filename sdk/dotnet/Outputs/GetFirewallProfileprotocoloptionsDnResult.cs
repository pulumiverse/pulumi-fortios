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
    public sealed class GetFirewallProfileprotocoloptionsDnResult
    {
        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 445).
        /// </summary>
        public readonly int Ports;
        /// <summary>
        /// Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetFirewallProfileprotocoloptionsDnResult(
            int ports,

            string status)
        {
            Ports = ports;
            Status = status;
        }
    }
}
