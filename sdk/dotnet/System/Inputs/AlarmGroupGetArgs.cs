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

    public sealed class AlarmGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Admin authentication failure threshold.
        /// </summary>
        [Input("adminAuthFailureThreshold")]
        public Input<int>? AdminAuthFailureThreshold { get; set; }

        /// <summary>
        /// Admin authentication lockout threshold.
        /// </summary>
        [Input("adminAuthLockoutThreshold")]
        public Input<int>? AdminAuthLockoutThreshold { get; set; }

        /// <summary>
        /// Decryption failure threshold.
        /// </summary>
        [Input("decryptionFailureThreshold")]
        public Input<int>? DecryptionFailureThreshold { get; set; }

        /// <summary>
        /// Encryption failure threshold.
        /// </summary>
        [Input("encryptionFailureThreshold")]
        public Input<int>? EncryptionFailureThreshold { get; set; }

        /// <summary>
        /// Firewall policy ID.
        /// </summary>
        [Input("fwPolicyId")]
        public Input<int>? FwPolicyId { get; set; }

        /// <summary>
        /// Firewall policy ID threshold.
        /// </summary>
        [Input("fwPolicyIdThreshold")]
        public Input<int>? FwPolicyIdThreshold { get; set; }

        [Input("fwPolicyViolations")]
        private InputList<Inputs.AlarmGroupFwPolicyViolationGetArgs>? _fwPolicyViolations;

        /// <summary>
        /// Firewall policy violations. The structure of `fw_policy_violations` block is documented below.
        /// </summary>
        public InputList<Inputs.AlarmGroupFwPolicyViolationGetArgs> FwPolicyViolations
        {
            get => _fwPolicyViolations ?? (_fwPolicyViolations = new InputList<Inputs.AlarmGroupFwPolicyViolationGetArgs>());
            set => _fwPolicyViolations = value;
        }

        /// <summary>
        /// Group ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Log full warning threshold.
        /// </summary>
        [Input("logFullWarningThreshold")]
        public Input<int>? LogFullWarningThreshold { get; set; }

        /// <summary>
        /// Time period in seconds (0 = from start up).
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Replay attempt threshold.
        /// </summary>
        [Input("replayAttemptThreshold")]
        public Input<int>? ReplayAttemptThreshold { get; set; }

        /// <summary>
        /// Self-test failure threshold.
        /// </summary>
        [Input("selfTestFailureThreshold")]
        public Input<int>? SelfTestFailureThreshold { get; set; }

        /// <summary>
        /// User authentication failure threshold.
        /// </summary>
        [Input("userAuthFailureThreshold")]
        public Input<int>? UserAuthFailureThreshold { get; set; }

        /// <summary>
        /// User authentication lockout threshold.
        /// </summary>
        [Input("userAuthLockoutThreshold")]
        public Input<int>? UserAuthLockoutThreshold { get; set; }

        public AlarmGroupGetArgs()
        {
        }
        public static new AlarmGroupGetArgs Empty => new AlarmGroupGetArgs();
    }
}
