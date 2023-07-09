import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  // FIX: Fix aliases
  resolve: {
    alias: {
      '@': '/src',
      '@components': '/src/Components',
      '@assets': '/src/assets',
      '@imgs': '/src/assets/imgs',
    },
  },
  plugins: [react()],
})
