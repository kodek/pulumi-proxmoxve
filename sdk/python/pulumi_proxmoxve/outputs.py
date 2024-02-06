# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'HostsEntry',
    'GetHostsEntryResult',
]

@pulumi.output_type
class HostsEntry(dict):
    def __init__(__self__, *,
                 address: str,
                 hostnames: Sequence[str]):
        """
        :param str address: The IP address.
        :param Sequence[str] hostnames: The hostnames.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "hostnames", hostnames)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The IP address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def hostnames(self) -> Sequence[str]:
        """
        The hostnames.
        """
        return pulumi.get(self, "hostnames")


@pulumi.output_type
class GetHostsEntryResult(dict):
    def __init__(__self__, *,
                 address: str,
                 hostnames: Sequence[str]):
        """
        :param str address: The address
        :param Sequence[str] hostnames: The hostnames associated with each of the IP addresses.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "hostnames", hostnames)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The address
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def hostnames(self) -> Sequence[str]:
        """
        The hostnames associated with each of the IP addresses.
        """
        return pulumi.get(self, "hostnames")


