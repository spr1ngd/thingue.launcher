import {fileURLToPath, URL} from 'node:url';
import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {quasar, transformAssetUrls} from '@quasar/vite-plugin'
import monacoEditorPlugin from 'vite-plugin-monaco-editor';
import {resolve} from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    server: {
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8877'
            },
            '/ws': {
                target: 'ws://127.0.0.1:8877',
                ws: true
            }
        }
    },
    build: {
        assetsDir: 'internal-assets',
        rollupOptions: {
            input: {
                main: resolve(__dirname, 'index.html'),
                player: resolve(__dirname, 'player.html'),
            }
        }
    },
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url)),
        },
    },
    plugins: [
        vue({
            template: {transformAssetUrls}
        }),
        quasar({
            sassVariables: 'src/quasar-variables.sass'
        }),
        monacoEditorPlugin({
            languageWorkers: []
        })
    ]
})