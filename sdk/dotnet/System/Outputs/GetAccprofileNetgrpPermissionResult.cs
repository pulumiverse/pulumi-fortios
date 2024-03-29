// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class GetAccprofileNetgrpPermissionResult
    {
        /// <summary>
        /// System Configuration.
        /// </summary>
        public readonly string Cfg;
        /// <summary>
        /// Packet Capture Configuration.
        /// </summary>
        public readonly string PacketCapture;
        /// <summary>
        /// Router Configuration.
        /// </summary>
        public readonly string RouteCfg;

        [OutputConstructor]
        private GetAccprofileNetgrpPermissionResult(
            string cfg,

            string packetCapture,

            string routeCfg)
        {
            Cfg = cfg;
            PacketCapture = packetCapture;
            RouteCfg = routeCfg;
        }
    }
}
