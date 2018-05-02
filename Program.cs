using System;
using Microsoft.Owin.Hosting;

namespace razweb
{
    class Program
    {
        static void Main(string[] args)
        {
            var url = "http://+:80";

            using (WebApp.Start<Startup>(url))
            {
                Console.WriteLine("Running on {0}", url);
                Console.WriteLine("Press enter to exit");
                Console.ReadLine();
            }
        }
    }
}
