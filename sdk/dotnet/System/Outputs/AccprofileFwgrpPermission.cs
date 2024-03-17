// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class AccprofileFwgrpPermission
    {
        /// <summary>
        /// Address Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// Other Firewall Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? Others;
        /// <summary>
        /// Policy Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? Policy;
        /// <summary>
        /// Schedule Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? Schedule;
        /// <summary>
        /// Service Configuration. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        public readonly string? Service;

        [OutputConstructor]
        private AccprofileFwgrpPermission(
            string? address,

            string? others,

            string? policy,

            string? schedule,

            string? service)
        {
            Address = address;
            Others = others;
            Policy = policy;
            Schedule = schedule;
            Service = service;
        }
    }
}
