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

    public sealed class IpsSensorFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action of selected rules. Valid values: `pass`, `block`, `reset`, `default`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Vulnerable application filter.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// Vulnerability location filter.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Enable/disable logging of selected rules. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable packet logging of selected rules. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        /// <summary>
        /// Filter name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Vulnerable OS filter.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// Vulnerable protocol filter.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Quarantine IP or interface. Valid values: `none`, `attacker`.
        /// </summary>
        [Input("quarantine")]
        public Input<string>? Quarantine { get; set; }

        /// <summary>
        /// Duration of quarantine in minute.
        /// </summary>
        [Input("quarantineExpiry")]
        public Input<int>? QuarantineExpiry { get; set; }

        /// <summary>
        /// Enable/disable logging of selected quarantine. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("quarantineLog")]
        public Input<string>? QuarantineLog { get; set; }

        /// <summary>
        /// Vulnerability severity filter.
        /// </summary>
        [Input("severity")]
        public Input<string>? Severity { get; set; }

        /// <summary>
        /// Selected rules status. Valid values: `disable`, `enable`, `default`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public IpsSensorFilterArgs()
        {
        }
        public static new IpsSensorFilterArgs Empty => new IpsSensorFilterArgs();
    }
}