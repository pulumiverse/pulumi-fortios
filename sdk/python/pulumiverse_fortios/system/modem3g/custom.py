# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['CustomArgs', 'Custom']

@pulumi.input_type
class CustomArgs:
    def __init__(__self__, *,
                 class_id: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 init_string: Optional[pulumi.Input[str]] = None,
                 model: Optional[pulumi.Input[str]] = None,
                 modeswitch_string: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None,
                 vendor_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Custom resource.
        :param pulumi.Input[str] class_id: USB interface class in hexadecimal format (00-ff).
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] init_string: Init string in hexadecimal format (even length).
        :param pulumi.Input[str] model: MODEM model name.
        :param pulumi.Input[str] modeswitch_string: USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        :param pulumi.Input[str] product_id: USB product ID in hexadecimal format (0000-ffff).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vendor: MODEM vendor name.
        :param pulumi.Input[str] vendor_id: USB vendor ID in hexadecimal format (0000-ffff).
        """
        if class_id is not None:
            pulumi.set(__self__, "class_id", class_id)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if init_string is not None:
            pulumi.set(__self__, "init_string", init_string)
        if model is not None:
            pulumi.set(__self__, "model", model)
        if modeswitch_string is not None:
            pulumi.set(__self__, "modeswitch_string", modeswitch_string)
        if product_id is not None:
            pulumi.set(__self__, "product_id", product_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vendor is not None:
            pulumi.set(__self__, "vendor", vendor)
        if vendor_id is not None:
            pulumi.set(__self__, "vendor_id", vendor_id)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> Optional[pulumi.Input[str]]:
        """
        USB interface class in hexadecimal format (00-ff).
        """
        return pulumi.get(self, "class_id")

    @class_id.setter
    def class_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "class_id", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="initString")
    def init_string(self) -> Optional[pulumi.Input[str]]:
        """
        Init string in hexadecimal format (even length).
        """
        return pulumi.get(self, "init_string")

    @init_string.setter
    def init_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "init_string", value)

    @property
    @pulumi.getter
    def model(self) -> Optional[pulumi.Input[str]]:
        """
        MODEM model name.
        """
        return pulumi.get(self, "model")

    @model.setter
    def model(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model", value)

    @property
    @pulumi.getter(name="modeswitchString")
    def modeswitch_string(self) -> Optional[pulumi.Input[str]]:
        """
        USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        """
        return pulumi.get(self, "modeswitch_string")

    @modeswitch_string.setter
    def modeswitch_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modeswitch_string", value)

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> Optional[pulumi.Input[str]]:
        """
        USB product ID in hexadecimal format (0000-ffff).
        """
        return pulumi.get(self, "product_id")

    @product_id.setter
    def product_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_id", value)

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
    @pulumi.getter
    def vendor(self) -> Optional[pulumi.Input[str]]:
        """
        MODEM vendor name.
        """
        return pulumi.get(self, "vendor")

    @vendor.setter
    def vendor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor", value)

    @property
    @pulumi.getter(name="vendorId")
    def vendor_id(self) -> Optional[pulumi.Input[str]]:
        """
        USB vendor ID in hexadecimal format (0000-ffff).
        """
        return pulumi.get(self, "vendor_id")

    @vendor_id.setter
    def vendor_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor_id", value)


@pulumi.input_type
class _CustomState:
    def __init__(__self__, *,
                 class_id: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 init_string: Optional[pulumi.Input[str]] = None,
                 model: Optional[pulumi.Input[str]] = None,
                 modeswitch_string: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None,
                 vendor_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Custom resources.
        :param pulumi.Input[str] class_id: USB interface class in hexadecimal format (00-ff).
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] init_string: Init string in hexadecimal format (even length).
        :param pulumi.Input[str] model: MODEM model name.
        :param pulumi.Input[str] modeswitch_string: USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        :param pulumi.Input[str] product_id: USB product ID in hexadecimal format (0000-ffff).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vendor: MODEM vendor name.
        :param pulumi.Input[str] vendor_id: USB vendor ID in hexadecimal format (0000-ffff).
        """
        if class_id is not None:
            pulumi.set(__self__, "class_id", class_id)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if init_string is not None:
            pulumi.set(__self__, "init_string", init_string)
        if model is not None:
            pulumi.set(__self__, "model", model)
        if modeswitch_string is not None:
            pulumi.set(__self__, "modeswitch_string", modeswitch_string)
        if product_id is not None:
            pulumi.set(__self__, "product_id", product_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vendor is not None:
            pulumi.set(__self__, "vendor", vendor)
        if vendor_id is not None:
            pulumi.set(__self__, "vendor_id", vendor_id)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> Optional[pulumi.Input[str]]:
        """
        USB interface class in hexadecimal format (00-ff).
        """
        return pulumi.get(self, "class_id")

    @class_id.setter
    def class_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "class_id", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="initString")
    def init_string(self) -> Optional[pulumi.Input[str]]:
        """
        Init string in hexadecimal format (even length).
        """
        return pulumi.get(self, "init_string")

    @init_string.setter
    def init_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "init_string", value)

    @property
    @pulumi.getter
    def model(self) -> Optional[pulumi.Input[str]]:
        """
        MODEM model name.
        """
        return pulumi.get(self, "model")

    @model.setter
    def model(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model", value)

    @property
    @pulumi.getter(name="modeswitchString")
    def modeswitch_string(self) -> Optional[pulumi.Input[str]]:
        """
        USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        """
        return pulumi.get(self, "modeswitch_string")

    @modeswitch_string.setter
    def modeswitch_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modeswitch_string", value)

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> Optional[pulumi.Input[str]]:
        """
        USB product ID in hexadecimal format (0000-ffff).
        """
        return pulumi.get(self, "product_id")

    @product_id.setter
    def product_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_id", value)

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
    @pulumi.getter
    def vendor(self) -> Optional[pulumi.Input[str]]:
        """
        MODEM vendor name.
        """
        return pulumi.get(self, "vendor")

    @vendor.setter
    def vendor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor", value)

    @property
    @pulumi.getter(name="vendorId")
    def vendor_id(self) -> Optional[pulumi.Input[str]]:
        """
        USB vendor ID in hexadecimal format (0000-ffff).
        """
        return pulumi.get(self, "vendor_id")

    @vendor_id.setter
    def vendor_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor_id", value)


class Custom(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 class_id: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 init_string: Optional[pulumi.Input[str]] = None,
                 model: Optional[pulumi.Input[str]] = None,
                 modeswitch_string: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None,
                 vendor_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        3G MODEM custom. Applies to FortiOS Version `7.0.4`.

        ## Import

        System3GModem Custom can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/modem3g/custom:Custom labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/modem3g/custom:Custom labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] class_id: USB interface class in hexadecimal format (00-ff).
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] init_string: Init string in hexadecimal format (even length).
        :param pulumi.Input[str] model: MODEM model name.
        :param pulumi.Input[str] modeswitch_string: USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        :param pulumi.Input[str] product_id: USB product ID in hexadecimal format (0000-ffff).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vendor: MODEM vendor name.
        :param pulumi.Input[str] vendor_id: USB vendor ID in hexadecimal format (0000-ffff).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CustomArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        3G MODEM custom. Applies to FortiOS Version `7.0.4`.

        ## Import

        System3GModem Custom can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/modem3g/custom:Custom labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/modem3g/custom:Custom labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param CustomArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 class_id: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 init_string: Optional[pulumi.Input[str]] = None,
                 model: Optional[pulumi.Input[str]] = None,
                 modeswitch_string: Optional[pulumi.Input[str]] = None,
                 product_id: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None,
                 vendor_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomArgs.__new__(CustomArgs)

            __props__.__dict__["class_id"] = class_id
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["init_string"] = init_string
            __props__.__dict__["model"] = model
            __props__.__dict__["modeswitch_string"] = modeswitch_string
            __props__.__dict__["product_id"] = product_id
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vendor"] = vendor
            __props__.__dict__["vendor_id"] = vendor_id
        super(Custom, __self__).__init__(
            'fortios:system/modem3g/custom:Custom',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            class_id: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            init_string: Optional[pulumi.Input[str]] = None,
            model: Optional[pulumi.Input[str]] = None,
            modeswitch_string: Optional[pulumi.Input[str]] = None,
            product_id: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vendor: Optional[pulumi.Input[str]] = None,
            vendor_id: Optional[pulumi.Input[str]] = None) -> 'Custom':
        """
        Get an existing Custom resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] class_id: USB interface class in hexadecimal format (00-ff).
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] init_string: Init string in hexadecimal format (even length).
        :param pulumi.Input[str] model: MODEM model name.
        :param pulumi.Input[str] modeswitch_string: USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        :param pulumi.Input[str] product_id: USB product ID in hexadecimal format (0000-ffff).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vendor: MODEM vendor name.
        :param pulumi.Input[str] vendor_id: USB vendor ID in hexadecimal format (0000-ffff).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomState.__new__(_CustomState)

        __props__.__dict__["class_id"] = class_id
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["init_string"] = init_string
        __props__.__dict__["model"] = model
        __props__.__dict__["modeswitch_string"] = modeswitch_string
        __props__.__dict__["product_id"] = product_id
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vendor"] = vendor
        __props__.__dict__["vendor_id"] = vendor_id
        return Custom(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> pulumi.Output[str]:
        """
        USB interface class in hexadecimal format (00-ff).
        """
        return pulumi.get(self, "class_id")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="initString")
    def init_string(self) -> pulumi.Output[str]:
        """
        Init string in hexadecimal format (even length).
        """
        return pulumi.get(self, "init_string")

    @property
    @pulumi.getter
    def model(self) -> pulumi.Output[str]:
        """
        MODEM model name.
        """
        return pulumi.get(self, "model")

    @property
    @pulumi.getter(name="modeswitchString")
    def modeswitch_string(self) -> pulumi.Output[str]:
        """
        USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
        """
        return pulumi.get(self, "modeswitch_string")

    @property
    @pulumi.getter(name="productId")
    def product_id(self) -> pulumi.Output[str]:
        """
        USB product ID in hexadecimal format (0000-ffff).
        """
        return pulumi.get(self, "product_id")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vendor(self) -> pulumi.Output[str]:
        """
        MODEM vendor name.
        """
        return pulumi.get(self, "vendor")

    @property
    @pulumi.getter(name="vendorId")
    def vendor_id(self) -> pulumi.Output[str]:
        """
        USB vendor ID in hexadecimal format (0000-ffff).
        """
        return pulumi.get(self, "vendor_id")

