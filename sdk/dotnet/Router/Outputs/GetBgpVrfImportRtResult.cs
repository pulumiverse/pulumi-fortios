// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class GetBgpVrfImportRtResult
    {
        /// <summary>
        /// Attribute: AA:NN|A.B.C.D:NN
        /// </summary>
        public readonly string RouteTarget;

        [OutputConstructor]
        private GetBgpVrfImportRtResult(string routeTarget)
        {
            RouteTarget = routeTarget;
        }
    }
}
