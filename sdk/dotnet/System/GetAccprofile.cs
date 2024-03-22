// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    public static class GetAccprofile
    {
        /// <summary>
        /// Use this data source to get information on an fortios system accprofile
        /// </summary>
        public static Task<GetAccprofileResult> InvokeAsync(GetAccprofileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccprofileResult>("fortios:system/getAccprofile:getAccprofile", args ?? new GetAccprofileArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system accprofile
        /// </summary>
        public static Output<GetAccprofileResult> Invoke(GetAccprofileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccprofileResult>("fortios:system/getAccprofile:getAccprofile", args ?? new GetAccprofileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccprofileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system accprofile.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAccprofileArgs()
        {
        }
        public static new GetAccprofileArgs Empty => new GetAccprofileArgs();
    }

    public sealed class GetAccprofileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system accprofile.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAccprofileInvokeArgs()
        {
        }
        public static new GetAccprofileInvokeArgs Empty => new GetAccprofileInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccprofileResult
    {
        /// <summary>
        /// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
        /// </summary>
        public readonly int Admintimeout;
        /// <summary>
        /// Enable/disable overriding the global administrator idle timeout.
        /// </summary>
        public readonly string AdmintimeoutOverride;
        /// <summary>
        /// Administrator access to Users and Devices.
        /// </summary>
        public readonly string Authgrp;
        /// <summary>
        /// Enable/disable permission to run config commands.
        /// </summary>
        public readonly string CliConfig;
        /// <summary>
        /// Enable/disable permission to run diagnostic commands.
        /// </summary>
        public readonly string CliDiagnose;
        /// <summary>
        /// Enable/disable permission to run execute commands.
        /// </summary>
        public readonly string CliExec;
        /// <summary>
        /// Enable/disable permission to run get commands.
        /// </summary>
        public readonly string CliGet;
        /// <summary>
        /// Enable/disable permission to run show commands.
        /// </summary>
        public readonly string CliShow;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// FortiView.
        /// </summary>
        public readonly string Ftviewgrp;
        /// <summary>
        /// Administrator access to the Firewall configuration.
        /// </summary>
        public readonly string Fwgrp;
        /// <summary>
        /// Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccprofileFwgrpPermissionResult> FwgrpPermissions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Administrator access to Logging and Reporting including viewing log messages.
        /// </summary>
        public readonly string Loggrp;
        /// <summary>
        /// Custom Log &amp; Report permission. The structure of `loggrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccprofileLoggrpPermissionResult> LoggrpPermissions;
        /// <summary>
        /// Profile name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network Configuration.
        /// </summary>
        public readonly string Netgrp;
        /// <summary>
        /// Custom network permission. The structure of `netgrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccprofileNetgrpPermissionResult> NetgrpPermissions;
        /// <summary>
        /// Scope of admin access: global or specific VDOM(s).
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// Security Fabric.
        /// </summary>
        public readonly string Secfabgrp;
        /// <summary>
        /// System Configuration.
        /// </summary>
        public readonly string Sysgrp;
        /// <summary>
        /// Custom system permission. The structure of `sysgrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccprofileSysgrpPermissionResult> SysgrpPermissions;
        /// <summary>
        /// Enable/disable permission to run system diagnostic commands.
        /// </summary>
        public readonly string SystemDiagnostics;
        /// <summary>
        /// Enable/disable permission to execute SSH commands.
        /// </summary>
        public readonly string SystemExecuteSsh;
        /// <summary>
        /// Enable/disable permission to execute TELNET commands.
        /// </summary>
        public readonly string SystemExecuteTelnet;
        /// <summary>
        /// Administrator access to Security Profiles.
        /// </summary>
        public readonly string Utmgrp;
        /// <summary>
        /// Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccprofileUtmgrpPermissionResult> UtmgrpPermissions;
        public readonly string? Vdomparam;
        /// <summary>
        /// Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        /// </summary>
        public readonly string Vpngrp;
        /// <summary>
        /// Administrator access to WAN Opt &amp; Cache.
        /// </summary>
        public readonly string Wanoptgrp;
        /// <summary>
        /// Administrator access to the WiFi controller and Switch controller.
        /// </summary>
        public readonly string Wifi;

        [OutputConstructor]
        private GetAccprofileResult(
            int admintimeout,

            string admintimeoutOverride,

            string authgrp,

            string cliConfig,

            string cliDiagnose,

            string cliExec,

            string cliGet,

            string cliShow,

            string comments,

            string ftviewgrp,

            string fwgrp,

            ImmutableArray<Outputs.GetAccprofileFwgrpPermissionResult> fwgrpPermissions,

            string id,

            string loggrp,

            ImmutableArray<Outputs.GetAccprofileLoggrpPermissionResult> loggrpPermissions,

            string name,

            string netgrp,

            ImmutableArray<Outputs.GetAccprofileNetgrpPermissionResult> netgrpPermissions,

            string scope,

            string secfabgrp,

            string sysgrp,

            ImmutableArray<Outputs.GetAccprofileSysgrpPermissionResult> sysgrpPermissions,

            string systemDiagnostics,

            string systemExecuteSsh,

            string systemExecuteTelnet,

            string utmgrp,

            ImmutableArray<Outputs.GetAccprofileUtmgrpPermissionResult> utmgrpPermissions,

            string? vdomparam,

            string vpngrp,

            string wanoptgrp,

            string wifi)
        {
            Admintimeout = admintimeout;
            AdmintimeoutOverride = admintimeoutOverride;
            Authgrp = authgrp;
            CliConfig = cliConfig;
            CliDiagnose = cliDiagnose;
            CliExec = cliExec;
            CliGet = cliGet;
            CliShow = cliShow;
            Comments = comments;
            Ftviewgrp = ftviewgrp;
            Fwgrp = fwgrp;
            FwgrpPermissions = fwgrpPermissions;
            Id = id;
            Loggrp = loggrp;
            LoggrpPermissions = loggrpPermissions;
            Name = name;
            Netgrp = netgrp;
            NetgrpPermissions = netgrpPermissions;
            Scope = scope;
            Secfabgrp = secfabgrp;
            Sysgrp = sysgrp;
            SysgrpPermissions = sysgrpPermissions;
            SystemDiagnostics = systemDiagnostics;
            SystemExecuteSsh = systemExecuteSsh;
            SystemExecuteTelnet = systemExecuteTelnet;
            Utmgrp = utmgrp;
            UtmgrpPermissions = utmgrpPermissions;
            Vdomparam = vdomparam;
            Vpngrp = vpngrp;
            Wanoptgrp = wanoptgrp;
            Wifi = wifi;
        }
    }
}
