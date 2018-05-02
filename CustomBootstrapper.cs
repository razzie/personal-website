using Nancy;
using Nancy.Bootstrapper;
using Nancy.Conventions;
using Nancy.Diagnostics;
using Nancy.Session;
using Nancy.TinyIoc;
using Nancy.ViewEngines;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;

//[assembly: IncludeInNancyAssemblyScanning]
namespace razweb
{
    public class CustomBootstrapper : DefaultNancyBootstrapper
    {
        private byte[] favicon;

        protected override byte[] FavIcon
        {
            get { return this.favicon ?? (this.favicon = LoadFavIcon()); }
        }

        private byte[] LoadFavIcon()
        {
            using (var resourceStream = GetType().Assembly.GetManifestResourceStream("razweb.Content.Images.favicon.png"))
            {
                var memoryStream = new MemoryStream();
                resourceStream.CopyTo(memoryStream);
                return memoryStream.GetBuffer();
            }
        }

        protected override NancyInternalConfiguration InternalConfiguration
        {
            get
            {
                return NancyInternalConfiguration.WithOverrides(OnConfigurationBuilder);
            }
        }

        void OnConfigurationBuilder(NancyInternalConfiguration x)
        {
            x.ViewLocationProvider = typeof(ResourceViewLocationProvider);
        }

        protected override void ConfigureApplicationContainer(TinyIoCContainer container)
        {
            Assembly assembly = Assembly.GetExecutingAssembly();
            string[] assemblyNames = assembly.GetManifestResourceNames();

            base.ConfigureApplicationContainer(container);
            ResourceViewLocationProvider.RootNamespaces.Add(GetType().Assembly, "razweb.Views");
        }

        protected override void ConfigureConventions(NancyConventions nancyConventions)
        {
            Assembly assembly = Assembly.GetExecutingAssembly();

            nancyConventions.StaticContentsConventions.Add(StaticResourceConventionBuilder.AddDirectory("/js", assembly, "razweb.Content.Javascript"));
            nancyConventions.StaticContentsConventions.Add(StaticResourceConventionBuilder.AddDirectory("/css", assembly, "razweb.Content.Stylesheet"));
            nancyConventions.StaticContentsConventions.Add(StaticResourceConventionBuilder.AddDirectory("/img", assembly, "razweb.Content.Images"));
        }

        protected override void ApplicationStartup(TinyIoCContainer container, IPipelines pipelines)
        {
            base.ApplicationStartup(container, pipelines);
            CookieBasedSessions.Enable(pipelines);
        }

        protected override IEnumerable<ModuleRegistration> Modules
        {
            get
            {
                return AppDomainAssemblyTypeScanner
                        .TypesOf<INancyModule>(ScanMode.All)
                        .NotOfType<DiagnosticModule>()
                        .Select(t => new ModuleRegistration(t))
                        .ToArray();
            }
        }
    }
}
