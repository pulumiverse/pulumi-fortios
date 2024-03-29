// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extensioncontroller.Outputs
{

    [OutputType]
    public sealed class ExtenderprofileCellularSmsNotification
    {
        /// <summary>
        /// SMS alert list. The structure of `alert` block is documented below.
        /// </summary>
        public readonly Outputs.ExtenderprofileCellularSmsNotificationAlert? Alert;
        /// <summary>
        /// SMS notification receiver list. The structure of `receiver` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExtenderprofileCellularSmsNotificationReceiver> Receivers;
        /// <summary>
        /// FortiExtender SMS notification status. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ExtenderprofileCellularSmsNotification(
            Outputs.ExtenderprofileCellularSmsNotificationAlert? alert,

            ImmutableArray<Outputs.ExtenderprofileCellularSmsNotificationReceiver> receivers,

            string? status)
        {
            Alert = alert;
            Receivers = receivers;
            Status = status;
        }
    }
}
