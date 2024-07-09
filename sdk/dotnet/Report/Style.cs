// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Report
{
    /// <summary>
    /// Report style configuration. Applies to FortiOS Version `&lt;= 7.0.0`.
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
    ///     var trname = new Fortios.Report.Style("trname", new()
    ///     {
    ///         BorderBottom = "\" none \"",
    ///         BorderLeft = "\" none \"",
    ///         BorderRight = "\" none \"",
    ///         BorderTop = "\" none \"",
    ///         ColumnSpan = "none",
    ///         FontStyle = "normal",
    ///         FontWeight = "normal",
    ///         Options = "font text color",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Report Style can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:report/style:Style labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:report/style:Style labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:report/style:Style")]
    public partial class Style : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alignment. Valid values: `left`, `center`, `right`, `justify`.
        /// </summary>
        [Output("align")]
        public Output<string> Align { get; private set; } = null!;

        /// <summary>
        /// Background color.
        /// </summary>
        [Output("bgColor")]
        public Output<string> BgColor { get; private set; } = null!;

        /// <summary>
        /// Border bottom.
        /// </summary>
        [Output("borderBottom")]
        public Output<string> BorderBottom { get; private set; } = null!;

        /// <summary>
        /// Border left.
        /// </summary>
        [Output("borderLeft")]
        public Output<string> BorderLeft { get; private set; } = null!;

        /// <summary>
        /// Border right.
        /// </summary>
        [Output("borderRight")]
        public Output<string> BorderRight { get; private set; } = null!;

        /// <summary>
        /// Border top.
        /// </summary>
        [Output("borderTop")]
        public Output<string> BorderTop { get; private set; } = null!;

        /// <summary>
        /// Column gap.
        /// </summary>
        [Output("columnGap")]
        public Output<string> ColumnGap { get; private set; } = null!;

        /// <summary>
        /// Column span. Valid values: `none`, `all`.
        /// </summary>
        [Output("columnSpan")]
        public Output<string> ColumnSpan { get; private set; } = null!;

        /// <summary>
        /// Foreground color.
        /// </summary>
        [Output("fgColor")]
        public Output<string> FgColor { get; private set; } = null!;

        /// <summary>
        /// Font family. Valid values: `Verdana`, `Arial`, `Helvetica`, `Courier`, `Times`.
        /// </summary>
        [Output("fontFamily")]
        public Output<string> FontFamily { get; private set; } = null!;

        /// <summary>
        /// Font size.
        /// </summary>
        [Output("fontSize")]
        public Output<string> FontSize { get; private set; } = null!;

        /// <summary>
        /// Font style. Valid values: `normal`, `italic`.
        /// </summary>
        [Output("fontStyle")]
        public Output<string> FontStyle { get; private set; } = null!;

        /// <summary>
        /// Font weight. Valid values: `normal`, `bold`.
        /// </summary>
        [Output("fontWeight")]
        public Output<string> FontWeight { get; private set; } = null!;

        /// <summary>
        /// Height.
        /// </summary>
        [Output("height")]
        public Output<string> Height { get; private set; } = null!;

        /// <summary>
        /// Text line height.
        /// </summary>
        [Output("lineHeight")]
        public Output<string> LineHeight { get; private set; } = null!;

        /// <summary>
        /// Margin bottom.
        /// </summary>
        [Output("marginBottom")]
        public Output<string> MarginBottom { get; private set; } = null!;

        /// <summary>
        /// Margin left.
        /// </summary>
        [Output("marginLeft")]
        public Output<string> MarginLeft { get; private set; } = null!;

        /// <summary>
        /// Margin right.
        /// </summary>
        [Output("marginRight")]
        public Output<string> MarginRight { get; private set; } = null!;

        /// <summary>
        /// Margin top.
        /// </summary>
        [Output("marginTop")]
        public Output<string> MarginTop { get; private set; } = null!;

        /// <summary>
        /// Report style name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Report style options. Valid values: `font`, `text`, `color`, `align`, `size`, `margin`, `border`, `padding`, `column`.
        /// </summary>
        [Output("options")]
        public Output<string> Options { get; private set; } = null!;

        /// <summary>
        /// Padding bottom.
        /// </summary>
        [Output("paddingBottom")]
        public Output<string> PaddingBottom { get; private set; } = null!;

        /// <summary>
        /// Padding left.
        /// </summary>
        [Output("paddingLeft")]
        public Output<string> PaddingLeft { get; private set; } = null!;

        /// <summary>
        /// Padding right.
        /// </summary>
        [Output("paddingRight")]
        public Output<string> PaddingRight { get; private set; } = null!;

        /// <summary>
        /// Padding top.
        /// </summary>
        [Output("paddingTop")]
        public Output<string> PaddingTop { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Width.
        /// </summary>
        [Output("width")]
        public Output<string> Width { get; private set; } = null!;


        /// <summary>
        /// Create a Style resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Style(string name, StyleArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:report/style:Style", name, args ?? new StyleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Style(string name, Input<string> id, StyleState? state = null, CustomResourceOptions? options = null)
            : base("fortios:report/style:Style", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Style resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Style Get(string name, Input<string> id, StyleState? state = null, CustomResourceOptions? options = null)
        {
            return new Style(name, id, state, options);
        }
    }

    public sealed class StyleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alignment. Valid values: `left`, `center`, `right`, `justify`.
        /// </summary>
        [Input("align")]
        public Input<string>? Align { get; set; }

        /// <summary>
        /// Background color.
        /// </summary>
        [Input("bgColor")]
        public Input<string>? BgColor { get; set; }

        /// <summary>
        /// Border bottom.
        /// </summary>
        [Input("borderBottom")]
        public Input<string>? BorderBottom { get; set; }

        /// <summary>
        /// Border left.
        /// </summary>
        [Input("borderLeft")]
        public Input<string>? BorderLeft { get; set; }

        /// <summary>
        /// Border right.
        /// </summary>
        [Input("borderRight")]
        public Input<string>? BorderRight { get; set; }

        /// <summary>
        /// Border top.
        /// </summary>
        [Input("borderTop")]
        public Input<string>? BorderTop { get; set; }

        /// <summary>
        /// Column gap.
        /// </summary>
        [Input("columnGap")]
        public Input<string>? ColumnGap { get; set; }

        /// <summary>
        /// Column span. Valid values: `none`, `all`.
        /// </summary>
        [Input("columnSpan")]
        public Input<string>? ColumnSpan { get; set; }

        /// <summary>
        /// Foreground color.
        /// </summary>
        [Input("fgColor")]
        public Input<string>? FgColor { get; set; }

        /// <summary>
        /// Font family. Valid values: `Verdana`, `Arial`, `Helvetica`, `Courier`, `Times`.
        /// </summary>
        [Input("fontFamily")]
        public Input<string>? FontFamily { get; set; }

        /// <summary>
        /// Font size.
        /// </summary>
        [Input("fontSize")]
        public Input<string>? FontSize { get; set; }

        /// <summary>
        /// Font style. Valid values: `normal`, `italic`.
        /// </summary>
        [Input("fontStyle")]
        public Input<string>? FontStyle { get; set; }

        /// <summary>
        /// Font weight. Valid values: `normal`, `bold`.
        /// </summary>
        [Input("fontWeight")]
        public Input<string>? FontWeight { get; set; }

        /// <summary>
        /// Height.
        /// </summary>
        [Input("height")]
        public Input<string>? Height { get; set; }

        /// <summary>
        /// Text line height.
        /// </summary>
        [Input("lineHeight")]
        public Input<string>? LineHeight { get; set; }

        /// <summary>
        /// Margin bottom.
        /// </summary>
        [Input("marginBottom")]
        public Input<string>? MarginBottom { get; set; }

        /// <summary>
        /// Margin left.
        /// </summary>
        [Input("marginLeft")]
        public Input<string>? MarginLeft { get; set; }

        /// <summary>
        /// Margin right.
        /// </summary>
        [Input("marginRight")]
        public Input<string>? MarginRight { get; set; }

        /// <summary>
        /// Margin top.
        /// </summary>
        [Input("marginTop")]
        public Input<string>? MarginTop { get; set; }

        /// <summary>
        /// Report style name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Report style options. Valid values: `font`, `text`, `color`, `align`, `size`, `margin`, `border`, `padding`, `column`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Padding bottom.
        /// </summary>
        [Input("paddingBottom")]
        public Input<string>? PaddingBottom { get; set; }

        /// <summary>
        /// Padding left.
        /// </summary>
        [Input("paddingLeft")]
        public Input<string>? PaddingLeft { get; set; }

        /// <summary>
        /// Padding right.
        /// </summary>
        [Input("paddingRight")]
        public Input<string>? PaddingRight { get; set; }

        /// <summary>
        /// Padding top.
        /// </summary>
        [Input("paddingTop")]
        public Input<string>? PaddingTop { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Width.
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public StyleArgs()
        {
        }
        public static new StyleArgs Empty => new StyleArgs();
    }

    public sealed class StyleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alignment. Valid values: `left`, `center`, `right`, `justify`.
        /// </summary>
        [Input("align")]
        public Input<string>? Align { get; set; }

        /// <summary>
        /// Background color.
        /// </summary>
        [Input("bgColor")]
        public Input<string>? BgColor { get; set; }

        /// <summary>
        /// Border bottom.
        /// </summary>
        [Input("borderBottom")]
        public Input<string>? BorderBottom { get; set; }

        /// <summary>
        /// Border left.
        /// </summary>
        [Input("borderLeft")]
        public Input<string>? BorderLeft { get; set; }

        /// <summary>
        /// Border right.
        /// </summary>
        [Input("borderRight")]
        public Input<string>? BorderRight { get; set; }

        /// <summary>
        /// Border top.
        /// </summary>
        [Input("borderTop")]
        public Input<string>? BorderTop { get; set; }

        /// <summary>
        /// Column gap.
        /// </summary>
        [Input("columnGap")]
        public Input<string>? ColumnGap { get; set; }

        /// <summary>
        /// Column span. Valid values: `none`, `all`.
        /// </summary>
        [Input("columnSpan")]
        public Input<string>? ColumnSpan { get; set; }

        /// <summary>
        /// Foreground color.
        /// </summary>
        [Input("fgColor")]
        public Input<string>? FgColor { get; set; }

        /// <summary>
        /// Font family. Valid values: `Verdana`, `Arial`, `Helvetica`, `Courier`, `Times`.
        /// </summary>
        [Input("fontFamily")]
        public Input<string>? FontFamily { get; set; }

        /// <summary>
        /// Font size.
        /// </summary>
        [Input("fontSize")]
        public Input<string>? FontSize { get; set; }

        /// <summary>
        /// Font style. Valid values: `normal`, `italic`.
        /// </summary>
        [Input("fontStyle")]
        public Input<string>? FontStyle { get; set; }

        /// <summary>
        /// Font weight. Valid values: `normal`, `bold`.
        /// </summary>
        [Input("fontWeight")]
        public Input<string>? FontWeight { get; set; }

        /// <summary>
        /// Height.
        /// </summary>
        [Input("height")]
        public Input<string>? Height { get; set; }

        /// <summary>
        /// Text line height.
        /// </summary>
        [Input("lineHeight")]
        public Input<string>? LineHeight { get; set; }

        /// <summary>
        /// Margin bottom.
        /// </summary>
        [Input("marginBottom")]
        public Input<string>? MarginBottom { get; set; }

        /// <summary>
        /// Margin left.
        /// </summary>
        [Input("marginLeft")]
        public Input<string>? MarginLeft { get; set; }

        /// <summary>
        /// Margin right.
        /// </summary>
        [Input("marginRight")]
        public Input<string>? MarginRight { get; set; }

        /// <summary>
        /// Margin top.
        /// </summary>
        [Input("marginTop")]
        public Input<string>? MarginTop { get; set; }

        /// <summary>
        /// Report style name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Report style options. Valid values: `font`, `text`, `color`, `align`, `size`, `margin`, `border`, `padding`, `column`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Padding bottom.
        /// </summary>
        [Input("paddingBottom")]
        public Input<string>? PaddingBottom { get; set; }

        /// <summary>
        /// Padding left.
        /// </summary>
        [Input("paddingLeft")]
        public Input<string>? PaddingLeft { get; set; }

        /// <summary>
        /// Padding right.
        /// </summary>
        [Input("paddingRight")]
        public Input<string>? PaddingRight { get; set; }

        /// <summary>
        /// Padding top.
        /// </summary>
        [Input("paddingTop")]
        public Input<string>? PaddingTop { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Width.
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public StyleState()
        {
        }
        public static new StyleState Empty => new StyleState();
    }
}
