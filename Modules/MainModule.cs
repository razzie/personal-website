using Nancy;
using System.Dynamic;

namespace razweb.Modules
{
    public class MainModule : NancyModule
    {
        public MainModule()
        {
            Get["/"] = args =>
            {
                //var requestEnvironment = (IDictionary<string, object>)Context.Items["OWIN_REQUEST_ENVIRONMENT"];
                //var user = (IPrincipal)requestEnvironment["server.User"];

                dynamic model = new ExpandoObject();
                model.FavoriteProjects = ProjectsDB.Projects;
                model.GithubReady = GithubDB.Ready;
                model.GithubProjects = GithubDB.Projects;
                model.GithubStars = GithubDB.Stars;

                return View["index", model];
            };

        }
    }
}
