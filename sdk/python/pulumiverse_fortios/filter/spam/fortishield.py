# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['FortishieldArgs', 'Fortishield']

@pulumi.input_type
class FortishieldArgs:
    def __init__(__self__, *,
                 spam_submit_force: Optional[pulumi.Input[str]] = None,
                 spam_submit_srv: Optional[pulumi.Input[str]] = None,
                 spam_submit_txt2htm: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Fortishield resource.
        :param pulumi.Input[str] spam_submit_force: Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] spam_submit_srv: Hostname of the spam submission server.
        :param pulumi.Input[str] spam_submit_txt2htm: Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if spam_submit_force is not None:
            pulumi.set(__self__, "spam_submit_force", spam_submit_force)
        if spam_submit_srv is not None:
            pulumi.set(__self__, "spam_submit_srv", spam_submit_srv)
        if spam_submit_txt2htm is not None:
            pulumi.set(__self__, "spam_submit_txt2htm", spam_submit_txt2htm)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="spamSubmitForce")
    def spam_submit_force(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "spam_submit_force")

    @spam_submit_force.setter
    def spam_submit_force(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spam_submit_force", value)

    @property
    @pulumi.getter(name="spamSubmitSrv")
    def spam_submit_srv(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the spam submission server.
        """
        return pulumi.get(self, "spam_submit_srv")

    @spam_submit_srv.setter
    def spam_submit_srv(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spam_submit_srv", value)

    @property
    @pulumi.getter(name="spamSubmitTxt2htm")
    def spam_submit_txt2htm(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "spam_submit_txt2htm")

    @spam_submit_txt2htm.setter
    def spam_submit_txt2htm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spam_submit_txt2htm", value)

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


@pulumi.input_type
class _FortishieldState:
    def __init__(__self__, *,
                 spam_submit_force: Optional[pulumi.Input[str]] = None,
                 spam_submit_srv: Optional[pulumi.Input[str]] = None,
                 spam_submit_txt2htm: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Fortishield resources.
        :param pulumi.Input[str] spam_submit_force: Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] spam_submit_srv: Hostname of the spam submission server.
        :param pulumi.Input[str] spam_submit_txt2htm: Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if spam_submit_force is not None:
            pulumi.set(__self__, "spam_submit_force", spam_submit_force)
        if spam_submit_srv is not None:
            pulumi.set(__self__, "spam_submit_srv", spam_submit_srv)
        if spam_submit_txt2htm is not None:
            pulumi.set(__self__, "spam_submit_txt2htm", spam_submit_txt2htm)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="spamSubmitForce")
    def spam_submit_force(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "spam_submit_force")

    @spam_submit_force.setter
    def spam_submit_force(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spam_submit_force", value)

    @property
    @pulumi.getter(name="spamSubmitSrv")
    def spam_submit_srv(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the spam submission server.
        """
        return pulumi.get(self, "spam_submit_srv")

    @spam_submit_srv.setter
    def spam_submit_srv(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spam_submit_srv", value)

    @property
    @pulumi.getter(name="spamSubmitTxt2htm")
    def spam_submit_txt2htm(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "spam_submit_txt2htm")

    @spam_submit_txt2htm.setter
    def spam_submit_txt2htm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spam_submit_txt2htm", value)

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


class Fortishield(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 spam_submit_force: Optional[pulumi.Input[str]] = None,
                 spam_submit_srv: Optional[pulumi.Input[str]] = None,
                 spam_submit_txt2htm: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiGuard - AntiSpam. Applies to FortiOS Version `<= 6.2.0`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.filter.spam.Fortishield("trname",
            spam_submit_force="enable",
            spam_submit_srv="www.nospammer.net",
            spam_submit_txt2htm="enable")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Spamfilter Fortishield can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:filter/spam/fortishield:Fortishield labelname SpamfilterFortishield
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:filter/spam/fortishield:Fortishield labelname SpamfilterFortishield
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] spam_submit_force: Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] spam_submit_srv: Hostname of the spam submission server.
        :param pulumi.Input[str] spam_submit_txt2htm: Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortishieldArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiGuard - AntiSpam. Applies to FortiOS Version `<= 6.2.0`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.filter.spam.Fortishield("trname",
            spam_submit_force="enable",
            spam_submit_srv="www.nospammer.net",
            spam_submit_txt2htm="enable")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Spamfilter Fortishield can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:filter/spam/fortishield:Fortishield labelname SpamfilterFortishield
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:filter/spam/fortishield:Fortishield labelname SpamfilterFortishield
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FortishieldArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortishieldArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 spam_submit_force: Optional[pulumi.Input[str]] = None,
                 spam_submit_srv: Optional[pulumi.Input[str]] = None,
                 spam_submit_txt2htm: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortishieldArgs.__new__(FortishieldArgs)

            __props__.__dict__["spam_submit_force"] = spam_submit_force
            __props__.__dict__["spam_submit_srv"] = spam_submit_srv
            __props__.__dict__["spam_submit_txt2htm"] = spam_submit_txt2htm
            __props__.__dict__["vdomparam"] = vdomparam
        super(Fortishield, __self__).__init__(
            'fortios:filter/spam/fortishield:Fortishield',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            spam_submit_force: Optional[pulumi.Input[str]] = None,
            spam_submit_srv: Optional[pulumi.Input[str]] = None,
            spam_submit_txt2htm: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Fortishield':
        """
        Get an existing Fortishield resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] spam_submit_force: Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] spam_submit_srv: Hostname of the spam submission server.
        :param pulumi.Input[str] spam_submit_txt2htm: Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortishieldState.__new__(_FortishieldState)

        __props__.__dict__["spam_submit_force"] = spam_submit_force
        __props__.__dict__["spam_submit_srv"] = spam_submit_srv
        __props__.__dict__["spam_submit_txt2htm"] = spam_submit_txt2htm
        __props__.__dict__["vdomparam"] = vdomparam
        return Fortishield(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="spamSubmitForce")
    def spam_submit_force(self) -> pulumi.Output[str]:
        """
        Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "spam_submit_force")

    @property
    @pulumi.getter(name="spamSubmitSrv")
    def spam_submit_srv(self) -> pulumi.Output[str]:
        """
        Hostname of the spam submission server.
        """
        return pulumi.get(self, "spam_submit_srv")

    @property
    @pulumi.getter(name="spamSubmitTxt2htm")
    def spam_submit_txt2htm(self) -> pulumi.Output[str]:
        """
        Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "spam_submit_txt2htm")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

