// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class HaSecondaryVclusterArgs : global::Pulumi.ResourceArgs
    {
        [Input("monitor")]
        public Input<string>? Monitor { get; set; }

        [Input("override")]
        public Input<string>? Override { get; set; }

        [Input("overrideWaitTime")]
        public Input<int>? OverrideWaitTime { get; set; }

        [Input("pingserverFailoverThreshold")]
        public Input<int>? PingserverFailoverThreshold { get; set; }

        [Input("pingserverMonitorInterface")]
        public Input<string>? PingserverMonitorInterface { get; set; }

        [Input("pingserverSecondaryForceReset")]
        public Input<string>? PingserverSecondaryForceReset { get; set; }

        [Input("pingserverSlaveForceReset")]
        public Input<string>? PingserverSlaveForceReset { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("vclusterId")]
        public Input<int>? VclusterId { get; set; }

        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        public HaSecondaryVclusterArgs()
        {
        }
        public static new HaSecondaryVclusterArgs Empty => new HaSecondaryVclusterArgs();
    }
}
