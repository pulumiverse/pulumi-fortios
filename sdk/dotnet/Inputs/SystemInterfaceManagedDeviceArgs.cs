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

    public sealed class SystemInterfaceManagedDeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Managed dev identifier.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SystemInterfaceManagedDeviceArgs()
        {
        }
        public static new SystemInterfaceManagedDeviceArgs Empty => new SystemInterfaceManagedDeviceArgs();
    }
}
