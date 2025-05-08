import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig(({ mode }) => ({
  plugins: [vue()],
  base: '/', // zmień jeśli serwujesz z podkatalogu
  server: {
    host: '0.0.0.0',
    port: 80,
    strictPort: true,
    cors: true,
    proxy: {
      '/api': {
        target: 'http://server:8437',
        changeOrigin: true,
      }
    }
  },
  build: {
    outDir: 'dist',
    sourcemap: false,
  }
}))