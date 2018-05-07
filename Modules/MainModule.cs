using Nancy;
using Nancy.Security;
using System;
using System.Dynamic;
using System.Linq;

namespace razweb.Modules
{
    public class MainModule : NancyModule
    {
        private Random _rnd = new Random(DateTime.Now.Millisecond);

        public MainModule()
        {
            Get["/"] = args =>
            {
                this.RequiresHttps();

                //var requestEnvironment = (IDictionary<string, object>)Context.Items["OWIN_REQUEST_ENVIRONMENT"];
                //var user = (IPrincipal)requestEnvironment["server.User"];

                dynamic model = new ExpandoObject();
                model.FavoriteProjects = ProjectsDB.Projects.OrderBy(x => _rnd.Next());
                model.GithubReady = GithubDB.Ready;
                model.GithubProjects = GithubDB.Projects.OrderBy(x => _rnd.Next()).Take(6);
                model.GithubStars = GithubDB.Stars.OrderBy(x => _rnd.Next()).Take(6);

                return View["index", model];
            };

        }
    }
}
