// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class GetRipRedistributeResult
    {
        /// <summary>
        /// Redistribute metric setting.
        /// </summary>
        public readonly int Metric;
        /// <summary>
        /// Interface name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Route map name.
        /// </summary>
        public readonly string Routemap;
        /// <summary>
        /// status
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetRipRedistributeResult(
            int metric,

            string name,

            string routemap,

            string status)
        {
            Metric = metric;
            Name = name;
            Routemap = routemap;
            Status = status;
        }
    }
}
