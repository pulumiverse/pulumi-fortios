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

    public sealed class UseractivityControlOptionOperationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CASB operation action. Valid values: `append`, `prepend`, `replace`, `new`, `new-on-not-found`, `delete`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// CASB operation search case sensitive. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("caseSensitive")]
        public Input<string>? CaseSensitive { get; set; }

        /// <summary>
        /// CASB operation direction. Valid values: `request`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// CASB operation header name to search.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// CASB control option operation name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// CASB operation key to search.
        /// </summary>
        [Input("searchKey")]
        public Input<string>? SearchKey { get; set; }

        /// <summary>
        /// CASB operation search pattern. Valid values: `simple`, `substr`, `regexp`.
        /// </summary>
        [Input("searchPattern")]
        public Input<string>? SearchPattern { get; set; }

        /// <summary>
        /// CASB operation target. Valid values: `header`, `path`.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// Enable/disable value from user input. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("valueFromInput")]
        public Input<string>? ValueFromInput { get; set; }

        [Input("values")]
        private InputList<Inputs.UseractivityControlOptionOperationValueGetArgs>? _values;

        /// <summary>
        /// CASB operation new values. The structure of `values` block is documented below.
        /// </summary>
        public InputList<Inputs.UseractivityControlOptionOperationValueGetArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.UseractivityControlOptionOperationValueGetArgs>());
            set => _values = value;
        }

        public UseractivityControlOptionOperationGetArgs()
        {
        }
        public static new UseractivityControlOptionOperationGetArgs Empty => new UseractivityControlOptionOperationGetArgs();
    }
}
