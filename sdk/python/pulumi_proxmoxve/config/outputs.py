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
    'Ssh',
    'SshNode',
]

@pulumi.output_type
class Ssh(dict):
    def __init__(__self__, *,
                 agent: Optional[bool] = None,
                 agent_socket: Optional[str] = None,
                 nodes: Optional[Sequence['outputs.SshNode']] = None,
                 password: Optional[str] = None,
                 username: Optional[str] = None):
        if agent is not None:
            pulumi.set(__self__, "agent", agent)
        if agent_socket is not None:
            pulumi.set(__self__, "agent_socket", agent_socket)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def agent(self) -> Optional[bool]:
        return pulumi.get(self, "agent")

    @property
    @pulumi.getter(name="agentSocket")
    def agent_socket(self) -> Optional[str]:
        return pulumi.get(self, "agent_socket")

    @property
    @pulumi.getter
    def nodes(self) -> Optional[Sequence['outputs.SshNode']]:
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")


@pulumi.output_type
class SshNode(dict):
    def __init__(__self__, *,
                 address: str,
                 name: str,
                 port: Optional[int] = None):
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        return pulumi.get(self, "port")


