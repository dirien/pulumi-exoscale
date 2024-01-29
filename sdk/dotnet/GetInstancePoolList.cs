// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Exoscale
{
    public static class GetInstancePoolList
    {
        /// <summary>
        /// List Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/).
        /// 
        /// Corresponding resource: exoscale_instance_pool.
        /// </summary>
        public static Task<GetInstancePoolListResult> InvokeAsync(GetInstancePoolListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancePoolListResult>("exoscale:index/getInstancePoolList:getInstancePoolList", args ?? new GetInstancePoolListArgs(), options.WithDefaults());

        /// <summary>
        /// List Exoscale [Instance Pools](https://community.exoscale.com/documentation/compute/instance-pools/).
        /// 
        /// Corresponding resource: exoscale_instance_pool.
        /// </summary>
        public static Output<GetInstancePoolListResult> Invoke(GetInstancePoolListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancePoolListResult>("exoscale:index/getInstancePoolList:getInstancePoolList", args ?? new GetInstancePoolListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancePoolListArgs : global::Pulumi.InvokeArgs
    {
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetInstancePoolListArgs()
        {
        }
        public static new GetInstancePoolListArgs Empty => new GetInstancePoolListArgs();
    }

    public sealed class GetInstancePoolListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetInstancePoolListInvokeArgs()
        {
        }
        public static new GetInstancePoolListInvokeArgs Empty => new GetInstancePoolListInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancePoolListResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of exoscale*instance*pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancePoolListPoolResult> Pools;
        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetInstancePoolListResult(
            string id,

            ImmutableArray<Outputs.GetInstancePoolListPoolResult> pools,

            string zone)
        {
            Id = id;
            Pools = pools;
            Zone = zone;
        }
    }
}
