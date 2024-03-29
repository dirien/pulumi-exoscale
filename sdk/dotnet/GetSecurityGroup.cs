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
    public static class GetSecurityGroup
    {
        /// <summary>
        /// Fetch Exoscale [Security Groups](https://community.exoscale.com/documentation/compute/security-groups/) data.
        /// 
        /// Corresponding resource: exoscale_security_group.
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
        ///     var mySecurityGroup = Exoscale.GetSecurityGroup.Invoke(new()
        ///     {
        ///         Name = "my-security-group",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mySecurityGroupId"] = mySecurityGroup.Apply(getSecurityGroupResult =&gt; getSecurityGroupResult.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// Please refer to the examples
        /// directory for complete configuration examples.
        /// </summary>
        public static Task<GetSecurityGroupResult> InvokeAsync(GetSecurityGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupResult>("exoscale:index/getSecurityGroup:getSecurityGroup", args ?? new GetSecurityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Fetch Exoscale [Security Groups](https://community.exoscale.com/documentation/compute/security-groups/) data.
        /// 
        /// Corresponding resource: exoscale_security_group.
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
        ///     var mySecurityGroup = Exoscale.GetSecurityGroup.Invoke(new()
        ///     {
        ///         Name = "my-security-group",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mySecurityGroupId"] = mySecurityGroup.Apply(getSecurityGroupResult =&gt; getSecurityGroupResult.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// Please refer to the examples
        /// directory for complete configuration examples.
        /// </summary>
        public static Output<GetSecurityGroupResult> Invoke(GetSecurityGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityGroupResult>("exoscale:index/getSecurityGroup:getSecurityGroup", args ?? new GetSecurityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The security group ID to match (conflicts with `name`)
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name to match (conflicts with `id`)
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetSecurityGroupArgs()
        {
        }
        public static new GetSecurityGroupArgs Empty => new GetSecurityGroupArgs();
    }

    public sealed class GetSecurityGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The security group ID to match (conflicts with `name`)
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name to match (conflicts with `id`)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetSecurityGroupInvokeArgs()
        {
        }
        public static new GetSecurityGroupInvokeArgs Empty => new GetSecurityGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityGroupResult
    {
        /// <summary>
        /// The list of external network sources, in [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notatio) notation.
        /// </summary>
        public readonly ImmutableArray<string> ExternalSources;
        /// <summary>
        /// The security group ID to match (conflicts with `name`)
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name to match (conflicts with `id`)
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetSecurityGroupResult(
            ImmutableArray<string> externalSources,

            string? id,

            string? name)
        {
            ExternalSources = externalSources;
            Id = id;
            Name = name;
        }
    }
}
