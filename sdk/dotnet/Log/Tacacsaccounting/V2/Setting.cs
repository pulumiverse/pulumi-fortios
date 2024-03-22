// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Log.Tacacsaccounting.V2
{
    /// <summary>
    /// Settings for TACACS+ accounting. Applies to FortiOS Version `&gt;= 7.0.2`.
    /// 
    /// ## Import
    /// 
    /// LogTacacsAccounting2 Setting can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/tacacsaccounting/v2/setting:Setting labelname LogTacacsAccounting2Setting
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:log/tacacsaccounting/v2/setting:Setting labelname LogTacacsAccounting2Setting
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:log/tacacsaccounting/v2/setting:Setting")]
    public partial class Setting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Address of TACACS+ server.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Key to access the TACACS+ server.
        /// </summary>
        [Output("serverKey")]
        public Output<string?> ServerKey { get; private set; } = null!;

        /// <summary>
        /// Source IP address for communication to TACACS+ server.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Setting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Setting(string name, SettingArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:log/tacacsaccounting/v2/setting:Setting", name, args ?? new SettingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Setting(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
            : base("fortios:log/tacacsaccounting/v2/setting:Setting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Setting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Setting Get(string name, Input<string> id, SettingState? state = null, CustomResourceOptions? options = null)
        {
            return new Setting(name, id, state, options);
        }
    }

    public sealed class SettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Address of TACACS+ server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Key to access the TACACS+ server.
        /// </summary>
        [Input("serverKey")]
        public Input<string>? ServerKey { get; set; }

        /// <summary>
        /// Source IP address for communication to TACACS+ server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingArgs()
        {
        }
        public static new SettingArgs Empty => new SettingArgs();
    }

    public sealed class SettingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Address of TACACS+ server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Key to access the TACACS+ server.
        /// </summary>
        [Input("serverKey")]
        public Input<string>? ServerKey { get; set; }

        /// <summary>
        /// Source IP address for communication to TACACS+ server.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable TACACS+ accounting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SettingState()
        {
        }
        public static new SettingState Empty => new SettingState();
    }
}
