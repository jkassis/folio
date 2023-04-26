package main

import (
	"fmt"
	"os"
	"sync"
)

// Write a simple browser history functionality for a single tab browser.
// The default webpage for every tab is "roblox.com".
// Allow for visiting any webpage, and browsing back and forward with variable depth.
// Your code should have the following public methods, but feel free to add any other methods
// or functions that can help you debug or check your solution.

// void visit(string destinationUrl)
//  - goes to a given URL
//  - starts new forward history

// string back(int steps)
//  - goes back in browsing history by the number of steps
//  - returns the URL at that point in history

// string forward(int steps)
//  - goes forward in history by the number of steps
//  - returns the URL at that point in history

// Support for multiple tabs: Now we want to support 1...n at the same time, with history tracked separately in all. You may modify existing methods to support this expanded behavior.

// Functionality to be added:
// - open a new tab
// - close an existing tab

// other methods as you see fit

// Allow restoring closed tabs. Should discuss both restoring a closed tab on a running browser AND restoring a closed tab from a previous browser session.

type Browser struct {
	tabs []*Tab
	mut  sync.Mutex
}

func (b *Browser) TabOpen() {
	defer b.mut.Unlock()
	b.mut.Lock()
	newTab := MakeTab()
	b.tabs = append(b.tabs, newTab)
}

func (b *Browser) TabClose(i int) {
	defer b.mut.Unlock()
	b.mut.Lock()

	if i < 0 {
		fmt.Fprintf(os.Stderr, "cannot close tab %d. index out of range.", i)
	} else if i > (len(b.tabs) - 1) {
		fmt.Fprintf(os.Stderr, "cannot close tab %d. index out of range.", i)
	}
	// tab := b.tabs[i]
	b.tabs = append(b.tabs[0:i], b.tabs[(i+1):]...)
}

func MakeBrowser() *Browser {
	initialTab := MakeTab()
	return &Browser{
		tabs: []*Tab{initialTab},
	}
}

type Tab struct {
	historyPage int
	history     []string
	mut         sync.Mutex
}

func MakeTab() *Tab {
	return &Tab{
		historyPage: 0,
		history:     []string{"roblox.com"},
	}
}

func (b *Tab) Visit(destinationUrl string) {
	defer b.mut.Unlock()
	b.mut.Lock()

	b.history = append(b.history[0:b.historyPage+1], destinationUrl)
	b.historyPage += 1
}

func (b *Tab) Back(steps int) string {
	defer b.mut.Unlock()
	b.mut.Lock()

	if steps > b.historyPage {
		b.historyPage = 0
	} else {
		b.historyPage -= steps
	}
	return b.history[b.historyPage]
}

func (b *Tab) Forward(steps int) string {
	defer b.mut.Unlock()
	b.mut.Lock()

	if b.historyPage+steps > (len(b.history) - 1) {
		b.historyPage = (len(b.history) - 1)
	} else {
		b.historyPage += steps
	}
	return b.history[b.historyPage]
}

func (b *Tab) Print() {
	defer b.mut.Unlock()
	b.mut.Lock()

	fmt.Printf("Browser:\n")
	fmt.Printf("  currentPage: %d\n", b.historyPage)
	fmt.Printf("  history: %v\n", b.history)
}

func main() {
	b := MakeBrowser()

	t := b.tabs[0]
	fmt.Println("Good Morning!")
	t.Print()
	t.Visit("www.meta.com")
	t.Print()
	t.Visit("www.airbnb.com")
	t.Print()
	t.Visit("www.gmail.com")
	t.Print()
	t.Visit("calendar.gmail.com")
	t.Print()
	t.Back(10)
	t.Print()
	t.Forward(2)
	t.Print()
	t.Forward(3)
	t.Print()
	t.Back(2)
	t.Print()
	t.Visit("www.roblox.com")
	t.Print()
}
