// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// SwitchControllerSecurityPolicy LocalAccess can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollersecuritypolicyLocalaccess:SwitchcontrollersecuritypolicyLocalaccess labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollersecuritypolicyLocalaccess:SwitchcontrollersecuritypolicyLocalaccess labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollersecuritypolicyLocalaccess:SwitchcontrollersecuritypolicyLocalaccess")]
    public partial class SwitchcontrollersecuritypolicyLocalaccess : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
        /// </summary>
        [Output("internalAllowaccess")]
        public Output<string> InternalAllowaccess { get; private set; } = null!;

        /// <summary>
        /// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
        /// </summary>
        [Output("mgmtAllowaccess")]
        public Output<string> MgmtAllowaccess { get; private set; } = null!;

        /// <summary>
        /// Policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollersecuritypolicyLocalaccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollersecuritypolicyLocalaccess(string name, SwitchcontrollersecuritypolicyLocalaccessArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollersecuritypolicyLocalaccess:SwitchcontrollersecuritypolicyLocalaccess", name, args ?? new SwitchcontrollersecuritypolicyLocalaccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollersecuritypolicyLocalaccess(string name, Input<string> id, SwitchcontrollersecuritypolicyLocalaccessState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollersecuritypolicyLocalaccess:SwitchcontrollersecuritypolicyLocalaccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchcontrollersecuritypolicyLocalaccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollersecuritypolicyLocalaccess Get(string name, Input<string> id, SwitchcontrollersecuritypolicyLocalaccessState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollersecuritypolicyLocalaccess(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollersecuritypolicyLocalaccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
        /// </summary>
        [Input("internalAllowaccess")]
        public Input<string>? InternalAllowaccess { get; set; }

        /// <summary>
        /// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
        /// </summary>
        [Input("mgmtAllowaccess")]
        public Input<string>? MgmtAllowaccess { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollersecuritypolicyLocalaccessArgs()
        {
        }
        public static new SwitchcontrollersecuritypolicyLocalaccessArgs Empty => new SwitchcontrollersecuritypolicyLocalaccessArgs();
    }

    public sealed class SwitchcontrollersecuritypolicyLocalaccessState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
        /// </summary>
        [Input("internalAllowaccess")]
        public Input<string>? InternalAllowaccess { get; set; }

        /// <summary>
        /// Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
        /// </summary>
        [Input("mgmtAllowaccess")]
        public Input<string>? MgmtAllowaccess { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollersecuritypolicyLocalaccessState()
        {
        }
        public static new SwitchcontrollersecuritypolicyLocalaccessState Empty => new SwitchcontrollersecuritypolicyLocalaccessState();
    }
}
