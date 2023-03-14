import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
	server: {
		watch: {
			exclude: 'node_modules/**',
			include: 'src/**',
		},
	},
	build: {
		chunkSizeWarningLimit: 2000,
	},
	plugins: [react()],
});
