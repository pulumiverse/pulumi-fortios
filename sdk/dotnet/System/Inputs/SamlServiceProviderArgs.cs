// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Inputs
{

    public sealed class SamlServiceProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("assertionAttributes")]
        private InputList<Inputs.SamlServiceProviderAssertionAttributeArgs>? _assertionAttributes;

        /// <summary>
        /// Customized SAML attributes to send along with assertion. The structure of `assertion_attributes` block is documented below.
        /// </summary>
        public InputList<Inputs.SamlServiceProviderAssertionAttributeArgs> AssertionAttributes
        {
            get => _assertionAttributes ?? (_assertionAttributes = new InputList<Inputs.SamlServiceProviderAssertionAttributeArgs>());
            set => _assertionAttributes = value;
        }

        /// <summary>
        /// IDP entity ID.
        /// </summary>
        [Input("idpEntityId")]
        public Input<string>? IdpEntityId { get; set; }

        /// <summary>
        /// IDP single logout URL.
        /// </summary>
        [Input("idpSingleLogoutUrl")]
        public Input<string>? IdpSingleLogoutUrl { get; set; }

        /// <summary>
        /// IDP single sign-on URL.
        /// </summary>
        [Input("idpSingleSignOnUrl")]
        public Input<string>? IdpSingleSignOnUrl { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Prefix.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// SP binding protocol. Valid values: `post`, `redirect`.
        /// </summary>
        [Input("spBindingProtocol")]
        public Input<string>? SpBindingProtocol { get; set; }

        /// <summary>
        /// SP certificate name.
        /// </summary>
        [Input("spCert")]
        public Input<string>? SpCert { get; set; }

        /// <summary>
        /// SP entity ID.
        /// </summary>
        [Input("spEntityId")]
        public Input<string>? SpEntityId { get; set; }

        /// <summary>
        /// SP portal URL.
        /// </summary>
        [Input("spPortalUrl")]
        public Input<string>? SpPortalUrl { get; set; }

        /// <summary>
        /// SP single logout URL.
        /// </summary>
        [Input("spSingleLogoutUrl")]
        public Input<string>? SpSingleLogoutUrl { get; set; }

        /// <summary>
        /// SP single sign-on URL.
        /// </summary>
        [Input("spSingleSignOnUrl")]
        public Input<string>? SpSingleSignOnUrl { get; set; }

        public SamlServiceProviderArgs()
        {
        }
        public static new SamlServiceProviderArgs Empty => new SamlServiceProviderArgs();
    }
}
