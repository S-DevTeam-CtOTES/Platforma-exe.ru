import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from "node:path";

// https://vitejs.dev/config/
export default defineConfig({
  // FIX: Fix aliases
  resolve: {
    alias: {
      "@": path.resolve(__dirname, './src'),
      "@components": path.resolve(__dirname, './src/Components'),
      "@assets": path.resolve(__dirname, './src/assets'),
      "@imgs": path.resolve(__dirname, './src/assets/imgs'),
    },
  },
  plugins: [react()],
})
