// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Casb.Inputs
{

    public sealed class UseractivityControlOptionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CASB control option name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operations")]
        private InputList<Inputs.UseractivityControlOptionOperationGetArgs>? _operations;

        /// <summary>
        /// CASB control option operations. The structure of `operations` block is documented below.
        /// </summary>
        public InputList<Inputs.UseractivityControlOptionOperationGetArgs> Operations
        {
            get => _operations ?? (_operations = new InputList<Inputs.UseractivityControlOptionOperationGetArgs>());
            set => _operations = value;
        }

        /// <summary>
        /// CASB control option status. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public UseractivityControlOptionGetArgs()
        {
        }
        public static new UseractivityControlOptionGetArgs Empty => new UseractivityControlOptionGetArgs();
    }
}
