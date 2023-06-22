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
    public sealed class Wirelesscontrollerhotspot20AnqpnairealmNaiListEapMethod
    {
        /// <summary>
        /// EAP auth param. The structure of `auth_param` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.Wirelesscontrollerhotspot20AnqpnairealmNaiListEapMethodAuthParam> AuthParams;
        /// <summary>
        /// EAP method index.
        /// </summary>
        public readonly int? Index;
        /// <summary>
        /// EAP method type. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`.
        /// </summary>
        public readonly string? Method;

        [OutputConstructor]
        private Wirelesscontrollerhotspot20AnqpnairealmNaiListEapMethod(
            ImmutableArray<Outputs.Wirelesscontrollerhotspot20AnqpnairealmNaiListEapMethodAuthParam> authParams,

            int? index,

            string? method)
        {
            AuthParams = authParams;
            Index = index;
            Method = method;
        }
    }
}
