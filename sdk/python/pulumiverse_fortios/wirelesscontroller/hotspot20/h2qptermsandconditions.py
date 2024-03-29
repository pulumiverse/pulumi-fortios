# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['H2qptermsandconditionsArgs', 'H2qptermsandconditions']

@pulumi.input_type
class H2qptermsandconditionsArgs:
    def __init__(__self__, *,
                 filename: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timestamp: Optional[pulumi.Input[int]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a H2qptermsandconditions resource.
        :param pulumi.Input[str] filename: Filename.
        :param pulumi.Input[str] name: Terms and Conditions ID.
        :param pulumi.Input[int] timestamp: Timestamp.
        :param pulumi.Input[str] url: URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if filename is not None:
            pulumi.set(__self__, "filename", filename)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if timestamp is not None:
            pulumi.set(__self__, "timestamp", timestamp)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[str]]:
        """
        Filename.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Terms and Conditions ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def timestamp(self) -> Optional[pulumi.Input[int]]:
        """
        Timestamp.
        """
        return pulumi.get(self, "timestamp")

    @timestamp.setter
    def timestamp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timestamp", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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
class _H2qptermsandconditionsState:
    def __init__(__self__, *,
                 filename: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timestamp: Optional[pulumi.Input[int]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering H2qptermsandconditions resources.
        :param pulumi.Input[str] filename: Filename.
        :param pulumi.Input[str] name: Terms and Conditions ID.
        :param pulumi.Input[int] timestamp: Timestamp.
        :param pulumi.Input[str] url: URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if filename is not None:
            pulumi.set(__self__, "filename", filename)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if timestamp is not None:
            pulumi.set(__self__, "timestamp", timestamp)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[str]]:
        """
        Filename.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Terms and Conditions ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def timestamp(self) -> Optional[pulumi.Input[int]]:
        """
        Timestamp.
        """
        return pulumi.get(self, "timestamp")

    @timestamp.setter
    def timestamp(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timestamp", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

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


class H2qptermsandconditions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timestamp: Optional[pulumi.Input[int]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure terms and conditions. Applies to FortiOS Version `>= 7.0.2`.

        ## Import

        WirelessControllerHotspot20 H2QpTermsAndConditions can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/h2qptermsandconditions:H2qptermsandconditions labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/h2qptermsandconditions:H2qptermsandconditions labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] filename: Filename.
        :param pulumi.Input[str] name: Terms and Conditions ID.
        :param pulumi.Input[int] timestamp: Timestamp.
        :param pulumi.Input[str] url: URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[H2qptermsandconditionsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure terms and conditions. Applies to FortiOS Version `>= 7.0.2`.

        ## Import

        WirelessControllerHotspot20 H2QpTermsAndConditions can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/h2qptermsandconditions:H2qptermsandconditions labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:wirelesscontroller/hotspot20/h2qptermsandconditions:H2qptermsandconditions labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param H2qptermsandconditionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(H2qptermsandconditionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timestamp: Optional[pulumi.Input[int]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = H2qptermsandconditionsArgs.__new__(H2qptermsandconditionsArgs)

            __props__.__dict__["filename"] = filename
            __props__.__dict__["name"] = name
            __props__.__dict__["timestamp"] = timestamp
            __props__.__dict__["url"] = url
            __props__.__dict__["vdomparam"] = vdomparam
        super(H2qptermsandconditions, __self__).__init__(
            'fortios:wirelesscontroller/hotspot20/h2qptermsandconditions:H2qptermsandconditions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            filename: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            timestamp: Optional[pulumi.Input[int]] = None,
            url: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'H2qptermsandconditions':
        """
        Get an existing H2qptermsandconditions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] filename: Filename.
        :param pulumi.Input[str] name: Terms and Conditions ID.
        :param pulumi.Input[int] timestamp: Timestamp.
        :param pulumi.Input[str] url: URL.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _H2qptermsandconditionsState.__new__(_H2qptermsandconditionsState)

        __props__.__dict__["filename"] = filename
        __props__.__dict__["name"] = name
        __props__.__dict__["timestamp"] = timestamp
        __props__.__dict__["url"] = url
        __props__.__dict__["vdomparam"] = vdomparam
        return H2qptermsandconditions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Output[str]:
        """
        Filename.
        """
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Terms and Conditions ID.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def timestamp(self) -> pulumi.Output[int]:
        """
        Timestamp.
        """
        return pulumi.get(self, "timestamp")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

