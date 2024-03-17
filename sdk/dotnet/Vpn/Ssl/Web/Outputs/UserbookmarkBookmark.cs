// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl.Web.Outputs
{

    [OutputType]
    public sealed class UserbookmarkBookmark
    {
        /// <summary>
        /// Additional parameters.
        /// </summary>
        public readonly string? AdditionalParams;
        /// <summary>
        /// Application type.
        /// </summary>
        public readonly string? Apptype;
        /// <summary>
        /// Color depth per pixel. Valid values: `32`, `16`, `8`.
        /// </summary>
        public readonly string? ColorDepth;
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Login domain.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// Network shared file folder parameter.
        /// </summary>
        public readonly string? Folder;
        /// <summary>
        /// Form data. The structure of `form_data` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.UserbookmarkBookmarkFormData> FormDatas;
        /// <summary>
        /// Screen height (range from 480 - 65535, default = 768).
        /// </summary>
        public readonly int? Height;
        /// <summary>
        /// Host name/IP parameter.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// Keyboard layout.
        /// </summary>
        public readonly string? KeyboardLayout;
        /// <summary>
        /// Listening port (0 - 65535).
        /// </summary>
        public readonly int? ListeningPort;
        /// <summary>
        /// The load balancing information or cookie which should be provided to the connection broker.
        /// </summary>
        public readonly string? LoadBalancingInfo;
        /// <summary>
        /// Logon password.
        /// </summary>
        public readonly string? LogonPassword;
        /// <summary>
        /// Logon user.
        /// </summary>
        public readonly string? LogonUser;
        /// <summary>
        /// Bookmark name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Remote port.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// An arbitrary string which identifies the RDP source.
        /// </summary>
        public readonly string? PreconnectionBlob;
        /// <summary>
        /// The numeric ID of the RDP source (0-2147483648).
        /// </summary>
        public readonly int? PreconnectionId;
        /// <summary>
        /// Remote port (0 - 65535).
        /// </summary>
        public readonly int? RemotePort;
        /// <summary>
        /// Enable/disable restricted admin mode for RDP. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? RestrictedAdmin;
        /// <summary>
        /// Security mode for RDP connection. Valid values: `rdp`, `nla`, `tls`, `any`.
        /// </summary>
        public readonly string? Security;
        /// <summary>
        /// Enable/disable sending of preconnection ID. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SendPreconnectionId;
        /// <summary>
        /// Server side keyboard layout.
        /// </summary>
        public readonly string? ServerLayout;
        /// <summary>
        /// Enable/disable showing of status window. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ShowStatusWindow;
        /// <summary>
        /// Single Sign-On. Valid values: `disable`, `static`, `auto`.
        /// </summary>
        public readonly string? Sso;
        /// <summary>
        /// Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.
        /// </summary>
        public readonly string? SsoCredential;
        /// <summary>
        /// Single sign-on credentials are only sent once to remote server. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? SsoCredentialSentOnce;
        /// <summary>
        /// SSO password.
        /// </summary>
        public readonly string? SsoPassword;
        /// <summary>
        /// SSO user name.
        /// </summary>
        public readonly string? SsoUsername;
        /// <summary>
        /// URL parameter.
        /// </summary>
        public readonly string? Url;
        /// <summary>
        /// Screen width (range from 640 - 65535, default = 1024).
        /// </summary>
        public readonly int? Width;

        [OutputConstructor]
        private UserbookmarkBookmark(
            string? additionalParams,

            string? apptype,

            string? colorDepth,

            string? description,

            string? domain,

            string? folder,

            ImmutableArray<Outputs.UserbookmarkBookmarkFormData> formDatas,

            int? height,

            string? host,

            string? keyboardLayout,

            int? listeningPort,

            string? loadBalancingInfo,

            string? logonPassword,

            string? logonUser,

            string? name,

            int? port,

            string? preconnectionBlob,

            int? preconnectionId,

            int? remotePort,

            string? restrictedAdmin,

            string? security,

            string? sendPreconnectionId,

            string? serverLayout,

            string? showStatusWindow,

            string? sso,

            string? ssoCredential,

            string? ssoCredentialSentOnce,

            string? ssoPassword,

            string? ssoUsername,

            string? url,

            int? width)
        {
            AdditionalParams = additionalParams;
            Apptype = apptype;
            ColorDepth = colorDepth;
            Description = description;
            Domain = domain;
            Folder = folder;
            FormDatas = formDatas;
            Height = height;
            Host = host;
            KeyboardLayout = keyboardLayout;
            ListeningPort = listeningPort;
            LoadBalancingInfo = loadBalancingInfo;
            LogonPassword = logonPassword;
            LogonUser = logonUser;
            Name = name;
            Port = port;
            PreconnectionBlob = preconnectionBlob;
            PreconnectionId = preconnectionId;
            RemotePort = remotePort;
            RestrictedAdmin = restrictedAdmin;
            Security = security;
            SendPreconnectionId = sendPreconnectionId;
            ServerLayout = serverLayout;
            ShowStatusWindow = showStatusWindow;
            Sso = sso;
            SsoCredential = ssoCredential;
            SsoCredentialSentOnce = ssoCredentialSentOnce;
            SsoPassword = ssoPassword;
            SsoUsername = ssoUsername;
            Url = url;
            Width = width;
        }
    }
}
