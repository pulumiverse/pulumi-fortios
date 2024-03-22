// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Virtualpatch.Inputs
{

    public sealed class ProfileExemptionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("devices")]
        private InputList<Inputs.ProfileExemptionDeviceGetArgs>? _devices;

        /// <summary>
        /// Device MAC addresses. The structure of `device` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileExemptionDeviceGetArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.ProfileExemptionDeviceGetArgs>());
            set => _devices = value;
        }

        /// <summary>
        /// IDs.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("rules")]
        private InputList<Inputs.ProfileExemptionRuleGetArgs>? _rules;

        /// <summary>
        /// Patch signature rule IDs. The structure of `rule` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileExemptionRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ProfileExemptionRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Enable/disable exemption. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ProfileExemptionGetArgs()
        {
        }
        public static new ProfileExemptionGetArgs Empty => new ProfileExemptionGetArgs();
    }
}
