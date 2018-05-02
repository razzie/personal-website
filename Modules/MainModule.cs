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

                dynamic viewbag = new ExpandoObject();
                viewbag.fav_projects = FavoriteProjects.Randomized;

                return View["index", viewbag];
            };

        }
    }
}
