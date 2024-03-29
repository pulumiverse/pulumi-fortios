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
    public static class GetVdomexception
    {
        /// <summary>
        /// Use this data source to get information on an fortios system vdomexception
        /// </summary>
        public static Task<GetVdomexceptionResult> InvokeAsync(GetVdomexceptionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVdomexceptionResult>("fortios:system/getVdomexception:getVdomexception", args ?? new GetVdomexceptionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system vdomexception
        /// </summary>
        public static Output<GetVdomexceptionResult> Invoke(GetVdomexceptionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVdomexceptionResult>("fortios:system/getVdomexception:getVdomexception", args ?? new GetVdomexceptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVdomexceptionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the fosid of the desired system vdomexception.
        /// </summary>
        [Input("fosid", required: true)]
        public int Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetVdomexceptionArgs()
        {
        }
        public static new GetVdomexceptionArgs Empty => new GetVdomexceptionArgs();
    }

    public sealed class GetVdomexceptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the fosid of the desired system vdomexception.
        /// </summary>
        [Input("fosid", required: true)]
        public Input<int> Fosid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetVdomexceptionInvokeArgs()
        {
        }
        public static new GetVdomexceptionInvokeArgs Empty => new GetVdomexceptionInvokeArgs();
    }


    [OutputType]
    public sealed class GetVdomexceptionResult
    {
        /// <summary>
        /// Index &lt;1-4096&gt;.
        /// </summary>
        public readonly int Fosid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the configuration object that can be configured independently for all VDOMs.
        /// </summary>
        public readonly string Object;
        /// <summary>
        /// Object ID.
        /// </summary>
        public readonly int Oid;
        /// <summary>
        /// Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration.
        /// </summary>
        public readonly string Scope;
        public readonly string? Vdomparam;
        /// <summary>
        /// Names of the VDOMs. The structure of `vdom` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVdomexceptionVdomResult> Vdoms;

        [OutputConstructor]
        private GetVdomexceptionResult(
            int fosid,

            string id,

            string @object,

            int oid,

            string scope,

            string? vdomparam,

            ImmutableArray<Outputs.GetVdomexceptionVdomResult> vdoms)
        {
            Fosid = fosid;
            Id = id;
            Object = @object;
            Oid = oid;
            Scope = scope;
            Vdomparam = vdomparam;
            Vdoms = vdoms;
        }
    }
}
