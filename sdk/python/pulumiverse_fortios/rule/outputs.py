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
    'FmwpMetadata',
    'OtdtMetadata',
    'OtdtParameter',
    'OtvpMetadata',
]

@pulumi.output_type
class FmwpMetadata(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None,
                 metaid: Optional[int] = None,
                 valueid: Optional[int] = None):
        """
        :param int id: ID.
        :param int metaid: Meta ID.
        :param int valueid: Value ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if metaid is not None:
            pulumi.set(__self__, "metaid", metaid)
        if valueid is not None:
            pulumi.set(__self__, "valueid", valueid)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metaid(self) -> Optional[int]:
        """
        Meta ID.
        """
        return pulumi.get(self, "metaid")

    @property
    @pulumi.getter
    def valueid(self) -> Optional[int]:
        """
        Value ID.
        """
        return pulumi.get(self, "valueid")


@pulumi.output_type
class OtdtMetadata(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None,
                 metaid: Optional[int] = None,
                 valueid: Optional[int] = None):
        """
        :param int id: ID.
        :param int metaid: Meta ID.
        :param int valueid: Value ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if metaid is not None:
            pulumi.set(__self__, "metaid", metaid)
        if valueid is not None:
            pulumi.set(__self__, "valueid", valueid)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metaid(self) -> Optional[int]:
        """
        Meta ID.
        """
        return pulumi.get(self, "metaid")

    @property
    @pulumi.getter
    def valueid(self) -> Optional[int]:
        """
        Value ID.
        """
        return pulumi.get(self, "valueid")


@pulumi.output_type
class OtdtParameter(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        :param str name: Parameter name.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Parameter name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class OtvpMetadata(dict):
    def __init__(__self__, *,
                 id: Optional[int] = None,
                 metaid: Optional[int] = None,
                 valueid: Optional[int] = None):
        """
        :param int id: ID.
        :param int metaid: Meta ID.
        :param int valueid: Value ID.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if metaid is not None:
            pulumi.set(__self__, "metaid", metaid)
        if valueid is not None:
            pulumi.set(__self__, "valueid", valueid)

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metaid(self) -> Optional[int]:
        """
        Meta ID.
        """
        return pulumi.get(self, "metaid")

    @property
    @pulumi.getter
    def valueid(self) -> Optional[int]:
        """
        Value ID.
        """
        return pulumi.get(self, "valueid")


