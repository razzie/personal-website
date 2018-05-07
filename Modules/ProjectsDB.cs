using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

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

            public Info(string id, string name, string img, string descr)
            {
                ID = id;
                Name = name;
                ImageUrl = img;
                Description = descr;
            }
        }

        private static List<Info> _projects = new List<Info>();

        static ProjectsDB()
        {
            _projects.Add(new Info("floating_islands", "Floating Islands", "img/floating_islands.png",
                "Floating Islands is yet another example of my experiments with procedurally generated terrains. I developed an algorithm that takes island positions and island connections as input parameters and generates a 3D grid world using a Perlin noise internally. I also focused on the cartoonish design with edge outlining and also had an A* pathfinder algorithm set up for the little red cube."));

            _projects.Add(new Info("flyff", "Flyff fansite", "img/flyff.png",
                "Flyff.extra.hu was a hungarian fan site of Flyff (Fly For Fun) online game. During the 3-4 years of its existence I developed more and more new features to make it a fully functional web portal. It was very interesting to see it slowly getting attention. Members started generating content in the forms of forum entries, polls, screenshots and news for the home page. The site even had a few moderators and admins. In 2010 extra.hu decided to close its free hosting service and the site went offline. I could have moved it to a different(even paid) hosting service, but most members and I have already lost interest in the game. I still have a backup copy of the site and all of its content."));

            _projects.Add(new Info("gglib", "gglib", "img/gglib.png",
                "GGlib was a library I used to develop between 2014 and 2016. It served as a sandbox to help me to learn about the new features of the C++11 standard. I also found a great use of it while working on other side projects. Due to some early design choices(including the heavy use of virtual function calls) I finally decided to send it into retirement and started working on its successor: <a href=\"https://github.com/razzie/librazzie\" target=\"_blank\">librazzie</a>."));

            _projects.Add(new Info("ground_tool", "Ground Tool", "img/ground_tool.png",
                "The idea behind the demo was a game with floating islands. Though this game never made it to the reality, Ground tool was born as a result of experimenting with the rendering of these islands. The user can draw a polygon border to the island and when all points are connected, the program generates an island which can be inspected by zooming and moving the camera around."));

            _projects.Add(new Info("gtaonline_videos", "GTA Online videos", "img/gtaonline_videos.png",
                "<a href=\"https://vimeo.com/razzie/videos\" target=\"_blank\">vimeo.com/razzie/videos</a><br />¯\\_(ツ)_/¯"));

            _projects.Add(new Info("hexagon", "Hexagon", "img/hexagon.png",
                "Hexagon is very similar to the Minecraft clone, but I wanted to try rendering hexagon blocks (hexagonal prisms to be correct) instead of cubes. The underlying datastructure is still a 3D grid of Perlin noise. One of the main differences is that this demo does not use textures, but has a custom shader that draws grass or rock detail based on the normal vectors."));

            _projects.Add(new Info("librazzie", "librazzie", "img/librazzie.png",
                "<a href=\"https://github.com/razzie/librazzie\" target=\"_blank\">github.com/razzie/librazzie</a><br />¯\\_(ツ)_/¯"));

            _projects.Add(new Info("logic_circuit_simulator", "Logic Circuit Simulator", "img/logic_circuit_simulator.png",
                "Logic Circuit Simulator was a university project I had to finish in a semester. It offers a variety of logic gates and a lot of visual customization options. Some special elements are also available, like seven-segment display, matrix display, JK flip-flop, adjustable timer and module. A selected part of a logic circuit can be converted to a module, where the circuit's input elements (buttons, switches) become the input pins and output elements (LEDs) become the output pins of the module. Modules also support recursion."));

            _projects.Add(new Info("minecraft_clone", "Minecraft clone", "img/minecraft_clone.png",
                "Minecraft clone is a graphics project I worked on during the final year of the university. Contrary to its name it is not a real clone, it features no game elements. My primary motivation was the recreation of a Minecraft-like rendering in C++."));

            _projects.Add(new Info("potato_game", "Potato Game", "img/potato_game.png",
                "¯\\_(ツ)_/¯"));

            _projects.Add(new Info("prepi", "Prepi", "img/prepi.png",
                "Prepi is an indie game project I was working on together with a friend. It was intended to be a rage game with puzzle elements. We took it so seriously we both quitted our jobs to be able to work on it full-time. After several month of development we reached a point where we wanted to introduce our game to the public via an Indiegogo campaign. Unfortunately the campaign failed, mostly because the lack of media attention (neither of us were expert on the topic). Also the game was probably in too early phase to show it to the public. Due to the failure of the campaign and smaller disagreements regarding the game design we decided not to continue the development."));

            _projects.Add(new Info("process_manager", "Process Manager", "img/process_manager.png",
                "Process Manager is a replacement application for the built-in Windows Task Manager with a lot of additional features. It keeps running in the background after closing the main window and can be brought back from the tray icon or by pressing Ctrl+F12."));

            _projects.Add(new Info("razzgravitas", "RazzGravitas", "img/razzgravitas.png",
                "<a href=\"https://github.com/razzie/razzgravitas\" target=\"_blank\">github.com/razzie/razzgravitas</a><br />¯\\_(ツ)_/¯"));

            _projects.Add(new Info("razzie_messenger", "Razzie Messenger", "img/razzie_messenger.png",
                "Razzie Messenger is a messenger client and server I started working on during a one-week vacation at lake Balaton (Hungary). It uses a very simple plain text based protocol for network communication, look at the picture to see it in details. (Since then I prefer binary protocols due to the ineffeciency of string parsing and large packet sizes in case of text protocols.) The user can chose a nickname and pick colors for the name and the text messages. There is no registration and password authentication, however the server rejects the connection if a nickname is already taken in the current session. It is possible to send a file to an other user: in this case the server opens a random port which receives and forwards the file. Both users have to connect to this port, but it is done automatically. This solution was necessary due to the limitations of the text based protocol."));

            _projects.Add(new Info("server_client_app", "Server - client application", "img/server_client_application.png",
                "Server - client application is a small tool which helps creating and debugging text based network protocols. I developed it while working on Razzie Messenger and it had a great use. On the server tab the user can select one of the connected clients as a message target or just close the connection to them."));

            _projects.Add(new Info("windows_manager", "Windows Manager", "img/windows_manager.png",
                "Windows Manager is a reworked and extended edition of Process Manager. Besides the improved visual appearance it introduces new window tweaking options which you can see in the picture. One can for example change a sticky (always on top) window to act as a normal one."));
        }
        
        public static IEnumerable<Info> Projects
        {
            get { return _projects; }
        }
    }
}
