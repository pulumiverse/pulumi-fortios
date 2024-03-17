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
    /// <summary>
    /// Configure access profiles for system administrators.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test12 = new Fortios.System.Accprofile("test12", new()
    ///     {
    ///         Admintimeout = 10,
    ///         AdmintimeoutOverride = "disable",
    ///         Authgrp = "read-write",
    ///         Ftviewgrp = "read-write",
    ///         Fwgrp = "custom",
    ///         FwgrpPermission = new Fortios.System.Inputs.AccprofileFwgrpPermissionArgs
    ///         {
    ///             Address = "read-write",
    ///             Policy = "read-write",
    ///             Schedule = "none",
    ///             Service = "none",
    ///         },
    ///         Loggrp = "read-write",
    ///         LoggrpPermission = new Fortios.System.Inputs.AccprofileLoggrpPermissionArgs
    ///         {
    ///             Config = "none",
    ///             DataAccess = "none",
    ///             ReportAccess = "none",
    ///             ThreatWeight = "none",
    ///         },
    ///         Netgrp = "read-write",
    ///         NetgrpPermission = new Fortios.System.Inputs.AccprofileNetgrpPermissionArgs
    ///         {
    ///             Cfg = "none",
    ///             PacketCapture = "none",
    ///             RouteCfg = "none",
    ///         },
    ///         Scope = "vdom",
    ///         Secfabgrp = "read-write",
    ///         Sysgrp = "read-write",
    ///         SysgrpPermission = new Fortios.System.Inputs.AccprofileSysgrpPermissionArgs
    ///         {
    ///             Admin = "none",
    ///             Cfg = "none",
    ///             Mnt = "none",
    ///             Upd = "none",
    ///         },
    ///         Utmgrp = "custom",
    ///         UtmgrpPermission = new Fortios.System.Inputs.AccprofileUtmgrpPermissionArgs
    ///         {
    ///             Antivirus = "read-write",
    ///             ApplicationControl = "none",
    ///             DataLossPrevention = "none",
    ///             Dnsfilter = "none",
    ///             EndpointControl = "none",
    ///             Icap = "none",
    ///             Ips = "read-write",
    ///             Voip = "none",
    ///             Waf = "none",
    ///             Webfilter = "none",
    ///         },
    ///         Vpngrp = "read-write",
    ///         Wanoptgrp = "read-write",
    ///         Wifi = "read-write",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// System Accprofile can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/accprofile:Accprofile labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:system/accprofile:Accprofile labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:system/accprofile:Accprofile")]
    public partial class Accprofile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
        /// </summary>
        [Output("admintimeout")]
        public Output<int> Admintimeout { get; private set; } = null!;

        /// <summary>
        /// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("admintimeoutOverride")]
        public Output<string> AdmintimeoutOverride { get; private set; } = null!;

        /// <summary>
        /// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Output("authgrp")]
        public Output<string> Authgrp { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comments")]
        public Output<string?> Comments { get; private set; } = null!;

        /// <summary>
        /// FortiView. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Output("ftviewgrp")]
        public Output<string> Ftviewgrp { get; private set; } = null!;

        /// <summary>
        /// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Output("fwgrp")]
        public Output<string> Fwgrp { get; private set; } = null!;

        /// <summary>
        /// Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
        /// </summary>
        [Output("fwgrpPermission")]
        public Output<Outputs.AccprofileFwgrpPermission> FwgrpPermission { get; private set; } = null!;

        /// <summary>
        /// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Output("loggrp")]
        public Output<string> Loggrp { get; private set; } = null!;

        /// <summary>
        /// Custom Log &amp; Report permission. The structure of `loggrp_permission` block is documented below.
        /// </summary>
        [Output("loggrpPermission")]
        public Output<Outputs.AccprofileLoggrpPermission> LoggrpPermission { get; private set; } = null!;

        /// <summary>
        /// Profile name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Output("netgrp")]
        public Output<string> Netgrp { get; private set; } = null!;

        /// <summary>
        /// Custom network permission. The structure of `netgrp_permission` block is documented below.
        /// </summary>
        [Output("netgrpPermission")]
        public Output<Outputs.AccprofileNetgrpPermission> NetgrpPermission { get; private set; } = null!;

        /// <summary>
        /// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// Security Fabric. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Output("secfabgrp")]
        public Output<string> Secfabgrp { get; private set; } = null!;

        /// <summary>
        /// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Output("sysgrp")]
        public Output<string> Sysgrp { get; private set; } = null!;

        /// <summary>
        /// Custom system permission. The structure of `sysgrp_permission` block is documented below.
        /// </summary>
        [Output("sysgrpPermission")]
        public Output<Outputs.AccprofileSysgrpPermission> SysgrpPermission { get; private set; } = null!;

        /// <summary>
        /// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("systemDiagnostics")]
        public Output<string> SystemDiagnostics { get; private set; } = null!;

        /// <summary>
        /// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("systemExecuteSsh")]
        public Output<string> SystemExecuteSsh { get; private set; } = null!;

        /// <summary>
        /// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("systemExecuteTelnet")]
        public Output<string> SystemExecuteTelnet { get; private set; } = null!;

        /// <summary>
        /// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Output("utmgrp")]
        public Output<string> Utmgrp { get; private set; } = null!;

        /// <summary>
        /// Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
        /// </summary>
        [Output("utmgrpPermission")]
        public Output<Outputs.AccprofileUtmgrpPermission> UtmgrpPermission { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Output("vpngrp")]
        public Output<string> Vpngrp { get; private set; } = null!;

        /// <summary>
        /// Administrator access to WAN Opt &amp; Cache. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Output("wanoptgrp")]
        public Output<string> Wanoptgrp { get; private set; } = null!;

        /// <summary>
        /// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Output("wifi")]
        public Output<string> Wifi { get; private set; } = null!;


        /// <summary>
        /// Create a Accprofile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Accprofile(string name, AccprofileArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:system/accprofile:Accprofile", name, args ?? new AccprofileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Accprofile(string name, Input<string> id, AccprofileState? state = null, CustomResourceOptions? options = null)
            : base("fortios:system/accprofile:Accprofile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Accprofile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Accprofile Get(string name, Input<string> id, AccprofileState? state = null, CustomResourceOptions? options = null)
        {
            return new Accprofile(name, id, state, options);
        }
    }

    public sealed class AccprofileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
        /// </summary>
        [Input("admintimeout")]
        public Input<int>? Admintimeout { get; set; }

        /// <summary>
        /// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("admintimeoutOverride")]
        public Input<string>? AdmintimeoutOverride { get; set; }

        /// <summary>
        /// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("authgrp")]
        public Input<string>? Authgrp { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// FortiView. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("ftviewgrp")]
        public Input<string>? Ftviewgrp { get; set; }

        /// <summary>
        /// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("fwgrp")]
        public Input<string>? Fwgrp { get; set; }

        /// <summary>
        /// Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
        /// </summary>
        [Input("fwgrpPermission")]
        public Input<Inputs.AccprofileFwgrpPermissionArgs>? FwgrpPermission { get; set; }

        /// <summary>
        /// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("loggrp")]
        public Input<string>? Loggrp { get; set; }

        /// <summary>
        /// Custom Log &amp; Report permission. The structure of `loggrp_permission` block is documented below.
        /// </summary>
        [Input("loggrpPermission")]
        public Input<Inputs.AccprofileLoggrpPermissionArgs>? LoggrpPermission { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("netgrp")]
        public Input<string>? Netgrp { get; set; }

        /// <summary>
        /// Custom network permission. The structure of `netgrp_permission` block is documented below.
        /// </summary>
        [Input("netgrpPermission")]
        public Input<Inputs.AccprofileNetgrpPermissionArgs>? NetgrpPermission { get; set; }

        /// <summary>
        /// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Security Fabric. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("secfabgrp")]
        public Input<string>? Secfabgrp { get; set; }

        /// <summary>
        /// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("sysgrp")]
        public Input<string>? Sysgrp { get; set; }

        /// <summary>
        /// Custom system permission. The structure of `sysgrp_permission` block is documented below.
        /// </summary>
        [Input("sysgrpPermission")]
        public Input<Inputs.AccprofileSysgrpPermissionArgs>? SysgrpPermission { get; set; }

        /// <summary>
        /// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("systemDiagnostics")]
        public Input<string>? SystemDiagnostics { get; set; }

        /// <summary>
        /// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("systemExecuteSsh")]
        public Input<string>? SystemExecuteSsh { get; set; }

        /// <summary>
        /// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("systemExecuteTelnet")]
        public Input<string>? SystemExecuteTelnet { get; set; }

        /// <summary>
        /// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("utmgrp")]
        public Input<string>? Utmgrp { get; set; }

        /// <summary>
        /// Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
        /// </summary>
        [Input("utmgrpPermission")]
        public Input<Inputs.AccprofileUtmgrpPermissionArgs>? UtmgrpPermission { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("vpngrp")]
        public Input<string>? Vpngrp { get; set; }

        /// <summary>
        /// Administrator access to WAN Opt &amp; Cache. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("wanoptgrp")]
        public Input<string>? Wanoptgrp { get; set; }

        /// <summary>
        /// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("wifi")]
        public Input<string>? Wifi { get; set; }

        public AccprofileArgs()
        {
        }
        public static new AccprofileArgs Empty => new AccprofileArgs();
    }

    public sealed class AccprofileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
        /// </summary>
        [Input("admintimeout")]
        public Input<int>? Admintimeout { get; set; }

        /// <summary>
        /// Enable/disable overriding the global administrator idle timeout. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("admintimeoutOverride")]
        public Input<string>? AdmintimeoutOverride { get; set; }

        /// <summary>
        /// Administrator access to Users and Devices. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("authgrp")]
        public Input<string>? Authgrp { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// FortiView. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("ftviewgrp")]
        public Input<string>? Ftviewgrp { get; set; }

        /// <summary>
        /// Administrator access to the Firewall configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("fwgrp")]
        public Input<string>? Fwgrp { get; set; }

        /// <summary>
        /// Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
        /// </summary>
        [Input("fwgrpPermission")]
        public Input<Inputs.AccprofileFwgrpPermissionGetArgs>? FwgrpPermission { get; set; }

        /// <summary>
        /// Administrator access to Logging and Reporting including viewing log messages. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("loggrp")]
        public Input<string>? Loggrp { get; set; }

        /// <summary>
        /// Custom Log &amp; Report permission. The structure of `loggrp_permission` block is documented below.
        /// </summary>
        [Input("loggrpPermission")]
        public Input<Inputs.AccprofileLoggrpPermissionGetArgs>? LoggrpPermission { get; set; }

        /// <summary>
        /// Profile name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("netgrp")]
        public Input<string>? Netgrp { get; set; }

        /// <summary>
        /// Custom network permission. The structure of `netgrp_permission` block is documented below.
        /// </summary>
        [Input("netgrpPermission")]
        public Input<Inputs.AccprofileNetgrpPermissionGetArgs>? NetgrpPermission { get; set; }

        /// <summary>
        /// Scope of admin access: global or specific VDOM(s). Valid values: `vdom`, `global`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Security Fabric. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("secfabgrp")]
        public Input<string>? Secfabgrp { get; set; }

        /// <summary>
        /// System Configuration. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("sysgrp")]
        public Input<string>? Sysgrp { get; set; }

        /// <summary>
        /// Custom system permission. The structure of `sysgrp_permission` block is documented below.
        /// </summary>
        [Input("sysgrpPermission")]
        public Input<Inputs.AccprofileSysgrpPermissionGetArgs>? SysgrpPermission { get; set; }

        /// <summary>
        /// Enable/disable permission to run system diagnostic commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("systemDiagnostics")]
        public Input<string>? SystemDiagnostics { get; set; }

        /// <summary>
        /// Enable/disable permission to execute SSH commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("systemExecuteSsh")]
        public Input<string>? SystemExecuteSsh { get; set; }

        /// <summary>
        /// Enable/disable permission to execute TELNET commands. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("systemExecuteTelnet")]
        public Input<string>? SystemExecuteTelnet { get; set; }

        /// <summary>
        /// Administrator access to Security Profiles. Valid values: `none`, `read`, `read-write`, `custom`.
        /// </summary>
        [Input("utmgrp")]
        public Input<string>? Utmgrp { get; set; }

        /// <summary>
        /// Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
        /// </summary>
        [Input("utmgrpPermission")]
        public Input<Inputs.AccprofileUtmgrpPermissionGetArgs>? UtmgrpPermission { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Administrator access to IPsec, SSL, PPTP, and L2TP VPN. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("vpngrp")]
        public Input<string>? Vpngrp { get; set; }

        /// <summary>
        /// Administrator access to WAN Opt &amp; Cache. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("wanoptgrp")]
        public Input<string>? Wanoptgrp { get; set; }

        /// <summary>
        /// Administrator access to the WiFi controller and Switch controller. Valid values: `none`, `read`, `read-write`.
        /// </summary>
        [Input("wifi")]
        public Input<string>? Wifi { get; set; }

        public AccprofileState()
        {
        }
        public static new AccprofileState Empty => new AccprofileState();
    }
}
