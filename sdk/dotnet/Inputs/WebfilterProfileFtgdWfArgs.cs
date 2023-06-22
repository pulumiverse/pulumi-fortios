// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class WebfilterProfileFtgdWfArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Do not stop quota for these categories.
        /// </summary>
        [Input("exemptQuota")]
        public Input<string>? ExemptQuota { get; set; }

        [Input("filters")]
        private InputList<Inputs.WebfilterProfileFtgdWfFilterArgs>? _filters;

        /// <summary>
        /// FortiGuard filters. The structure of `filters` block is documented below.
        /// </summary>
        public InputList<Inputs.WebfilterProfileFtgdWfFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.WebfilterProfileFtgdWfFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Maximum FortiGuard quota used by single page view in seconds (excludes streams).
        /// </summary>
        [Input("maxQuotaTimeout")]
        public Input<int>? MaxQuotaTimeout { get; set; }

        /// <summary>
        /// Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Allow web filter profile overrides.
        /// </summary>
        [Input("ovrd")]
        public Input<string>? Ovrd { get; set; }

        [Input("quotas")]
        private InputList<Inputs.WebfilterProfileFtgdWfQuotaArgs>? _quotas;

        /// <summary>
        /// FortiGuard traffic quota settings. The structure of `quota` block is documented below.
        /// </summary>
        public InputList<Inputs.WebfilterProfileFtgdWfQuotaArgs> Quotas
        {
            get => _quotas ?? (_quotas = new InputList<Inputs.WebfilterProfileFtgdWfQuotaArgs>());
            set => _quotas = value;
        }

        /// <summary>
        /// Enable/disable rating CRL by URL. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("rateCrlUrls")]
        public Input<string>? RateCrlUrls { get; set; }

        /// <summary>
        /// Enable/disable rating CSS by URL. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("rateCssUrls")]
        public Input<string>? RateCssUrls { get; set; }

        /// <summary>
        /// Enable/disable rating images by URL. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("rateImageUrls")]
        public Input<string>? RateImageUrls { get; set; }

        /// <summary>
        /// Enable/disable rating JavaScript by URL. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("rateJavascriptUrls")]
        public Input<string>? RateJavascriptUrls { get; set; }

        public WebfilterProfileFtgdWfArgs()
        {
        }
        public static new WebfilterProfileFtgdWfArgs Empty => new WebfilterProfileFtgdWfArgs();
    }
}
