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

import types

__config__ = pulumi.Config('proxmoxve')


class _ExportableConfig(types.ModuleType):
    @property
    def api_token(self) -> Optional[str]:
        """
        The API token for the Proxmox Virtual Environment API
        """
        return __config__.get('apiToken')

    @property
    def endpoint(self) -> Optional[str]:
        """
        The endpoint for the Proxmox Virtual Environment API
        """
        return __config__.get('endpoint')

    @property
    def insecure(self) -> Optional[bool]:
        """
        Whether to skip the TLS verification step
        """
        return __config__.get_bool('insecure')

    @property
    def otp(self) -> Optional[str]:
        """
        The one-time password for the Proxmox Virtual Environment API
        """
        return __config__.get('otp')

    @property
    def password(self) -> Optional[str]:
        """
        The password for the Proxmox Virtual Environment API
        """
        return __config__.get('password')

    @property
    def ssh(self) -> Optional[str]:
        """
        The SSH connection configuration to a Proxmox node
        """
        return __config__.get('ssh')

    @property
    def username(self) -> Optional[str]:
        """
        The username for the Proxmox Virtual Environment API
        """
        return __config__.get('username')

    @property
    def virtual_environment(self) -> Optional[str]:
        return __config__.get('virtualEnvironment')

