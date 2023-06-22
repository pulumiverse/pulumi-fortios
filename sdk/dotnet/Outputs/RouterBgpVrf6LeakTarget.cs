// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class RouterBgpVrf6LeakTarget
    {
        /// <summary>
        /// Interface which is used to leak routes to target VRF.
        /// </summary>
        public readonly string? Interface;
        /// <summary>
        /// Route map of VRF leaking.
        /// </summary>
        public readonly string? RouteMap;
        /// <summary>
        /// Target VRF ID (0 - 63).
        /// </summary>
        public readonly string? Vrf;

        [OutputConstructor]
        private RouterBgpVrf6LeakTarget(
            string? @interface,

            string? routeMap,

            string? vrf)
        {
            Interface = @interface;
            RouteMap = routeMap;
            Vrf = vrf;
        }
    }
}
