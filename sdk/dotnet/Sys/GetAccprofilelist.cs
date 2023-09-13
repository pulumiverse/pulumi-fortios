// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys
{
    public static class GetAccprofilelist
    {
        /// <summary>
        /// Provides a list of `fortios.sys.Accprofile`.
        /// </summary>
        public static Task<GetAccprofilelistResult> InvokeAsync(GetAccprofilelistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccprofilelistResult>("fortios:sys/getAccprofilelist:getAccprofilelist", args ?? new GetAccprofilelistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.sys.Accprofile`.
        /// </summary>
        public static Output<GetAccprofilelistResult> Invoke(GetAccprofilelistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccprofilelistResult>("fortios:sys/getAccprofilelist:getAccprofilelist", args ?? new GetAccprofilelistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccprofilelistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAccprofilelistArgs()
        {
        }
        public static new GetAccprofilelistArgs Empty => new GetAccprofilelistArgs();
    }

    public sealed class GetAccprofilelistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAccprofilelistInvokeArgs()
        {
        }
        public static new GetAccprofilelistInvokeArgs Empty => new GetAccprofilelistInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccprofilelistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.sys.Accprofile`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAccprofilelistResult(
            string? filter,

            string id,

            ImmutableArray<string> namelists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Namelists = namelists;
            Vdomparam = vdomparam;
        }
    }
}
