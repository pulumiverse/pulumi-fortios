// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Outputs
{

    [OutputType]
    public sealed class ProfileprotocoloptionsMailSignature
    {
        /// <summary>
        /// Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).
        /// </summary>
        public readonly string? Signature;
        /// <summary>
        /// Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ProfileprotocoloptionsMailSignature(
            string? signature,

            string? status)
        {
            Signature = signature;
            Status = status;
        }
    }
}
