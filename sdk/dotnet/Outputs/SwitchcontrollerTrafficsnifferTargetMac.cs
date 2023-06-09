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
    public sealed class SwitchcontrollerTrafficsnifferTargetMac
    {
        /// <summary>
        /// Description for the sniffer MAC.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Sniffer MAC.
        /// </summary>
        public readonly string? Mac;

        [OutputConstructor]
        private SwitchcontrollerTrafficsnifferTargetMac(
            string? description,

            string? mac)
        {
            Description = description;
            Mac = mac;
        }
    }
}
