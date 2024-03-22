# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'KeywordWord',
    'ProfileFilter',
    'ProfileFortiguardCategory',
    'ProfileFortiguardCategoryFilter',
    'YoutubechannelfilterEntry',
]

@pulumi.output_type
class KeywordWord(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "patternType":
            suggest = "pattern_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KeywordWord. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KeywordWord.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KeywordWord.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 comment: Optional[str] = None,
                 name: Optional[str] = None,
                 pattern_type: Optional[str] = None,
                 status: Optional[str] = None):
        """
        :param str comment: Comment.
        :param str name: Name.
        :param str pattern_type: Pattern type. Valid values: `wildcard`, `regex`.
        :param str status: Enable(consider)/disable(ignore) this keyword. Valid values: `enable`, `disable`.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pattern_type is not None:
            pulumi.set(__self__, "pattern_type", pattern_type)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def comment(self) -> Optional[str]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="patternType")
    def pattern_type(self) -> Optional[str]:
        """
        Pattern type. Valid values: `wildcard`, `regex`.
        """
        return pulumi.get(self, "pattern_type")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Enable(consider)/disable(ignore) this keyword. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ProfileFilter(dict):
    def __init__(__self__, *,
                 action: Optional[str] = None,
                 category: Optional[str] = None,
                 channel: Optional[str] = None,
                 comment: Optional[str] = None,
                 id: Optional[int] = None,
                 keyword: Optional[int] = None,
                 log: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param str action: VideoFilter action. Valid values: `allow`, `monitor`, `block`.
        :param str category: FortiGuard category ID.
        :param str channel: Channel ID.
        :param str comment: Comment.
        :param int id: ID.
        :param int keyword: Video filter keyword ID.
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        :param str type: Filter type. Valid values: `category`, `channel`, `title`, `description`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if channel is not None:
            pulumi.set(__self__, "channel", channel)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if keyword is not None:
            pulumi.set(__self__, "keyword", keyword)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        VideoFilter action. Valid values: `allow`, `monitor`, `block`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        """
        FortiGuard category ID.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def channel(self) -> Optional[str]:
        """
        Channel ID.
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter
    def comment(self) -> Optional[str]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keyword(self) -> Optional[int]:
        """
        Video filter keyword ID.
        """
        return pulumi.get(self, "keyword")

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Filter type. Valid values: `category`, `channel`, `title`, `description`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ProfileFortiguardCategory(dict):
    def __init__(__self__, *,
                 filters: Optional[Sequence['outputs.ProfileFortiguardCategoryFilter']] = None):
        """
        :param Sequence['ProfileFortiguardCategoryFilterArgs'] filters: Configure VideoFilter FortiGuard category. The structure of `filters` block is documented below.
        """
        if filters is not None:
            pulumi.set(__self__, "filters", filters)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.ProfileFortiguardCategoryFilter']]:
        """
        Configure VideoFilter FortiGuard category. The structure of `filters` block is documented below.
        """
        return pulumi.get(self, "filters")


@pulumi.output_type
class ProfileFortiguardCategoryFilter(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "categoryId":
            suggest = "category_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileFortiguardCategoryFilter. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileFortiguardCategoryFilter.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileFortiguardCategoryFilter.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 category_id: Optional[int] = None,
                 id: Optional[int] = None,
                 log: Optional[str] = None):
        """
        :param str action: VideoFilter action. Valid values: `allow`, `monitor`, `block`.
        :param int category_id: Category ID.
        :param int id: ID.
        :param str log: Enable/disable logging. Valid values: `enable`, `disable`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if category_id is not None:
            pulumi.set(__self__, "category_id", category_id)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if log is not None:
            pulumi.set(__self__, "log", log)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        VideoFilter action. Valid values: `allow`, `monitor`, `block`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="categoryId")
    def category_id(self) -> Optional[int]:
        """
        Category ID.
        """
        return pulumi.get(self, "category_id")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def log(self) -> Optional[str]:
        """
        Enable/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")


@pulumi.output_type
class YoutubechannelfilterEntry(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "channelId":
            suggest = "channel_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in YoutubechannelfilterEntry. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        YoutubechannelfilterEntry.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        YoutubechannelfilterEntry.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: Optional[str] = None,
                 channel_id: Optional[str] = None,
                 comment: Optional[str] = None,
                 id: Optional[int] = None):
        """
        :param str action: YouTube channel filter action. Valid values: `allow`, `monitor`, `block`.
        :param str channel_id: Channel ID.
        :param str comment: Comment.
        :param int id: ID.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if channel_id is not None:
            pulumi.set(__self__, "channel_id", channel_id)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        YouTube channel filter action. Valid values: `allow`, `monitor`, `block`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="channelId")
    def channel_id(self) -> Optional[str]:
        """
        Channel ID.
        """
        return pulumi.get(self, "channel_id")

    @property
    @pulumi.getter
    def comment(self) -> Optional[str]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        ID.
        """
        return pulumi.get(self, "id")


