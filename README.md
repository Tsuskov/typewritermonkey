# typewriter monkey

```
                                     ██████                             
                  ███     ████                          
                 ██  █████   ███                        
                 █       ██     ████               ████ 
                 ███     ██         ███           █ █   
                  ██                   ███        ███   
                  ███                     ██      ███   
                    ███████      ██         ██    ███   
                          ██      ██         ██   ███   
                           ██      ██         ██   ███  
 ██                       ████      ██         ██   ██  
  ██                 ██████  █      ████        █   ███ 
   ██            █████     ████      █  █       █    ██ 
   ███         ████  █████████      ██          █   ███ 
   ██████████████████     ██     ██████         █  ██ █ 
  ██ ██ █ ██ █████      ██     ███    █        ███████  
  █   ██        ████   ██   ██████     ██    ███████    
  █ █████████████████  ██  ███          █████           
  ███████████████████   ███████ ████████                
```

You know the thing about infinite monkeys and infinite typewriters eventually
producing Shakespeare? This monkey skips the infinite part and just hands you
a side project idea, one keystroke at a time, like it's typing it out for you
right there in your terminal.

Run it, watch the idea type itself out, then decide:

- **y** — yes, put it on the build list
- **m** — maybe, save it for later
- **n** — nah, next
- **q** — quit, and the monkey shows you everything you saved

That's it. No accounts, no database, no analytics. Just a monkey, a typewriter,
and a shuffled deck of ideas spanning CLI tools, web apps, APIs, and more.

## Running it

```
go run monkey.go
```

or build it:

```
go build -o typemonkey monkey.go
./typemonkey
```

## Installing it globally

Want to type `typemonkey` from any terminal, in any directory? Build it
straight into a folder that's already on your `PATH`:

```
mkdir -p ~/.local/bin
go build -o ~/.local/bin/typemonkey monkey.go
```

That's it — open a new terminal and run `typemonkey` from anywhere. The
monkey art is embedded into the binary at build time, so the resulting
executable is fully self-contained; you don't need to keep this repo around
afterwards.

If `~/.local/bin` isn't on your `PATH`, add this to your shell config
(`~/.zshrc` or `~/.bashrc`) and restart your terminal:

```
export PATH="$HOME/.local/bin:$PATH"
```

## Why

Sometimes the hardest part of starting a side project is just picking one.
This is a small, dumb, fun way to let chance make the first move — you still
get the final say.
