using System;
using Microsoft.Owin.Hosting;
using razweb.Modules;

namespace razweb
{
    class Program
    {
        static void Main(string[] args)
        {
            var options = new StartOptions();
            options.Urls.Add("http://+:80");
            options.Urls.Add("https://+:443");

            using (WebApp.Start<Startup>(options))
            {
                GithubDB.Update();

                Console.WriteLine("Press enter to exit");
                Console.ReadLine();
            }
        }
    }
}
