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
    public static class GetInterfacelist
    {
        /// <summary>
        /// Provides a list of `fortios.system.Interface`.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fortios = Pulumi.Fortios;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample1 = Fortios.System.GetInterfacelist.Invoke(new()
        ///     {
        ///         Filter = "name!=port1",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["output1"] = data.Fortios_system_interfacelist.Sample2.Namelist,
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetInterfacelistResult> InvokeAsync(GetInterfacelistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInterfacelistResult>("fortios:system/getInterfacelist:getInterfacelist", args ?? new GetInterfacelistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system.Interface`.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fortios = Pulumi.Fortios;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample1 = Fortios.System.GetInterfacelist.Invoke(new()
        ///     {
        ///         Filter = "name!=port1",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["output1"] = data.Fortios_system_interfacelist.Sample2.Namelist,
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetInterfacelistResult> Invoke(GetInterfacelistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInterfacelistResult>("fortios:system/getInterfacelist:getInterfacelist", args ?? new GetInterfacelistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInterfacelistArgs : global::Pulumi.InvokeArgs
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

        public GetInterfacelistArgs()
        {
        }
        public static new GetInterfacelistArgs Empty => new GetInterfacelistArgs();
    }

    public sealed class GetInterfacelistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetInterfacelistInvokeArgs()
        {
        }
        public static new GetInterfacelistInvokeArgs Empty => new GetInterfacelistInvokeArgs();
    }


    [OutputType]
    public sealed class GetInterfacelistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.system.Interface`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetInterfacelistResult(
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
