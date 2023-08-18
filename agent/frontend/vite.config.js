import {fileURLToPath, URL} from 'node:url';
import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {quasar, transformAssetUrls} from '@quasar/vite-plugin'
import monacoEditorPlugin from 'vite-plugin-monaco-editor';

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url)),
            '@wails': fileURLToPath(new URL('./wailsjs', import.meta.url)),
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
    ],
    server: {
        host: '0.0.0.0',
        port: 7789
    }
})
