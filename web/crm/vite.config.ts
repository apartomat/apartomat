import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from "path"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: [
      {"find": "common", "replacement": path.resolve(__dirname, "./src/common")},
      {"find": "context", "replacement": path.resolve(__dirname, "./src/context")},
      {"find": "screen", "replacement": path.resolve(__dirname, "./src/screen")},
      {"find": "api", "replacement": path.resolve(__dirname, "./api")},
    ],
  },
})
