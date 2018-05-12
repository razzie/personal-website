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
        private static GithubDB _github = new GithubDB("razzie", "b37e88bef5b39e88dab437fc49351aff1c29d853");
        private static ProjectsDB _projects = new ProjectsDB();

        static MainModule()
        {
            _projects.LoadFromAssembly("razweb.Modules.ProjectsDB.xml");
        }

        public MainModule()
        {
            Get["/"] = args =>
            {
                this.RequiresHttps();

                dynamic model = new ExpandoObject();
                model.FavoriteProjects = _projects.Projects.OrderBy(x => _rnd.Next());
                model.GithubReady = _github.Ready;
                if (_github.Ready)
                {
                    model.GithubProjects = _github.Projects.OrderBy(x => _rnd.Next()).Take(6);
                    model.GithubStars = _github.Stars.OrderBy(x => _rnd.Next()).Take(6);
                }

                return View["index", model];
            };

        }
    }
}
