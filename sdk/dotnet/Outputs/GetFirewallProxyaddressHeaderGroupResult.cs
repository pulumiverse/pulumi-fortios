// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class GetFirewallProxyaddressHeaderGroupResult
    {
        /// <summary>
        /// Case sensitivity in pattern.
        /// </summary>
        public readonly string CaseSensitivity;
        /// <summary>
        /// HTTP header regular expression.
        /// </summary>
        public readonly string Header;
        /// <summary>
        /// HTTP header.
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int Id;

        [OutputConstructor]
        private GetFirewallProxyaddressHeaderGroupResult(
            string caseSensitivity,

            string header,

            string headerName,

            int id)
        {
            CaseSensitivity = caseSensitivity;
            Header = header;
            HeaderName = headerName;
            Id = id;
        }
    }
}
