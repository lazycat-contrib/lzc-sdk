import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  server: {
    watch: {
      ignored: ["**/node_modules/**"],
    },
  },
  build: {
    chunkSizeWarningLimit: 2000,
  },
  plugins: [react()],
});
