using Nancy.Hosting.Self;
using razweb.Modules;
using System;

namespace razweb
{
    class Program
    {
        static void Main(string[] args)
        {
            var uri = new Uri("http://localhost:8080");
            var bootstrapper = new CustomBootstrapper();
            var hostConfiguration = new HostConfiguration
            {
                UrlReservations = new UrlReservations() { CreateAutomatically = true }
            };

            using (var host = new NancyHost(bootstrapper, hostConfiguration, uri))
            {
                host.Start();
                GithubDB.Update();

                Console.WriteLine("Press enter to exit");
                Console.ReadLine();
            }
        }
    }
}
