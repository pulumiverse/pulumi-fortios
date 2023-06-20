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
    public sealed class SwitchcontrollerManagedswitchSwitchLog
    {
        /// <summary>
        /// Enable to configure local logging settings that override global logging settings. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? LocalOverride;
        /// <summary>
        /// Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        /// </summary>
        public readonly string? Severity;
        /// <summary>
        /// Enable/disable adding FortiSwitch logs to the FortiGate event log. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private SwitchcontrollerManagedswitchSwitchLog(
            string? localOverride,

            string? severity,

            string? status)
        {
            LocalOverride = localOverride;
            Severity = severity;
            Status = status;
        }
    }
}