# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetGretunnelResult',
    'AwaitableGetGretunnelResult',
    'get_gretunnel',
    'get_gretunnel_output',
]

@pulumi.output_type
class GetGretunnelResult:
    """
    A collection of values returned by getGretunnel.
    """
    def __init__(__self__, checksum_reception=None, checksum_transmission=None, diffservcode=None, dscp_copying=None, id=None, interface=None, ip_version=None, keepalive_failtimes=None, keepalive_interval=None, key_inbound=None, key_outbound=None, local_gw=None, local_gw6=None, name=None, remote_gw=None, remote_gw6=None, sequence_number_reception=None, sequence_number_transmission=None, use_sdwan=None, vdomparam=None):
        if checksum_reception and not isinstance(checksum_reception, str):
            raise TypeError("Expected argument 'checksum_reception' to be a str")
        pulumi.set(__self__, "checksum_reception", checksum_reception)
        if checksum_transmission and not isinstance(checksum_transmission, str):
            raise TypeError("Expected argument 'checksum_transmission' to be a str")
        pulumi.set(__self__, "checksum_transmission", checksum_transmission)
        if diffservcode and not isinstance(diffservcode, str):
            raise TypeError("Expected argument 'diffservcode' to be a str")
        pulumi.set(__self__, "diffservcode", diffservcode)
        if dscp_copying and not isinstance(dscp_copying, str):
            raise TypeError("Expected argument 'dscp_copying' to be a str")
        pulumi.set(__self__, "dscp_copying", dscp_copying)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if ip_version and not isinstance(ip_version, str):
            raise TypeError("Expected argument 'ip_version' to be a str")
        pulumi.set(__self__, "ip_version", ip_version)
        if keepalive_failtimes and not isinstance(keepalive_failtimes, int):
            raise TypeError("Expected argument 'keepalive_failtimes' to be a int")
        pulumi.set(__self__, "keepalive_failtimes", keepalive_failtimes)
        if keepalive_interval and not isinstance(keepalive_interval, int):
            raise TypeError("Expected argument 'keepalive_interval' to be a int")
        pulumi.set(__self__, "keepalive_interval", keepalive_interval)
        if key_inbound and not isinstance(key_inbound, int):
            raise TypeError("Expected argument 'key_inbound' to be a int")
        pulumi.set(__self__, "key_inbound", key_inbound)
        if key_outbound and not isinstance(key_outbound, int):
            raise TypeError("Expected argument 'key_outbound' to be a int")
        pulumi.set(__self__, "key_outbound", key_outbound)
        if local_gw and not isinstance(local_gw, str):
            raise TypeError("Expected argument 'local_gw' to be a str")
        pulumi.set(__self__, "local_gw", local_gw)
        if local_gw6 and not isinstance(local_gw6, str):
            raise TypeError("Expected argument 'local_gw6' to be a str")
        pulumi.set(__self__, "local_gw6", local_gw6)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if remote_gw and not isinstance(remote_gw, str):
            raise TypeError("Expected argument 'remote_gw' to be a str")
        pulumi.set(__self__, "remote_gw", remote_gw)
        if remote_gw6 and not isinstance(remote_gw6, str):
            raise TypeError("Expected argument 'remote_gw6' to be a str")
        pulumi.set(__self__, "remote_gw6", remote_gw6)
        if sequence_number_reception and not isinstance(sequence_number_reception, str):
            raise TypeError("Expected argument 'sequence_number_reception' to be a str")
        pulumi.set(__self__, "sequence_number_reception", sequence_number_reception)
        if sequence_number_transmission and not isinstance(sequence_number_transmission, str):
            raise TypeError("Expected argument 'sequence_number_transmission' to be a str")
        pulumi.set(__self__, "sequence_number_transmission", sequence_number_transmission)
        if use_sdwan and not isinstance(use_sdwan, str):
            raise TypeError("Expected argument 'use_sdwan' to be a str")
        pulumi.set(__self__, "use_sdwan", use_sdwan)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="checksumReception")
    def checksum_reception(self) -> str:
        """
        Enable/disable validating checksums in received GRE packets.
        """
        return pulumi.get(self, "checksum_reception")

    @property
    @pulumi.getter(name="checksumTransmission")
    def checksum_transmission(self) -> str:
        """
        Enable/disable including checksums in transmitted GRE packets.
        """
        return pulumi.get(self, "checksum_transmission")

    @property
    @pulumi.getter
    def diffservcode(self) -> str:
        """
        DiffServ setting to be applied to GRE tunnel outer IP header.
        """
        return pulumi.get(self, "diffservcode")

    @property
    @pulumi.getter(name="dscpCopying")
    def dscp_copying(self) -> str:
        """
        Enable/disable DSCP copying.
        """
        return pulumi.get(self, "dscp_copying")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        Interface name.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> str:
        """
        IP version to use for VPN interface.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="keepaliveFailtimes")
    def keepalive_failtimes(self) -> int:
        """
        Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
        """
        return pulumi.get(self, "keepalive_failtimes")

    @property
    @pulumi.getter(name="keepaliveInterval")
    def keepalive_interval(self) -> int:
        """
        Keepalive message interval (0 - 32767, 0 = disabled).
        """
        return pulumi.get(self, "keepalive_interval")

    @property
    @pulumi.getter(name="keyInbound")
    def key_inbound(self) -> int:
        """
        Require received GRE packets contain this key (0 - 4294967295).
        """
        return pulumi.get(self, "key_inbound")

    @property
    @pulumi.getter(name="keyOutbound")
    def key_outbound(self) -> int:
        """
        Include this key in transmitted GRE packets (0 - 4294967295).
        """
        return pulumi.get(self, "key_outbound")

    @property
    @pulumi.getter(name="localGw")
    def local_gw(self) -> str:
        """
        IP address of the local gateway.
        """
        return pulumi.get(self, "local_gw")

    @property
    @pulumi.getter(name="localGw6")
    def local_gw6(self) -> str:
        """
        IPv6 address of the local gateway.
        """
        return pulumi.get(self, "local_gw6")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Tunnel name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteGw")
    def remote_gw(self) -> str:
        """
        IP address of the remote gateway.
        """
        return pulumi.get(self, "remote_gw")

    @property
    @pulumi.getter(name="remoteGw6")
    def remote_gw6(self) -> str:
        """
        IPv6 address of the remote gateway.
        """
        return pulumi.get(self, "remote_gw6")

    @property
    @pulumi.getter(name="sequenceNumberReception")
    def sequence_number_reception(self) -> str:
        """
        Enable/disable validating sequence numbers in received GRE packets.
        """
        return pulumi.get(self, "sequence_number_reception")

    @property
    @pulumi.getter(name="sequenceNumberTransmission")
    def sequence_number_transmission(self) -> str:
        """
        Enable/disable including of sequence numbers in transmitted GRE packets.
        """
        return pulumi.get(self, "sequence_number_transmission")

    @property
    @pulumi.getter(name="useSdwan")
    def use_sdwan(self) -> str:
        """
        Enable/disable use of SD-WAN to reach remote gateway.
        """
        return pulumi.get(self, "use_sdwan")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetGretunnelResult(GetGretunnelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGretunnelResult(
            checksum_reception=self.checksum_reception,
            checksum_transmission=self.checksum_transmission,
            diffservcode=self.diffservcode,
            dscp_copying=self.dscp_copying,
            id=self.id,
            interface=self.interface,
            ip_version=self.ip_version,
            keepalive_failtimes=self.keepalive_failtimes,
            keepalive_interval=self.keepalive_interval,
            key_inbound=self.key_inbound,
            key_outbound=self.key_outbound,
            local_gw=self.local_gw,
            local_gw6=self.local_gw6,
            name=self.name,
            remote_gw=self.remote_gw,
            remote_gw6=self.remote_gw6,
            sequence_number_reception=self.sequence_number_reception,
            sequence_number_transmission=self.sequence_number_transmission,
            use_sdwan=self.use_sdwan,
            vdomparam=self.vdomparam)


def get_gretunnel(name: Optional[str] = None,
                  vdomparam: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGretunnelResult:
    """
    Use this data source to get information on an fortios system gretunnel


    :param str name: Specify the name of the desired system gretunnel.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getGretunnel:getGretunnel', __args__, opts=opts, typ=GetGretunnelResult).value

    return AwaitableGetGretunnelResult(
        checksum_reception=pulumi.get(__ret__, 'checksum_reception'),
        checksum_transmission=pulumi.get(__ret__, 'checksum_transmission'),
        diffservcode=pulumi.get(__ret__, 'diffservcode'),
        dscp_copying=pulumi.get(__ret__, 'dscp_copying'),
        id=pulumi.get(__ret__, 'id'),
        interface=pulumi.get(__ret__, 'interface'),
        ip_version=pulumi.get(__ret__, 'ip_version'),
        keepalive_failtimes=pulumi.get(__ret__, 'keepalive_failtimes'),
        keepalive_interval=pulumi.get(__ret__, 'keepalive_interval'),
        key_inbound=pulumi.get(__ret__, 'key_inbound'),
        key_outbound=pulumi.get(__ret__, 'key_outbound'),
        local_gw=pulumi.get(__ret__, 'local_gw'),
        local_gw6=pulumi.get(__ret__, 'local_gw6'),
        name=pulumi.get(__ret__, 'name'),
        remote_gw=pulumi.get(__ret__, 'remote_gw'),
        remote_gw6=pulumi.get(__ret__, 'remote_gw6'),
        sequence_number_reception=pulumi.get(__ret__, 'sequence_number_reception'),
        sequence_number_transmission=pulumi.get(__ret__, 'sequence_number_transmission'),
        use_sdwan=pulumi.get(__ret__, 'use_sdwan'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_gretunnel)
def get_gretunnel_output(name: Optional[pulumi.Input[str]] = None,
                         vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGretunnelResult]:
    """
    Use this data source to get information on an fortios system gretunnel


    :param str name: Specify the name of the desired system gretunnel.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
