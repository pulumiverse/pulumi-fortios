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

    public sealed class Wirelesscontrollerhotspot20H2qpoperatornameValueListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value index.
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// Language code.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        /// <summary>
        /// Friendly name value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public Wirelesscontrollerhotspot20H2qpoperatornameValueListGetArgs()
        {
        }
        public static new Wirelesscontrollerhotspot20H2qpoperatornameValueListGetArgs Empty => new Wirelesscontrollerhotspot20H2qpoperatornameValueListGetArgs();
    }
}
