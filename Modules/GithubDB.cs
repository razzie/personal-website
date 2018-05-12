using Octokit;
using System;
using System.Collections.Generic;
using System.Threading;

namespace razweb.Modules
{
    public class GithubDB
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

            public static Repo Create(GitHubClient github, Repository repo)
            {
                var top_commits = new List<Commit>();

                var api_options = new ApiOptions();
                api_options.PageCount = 1;
                api_options.PageSize = 5;

                var commits = github.Repository.Commit.GetAll(repo.Id, api_options).Result;
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

        private List<Repo> _projects;
        private List<Repo> _stars;
        private string _user;
        private GitHubClient _github;
        private Timer _updater;
        private object _lock = new object();

        public GithubDB(string user, string credentials = null)
        {
            _user = user;

            _github = new GitHubClient(new ProductHeaderValue(user + "-updater"));
            if (credentials != null)
                _github.Credentials = new Credentials(credentials);

            _updater = new Timer(_ => Update(), null, TimeSpan.Zero, TimeSpan.FromMinutes(30));
        }

        public bool Ready
        {
            get
            {
                return (_projects != null && _stars != null);
            }
        }

        public IEnumerable<Repo> Projects
        {
            get { return _projects; }
        }

        public IEnumerable<Repo> Stars
        {
            get { return _stars; }
        }

        private void Update()
        {
            var projects = new List<Repo>();
            var stars = new List<Repo>();

            try
            {
                var user_repos = _github.Repository.GetAllForUser(_user).Result;
                foreach (var repo in user_repos)
                {
                    if (!repo.Fork)
                        projects.Add(Repo.Create(_github, repo));
                }

                var api_connection = new ApiConnection(_github.Connection);
                var starred_client = new StarredClient(api_connection);
                var starred_repos = starred_client.GetAllForUser(_user).Result;
                foreach (var repo in starred_repos)
                {
                    stars.Add(Repo.Create(_github, repo));
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
        }
    }
}
