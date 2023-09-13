// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class TrafficsnifferTargetPortGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the sniffer port entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inPorts")]
        private InputList<Inputs.TrafficsnifferTargetPortInPortGetArgs>? _inPorts;

        /// <summary>
        /// Configure source ingress port interfaces. The structure of `in_ports` block is documented below.
        /// </summary>
        public InputList<Inputs.TrafficsnifferTargetPortInPortGetArgs> InPorts
        {
            get => _inPorts ?? (_inPorts = new InputList<Inputs.TrafficsnifferTargetPortInPortGetArgs>());
            set => _inPorts = value;
        }

        [Input("outPorts")]
        private InputList<Inputs.TrafficsnifferTargetPortOutPortGetArgs>? _outPorts;

        /// <summary>
        /// Configure source egress port interfaces. The structure of `out_ports` block is documented below.
        /// </summary>
        public InputList<Inputs.TrafficsnifferTargetPortOutPortGetArgs> OutPorts
        {
            get => _outPorts ?? (_outPorts = new InputList<Inputs.TrafficsnifferTargetPortOutPortGetArgs>());
            set => _outPorts = value;
        }

        /// <summary>
        /// Managed-switch ID.
        /// </summary>
        [Input("switchId")]
        public Input<string>? SwitchId { get; set; }

        public TrafficsnifferTargetPortGetArgs()
        {
        }
        public static new TrafficsnifferTargetPortGetArgs Empty => new TrafficsnifferTargetPortGetArgs();
    }
}
