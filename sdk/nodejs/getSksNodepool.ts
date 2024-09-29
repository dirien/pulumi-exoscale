// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getSksNodepool(args: GetSksNodepoolArgs, opts?: pulumi.InvokeOptions): Promise<GetSksNodepoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getSksNodepool:getSksNodepool", {
        "antiAffinityGroupIds": args.antiAffinityGroupIds,
        "clusterId": args.clusterId,
        "createdAt": args.createdAt,
        "deployTargetId": args.deployTargetId,
        "description": args.description,
        "diskSize": args.diskSize,
        "id": args.id,
        "instancePoolId": args.instancePoolId,
        "instancePrefix": args.instancePrefix,
        "instanceType": args.instanceType,
        "kubeletImageGcs": args.kubeletImageGcs,
        "labels": args.labels,
        "name": args.name,
        "privateNetworkIds": args.privateNetworkIds,
        "securityGroupIds": args.securityGroupIds,
        "size": args.size,
        "state": args.state,
        "storageLvm": args.storageLvm,
        "taints": args.taints,
        "templateId": args.templateId,
        "version": args.version,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSksNodepool.
 */
export interface GetSksNodepoolArgs {
    /**
     * A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
     */
    antiAffinityGroupIds?: string[];
    clusterId: string;
    /**
     * The pool creation date.
     */
    createdAt?: string;
    /**
     * A deploy target ID.
     */
    deployTargetId?: string;
    /**
     * A free-form text describing the pool.
     */
    description?: string;
    /**
     * The managed instances disk size (GiB; default: `50`).
     */
    diskSize?: number;
    /**
     * The ID of this resource.
     */
    id?: string;
    /**
     * The underlying exoscale*instance*pool ID.
     */
    instancePoolId?: string;
    /**
     * The string used to prefix the managed instances name (default `pool`).
     */
    instancePrefix?: string;
    /**
     * The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
     */
    instanceType?: string;
    /**
     * Configuration for this nodepool's kubelet image garbage collector
     */
    kubeletImageGcs?: inputs.GetSksNodepoolKubeletImageGc[];
    /**
     * A map of key/value labels.
     */
    labels?: {[key: string]: string};
    name?: string;
    /**
     * A list of exoscale*private*network (IDs) to be attached to the managed instances.
     */
    privateNetworkIds?: string[];
    /**
     * A list of exoscale*security*group (IDs) to be attached to the managed instances.
     */
    securityGroupIds?: string[];
    size?: number;
    /**
     * The current pool state.
     */
    state?: string;
    /**
     * Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
     */
    storageLvm?: boolean;
    /**
     * A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) ('taints = { \n\n = "\n\n:\n\n" }').
     */
    taints?: {[key: string]: string};
    /**
     * The managed instances template ID.
     */
    templateId?: string;
    /**
     * The managed instances version.
     */
    version?: string;
    zone: string;
}

/**
 * A collection of values returned by getSksNodepool.
 */
export interface GetSksNodepoolResult {
    /**
     * A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
     */
    readonly antiAffinityGroupIds?: string[];
    readonly clusterId: string;
    /**
     * The pool creation date.
     */
    readonly createdAt: string;
    /**
     * A deploy target ID.
     */
    readonly deployTargetId?: string;
    /**
     * A free-form text describing the pool.
     */
    readonly description?: string;
    /**
     * The managed instances disk size (GiB; default: `50`).
     */
    readonly diskSize?: number;
    /**
     * The ID of this resource.
     */
    readonly id?: string;
    /**
     * The underlying exoscale*instance*pool ID.
     */
    readonly instancePoolId: string;
    /**
     * The string used to prefix the managed instances name (default `pool`).
     */
    readonly instancePrefix?: string;
    /**
     * The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
     */
    readonly instanceType?: string;
    /**
     * Configuration for this nodepool's kubelet image garbage collector
     */
    readonly kubeletImageGcs?: outputs.GetSksNodepoolKubeletImageGc[];
    /**
     * A map of key/value labels.
     */
    readonly labels?: {[key: string]: string};
    readonly name?: string;
    /**
     * A list of exoscale*private*network (IDs) to be attached to the managed instances.
     */
    readonly privateNetworkIds?: string[];
    /**
     * A list of exoscale*security*group (IDs) to be attached to the managed instances.
     */
    readonly securityGroupIds?: string[];
    readonly size?: number;
    /**
     * The current pool state.
     */
    readonly state: string;
    /**
     * Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
     */
    readonly storageLvm?: boolean;
    /**
     * A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) ('taints = { \n\n = "\n\n:\n\n" }').
     */
    readonly taints?: {[key: string]: string};
    /**
     * The managed instances template ID.
     */
    readonly templateId: string;
    /**
     * The managed instances version.
     */
    readonly version: string;
    readonly zone: string;
}
export function getSksNodepoolOutput(args: GetSksNodepoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSksNodepoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getSksNodepool:getSksNodepool", {
        "antiAffinityGroupIds": args.antiAffinityGroupIds,
        "clusterId": args.clusterId,
        "createdAt": args.createdAt,
        "deployTargetId": args.deployTargetId,
        "description": args.description,
        "diskSize": args.diskSize,
        "id": args.id,
        "instancePoolId": args.instancePoolId,
        "instancePrefix": args.instancePrefix,
        "instanceType": args.instanceType,
        "kubeletImageGcs": args.kubeletImageGcs,
        "labels": args.labels,
        "name": args.name,
        "privateNetworkIds": args.privateNetworkIds,
        "securityGroupIds": args.securityGroupIds,
        "size": args.size,
        "state": args.state,
        "storageLvm": args.storageLvm,
        "taints": args.taints,
        "templateId": args.templateId,
        "version": args.version,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSksNodepool.
 */
export interface GetSksNodepoolOutputArgs {
    /**
     * A list of exoscale*anti*affinity_group (IDs) to be attached to the managed instances.
     */
    antiAffinityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    clusterId: pulumi.Input<string>;
    /**
     * The pool creation date.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A deploy target ID.
     */
    deployTargetId?: pulumi.Input<string>;
    /**
     * A free-form text describing the pool.
     */
    description?: pulumi.Input<string>;
    /**
     * The managed instances disk size (GiB; default: `50`).
     */
    diskSize?: pulumi.Input<number>;
    /**
     * The ID of this resource.
     */
    id?: pulumi.Input<string>;
    /**
     * The underlying exoscale*instance*pool ID.
     */
    instancePoolId?: pulumi.Input<string>;
    /**
     * The string used to prefix the managed instances name (default `pool`).
     */
    instancePrefix?: pulumi.Input<string>;
    /**
     * The managed compute instances type (`<family>.<size>`, e.g. `standard.medium`; use the [Exoscale CLI](https://github.com/exoscale/cli/) - `exo compute instance-type list` - for the list of available types).
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Configuration for this nodepool's kubelet image garbage collector
     */
    kubeletImageGcs?: pulumi.Input<pulumi.Input<inputs.GetSksNodepoolKubeletImageGcArgs>[]>;
    /**
     * A map of key/value labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    /**
     * A list of exoscale*private*network (IDs) to be attached to the managed instances.
     */
    privateNetworkIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of exoscale*security*group (IDs) to be attached to the managed instances.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    size?: pulumi.Input<number>;
    /**
     * The current pool state.
     */
    state?: pulumi.Input<string>;
    /**
     * Create nodes with non-standard partitioning for persistent storage (requires min 100G of disk space) (may only be set at creation time).
     */
    storageLvm?: pulumi.Input<boolean>;
    /**
     * A map of key/value Kubernetes [taints](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) ('taints = { \n\n = "\n\n:\n\n" }').
     */
    taints?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed instances template ID.
     */
    templateId?: pulumi.Input<string>;
    /**
     * The managed instances version.
     */
    version?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
