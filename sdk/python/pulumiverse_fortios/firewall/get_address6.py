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

__all__ = [
    'GetAddress6Result',
    'AwaitableGetAddress6Result',
    'get_address6',
    'get_address6_output',
]

@pulumi.output_type
class GetAddress6Result:
    """
    A collection of values returned by getAddress6.
    """
    def __init__(__self__, cache_ttl=None, color=None, comment=None, country=None, end_ip=None, end_mac=None, epg_name=None, fabric_object=None, fqdn=None, host=None, host_type=None, id=None, ip6=None, lists=None, macaddrs=None, name=None, obj_id=None, route_tag=None, sdn=None, sdn_tag=None, start_ip=None, start_mac=None, subnet_segments=None, taggings=None, template=None, tenant=None, type=None, uuid=None, vdomparam=None, visibility=None):
        if cache_ttl and not isinstance(cache_ttl, int):
            raise TypeError("Expected argument 'cache_ttl' to be a int")
        pulumi.set(__self__, "cache_ttl", cache_ttl)
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if country and not isinstance(country, str):
            raise TypeError("Expected argument 'country' to be a str")
        pulumi.set(__self__, "country", country)
        if end_ip and not isinstance(end_ip, str):
            raise TypeError("Expected argument 'end_ip' to be a str")
        pulumi.set(__self__, "end_ip", end_ip)
        if end_mac and not isinstance(end_mac, str):
            raise TypeError("Expected argument 'end_mac' to be a str")
        pulumi.set(__self__, "end_mac", end_mac)
        if epg_name and not isinstance(epg_name, str):
            raise TypeError("Expected argument 'epg_name' to be a str")
        pulumi.set(__self__, "epg_name", epg_name)
        if fabric_object and not isinstance(fabric_object, str):
            raise TypeError("Expected argument 'fabric_object' to be a str")
        pulumi.set(__self__, "fabric_object", fabric_object)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if host_type and not isinstance(host_type, str):
            raise TypeError("Expected argument 'host_type' to be a str")
        pulumi.set(__self__, "host_type", host_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip6 and not isinstance(ip6, str):
            raise TypeError("Expected argument 'ip6' to be a str")
        pulumi.set(__self__, "ip6", ip6)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if macaddrs and not isinstance(macaddrs, list):
            raise TypeError("Expected argument 'macaddrs' to be a list")
        pulumi.set(__self__, "macaddrs", macaddrs)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if obj_id and not isinstance(obj_id, str):
            raise TypeError("Expected argument 'obj_id' to be a str")
        pulumi.set(__self__, "obj_id", obj_id)
        if route_tag and not isinstance(route_tag, int):
            raise TypeError("Expected argument 'route_tag' to be a int")
        pulumi.set(__self__, "route_tag", route_tag)
        if sdn and not isinstance(sdn, str):
            raise TypeError("Expected argument 'sdn' to be a str")
        pulumi.set(__self__, "sdn", sdn)
        if sdn_tag and not isinstance(sdn_tag, str):
            raise TypeError("Expected argument 'sdn_tag' to be a str")
        pulumi.set(__self__, "sdn_tag", sdn_tag)
        if start_ip and not isinstance(start_ip, str):
            raise TypeError("Expected argument 'start_ip' to be a str")
        pulumi.set(__self__, "start_ip", start_ip)
        if start_mac and not isinstance(start_mac, str):
            raise TypeError("Expected argument 'start_mac' to be a str")
        pulumi.set(__self__, "start_mac", start_mac)
        if subnet_segments and not isinstance(subnet_segments, list):
            raise TypeError("Expected argument 'subnet_segments' to be a list")
        pulumi.set(__self__, "subnet_segments", subnet_segments)
        if taggings and not isinstance(taggings, list):
            raise TypeError("Expected argument 'taggings' to be a list")
        pulumi.set(__self__, "taggings", taggings)
        if template and not isinstance(template, str):
            raise TypeError("Expected argument 'template' to be a str")
        pulumi.set(__self__, "template", template)
        if tenant and not isinstance(tenant, str):
            raise TypeError("Expected argument 'tenant' to be a str")
        pulumi.set(__self__, "tenant", tenant)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter(name="cacheTtl")
    def cache_ttl(self) -> int:
        """
        Minimal TTL of individual IPv6 addresses in FQDN cache.
        """
        return pulumi.get(self, "cache_ttl")

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def country(self) -> str:
        """
        IPv6 addresses associated to a specific country.
        """
        return pulumi.get(self, "country")

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> str:
        """
        Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        """
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter(name="endMac")
    def end_mac(self) -> str:
        """
        Last MAC address in the range.
        """
        return pulumi.get(self, "end_mac")

    @property
    @pulumi.getter(name="epgName")
    def epg_name(self) -> str:
        """
        Endpoint group name.
        """
        return pulumi.get(self, "epg_name")

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> str:
        """
        Security Fabric global object setting.
        """
        return pulumi.get(self, "fabric_object")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        Fully qualified domain name.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        Host Address.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> str:
        """
        Host type.
        """
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip6(self) -> str:
        """
        IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
        """
        return pulumi.get(self, "ip6")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetAddress6ListResult']:
        """
        IP address list. The structure of `list` block is documented below.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def macaddrs(self) -> Sequence['outputs.GetAddress6MacaddrResult']:
        """
        MAC address ranges <start>[-<end>] separated by space.
        """
        return pulumi.get(self, "macaddrs")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objId")
    def obj_id(self) -> str:
        """
        Object ID for NSX.
        """
        return pulumi.get(self, "obj_id")

    @property
    @pulumi.getter(name="routeTag")
    def route_tag(self) -> int:
        """
        route-tag address.
        """
        return pulumi.get(self, "route_tag")

    @property
    @pulumi.getter
    def sdn(self) -> str:
        """
        SDN.
        """
        return pulumi.get(self, "sdn")

    @property
    @pulumi.getter(name="sdnTag")
    def sdn_tag(self) -> str:
        """
        SDN Tag.
        """
        return pulumi.get(self, "sdn_tag")

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> str:
        """
        First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        """
        return pulumi.get(self, "start_ip")

    @property
    @pulumi.getter(name="startMac")
    def start_mac(self) -> str:
        """
        First MAC address in the range.
        """
        return pulumi.get(self, "start_mac")

    @property
    @pulumi.getter(name="subnetSegments")
    def subnet_segments(self) -> Sequence['outputs.GetAddress6SubnetSegmentResult']:
        """
        IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
        """
        return pulumi.get(self, "subnet_segments")

    @property
    @pulumi.getter
    def taggings(self) -> Sequence['outputs.GetAddress6TaggingResult']:
        """
        Config object tagging The structure of `tagging` block is documented below.
        """
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def template(self) -> str:
        """
        IPv6 address template.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def tenant(self) -> str:
        """
        Tenant.
        """
        return pulumi.get(self, "tenant")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Subnet segment type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        Enable/disable the visibility of the object in the GUI.
        """
        return pulumi.get(self, "visibility")


class AwaitableGetAddress6Result(GetAddress6Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddress6Result(
            cache_ttl=self.cache_ttl,
            color=self.color,
            comment=self.comment,
            country=self.country,
            end_ip=self.end_ip,
            end_mac=self.end_mac,
            epg_name=self.epg_name,
            fabric_object=self.fabric_object,
            fqdn=self.fqdn,
            host=self.host,
            host_type=self.host_type,
            id=self.id,
            ip6=self.ip6,
            lists=self.lists,
            macaddrs=self.macaddrs,
            name=self.name,
            obj_id=self.obj_id,
            route_tag=self.route_tag,
            sdn=self.sdn,
            sdn_tag=self.sdn_tag,
            start_ip=self.start_ip,
            start_mac=self.start_mac,
            subnet_segments=self.subnet_segments,
            taggings=self.taggings,
            template=self.template,
            tenant=self.tenant,
            type=self.type,
            uuid=self.uuid,
            vdomparam=self.vdomparam,
            visibility=self.visibility)


def get_address6(name: Optional[str] = None,
                 vdomparam: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddress6Result:
    """
    Use this data source to get information on an fortios firewall address6


    :param str name: Specify the name of the desired firewall address6.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:firewall/getAddress6:getAddress6', __args__, opts=opts, typ=GetAddress6Result).value

    return AwaitableGetAddress6Result(
        cache_ttl=pulumi.get(__ret__, 'cache_ttl'),
        color=pulumi.get(__ret__, 'color'),
        comment=pulumi.get(__ret__, 'comment'),
        country=pulumi.get(__ret__, 'country'),
        end_ip=pulumi.get(__ret__, 'end_ip'),
        end_mac=pulumi.get(__ret__, 'end_mac'),
        epg_name=pulumi.get(__ret__, 'epg_name'),
        fabric_object=pulumi.get(__ret__, 'fabric_object'),
        fqdn=pulumi.get(__ret__, 'fqdn'),
        host=pulumi.get(__ret__, 'host'),
        host_type=pulumi.get(__ret__, 'host_type'),
        id=pulumi.get(__ret__, 'id'),
        ip6=pulumi.get(__ret__, 'ip6'),
        lists=pulumi.get(__ret__, 'lists'),
        macaddrs=pulumi.get(__ret__, 'macaddrs'),
        name=pulumi.get(__ret__, 'name'),
        obj_id=pulumi.get(__ret__, 'obj_id'),
        route_tag=pulumi.get(__ret__, 'route_tag'),
        sdn=pulumi.get(__ret__, 'sdn'),
        sdn_tag=pulumi.get(__ret__, 'sdn_tag'),
        start_ip=pulumi.get(__ret__, 'start_ip'),
        start_mac=pulumi.get(__ret__, 'start_mac'),
        subnet_segments=pulumi.get(__ret__, 'subnet_segments'),
        taggings=pulumi.get(__ret__, 'taggings'),
        template=pulumi.get(__ret__, 'template'),
        tenant=pulumi.get(__ret__, 'tenant'),
        type=pulumi.get(__ret__, 'type'),
        uuid=pulumi.get(__ret__, 'uuid'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'),
        visibility=pulumi.get(__ret__, 'visibility'))


@_utilities.lift_output_func(get_address6)
def get_address6_output(name: Optional[pulumi.Input[str]] = None,
                        vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAddress6Result]:
    """
    Use this data source to get information on an fortios firewall address6


    :param str name: Specify the name of the desired firewall address6.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
