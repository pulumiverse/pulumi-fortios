// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Schedule
{
    public static class GetGroup
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewallschedule group
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("fortios:firewall/schedule/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewallschedule group
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("fortios:firewall/schedule/getGroup:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewallschedule group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewallschedule group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        public readonly int Color;
        /// <summary>
        /// Security Fabric global object setting.
        /// </summary>
        public readonly string FabricObject;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Schedules added to the schedule group. The structure of `member` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupMemberResult> Members;
        /// <summary>
        /// Schedule name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetGroupResult(
            int color,

            string fabricObject,

            string id,

            ImmutableArray<Outputs.GetGroupMemberResult> members,

            string name,

            string? vdomparam)
        {
            Color = color;
            FabricObject = fabricObject;
            Id = id;
            Members = members;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
