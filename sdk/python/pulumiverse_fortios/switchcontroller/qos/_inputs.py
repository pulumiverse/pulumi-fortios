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
    'IpdscpmapMapArgs',
    'QueuepolicyCosQueueArgs',
]

@pulumi.input_type
class IpdscpmapMapArgs:
    def __init__(__self__, *,
                 cos_queue: Optional[pulumi.Input[int]] = None,
                 diffserv: Optional[pulumi.Input[str]] = None,
                 ip_precedence: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] cos_queue: COS queue number.
        :param pulumi.Input[str] diffserv: Differentiated service. Valid values: `CS0`, `CS1`, `AF11`, `AF12`, `AF13`, `CS2`, `AF21`, `AF22`, `AF23`, `CS3`, `AF31`, `AF32`, `AF33`, `CS4`, `AF41`, `AF42`, `AF43`, `CS5`, `EF`, `CS6`, `CS7`.
        :param pulumi.Input[str] ip_precedence: IP Precedence. Valid values: `network-control`, `internetwork-control`, `critic-ecp`, `flashoverride`, `flash`, `immediate`, `priority`, `routine`.
        :param pulumi.Input[str] name: Dscp mapping entry name.
        :param pulumi.Input[str] value: Raw values of DSCP (0 - 63).
        """
        if cos_queue is not None:
            pulumi.set(__self__, "cos_queue", cos_queue)
        if diffserv is not None:
            pulumi.set(__self__, "diffserv", diffserv)
        if ip_precedence is not None:
            pulumi.set(__self__, "ip_precedence", ip_precedence)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="cosQueue")
    def cos_queue(self) -> Optional[pulumi.Input[int]]:
        """
        COS queue number.
        """
        return pulumi.get(self, "cos_queue")

    @cos_queue.setter
    def cos_queue(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cos_queue", value)

    @property
    @pulumi.getter
    def diffserv(self) -> Optional[pulumi.Input[str]]:
        """
        Differentiated service. Valid values: `CS0`, `CS1`, `AF11`, `AF12`, `AF13`, `CS2`, `AF21`, `AF22`, `AF23`, `CS3`, `AF31`, `AF32`, `AF33`, `CS4`, `AF41`, `AF42`, `AF43`, `CS5`, `EF`, `CS6`, `CS7`.
        """
        return pulumi.get(self, "diffserv")

    @diffserv.setter
    def diffserv(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "diffserv", value)

    @property
    @pulumi.getter(name="ipPrecedence")
    def ip_precedence(self) -> Optional[pulumi.Input[str]]:
        """
        IP Precedence. Valid values: `network-control`, `internetwork-control`, `critic-ecp`, `flashoverride`, `flash`, `immediate`, `priority`, `routine`.
        """
        return pulumi.get(self, "ip_precedence")

    @ip_precedence.setter
    def ip_precedence(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_precedence", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Dscp mapping entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Raw values of DSCP (0 - 63).
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class QueuepolicyCosQueueArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 drop_policy: Optional[pulumi.Input[str]] = None,
                 ecn: Optional[pulumi.Input[str]] = None,
                 max_rate: Optional[pulumi.Input[int]] = None,
                 max_rate_percent: Optional[pulumi.Input[int]] = None,
                 min_rate: Optional[pulumi.Input[int]] = None,
                 min_rate_percent: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] description: Description of the COS queue.
        :param pulumi.Input[str] drop_policy: COS queue drop policy. Valid values: `taildrop`, `weighted-random-early-detection`.
        :param pulumi.Input[str] ecn: Enable/disable ECN packet marking to drop eligible packets. Valid values: `disable`, `enable`.
        :param pulumi.Input[int] max_rate: Maximum rate (0 - 4294967295 kbps, 0 to disable).
        :param pulumi.Input[int] max_rate_percent: Maximum rate (% of link speed).
        :param pulumi.Input[int] min_rate: Minimum rate (0 - 4294967295 kbps, 0 to disable).
        :param pulumi.Input[int] min_rate_percent: Minimum rate (% of link speed).
        :param pulumi.Input[str] name: Cos queue ID.
        :param pulumi.Input[int] weight: Weight of weighted round robin scheduling.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if drop_policy is not None:
            pulumi.set(__self__, "drop_policy", drop_policy)
        if ecn is not None:
            pulumi.set(__self__, "ecn", ecn)
        if max_rate is not None:
            pulumi.set(__self__, "max_rate", max_rate)
        if max_rate_percent is not None:
            pulumi.set(__self__, "max_rate_percent", max_rate_percent)
        if min_rate is not None:
            pulumi.set(__self__, "min_rate", min_rate)
        if min_rate_percent is not None:
            pulumi.set(__self__, "min_rate_percent", min_rate_percent)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the COS queue.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dropPolicy")
    def drop_policy(self) -> Optional[pulumi.Input[str]]:
        """
        COS queue drop policy. Valid values: `taildrop`, `weighted-random-early-detection`.
        """
        return pulumi.get(self, "drop_policy")

    @drop_policy.setter
    def drop_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "drop_policy", value)

    @property
    @pulumi.getter
    def ecn(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable ECN packet marking to drop eligible packets. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "ecn")

    @ecn.setter
    def ecn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ecn", value)

    @property
    @pulumi.getter(name="maxRate")
    def max_rate(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum rate (0 - 4294967295 kbps, 0 to disable).
        """
        return pulumi.get(self, "max_rate")

    @max_rate.setter
    def max_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_rate", value)

    @property
    @pulumi.getter(name="maxRatePercent")
    def max_rate_percent(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum rate (% of link speed).
        """
        return pulumi.get(self, "max_rate_percent")

    @max_rate_percent.setter
    def max_rate_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_rate_percent", value)

    @property
    @pulumi.getter(name="minRate")
    def min_rate(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum rate (0 - 4294967295 kbps, 0 to disable).
        """
        return pulumi.get(self, "min_rate")

    @min_rate.setter
    def min_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_rate", value)

    @property
    @pulumi.getter(name="minRatePercent")
    def min_rate_percent(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum rate (% of link speed).
        """
        return pulumi.get(self, "min_rate_percent")

    @min_rate_percent.setter
    def min_rate_percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_rate_percent", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Cos queue ID.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        Weight of weighted round robin scheduling.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


