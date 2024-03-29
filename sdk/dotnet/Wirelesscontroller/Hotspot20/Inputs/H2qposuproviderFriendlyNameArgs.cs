// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20.Inputs
{

    public sealed class H2qposuproviderFriendlyNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// OSU provider friendly name.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// OSU provider friendly name index.
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// Language code.
        /// </summary>
        [Input("lang")]
        public Input<string>? Lang { get; set; }

        public H2qposuproviderFriendlyNameArgs()
        {
        }
        public static new H2qposuproviderFriendlyNameArgs Empty => new H2qposuproviderFriendlyNameArgs();
    }
}
