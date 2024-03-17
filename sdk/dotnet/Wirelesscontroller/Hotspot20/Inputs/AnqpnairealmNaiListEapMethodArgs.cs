// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wirelesscontroller.Hotspot20.Inputs
{

    public sealed class AnqpnairealmNaiListEapMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("authParams")]
        private InputList<Inputs.AnqpnairealmNaiListEapMethodAuthParamArgs>? _authParams;

        /// <summary>
        /// EAP auth param. The structure of `auth_param` block is documented below.
        /// </summary>
        public InputList<Inputs.AnqpnairealmNaiListEapMethodAuthParamArgs> AuthParams
        {
            get => _authParams ?? (_authParams = new InputList<Inputs.AnqpnairealmNaiListEapMethodAuthParamArgs>());
            set => _authParams = value;
        }

        /// <summary>
        /// EAP method index.
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// EAP method type. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        public AnqpnairealmNaiListEapMethodArgs()
        {
        }
        public static new AnqpnairealmNaiListEapMethodArgs Empty => new AnqpnairealmNaiListEapMethodArgs();
    }
}
