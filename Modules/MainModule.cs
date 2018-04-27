using Nancy;

namespace Modules
{
    public class MainModule : NancyModule
    {
        public MainModule()
        {
            Get["/"] = args =>
            {
                //var requestEnvironment = (IDictionary<string, object>)Context.Items["OWIN_REQUEST_ENVIRONMENT"];
                //var user = (IPrincipal)requestEnvironment["server.User"];
                return View["index"];
            };

        }
    }
}
