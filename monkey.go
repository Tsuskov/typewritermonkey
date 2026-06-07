package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//go:embed ape.md
var frame1 string

//go:embed apetype.md
var frame2 string

// ANSI escape codes
const (
	clearScreen    = "\033[2J\033[H"
	hideCursor     = "\033[?25l"
	showCursor     = "\033[?25h"
	saveCursor     = "\0337"
	restoreCursor  = "\0338"
	enterAltScreen = "\033[?1049h"
	exitAltScreen  = "\033[?1049l"
	bold           = "\033[1m"
	dim            = "\033[2m"
	reset          = "\033[0m"
	green          = "\033[32m"
	yellow         = "\033[33m"
	cyan           = "\033[36m"
	white          = "\033[97m"
	gray           = "\033[90m"
	red            = "\033[31m"
	magenta        = "\033[35m"
)

var ideas = []string{
	// CLI Tools
	"a CLI tool that converts markdown to beautiful terminal output",
	"a terminal pomodoro timer with ASCII progress bars",
	"a CLI password manager with AES encryption",
	"a command-line RSS reader with vim keybindings",
	"a terminal-based Kanban board",
	"a CLI tool that generates .gitignore files from templates",
	"a terminal file explorer with preview pane",
	"a CLI tool that summarizes git history in plain language",
	"a terminal-based budget tracker",
	"a CLI tool that watches a folder and auto-commits changes",
	"a terminal multiplexer plugin that shows system stats",
	"a CLI tool that converts CSV to beautiful ASCII tables",
	"a command-line flashcard app for spaced repetition learning",
	"a terminal-based habit tracker with streaks",
	"a CLI tool that generates a daily todo from your calendar",
	"a terminal diff viewer with syntax highlighting",
	"a CLI port scanner with a live updating table",
	"a command-line tool that mangles code for obfuscation",
	"a terminal weather app with ASCII art forecasts",
	"a CLI tool that formats SQL queries beautifully",
	"a terminal-based journal with encryption",
	"a CLI tool that auto-generates changelogs from git commits",
	"a command-line image resizer with batch processing",
	"a terminal-based music visualizer using FFT",
	"a CLI tool that checks for broken links in markdown files",
	"a terminal IRC client with notification support",
	"a CLI book tracker with reading progress and notes",
	"a command-line tool that converts JSON to Go structs",
	"a terminal-based time zone converter for remote teams",
	"a CLI tool that lints commit messages",

	// Web Apps
	"a real-time collaborative whiteboard app",
	"a web app that turns any URL into a distraction-free reading view",
	"a micro-blogging platform with no JavaScript on the frontend",
	"a web app that generates color palettes from uploaded images",
	"a recipe manager that scales ingredients automatically",
	"a web-based Gantt chart builder that exports to SVG",
	"a link-in-bio page builder with live preview",
	"a web app that converts voice notes to structured text",
	"a minimalist habit tracker as a single HTML file",
	"a web app that generates wireframes from text descriptions",
	"a flashcard SRS app that works offline as a PWA",
	"a web-based SQL playground with shareable queries",
	"a markdown blog engine with zero dependencies",
	"a web app that monitors website uptime and sends alerts",
	"a browser-based pixel art editor that exports to SVG",
	"a personal finance dashboard that imports bank CSV exports",
	"a web app that creates printable weekly planners",
	"a URL shortener with click analytics",
	"a web-based Pomodoro timer with task queue",
	"a collaborative decision-making app using ranked voting",
	"a web app that tracks your subscriptions and monthly spend",
	"a minimal pastebin clone with syntax highlighting",
	"a web app that generates fake but realistic test data",
	"a browser extension that blocks distracting sites on a schedule",
	"a web app for tracking reading lists with progress bars",
	"a one-page invoice generator that exports to PDF",
	"a web app that converts spreadsheets to interactive charts",
	"a meeting cost calculator that runs in the browser",
	"a web app that generates privacy policies from a questionnaire",
	"a browser-based MIDI sequencer with a step grid",

	// APIs & Backend
	"a REST API that serves random programming challenges",
	"a webhook relay service that logs and replays requests",
	"a GraphQL API for personal finance tracking",
	"an API that converts HTML emails to plain text",
	"a rate-limiting middleware library for Go",
	"an API that generates placeholder images with custom text",
	"a service that sends daily digest emails from RSS feeds",
	"a backend that manages feature flags for your apps",
	"an API gateway with built-in request caching",
	"a service that archives tweets to a personal database",
	"a backend for a multi-tenant note-taking app",
	"an API that parses and normalizes addresses",
	"a service that compresses images on upload automatically",
	"a backend that handles file uploads to multiple cloud providers",
	"an API that checks passwords against breach databases",
	"a service that generates PDF invoices from JSON",
	"a backend for a simple event ticketing system",
	"an API that extracts metadata from URLs (og:title, etc.)",
	"a service that translates error messages to plain English",
	"a backend that sends smart reminders based on calendar events",

	// Dev Tools
	"a VS Code extension that shows git blame inline",
	"a code review bot that checks for common mistakes",
	"a tool that generates API documentation from code comments",
	"a local development HTTPS proxy with auto-certificates",
	"a tool that visualizes your dependency graph as a web app",
	"a linter that checks for accessibility issues in JSX",
	"a tool that benchmarks your CI pipeline step by step",
	"a code snippet manager with a CLI and web UI",
	"a tool that detects dead code across a monorepo",
	"a local mock server that serves fixtures from a folder",
	"a tool that generates database migrations from schema diffs",
	"a script that audits your npm dependencies for licenses",
	"a tool that auto-generates TypeScript types from API responses",
	"a pre-commit hook that enforces consistent file naming",
	"a tool that profiles startup time of Node.js apps",
	"a script that converts Swagger specs to Postman collections",
	"a tool that enforces PR size limits and suggests splitting",
	"a local tunnel service like ngrok but self-hosted",
	"a tool that checks for hardcoded secrets in your codebase",
	"a script that generates architecture diagrams from code",

	// Games & Fun
	"a terminal-based roguelike dungeon crawler",
	"a text adventure game engine with a simple scripting language",
	"a multiplayer word game over WebSockets",
	"a browser-based rhythm game with keyboard input",
	"a terminal chess game with an AI opponent",
	"a Wordle clone with multiple difficulty modes",
	"a browser game where you manage a virtual startup",
	"a terminal snake game with increasing speed",
	"a typing speed trainer with code snippets",
	"a browser-based tower defense game on a grid",
	"a terminal-based trivia game that pulls from an API",
	"a multiplayer drawing game in the browser",
	"a text-based stock market simulator",
	"a terminal Tetris clone with color support",
	"a browser game where you debug intentionally broken code",
	"a daily crossword generator from a custom word list",
	"a terminal-based Minesweeper with high score tracking",
	"a browser-based typing game with code completion",
	"a terminal ASCII art animation framework",
	"a multiplayer quiz app with live leaderboard",

	// AI & ML Projects
	"a local chatbot that uses Ollama as the backend",
	"a tool that auto-tags photos using a local vision model",
	"a writing assistant that suggests continuations offline",
	"a CLI tool that summarizes long documents locally",
	"a tool that detects duplicate content across a folder of text files",
	"a local spell checker trained on your own writing style",
	"a tool that classifies incoming emails into categories",
	"a script that generates alt text for images automatically",
	"a local search engine for your personal notes",
	"a tool that extracts action items from meeting transcripts",
	"a sentiment analyzer for customer feedback spreadsheets",
	"a tool that generates commit messages from diffs",
	"a local voice-to-text tool that runs without the cloud",
	"a tool that clusters similar GitHub issues automatically",
	"a script that generates test cases from function signatures",
	"a tool that rewrites code comments to be clearer",
	"a local image search engine using embeddings",
	"a tool that detects plagiarism across student submissions",
	"a browser extension that summarizes the current page",
	"a tool that turns your highlights into a spaced repetition deck",

	// Mobile / Cross-platform
	"a cross-platform habit tracker built with Flutter",
	"a mobile app that scans receipts and extracts totals",
	"a React Native app for tracking daily water intake",
	"a mobile app that identifies plants from photos",
	"a cross-platform timer app with custom interval presets",
	"a mobile app that generates workout plans from your schedule",
	"a React Native clipboard manager with search",
	"a mobile app for tracking hiking routes offline",
	"a cross-platform expense tracker with currency conversion",
	"a mobile app that reads QR codes and saves the history",

	// Homelab / Self-hosted
	"a self-hosted bookmark manager with full-text search",
	"a personal dashboard that aggregates all your services",
	"a self-hosted URL shortener with analytics",
	"a home media server that generates thumbnails automatically",
	"a self-hosted password manager with a clean web UI",
	"a Raspberry Pi dashboard that shows home sensor data",
	"a self-hosted git hosting platform in under 500 lines",
	"a home automation hub that triggers webhooks on events",
	"a self-hosted photo gallery with face recognition",
	"a personal search engine that indexes your local files",
	"a self-hosted read-later app with browser extension",
	"a home network monitor that alerts on new devices",
	"a self-hosted email digest service for newsletters",
	"a personal wiki engine with markdown and backlinks",
	"a self-hosted uptime monitor with a status page",

	// Data & Visualization
	"a tool that visualizes your GitHub contribution graph in 3D",
	"a personal analytics dashboard from browser history",
	"a tool that turns your Spotify history into listening reports",
	"a script that visualizes your email response time patterns",
	"a tool that maps your file system usage as a treemap",
	"a dashboard for tracking your personal OKRs",
	"a tool that visualizes sleep data from a fitness tracker CSV",
	"a script that graphs your commit activity over time",
	"a personal finance chart generator from bank statements",
	"a tool that visualizes network latency over time",

	// Automation & Bots
	"a Telegram bot that summarizes news from your chosen sources",
	"a Discord bot that tracks server activity stats",
	"a bot that posts a daily challenge to a Slack channel",
	"a script that auto-organizes your Downloads folder by file type",
	"a bot that monitors a subreddit and notifies on keywords",
	"a script that backs up all your dotfiles to a git repo automatically",
	"a Telegram bot that converts currency on demand",
	"a bot that reminds you to drink water every hour",
	"a script that auto-generates weekly reports from Jira",
	"a bot that posts your daily schedule to a Slack channel",
	"a script that renames photo files by EXIF date automatically",
	"a bot that monitors stock prices and alerts on thresholds",
	"a script that syncs your local notes to a git repo",
	"a bot that schedules your tweets for optimal engagement times",
	"a script that generates a reading list from your bookmarks",

	// Educational / Learning Tools
	"a spaced repetition app for learning command-line shortcuts",
	"a tool that generates coding challenges from your weak areas",
	"a web app that teaches regex through interactive puzzles",
	"a tool that quizzes you on keyboard shortcuts for your IDE",
	"a browser app for practicing mental math",
	"a tool that teaches Git through an interactive simulation",
	"a web app for learning touch typing with custom texts",
	"a tool that generates SQL exercises from a real schema",
	"a browser game that teaches binary numbers through puzzles",
	"a tool that creates mnemonics for things you need to memorize",

	// Security & Privacy
	"a tool that audits browser extension permissions",
	"a script that checks your email against known breaches",
	"a local VPN kill-switch that blocks traffic when VPN drops",
	"a tool that generates strong passphrases from word lists",
	"a script that removes EXIF data from photos in bulk",
	"a tool that monitors for changes on websites you care about",
	"a browser extension that warns about dark patterns",
	"a script that audits your SSH authorized keys",
	"a tool that encrypts specific files in a folder transparently",
	"a personal threat model builder as a web app",

	// Open Source / Community
	"a tool that helps maintainers triage GitHub issues faster",
	"a web app that matches contributors to open source projects",
	"a bot that welcomes first-time contributors to your repo",
	"a tool that generates good first issue tickets from a backlog",
	"a web app for crowdsourcing translations of open source docs",
	"a tool that summarizes a project's changelog in plain language",
	"a bot that auto-labels GitHub issues by content",
	"a tool that generates contributor stats for your org",
	"a web app for coordinating open source sprints",
	"a tool that checks if your README follows best practices",

	// Niche / Creative
	"a tool that generates procedural dungeon maps as ASCII art",
	"a web app that creates generative album cover art",
	"a CLI tool that writes poetry using a Markov chain",
	"a tool that converts code to a different programming language style",
	"a web app that generates realistic fake social media profiles",
	"a tool that creates crossword puzzles from your own vocabulary list",
	"a CLI tool that generates haiku from any text",
	"a web app that turns a URL into a printable zine",
	"a tool that generates creative project names from themes",
	"a web app that creates procedural pixel art sprites",
	"a CLI tool that translates error messages to different tones",
	"a tool that generates fictional world maps as SVG",
	"a web app that creates ASCII art from any image",
	"a tool that generates lorem ipsum in different programming languages",
	"a CLI tool that creates animated ASCII banners",
	"a web app that generates tarot card readings from code errors",
	"a tool that converts sheet music to ASCII tablature",
	"a web app for collaborative storytelling with branching paths",
	"a CLI tool that generates startup names from syllable rules",
	"a web app that makes generative art from your typing speed",

	// Infrastructure / DevOps
	"a self-hosted status page that reads from your monitoring tools",
	"a tool that diffs Kubernetes manifests across environments",
	"a script that auto-scales servers based on a schedule",
	"a tool that generates Terraform configs from a web UI",
	"a dashboard that shows costs broken down by service and team",
	"a tool that validates Helm charts against your policies",
	"a script that rotates credentials on a schedule automatically",
	"a tool that generates runbooks from incident postmortems",
	"a dashboard showing deploy frequency and lead time metrics",
	"a tool that tests your disaster recovery playbooks automatically",

	// Productivity
	"a tool that blocks distracting sites during focus sessions",
	"a web app for tracking the books you want to read",
	"a personal CRM to remember details about your contacts",
	"a tool that generates weekly reviews from your task list",
	"a web app for tracking goals with a visual progress timeline",
	"a tool that turns meeting notes into structured action items",
	"a browser extension that saves reading position on any page",
	"a web app for planning your week using time blocking",
	"a tool that generates a daily briefing from your inbox",
	"a personal changelog where you log what you shipped each week",
	"a web app for tracking which ideas you want to explore",
	"a tool that converts voice memos to organized notes",
	"a browser extension that estimates reading time for pages",
	"a web app for tracking which restaurants you want to try",
	"a tool that generates retrospective prompts for your team",

	// Hardware / IoT
	"a Raspberry Pi dashboard showing live CPU and temp graphs",
	"a tool that logs and visualizes data from a soil moisture sensor",
	"a home weather station that uploads data to a personal site",
	"a script that controls a smart plug based on your calendar",
	"a tool that reads from a CO2 sensor and alerts above threshold",
	"a Raspberry Pi clock that shows your calendar on an e-ink display",
	"a tool that tracks package deliveries and alerts via notification",
	"a home plant watering system with a web dashboard",
	"a script that monitors your 3D printer and sends photo updates",
	"a tool that auto-dims your monitor based on ambient light",

	// Miscellaneous but Brilliant
	"a tool that generates a diff between two grocery lists",
	"a web app that turns a GitHub repo into a visual timeline",
	"a tool that generates a project skeleton from a description",
	"a web app for tracking which movies you want to watch together",
	"a tool that finds the best meeting time across time zones",
	"a script that generates a personal year in review from your data",
	"a web app for planning a trip itinerary collaboratively",
	"a tool that converts a job description into a checklist of skills",
	"a web app that generates a reading schedule for a list of books",
	"a tool that turns any Wikipedia article into a quiz",
	"a web app for tracking which podcasts you want to listen to",
	"a tool that generates a meal plan from ingredients you have",
	"a web app for coordinating Secret Santa gift exchanges",
	"a tool that converts a PDF resume to a clean web portfolio",
	"a web app that creates a printable contact sheet from photos",
	"a tool that tracks how long you spend in meetings each week",
	"a web app for building and sharing custom keyboard shortcuts",
	"a tool that generates a bingo card from a list of items",
	"a web app for tracking which albums you want to listen to",
	"a tool that turns a list of tasks into an estimated timeline",
	"a web app for tracking your personal records in any activity",
	"a tool that generates a family tree from a structured text file",
	"a web app for planning a wedding seating arrangement",
	"a tool that converts a CSV of contacts to vCard format",
	"a web app for tracking which games you want to play",
	"a tool that generates a gift wishlist page from product URLs",
	"a script that backs up your browser bookmarks to markdown",
	"a web app for tracking your personal knowledge base evolution",
	"a tool that generates a daily journaling prompt from keywords",
	"a web app for coordinating a book club reading schedule",
}

var charset = "abcdefghijklmnopqrstuvwxyz abcdefghijklmnopqrstuvwxyz    !!??--"

func randomChar() string {
	return string(charset[rand.Intn(len(charset))])
}

func moveCursor(row, col int) string {
	return fmt.Sprintf("\033[%d;%dH", row, col)
}

func clearLine() string {
	return "\033[2K"
}

// The header occupies rows 1-2; the monkey art starts at row 3.
const artTopRow = 3

var animFrames = []string{frame1, frame2}

func frameLines(frame string) []string {
	return strings.Split(strings.TrimRight(frame, "\n"), "\n")
}

// paintArt redraws just the monkey rows in place (fixed positions), so the two
// frames can be swapped to animate without clearing and scrolling the screen.
func paintArt(frame string) {
	for i, line := range frameLines(frame) {
		fmt.Print(moveCursor(artTopRow+i, 1))
		fmt.Print(clearLine())
		fmt.Printf("%s%s%s", yellow, line, reset)
	}
}

// drawMonkey paints the header and first frame once, then parks the cursor on
// the (empty) typing line and saves that position so the typing line can be
// updated in place.
func drawMonkey() {
	fmt.Print(clearScreen)
	fmt.Print(moveCursor(1, 1))
	fmt.Printf("%s THE INFINITE MONKEY%s", bold+yellow, reset)
	fmt.Print(moveCursor(2, 1))
	fmt.Printf("%s--------------------------------------------------%s", gray, reset)
	paintArt(frame1)
	n := len(frameLines(frame1))
	fmt.Print(moveCursor(artTopRow+n+1, 1))
	fmt.Printf("%s Typing:%s", gray, reset)
	fmt.Print(moveCursor(artTopRow+n+2, 1))
	fmt.Print(saveCursor)
}

// updateTyping rewrites just the typing line in place.
func updateTyping(line string) {
	fmt.Print(restoreCursor)
	fmt.Print(clearLine())
	fmt.Printf("\r  %s", line)
}

func animateTyping(target string) {
	fmt.Print(hideCursor)
	drawMonkey()

	// Alternate the two frames as the monkey "types". swap only repaints the
	// art when the frame actually changes, and updateTyping returns the cursor
	// to the typing line afterwards.
	cur := 0
	step := 0
	tick := func(line string) {
		if step%2 != cur {
			cur = step % 2
			paintArt(animFrames[cur])
		}
		step++
		updateTyping(line)
	}

	// Phase 1: chaotic random typing
	displayed := ""
	chaosDuration := len(target)*2 + rand.Intn(len(target))
	for i := 0; i < chaosDuration; i++ {
		if len(displayed) > 0 && rand.Float32() < 0.15 {
			displayed = displayed[:len(displayed)-1]
		} else {
			displayed += randomChar()
			if len(displayed) > 50 {
				displayed = displayed[len(displayed)-50:]
			}
		}
		tick(fmt.Sprintf("%s%s%s%s_", dim+gray, displayed, reset, cyan))
		time.Sleep(time.Duration(20+rand.Intn(60)) * time.Millisecond)
	}

	// Phase 2: type the actual idea with occasional typos
	realTyped := ""
	for _, ch := range target {
		if rand.Float32() < 0.08 && len(realTyped) > 0 {
			preview := realTyped + randomChar()
			if len(preview) > 55 {
				preview = "..." + preview[len(preview)-52:]
			}
			tick(fmt.Sprintf("%s%s%s_", white, preview, reset))
			time.Sleep(80 * time.Millisecond)

			disp := realTyped
			if len(disp) > 55 {
				disp = "..." + disp[len(disp)-52:]
			}
			tick(fmt.Sprintf("%s%s%s_", white, disp, reset))
			time.Sleep(60 * time.Millisecond)
		}

		realTyped += string(ch)
		disp := realTyped
		if len(disp) > 55 {
			disp = "..." + disp[len(disp)-52:]
		}
		tick(fmt.Sprintf("%s%s%s_", white, disp, reset))
		time.Sleep(time.Duration(30+rand.Intn(70)) * time.Millisecond)
	}

	updateTyping(fmt.Sprintf("%s%s%s", bold+cyan, target, reset))
	time.Sleep(300 * time.Millisecond)
}

func showChoices(idea string) string {
	fmt.Print(clearScreen)
	typeRow := 1

	// Draw the full idea box
	fmt.Print(moveCursor(typeRow+3, 1))
	fmt.Printf("%s+----------------------------------------------------%s\n", gray, reset)
	fmt.Print(moveCursor(typeRow+4, 1))
	wrapped := wrapText(idea, 50)
	for i, line := range wrapped {
		fmt.Print(moveCursor(typeRow+4+i, 1))
		fmt.Printf("%s|%s  %-50s  %s|%s\n", gray, bold+white, line, reset, gray+reset)
	}
	extraRows := len(wrapped) - 1
	fmt.Print(moveCursor(typeRow+5+extraRows, 1))
	fmt.Printf("%s+----------------------------------------------------%s\n", gray, reset)

	fmt.Print(moveCursor(typeRow+7+extraRows, 1))
	fmt.Printf("%sWhat do you want to do?%s\n\n", dim+white, reset)

	fmt.Print(moveCursor(typeRow+9+extraRows, 1))
	fmt.Printf("  %s[y]%s %sYes, I'll build this!%s\n", bold+green, reset, white, reset)
	fmt.Print(moveCursor(typeRow+10+extraRows, 1))
	fmt.Printf("  %s[m]%s %sMaybe later%s\n", bold+yellow, reset, white, reset)
	fmt.Print(moveCursor(typeRow+11+extraRows, 1))
	fmt.Printf("  %s[n]%s %sNah, next idea%s\n", bold+red, reset, white, reset)
	fmt.Print(moveCursor(typeRow+12+extraRows, 1))
	fmt.Printf("  %s[q]%s %sQuit%s\n", bold+magenta, reset, white, reset)

	fmt.Print(moveCursor(typeRow+14+extraRows, 1))
	fmt.Printf("%s> %s", bold+cyan, reset)

	reader := bufio.NewReader(os.Stdin)

	for {
		// Make terminal raw-ish by reading a line
		input, err := reader.ReadString('\n')
		if err != nil {
			return "q"
		}
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "" {
			input = "n"
		}
		switch input {
		case "y", "yes":
			return "y"
		case "m", "maybe", "later":
			return "m"
		case "n", "no", "next":
			return "n"
		case "q", "quit", "exit":
			return "q"
		default:
			fmt.Print(moveCursor(typeRow+14+extraRows, 1))
			fmt.Print(clearLine())
			fmt.Printf("%s  Type y / m / n / q then Enter%s\n", gray, reset)
			fmt.Print(moveCursor(typeRow+15+extraRows, 1))
			fmt.Printf("%s> %s", bold+cyan, reset)
		}
	}
}

func wrapText(text string, width int) []string {
	words := strings.Fields(text)
	var lines []string
	current := ""
	for _, w := range words {
		if len(current)+len(w)+1 > width {
			if current != "" {
				lines = append(lines, current)
			}
			current = w
		} else {
			if current == "" {
				current = w
			} else {
				current += " " + w
			}
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func showYesList(list []string) {
	fmt.Print(clearScreen)
	fmt.Print(moveCursor(1, 1))
	fmt.Printf("%s YOUR BUILD LIST%s\n", bold+green, reset)
	fmt.Print(moveCursor(2, 1))
	fmt.Printf("%s--------------------------------------------------%s\n", gray, reset)
	for i, item := range list {
		fmt.Print(moveCursor(3+i, 1))
		fmt.Printf("  %s%d.%s %s%s%s\n", bold+cyan, i+1, reset, white, item, reset)
	}
	fmt.Print(moveCursor(4+len(list), 1))
	fmt.Printf("%s--------------------------------------------------%s\n", gray, reset)
}

func showMaybeList(list []string) {
	if len(list) == 0 {
		return
	}
	offset := 0
	fmt.Print(moveCursor(20+offset, 1))
	fmt.Printf("\n%s MAYBE LATER%s\n", dim+yellow, reset)
	for _, item := range list {
		fmt.Printf("  %s- %s%s\n", gray, item, reset)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Draw on the alternate screen so the animation never pollutes the
	// terminal's scrollback; the original screen is restored on exit.
	fmt.Print(enterAltScreen)

	// Restore the terminal if interrupted (Ctrl-C) instead of leaving the
	// user stranded on the alternate screen with a hidden cursor.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Print(showCursor)
		fmt.Print(exitAltScreen)
		os.Exit(0)
	}()

	// Shuffle ideas
	rand.Shuffle(len(ideas), func(i, j int) { ideas[i], ideas[j] = ideas[j], ideas[i] })

	var yesList []string
	var maybeList []string

	idx := 0
	for {
		if idx >= len(ideas) {
			rand.Shuffle(len(ideas), func(i, j int) { ideas[i], ideas[j] = ideas[j], ideas[i] })
			idx = 0
		}

		idea := ideas[idx]
		idx++

		animateTyping(idea)
		choice := showChoices(idea)

		switch choice {
		case "y":
			yesList = append(yesList, idea)
			fmt.Print(clearScreen)
			fmt.Print(moveCursor(1, 1))
			fmt.Printf("\n  %s Let's go! Added to your build list.%s\n", bold+green, reset)
			time.Sleep(600 * time.Millisecond)

		case "m":
			maybeList = append(maybeList, idea)
			fmt.Print(clearScreen)
			fmt.Print(moveCursor(1, 1))
			fmt.Printf("\n  %s Saved for later.%s\n", bold+yellow, reset)
			time.Sleep(400 * time.Millisecond)

		case "n":
			// just continue

		case "q":
			// Leave the alternate screen first so the summary prints on the
			// real terminal and stays visible after the program exits.
			fmt.Print(exitAltScreen)
			fmt.Print(showCursor)
			fmt.Printf("%s The monkey takes a break.%s\n\n", bold+yellow, reset)

			if len(yesList) > 0 {
				showYesList(yesList)
			}
			if len(maybeList) > 0 {
				showMaybeList(maybeList)
			}
			if len(yesList) == 0 && len(maybeList) == 0 {
				fmt.Printf("  %sYou didn't save anything. Come back anytime!%s\n", gray, reset)
			}
			fmt.Println()
			return
		}
	}
}
