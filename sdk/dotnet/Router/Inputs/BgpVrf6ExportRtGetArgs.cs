// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class BgpVrf6ExportRtGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Attribute: AA:NN|A.B.C.D:NN.
        /// </summary>
        [Input("routeTarget")]
        public Input<string>? RouteTarget { get; set; }

        public BgpVrf6ExportRtGetArgs()
        {
        }
        public static new BgpVrf6ExportRtGetArgs Empty => new BgpVrf6ExportRtGetArgs();
    }
}
