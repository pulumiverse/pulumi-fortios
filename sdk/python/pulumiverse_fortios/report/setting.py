# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SettingArgs', 'Setting']

@pulumi.input_type
class SettingArgs:
    def __init__(__self__, *,
                 fortiview: Optional[pulumi.Input[str]] = None,
                 pdf_report: Optional[pulumi.Input[str]] = None,
                 report_source: Optional[pulumi.Input[str]] = None,
                 top_n: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 web_browsing_threshold: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Setting resource.
        :param pulumi.Input[str] fortiview: Enable/disable historical FortiView. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pdf_report: Enable/disable PDF report. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] report_source: Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
        :param pulumi.Input[int] top_n: Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] web_browsing_threshold: Web browsing time calculation threshold (3 - 15 min).
        """
        if fortiview is not None:
            pulumi.set(__self__, "fortiview", fortiview)
        if pdf_report is not None:
            pulumi.set(__self__, "pdf_report", pdf_report)
        if report_source is not None:
            pulumi.set(__self__, "report_source", report_source)
        if top_n is not None:
            pulumi.set(__self__, "top_n", top_n)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if web_browsing_threshold is not None:
            pulumi.set(__self__, "web_browsing_threshold", web_browsing_threshold)

    @property
    @pulumi.getter
    def fortiview(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable historical FortiView. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiview")

    @fortiview.setter
    def fortiview(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiview", value)

    @property
    @pulumi.getter(name="pdfReport")
    def pdf_report(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable PDF report. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pdf_report")

    @pdf_report.setter
    def pdf_report(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pdf_report", value)

    @property
    @pulumi.getter(name="reportSource")
    def report_source(self) -> Optional[pulumi.Input[str]]:
        """
        Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
        """
        return pulumi.get(self, "report_source")

    @report_source.setter
    def report_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "report_source", value)

    @property
    @pulumi.getter(name="topN")
    def top_n(self) -> Optional[pulumi.Input[int]]:
        """
        Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
        """
        return pulumi.get(self, "top_n")

    @top_n.setter
    def top_n(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "top_n", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="webBrowsingThreshold")
    def web_browsing_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Web browsing time calculation threshold (3 - 15 min).
        """
        return pulumi.get(self, "web_browsing_threshold")

    @web_browsing_threshold.setter
    def web_browsing_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "web_browsing_threshold", value)


@pulumi.input_type
class _SettingState:
    def __init__(__self__, *,
                 fortiview: Optional[pulumi.Input[str]] = None,
                 pdf_report: Optional[pulumi.Input[str]] = None,
                 report_source: Optional[pulumi.Input[str]] = None,
                 top_n: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 web_browsing_threshold: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Setting resources.
        :param pulumi.Input[str] fortiview: Enable/disable historical FortiView. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pdf_report: Enable/disable PDF report. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] report_source: Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
        :param pulumi.Input[int] top_n: Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] web_browsing_threshold: Web browsing time calculation threshold (3 - 15 min).
        """
        if fortiview is not None:
            pulumi.set(__self__, "fortiview", fortiview)
        if pdf_report is not None:
            pulumi.set(__self__, "pdf_report", pdf_report)
        if report_source is not None:
            pulumi.set(__self__, "report_source", report_source)
        if top_n is not None:
            pulumi.set(__self__, "top_n", top_n)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if web_browsing_threshold is not None:
            pulumi.set(__self__, "web_browsing_threshold", web_browsing_threshold)

    @property
    @pulumi.getter
    def fortiview(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable historical FortiView. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiview")

    @fortiview.setter
    def fortiview(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiview", value)

    @property
    @pulumi.getter(name="pdfReport")
    def pdf_report(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable PDF report. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pdf_report")

    @pdf_report.setter
    def pdf_report(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pdf_report", value)

    @property
    @pulumi.getter(name="reportSource")
    def report_source(self) -> Optional[pulumi.Input[str]]:
        """
        Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
        """
        return pulumi.get(self, "report_source")

    @report_source.setter
    def report_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "report_source", value)

    @property
    @pulumi.getter(name="topN")
    def top_n(self) -> Optional[pulumi.Input[int]]:
        """
        Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
        """
        return pulumi.get(self, "top_n")

    @top_n.setter
    def top_n(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "top_n", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="webBrowsingThreshold")
    def web_browsing_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Web browsing time calculation threshold (3 - 15 min).
        """
        return pulumi.get(self, "web_browsing_threshold")

    @web_browsing_threshold.setter
    def web_browsing_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "web_browsing_threshold", value)


class Setting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fortiview: Optional[pulumi.Input[str]] = None,
                 pdf_report: Optional[pulumi.Input[str]] = None,
                 report_source: Optional[pulumi.Input[str]] = None,
                 top_n: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 web_browsing_threshold: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Report setting configuration.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.report.Setting("trname",
            fortiview="enable",
            pdf_report="enable",
            report_source="forward-traffic",
            top_n=1000,
            web_browsing_threshold=3)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Report Setting can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:report/setting:Setting labelname ReportSetting
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:report/setting:Setting labelname ReportSetting
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fortiview: Enable/disable historical FortiView. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pdf_report: Enable/disable PDF report. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] report_source: Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
        :param pulumi.Input[int] top_n: Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] web_browsing_threshold: Web browsing time calculation threshold (3 - 15 min).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Report setting configuration.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.report.Setting("trname",
            fortiview="enable",
            pdf_report="enable",
            report_source="forward-traffic",
            top_n=1000,
            web_browsing_threshold=3)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Report Setting can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:report/setting:Setting labelname ReportSetting
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:report/setting:Setting labelname ReportSetting
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fortiview: Optional[pulumi.Input[str]] = None,
                 pdf_report: Optional[pulumi.Input[str]] = None,
                 report_source: Optional[pulumi.Input[str]] = None,
                 top_n: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 web_browsing_threshold: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SettingArgs.__new__(SettingArgs)

            __props__.__dict__["fortiview"] = fortiview
            __props__.__dict__["pdf_report"] = pdf_report
            __props__.__dict__["report_source"] = report_source
            __props__.__dict__["top_n"] = top_n
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["web_browsing_threshold"] = web_browsing_threshold
        super(Setting, __self__).__init__(
            'fortios:report/setting:Setting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fortiview: Optional[pulumi.Input[str]] = None,
            pdf_report: Optional[pulumi.Input[str]] = None,
            report_source: Optional[pulumi.Input[str]] = None,
            top_n: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            web_browsing_threshold: Optional[pulumi.Input[int]] = None) -> 'Setting':
        """
        Get an existing Setting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fortiview: Enable/disable historical FortiView. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pdf_report: Enable/disable PDF report. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] report_source: Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
        :param pulumi.Input[int] top_n: Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] web_browsing_threshold: Web browsing time calculation threshold (3 - 15 min).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SettingState.__new__(_SettingState)

        __props__.__dict__["fortiview"] = fortiview
        __props__.__dict__["pdf_report"] = pdf_report
        __props__.__dict__["report_source"] = report_source
        __props__.__dict__["top_n"] = top_n
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["web_browsing_threshold"] = web_browsing_threshold
        return Setting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fortiview(self) -> pulumi.Output[str]:
        """
        Enable/disable historical FortiView. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiview")

    @property
    @pulumi.getter(name="pdfReport")
    def pdf_report(self) -> pulumi.Output[str]:
        """
        Enable/disable PDF report. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pdf_report")

    @property
    @pulumi.getter(name="reportSource")
    def report_source(self) -> pulumi.Output[str]:
        """
        Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
        """
        return pulumi.get(self, "report_source")

    @property
    @pulumi.getter(name="topN")
    def top_n(self) -> pulumi.Output[int]:
        """
        Number of items to populate. On FortiOS versions 6.2.0: 100 - 4000. On FortiOS versions >= 6.2.4: 1000 - 20000.
        """
        return pulumi.get(self, "top_n")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="webBrowsingThreshold")
    def web_browsing_threshold(self) -> pulumi.Output[int]:
        """
        Web browsing time calculation threshold (3 - 15 min).
        """
        return pulumi.get(self, "web_browsing_threshold")

