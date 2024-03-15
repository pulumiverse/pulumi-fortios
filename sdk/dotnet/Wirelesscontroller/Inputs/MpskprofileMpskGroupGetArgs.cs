// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Inputs
{

    public sealed class MpskprofileMpskGroupGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("mpskKeys")]
        private InputList<Inputs.MpskprofileMpskGroupMpskKeyGetArgs>? _mpskKeys;

        /// <summary>
        /// List of multiple PSK entries. The structure of `mpsk_key` block is documented below.
        /// </summary>
        public InputList<Inputs.MpskprofileMpskGroupMpskKeyGetArgs> MpskKeys
        {
            get => _mpskKeys ?? (_mpskKeys = new InputList<Inputs.MpskprofileMpskGroupMpskKeyGetArgs>());
            set => _mpskKeys = value;
        }

        /// <summary>
        /// MPSK group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional VLAN ID.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        /// <summary>
        /// MPSK group VLAN options. Valid values: `no-vlan`, `fixed-vlan`.
        /// </summary>
        [Input("vlanType")]
        public Input<string>? VlanType { get; set; }

        public MpskprofileMpskGroupGetArgs()
        {
        }
        public static new MpskprofileMpskGroupGetArgs Empty => new MpskprofileMpskGroupGetArgs();
    }
}