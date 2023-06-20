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

    public sealed class SwitchcontrollerLldpprofileMedLocationServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location service type name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable or disable this TLV. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Location service ID.
        /// </summary>
        [Input("sysLocationId")]
        public Input<string>? SysLocationId { get; set; }

        public SwitchcontrollerLldpprofileMedLocationServiceArgs()
        {
        }
        public static new SwitchcontrollerLldpprofileMedLocationServiceArgs Empty => new SwitchcontrollerLldpprofileMedLocationServiceArgs();
    }
}