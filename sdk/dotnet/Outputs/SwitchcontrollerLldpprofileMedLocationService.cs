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
    public sealed class SwitchcontrollerLldpprofileMedLocationService
    {
        /// <summary>
        /// Location service type name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Enable or disable this TLV. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Location service ID.
        /// </summary>
        public readonly string? SysLocationId;

        [OutputConstructor]
        private SwitchcontrollerLldpprofileMedLocationService(
            string? name,

            string? status,

            string? sysLocationId)
        {
            Name = name;
            Status = status;
            SysLocationId = sysLocationId;
        }
    }
}