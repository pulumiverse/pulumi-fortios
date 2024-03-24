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

    public sealed class ProfileSaasApplicationCustomControlOptionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CASB custom control option name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("userInputs")]
        private InputList<Inputs.ProfileSaasApplicationCustomControlOptionUserInputGetArgs>? _userInputs;

        /// <summary>
        /// CASB custom control user input. The structure of `user_input` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSaasApplicationCustomControlOptionUserInputGetArgs> UserInputs
        {
            get => _userInputs ?? (_userInputs = new InputList<Inputs.ProfileSaasApplicationCustomControlOptionUserInputGetArgs>());
            set => _userInputs = value;
        }

        public ProfileSaasApplicationCustomControlOptionGetArgs()
        {
        }
        public static new ProfileSaasApplicationCustomControlOptionGetArgs Empty => new ProfileSaasApplicationCustomControlOptionGetArgs();
    }
}