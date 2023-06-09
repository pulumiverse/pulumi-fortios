// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    public static class GetSystemInterfacelist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemInterface`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fortios = Pulumi.Fortios;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample1 = Fortios.GetSystemInterfacelist.Invoke(new()
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
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSystemInterfacelistResult> InvokeAsync(GetSystemInterfacelistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemInterfacelistResult>("fortios:index/getSystemInterfacelist:getSystemInterfacelist", args ?? new GetSystemInterfacelistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemInterface`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Fortios = Pulumi.Fortios;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sample1 = Fortios.GetSystemInterfacelist.Invoke(new()
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
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSystemInterfacelistResult> Invoke(GetSystemInterfacelistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemInterfacelistResult>("fortios:index/getSystemInterfacelist:getSystemInterfacelist", args ?? new GetSystemInterfacelistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemInterfacelistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemInterfacelistArgs()
        {
        }
        public static new GetSystemInterfacelistArgs Empty => new GetSystemInterfacelistArgs();
    }

    public sealed class GetSystemInterfacelistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemInterfacelistInvokeArgs()
        {
        }
        public static new GetSystemInterfacelistInvokeArgs Empty => new GetSystemInterfacelistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemInterfacelistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.SystemInterface`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemInterfacelistResult(
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
