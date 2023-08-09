import { fileURLToPath, URL } from 'node:url';

import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { quasar, transformAssetUrls } from '@quasar/vite-plugin';
import monacoEditorPlugin from 'vite-plugin-monaco-editor';

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      // '/stomp': {
      //   target: 'ws://127.0.0.1:8888',
      //   changeOrigin: true,
      //   ws: true
      // },
      // '/player': {
      //   target: 'ws://127.0.0.1:8888',
      //   changeOrigin: true,
      //   ws: true
      // },
      '/ws': {
        target: 'ws://127.0.0.1:8877',
        changeOrigin: true,
        ws: true
      },
      '/ue': {
        target: 'http://127.0.0.1:8877',
        // changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  },
  plugins: [
    vue({
      template: { transformAssetUrls }
    }),
    quasar({
      sassVariables: 'src/quasar-variables.sass'
    }),
    monacoEditorPlugin({
      languageWorkers: ['editorWorkerService', 'css', 'html', 'json', 'typescript']
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
});
