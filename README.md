jkassis/folio
=============
A guided-tour of some interesting tidbits for you to get to know my software engineering skills.

Hopefully, after reviewing this work, we can fast-forward through coding tests to talk about your projects, system design, and how we can help each other.

About
-----
I've done software development for 20+ years with extensive experience (at one time or another) in GoLang, Typescript, Python, Java, Javascript, and Perl.

I'm fresh in GoLang and Typescript at the moment.


Code
----
Here are some interesting links:

* [GAS: Game Animation System](https://github.com/jkassis/gas)
  I just put this together to give you a fun glimpse of my game-development and front-end experience.

  I also intend to use this as a pivot-point for learning other languages. While I've done C++ before, translating this will freshen those chops. I'm also planning for a Rust translation and more.

  [Adam Rogoyski](https://github.com/adamrogoyski), a former colleague at Google inspired this with his [Polyglot Tetris](https://github.com/adamrogoyski/tetris) experiment, which I admire so much.

  Game-programming requires a special approach. If this doesn't feel familiar to you, that's ok. Don't confuse it with server code, build and release, or tools.

* [Pokemon API Client](https://github.com/jkassis/pokemoncli)
  While I'm not focused on gaming at the moment, my history tends to draw a lot of attention from game companies. I did this for Gearbox not long ago. It shows off [OpenTelemetry](https://opentelemetry.io/) integration in GoLang. OpenTelemetry is all the rage in observability right now.

  It also shows off a clean, lightweight CLI, consumption of a chaotic REST API, and a lightweight container build with Docker.

  How bout dem buzzwords?

* [JerrieDr](https://github.com/jkassis/jerriedr)
  You won't know what to do with this GoLang ops tool, but you can take a look at the code.

  Notably, it de-risks backup and restore operations on my custom databases by automating those processes.

  It uses the k8s API to control StatefulSets and Ingresses of live cloud services and a terminal GUI framework for user control.

  It's an interesting study in how to backup microservices and maintain consistency of the data set. Short answer... it ain't easy.

* [GitAll](https://github.com/jkassis/gitall)
  I have 137 git repositories. 63 are source repos that I created and the rest are forks. That's a lot of repos.

  And at any given time, I have multiple dev laptops (desk / bed) going, which I find healthy because it forces me to check in more frequently.

  But checking on all those repos to see if they are synced... what a PITA!

  While I Google I got into the convenience of Blaze and the monorepo... so I wrote `gitall` to reduce the toil of maintaining all of *my* repos.

  It's a work in progress, but I imagine that one day it could serve as the mono-repo layer on top of git.

  A few interesting tidbits here...
  * It uses a PureGo git client (i have mixed feelings) for git operations (does not require git to be installed)
  * It has a custom, PureGo makefile [make.go](https://github.com/jkassis/gitall/blob/master/bin/make.go) that...
    * does cross-platform builds using docker, like [xgo](https://github.com/karalabe/xgo), but cleaner
    * packages the binary with [nFPM](https://github.com/goreleaser/nfpm), a subcomponent of [GoReleaser](https://github.com/goreleaser)
    * releases to github with the [GitHub CLI](https://cli.github.com/)

  If you're really interested, [My XGO Fork](https://github.com/jkassis/xgo) includes the Dockerfile for the Ubuntu cross-build server image.

* [DragonBoat Issues](https://github.com/lni/dragonboat/issues?q=jkassis)
  Check out a few issues I raised againt the [DragonBoat Implementation of RAFT](https://github.com/lni/dragonboat) protocol in GoLang

* [NextTime](https://github.com/jkassis/nexttime)
  Here's a Typescript iterator for luxon Datetimes. Imagine why I did this.

* [EzGo Proxy](https://github.com/jkassis/ezgo/blob/master/proxy/client.go)
  For those interested in networking... I wrote this client to connect IoT devices to a custom gateway.

* [My Dotfiles](https://github.com/jkassis/home)
  I keep my dotfiles and dev box installation scripts checked into git.  I keep them up-to-date and they change all the time.  This, alone, should convince you that I'm a hardcore software developer.

 * [Code Screens] (https://github.com/jkassis/folio/code)
   I've taken many coding tests in my career. I wish I had saved them all. Recently, I decided to start collecting. This repo contains a few choice examples...



System Design
-------------
System Design interviews tend to produce good conversations, but not good assets for show.

And while I've been building and working at companies that do distributed systems since the beginning of the Internet, I can't share confidential System Design artifacts from those companies.

Still, I have a [few artifacts](https://github.com/jkassis/folio/systemdesign) to show...




Projects
--------
### eCommerce
I created a full-stack eCommerce solution (soon to be released) using GoLang data services, typescript node.js server, typescript front-end with web technologies.

GoLang data services include a custom, composable microservices framework that allow the following services to run individually or in a single process with optional [Dragonboat RAFT](https://github.com/lni/dragonboat) redundancy on localhost or the cloud...
* dockie
  A structured DB topping an open-source KV DB
* pubsub
  A PubSub service
* tickie
  A multi-threaded timer / callback service
* ledgie
  A basic, lightweight accounting ledger for high transaction volumes
* permie
  An RBAC (role-based authentication) server/db



### Infrastructure
I run my own k8s cluster in AWS using [OKD](https://www.okd.io/).

I decided to do this after a year and a half with [MachineZone](https://mz.com) where I helped migrate their multi-player game stack from metal to kubernetes and public / hybrid cloud.

I was involved in testing various k8s control planes including [KOPS](https://github.com/kubernetes/kops) Amazon EKS and Google GKE.

We stood the game server and client up on each of these and weighed pros and cons.

We eventually chose OKD for the launch of a new title and benchmarked up to 10k concurrent users. I contributed to the development of the logging system and migration of the Python load testing framework to kubernetes (and updating it with asyncio) so that we could load test without extreme Ingress fees.

I was not the chief strategist, but involved in all aspects of the migration.

### IoT (Internet of Things)
I do some embedded / small device programming as well. Here are a few...

 * RaspberryPi OpenCV Security Camera
 * RaspberryPi Controlled Mechanum Autonomous Vehicle
 * Arduino Pulse Controller / Modulator
 * RaspberryPi Point of Sale device with Touch Display
 * FreeSWAN VPN Firewall


### Word Games
I created two full-stack, multi-platform word games using a common engine and launched them on iOS, Android, and Facebook Web. That's 2 titles on three platforms using advanced web technology including WebGL, node.js server, and MongoDB.


Editor / IDE
------------
While once a devout emacs user, I transitioned to Vim in 2016 with [Vim Adventures](https://vim-adventures.com/) and [NeoVim](https://neovim.io/).  Today, I use [VSCode](https://code.visualstudio.com/) and edit in VSCode Vim(https://github.com/VSCodeVim/Vim) and don't see a future in which I don't.


Other Tools
-----------
### Keyboard
I use a [Dygma Raise](https://dygma.com/).

I got into mechanical keyboards again in 2016 (around the time I got into Vim) and have used many split keyboards (Goldtouch), etc.

Without getting into details, the Raise has a learning curve, but it is also remarkable gear.

### Mouse
The Magic mouse with touch scroll is a game changer. Mid/old-timers like me agree on this one. Ever had to scroll your code with a mouse pointer and a scrollbar? I have. I wonder why it took so long to get here.

### Monitor
4k curved.

### Desk
I custom built my desk from walnut and a sit-stand workstation base. I was one of the first into sit-sand workstations and had a Boston lift system. These days what you get on Amazon for $200 gives you quality and clean.


Memories
--------
VSCode is a remarkable product, so remarkable that Google now provides it internally as CiderV.  I've used many editors over the years...
 * [Merlin](https://brutaldeluxe.fr/products/crossdevtools/merlin/)
   for 65C816 assembly on the Apple IIgs.
 * [Eclipse](https://eclipse.org)
 * [IntelliJ IDEA](https://www.jetbrains.com/idea/)
 * [Apache NetBeans](https://netbeans.apache.org/)
 * [XCode](https://developer.apple.com/xcode/)
 * [Android Studio](https://developer.android.com/studio)

Boxes
* Apple IIe
* Apple IIgs
* Apple IIsi
* Apple Mac Quadra 950
* iMac (a few)
* 2008 MacBook
* Many more MacBooks
* Some Dells
* Homebuilts
