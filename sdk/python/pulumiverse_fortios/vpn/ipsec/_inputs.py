# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'ConcentratorMemberArgs',
    'FecMappingArgs',
    'Phase1BackupGatewayArgs',
    'Phase1CertificateArgs',
    'Phase1Ipv4ExcludeRangeArgs',
    'Phase1Ipv6ExcludeRangeArgs',
    'Phase1interfaceBackupGatewayArgs',
    'Phase1interfaceCertificateArgs',
    'Phase1interfaceIpv4ExcludeRangeArgs',
    'Phase1interfaceIpv6ExcludeRangeArgs',
]

@pulumi.input_type
class ConcentratorMemberArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Member name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Member name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class FecMappingArgs:
    def __init__(__self__, *,
                 bandwidth_bi_threshold: Optional[pulumi.Input[int]] = None,
                 bandwidth_down_threshold: Optional[pulumi.Input[int]] = None,
                 bandwidth_up_threshold: Optional[pulumi.Input[int]] = None,
                 base: Optional[pulumi.Input[int]] = None,
                 latency_threshold: Optional[pulumi.Input[int]] = None,
                 packet_loss_threshold: Optional[pulumi.Input[int]] = None,
                 redundant: Optional[pulumi.Input[int]] = None,
                 seqno: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] bandwidth_bi_threshold: Apply FEC parameters when available bi-bandwidth is >= threshold (kbps, 0 means no threshold).
        :param pulumi.Input[int] bandwidth_down_threshold: Apply FEC parameters when available down bandwidth is >= threshold (kbps, 0 means no threshold).
        :param pulumi.Input[int] bandwidth_up_threshold: Apply FEC parameters when available up bandwidth is >= threshold (kbps, 0 means no threshold).
        :param pulumi.Input[int] base: Number of base FEC packets (1 - 20).
        :param pulumi.Input[int] latency_threshold: Apply FEC parameters when latency is <= threshold (0 means no threshold).
        :param pulumi.Input[int] packet_loss_threshold: Apply FEC parameters when packet loss is >= threshold (0 - 100, 0 means no threshold).
        :param pulumi.Input[int] redundant: Number of redundant FEC packets (1 - 5).
        :param pulumi.Input[int] seqno: Sequence number (1 - 64).
        """
        if bandwidth_bi_threshold is not None:
            pulumi.set(__self__, "bandwidth_bi_threshold", bandwidth_bi_threshold)
        if bandwidth_down_threshold is not None:
            pulumi.set(__self__, "bandwidth_down_threshold", bandwidth_down_threshold)
        if bandwidth_up_threshold is not None:
            pulumi.set(__self__, "bandwidth_up_threshold", bandwidth_up_threshold)
        if base is not None:
            pulumi.set(__self__, "base", base)
        if latency_threshold is not None:
            pulumi.set(__self__, "latency_threshold", latency_threshold)
        if packet_loss_threshold is not None:
            pulumi.set(__self__, "packet_loss_threshold", packet_loss_threshold)
        if redundant is not None:
            pulumi.set(__self__, "redundant", redundant)
        if seqno is not None:
            pulumi.set(__self__, "seqno", seqno)

    @property
    @pulumi.getter(name="bandwidthBiThreshold")
    def bandwidth_bi_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Apply FEC parameters when available bi-bandwidth is >= threshold (kbps, 0 means no threshold).
        """
        return pulumi.get(self, "bandwidth_bi_threshold")

    @bandwidth_bi_threshold.setter
    def bandwidth_bi_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth_bi_threshold", value)

    @property
    @pulumi.getter(name="bandwidthDownThreshold")
    def bandwidth_down_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Apply FEC parameters when available down bandwidth is >= threshold (kbps, 0 means no threshold).
        """
        return pulumi.get(self, "bandwidth_down_threshold")

    @bandwidth_down_threshold.setter
    def bandwidth_down_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth_down_threshold", value)

    @property
    @pulumi.getter(name="bandwidthUpThreshold")
    def bandwidth_up_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Apply FEC parameters when available up bandwidth is >= threshold (kbps, 0 means no threshold).
        """
        return pulumi.get(self, "bandwidth_up_threshold")

    @bandwidth_up_threshold.setter
    def bandwidth_up_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth_up_threshold", value)

    @property
    @pulumi.getter
    def base(self) -> Optional[pulumi.Input[int]]:
        """
        Number of base FEC packets (1 - 20).
        """
        return pulumi.get(self, "base")

    @base.setter
    def base(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "base", value)

    @property
    @pulumi.getter(name="latencyThreshold")
    def latency_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Apply FEC parameters when latency is <= threshold (0 means no threshold).
        """
        return pulumi.get(self, "latency_threshold")

    @latency_threshold.setter
    def latency_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "latency_threshold", value)

    @property
    @pulumi.getter(name="packetLossThreshold")
    def packet_loss_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Apply FEC parameters when packet loss is >= threshold (0 - 100, 0 means no threshold).
        """
        return pulumi.get(self, "packet_loss_threshold")

    @packet_loss_threshold.setter
    def packet_loss_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_loss_threshold", value)

    @property
    @pulumi.getter
    def redundant(self) -> Optional[pulumi.Input[int]]:
        """
        Number of redundant FEC packets (1 - 5).
        """
        return pulumi.get(self, "redundant")

    @redundant.setter
    def redundant(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "redundant", value)

    @property
    @pulumi.getter
    def seqno(self) -> Optional[pulumi.Input[int]]:
        """
        Sequence number (1 - 64).
        """
        return pulumi.get(self, "seqno")

    @seqno.setter
    def seqno(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "seqno", value)


@pulumi.input_type
class Phase1BackupGatewayArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] address: Address of backup gateway.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Address of backup gateway.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)


@pulumi.input_type
class Phase1CertificateArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Certificate name.
               
               The `ipv4_exclude_range` block supports:
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate name.

        The `ipv4_exclude_range` block supports:
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class Phase1Ipv4ExcludeRangeArgs:
    def __init__(__self__, *,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] end_ip: End of IPv6 exclusive range.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] start_ip: Start of IPv6 exclusive range.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        End of IPv6 exclusive range.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Start of IPv6 exclusive range.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)


@pulumi.input_type
class Phase1Ipv6ExcludeRangeArgs:
    def __init__(__self__, *,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] end_ip: End of IPv6 exclusive range.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] start_ip: Start of IPv6 exclusive range.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        End of IPv6 exclusive range.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Start of IPv6 exclusive range.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)


@pulumi.input_type
class Phase1interfaceBackupGatewayArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] address: Address of backup gateway.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Address of backup gateway.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)


@pulumi.input_type
class Phase1interfaceCertificateArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Certificate name.
               
               The `ipv4_exclude_range` block supports:
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate name.

        The `ipv4_exclude_range` block supports:
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class Phase1interfaceIpv4ExcludeRangeArgs:
    def __init__(__self__, *,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] end_ip: End of IPv6 exclusive range.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] start_ip: Start of IPv6 exclusive range.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        End of IPv6 exclusive range.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Start of IPv6 exclusive range.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)


@pulumi.input_type
class Phase1interfaceIpv6ExcludeRangeArgs:
    def __init__(__self__, *,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[int]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] end_ip: End of IPv6 exclusive range.
        :param pulumi.Input[int] id: ID.
        :param pulumi.Input[str] start_ip: Start of IPv6 exclusive range.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        """
        End of IPv6 exclusive range.
        """
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Start of IPv6 exclusive range.
        """
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)


