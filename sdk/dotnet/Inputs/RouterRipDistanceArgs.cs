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

    public sealed class RouterRipDistanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access list for route destination.
        /// </summary>
        [Input("accessList")]
        public Input<string>? AccessList { get; set; }

        /// <summary>
        /// Distance (1 - 255).
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// Distance ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Distance prefix.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public RouterRipDistanceArgs()
        {
        }
        public static new RouterRipDistanceArgs Empty => new RouterRipDistanceArgs();
    }
}
