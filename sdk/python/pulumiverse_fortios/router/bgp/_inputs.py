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
    'NeighborConditionalAdvertise6Args',
    'NeighborConditionalAdvertiseArgs',
]

@pulumi.input_type
class NeighborConditionalAdvertise6Args:
    def __init__(__self__, *,
                 advertise_routemap: Optional[pulumi.Input[str]] = None,
                 condition_routemap: Optional[pulumi.Input[str]] = None,
                 condition_type: Optional[pulumi.Input[str]] = None):
        if advertise_routemap is not None:
            pulumi.set(__self__, "advertise_routemap", advertise_routemap)
        if condition_routemap is not None:
            pulumi.set(__self__, "condition_routemap", condition_routemap)
        if condition_type is not None:
            pulumi.set(__self__, "condition_type", condition_type)

    @property
    @pulumi.getter(name="advertiseRoutemap")
    def advertise_routemap(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "advertise_routemap")

    @advertise_routemap.setter
    def advertise_routemap(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "advertise_routemap", value)

    @property
    @pulumi.getter(name="conditionRoutemap")
    def condition_routemap(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "condition_routemap")

    @condition_routemap.setter
    def condition_routemap(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_routemap", value)

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "condition_type")

    @condition_type.setter
    def condition_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_type", value)


@pulumi.input_type
class NeighborConditionalAdvertiseArgs:
    def __init__(__self__, *,
                 advertise_routemap: Optional[pulumi.Input[str]] = None,
                 condition_routemap: Optional[pulumi.Input[str]] = None,
                 condition_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] advertise_routemap: Name of advertising route map.
        :param pulumi.Input[str] condition_routemap: Name of condition route map.
        :param pulumi.Input[str] condition_type: Type of condition. Valid values: `exist`, `non-exist`.
        """
        if advertise_routemap is not None:
            pulumi.set(__self__, "advertise_routemap", advertise_routemap)
        if condition_routemap is not None:
            pulumi.set(__self__, "condition_routemap", condition_routemap)
        if condition_type is not None:
            pulumi.set(__self__, "condition_type", condition_type)

    @property
    @pulumi.getter(name="advertiseRoutemap")
    def advertise_routemap(self) -> Optional[pulumi.Input[str]]:
        """
        Name of advertising route map.
        """
        return pulumi.get(self, "advertise_routemap")

    @advertise_routemap.setter
    def advertise_routemap(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "advertise_routemap", value)

    @property
    @pulumi.getter(name="conditionRoutemap")
    def condition_routemap(self) -> Optional[pulumi.Input[str]]:
        """
        Name of condition route map.
        """
        return pulumi.get(self, "condition_routemap")

    @condition_routemap.setter
    def condition_routemap(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_routemap", value)

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of condition. Valid values: `exist`, `non-exist`.
        """
        return pulumi.get(self, "condition_type")

    @condition_type.setter
    def condition_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_type", value)


