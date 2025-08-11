import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'  // or your chosen frontend framework plugin

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Proxy all API calls starting with /api to localhost:8080 (your Go backend)
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  }
})
