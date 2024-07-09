# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['IpsurlfiltersettingArgs', 'Ipsurlfiltersetting']

@pulumi.input_type
class IpsurlfiltersettingArgs:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 geo_filter: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ipsurlfiltersetting resource.
        :param pulumi.Input[str] device: Interface for this route.
        :param pulumi.Input[int] distance: Administrative distance (1 - 255) for this route.
        :param pulumi.Input[str] gateway: Gateway IP address for this route.
        :param pulumi.Input[str] geo_filter: Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if geo_filter is not None:
            pulumi.set(__self__, "geo_filter", geo_filter)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        Interface for this route.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        """
        Administrative distance (1 - 255) for this route.
        """
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        Gateway IP address for this route.
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="geoFilter")
    def geo_filter(self) -> Optional[pulumi.Input[str]]:
        """
        Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        """
        return pulumi.get(self, "geo_filter")

    @geo_filter.setter
    def geo_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geo_filter", value)

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
class _IpsurlfiltersettingState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 geo_filter: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ipsurlfiltersetting resources.
        :param pulumi.Input[str] device: Interface for this route.
        :param pulumi.Input[int] distance: Administrative distance (1 - 255) for this route.
        :param pulumi.Input[str] gateway: Gateway IP address for this route.
        :param pulumi.Input[str] geo_filter: Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if geo_filter is not None:
            pulumi.set(__self__, "geo_filter", geo_filter)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        Interface for this route.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        """
        Administrative distance (1 - 255) for this route.
        """
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        Gateway IP address for this route.
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="geoFilter")
    def geo_filter(self) -> Optional[pulumi.Input[str]]:
        """
        Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        """
        return pulumi.get(self, "geo_filter")

    @geo_filter.setter
    def geo_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geo_filter", value)

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


class Ipsurlfiltersetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 geo_filter: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPS URL filter settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.filter.web.Ipsurlfiltersetting("trname",
            distance=1,
            gateway="0.0.0.0")
        ```

        ## Import

        Webfilter IpsUrlfilterSetting can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:filter/web/ipsurlfiltersetting:Ipsurlfiltersetting labelname WebfilterIpsUrlfilterSetting
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:filter/web/ipsurlfiltersetting:Ipsurlfiltersetting labelname WebfilterIpsUrlfilterSetting
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: Interface for this route.
        :param pulumi.Input[int] distance: Administrative distance (1 - 255) for this route.
        :param pulumi.Input[str] gateway: Gateway IP address for this route.
        :param pulumi.Input[str] geo_filter: Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IpsurlfiltersettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPS URL filter settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.filter.web.Ipsurlfiltersetting("trname",
            distance=1,
            gateway="0.0.0.0")
        ```

        ## Import

        Webfilter IpsUrlfilterSetting can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:filter/web/ipsurlfiltersetting:Ipsurlfiltersetting labelname WebfilterIpsUrlfilterSetting
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:filter/web/ipsurlfiltersetting:Ipsurlfiltersetting labelname WebfilterIpsUrlfilterSetting
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param IpsurlfiltersettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpsurlfiltersettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 geo_filter: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpsurlfiltersettingArgs.__new__(IpsurlfiltersettingArgs)

            __props__.__dict__["device"] = device
            __props__.__dict__["distance"] = distance
            __props__.__dict__["gateway"] = gateway
            __props__.__dict__["geo_filter"] = geo_filter
            __props__.__dict__["vdomparam"] = vdomparam
        super(Ipsurlfiltersetting, __self__).__init__(
            'fortios:filter/web/ipsurlfiltersetting:Ipsurlfiltersetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            distance: Optional[pulumi.Input[int]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            geo_filter: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Ipsurlfiltersetting':
        """
        Get an existing Ipsurlfiltersetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: Interface for this route.
        :param pulumi.Input[int] distance: Administrative distance (1 - 255) for this route.
        :param pulumi.Input[str] gateway: Gateway IP address for this route.
        :param pulumi.Input[str] geo_filter: Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpsurlfiltersettingState.__new__(_IpsurlfiltersettingState)

        __props__.__dict__["device"] = device
        __props__.__dict__["distance"] = distance
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["geo_filter"] = geo_filter
        __props__.__dict__["vdomparam"] = vdomparam
        return Ipsurlfiltersetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[str]:
        """
        Interface for this route.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def distance(self) -> pulumi.Output[int]:
        """
        Administrative distance (1 - 255) for this route.
        """
        return pulumi.get(self, "distance")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        """
        Gateway IP address for this route.
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="geoFilter")
    def geo_filter(self) -> pulumi.Output[Optional[str]]:
        """
        Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
        """
        return pulumi.get(self, "geo_filter")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

