// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Extendercontroller
{
    /// <summary>
    /// Extender controller configuration.
    /// 
    /// &gt; The resource applies to FortiOS Version &gt;= 6.4.2. For FortiOS Version &lt; 6.4.2, see `fortios.extendercontroller.Extender`.
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
    ///     var trname = new Fortios.Extendercontroller.Extender1("trname", new()
    ///     {
    ///         Authorized = "disable",
    ///         ControllerReport = new Fortios.Extendercontroller.Inputs.Extender1ControllerReportArgs
    ///         {
    ///             Interval = 300,
    ///             SignalThreshold = 10,
    ///             Status = "disable",
    ///         },
    ///         ExtName = "2932",
    ///         Fosid = "FX201E5919004031",
    ///         Modem1 = new Fortios.Extendercontroller.Inputs.Extender1Modem1Args
    ///         {
    ///             AutoSwitch = new Fortios.Extendercontroller.Inputs.Extender1Modem1AutoSwitchArgs
    ///             {
    ///                 Dataplan = "disable",
    ///                 Disconnect = "disable",
    ///                 DisconnectPeriod = 600,
    ///                 DisconnectThreshold = 3,
    ///                 Signal = "disable",
    ///                 SwitchBack = "timer",
    ///                 SwitchBackTime = "00:01",
    ///                 SwitchBackTimer = 86400,
    ///             },
    ///             ConnStatus = 0,
    ///             DefaultSim = "sim2",
    ///             Gps = "enable",
    ///             RedundantIntf = "s1",
    ///             RedundantMode = "enable",
    ///             Sim1Pin = "disable",
    ///             Sim1PinCode = "testpincode",
    ///             Sim2Pin = "disable",
    ///         },
    ///         Modem2 = new Fortios.Extendercontroller.Inputs.Extender1Modem2Args
    ///         {
    ///             AutoSwitch = new Fortios.Extendercontroller.Inputs.Extender1Modem2AutoSwitchArgs
    ///             {
    ///                 Dataplan = "disable",
    ///                 Disconnect = "disable",
    ///                 DisconnectPeriod = 600,
    ///                 DisconnectThreshold = 3,
    ///                 Signal = "disable",
    ///                 SwitchBackTime = "00:01",
    ///                 SwitchBackTimer = 86400,
    ///             },
    ///             ConnStatus = 0,
    ///             DefaultSim = "sim1",
    ///             Gps = "enable",
    ///             RedundantMode = "disable",
    ///             Sim1Pin = "disable",
    ///             Sim2Pin = "disable",
    ///         },
    ///         Vdom = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ExtenderController Extender1 can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:extendercontroller/extender1:Extender1 labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:extendercontroller/extender1:Extender1 labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:extendercontroller/extender1:Extender1")]
    public partial class Extender1 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        /// </summary>
        [Output("authorized")]
        public Output<string> Authorized { get; private set; } = null!;

        /// <summary>
        /// FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        /// </summary>
        [Output("controllerReport")]
        public Output<Outputs.Extender1ControllerReport> ControllerReport { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// FortiExtender name.
        /// </summary>
        [Output("extName")]
        public Output<string> ExtName { get; private set; } = null!;

        /// <summary>
        /// FortiExtender serial number.
        /// </summary>
        [Output("fosid")]
        public Output<string> Fosid { get; private set; } = null!;

        /// <summary>
        /// FortiExtender login password.
        /// </summary>
        [Output("loginPassword")]
        public Output<string?> LoginPassword { get; private set; } = null!;

        /// <summary>
        /// Configuration options for modem 1. The structure of `modem1` block is documented below.
        /// </summary>
        [Output("modem1")]
        public Output<Outputs.Extender1Modem1> Modem1 { get; private set; } = null!;

        /// <summary>
        /// Configuration options for modem 2. The structure of `modem2` block is documented below.
        /// </summary>
        [Output("modem2")]
        public Output<Outputs.Extender1Modem2> Modem2 { get; private set; } = null!;

        /// <summary>
        /// FortiExtender entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// VDOM
        /// </summary>
        [Output("vdom")]
        public Output<int> Vdom { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Extender1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Extender1(string name, Extender1Args args, CustomResourceOptions? options = null)
            : base("fortios:extendercontroller/extender1:Extender1", name, args ?? new Extender1Args(), MakeResourceOptions(options, ""))
        {
        }

        private Extender1(string name, Input<string> id, Extender1State? state = null, CustomResourceOptions? options = null)
            : base("fortios:extendercontroller/extender1:Extender1", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "loginPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Extender1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Extender1 Get(string name, Input<string> id, Extender1State? state = null, CustomResourceOptions? options = null)
        {
            return new Extender1(name, id, state, options);
        }
    }

    public sealed class Extender1Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authorized", required: true)]
        public Input<string> Authorized { get; set; } = null!;

        /// <summary>
        /// FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        /// </summary>
        [Input("controllerReport")]
        public Input<Inputs.Extender1ControllerReportArgs>? ControllerReport { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// FortiExtender name.
        /// </summary>
        [Input("extName")]
        public Input<string>? ExtName { get; set; }

        /// <summary>
        /// FortiExtender serial number.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        [Input("loginPassword")]
        private Input<string>? _loginPassword;

        /// <summary>
        /// FortiExtender login password.
        /// </summary>
        public Input<string>? LoginPassword
        {
            get => _loginPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _loginPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Configuration options for modem 1. The structure of `modem1` block is documented below.
        /// </summary>
        [Input("modem1")]
        public Input<Inputs.Extender1Modem1Args>? Modem1 { get; set; }

        /// <summary>
        /// Configuration options for modem 2. The structure of `modem2` block is documented below.
        /// </summary>
        [Input("modem2")]
        public Input<Inputs.Extender1Modem2Args>? Modem2 { get; set; }

        /// <summary>
        /// FortiExtender entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// VDOM
        /// </summary>
        [Input("vdom")]
        public Input<int>? Vdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Extender1Args()
        {
        }
        public static new Extender1Args Empty => new Extender1Args();
    }

    public sealed class Extender1State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
        /// </summary>
        [Input("authorized")]
        public Input<string>? Authorized { get; set; }

        /// <summary>
        /// FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
        /// </summary>
        [Input("controllerReport")]
        public Input<Inputs.Extender1ControllerReportGetArgs>? ControllerReport { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// FortiExtender name.
        /// </summary>
        [Input("extName")]
        public Input<string>? ExtName { get; set; }

        /// <summary>
        /// FortiExtender serial number.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        [Input("loginPassword")]
        private Input<string>? _loginPassword;

        /// <summary>
        /// FortiExtender login password.
        /// </summary>
        public Input<string>? LoginPassword
        {
            get => _loginPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _loginPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Configuration options for modem 1. The structure of `modem1` block is documented below.
        /// </summary>
        [Input("modem1")]
        public Input<Inputs.Extender1Modem1GetArgs>? Modem1 { get; set; }

        /// <summary>
        /// Configuration options for modem 2. The structure of `modem2` block is documented below.
        /// </summary>
        [Input("modem2")]
        public Input<Inputs.Extender1Modem2GetArgs>? Modem2 { get; set; }

        /// <summary>
        /// FortiExtender entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// VDOM
        /// </summary>
        [Input("vdom")]
        public Input<int>? Vdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Extender1State()
        {
        }
        public static new Extender1State Empty => new Extender1State();
    }
}
