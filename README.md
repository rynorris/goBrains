#goBrains

Real-time neural network controlled creatures, implemented in Go.

##Status
This project is complete, barring feature additions.
The creatures are capable of learning to search for, and find food.

##Aims
We aim to create a general framework on which basic evolutionary creatures can be developed.

Basic creatures are generated with no knowledege or understanding and random brains.
No feedback is given, other than that creatures which live longer will have more chances to reproduce.
Over time, they evolve to be extremely efficient at finding and consuming food.

##Usage
Simply run with `goBrains`.

To use the web view (recommended), run with `goBrains --headless` and then connect at `localhost:9999/tank`.
Then use Q and W to adjust the speed of the simulation.  Or Z to toggle the frame limit altogether.  You can scroll by click&drag, and zoom with the mouse wheel.

##Dependencies
goSDL is used to visulize creatures in their native habitat (your computer).
golang.org/x/net/websocket is used to stream the data for viewing in a browser.

##Additional disclaimers
In the possibility that the basic creatures become intelligent enough to escape the confines of the simulation the developers would like to congratulate you on your discovery and will retreat to a safe distance.

##Related projects
Please see DiscoViking/pyBrains for the original implementation of neural network controlled creatures, implemented in Python.
