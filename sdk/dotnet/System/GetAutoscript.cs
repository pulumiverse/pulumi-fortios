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
    public static class GetAutoscript
    {
        /// <summary>
        /// Use this data source to get information on an fortios system autoscript
        /// </summary>
        public static Task<GetAutoscriptResult> InvokeAsync(GetAutoscriptArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutoscriptResult>("fortios:system/getAutoscript:getAutoscript", args ?? new GetAutoscriptArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system autoscript
        /// </summary>
        public static Output<GetAutoscriptResult> Invoke(GetAutoscriptInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutoscriptResult>("fortios:system/getAutoscript:getAutoscript", args ?? new GetAutoscriptInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutoscriptArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system autoscript.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAutoscriptArgs()
        {
        }
        public static new GetAutoscriptArgs Empty => new GetAutoscriptArgs();
    }

    public sealed class GetAutoscriptInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system autoscript.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAutoscriptInvokeArgs()
        {
        }
        public static new GetAutoscriptInvokeArgs Empty => new GetAutoscriptInvokeArgs();
    }


    [OutputType]
    public sealed class GetAutoscriptResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Repeat interval in seconds.
        /// </summary>
        public readonly int Interval;
        /// <summary>
        /// Auto script name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of megabytes to limit script output to (10 - 1024, default = 10).
        /// </summary>
        public readonly int OutputSize;
        /// <summary>
        /// Number of times to repeat this script (0 = infinite).
        /// </summary>
        public readonly int Repeat;
        /// <summary>
        /// List of FortiOS CLI commands to repeat.
        /// </summary>
        public readonly string Script;
        /// <summary>
        /// Script starting mode.
        /// </summary>
        public readonly string Start;
        /// <summary>
        /// Maximum running time for this script in seconds (0 = no timeout).
        /// </summary>
        public readonly int Timeout;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAutoscriptResult(
            string id,

            int interval,

            string name,

            int outputSize,

            int repeat,

            string script,

            string start,

            int timeout,

            string? vdomparam)
        {
            Id = id;
            Interval = interval;
            Name = name;
            OutputSize = outputSize;
            Repeat = repeat;
            Script = script;
            Start = start;
            Timeout = timeout;
            Vdomparam = vdomparam;
        }
    }
}
