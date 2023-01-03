# Go Svelte SPA

- A sample fullstack application with Go as server (spa files + API if any).
- `frontend` folder:
  - contains a Vite project with Svelte app and has TailwindCSS integration.
  - is also a Go package `frontend` that allows embedding the complete svelte build into `frontend.DistFS` variable.
- `make` will build frontend and then trigger go-build which embeds the frontend files and produces one final executable.

## Usage

1. Clone the repository.
2. Replace all instances of `go-svelte` string with your app name.
3. In the root directory, run `make` to build everything (frontend + backend).
4. A single executable binary will be generated in `dist/`
