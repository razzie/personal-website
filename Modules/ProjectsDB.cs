using System;
using System.Collections.Generic;
using System.IO;
using System.Reflection;
using System.Xml;

namespace razweb.Modules
{
    public static class ProjectsDB
    {
        public class Info
        {
            public string ID { get; private set; }
            public string Name { get; private set; }
            public string ImageUrl { get; private set; }
            public string Description { get; private set; }

            public static Info Read(XmlNode node)
            {
                try
                {
                    var info = new Info();
                    info.ID = node.SelectSingleNode("id").InnerXml;
                    info.Name = node.SelectSingleNode("name").InnerXml;
                    info.ImageUrl = node.SelectSingleNode("img").InnerXml;
                    info.Description = node.SelectSingleNode("description").InnerXml;
                    return info;
                }
                catch (Exception e)
                {
                    Console.WriteLine("Failed to read project from xml: " + e.Message);
                    return null;
                }
            }
        }

        private static List<Info> _projects = new List<Info>();

        static ProjectsDB()
        {
            var assembly = Assembly.GetExecutingAssembly();
            using (Stream stream = assembly.GetManifestResourceStream("razweb.Modules.ProjectsDB.xml"))
            {
                XmlDocument doc = new XmlDocument();
                doc.Load(stream);

                foreach (XmlNode node in doc.SelectNodes("/projects/project"))
                {
                    var project = Info.Read(node);
                    if (project != null)
                        _projects.Add(project);
                }
            }
        }
        
        public static IEnumerable<Info> Projects
        {
            get { return _projects; }
        }
    }
}
