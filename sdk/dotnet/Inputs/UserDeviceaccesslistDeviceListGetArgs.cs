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

    public sealed class UserDeviceaccesslistDeviceListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow or block device. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Firewall device or device group.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public UserDeviceaccesslistDeviceListGetArgs()
        {
        }
        public static new UserDeviceaccesslistDeviceListGetArgs Empty => new UserDeviceaccesslistDeviceListGetArgs();
    }
}