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
    public sealed class SwitchcontrollerManagedswitchSnmpSysinfo
    {
        /// <summary>
        /// Contact information.
        /// </summary>
        public readonly string? ContactInfo;
        /// <summary>
        /// System description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Local SNMP engine ID string (max 24 char).
        /// </summary>
        public readonly string? EngineId;
        /// <summary>
        /// System location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Enable/disable SNMP. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private SwitchcontrollerManagedswitchSnmpSysinfo(
            string? contactInfo,

            string? description,

            string? engineId,

            string? location,

            string? status)
        {
            ContactInfo = contactInfo;
            Description = description;
            EngineId = engineId;
            Location = location;
            Status = status;
        }
    }
}
