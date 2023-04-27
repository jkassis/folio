jkassis/folio
=============
A collection of projects to demonstrate software engineering skills. Hopefully, after review, we can shortcut to your systems, people, and projects... and how we can help each other... and skip the coding tests.


About
-----
I've done development, ops, and scaling for 20+ years with extensive experience (at one time or another) in GoLang, Typescript, Python, Java, Javascript, and Perl.  I'm fresh in GoLang and Typescript at the moment.


Code Samples
------------

* [GAS: Game Animation System](https://github.com/jkassis/gas)  
  A fun glimpse of my game-development and front-end skills. I'll use this as a pivot-point for learning other languages as needed.

  [Adam Rogoyski](https://github.com/adamrogoyski), a former colleague at Google inspired this with his [Polyglot Tetris](https://github.com/adamrogoyski/tetris) experiment.

  Game-programming requires a special approach. If this doesn't feel familiar to you, that's ok. Don't confuse it with server code, build, release, or tools.

* [Pokemon API Client](https://github.com/jkassis/pokemoncli)  
  While I'm not focused on gaming at the moment, my history draws attention from game companies. I did this for Gearbox not long ago. It shows GoLang integration of [OpenTelemetry](https://opentelemetry.io/)... the new hotness in observability.

  It also shows a clean, lightweight CLI, consumption of a chaotic REST API, and a lightweight container build with Docker.

* [JerrieDr](https://github.com/jkassis/jerriedr)  
  JerrieDR de-risks backup and restore operations on my custom databases by automating those processes. It uses the k8s API to control StatefulSets and Ingresses of live cloud services and a terminal GUI framework for user control.

  It's an interesting study in how to backup microservices and maintain consistency of the data set. Short answer... it ain't easy.

* [GitAll](https://github.com/jkassis/gitall)  
  At any given time, I work from multiple dev laptops (desk / bed) to force myself to check in more frequently. But I have 137 git repositories. 63 are source repos that I created and the rest are forks. And syncing all those repos... what a PITA!
  
  At Google I got into the convenience of Blaze and the monorepo... so I wrote `gitall` to reduce the toil of maintaining all of *my* repos.
  
  * It uses a PureGo git client for git operations. I have mixed feelings, but it eliminates a dependency on git.
  * It has a custom, PureGo makefile [make.go](https://github.com/jkassis/gitall/blob/master/bin/make.go) that...
    * does cross-platform builds using docker, like [xgo](https://github.com/karalabe/xgo), but cleaner
    * packages the binary with [nFPM](https://github.com/goreleaser/nfpm), a subcomponent of [GoReleaser](https://github.com/goreleaser)
    * releases to github with the [GitHub CLI](https://cli.github.com/)

  Going deeper... [My XGO Fork](https://github.com/jkassis/xgo) includes the Dockerfile for the Ubuntu cross-build server image.

* [DragonBoat](https://github.com/lni/dragonboat/issues?q=jkassis)  
  A few issues I raised againt the [DragonBoat Implementation of RAFT](https://github.com/lni/dragonboat) protocol in GoLang.

* [NextTime](https://github.com/jkassis/nexttime)  
  A Typescript iterator for luxon Datetimes. Imagine why I did this.

* [EzGo Proxy](https://github.com/jkassis/ezgo/blob/master/proxy/client.go)  
  A light client that connects IoT devices to a custom IoT gateway.

* [My Dotfiles](https://github.com/jkassis/home)  
  For config and dev box installation. I keep them up-to-date and they change all the time.

 * [Code Screens] (https://github.com/jkassis/folio/code)  
   A few recent examples.


System Design
-------------
System Design interviews produce good conversations, but not always much to show. And I obviously can't share confidential System Design artifacts from previous employers.

Still, I have a few [System Design Artifacts](https://github.com/jkassis/folio/systemdesign) to show.




Projects
--------

### Infrastructure
I run my own k8s cluster in AWS using [OKD](https://www.okd.io/).

I decided to do this after a year and a half with [MachineZone](https://mz.com) where I helped migrate their multi-player game stack from metal to kubernetes and public / hybrid cloud.

I did testing on various k8s control planes, including [KOPS](https://github.com/kubernetes/kops), Amazon EKS, and Google GKE. The team and I stood the game server and client up on each of these and weighed pros and cons.

We eventually chose OKD for the next title launch and benchmarked up to 10k concurrent users. I contributed to the development of the logging system and migration of the Python load testing framework to kubernetes (and updating it with asyncio) so that we could load test without extreme Ingress and Compute fees.

I was not the chief strategist, but involved in all aspects of the migration.

### eCommerce
I created a full-stack eCommerce solution (soon to be released) using GoLang data services, typescript node.js server, typescript front-end with web technologies.

GoLang data services include a custom, composable microservices framework that allow arbitrary compositions of the following services (single-process, multi-process) with optional [Dragonboat RAFT](https://github.com/lni/dragonboat) redundancy on localhost or the cloud...
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

### IoT (Internet of Things)
I do some embedded / small device programming as well. Here are a few...

 * RaspberryPi OpenCV Security Camera
 * RaspberryPi Mechanum Rover
 * Arduino pulse Controller / Modulator
 * RaspberryPi Point of Sale Terminal with Touch Display
 * FreeSWAN VPN Firewall

### Games
I created two full-stack, multi-platform word games using a common engine and launched them on iOS, Android, and Facebook Web. That's 2 titles on three platforms using advanced web technology including WebGL, node.js, and MongoDB.


Editor / IDE
------------
While once a devout emacs user, I transitioned to Vim in 2016 with [Vim Adventures](https://vim-adventures.com/) and [NeoVim](https://neovim.io/).  Today, I use [VSCode](https://code.visualstudio.com/) and edit in [VSCodeVim](https://github.com/VSCodeVim/Vim).

VSCode is so remarkable that Google now provides it internally as CiderV.  I've used many editors over the years... and don't see coding without it.


Other Tools
-----------
### KVM
* [Dygma Raise](https://dygma.com/). The Raise has a learning curve, but is remarkable.
* 4k curved Phillips on a VESA mount.
* Magic mouse with touch scroll. This was such a game-changer. Ever had to scroll with a mouse, pointer, and scrollbar?

### Desk
I built my desk from walnut and a sit-stand workstation base. As an early adopter of sit-stand workstations, I had a Boston lift system. These days, $200 on Amazon gives you quality and clean.


Memories
--------
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
