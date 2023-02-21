# Go Svelte SPA

- A sample fullstack application with Go as backend (Responsible for serving frontend static files + API if any).
- `ui` folder:
  - contains a multi-page Vite project with Svelte (not SvelteKit)
    - has TailwindCSS + DaisyUI
    - `src/index.html` is static (useful as landing page)
    - `src/app/index.html` is SPA and uses `svelte-spa-router` for Routing (hash-based)
  - is also a Go package `ui` that embeds svelte build and exposes `Handler()` which can be mounted on any router for serving frontend. 
- `make` will build frontend and then trigger go-build which embeds the frontend files and produces one final executable.
- A 3-stage `Dockerfile` is included that builds frontend, backend and then produces a very small final alpine image with just the binary. 

## Usage

1. Clone the repository.
2. Replace all instances of `go-svelte` string with your app name.
3. In the root directory, run `make` to build everything (frontend + backend).
4. A single executable binary will be generated in `dist/`
