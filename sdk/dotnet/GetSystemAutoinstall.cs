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
    public static class GetSystemAutoinstall
    {
        /// <summary>
        /// Use this data source to get information on fortios system autoinstall
        /// </summary>
        public static Task<GetSystemAutoinstallResult> InvokeAsync(GetSystemAutoinstallArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemAutoinstallResult>("fortios:index/getSystemAutoinstall:getSystemAutoinstall", args ?? new GetSystemAutoinstallArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system autoinstall
        /// </summary>
        public static Output<GetSystemAutoinstallResult> Invoke(GetSystemAutoinstallInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemAutoinstallResult>("fortios:index/getSystemAutoinstall:getSystemAutoinstall", args ?? new GetSystemAutoinstallInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemAutoinstallArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemAutoinstallArgs()
        {
        }
        public static new GetSystemAutoinstallArgs Empty => new GetSystemAutoinstallArgs();
    }

    public sealed class GetSystemAutoinstallInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemAutoinstallInvokeArgs()
        {
        }
        public static new GetSystemAutoinstallInvokeArgs Empty => new GetSystemAutoinstallInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemAutoinstallResult
    {
        /// <summary>
        /// Enable/disable auto install the config in USB disk.
        /// </summary>
        public readonly string AutoInstallConfig;
        /// <summary>
        /// Enable/disable auto install the image in USB disk.
        /// </summary>
        public readonly string AutoInstallImage;
        /// <summary>
        /// Default config file name in USB disk.
        /// </summary>
        public readonly string DefaultConfigFile;
        /// <summary>
        /// Default image file name in USB disk.
        /// </summary>
        public readonly string DefaultImageFile;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemAutoinstallResult(
            string autoInstallConfig,

            string autoInstallImage,

            string defaultConfigFile,

            string defaultImageFile,

            string id,

            string? vdomparam)
        {
            AutoInstallConfig = autoInstallConfig;
            AutoInstallImage = autoInstallImage;
            DefaultConfigFile = defaultConfigFile;
            DefaultImageFile = defaultImageFile;
            Id = id;
            Vdomparam = vdomparam;
        }
    }
}
