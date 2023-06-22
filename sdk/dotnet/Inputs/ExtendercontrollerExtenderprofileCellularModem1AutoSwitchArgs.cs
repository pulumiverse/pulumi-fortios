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

    public sealed class ExtendercontrollerExtenderprofileCellularModem1AutoSwitchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Automatically switch based on data usage. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("dataplan")]
        public Input<string>? Dataplan { get; set; }

        /// <summary>
        /// Auto switch by disconnect. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("disconnect")]
        public Input<string>? Disconnect { get; set; }

        /// <summary>
        /// Automatically switch based on disconnect period.
        /// </summary>
        [Input("disconnectPeriod")]
        public Input<int>? DisconnectPeriod { get; set; }

        /// <summary>
        /// Automatically switch based on disconnect threshold.
        /// </summary>
        [Input("disconnectThreshold")]
        public Input<int>? DisconnectThreshold { get; set; }

        /// <summary>
        /// Automatically switch based on signal strength. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("signal")]
        public Input<string>? Signal { get; set; }

        /// <summary>
        /// Auto switch with switch back multi-options. Valid values: `time`, `timer`.
        /// </summary>
        [Input("switchBack")]
        public Input<string>? SwitchBack { get; set; }

        /// <summary>
        /// Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
        /// </summary>
        [Input("switchBackTime")]
        public Input<string>? SwitchBackTime { get; set; }

        /// <summary>
        /// Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).
        /// </summary>
        [Input("switchBackTimer")]
        public Input<int>? SwitchBackTimer { get; set; }

        public ExtendercontrollerExtenderprofileCellularModem1AutoSwitchArgs()
        {
        }
        public static new ExtendercontrollerExtenderprofileCellularModem1AutoSwitchArgs Empty => new ExtendercontrollerExtenderprofileCellularModem1AutoSwitchArgs();
    }
}
