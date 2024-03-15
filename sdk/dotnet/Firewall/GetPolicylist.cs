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
    public static class GetPolicylist
    {
        /// <summary>
        /// Provides a list of `fortios.firewall.Policy`.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fortios = Pulumi.Fortios;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample1 = Fortios.Firewall.GetPolicylist.Invoke();
        /// 
        ///     var sample2 = Fortios.Firewall.GetPolicylist.Invoke(new()
        ///     {
        ///         Filter = "schedule==always&amp;action==accept,action==deny",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["output1"] = sample1,
        ///         ["sample2Output"] = sample2.Apply(getPolicylistResult =&gt; getPolicylistResult.Policyidlists),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPolicylistResult> InvokeAsync(GetPolicylistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPolicylistResult>("fortios:firewall/getPolicylist:getPolicylist", args ?? new GetPolicylistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.firewall.Policy`.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fortios = Pulumi.Fortios;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample1 = Fortios.Firewall.GetPolicylist.Invoke();
        /// 
        ///     var sample2 = Fortios.Firewall.GetPolicylist.Invoke(new()
        ///     {
        ///         Filter = "schedule==always&amp;action==accept,action==deny",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["output1"] = sample1,
        ///         ["sample2Output"] = sample2.Apply(getPolicylistResult =&gt; getPolicylistResult.Policyidlists),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPolicylistResult> Invoke(GetPolicylistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPolicylistResult>("fortios:firewall/getPolicylist:getPolicylist", args ?? new GetPolicylistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPolicylistArgs : global::Pulumi.InvokeArgs
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

        public GetPolicylistArgs()
        {
        }
        public static new GetPolicylistArgs Empty => new GetPolicylistArgs();
    }

    public sealed class GetPolicylistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetPolicylistInvokeArgs()
        {
        }
        public static new GetPolicylistInvokeArgs Empty => new GetPolicylistInvokeArgs();
    }


    [OutputType]
    public sealed class GetPolicylistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.firewall.Policy`.
        /// </summary>
        public readonly ImmutableArray<int> Policyidlists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetPolicylistResult(
            string? filter,

            string id,

            ImmutableArray<int> policyidlists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Policyidlists = policyidlists;
            Vdomparam = vdomparam;
        }
    }
}