# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DeviceArgs', 'Device']

@pulumi.input_type
class DeviceArgs:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 avatar: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 master_device: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input['DeviceTaggingArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Device resource.
        :param pulumi.Input[str] alias: Device alias.
        :param pulumi.Input[str] avatar: Image file for avatar (maximum 4K base64 encoded).
        :param pulumi.Input[str] category: Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] mac: Device MAC address.
        :param pulumi.Input[str] master_device: Master device (optional).
        :param pulumi.Input[Sequence[pulumi.Input['DeviceTaggingArgs']]] taggings: Config object tagging. The structure of `tagging` block is documented below.
        :param pulumi.Input[str] type: Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        :param pulumi.Input[str] user: User name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if avatar is not None:
            pulumi.set(__self__, "avatar", avatar)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if master_device is not None:
            pulumi.set(__self__, "master_device", master_device)
        if taggings is not None:
            pulumi.set(__self__, "taggings", taggings)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user is not None:
            pulumi.set(__self__, "user", user)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        Device alias.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def avatar(self) -> Optional[pulumi.Input[str]]:
        """
        Image file for avatar (maximum 4K base64 encoded).
        """
        return pulumi.get(self, "avatar")

    @avatar.setter
    def avatar(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "avatar", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        Device MAC address.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter(name="masterDevice")
    def master_device(self) -> Optional[pulumi.Input[str]]:
        """
        Master device (optional).
        """
        return pulumi.get(self, "master_device")

    @master_device.setter
    def master_device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_device", value)

    @property
    @pulumi.getter
    def taggings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeviceTaggingArgs']]]]:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @taggings.setter
    def taggings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeviceTaggingArgs']]]]):
        pulumi.set(self, "taggings", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        User name.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)

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
class _DeviceState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 avatar: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 master_device: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input['DeviceTaggingArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Device resources.
        :param pulumi.Input[str] alias: Device alias.
        :param pulumi.Input[str] avatar: Image file for avatar (maximum 4K base64 encoded).
        :param pulumi.Input[str] category: Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] mac: Device MAC address.
        :param pulumi.Input[str] master_device: Master device (optional).
        :param pulumi.Input[Sequence[pulumi.Input['DeviceTaggingArgs']]] taggings: Config object tagging. The structure of `tagging` block is documented below.
        :param pulumi.Input[str] type: Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        :param pulumi.Input[str] user: User name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if avatar is not None:
            pulumi.set(__self__, "avatar", avatar)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if master_device is not None:
            pulumi.set(__self__, "master_device", master_device)
        if taggings is not None:
            pulumi.set(__self__, "taggings", taggings)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user is not None:
            pulumi.set(__self__, "user", user)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        Device alias.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def avatar(self) -> Optional[pulumi.Input[str]]:
        """
        Image file for avatar (maximum 4K base64 encoded).
        """
        return pulumi.get(self, "avatar")

    @avatar.setter
    def avatar(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "avatar", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        Device MAC address.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter(name="masterDevice")
    def master_device(self) -> Optional[pulumi.Input[str]]:
        """
        Master device (optional).
        """
        return pulumi.get(self, "master_device")

    @master_device.setter
    def master_device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_device", value)

    @property
    @pulumi.getter
    def taggings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeviceTaggingArgs']]]]:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @taggings.setter
    def taggings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeviceTaggingArgs']]]]):
        pulumi.set(self, "taggings", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        User name.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)

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


class Device(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 avatar: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 master_device: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeviceTaggingArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure devices. Applies to FortiOS Version `<= 6.2.0`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.user.Device("trname",
            alias="1",
            category="amazon-device",
            mac="08:00:20:0a:8c:6d",
            type="unknown")
        ```

        ## Import

        User Device can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:user/device:Device labelname {{alias}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:user/device:Device labelname {{alias}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Device alias.
        :param pulumi.Input[str] avatar: Image file for avatar (maximum 4K base64 encoded).
        :param pulumi.Input[str] category: Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] mac: Device MAC address.
        :param pulumi.Input[str] master_device: Master device (optional).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeviceTaggingArgs']]]] taggings: Config object tagging. The structure of `tagging` block is documented below.
        :param pulumi.Input[str] type: Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        :param pulumi.Input[str] user: User name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DeviceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure devices. Applies to FortiOS Version `<= 6.2.0`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.user.Device("trname",
            alias="1",
            category="amazon-device",
            mac="08:00:20:0a:8c:6d",
            type="unknown")
        ```

        ## Import

        User Device can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:user/device:Device labelname {{alias}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:user/device:Device labelname {{alias}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param DeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 avatar: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 master_device: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeviceTaggingArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceArgs.__new__(DeviceArgs)

            __props__.__dict__["alias"] = alias
            __props__.__dict__["avatar"] = avatar
            __props__.__dict__["category"] = category
            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["mac"] = mac
            __props__.__dict__["master_device"] = master_device
            __props__.__dict__["taggings"] = taggings
            __props__.__dict__["type"] = type
            __props__.__dict__["user"] = user
            __props__.__dict__["vdomparam"] = vdomparam
        super(Device, __self__).__init__(
            'fortios:user/device:Device',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            avatar: Optional[pulumi.Input[str]] = None,
            category: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            master_device: Optional[pulumi.Input[str]] = None,
            taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeviceTaggingArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Device':
        """
        Get an existing Device resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Device alias.
        :param pulumi.Input[str] avatar: Image file for avatar (maximum 4K base64 encoded).
        :param pulumi.Input[str] category: Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] get_all_tables: Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        :param pulumi.Input[str] mac: Device MAC address.
        :param pulumi.Input[str] master_device: Master device (optional).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DeviceTaggingArgs']]]] taggings: Config object tagging. The structure of `tagging` block is documented below.
        :param pulumi.Input[str] type: Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        :param pulumi.Input[str] user: User name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceState.__new__(_DeviceState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["avatar"] = avatar
        __props__.__dict__["category"] = category
        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["mac"] = mac
        __props__.__dict__["master_device"] = master_device
        __props__.__dict__["taggings"] = taggings
        __props__.__dict__["type"] = type
        __props__.__dict__["user"] = user
        __props__.__dict__["vdomparam"] = vdomparam
        return Device(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        Device alias.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def avatar(self) -> pulumi.Output[Optional[str]]:
        """
        Image file for avatar (maximum 4K base64 encoded).
        """
        return pulumi.get(self, "avatar")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        """
        Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        """
        Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
        """
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        Device MAC address.
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter(name="masterDevice")
    def master_device(self) -> pulumi.Output[str]:
        """
        Master device (optional).
        """
        return pulumi.get(self, "master_device")

    @property
    @pulumi.getter
    def taggings(self) -> pulumi.Output[Optional[Sequence['outputs.DeviceTagging']]]:
        """
        Config object tagging. The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        User name.
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

