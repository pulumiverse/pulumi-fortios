// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Fmg
{
    /// <summary>
    /// This resource supports Create/Read/Update/Delete system adom for FortiManager.
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
    ///     var test1 = new Fortios.Fmg.SystemAdom("test1", new()
    ///     {
    ///         ActionWhenConflictsOccurDuringPolicyCheck = "Continue",
    ///         AutoPushPolicyPackagesWhenDeviceBackOnline = "Enable",
    ///         CentralManagementFortiap = true,
    ///         CentralManagementSdwan = false,
    ///         CentralManagementVpn = false,
    ///         Mode = "Normal",
    ///         PerformPolicyCheckBeforeEveryInstall = true,
    ///         Status = 1,
    ///         Type = "FortiCarrier",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [FortiosResourceType("fortios:fmg/systemAdom:SystemAdom")]
    public partial class SystemAdom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// True or False.
        /// </summary>
        [Output("actionWhenConflictsOccurDuringPolicyCheck")]
        public Output<string?> ActionWhenConflictsOccurDuringPolicyCheck { get; private set; } = null!;

        /// <summary>
        /// True or False.
        /// </summary>
        [Output("autoPushPolicyPackagesWhenDeviceBackOnline")]
        public Output<string?> AutoPushPolicyPackagesWhenDeviceBackOnline { get; private set; } = null!;

        /// <summary>
        /// True or False.
        /// </summary>
        [Output("centralManagementFortiap")]
        public Output<bool?> CentralManagementFortiap { get; private set; } = null!;

        /// <summary>
        /// True or False.
        /// </summary>
        [Output("centralManagementSdwan")]
        public Output<bool?> CentralManagementSdwan { get; private set; } = null!;

        /// <summary>
        /// True or False.
        /// </summary>
        [Output("centralManagementVpn")]
        public Output<bool?> CentralManagementVpn { get; private set; } = null!;

        /// <summary>
        /// Adom mode: Normal or Backup.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// Administrative Domain name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// True or False.
        /// </summary>
        [Output("performPolicyCheckBeforeEveryInstall")]
        public Output<bool?> PerformPolicyCheckBeforeEveryInstall { get; private set; } = null!;

        /// <summary>
        /// Adom status. 0 means off and 1 means on.
        /// </summary>
        [Output("status")]
        public Output<int?> Status { get; private set; } = null!;

        /// <summary>
        /// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SystemAdom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemAdom(string name, SystemAdomArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:fmg/systemAdom:SystemAdom", name, args ?? new SystemAdomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemAdom(string name, Input<string> id, SystemAdomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:fmg/systemAdom:SystemAdom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemAdom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemAdom Get(string name, Input<string> id, SystemAdomState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemAdom(name, id, state, options);
        }
    }

    public sealed class SystemAdomArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True or False.
        /// </summary>
        [Input("actionWhenConflictsOccurDuringPolicyCheck")]
        public Input<string>? ActionWhenConflictsOccurDuringPolicyCheck { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("autoPushPolicyPackagesWhenDeviceBackOnline")]
        public Input<string>? AutoPushPolicyPackagesWhenDeviceBackOnline { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("centralManagementFortiap")]
        public Input<bool>? CentralManagementFortiap { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("centralManagementSdwan")]
        public Input<bool>? CentralManagementSdwan { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("centralManagementVpn")]
        public Input<bool>? CentralManagementVpn { get; set; }

        /// <summary>
        /// Adom mode: Normal or Backup.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Administrative Domain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("performPolicyCheckBeforeEveryInstall")]
        public Input<bool>? PerformPolicyCheckBeforeEveryInstall { get; set; }

        /// <summary>
        /// Adom status. 0 means off and 1 means on.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SystemAdomArgs()
        {
        }
        public static new SystemAdomArgs Empty => new SystemAdomArgs();
    }

    public sealed class SystemAdomState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// True or False.
        /// </summary>
        [Input("actionWhenConflictsOccurDuringPolicyCheck")]
        public Input<string>? ActionWhenConflictsOccurDuringPolicyCheck { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("autoPushPolicyPackagesWhenDeviceBackOnline")]
        public Input<string>? AutoPushPolicyPackagesWhenDeviceBackOnline { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("centralManagementFortiap")]
        public Input<bool>? CentralManagementFortiap { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("centralManagementSdwan")]
        public Input<bool>? CentralManagementSdwan { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("centralManagementVpn")]
        public Input<bool>? CentralManagementVpn { get; set; }

        /// <summary>
        /// Adom mode: Normal or Backup.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Administrative Domain name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// True or False.
        /// </summary>
        [Input("performPolicyCheckBeforeEveryInstall")]
        public Input<bool>? PerformPolicyCheckBeforeEveryInstall { get; set; }

        /// <summary>
        /// Adom status. 0 means off and 1 means on.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Domain type, Enum: ["FortiGate", "FortiCarrier], default is "FortiCarrier".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SystemAdomState()
        {
        }
        public static new SystemAdomState Empty => new SystemAdomState();
    }
}
