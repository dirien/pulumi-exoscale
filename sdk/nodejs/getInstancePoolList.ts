// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * List Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/).
 *
 * Corresponding resource: exoscale_instance_pool.
 */
export function getInstancePoolList(args: GetInstancePoolListArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancePoolListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("exoscale:index/getInstancePoolList:getInstancePoolList", {
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstancePoolList.
 */
export interface GetInstancePoolListArgs {
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: string;
}

/**
 * A collection of values returned by getInstancePoolList.
 */
export interface GetInstancePoolListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of exoscale*instance*pool.
     */
    readonly pools: outputs.GetInstancePoolListPool[];
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    readonly zone: string;
}
/**
 * List Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/).
 *
 * Corresponding resource: exoscale_instance_pool.
 */
export function getInstancePoolListOutput(args: GetInstancePoolListOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetInstancePoolListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("exoscale:index/getInstancePoolList:getInstancePoolList", {
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstancePoolList.
 */
export interface GetInstancePoolListOutputArgs {
    /**
     * The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
     */
    zone: pulumi.Input<string>;
}
