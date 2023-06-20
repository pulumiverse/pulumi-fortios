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

    public sealed class ApplicationListEntryParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("members")]
        private InputList<Inputs.ApplicationListEntryParameterMemberArgs>? _members;

        /// <summary>
        /// Parameter tuple members. The structure of `members` block is documented below.
        /// </summary>
        public InputList<Inputs.ApplicationListEntryParameterMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.ApplicationListEntryParameterMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Parameter value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApplicationListEntryParameterArgs()
        {
        }
        public static new ApplicationListEntryParameterArgs Empty => new ApplicationListEntryParameterArgs();
    }
}