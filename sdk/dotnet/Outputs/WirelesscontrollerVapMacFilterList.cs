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
    public sealed class WirelesscontrollerVapMacFilterList
    {
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// MAC address.
        /// </summary>
        public readonly string? Mac;
        /// <summary>
        /// Deny or allow the client with this MAC address. Valid values: `allow`, `deny`.
        /// </summary>
        public readonly string? MacFilterPolicy;

        [OutputConstructor]
        private WirelesscontrollerVapMacFilterList(
            int? id,

            string? mac,

            string? macFilterPolicy)
        {
            Id = id;
            Mac = mac;
            MacFilterPolicy = macFilterPolicy;
        }
    }
}