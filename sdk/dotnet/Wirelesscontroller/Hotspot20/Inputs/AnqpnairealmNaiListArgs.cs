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

    public sealed class AnqpnairealmNaiListArgs : global::Pulumi.ResourceArgs
    {
        [Input("eapMethods")]
        private InputList<Inputs.AnqpnairealmNaiListEapMethodArgs>? _eapMethods;

        /// <summary>
        /// EAP Methods. The structure of `eap_method` block is documented below.
        /// </summary>
        public InputList<Inputs.AnqpnairealmNaiListEapMethodArgs> EapMethods
        {
            get => _eapMethods ?? (_eapMethods = new InputList<Inputs.AnqpnairealmNaiListEapMethodArgs>());
            set => _eapMethods = value;
        }

        /// <summary>
        /// Enable/disable format in accordance with IETF RFC 4282. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Configure NAI realms (delimited by a semi-colon character).
        /// </summary>
        [Input("naiRealm")]
        public Input<string>? NaiRealm { get; set; }

        /// <summary>
        /// NAI realm name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AnqpnairealmNaiListArgs()
        {
        }
        public static new AnqpnairealmNaiListArgs Empty => new AnqpnairealmNaiListArgs();
    }
}
