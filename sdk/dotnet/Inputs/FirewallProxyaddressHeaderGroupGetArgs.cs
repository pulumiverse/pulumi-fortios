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

    public sealed class FirewallProxyaddressHeaderGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Case sensitivity in pattern. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("caseSensitivity")]
        public Input<string>? CaseSensitivity { get; set; }

        /// <summary>
        /// HTTP header regular expression.
        /// </summary>
        [Input("header")]
        public Input<string>? Header { get; set; }

        /// <summary>
        /// HTTP header.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public FirewallProxyaddressHeaderGroupGetArgs()
        {
        }
        public static new FirewallProxyaddressHeaderGroupGetArgs Empty => new FirewallProxyaddressHeaderGroupGetArgs();
    }
}
