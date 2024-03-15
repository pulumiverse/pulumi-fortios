// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Voip.Outputs
{

    [OutputType]
    public sealed class ProfileMsrp
    {
        /// <summary>
        /// Enable/disable logging of MSRP violations. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? LogViolations;
        /// <summary>
        /// Maximum allowable MSRP message size (1-65535).
        /// </summary>
        public readonly int? MaxMsgSize;
        /// <summary>
        /// Action for violation of max-msg-size. Valid values: `pass`, `block`, `reset`, `monitor`.
        /// </summary>
        public readonly string? MaxMsgSizeAction;
        /// <summary>
        /// Enable/disable MSRP. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ProfileMsrp(
            string? logViolations,

            int? maxMsgSize,

            string? maxMsgSizeAction,

            string? status)
        {
            LogViolations = logViolations;
            MaxMsgSize = maxMsgSize;
            MaxMsgSizeAction = maxMsgSizeAction;
            Status = status;
        }
    }
}