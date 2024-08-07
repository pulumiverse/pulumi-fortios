// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Securitypolicy
{
    /// <summary>
    /// Names of VLANs that use captive portal authentication. Applies to FortiOS Version `&lt;= 6.2.0`.
    /// 
    /// ## Import
    /// 
    /// SwitchControllerSecurityPolicy CaptivePortal can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/securitypolicy/captiveportal:Captiveportal labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:switchcontroller/securitypolicy/captiveportal:Captiveportal labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:switchcontroller/securitypolicy/captiveportal:Captiveportal")]
    public partial class Captiveportal : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Policy type. Valid values: `captive-portal`.
        /// </summary>
        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Names of VLANs that use captive portal authentication.
        /// </summary>
        [Output("vlan")]
        public Output<string> Vlan { get; private set; } = null!;


        /// <summary>
        /// Create a Captiveportal resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Captiveportal(string name, CaptiveportalArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/securitypolicy/captiveportal:Captiveportal", name, args ?? new CaptiveportalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Captiveportal(string name, Input<string> id, CaptiveportalState? state = null, CustomResourceOptions? options = null)
            : base("fortios:switchcontroller/securitypolicy/captiveportal:Captiveportal", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Captiveportal resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Captiveportal Get(string name, Input<string> id, CaptiveportalState? state = null, CustomResourceOptions? options = null)
        {
            return new Captiveportal(name, id, state, options);
        }
    }

    public sealed class CaptiveportalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy type. Valid values: `captive-portal`.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Names of VLANs that use captive portal authentication.
        /// </summary>
        [Input("vlan")]
        public Input<string>? Vlan { get; set; }

        public CaptiveportalArgs()
        {
        }
        public static new CaptiveportalArgs Empty => new CaptiveportalArgs();
    }

    public sealed class CaptiveportalState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy type. Valid values: `captive-portal`.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Names of VLANs that use captive portal authentication.
        /// </summary>
        [Input("vlan")]
        public Input<string>? Vlan { get; set; }

        public CaptiveportalState()
        {
        }
        public static new CaptiveportalState Empty => new CaptiveportalState();
    }
}
