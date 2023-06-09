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
    public sealed class IcapProfileRespmodForwardRule
    {
        /// <summary>
        /// Action to be taken for ICAP server. Valid values: `forward`, `bypass`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// HTTP header group. The structure of `header_group` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.IcapProfileRespmodForwardRuleHeaderGroup> HeaderGroups;
        /// <summary>
        /// Address object for the host.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// HTTP response status code. The structure of `http_resp_status_code` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.IcapProfileRespmodForwardRuleHttpRespStatusCode> HttpRespStatusCodes;
        /// <summary>
        /// Address name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private IcapProfileRespmodForwardRule(
            string? action,

            ImmutableArray<Outputs.IcapProfileRespmodForwardRuleHeaderGroup> headerGroups,

            string? host,

            ImmutableArray<Outputs.IcapProfileRespmodForwardRuleHttpRespStatusCode> httpRespStatusCodes,

            string? name)
        {
            Action = action;
            HeaderGroups = headerGroups;
            Host = host;
            HttpRespStatusCodes = httpRespStatusCodes;
            Name = name;
        }
    }
}
