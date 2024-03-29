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
    public static class GetAntiAffinityGroup
    {
        /// <summary>
        /// Fetch Exoscale [Anti-Affinity Groups](https://community.exoscale.com/documentation/compute/anti-affinity-groups/) data.
        /// 
        /// Corresponding resource: exoscale_anti_affinity_group.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Exoscale = Pulumi.Exoscale;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myAntiAffinityGroup = Exoscale.GetAntiAffinityGroup.Invoke(new()
        ///     {
        ///         Name = "my-anti-affinity-group",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["myAntiAffinityGroupId"] = myAntiAffinityGroup.Apply(getAntiAffinityGroupResult =&gt; getAntiAffinityGroupResult.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// Please refer to the examples
        /// directory for complete configuration examples.
        /// </summary>
        public static Task<GetAntiAffinityGroupResult> InvokeAsync(GetAntiAffinityGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAntiAffinityGroupResult>("exoscale:index/getAntiAffinityGroup:getAntiAffinityGroup", args ?? new GetAntiAffinityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Fetch Exoscale [Anti-Affinity Groups](https://community.exoscale.com/documentation/compute/anti-affinity-groups/) data.
        /// 
        /// Corresponding resource: exoscale_anti_affinity_group.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Exoscale = Pulumi.Exoscale;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myAntiAffinityGroup = Exoscale.GetAntiAffinityGroup.Invoke(new()
        ///     {
        ///         Name = "my-anti-affinity-group",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["myAntiAffinityGroupId"] = myAntiAffinityGroup.Apply(getAntiAffinityGroupResult =&gt; getAntiAffinityGroupResult.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// Please refer to the examples
        /// directory for complete configuration examples.
        /// </summary>
        public static Output<GetAntiAffinityGroupResult> Invoke(GetAntiAffinityGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAntiAffinityGroupResult>("exoscale:index/getAntiAffinityGroup:getAntiAffinityGroup", args ?? new GetAntiAffinityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAntiAffinityGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The anti-affinity group ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The group name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetAntiAffinityGroupArgs()
        {
        }
        public static new GetAntiAffinityGroupArgs Empty => new GetAntiAffinityGroupArgs();
    }

    public sealed class GetAntiAffinityGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The anti-affinity group ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The group name to match (conflicts with `id`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetAntiAffinityGroupInvokeArgs()
        {
        }
        public static new GetAntiAffinityGroupInvokeArgs Empty => new GetAntiAffinityGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetAntiAffinityGroupResult
    {
        /// <summary>
        /// The anti-affinity group ID to match (conflicts with `name`).
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The list of attached exoscale*compute*instance (IDs).
        /// </summary>
        public readonly ImmutableArray<string> Instances;
        /// <summary>
        /// The group name to match (conflicts with `id`).
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetAntiAffinityGroupResult(
            string? id,

            ImmutableArray<string> instances,

            string? name)
        {
            Id = id;
            Instances = instances;
            Name = name;
        }
    }
}
