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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, auth_proto=None, auth_pwd=None, events=None, ha_direct=None, id=None, mib_view=None, name=None, notify_hosts=None, notify_hosts6=None, priv_proto=None, priv_pwd=None, queries=None, query_port=None, security_level=None, source_ip=None, source_ipv6=None, status=None, trap_lport=None, trap_rport=None, trap_status=None, vdomparam=None, vdoms=None):
        if auth_proto and not isinstance(auth_proto, str):
            raise TypeError("Expected argument 'auth_proto' to be a str")
        pulumi.set(__self__, "auth_proto", auth_proto)
        if auth_pwd and not isinstance(auth_pwd, str):
            raise TypeError("Expected argument 'auth_pwd' to be a str")
        pulumi.set(__self__, "auth_pwd", auth_pwd)
        if events and not isinstance(events, str):
            raise TypeError("Expected argument 'events' to be a str")
        pulumi.set(__self__, "events", events)
        if ha_direct and not isinstance(ha_direct, str):
            raise TypeError("Expected argument 'ha_direct' to be a str")
        pulumi.set(__self__, "ha_direct", ha_direct)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mib_view and not isinstance(mib_view, str):
            raise TypeError("Expected argument 'mib_view' to be a str")
        pulumi.set(__self__, "mib_view", mib_view)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notify_hosts and not isinstance(notify_hosts, str):
            raise TypeError("Expected argument 'notify_hosts' to be a str")
        pulumi.set(__self__, "notify_hosts", notify_hosts)
        if notify_hosts6 and not isinstance(notify_hosts6, str):
            raise TypeError("Expected argument 'notify_hosts6' to be a str")
        pulumi.set(__self__, "notify_hosts6", notify_hosts6)
        if priv_proto and not isinstance(priv_proto, str):
            raise TypeError("Expected argument 'priv_proto' to be a str")
        pulumi.set(__self__, "priv_proto", priv_proto)
        if priv_pwd and not isinstance(priv_pwd, str):
            raise TypeError("Expected argument 'priv_pwd' to be a str")
        pulumi.set(__self__, "priv_pwd", priv_pwd)
        if queries and not isinstance(queries, str):
            raise TypeError("Expected argument 'queries' to be a str")
        pulumi.set(__self__, "queries", queries)
        if query_port and not isinstance(query_port, int):
            raise TypeError("Expected argument 'query_port' to be a int")
        pulumi.set(__self__, "query_port", query_port)
        if security_level and not isinstance(security_level, str):
            raise TypeError("Expected argument 'security_level' to be a str")
        pulumi.set(__self__, "security_level", security_level)
        if source_ip and not isinstance(source_ip, str):
            raise TypeError("Expected argument 'source_ip' to be a str")
        pulumi.set(__self__, "source_ip", source_ip)
        if source_ipv6 and not isinstance(source_ipv6, str):
            raise TypeError("Expected argument 'source_ipv6' to be a str")
        pulumi.set(__self__, "source_ipv6", source_ipv6)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if trap_lport and not isinstance(trap_lport, int):
            raise TypeError("Expected argument 'trap_lport' to be a int")
        pulumi.set(__self__, "trap_lport", trap_lport)
        if trap_rport and not isinstance(trap_rport, int):
            raise TypeError("Expected argument 'trap_rport' to be a int")
        pulumi.set(__self__, "trap_rport", trap_rport)
        if trap_status and not isinstance(trap_status, str):
            raise TypeError("Expected argument 'trap_status' to be a str")
        pulumi.set(__self__, "trap_status", trap_status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms and not isinstance(vdoms, list):
            raise TypeError("Expected argument 'vdoms' to be a list")
        pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter(name="authProto")
    def auth_proto(self) -> str:
        """
        Authentication protocol.
        """
        return pulumi.get(self, "auth_proto")

    @property
    @pulumi.getter(name="authPwd")
    def auth_pwd(self) -> str:
        """
        Password for authentication protocol.
        """
        return pulumi.get(self, "auth_pwd")

    @property
    @pulumi.getter
    def events(self) -> str:
        """
        SNMP notifications (traps) to send.
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter(name="haDirect")
    def ha_direct(self) -> str:
        """
        Enable/disable direct management of HA cluster members.
        """
        return pulumi.get(self, "ha_direct")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mibView")
    def mib_view(self) -> str:
        """
        SNMP access control MIB view.
        """
        return pulumi.get(self, "mib_view")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        VDOM name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyHosts")
    def notify_hosts(self) -> str:
        """
        SNMP managers to send notifications (traps) to.
        """
        return pulumi.get(self, "notify_hosts")

    @property
    @pulumi.getter(name="notifyHosts6")
    def notify_hosts6(self) -> str:
        """
        IPv6 SNMP managers to send notifications (traps) to.
        """
        return pulumi.get(self, "notify_hosts6")

    @property
    @pulumi.getter(name="privProto")
    def priv_proto(self) -> str:
        """
        Privacy (encryption) protocol.
        """
        return pulumi.get(self, "priv_proto")

    @property
    @pulumi.getter(name="privPwd")
    def priv_pwd(self) -> str:
        """
        Password for privacy (encryption) protocol.
        """
        return pulumi.get(self, "priv_pwd")

    @property
    @pulumi.getter
    def queries(self) -> str:
        """
        Enable/disable SNMP queries for this user.
        """
        return pulumi.get(self, "queries")

    @property
    @pulumi.getter(name="queryPort")
    def query_port(self) -> int:
        """
        SNMPv3 query port (default = 161).
        """
        return pulumi.get(self, "query_port")

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> str:
        """
        Security level for message authentication and encryption.
        """
        return pulumi.get(self, "security_level")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> str:
        """
        Source IP for SNMP trap.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sourceIpv6")
    def source_ipv6(self) -> str:
        """
        Source IPv6 for SNMP trap.
        """
        return pulumi.get(self, "source_ipv6")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable this SNMP user.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="trapLport")
    def trap_lport(self) -> int:
        """
        SNMPv3 local trap port (default = 162).
        """
        return pulumi.get(self, "trap_lport")

    @property
    @pulumi.getter(name="trapRport")
    def trap_rport(self) -> int:
        """
        SNMPv3 trap remote port (default = 162).
        """
        return pulumi.get(self, "trap_rport")

    @property
    @pulumi.getter(name="trapStatus")
    def trap_status(self) -> str:
        """
        Enable/disable traps for this SNMP user.
        """
        return pulumi.get(self, "trap_status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vdoms(self) -> Sequence['outputs.GetUserVdomResult']:
        """
        SNMP access control VDOMs. The structure of `vdoms` block is documented below.
        """
        return pulumi.get(self, "vdoms")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            auth_proto=self.auth_proto,
            auth_pwd=self.auth_pwd,
            events=self.events,
            ha_direct=self.ha_direct,
            id=self.id,
            mib_view=self.mib_view,
            name=self.name,
            notify_hosts=self.notify_hosts,
            notify_hosts6=self.notify_hosts6,
            priv_proto=self.priv_proto,
            priv_pwd=self.priv_pwd,
            queries=self.queries,
            query_port=self.query_port,
            security_level=self.security_level,
            source_ip=self.source_ip,
            source_ipv6=self.source_ipv6,
            status=self.status,
            trap_lport=self.trap_lport,
            trap_rport=self.trap_rport,
            trap_status=self.trap_status,
            vdomparam=self.vdomparam,
            vdoms=self.vdoms)


def get_user(name: Optional[str] = None,
             vdomparam: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Use this data source to get information on an fortios systemsnmp user


    :param str name: Specify the name of the desired systemsnmp user.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:systemsnmp/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        auth_proto=__ret__.auth_proto,
        auth_pwd=__ret__.auth_pwd,
        events=__ret__.events,
        ha_direct=__ret__.ha_direct,
        id=__ret__.id,
        mib_view=__ret__.mib_view,
        name=__ret__.name,
        notify_hosts=__ret__.notify_hosts,
        notify_hosts6=__ret__.notify_hosts6,
        priv_proto=__ret__.priv_proto,
        priv_pwd=__ret__.priv_pwd,
        queries=__ret__.queries,
        query_port=__ret__.query_port,
        security_level=__ret__.security_level,
        source_ip=__ret__.source_ip,
        source_ipv6=__ret__.source_ipv6,
        status=__ret__.status,
        trap_lport=__ret__.trap_lport,
        trap_rport=__ret__.trap_rport,
        trap_status=__ret__.trap_status,
        vdomparam=__ret__.vdomparam,
        vdoms=__ret__.vdoms)


@_utilities.lift_output_func(get_user)
def get_user_output(name: Optional[pulumi.Input[str]] = None,
                    vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Use this data source to get information on an fortios systemsnmp user


    :param str name: Specify the name of the desired systemsnmp user.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
