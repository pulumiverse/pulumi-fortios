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

    public sealed class SwitchcontrollerVlanPortalMessageOverridesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Override auth-disclaimer-page message with message from portal-message-overrides group.
        /// </summary>
        [Input("authDisclaimerPage")]
        public Input<string>? AuthDisclaimerPage { get; set; }

        /// <summary>
        /// Override auth-login-failed-page message with message from portal-message-overrides group.
        /// </summary>
        [Input("authLoginFailedPage")]
        public Input<string>? AuthLoginFailedPage { get; set; }

        /// <summary>
        /// Override auth-login-page message with message from portal-message-overrides group.
        /// </summary>
        [Input("authLoginPage")]
        public Input<string>? AuthLoginPage { get; set; }

        /// <summary>
        /// Override auth-reject-page message with message from portal-message-overrides group.
        /// </summary>
        [Input("authRejectPage")]
        public Input<string>? AuthRejectPage { get; set; }

        public SwitchcontrollerVlanPortalMessageOverridesArgs()
        {
        }
        public static new SwitchcontrollerVlanPortalMessageOverridesArgs Empty => new SwitchcontrollerVlanPortalMessageOverridesArgs();
    }
}
