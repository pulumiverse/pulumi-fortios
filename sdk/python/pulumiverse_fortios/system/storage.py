# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StorageArgs', 'Storage']

@pulumi.input_type
class StorageArgs:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 media_status: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wanopt_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Storage resource.
        :param pulumi.Input[str] device: Partition device.
        :param pulumi.Input[str] media_status: The physical status of current media. Valid values: `enable`, `disable`, `fail`.
        :param pulumi.Input[str] name: Storage name.
        :param pulumi.Input[int] order: Set storage order.
        :param pulumi.Input[str] partition: Label of underlying partition.
        :param pulumi.Input[int] size: Partition size.
        :param pulumi.Input[str] status: Enable/disable storage. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usage: Use hard disk for logging or WAN Optimization (default = log).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] wanopt_mode: WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if media_status is not None:
            pulumi.set(__self__, "media_status", media_status)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if usage is not None:
            pulumi.set(__self__, "usage", usage)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wanopt_mode is not None:
            pulumi.set(__self__, "wanopt_mode", wanopt_mode)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        Partition device.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="mediaStatus")
    def media_status(self) -> Optional[pulumi.Input[str]]:
        """
        The physical status of current media. Valid values: `enable`, `disable`, `fail`.
        """
        return pulumi.get(self, "media_status")

    @media_status.setter
    def media_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "media_status", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Storage name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        """
        Set storage order.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Label of underlying partition.
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Partition size.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable storage. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def usage(self) -> Optional[pulumi.Input[str]]:
        """
        Use hard disk for logging or WAN Optimization (default = log).
        """
        return pulumi.get(self, "usage")

    @usage.setter
    def usage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage", value)

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
    @pulumi.getter(name="wanoptMode")
    def wanopt_mode(self) -> Optional[pulumi.Input[str]]:
        """
        WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
        """
        return pulumi.get(self, "wanopt_mode")

    @wanopt_mode.setter
    def wanopt_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wanopt_mode", value)


@pulumi.input_type
class _StorageState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 media_status: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wanopt_mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Storage resources.
        :param pulumi.Input[str] device: Partition device.
        :param pulumi.Input[str] media_status: The physical status of current media. Valid values: `enable`, `disable`, `fail`.
        :param pulumi.Input[str] name: Storage name.
        :param pulumi.Input[int] order: Set storage order.
        :param pulumi.Input[str] partition: Label of underlying partition.
        :param pulumi.Input[int] size: Partition size.
        :param pulumi.Input[str] status: Enable/disable storage. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usage: Use hard disk for logging or WAN Optimization (default = log).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] wanopt_mode: WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if media_status is not None:
            pulumi.set(__self__, "media_status", media_status)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if usage is not None:
            pulumi.set(__self__, "usage", usage)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wanopt_mode is not None:
            pulumi.set(__self__, "wanopt_mode", wanopt_mode)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        Partition device.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="mediaStatus")
    def media_status(self) -> Optional[pulumi.Input[str]]:
        """
        The physical status of current media. Valid values: `enable`, `disable`, `fail`.
        """
        return pulumi.get(self, "media_status")

    @media_status.setter
    def media_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "media_status", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Storage name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        """
        Set storage order.
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Label of underlying partition.
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Partition size.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable storage. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def usage(self) -> Optional[pulumi.Input[str]]:
        """
        Use hard disk for logging or WAN Optimization (default = log).
        """
        return pulumi.get(self, "usage")

    @usage.setter
    def usage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage", value)

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
    @pulumi.getter(name="wanoptMode")
    def wanopt_mode(self) -> Optional[pulumi.Input[str]]:
        """
        WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
        """
        return pulumi.get(self, "wanopt_mode")

    @wanopt_mode.setter
    def wanopt_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wanopt_mode", value)


class Storage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 media_status: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wanopt_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure logical storage.

        ## Import

        System Storage can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/storage:Storage labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/storage:Storage labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: Partition device.
        :param pulumi.Input[str] media_status: The physical status of current media. Valid values: `enable`, `disable`, `fail`.
        :param pulumi.Input[str] name: Storage name.
        :param pulumi.Input[int] order: Set storage order.
        :param pulumi.Input[str] partition: Label of underlying partition.
        :param pulumi.Input[int] size: Partition size.
        :param pulumi.Input[str] status: Enable/disable storage. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usage: Use hard disk for logging or WAN Optimization (default = log).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] wanopt_mode: WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[StorageArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure logical storage.

        ## Import

        System Storage can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/storage:Storage labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/storage:Storage labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param StorageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 media_status: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wanopt_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StorageArgs.__new__(StorageArgs)

            __props__.__dict__["device"] = device
            __props__.__dict__["media_status"] = media_status
            __props__.__dict__["name"] = name
            __props__.__dict__["order"] = order
            __props__.__dict__["partition"] = partition
            __props__.__dict__["size"] = size
            __props__.__dict__["status"] = status
            __props__.__dict__["usage"] = usage
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["wanopt_mode"] = wanopt_mode
        super(Storage, __self__).__init__(
            'fortios:system/storage:Storage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            media_status: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            order: Optional[pulumi.Input[int]] = None,
            partition: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            usage: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            wanopt_mode: Optional[pulumi.Input[str]] = None) -> 'Storage':
        """
        Get an existing Storage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: Partition device.
        :param pulumi.Input[str] media_status: The physical status of current media. Valid values: `enable`, `disable`, `fail`.
        :param pulumi.Input[str] name: Storage name.
        :param pulumi.Input[int] order: Set storage order.
        :param pulumi.Input[str] partition: Label of underlying partition.
        :param pulumi.Input[int] size: Partition size.
        :param pulumi.Input[str] status: Enable/disable storage. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] usage: Use hard disk for logging or WAN Optimization (default = log).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] wanopt_mode: WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StorageState.__new__(_StorageState)

        __props__.__dict__["device"] = device
        __props__.__dict__["media_status"] = media_status
        __props__.__dict__["name"] = name
        __props__.__dict__["order"] = order
        __props__.__dict__["partition"] = partition
        __props__.__dict__["size"] = size
        __props__.__dict__["status"] = status
        __props__.__dict__["usage"] = usage
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["wanopt_mode"] = wanopt_mode
        return Storage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[str]:
        """
        Partition device.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="mediaStatus")
    def media_status(self) -> pulumi.Output[str]:
        """
        The physical status of current media. Valid values: `enable`, `disable`, `fail`.
        """
        return pulumi.get(self, "media_status")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Storage name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> pulumi.Output[int]:
        """
        Set storage order.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def partition(self) -> pulumi.Output[str]:
        """
        Label of underlying partition.
        """
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Partition size.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable storage. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def usage(self) -> pulumi.Output[str]:
        """
        Use hard disk for logging or WAN Optimization (default = log).
        """
        return pulumi.get(self, "usage")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[str]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="wanoptMode")
    def wanopt_mode(self) -> pulumi.Output[str]:
        """
        WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
        """
        return pulumi.get(self, "wanopt_mode")

