using Nancy.Hosting.Self;
using razweb.Modules;
using System;

namespace razweb
{
    class Program
    {
        static void Main(string[] args)
        {
            using (var host = new NancyHost(new Uri("http://localhost:8080")))
            {
                host.Start();
                GithubDB.Update();

                Console.WriteLine("Press enter to exit");
                Console.ReadLine();
            }
        }
    }
}
