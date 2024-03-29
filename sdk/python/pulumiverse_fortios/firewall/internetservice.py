# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InternetserviceArgs', 'Internetservice']

@pulumi.input_type
class InternetserviceArgs:
    def __init__(__self__, *,
                 database: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 extra_ip6_range_number: Optional[pulumi.Input[int]] = None,
                 extra_ip_range_number: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 icon_id: Optional[pulumi.Input[int]] = None,
                 ip6_range_number: Optional[pulumi.Input[int]] = None,
                 ip_number: Optional[pulumi.Input[int]] = None,
                 ip_range_number: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obsolete: Optional[pulumi.Input[int]] = None,
                 reputation: Optional[pulumi.Input[int]] = None,
                 singularity: Optional[pulumi.Input[int]] = None,
                 sld_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Internetservice resource.
        :param pulumi.Input[str] database: Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
        :param pulumi.Input[str] direction: How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
        :param pulumi.Input[int] extra_ip6_range_number: Extra number of IPv6 ranges.
        :param pulumi.Input[int] extra_ip_range_number: Extra number of IP ranges.
        :param pulumi.Input[int] fosid: Internet Service ID.
        :param pulumi.Input[int] icon_id: Icon ID of Internet Service.
        :param pulumi.Input[int] ip6_range_number: Number of IPv6 ranges.
        :param pulumi.Input[int] ip_number: Total number of IP addresses.
        :param pulumi.Input[int] ip_range_number: Total number of IP ranges.
        :param pulumi.Input[str] name: Internet Service name.
        :param pulumi.Input[int] obsolete: Indicates whether the Internet Service can be used.
        :param pulumi.Input[int] reputation: Reputation level of the Internet Service.
        :param pulumi.Input[int] singularity: Singular level of the Internet Service.
        :param pulumi.Input[int] sld_id: Second Level Domain.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if database is not None:
            pulumi.set(__self__, "database", database)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if extra_ip6_range_number is not None:
            pulumi.set(__self__, "extra_ip6_range_number", extra_ip6_range_number)
        if extra_ip_range_number is not None:
            pulumi.set(__self__, "extra_ip_range_number", extra_ip_range_number)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if icon_id is not None:
            pulumi.set(__self__, "icon_id", icon_id)
        if ip6_range_number is not None:
            pulumi.set(__self__, "ip6_range_number", ip6_range_number)
        if ip_number is not None:
            pulumi.set(__self__, "ip_number", ip_number)
        if ip_range_number is not None:
            pulumi.set(__self__, "ip_range_number", ip_range_number)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if obsolete is not None:
            pulumi.set(__self__, "obsolete", obsolete)
        if reputation is not None:
            pulumi.set(__self__, "reputation", reputation)
        if singularity is not None:
            pulumi.set(__self__, "singularity", singularity)
        if sld_id is not None:
            pulumi.set(__self__, "sld_id", sld_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="extraIp6RangeNumber")
    def extra_ip6_range_number(self) -> Optional[pulumi.Input[int]]:
        """
        Extra number of IPv6 ranges.
        """
        return pulumi.get(self, "extra_ip6_range_number")

    @extra_ip6_range_number.setter
    def extra_ip6_range_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "extra_ip6_range_number", value)

    @property
    @pulumi.getter(name="extraIpRangeNumber")
    def extra_ip_range_number(self) -> Optional[pulumi.Input[int]]:
        """
        Extra number of IP ranges.
        """
        return pulumi.get(self, "extra_ip_range_number")

    @extra_ip_range_number.setter
    def extra_ip_range_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "extra_ip_range_number", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Internet Service ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="iconId")
    def icon_id(self) -> Optional[pulumi.Input[int]]:
        """
        Icon ID of Internet Service.
        """
        return pulumi.get(self, "icon_id")

    @icon_id.setter
    def icon_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icon_id", value)

    @property
    @pulumi.getter(name="ip6RangeNumber")
    def ip6_range_number(self) -> Optional[pulumi.Input[int]]:
        """
        Number of IPv6 ranges.
        """
        return pulumi.get(self, "ip6_range_number")

    @ip6_range_number.setter
    def ip6_range_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip6_range_number", value)

    @property
    @pulumi.getter(name="ipNumber")
    def ip_number(self) -> Optional[pulumi.Input[int]]:
        """
        Total number of IP addresses.
        """
        return pulumi.get(self, "ip_number")

    @ip_number.setter
    def ip_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_number", value)

    @property
    @pulumi.getter(name="ipRangeNumber")
    def ip_range_number(self) -> Optional[pulumi.Input[int]]:
        """
        Total number of IP ranges.
        """
        return pulumi.get(self, "ip_range_number")

    @ip_range_number.setter
    def ip_range_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_range_number", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Internet Service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def obsolete(self) -> Optional[pulumi.Input[int]]:
        """
        Indicates whether the Internet Service can be used.
        """
        return pulumi.get(self, "obsolete")

    @obsolete.setter
    def obsolete(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "obsolete", value)

    @property
    @pulumi.getter
    def reputation(self) -> Optional[pulumi.Input[int]]:
        """
        Reputation level of the Internet Service.
        """
        return pulumi.get(self, "reputation")

    @reputation.setter
    def reputation(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reputation", value)

    @property
    @pulumi.getter
    def singularity(self) -> Optional[pulumi.Input[int]]:
        """
        Singular level of the Internet Service.
        """
        return pulumi.get(self, "singularity")

    @singularity.setter
    def singularity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "singularity", value)

    @property
    @pulumi.getter(name="sldId")
    def sld_id(self) -> Optional[pulumi.Input[int]]:
        """
        Second Level Domain.
        """
        return pulumi.get(self, "sld_id")

    @sld_id.setter
    def sld_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sld_id", value)

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
class _InternetserviceState:
    def __init__(__self__, *,
                 database: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 extra_ip6_range_number: Optional[pulumi.Input[int]] = None,
                 extra_ip_range_number: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 icon_id: Optional[pulumi.Input[int]] = None,
                 ip6_range_number: Optional[pulumi.Input[int]] = None,
                 ip_number: Optional[pulumi.Input[int]] = None,
                 ip_range_number: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obsolete: Optional[pulumi.Input[int]] = None,
                 reputation: Optional[pulumi.Input[int]] = None,
                 singularity: Optional[pulumi.Input[int]] = None,
                 sld_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Internetservice resources.
        :param pulumi.Input[str] database: Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
        :param pulumi.Input[str] direction: How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
        :param pulumi.Input[int] extra_ip6_range_number: Extra number of IPv6 ranges.
        :param pulumi.Input[int] extra_ip_range_number: Extra number of IP ranges.
        :param pulumi.Input[int] fosid: Internet Service ID.
        :param pulumi.Input[int] icon_id: Icon ID of Internet Service.
        :param pulumi.Input[int] ip6_range_number: Number of IPv6 ranges.
        :param pulumi.Input[int] ip_number: Total number of IP addresses.
        :param pulumi.Input[int] ip_range_number: Total number of IP ranges.
        :param pulumi.Input[str] name: Internet Service name.
        :param pulumi.Input[int] obsolete: Indicates whether the Internet Service can be used.
        :param pulumi.Input[int] reputation: Reputation level of the Internet Service.
        :param pulumi.Input[int] singularity: Singular level of the Internet Service.
        :param pulumi.Input[int] sld_id: Second Level Domain.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if database is not None:
            pulumi.set(__self__, "database", database)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if extra_ip6_range_number is not None:
            pulumi.set(__self__, "extra_ip6_range_number", extra_ip6_range_number)
        if extra_ip_range_number is not None:
            pulumi.set(__self__, "extra_ip_range_number", extra_ip_range_number)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if icon_id is not None:
            pulumi.set(__self__, "icon_id", icon_id)
        if ip6_range_number is not None:
            pulumi.set(__self__, "ip6_range_number", ip6_range_number)
        if ip_number is not None:
            pulumi.set(__self__, "ip_number", ip_number)
        if ip_range_number is not None:
            pulumi.set(__self__, "ip_range_number", ip_range_number)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if obsolete is not None:
            pulumi.set(__self__, "obsolete", obsolete)
        if reputation is not None:
            pulumi.set(__self__, "reputation", reputation)
        if singularity is not None:
            pulumi.set(__self__, "singularity", singularity)
        if sld_id is not None:
            pulumi.set(__self__, "sld_id", sld_id)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="extraIp6RangeNumber")
    def extra_ip6_range_number(self) -> Optional[pulumi.Input[int]]:
        """
        Extra number of IPv6 ranges.
        """
        return pulumi.get(self, "extra_ip6_range_number")

    @extra_ip6_range_number.setter
    def extra_ip6_range_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "extra_ip6_range_number", value)

    @property
    @pulumi.getter(name="extraIpRangeNumber")
    def extra_ip_range_number(self) -> Optional[pulumi.Input[int]]:
        """
        Extra number of IP ranges.
        """
        return pulumi.get(self, "extra_ip_range_number")

    @extra_ip_range_number.setter
    def extra_ip_range_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "extra_ip_range_number", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Internet Service ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="iconId")
    def icon_id(self) -> Optional[pulumi.Input[int]]:
        """
        Icon ID of Internet Service.
        """
        return pulumi.get(self, "icon_id")

    @icon_id.setter
    def icon_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "icon_id", value)

    @property
    @pulumi.getter(name="ip6RangeNumber")
    def ip6_range_number(self) -> Optional[pulumi.Input[int]]:
        """
        Number of IPv6 ranges.
        """
        return pulumi.get(self, "ip6_range_number")

    @ip6_range_number.setter
    def ip6_range_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip6_range_number", value)

    @property
    @pulumi.getter(name="ipNumber")
    def ip_number(self) -> Optional[pulumi.Input[int]]:
        """
        Total number of IP addresses.
        """
        return pulumi.get(self, "ip_number")

    @ip_number.setter
    def ip_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_number", value)

    @property
    @pulumi.getter(name="ipRangeNumber")
    def ip_range_number(self) -> Optional[pulumi.Input[int]]:
        """
        Total number of IP ranges.
        """
        return pulumi.get(self, "ip_range_number")

    @ip_range_number.setter
    def ip_range_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_range_number", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Internet Service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def obsolete(self) -> Optional[pulumi.Input[int]]:
        """
        Indicates whether the Internet Service can be used.
        """
        return pulumi.get(self, "obsolete")

    @obsolete.setter
    def obsolete(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "obsolete", value)

    @property
    @pulumi.getter
    def reputation(self) -> Optional[pulumi.Input[int]]:
        """
        Reputation level of the Internet Service.
        """
        return pulumi.get(self, "reputation")

    @reputation.setter
    def reputation(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reputation", value)

    @property
    @pulumi.getter
    def singularity(self) -> Optional[pulumi.Input[int]]:
        """
        Singular level of the Internet Service.
        """
        return pulumi.get(self, "singularity")

    @singularity.setter
    def singularity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "singularity", value)

    @property
    @pulumi.getter(name="sldId")
    def sld_id(self) -> Optional[pulumi.Input[int]]:
        """
        Second Level Domain.
        """
        return pulumi.get(self, "sld_id")

    @sld_id.setter
    def sld_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sld_id", value)

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


class Internetservice(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 extra_ip6_range_number: Optional[pulumi.Input[int]] = None,
                 extra_ip_range_number: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 icon_id: Optional[pulumi.Input[int]] = None,
                 ip6_range_number: Optional[pulumi.Input[int]] = None,
                 ip_number: Optional[pulumi.Input[int]] = None,
                 ip_range_number: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obsolete: Optional[pulumi.Input[int]] = None,
                 reputation: Optional[pulumi.Input[int]] = None,
                 singularity: Optional[pulumi.Input[int]] = None,
                 sld_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Show Internet Service application.

        ## Import

        Firewall InternetService can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/internetservice:Internetservice labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/internetservice:Internetservice labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
        :param pulumi.Input[str] direction: How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
        :param pulumi.Input[int] extra_ip6_range_number: Extra number of IPv6 ranges.
        :param pulumi.Input[int] extra_ip_range_number: Extra number of IP ranges.
        :param pulumi.Input[int] fosid: Internet Service ID.
        :param pulumi.Input[int] icon_id: Icon ID of Internet Service.
        :param pulumi.Input[int] ip6_range_number: Number of IPv6 ranges.
        :param pulumi.Input[int] ip_number: Total number of IP addresses.
        :param pulumi.Input[int] ip_range_number: Total number of IP ranges.
        :param pulumi.Input[str] name: Internet Service name.
        :param pulumi.Input[int] obsolete: Indicates whether the Internet Service can be used.
        :param pulumi.Input[int] reputation: Reputation level of the Internet Service.
        :param pulumi.Input[int] singularity: Singular level of the Internet Service.
        :param pulumi.Input[int] sld_id: Second Level Domain.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InternetserviceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Show Internet Service application.

        ## Import

        Firewall InternetService can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:firewall/internetservice:Internetservice labelname {{fosid}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:firewall/internetservice:Internetservice labelname {{fosid}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param InternetserviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InternetserviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 extra_ip6_range_number: Optional[pulumi.Input[int]] = None,
                 extra_ip_range_number: Optional[pulumi.Input[int]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 icon_id: Optional[pulumi.Input[int]] = None,
                 ip6_range_number: Optional[pulumi.Input[int]] = None,
                 ip_number: Optional[pulumi.Input[int]] = None,
                 ip_range_number: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 obsolete: Optional[pulumi.Input[int]] = None,
                 reputation: Optional[pulumi.Input[int]] = None,
                 singularity: Optional[pulumi.Input[int]] = None,
                 sld_id: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InternetserviceArgs.__new__(InternetserviceArgs)

            __props__.__dict__["database"] = database
            __props__.__dict__["direction"] = direction
            __props__.__dict__["extra_ip6_range_number"] = extra_ip6_range_number
            __props__.__dict__["extra_ip_range_number"] = extra_ip_range_number
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["icon_id"] = icon_id
            __props__.__dict__["ip6_range_number"] = ip6_range_number
            __props__.__dict__["ip_number"] = ip_number
            __props__.__dict__["ip_range_number"] = ip_range_number
            __props__.__dict__["name"] = name
            __props__.__dict__["obsolete"] = obsolete
            __props__.__dict__["reputation"] = reputation
            __props__.__dict__["singularity"] = singularity
            __props__.__dict__["sld_id"] = sld_id
            __props__.__dict__["vdomparam"] = vdomparam
        super(Internetservice, __self__).__init__(
            'fortios:firewall/internetservice:Internetservice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database: Optional[pulumi.Input[str]] = None,
            direction: Optional[pulumi.Input[str]] = None,
            extra_ip6_range_number: Optional[pulumi.Input[int]] = None,
            extra_ip_range_number: Optional[pulumi.Input[int]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            icon_id: Optional[pulumi.Input[int]] = None,
            ip6_range_number: Optional[pulumi.Input[int]] = None,
            ip_number: Optional[pulumi.Input[int]] = None,
            ip_range_number: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            obsolete: Optional[pulumi.Input[int]] = None,
            reputation: Optional[pulumi.Input[int]] = None,
            singularity: Optional[pulumi.Input[int]] = None,
            sld_id: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Internetservice':
        """
        Get an existing Internetservice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
        :param pulumi.Input[str] direction: How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
        :param pulumi.Input[int] extra_ip6_range_number: Extra number of IPv6 ranges.
        :param pulumi.Input[int] extra_ip_range_number: Extra number of IP ranges.
        :param pulumi.Input[int] fosid: Internet Service ID.
        :param pulumi.Input[int] icon_id: Icon ID of Internet Service.
        :param pulumi.Input[int] ip6_range_number: Number of IPv6 ranges.
        :param pulumi.Input[int] ip_number: Total number of IP addresses.
        :param pulumi.Input[int] ip_range_number: Total number of IP ranges.
        :param pulumi.Input[str] name: Internet Service name.
        :param pulumi.Input[int] obsolete: Indicates whether the Internet Service can be used.
        :param pulumi.Input[int] reputation: Reputation level of the Internet Service.
        :param pulumi.Input[int] singularity: Singular level of the Internet Service.
        :param pulumi.Input[int] sld_id: Second Level Domain.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InternetserviceState.__new__(_InternetserviceState)

        __props__.__dict__["database"] = database
        __props__.__dict__["direction"] = direction
        __props__.__dict__["extra_ip6_range_number"] = extra_ip6_range_number
        __props__.__dict__["extra_ip_range_number"] = extra_ip_range_number
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["icon_id"] = icon_id
        __props__.__dict__["ip6_range_number"] = ip6_range_number
        __props__.__dict__["ip_number"] = ip_number
        __props__.__dict__["ip_range_number"] = ip_range_number
        __props__.__dict__["name"] = name
        __props__.__dict__["obsolete"] = obsolete
        __props__.__dict__["reputation"] = reputation
        __props__.__dict__["singularity"] = singularity
        __props__.__dict__["sld_id"] = sld_id
        __props__.__dict__["vdomparam"] = vdomparam
        return Internetservice(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[str]:
        """
        How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter(name="extraIp6RangeNumber")
    def extra_ip6_range_number(self) -> pulumi.Output[int]:
        """
        Extra number of IPv6 ranges.
        """
        return pulumi.get(self, "extra_ip6_range_number")

    @property
    @pulumi.getter(name="extraIpRangeNumber")
    def extra_ip_range_number(self) -> pulumi.Output[int]:
        """
        Extra number of IP ranges.
        """
        return pulumi.get(self, "extra_ip_range_number")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Internet Service ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="iconId")
    def icon_id(self) -> pulumi.Output[int]:
        """
        Icon ID of Internet Service.
        """
        return pulumi.get(self, "icon_id")

    @property
    @pulumi.getter(name="ip6RangeNumber")
    def ip6_range_number(self) -> pulumi.Output[int]:
        """
        Number of IPv6 ranges.
        """
        return pulumi.get(self, "ip6_range_number")

    @property
    @pulumi.getter(name="ipNumber")
    def ip_number(self) -> pulumi.Output[int]:
        """
        Total number of IP addresses.
        """
        return pulumi.get(self, "ip_number")

    @property
    @pulumi.getter(name="ipRangeNumber")
    def ip_range_number(self) -> pulumi.Output[int]:
        """
        Total number of IP ranges.
        """
        return pulumi.get(self, "ip_range_number")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Internet Service name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def obsolete(self) -> pulumi.Output[int]:
        """
        Indicates whether the Internet Service can be used.
        """
        return pulumi.get(self, "obsolete")

    @property
    @pulumi.getter
    def reputation(self) -> pulumi.Output[int]:
        """
        Reputation level of the Internet Service.
        """
        return pulumi.get(self, "reputation")

    @property
    @pulumi.getter
    def singularity(self) -> pulumi.Output[int]:
        """
        Singular level of the Internet Service.
        """
        return pulumi.get(self, "singularity")

    @property
    @pulumi.getter(name="sldId")
    def sld_id(self) -> pulumi.Output[int]:
        """
        Second Level Domain.
        """
        return pulumi.get(self, "sld_id")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

