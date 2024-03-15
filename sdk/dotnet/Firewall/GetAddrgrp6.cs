// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    public static class GetAddrgrp6
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall addrgrp6
        /// </summary>
        public static Task<GetAddrgrp6Result> InvokeAsync(GetAddrgrp6Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAddrgrp6Result>("fortios:firewall/getAddrgrp6:getAddrgrp6", args ?? new GetAddrgrp6Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall addrgrp6
        /// </summary>
        public static Output<GetAddrgrp6Result> Invoke(GetAddrgrp6InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAddrgrp6Result>("fortios:firewall/getAddrgrp6:getAddrgrp6", args ?? new GetAddrgrp6InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAddrgrp6Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall addrgrp6.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAddrgrp6Args()
        {
        }
        public static new GetAddrgrp6Args Empty => new GetAddrgrp6Args();
    }

    public sealed class GetAddrgrp6InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall addrgrp6.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAddrgrp6InvokeArgs()
        {
        }
        public static new GetAddrgrp6InvokeArgs Empty => new GetAddrgrp6InvokeArgs();
    }


    [OutputType]
    public sealed class GetAddrgrp6Result
    {
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
        /// </summary>
        public readonly int Color;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Security Fabric global object setting.
        /// </summary>
        public readonly string FabricObject;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Address objects contained within the group. The structure of `member` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAddrgrp6MemberResult> Members;
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAddrgrp6TaggingResult> Taggings;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable address group6 visibility in the GUI.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetAddrgrp6Result(
            int color,

            string comment,

            string fabricObject,

            string id,

            ImmutableArray<Outputs.GetAddrgrp6MemberResult> members,

            string name,

            ImmutableArray<Outputs.GetAddrgrp6TaggingResult> taggings,

            string uuid,

            string? vdomparam,

            string visibility)
        {
            Color = color;
            Comment = comment;
            FabricObject = fabricObject;
            Id = id;
            Members = members;
            Name = name;
            Taggings = taggings;
            Uuid = uuid;
            Vdomparam = vdomparam;
            Visibility = visibility;
        }
    }
}