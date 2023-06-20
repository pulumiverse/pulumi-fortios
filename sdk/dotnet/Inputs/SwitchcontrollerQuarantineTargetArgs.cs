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

    public sealed class SwitchcontrollerQuarantineTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the quarantine MAC.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// FSW entry id for the quarantine MAC.
        /// </summary>
        [Input("entryId")]
        public Input<int>? EntryId { get; set; }

        /// <summary>
        /// Quarantine MAC.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        [Input("tags")]
        private InputList<Inputs.SwitchcontrollerQuarantineTargetTagArgs>? _tags;

        /// <summary>
        /// Tags for the quarantine MAC. The structure of `tag` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerQuarantineTargetTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.SwitchcontrollerQuarantineTargetTagArgs>());
            set => _tags = value;
        }

        public SwitchcontrollerQuarantineTargetArgs()
        {
        }
        public static new SwitchcontrollerQuarantineTargetArgs Empty => new SwitchcontrollerQuarantineTargetArgs();
    }
}