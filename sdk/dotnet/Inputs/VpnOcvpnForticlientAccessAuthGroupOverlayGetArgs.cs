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

    public sealed class VpnOcvpnForticlientAccessAuthGroupOverlayGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Overlay name.
        /// </summary>
        [Input("overlayName")]
        public Input<string>? OverlayName { get; set; }

        public VpnOcvpnForticlientAccessAuthGroupOverlayGetArgs()
        {
        }
        public static new VpnOcvpnForticlientAccessAuthGroupOverlayGetArgs Empty => new VpnOcvpnForticlientAccessAuthGroupOverlayGetArgs();
    }
}
