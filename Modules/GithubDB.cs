using Octokit;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Timers;

namespace razweb.Modules
{
    public static class GithubDB
    {
        public class Repo
        {
            public class Commit
            {
                public string CommitID { get; private set; }
                public string CommitMessage { get; private set; }
                public string User { get; private set; }
                public DateTimeOffset Date { get; private set; }

                public Commit(string id, string message, string user, DateTimeOffset date)
                {
                    CommitID = id.Substring(0, 7);
                    CommitMessage = message.Length > 100 ? message.Substring(0, 100) + "..." : message;
                    User = user;
                    Date = date;
                }
            }

            public string Name { get; private set; }
            public string Description { get; private set; }
            public string Owner { get; private set; }
            public string Url { get; private set; }
            public IEnumerable<Commit> TopCommits { get; private set; }

            public static Repo Create(Repository repo)
            {
                var top_commits = new List<Commit>();

                var api_options = new ApiOptions();
                api_options.PageCount = 1;
                api_options.PageSize = 5;

                var commits = _github.Repository.Commit.GetAll(repo.Id, api_options).Result;
                foreach (var commit in commits)
                {
                    top_commits.Add(new Commit(commit.Sha, commit.Commit.Message, commit.Commit.Author.Name, commit.Commit.Author.Date));

                    if (top_commits.Count >= 5)
                        break;
                }

                var result = new Repo();
                result.Name = repo.Name;
                result.Description = repo.Description;
                result.Owner = repo.Owner.Login;
                result.Url = repo.HtmlUrl;
                result.TopCommits = top_commits;

                return result;
            }
        }

        private static GitHubClient _github = new GitHubClient(new ProductHeaderValue("razzie-updater"));
        private static List<Repo> _projects = new List<Repo>();
        private static List<Repo> _stars = new List<Repo>();
        private static Timer _updater = new Timer();

        static GithubDB()
        {
            _github.Credentials = new Credentials("b37e88bef5b39e88dab437fc49351aff1c29d853");

            _updater.Interval = 1800000; // 30 minutes
            _updater.Elapsed += (sender, args) => Update();
            _updater.Start();
        }

        public static bool Ready
        {
            get; private set;
        }

        public static IEnumerable<Repo> Projects
        {
            get { return _projects; }
        }

        public static IEnumerable<Repo> Stars
        {
            get { return _stars; }
        }

        public static void Update()
        {
            var projects = new List<Repo>();
            var stars = new List<Repo>();

            try
            {
                var user_repos = _github.Repository.GetAllForUser("razzie").Result;
                foreach (var repo in user_repos)
                {
                    if (!repo.Fork)
                        projects.Add(Repo.Create(repo));
                }

                var api_connection = new ApiConnection(_github.Connection);
                var starred_client = new StarredClient(api_connection);
                var starred_repos = starred_client.GetAllForUser("razzie").Result;
                foreach (var repo in starred_repos)
                {
                    stars.Add(Repo.Create(repo));
                }
            }
            catch (Exception e)
            {
                for (; e != null; e = e.InnerException)
                    Console.WriteLine(e.Message);

                return;
            }

            _projects = projects;
            _stars = stars;
            Ready = true;
        }
    }
}
