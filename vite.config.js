import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import path from 'path'
import 'dotenv/config'

export default defineConfig({
	plugins: [vue()],
	define: {
		'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV)
	},
	build: {
		outDir: "public/js",
		lib: {
			entry: path.resolve(__dirname + '/resources/js/app.js'),
			name: 'app',
			fileName: (_, entryName) => `${entryName}.js`,
			formats: ['es']
		}
	}
})
