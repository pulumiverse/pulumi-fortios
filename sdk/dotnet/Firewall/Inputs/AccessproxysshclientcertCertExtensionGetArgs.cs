// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class AccessproxysshclientcertCertExtensionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Critical option. Valid values: `no`, `yes`.
        /// </summary>
        [Input("critical")]
        public Input<string>? Critical { get; set; }

        /// <summary>
        /// Data of certificate extension.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// Name of certificate extension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of certificate extension. Valid values: `fixed`, `user`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AccessproxysshclientcertCertExtensionGetArgs()
        {
        }
        public static new AccessproxysshclientcertCertExtensionGetArgs Empty => new AccessproxysshclientcertCertExtensionGetArgs();
    }
}
