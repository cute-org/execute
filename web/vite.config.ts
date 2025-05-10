import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  base: '/', // lub inna ścieżka, jeśli serwujesz z podkatalogu
  server: {
    host: '0.0.0.0',
    port: 5173,  // Port w trybie deweloperskim
    strictPort: true,
    cors: true,
    proxy: {
      '/api': {
        target: 'http://server:8437', // API w trybie deweloperskim
        changeOrigin: true,
      }
    }
  },
  build: {
    outDir: 'dist',
    sourcemap: false,
  }
})
