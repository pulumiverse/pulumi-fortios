// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class EvpnExportRtGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Route target: AA|AA:NN.
        /// </summary>
        [Input("routeTarget")]
        public Input<string>? RouteTarget { get; set; }

        public EvpnExportRtGetArgs()
        {
        }
        public static new EvpnExportRtGetArgs Empty => new EvpnExportRtGetArgs();
    }
}
