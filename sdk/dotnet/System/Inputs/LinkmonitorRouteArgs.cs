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

    public sealed class LinkmonitorRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP and netmask (x.x.x.x/y).
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        public LinkmonitorRouteArgs()
        {
        }
        public static new LinkmonitorRouteArgs Empty => new LinkmonitorRouteArgs();
    }
}
