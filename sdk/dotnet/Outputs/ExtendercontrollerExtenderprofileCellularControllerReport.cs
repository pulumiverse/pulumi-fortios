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
    public sealed class ExtendercontrollerExtenderprofileCellularControllerReport
    {
        /// <summary>
        /// Controller report interval.
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// Controller report signal threshold.
        /// </summary>
        public readonly int? SignalThreshold;
        /// <summary>
        /// FortiExtender controller report status. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ExtendercontrollerExtenderprofileCellularControllerReport(
            int? interval,

            int? signalThreshold,

            string? status)
        {
            Interval = interval;
            SignalThreshold = signalThreshold;
            Status = status;
        }
    }
}
