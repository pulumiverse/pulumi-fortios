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

    public sealed class QosprofileDscpWmmBkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DSCP WMM mapping numbers (0 - 63).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public QosprofileDscpWmmBkArgs()
        {
        }
        public static new QosprofileDscpWmmBkArgs Empty => new QosprofileDscpWmmBkArgs();
    }
}
