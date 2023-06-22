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

    public sealed class RouterBgpVrf6LeakTargetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface which is used to leak routes to target VRF.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Route map of VRF leaking.
        /// </summary>
        [Input("routeMap")]
        public Input<string>? RouteMap { get; set; }

        /// <summary>
        /// Target VRF ID (0 - 63).
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public RouterBgpVrf6LeakTargetGetArgs()
        {
        }
        public static new RouterBgpVrf6LeakTargetGetArgs Empty => new RouterBgpVrf6LeakTargetGetArgs();
    }
}
