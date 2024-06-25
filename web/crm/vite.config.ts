import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"
import path from "path"

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
    resolve: {
        alias: [
            { find: "api", replacement: path.resolve(__dirname, "./api") },
            { find: "features", replacement: path.resolve(__dirname, "./src/features") },
            { find: "pages", replacement: path.resolve(__dirname, "./src/pages") },
            { find: "shared", replacement: path.resolve(__dirname, "./src/shared") },
            { find: "widgets", replacement: path.resolve(__dirname, "./src/widgets") },
        ],
    },
})
