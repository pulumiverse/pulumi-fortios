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
    public sealed class ExtensioncontrollerExtenderprofileCellularSmsNotification
    {
        /// <summary>
        /// SMS alert list. The structure of `alert` block is documented below.
        /// </summary>
        public readonly Outputs.ExtensioncontrollerExtenderprofileCellularSmsNotificationAlert? Alert;
        /// <summary>
        /// SMS notification receiver list. The structure of `receiver` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExtensioncontrollerExtenderprofileCellularSmsNotificationReceiver> Receivers;
        /// <summary>
        /// FortiExtender SMS notification status. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ExtensioncontrollerExtenderprofileCellularSmsNotification(
            Outputs.ExtensioncontrollerExtenderprofileCellularSmsNotificationAlert? alert,

            ImmutableArray<Outputs.ExtensioncontrollerExtenderprofileCellularSmsNotificationReceiver> receivers,

            string? status)
        {
            Alert = alert;
            Receivers = receivers;
            Status = status;
        }
    }
}