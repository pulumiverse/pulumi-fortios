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

    public sealed class PortalBookmarkGroupBookmarkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Additional parameters.
        /// </summary>
        [Input("additionalParams")]
        public Input<string>? AdditionalParams { get; set; }

        /// <summary>
        /// Application type.
        /// </summary>
        [Input("apptype")]
        public Input<string>? Apptype { get; set; }

        /// <summary>
        /// Color depth per pixel. Valid values: `32`, `16`, `8`.
        /// </summary>
        [Input("colorDepth")]
        public Input<string>? ColorDepth { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Login domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Network shared file folder parameter.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        [Input("formDatas")]
        private InputList<Inputs.PortalBookmarkGroupBookmarkFormDataGetArgs>? _formDatas;

        /// <summary>
        /// Form data. The structure of `form_data` block is documented below.
        /// </summary>
        public InputList<Inputs.PortalBookmarkGroupBookmarkFormDataGetArgs> FormDatas
        {
            get => _formDatas ?? (_formDatas = new InputList<Inputs.PortalBookmarkGroupBookmarkFormDataGetArgs>());
            set => _formDatas = value;
        }

        /// <summary>
        /// Screen height. On FortiOS versions 7.0.4-7.0.5: range from 480 - 65535, default = 768. On FortiOS versions &gt;= 7.0.6: range from 0 - 65535, default = 0.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// Host name/IP parameter.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Keyboard layout.
        /// </summary>
        [Input("keyboardLayout")]
        public Input<string>? KeyboardLayout { get; set; }

        /// <summary>
        /// Listening port (0 - 65535).
        /// </summary>
        [Input("listeningPort")]
        public Input<int>? ListeningPort { get; set; }

        /// <summary>
        /// The load balancing information or cookie which should be provided to the connection broker.
        /// </summary>
        [Input("loadBalancingInfo")]
        public Input<string>? LoadBalancingInfo { get; set; }

        [Input("logonPassword")]
        private Input<string>? _logonPassword;

        /// <summary>
        /// Logon password.
        /// </summary>
        public Input<string>? LogonPassword
        {
            get => _logonPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _logonPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Logon user.
        /// </summary>
        [Input("logonUser")]
        public Input<string>? LogonUser { get; set; }

        /// <summary>
        /// Bookmark name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Remote port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// An arbitrary string which identifies the RDP source.
        /// </summary>
        [Input("preconnectionBlob")]
        public Input<string>? PreconnectionBlob { get; set; }

        /// <summary>
        /// The numeric ID of the RDP source. On FortiOS versions 6.2.0-6.4.2, 7.0.0: 0-2147483648. On FortiOS versions 6.4.10-6.4.15, &gt;= 7.0.1: 0-4294967295.
        /// </summary>
        [Input("preconnectionId")]
        public Input<int>? PreconnectionId { get; set; }

        /// <summary>
        /// Remote port (0 - 65535).
        /// </summary>
        [Input("remotePort")]
        public Input<int>? RemotePort { get; set; }

        /// <summary>
        /// Enable/disable restricted admin mode for RDP. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("restrictedAdmin")]
        public Input<string>? RestrictedAdmin { get; set; }

        /// <summary>
        /// Security mode for RDP connection (default = any). Valid values: `rdp`, `nla`, `tls`, `any`.
        /// </summary>
        [Input("security")]
        public Input<string>? Security { get; set; }

        /// <summary>
        /// Enable/disable sending of preconnection ID. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("sendPreconnectionId")]
        public Input<string>? SendPreconnectionId { get; set; }

        /// <summary>
        /// Server side keyboard layout.
        /// </summary>
        [Input("serverLayout")]
        public Input<string>? ServerLayout { get; set; }

        /// <summary>
        /// Enable/disable showing of status window. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("showStatusWindow")]
        public Input<string>? ShowStatusWindow { get; set; }

        /// <summary>
        /// Single Sign-On. Valid values: `disable`, `static`, `auto`.
        /// </summary>
        [Input("sso")]
        public Input<string>? Sso { get; set; }

        /// <summary>
        /// Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.
        /// </summary>
        [Input("ssoCredential")]
        public Input<string>? SsoCredential { get; set; }

        /// <summary>
        /// Single sign-on credentials are only sent once to remote server. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ssoCredentialSentOnce")]
        public Input<string>? SsoCredentialSentOnce { get; set; }

        [Input("ssoPassword")]
        private Input<string>? _ssoPassword;

        /// <summary>
        /// SSO password.
        /// </summary>
        public Input<string>? SsoPassword
        {
            get => _ssoPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ssoPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// SSO user name.
        /// </summary>
        [Input("ssoUsername")]
        public Input<string>? SsoUsername { get; set; }

        /// <summary>
        /// URL parameter.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Keyboard layout. Valid values: `default`, `da`, `nl`, `en-uk`, `en-uk-ext`, `fi`, `fr`, `fr-be`, `fr-ca-mul`, `de`, `de-ch`, `it`, `it-142`, `pt`, `pt-br-abnt2`, `no`, `gd`, `es`, `sv`, `us-intl`.
        /// </summary>
        [Input("vncKeyboardLayout")]
        public Input<string>? VncKeyboardLayout { get; set; }

        /// <summary>
        /// Screen width. On FortiOS versions 7.0.4-7.0.5: range from 640 - 65535, default = 1024. On FortiOS versions &gt;= 7.0.6: range from 0 - 65535, default = 0.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public PortalBookmarkGroupBookmarkGetArgs()
        {
        }
        public static new PortalBookmarkGroupBookmarkGetArgs Empty => new PortalBookmarkGroupBookmarkGetArgs();
    }
}
