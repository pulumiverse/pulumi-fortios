// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Waf.Inputs
{

    public sealed class ProfileSignatureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The minimum number of Credit cards to detect violation.
        /// </summary>
        [Input("creditCardDetectionThreshold")]
        public Input<int>? CreditCardDetectionThreshold { get; set; }

        [Input("customSignatures")]
        private InputList<Inputs.ProfileSignatureCustomSignatureArgs>? _customSignatures;

        /// <summary>
        /// Custom signature. The structure of `custom_signature` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSignatureCustomSignatureArgs> CustomSignatures
        {
            get => _customSignatures ?? (_customSignatures = new InputList<Inputs.ProfileSignatureCustomSignatureArgs>());
            set => _customSignatures = value;
        }

        [Input("disabledSignatures")]
        private InputList<Inputs.ProfileSignatureDisabledSignatureArgs>? _disabledSignatures;

        /// <summary>
        /// Disabled signatures The structure of `disabled_signature` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSignatureDisabledSignatureArgs> DisabledSignatures
        {
            get => _disabledSignatures ?? (_disabledSignatures = new InputList<Inputs.ProfileSignatureDisabledSignatureArgs>());
            set => _disabledSignatures = value;
        }

        [Input("disabledSubClasses")]
        private InputList<Inputs.ProfileSignatureDisabledSubClassArgs>? _disabledSubClasses;

        /// <summary>
        /// Disabled signature subclasses. The structure of `disabled_sub_class` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSignatureDisabledSubClassArgs> DisabledSubClasses
        {
            get => _disabledSubClasses ?? (_disabledSubClasses = new InputList<Inputs.ProfileSignatureDisabledSubClassArgs>());
            set => _disabledSubClasses = value;
        }

        [Input("mainClasses")]
        private InputList<Inputs.ProfileSignatureMainClassArgs>? _mainClasses;

        /// <summary>
        /// Main signature class. The structure of `main_class` block is documented below.
        /// </summary>
        public InputList<Inputs.ProfileSignatureMainClassArgs> MainClasses
        {
            get => _mainClasses ?? (_mainClasses = new InputList<Inputs.ProfileSignatureMainClassArgs>());
            set => _mainClasses = value;
        }

        public ProfileSignatureArgs()
        {
        }
        public static new ProfileSignatureArgs Empty => new ProfileSignatureArgs();
    }
}