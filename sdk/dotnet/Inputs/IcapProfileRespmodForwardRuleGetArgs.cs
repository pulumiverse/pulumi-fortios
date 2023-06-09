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

    public sealed class IcapProfileRespmodForwardRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be taken for ICAP server. Valid values: `forward`, `bypass`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("headerGroups")]
        private InputList<Inputs.IcapProfileRespmodForwardRuleHeaderGroupGetArgs>? _headerGroups;

        /// <summary>
        /// HTTP header group. The structure of `header_group` block is documented below.
        /// </summary>
        public InputList<Inputs.IcapProfileRespmodForwardRuleHeaderGroupGetArgs> HeaderGroups
        {
            get => _headerGroups ?? (_headerGroups = new InputList<Inputs.IcapProfileRespmodForwardRuleHeaderGroupGetArgs>());
            set => _headerGroups = value;
        }

        /// <summary>
        /// Address object for the host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("httpRespStatusCodes")]
        private InputList<Inputs.IcapProfileRespmodForwardRuleHttpRespStatusCodeGetArgs>? _httpRespStatusCodes;

        /// <summary>
        /// HTTP response status code. The structure of `http_resp_status_code` block is documented below.
        /// </summary>
        public InputList<Inputs.IcapProfileRespmodForwardRuleHttpRespStatusCodeGetArgs> HttpRespStatusCodes
        {
            get => _httpRespStatusCodes ?? (_httpRespStatusCodes = new InputList<Inputs.IcapProfileRespmodForwardRuleHttpRespStatusCodeGetArgs>());
            set => _httpRespStatusCodes = value;
        }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IcapProfileRespmodForwardRuleGetArgs()
        {
        }
        public static new IcapProfileRespmodForwardRuleGetArgs Empty => new IcapProfileRespmodForwardRuleGetArgs();
    }
}
