// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SwitchcontrollerManagedswitchSnmpTrapThresholdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CPU usage when trap is sent.
        /// </summary>
        [Input("trapHighCpuThreshold")]
        public Input<int>? TrapHighCpuThreshold { get; set; }

        /// <summary>
        /// Log disk usage when trap is sent.
        /// </summary>
        [Input("trapLogFullThreshold")]
        public Input<int>? TrapLogFullThreshold { get; set; }

        /// <summary>
        /// Memory usage when trap is sent.
        /// </summary>
        [Input("trapLowMemoryThreshold")]
        public Input<int>? TrapLowMemoryThreshold { get; set; }

        public SwitchcontrollerManagedswitchSnmpTrapThresholdArgs()
        {
        }
        public static new SwitchcontrollerManagedswitchSnmpTrapThresholdArgs Empty => new SwitchcontrollerManagedswitchSnmpTrapThresholdArgs();
    }
}
