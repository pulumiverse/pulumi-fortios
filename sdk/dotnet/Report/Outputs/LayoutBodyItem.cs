// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report.Outputs
{

    [OutputType]
    public sealed class LayoutBodyItem
    {
        /// <summary>
        /// Report item chart name.
        /// </summary>
        public readonly string? Chart;
        /// <summary>
        /// Report chart options. Valid values: `include-no-data`, `hide-title`, `show-caption`.
        /// </summary>
        public readonly string? ChartOptions;
        /// <summary>
        /// Report section column number.
        /// </summary>
        public readonly int? Column;
        /// <summary>
        /// Report item text content.
        /// </summary>
        public readonly string? Content;
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Control how drill down charts are shown.
        /// </summary>
        public readonly string? DrillDownItems;
        /// <summary>
        /// Control whether keys from the parent being combined or not.
        /// </summary>
        public readonly string? DrillDownTypes;
        /// <summary>
        /// Enable/disable hide item in report. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Hide;
        /// <summary>
        /// Report item ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Report item image file name.
        /// </summary>
        public readonly string? ImgSrc;
        /// <summary>
        /// Report item list component. Valid values: `bullet`, `numbered`.
        /// </summary>
        public readonly string? ListComponent;
        /// <summary>
        /// Configure report list item. The structure of `list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.LayoutBodyItemList> Lists;
        /// <summary>
        /// Report item miscellaneous component. Valid values: `hline`, `page-break`, `column-break`, `section-start`.
        /// </summary>
        public readonly string? MiscComponent;
        /// <summary>
        /// Parameters. The structure of `parameters` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.LayoutBodyItemParameter> Parameters;
        /// <summary>
        /// Report item style.
        /// </summary>
        public readonly string? Style;
        /// <summary>
        /// Table chart caption style.
        /// </summary>
        public readonly string? TableCaptionStyle;
        /// <summary>
        /// Report item table column widths.
        /// </summary>
        public readonly string? TableColumnWidths;
        /// <summary>
        /// Table chart even row style.
        /// </summary>
        public readonly string? TableEvenRowStyle;
        /// <summary>
        /// Table chart head style.
        /// </summary>
        public readonly string? TableHeadStyle;
        /// <summary>
        /// Table chart odd row style.
        /// </summary>
        public readonly string? TableOddRowStyle;
        /// <summary>
        /// Report item text component. Valid values: `text`, `heading1`, `heading2`, `heading3`.
        /// </summary>
        public readonly string? TextComponent;
        /// <summary>
        /// Report section title.
        /// </summary>
        public readonly string? Title;
        /// <summary>
        /// Value of top.
        /// </summary>
        public readonly int? TopN;
        /// <summary>
        /// Report item type. Valid values: `text`, `image`, `chart`, `misc`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private LayoutBodyItem(
            string? chart,

            string? chartOptions,

            int? column,

            string? content,

            string? description,

            string? drillDownItems,

            string? drillDownTypes,

            string? hide,

            int? id,

            string? imgSrc,

            string? listComponent,

            ImmutableArray<Outputs.LayoutBodyItemList> lists,

            string? miscComponent,

            ImmutableArray<Outputs.LayoutBodyItemParameter> parameters,

            string? style,

            string? tableCaptionStyle,

            string? tableColumnWidths,

            string? tableEvenRowStyle,

            string? tableHeadStyle,

            string? tableOddRowStyle,

            string? textComponent,

            string? title,

            int? topN,

            string? type)
        {
            Chart = chart;
            ChartOptions = chartOptions;
            Column = column;
            Content = content;
            Description = description;
            DrillDownItems = drillDownItems;
            DrillDownTypes = drillDownTypes;
            Hide = hide;
            Id = id;
            ImgSrc = imgSrc;
            ListComponent = listComponent;
            Lists = lists;
            MiscComponent = miscComponent;
            Parameters = parameters;
            Style = style;
            TableCaptionStyle = tableCaptionStyle;
            TableColumnWidths = tableColumnWidths;
            TableEvenRowStyle = tableEvenRowStyle;
            TableHeadStyle = tableHeadStyle;
            TableOddRowStyle = tableOddRowStyle;
            TextComponent = textComponent;
            Title = title;
            TopN = topN;
            Type = type;
        }
    }
}
