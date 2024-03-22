// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Web.Inputs
{

    public sealed class PortalLandingPageGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("formDatas")]
        private InputList<Inputs.PortalLandingPageFormDataGetArgs>? _formDatas;

        /// <summary>
        /// Form data. The structure of `form_data` block is documented below.
        /// </summary>
        public InputList<Inputs.PortalLandingPageFormDataGetArgs> FormDatas
        {
            get => _formDatas ?? (_formDatas = new InputList<Inputs.PortalLandingPageFormDataGetArgs>());
            set => _formDatas = value;
        }

        /// <summary>
        /// Landing page log out URL.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        /// <summary>
        /// Single sign-on. Valid values: `disable`, `static`, `auto`.
        /// </summary>
        [Input("sso")]
        public Input<string>? Sso { get; set; }

        /// <summary>
        /// Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.
        /// </summary>
        [Input("ssoCredential")]
        public Input<string>? SsoCredential { get; set; }

        /// <summary>
        /// SSO password.
        /// </summary>
        [Input("ssoPassword")]
        public Input<string>? SsoPassword { get; set; }

        /// <summary>
        /// SSO user name.
        /// </summary>
        [Input("ssoUsername")]
        public Input<string>? SsoUsername { get; set; }

        /// <summary>
        /// Landing page URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public PortalLandingPageGetArgs()
        {
        }
        public static new PortalLandingPageGetArgs Empty => new PortalLandingPageGetArgs();
    }
}
