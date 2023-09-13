// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class GetAccprofileFwgrpPermissionResult
    {
        /// <summary>
        /// Address Configuration.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// Other Firewall Configuration.
        /// </summary>
        public readonly string Others;
        /// <summary>
        /// Policy Configuration.
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// Schedule Configuration.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// Service Configuration.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private GetAccprofileFwgrpPermissionResult(
            string address,

            string others,

            string policy,

            string schedule,

            string service)
        {
            Address = address;
            Others = others;
            Policy = policy;
            Schedule = schedule;
            Service = service;
        }
    }
}
