// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Report style configuration. Applies to FortiOS Version `<= 7.0.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.report.Style("trname", {
 *     borderBottom: "\" none \"",
 *     borderLeft: "\" none \"",
 *     borderRight: "\" none \"",
 *     borderTop: "\" none \"",
 *     columnSpan: "none",
 *     fontStyle: "normal",
 *     fontWeight: "normal",
 *     options: "font text color",
 * });
 * ```
 *
 * ## Import
 *
 * Report Style can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:report/style:Style labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:report/style:Style labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Style extends pulumi.CustomResource {
    /**
     * Get an existing Style resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StyleState, opts?: pulumi.CustomResourceOptions): Style {
        return new Style(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:report/style:Style';

    /**
     * Returns true if the given object is an instance of Style.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Style {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Style.__pulumiType;
    }

    /**
     * Alignment. Valid values: `left`, `center`, `right`, `justify`.
     */
    public readonly align!: pulumi.Output<string>;
    /**
     * Background color.
     */
    public readonly bgColor!: pulumi.Output<string>;
    /**
     * Border bottom.
     */
    public readonly borderBottom!: pulumi.Output<string>;
    /**
     * Border left.
     */
    public readonly borderLeft!: pulumi.Output<string>;
    /**
     * Border right.
     */
    public readonly borderRight!: pulumi.Output<string>;
    /**
     * Border top.
     */
    public readonly borderTop!: pulumi.Output<string>;
    /**
     * Column gap.
     */
    public readonly columnGap!: pulumi.Output<string>;
    /**
     * Column span. Valid values: `none`, `all`.
     */
    public readonly columnSpan!: pulumi.Output<string>;
    /**
     * Foreground color.
     */
    public readonly fgColor!: pulumi.Output<string>;
    /**
     * Font family. Valid values: `Verdana`, `Arial`, `Helvetica`, `Courier`, `Times`.
     */
    public readonly fontFamily!: pulumi.Output<string>;
    /**
     * Font size.
     */
    public readonly fontSize!: pulumi.Output<string>;
    /**
     * Font style. Valid values: `normal`, `italic`.
     */
    public readonly fontStyle!: pulumi.Output<string>;
    /**
     * Font weight. Valid values: `normal`, `bold`.
     */
    public readonly fontWeight!: pulumi.Output<string>;
    /**
     * Height.
     */
    public readonly height!: pulumi.Output<string>;
    /**
     * Text line height.
     */
    public readonly lineHeight!: pulumi.Output<string>;
    /**
     * Margin bottom.
     */
    public readonly marginBottom!: pulumi.Output<string>;
    /**
     * Margin left.
     */
    public readonly marginLeft!: pulumi.Output<string>;
    /**
     * Margin right.
     */
    public readonly marginRight!: pulumi.Output<string>;
    /**
     * Margin top.
     */
    public readonly marginTop!: pulumi.Output<string>;
    /**
     * Report style name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Report style options. Valid values: `font`, `text`, `color`, `align`, `size`, `margin`, `border`, `padding`, `column`.
     */
    public readonly options!: pulumi.Output<string>;
    /**
     * Padding bottom.
     */
    public readonly paddingBottom!: pulumi.Output<string>;
    /**
     * Padding left.
     */
    public readonly paddingLeft!: pulumi.Output<string>;
    /**
     * Padding right.
     */
    public readonly paddingRight!: pulumi.Output<string>;
    /**
     * Padding top.
     */
    public readonly paddingTop!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string>;
    /**
     * Width.
     */
    public readonly width!: pulumi.Output<string>;

    /**
     * Create a Style resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StyleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StyleArgs | StyleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StyleState | undefined;
            resourceInputs["align"] = state ? state.align : undefined;
            resourceInputs["bgColor"] = state ? state.bgColor : undefined;
            resourceInputs["borderBottom"] = state ? state.borderBottom : undefined;
            resourceInputs["borderLeft"] = state ? state.borderLeft : undefined;
            resourceInputs["borderRight"] = state ? state.borderRight : undefined;
            resourceInputs["borderTop"] = state ? state.borderTop : undefined;
            resourceInputs["columnGap"] = state ? state.columnGap : undefined;
            resourceInputs["columnSpan"] = state ? state.columnSpan : undefined;
            resourceInputs["fgColor"] = state ? state.fgColor : undefined;
            resourceInputs["fontFamily"] = state ? state.fontFamily : undefined;
            resourceInputs["fontSize"] = state ? state.fontSize : undefined;
            resourceInputs["fontStyle"] = state ? state.fontStyle : undefined;
            resourceInputs["fontWeight"] = state ? state.fontWeight : undefined;
            resourceInputs["height"] = state ? state.height : undefined;
            resourceInputs["lineHeight"] = state ? state.lineHeight : undefined;
            resourceInputs["marginBottom"] = state ? state.marginBottom : undefined;
            resourceInputs["marginLeft"] = state ? state.marginLeft : undefined;
            resourceInputs["marginRight"] = state ? state.marginRight : undefined;
            resourceInputs["marginTop"] = state ? state.marginTop : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["paddingBottom"] = state ? state.paddingBottom : undefined;
            resourceInputs["paddingLeft"] = state ? state.paddingLeft : undefined;
            resourceInputs["paddingRight"] = state ? state.paddingRight : undefined;
            resourceInputs["paddingTop"] = state ? state.paddingTop : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["width"] = state ? state.width : undefined;
        } else {
            const args = argsOrState as StyleArgs | undefined;
            resourceInputs["align"] = args ? args.align : undefined;
            resourceInputs["bgColor"] = args ? args.bgColor : undefined;
            resourceInputs["borderBottom"] = args ? args.borderBottom : undefined;
            resourceInputs["borderLeft"] = args ? args.borderLeft : undefined;
            resourceInputs["borderRight"] = args ? args.borderRight : undefined;
            resourceInputs["borderTop"] = args ? args.borderTop : undefined;
            resourceInputs["columnGap"] = args ? args.columnGap : undefined;
            resourceInputs["columnSpan"] = args ? args.columnSpan : undefined;
            resourceInputs["fgColor"] = args ? args.fgColor : undefined;
            resourceInputs["fontFamily"] = args ? args.fontFamily : undefined;
            resourceInputs["fontSize"] = args ? args.fontSize : undefined;
            resourceInputs["fontStyle"] = args ? args.fontStyle : undefined;
            resourceInputs["fontWeight"] = args ? args.fontWeight : undefined;
            resourceInputs["height"] = args ? args.height : undefined;
            resourceInputs["lineHeight"] = args ? args.lineHeight : undefined;
            resourceInputs["marginBottom"] = args ? args.marginBottom : undefined;
            resourceInputs["marginLeft"] = args ? args.marginLeft : undefined;
            resourceInputs["marginRight"] = args ? args.marginRight : undefined;
            resourceInputs["marginTop"] = args ? args.marginTop : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["paddingBottom"] = args ? args.paddingBottom : undefined;
            resourceInputs["paddingLeft"] = args ? args.paddingLeft : undefined;
            resourceInputs["paddingRight"] = args ? args.paddingRight : undefined;
            resourceInputs["paddingTop"] = args ? args.paddingTop : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["width"] = args ? args.width : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Style.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Style resources.
 */
export interface StyleState {
    /**
     * Alignment. Valid values: `left`, `center`, `right`, `justify`.
     */
    align?: pulumi.Input<string>;
    /**
     * Background color.
     */
    bgColor?: pulumi.Input<string>;
    /**
     * Border bottom.
     */
    borderBottom?: pulumi.Input<string>;
    /**
     * Border left.
     */
    borderLeft?: pulumi.Input<string>;
    /**
     * Border right.
     */
    borderRight?: pulumi.Input<string>;
    /**
     * Border top.
     */
    borderTop?: pulumi.Input<string>;
    /**
     * Column gap.
     */
    columnGap?: pulumi.Input<string>;
    /**
     * Column span. Valid values: `none`, `all`.
     */
    columnSpan?: pulumi.Input<string>;
    /**
     * Foreground color.
     */
    fgColor?: pulumi.Input<string>;
    /**
     * Font family. Valid values: `Verdana`, `Arial`, `Helvetica`, `Courier`, `Times`.
     */
    fontFamily?: pulumi.Input<string>;
    /**
     * Font size.
     */
    fontSize?: pulumi.Input<string>;
    /**
     * Font style. Valid values: `normal`, `italic`.
     */
    fontStyle?: pulumi.Input<string>;
    /**
     * Font weight. Valid values: `normal`, `bold`.
     */
    fontWeight?: pulumi.Input<string>;
    /**
     * Height.
     */
    height?: pulumi.Input<string>;
    /**
     * Text line height.
     */
    lineHeight?: pulumi.Input<string>;
    /**
     * Margin bottom.
     */
    marginBottom?: pulumi.Input<string>;
    /**
     * Margin left.
     */
    marginLeft?: pulumi.Input<string>;
    /**
     * Margin right.
     */
    marginRight?: pulumi.Input<string>;
    /**
     * Margin top.
     */
    marginTop?: pulumi.Input<string>;
    /**
     * Report style name.
     */
    name?: pulumi.Input<string>;
    /**
     * Report style options. Valid values: `font`, `text`, `color`, `align`, `size`, `margin`, `border`, `padding`, `column`.
     */
    options?: pulumi.Input<string>;
    /**
     * Padding bottom.
     */
    paddingBottom?: pulumi.Input<string>;
    /**
     * Padding left.
     */
    paddingLeft?: pulumi.Input<string>;
    /**
     * Padding right.
     */
    paddingRight?: pulumi.Input<string>;
    /**
     * Padding top.
     */
    paddingTop?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Width.
     */
    width?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Style resource.
 */
export interface StyleArgs {
    /**
     * Alignment. Valid values: `left`, `center`, `right`, `justify`.
     */
    align?: pulumi.Input<string>;
    /**
     * Background color.
     */
    bgColor?: pulumi.Input<string>;
    /**
     * Border bottom.
     */
    borderBottom?: pulumi.Input<string>;
    /**
     * Border left.
     */
    borderLeft?: pulumi.Input<string>;
    /**
     * Border right.
     */
    borderRight?: pulumi.Input<string>;
    /**
     * Border top.
     */
    borderTop?: pulumi.Input<string>;
    /**
     * Column gap.
     */
    columnGap?: pulumi.Input<string>;
    /**
     * Column span. Valid values: `none`, `all`.
     */
    columnSpan?: pulumi.Input<string>;
    /**
     * Foreground color.
     */
    fgColor?: pulumi.Input<string>;
    /**
     * Font family. Valid values: `Verdana`, `Arial`, `Helvetica`, `Courier`, `Times`.
     */
    fontFamily?: pulumi.Input<string>;
    /**
     * Font size.
     */
    fontSize?: pulumi.Input<string>;
    /**
     * Font style. Valid values: `normal`, `italic`.
     */
    fontStyle?: pulumi.Input<string>;
    /**
     * Font weight. Valid values: `normal`, `bold`.
     */
    fontWeight?: pulumi.Input<string>;
    /**
     * Height.
     */
    height?: pulumi.Input<string>;
    /**
     * Text line height.
     */
    lineHeight?: pulumi.Input<string>;
    /**
     * Margin bottom.
     */
    marginBottom?: pulumi.Input<string>;
    /**
     * Margin left.
     */
    marginLeft?: pulumi.Input<string>;
    /**
     * Margin right.
     */
    marginRight?: pulumi.Input<string>;
    /**
     * Margin top.
     */
    marginTop?: pulumi.Input<string>;
    /**
     * Report style name.
     */
    name?: pulumi.Input<string>;
    /**
     * Report style options. Valid values: `font`, `text`, `color`, `align`, `size`, `margin`, `border`, `padding`, `column`.
     */
    options?: pulumi.Input<string>;
    /**
     * Padding bottom.
     */
    paddingBottom?: pulumi.Input<string>;
    /**
     * Padding left.
     */
    paddingLeft?: pulumi.Input<string>;
    /**
     * Padding right.
     */
    paddingRight?: pulumi.Input<string>;
    /**
     * Padding top.
     */
    paddingTop?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Width.
     */
    width?: pulumi.Input<string>;
}
