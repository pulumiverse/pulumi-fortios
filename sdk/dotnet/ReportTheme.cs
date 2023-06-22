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
    /// <summary>
    /// Report themes configuration Applies to FortiOS Version `&lt;= 7.0.0`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.ReportTheme("trname", new()
    ///     {
    ///         ColumnCount = "1",
    ///         GraphChartStyle = "PS",
    ///         PageOrient = "portrait",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Report Theme can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/reportTheme:ReportTheme labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/reportTheme:ReportTheme labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/reportTheme:ReportTheme")]
    public partial class ReportTheme : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Bullet list style.
        /// </summary>
        [Output("bulletListStyle")]
        public Output<string> BulletListStyle { get; private set; } = null!;

        /// <summary>
        /// Report page column count. Valid values: `1`, `2`, `3`.
        /// </summary>
        [Output("columnCount")]
        public Output<string> ColumnCount { get; private set; } = null!;

        /// <summary>
        /// Default HTML report style.
        /// </summary>
        [Output("defaultHtmlStyle")]
        public Output<string> DefaultHtmlStyle { get; private set; } = null!;

        /// <summary>
        /// Default PDF report style.
        /// </summary>
        [Output("defaultPdfStyle")]
        public Output<string> DefaultPdfStyle { get; private set; } = null!;

        /// <summary>
        /// Graph chart style.
        /// </summary>
        [Output("graphChartStyle")]
        public Output<string> GraphChartStyle { get; private set; } = null!;

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Output("heading1Style")]
        public Output<string> Heading1Style { get; private set; } = null!;

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Output("heading2Style")]
        public Output<string> Heading2Style { get; private set; } = null!;

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Output("heading3Style")]
        public Output<string> Heading3Style { get; private set; } = null!;

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Output("heading4Style")]
        public Output<string> Heading4Style { get; private set; } = null!;

        /// <summary>
        /// Horizontal line style.
        /// </summary>
        [Output("hlineStyle")]
        public Output<string> HlineStyle { get; private set; } = null!;

        /// <summary>
        /// Image style.
        /// </summary>
        [Output("imageStyle")]
        public Output<string> ImageStyle { get; private set; } = null!;

        /// <summary>
        /// Report theme name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Normal text style.
        /// </summary>
        [Output("normalTextStyle")]
        public Output<string> NormalTextStyle { get; private set; } = null!;

        /// <summary>
        /// Numbered list style.
        /// </summary>
        [Output("numberedListStyle")]
        public Output<string> NumberedListStyle { get; private set; } = null!;

        /// <summary>
        /// Report page footer style.
        /// </summary>
        [Output("pageFooterStyle")]
        public Output<string> PageFooterStyle { get; private set; } = null!;

        /// <summary>
        /// Report page header style.
        /// </summary>
        [Output("pageHeaderStyle")]
        public Output<string> PageHeaderStyle { get; private set; } = null!;

        /// <summary>
        /// Report page orientation. Valid values: `portrait`, `landscape`.
        /// </summary>
        [Output("pageOrient")]
        public Output<string> PageOrient { get; private set; } = null!;

        /// <summary>
        /// Report page style.
        /// </summary>
        [Output("pageStyle")]
        public Output<string> PageStyle { get; private set; } = null!;

        /// <summary>
        /// Report subtitle style.
        /// </summary>
        [Output("reportSubtitleStyle")]
        public Output<string> ReportSubtitleStyle { get; private set; } = null!;

        /// <summary>
        /// Report title style.
        /// </summary>
        [Output("reportTitleStyle")]
        public Output<string> ReportTitleStyle { get; private set; } = null!;

        /// <summary>
        /// Table chart caption style.
        /// </summary>
        [Output("tableChartCaptionStyle")]
        public Output<string> TableChartCaptionStyle { get; private set; } = null!;

        /// <summary>
        /// Table chart even row style.
        /// </summary>
        [Output("tableChartEvenRowStyle")]
        public Output<string> TableChartEvenRowStyle { get; private set; } = null!;

        /// <summary>
        /// Table chart head row style.
        /// </summary>
        [Output("tableChartHeadStyle")]
        public Output<string> TableChartHeadStyle { get; private set; } = null!;

        /// <summary>
        /// Table chart odd row style.
        /// </summary>
        [Output("tableChartOddRowStyle")]
        public Output<string> TableChartOddRowStyle { get; private set; } = null!;

        /// <summary>
        /// Table chart style.
        /// </summary>
        [Output("tableChartStyle")]
        public Output<string> TableChartStyle { get; private set; } = null!;

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Output("tocHeading1Style")]
        public Output<string> TocHeading1Style { get; private set; } = null!;

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Output("tocHeading2Style")]
        public Output<string> TocHeading2Style { get; private set; } = null!;

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Output("tocHeading3Style")]
        public Output<string> TocHeading3Style { get; private set; } = null!;

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Output("tocHeading4Style")]
        public Output<string> TocHeading4Style { get; private set; } = null!;

        /// <summary>
        /// Table of contents title style.
        /// </summary>
        [Output("tocTitleStyle")]
        public Output<string> TocTitleStyle { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a ReportTheme resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReportTheme(string name, ReportThemeArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/reportTheme:ReportTheme", name, args ?? new ReportThemeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReportTheme(string name, Input<string> id, ReportThemeState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/reportTheme:ReportTheme", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReportTheme resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReportTheme Get(string name, Input<string> id, ReportThemeState? state = null, CustomResourceOptions? options = null)
        {
            return new ReportTheme(name, id, state, options);
        }
    }

    public sealed class ReportThemeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bullet list style.
        /// </summary>
        [Input("bulletListStyle")]
        public Input<string>? BulletListStyle { get; set; }

        /// <summary>
        /// Report page column count. Valid values: `1`, `2`, `3`.
        /// </summary>
        [Input("columnCount")]
        public Input<string>? ColumnCount { get; set; }

        /// <summary>
        /// Default HTML report style.
        /// </summary>
        [Input("defaultHtmlStyle")]
        public Input<string>? DefaultHtmlStyle { get; set; }

        /// <summary>
        /// Default PDF report style.
        /// </summary>
        [Input("defaultPdfStyle")]
        public Input<string>? DefaultPdfStyle { get; set; }

        /// <summary>
        /// Graph chart style.
        /// </summary>
        [Input("graphChartStyle")]
        public Input<string>? GraphChartStyle { get; set; }

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Input("heading1Style")]
        public Input<string>? Heading1Style { get; set; }

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Input("heading2Style")]
        public Input<string>? Heading2Style { get; set; }

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Input("heading3Style")]
        public Input<string>? Heading3Style { get; set; }

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Input("heading4Style")]
        public Input<string>? Heading4Style { get; set; }

        /// <summary>
        /// Horizontal line style.
        /// </summary>
        [Input("hlineStyle")]
        public Input<string>? HlineStyle { get; set; }

        /// <summary>
        /// Image style.
        /// </summary>
        [Input("imageStyle")]
        public Input<string>? ImageStyle { get; set; }

        /// <summary>
        /// Report theme name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Normal text style.
        /// </summary>
        [Input("normalTextStyle")]
        public Input<string>? NormalTextStyle { get; set; }

        /// <summary>
        /// Numbered list style.
        /// </summary>
        [Input("numberedListStyle")]
        public Input<string>? NumberedListStyle { get; set; }

        /// <summary>
        /// Report page footer style.
        /// </summary>
        [Input("pageFooterStyle")]
        public Input<string>? PageFooterStyle { get; set; }

        /// <summary>
        /// Report page header style.
        /// </summary>
        [Input("pageHeaderStyle")]
        public Input<string>? PageHeaderStyle { get; set; }

        /// <summary>
        /// Report page orientation. Valid values: `portrait`, `landscape`.
        /// </summary>
        [Input("pageOrient")]
        public Input<string>? PageOrient { get; set; }

        /// <summary>
        /// Report page style.
        /// </summary>
        [Input("pageStyle")]
        public Input<string>? PageStyle { get; set; }

        /// <summary>
        /// Report subtitle style.
        /// </summary>
        [Input("reportSubtitleStyle")]
        public Input<string>? ReportSubtitleStyle { get; set; }

        /// <summary>
        /// Report title style.
        /// </summary>
        [Input("reportTitleStyle")]
        public Input<string>? ReportTitleStyle { get; set; }

        /// <summary>
        /// Table chart caption style.
        /// </summary>
        [Input("tableChartCaptionStyle")]
        public Input<string>? TableChartCaptionStyle { get; set; }

        /// <summary>
        /// Table chart even row style.
        /// </summary>
        [Input("tableChartEvenRowStyle")]
        public Input<string>? TableChartEvenRowStyle { get; set; }

        /// <summary>
        /// Table chart head row style.
        /// </summary>
        [Input("tableChartHeadStyle")]
        public Input<string>? TableChartHeadStyle { get; set; }

        /// <summary>
        /// Table chart odd row style.
        /// </summary>
        [Input("tableChartOddRowStyle")]
        public Input<string>? TableChartOddRowStyle { get; set; }

        /// <summary>
        /// Table chart style.
        /// </summary>
        [Input("tableChartStyle")]
        public Input<string>? TableChartStyle { get; set; }

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Input("tocHeading1Style")]
        public Input<string>? TocHeading1Style { get; set; }

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Input("tocHeading2Style")]
        public Input<string>? TocHeading2Style { get; set; }

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Input("tocHeading3Style")]
        public Input<string>? TocHeading3Style { get; set; }

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Input("tocHeading4Style")]
        public Input<string>? TocHeading4Style { get; set; }

        /// <summary>
        /// Table of contents title style.
        /// </summary>
        [Input("tocTitleStyle")]
        public Input<string>? TocTitleStyle { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ReportThemeArgs()
        {
        }
        public static new ReportThemeArgs Empty => new ReportThemeArgs();
    }

    public sealed class ReportThemeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bullet list style.
        /// </summary>
        [Input("bulletListStyle")]
        public Input<string>? BulletListStyle { get; set; }

        /// <summary>
        /// Report page column count. Valid values: `1`, `2`, `3`.
        /// </summary>
        [Input("columnCount")]
        public Input<string>? ColumnCount { get; set; }

        /// <summary>
        /// Default HTML report style.
        /// </summary>
        [Input("defaultHtmlStyle")]
        public Input<string>? DefaultHtmlStyle { get; set; }

        /// <summary>
        /// Default PDF report style.
        /// </summary>
        [Input("defaultPdfStyle")]
        public Input<string>? DefaultPdfStyle { get; set; }

        /// <summary>
        /// Graph chart style.
        /// </summary>
        [Input("graphChartStyle")]
        public Input<string>? GraphChartStyle { get; set; }

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Input("heading1Style")]
        public Input<string>? Heading1Style { get; set; }

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Input("heading2Style")]
        public Input<string>? Heading2Style { get; set; }

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Input("heading3Style")]
        public Input<string>? Heading3Style { get; set; }

        /// <summary>
        /// Report heading style.
        /// </summary>
        [Input("heading4Style")]
        public Input<string>? Heading4Style { get; set; }

        /// <summary>
        /// Horizontal line style.
        /// </summary>
        [Input("hlineStyle")]
        public Input<string>? HlineStyle { get; set; }

        /// <summary>
        /// Image style.
        /// </summary>
        [Input("imageStyle")]
        public Input<string>? ImageStyle { get; set; }

        /// <summary>
        /// Report theme name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Normal text style.
        /// </summary>
        [Input("normalTextStyle")]
        public Input<string>? NormalTextStyle { get; set; }

        /// <summary>
        /// Numbered list style.
        /// </summary>
        [Input("numberedListStyle")]
        public Input<string>? NumberedListStyle { get; set; }

        /// <summary>
        /// Report page footer style.
        /// </summary>
        [Input("pageFooterStyle")]
        public Input<string>? PageFooterStyle { get; set; }

        /// <summary>
        /// Report page header style.
        /// </summary>
        [Input("pageHeaderStyle")]
        public Input<string>? PageHeaderStyle { get; set; }

        /// <summary>
        /// Report page orientation. Valid values: `portrait`, `landscape`.
        /// </summary>
        [Input("pageOrient")]
        public Input<string>? PageOrient { get; set; }

        /// <summary>
        /// Report page style.
        /// </summary>
        [Input("pageStyle")]
        public Input<string>? PageStyle { get; set; }

        /// <summary>
        /// Report subtitle style.
        /// </summary>
        [Input("reportSubtitleStyle")]
        public Input<string>? ReportSubtitleStyle { get; set; }

        /// <summary>
        /// Report title style.
        /// </summary>
        [Input("reportTitleStyle")]
        public Input<string>? ReportTitleStyle { get; set; }

        /// <summary>
        /// Table chart caption style.
        /// </summary>
        [Input("tableChartCaptionStyle")]
        public Input<string>? TableChartCaptionStyle { get; set; }

        /// <summary>
        /// Table chart even row style.
        /// </summary>
        [Input("tableChartEvenRowStyle")]
        public Input<string>? TableChartEvenRowStyle { get; set; }

        /// <summary>
        /// Table chart head row style.
        /// </summary>
        [Input("tableChartHeadStyle")]
        public Input<string>? TableChartHeadStyle { get; set; }

        /// <summary>
        /// Table chart odd row style.
        /// </summary>
        [Input("tableChartOddRowStyle")]
        public Input<string>? TableChartOddRowStyle { get; set; }

        /// <summary>
        /// Table chart style.
        /// </summary>
        [Input("tableChartStyle")]
        public Input<string>? TableChartStyle { get; set; }

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Input("tocHeading1Style")]
        public Input<string>? TocHeading1Style { get; set; }

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Input("tocHeading2Style")]
        public Input<string>? TocHeading2Style { get; set; }

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Input("tocHeading3Style")]
        public Input<string>? TocHeading3Style { get; set; }

        /// <summary>
        /// Table of contents heading style.
        /// </summary>
        [Input("tocHeading4Style")]
        public Input<string>? TocHeading4Style { get; set; }

        /// <summary>
        /// Table of contents title style.
        /// </summary>
        [Input("tocTitleStyle")]
        public Input<string>? TocTitleStyle { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public ReportThemeState()
        {
        }
        public static new ReportThemeState Empty => new ReportThemeState();
    }
}
