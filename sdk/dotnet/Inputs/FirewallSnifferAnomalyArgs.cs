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

    public sealed class FirewallSnifferAnomalyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action taken when the threshold is reached.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Enable/disable anomaly logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Anomaly name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Quarantine method. Valid values: `none`, `attacker`.
        /// </summary>
        [Input("quarantine")]
        public Input<string>? Quarantine { get; set; }

        /// <summary>
        /// Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
        /// </summary>
        [Input("quarantineExpiry")]
        public Input<string>? QuarantineExpiry { get; set; }

        /// <summary>
        /// Enable/disable quarantine logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("quarantineLog")]
        public Input<string>? QuarantineLog { get; set; }

        /// <summary>
        /// Enable/disable this anomaly. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.
        /// </summary>
        [Input("threshold")]
        public Input<int>? Threshold { get; set; }

        /// <summary>
        /// Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.
        /// </summary>
        [Input("thresholddefault")]
        public Input<int>? Thresholddefault { get; set; }

        public FirewallSnifferAnomalyArgs()
        {
        }
        public static new FirewallSnifferAnomalyArgs Empty => new FirewallSnifferAnomalyArgs();
    }
}
