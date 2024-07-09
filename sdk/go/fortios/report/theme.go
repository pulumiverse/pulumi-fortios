// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package report

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Report themes configuration Applies to FortiOS Version `<= 7.0.0`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/report"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := report.NewTheme(ctx, "trname", &report.ThemeArgs{
//				ColumnCount:     pulumi.String("1"),
//				GraphChartStyle: pulumi.String("PS"),
//				PageOrient:      pulumi.String("portrait"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Report Theme can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:report/theme:Theme labelname {{name}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:report/theme:Theme labelname {{name}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Theme struct {
	pulumi.CustomResourceState

	// Bullet list style.
	BulletListStyle pulumi.StringOutput `pulumi:"bulletListStyle"`
	// Report page column count. Valid values: `1`, `2`, `3`.
	ColumnCount pulumi.StringOutput `pulumi:"columnCount"`
	// Default HTML report style.
	DefaultHtmlStyle pulumi.StringOutput `pulumi:"defaultHtmlStyle"`
	// Default PDF report style.
	DefaultPdfStyle pulumi.StringOutput `pulumi:"defaultPdfStyle"`
	// Graph chart style.
	GraphChartStyle pulumi.StringOutput `pulumi:"graphChartStyle"`
	// Report heading style.
	Heading1Style pulumi.StringOutput `pulumi:"heading1Style"`
	// Report heading style.
	Heading2Style pulumi.StringOutput `pulumi:"heading2Style"`
	// Report heading style.
	Heading3Style pulumi.StringOutput `pulumi:"heading3Style"`
	// Report heading style.
	Heading4Style pulumi.StringOutput `pulumi:"heading4Style"`
	// Horizontal line style.
	HlineStyle pulumi.StringOutput `pulumi:"hlineStyle"`
	// Image style.
	ImageStyle pulumi.StringOutput `pulumi:"imageStyle"`
	// Report theme name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Normal text style.
	NormalTextStyle pulumi.StringOutput `pulumi:"normalTextStyle"`
	// Numbered list style.
	NumberedListStyle pulumi.StringOutput `pulumi:"numberedListStyle"`
	// Report page footer style.
	PageFooterStyle pulumi.StringOutput `pulumi:"pageFooterStyle"`
	// Report page header style.
	PageHeaderStyle pulumi.StringOutput `pulumi:"pageHeaderStyle"`
	// Report page orientation. Valid values: `portrait`, `landscape`.
	PageOrient pulumi.StringOutput `pulumi:"pageOrient"`
	// Report page style.
	PageStyle pulumi.StringOutput `pulumi:"pageStyle"`
	// Report subtitle style.
	ReportSubtitleStyle pulumi.StringOutput `pulumi:"reportSubtitleStyle"`
	// Report title style.
	ReportTitleStyle pulumi.StringOutput `pulumi:"reportTitleStyle"`
	// Table chart caption style.
	TableChartCaptionStyle pulumi.StringOutput `pulumi:"tableChartCaptionStyle"`
	// Table chart even row style.
	TableChartEvenRowStyle pulumi.StringOutput `pulumi:"tableChartEvenRowStyle"`
	// Table chart head row style.
	TableChartHeadStyle pulumi.StringOutput `pulumi:"tableChartHeadStyle"`
	// Table chart odd row style.
	TableChartOddRowStyle pulumi.StringOutput `pulumi:"tableChartOddRowStyle"`
	// Table chart style.
	TableChartStyle pulumi.StringOutput `pulumi:"tableChartStyle"`
	// Table of contents heading style.
	TocHeading1Style pulumi.StringOutput `pulumi:"tocHeading1Style"`
	// Table of contents heading style.
	TocHeading2Style pulumi.StringOutput `pulumi:"tocHeading2Style"`
	// Table of contents heading style.
	TocHeading3Style pulumi.StringOutput `pulumi:"tocHeading3Style"`
	// Table of contents heading style.
	TocHeading4Style pulumi.StringOutput `pulumi:"tocHeading4Style"`
	// Table of contents title style.
	TocTitleStyle pulumi.StringOutput `pulumi:"tocTitleStyle"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewTheme registers a new resource with the given unique name, arguments, and options.
func NewTheme(ctx *pulumi.Context,
	name string, args *ThemeArgs, opts ...pulumi.ResourceOption) (*Theme, error) {
	if args == nil {
		args = &ThemeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Theme
	err := ctx.RegisterResource("fortios:report/theme:Theme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTheme gets an existing Theme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTheme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThemeState, opts ...pulumi.ResourceOption) (*Theme, error) {
	var resource Theme
	err := ctx.ReadResource("fortios:report/theme:Theme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Theme resources.
type themeState struct {
	// Bullet list style.
	BulletListStyle *string `pulumi:"bulletListStyle"`
	// Report page column count. Valid values: `1`, `2`, `3`.
	ColumnCount *string `pulumi:"columnCount"`
	// Default HTML report style.
	DefaultHtmlStyle *string `pulumi:"defaultHtmlStyle"`
	// Default PDF report style.
	DefaultPdfStyle *string `pulumi:"defaultPdfStyle"`
	// Graph chart style.
	GraphChartStyle *string `pulumi:"graphChartStyle"`
	// Report heading style.
	Heading1Style *string `pulumi:"heading1Style"`
	// Report heading style.
	Heading2Style *string `pulumi:"heading2Style"`
	// Report heading style.
	Heading3Style *string `pulumi:"heading3Style"`
	// Report heading style.
	Heading4Style *string `pulumi:"heading4Style"`
	// Horizontal line style.
	HlineStyle *string `pulumi:"hlineStyle"`
	// Image style.
	ImageStyle *string `pulumi:"imageStyle"`
	// Report theme name.
	Name *string `pulumi:"name"`
	// Normal text style.
	NormalTextStyle *string `pulumi:"normalTextStyle"`
	// Numbered list style.
	NumberedListStyle *string `pulumi:"numberedListStyle"`
	// Report page footer style.
	PageFooterStyle *string `pulumi:"pageFooterStyle"`
	// Report page header style.
	PageHeaderStyle *string `pulumi:"pageHeaderStyle"`
	// Report page orientation. Valid values: `portrait`, `landscape`.
	PageOrient *string `pulumi:"pageOrient"`
	// Report page style.
	PageStyle *string `pulumi:"pageStyle"`
	// Report subtitle style.
	ReportSubtitleStyle *string `pulumi:"reportSubtitleStyle"`
	// Report title style.
	ReportTitleStyle *string `pulumi:"reportTitleStyle"`
	// Table chart caption style.
	TableChartCaptionStyle *string `pulumi:"tableChartCaptionStyle"`
	// Table chart even row style.
	TableChartEvenRowStyle *string `pulumi:"tableChartEvenRowStyle"`
	// Table chart head row style.
	TableChartHeadStyle *string `pulumi:"tableChartHeadStyle"`
	// Table chart odd row style.
	TableChartOddRowStyle *string `pulumi:"tableChartOddRowStyle"`
	// Table chart style.
	TableChartStyle *string `pulumi:"tableChartStyle"`
	// Table of contents heading style.
	TocHeading1Style *string `pulumi:"tocHeading1Style"`
	// Table of contents heading style.
	TocHeading2Style *string `pulumi:"tocHeading2Style"`
	// Table of contents heading style.
	TocHeading3Style *string `pulumi:"tocHeading3Style"`
	// Table of contents heading style.
	TocHeading4Style *string `pulumi:"tocHeading4Style"`
	// Table of contents title style.
	TocTitleStyle *string `pulumi:"tocTitleStyle"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ThemeState struct {
	// Bullet list style.
	BulletListStyle pulumi.StringPtrInput
	// Report page column count. Valid values: `1`, `2`, `3`.
	ColumnCount pulumi.StringPtrInput
	// Default HTML report style.
	DefaultHtmlStyle pulumi.StringPtrInput
	// Default PDF report style.
	DefaultPdfStyle pulumi.StringPtrInput
	// Graph chart style.
	GraphChartStyle pulumi.StringPtrInput
	// Report heading style.
	Heading1Style pulumi.StringPtrInput
	// Report heading style.
	Heading2Style pulumi.StringPtrInput
	// Report heading style.
	Heading3Style pulumi.StringPtrInput
	// Report heading style.
	Heading4Style pulumi.StringPtrInput
	// Horizontal line style.
	HlineStyle pulumi.StringPtrInput
	// Image style.
	ImageStyle pulumi.StringPtrInput
	// Report theme name.
	Name pulumi.StringPtrInput
	// Normal text style.
	NormalTextStyle pulumi.StringPtrInput
	// Numbered list style.
	NumberedListStyle pulumi.StringPtrInput
	// Report page footer style.
	PageFooterStyle pulumi.StringPtrInput
	// Report page header style.
	PageHeaderStyle pulumi.StringPtrInput
	// Report page orientation. Valid values: `portrait`, `landscape`.
	PageOrient pulumi.StringPtrInput
	// Report page style.
	PageStyle pulumi.StringPtrInput
	// Report subtitle style.
	ReportSubtitleStyle pulumi.StringPtrInput
	// Report title style.
	ReportTitleStyle pulumi.StringPtrInput
	// Table chart caption style.
	TableChartCaptionStyle pulumi.StringPtrInput
	// Table chart even row style.
	TableChartEvenRowStyle pulumi.StringPtrInput
	// Table chart head row style.
	TableChartHeadStyle pulumi.StringPtrInput
	// Table chart odd row style.
	TableChartOddRowStyle pulumi.StringPtrInput
	// Table chart style.
	TableChartStyle pulumi.StringPtrInput
	// Table of contents heading style.
	TocHeading1Style pulumi.StringPtrInput
	// Table of contents heading style.
	TocHeading2Style pulumi.StringPtrInput
	// Table of contents heading style.
	TocHeading3Style pulumi.StringPtrInput
	// Table of contents heading style.
	TocHeading4Style pulumi.StringPtrInput
	// Table of contents title style.
	TocTitleStyle pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ThemeState) ElementType() reflect.Type {
	return reflect.TypeOf((*themeState)(nil)).Elem()
}

type themeArgs struct {
	// Bullet list style.
	BulletListStyle *string `pulumi:"bulletListStyle"`
	// Report page column count. Valid values: `1`, `2`, `3`.
	ColumnCount *string `pulumi:"columnCount"`
	// Default HTML report style.
	DefaultHtmlStyle *string `pulumi:"defaultHtmlStyle"`
	// Default PDF report style.
	DefaultPdfStyle *string `pulumi:"defaultPdfStyle"`
	// Graph chart style.
	GraphChartStyle *string `pulumi:"graphChartStyle"`
	// Report heading style.
	Heading1Style *string `pulumi:"heading1Style"`
	// Report heading style.
	Heading2Style *string `pulumi:"heading2Style"`
	// Report heading style.
	Heading3Style *string `pulumi:"heading3Style"`
	// Report heading style.
	Heading4Style *string `pulumi:"heading4Style"`
	// Horizontal line style.
	HlineStyle *string `pulumi:"hlineStyle"`
	// Image style.
	ImageStyle *string `pulumi:"imageStyle"`
	// Report theme name.
	Name *string `pulumi:"name"`
	// Normal text style.
	NormalTextStyle *string `pulumi:"normalTextStyle"`
	// Numbered list style.
	NumberedListStyle *string `pulumi:"numberedListStyle"`
	// Report page footer style.
	PageFooterStyle *string `pulumi:"pageFooterStyle"`
	// Report page header style.
	PageHeaderStyle *string `pulumi:"pageHeaderStyle"`
	// Report page orientation. Valid values: `portrait`, `landscape`.
	PageOrient *string `pulumi:"pageOrient"`
	// Report page style.
	PageStyle *string `pulumi:"pageStyle"`
	// Report subtitle style.
	ReportSubtitleStyle *string `pulumi:"reportSubtitleStyle"`
	// Report title style.
	ReportTitleStyle *string `pulumi:"reportTitleStyle"`
	// Table chart caption style.
	TableChartCaptionStyle *string `pulumi:"tableChartCaptionStyle"`
	// Table chart even row style.
	TableChartEvenRowStyle *string `pulumi:"tableChartEvenRowStyle"`
	// Table chart head row style.
	TableChartHeadStyle *string `pulumi:"tableChartHeadStyle"`
	// Table chart odd row style.
	TableChartOddRowStyle *string `pulumi:"tableChartOddRowStyle"`
	// Table chart style.
	TableChartStyle *string `pulumi:"tableChartStyle"`
	// Table of contents heading style.
	TocHeading1Style *string `pulumi:"tocHeading1Style"`
	// Table of contents heading style.
	TocHeading2Style *string `pulumi:"tocHeading2Style"`
	// Table of contents heading style.
	TocHeading3Style *string `pulumi:"tocHeading3Style"`
	// Table of contents heading style.
	TocHeading4Style *string `pulumi:"tocHeading4Style"`
	// Table of contents title style.
	TocTitleStyle *string `pulumi:"tocTitleStyle"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Theme resource.
type ThemeArgs struct {
	// Bullet list style.
	BulletListStyle pulumi.StringPtrInput
	// Report page column count. Valid values: `1`, `2`, `3`.
	ColumnCount pulumi.StringPtrInput
	// Default HTML report style.
	DefaultHtmlStyle pulumi.StringPtrInput
	// Default PDF report style.
	DefaultPdfStyle pulumi.StringPtrInput
	// Graph chart style.
	GraphChartStyle pulumi.StringPtrInput
	// Report heading style.
	Heading1Style pulumi.StringPtrInput
	// Report heading style.
	Heading2Style pulumi.StringPtrInput
	// Report heading style.
	Heading3Style pulumi.StringPtrInput
	// Report heading style.
	Heading4Style pulumi.StringPtrInput
	// Horizontal line style.
	HlineStyle pulumi.StringPtrInput
	// Image style.
	ImageStyle pulumi.StringPtrInput
	// Report theme name.
	Name pulumi.StringPtrInput
	// Normal text style.
	NormalTextStyle pulumi.StringPtrInput
	// Numbered list style.
	NumberedListStyle pulumi.StringPtrInput
	// Report page footer style.
	PageFooterStyle pulumi.StringPtrInput
	// Report page header style.
	PageHeaderStyle pulumi.StringPtrInput
	// Report page orientation. Valid values: `portrait`, `landscape`.
	PageOrient pulumi.StringPtrInput
	// Report page style.
	PageStyle pulumi.StringPtrInput
	// Report subtitle style.
	ReportSubtitleStyle pulumi.StringPtrInput
	// Report title style.
	ReportTitleStyle pulumi.StringPtrInput
	// Table chart caption style.
	TableChartCaptionStyle pulumi.StringPtrInput
	// Table chart even row style.
	TableChartEvenRowStyle pulumi.StringPtrInput
	// Table chart head row style.
	TableChartHeadStyle pulumi.StringPtrInput
	// Table chart odd row style.
	TableChartOddRowStyle pulumi.StringPtrInput
	// Table chart style.
	TableChartStyle pulumi.StringPtrInput
	// Table of contents heading style.
	TocHeading1Style pulumi.StringPtrInput
	// Table of contents heading style.
	TocHeading2Style pulumi.StringPtrInput
	// Table of contents heading style.
	TocHeading3Style pulumi.StringPtrInput
	// Table of contents heading style.
	TocHeading4Style pulumi.StringPtrInput
	// Table of contents title style.
	TocTitleStyle pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ThemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*themeArgs)(nil)).Elem()
}

type ThemeInput interface {
	pulumi.Input

	ToThemeOutput() ThemeOutput
	ToThemeOutputWithContext(ctx context.Context) ThemeOutput
}

func (*Theme) ElementType() reflect.Type {
	return reflect.TypeOf((**Theme)(nil)).Elem()
}

func (i *Theme) ToThemeOutput() ThemeOutput {
	return i.ToThemeOutputWithContext(context.Background())
}

func (i *Theme) ToThemeOutputWithContext(ctx context.Context) ThemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThemeOutput)
}

// ThemeArrayInput is an input type that accepts ThemeArray and ThemeArrayOutput values.
// You can construct a concrete instance of `ThemeArrayInput` via:
//
//	ThemeArray{ ThemeArgs{...} }
type ThemeArrayInput interface {
	pulumi.Input

	ToThemeArrayOutput() ThemeArrayOutput
	ToThemeArrayOutputWithContext(context.Context) ThemeArrayOutput
}

type ThemeArray []ThemeInput

func (ThemeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Theme)(nil)).Elem()
}

func (i ThemeArray) ToThemeArrayOutput() ThemeArrayOutput {
	return i.ToThemeArrayOutputWithContext(context.Background())
}

func (i ThemeArray) ToThemeArrayOutputWithContext(ctx context.Context) ThemeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThemeArrayOutput)
}

// ThemeMapInput is an input type that accepts ThemeMap and ThemeMapOutput values.
// You can construct a concrete instance of `ThemeMapInput` via:
//
//	ThemeMap{ "key": ThemeArgs{...} }
type ThemeMapInput interface {
	pulumi.Input

	ToThemeMapOutput() ThemeMapOutput
	ToThemeMapOutputWithContext(context.Context) ThemeMapOutput
}

type ThemeMap map[string]ThemeInput

func (ThemeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Theme)(nil)).Elem()
}

func (i ThemeMap) ToThemeMapOutput() ThemeMapOutput {
	return i.ToThemeMapOutputWithContext(context.Background())
}

func (i ThemeMap) ToThemeMapOutputWithContext(ctx context.Context) ThemeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThemeMapOutput)
}

type ThemeOutput struct{ *pulumi.OutputState }

func (ThemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Theme)(nil)).Elem()
}

func (o ThemeOutput) ToThemeOutput() ThemeOutput {
	return o
}

func (o ThemeOutput) ToThemeOutputWithContext(ctx context.Context) ThemeOutput {
	return o
}

// Bullet list style.
func (o ThemeOutput) BulletListStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.BulletListStyle }).(pulumi.StringOutput)
}

// Report page column count. Valid values: `1`, `2`, `3`.
func (o ThemeOutput) ColumnCount() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.ColumnCount }).(pulumi.StringOutput)
}

// Default HTML report style.
func (o ThemeOutput) DefaultHtmlStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.DefaultHtmlStyle }).(pulumi.StringOutput)
}

// Default PDF report style.
func (o ThemeOutput) DefaultPdfStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.DefaultPdfStyle }).(pulumi.StringOutput)
}

// Graph chart style.
func (o ThemeOutput) GraphChartStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.GraphChartStyle }).(pulumi.StringOutput)
}

// Report heading style.
func (o ThemeOutput) Heading1Style() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.Heading1Style }).(pulumi.StringOutput)
}

// Report heading style.
func (o ThemeOutput) Heading2Style() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.Heading2Style }).(pulumi.StringOutput)
}

// Report heading style.
func (o ThemeOutput) Heading3Style() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.Heading3Style }).(pulumi.StringOutput)
}

// Report heading style.
func (o ThemeOutput) Heading4Style() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.Heading4Style }).(pulumi.StringOutput)
}

// Horizontal line style.
func (o ThemeOutput) HlineStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.HlineStyle }).(pulumi.StringOutput)
}

// Image style.
func (o ThemeOutput) ImageStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.ImageStyle }).(pulumi.StringOutput)
}

// Report theme name.
func (o ThemeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Normal text style.
func (o ThemeOutput) NormalTextStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.NormalTextStyle }).(pulumi.StringOutput)
}

// Numbered list style.
func (o ThemeOutput) NumberedListStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.NumberedListStyle }).(pulumi.StringOutput)
}

// Report page footer style.
func (o ThemeOutput) PageFooterStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.PageFooterStyle }).(pulumi.StringOutput)
}

// Report page header style.
func (o ThemeOutput) PageHeaderStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.PageHeaderStyle }).(pulumi.StringOutput)
}

// Report page orientation. Valid values: `portrait`, `landscape`.
func (o ThemeOutput) PageOrient() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.PageOrient }).(pulumi.StringOutput)
}

// Report page style.
func (o ThemeOutput) PageStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.PageStyle }).(pulumi.StringOutput)
}

// Report subtitle style.
func (o ThemeOutput) ReportSubtitleStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.ReportSubtitleStyle }).(pulumi.StringOutput)
}

// Report title style.
func (o ThemeOutput) ReportTitleStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.ReportTitleStyle }).(pulumi.StringOutput)
}

// Table chart caption style.
func (o ThemeOutput) TableChartCaptionStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TableChartCaptionStyle }).(pulumi.StringOutput)
}

// Table chart even row style.
func (o ThemeOutput) TableChartEvenRowStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TableChartEvenRowStyle }).(pulumi.StringOutput)
}

// Table chart head row style.
func (o ThemeOutput) TableChartHeadStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TableChartHeadStyle }).(pulumi.StringOutput)
}

// Table chart odd row style.
func (o ThemeOutput) TableChartOddRowStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TableChartOddRowStyle }).(pulumi.StringOutput)
}

// Table chart style.
func (o ThemeOutput) TableChartStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TableChartStyle }).(pulumi.StringOutput)
}

// Table of contents heading style.
func (o ThemeOutput) TocHeading1Style() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TocHeading1Style }).(pulumi.StringOutput)
}

// Table of contents heading style.
func (o ThemeOutput) TocHeading2Style() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TocHeading2Style }).(pulumi.StringOutput)
}

// Table of contents heading style.
func (o ThemeOutput) TocHeading3Style() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TocHeading3Style }).(pulumi.StringOutput)
}

// Table of contents heading style.
func (o ThemeOutput) TocHeading4Style() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TocHeading4Style }).(pulumi.StringOutput)
}

// Table of contents title style.
func (o ThemeOutput) TocTitleStyle() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.TocTitleStyle }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ThemeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ThemeArrayOutput struct{ *pulumi.OutputState }

func (ThemeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Theme)(nil)).Elem()
}

func (o ThemeArrayOutput) ToThemeArrayOutput() ThemeArrayOutput {
	return o
}

func (o ThemeArrayOutput) ToThemeArrayOutputWithContext(ctx context.Context) ThemeArrayOutput {
	return o
}

func (o ThemeArrayOutput) Index(i pulumi.IntInput) ThemeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Theme {
		return vs[0].([]*Theme)[vs[1].(int)]
	}).(ThemeOutput)
}

type ThemeMapOutput struct{ *pulumi.OutputState }

func (ThemeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Theme)(nil)).Elem()
}

func (o ThemeMapOutput) ToThemeMapOutput() ThemeMapOutput {
	return o
}

func (o ThemeMapOutput) ToThemeMapOutputWithContext(ctx context.Context) ThemeMapOutput {
	return o
}

func (o ThemeMapOutput) MapIndex(k pulumi.StringInput) ThemeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Theme {
		return vs[0].(map[string]*Theme)[vs[1].(string)]
	}).(ThemeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ThemeInput)(nil)).Elem(), &Theme{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThemeArrayInput)(nil)).Elem(), ThemeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThemeMapInput)(nil)).Elem(), ThemeMap{})
	pulumi.RegisterOutputType(ThemeOutput{})
	pulumi.RegisterOutputType(ThemeArrayOutput{})
	pulumi.RegisterOutputType(ThemeMapOutput{})
}
