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
    public static class GetTemplate
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Exoscale = Pulumi.Exoscale;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myTemplate = Exoscale.GetTemplate.Invoke(new()
        ///     {
        ///         Zone = "ch-gva-2",
        ///         Name = "Linux Ubuntu 22.04 LTS 64-bit",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["myTemplateId"] = myTemplate.Apply(getTemplateResult =&gt; getTemplateResult.Id),
        ///     };
        /// });
        /// ```
        /// 
        /// Please refer to the examples
        /// directory for complete configuration examples.
        /// </summary>
        public static Task<GetTemplateResult> InvokeAsync(GetTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTemplateResult>("exoscale:index/getTemplate:getTemplate", args ?? new GetTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Exoscale = Pulumi.Exoscale;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myTemplate = Exoscale.GetTemplate.Invoke(new()
        ///     {
        ///         Zone = "ch-gva-2",
        ///         Name = "Linux Ubuntu 22.04 LTS 64-bit",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["myTemplateId"] = myTemplate.Apply(getTemplateResult =&gt; getTemplateResult.Id),
        ///     };
        /// });
        /// ```
        /// 
        /// Please refer to the examples
        /// directory for complete configuration examples.
        /// </summary>
        public static Output<GetTemplateResult> Invoke(GetTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTemplateResult>("exoscale:index/getTemplate:getTemplate", args ?? new GetTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The compute instance template ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The template name to match (conflicts with `id`) (when multiple templates have the same name, the newest one will be returned).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A template category filter (default: `public`); among: - `public` - official Exoscale templates - `private` - custom templates private to my organization
        /// </summary>
        [Input("visibility")]
        public string? Visibility { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetTemplateArgs()
        {
        }
        public static new GetTemplateArgs Empty => new GetTemplateArgs();
    }

    public sealed class GetTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The compute instance template ID to match (conflicts with `name`).
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The template name to match (conflicts with `id`) (when multiple templates have the same name, the newest one will be returned).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A template category filter (default: `public`); among: - `public` - official Exoscale templates - `private` - custom templates private to my organization
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetTemplateInvokeArgs()
        {
        }
        public static new GetTemplateInvokeArgs Empty => new GetTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetTemplateResult
    {
        /// <summary>
        /// Username to use to log into a compute instance based on this template
        /// </summary>
        public readonly string DefaultUser;
        /// <summary>
        /// The compute instance template ID to match (conflicts with `name`).
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The template name to match (conflicts with `id`) (when multiple templates have the same name, the newest one will be returned).
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A template category filter (default: `public`); among: - `public` - official Exoscale templates - `private` - custom templates private to my organization
        /// </summary>
        public readonly string? Visibility;
        /// <summary>
        /// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetTemplateResult(
            string defaultUser,

            string? id,

            string? name,

            string? visibility,

            string zone)
        {
            DefaultUser = defaultUser;
            Id = id;
            Name = name;
            Visibility = visibility;
            Zone = zone;
        }
    }
}
