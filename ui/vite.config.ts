import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { resolve } from 'path'

const root = resolve(__dirname, 'src')
const outDir = resolve(__dirname, 'dist')

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    proxy: {
      "/api": "http://localhost:8080"
    },
  },
  resolve: {
    alias: {
      "$lib": resolve(root, 'lib'),
    }
  },
  root: "src",
  publicDir: resolve(__dirname, "public/"),
  build: {
    outDir,
    copyPublicDir: true,
    emptyOutDir: true,
    rollupOptions: {
      input: {
        main: resolve(root, "index.html"),
        app: resolve(root, "app/index.html")
      }
    }
  }
})
