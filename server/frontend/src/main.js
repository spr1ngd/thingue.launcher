import {createApp} from 'vue'
import {Dialog, Notify, Quasar} from 'quasar'
import quasarLang from 'quasar/lang/zh-CN'
import {connectWebSocket} from "@/ws";
import router from './router'
import {createPinia} from 'pinia'

// Import icon libraries
import '@quasar/extras/material-icons/material-icons.css'
import '@quasar/extras/material-icons-round/material-icons-round.css'
import '@quasar/extras/material-symbols-outlined/material-symbols-outlined.css'

// Import Quasar css
import 'quasar/src/css/index.sass'
// Assumes your root component is App.vue
// and placed in same folder as main.js
import App from './App.vue'

const myApp = createApp(App)

myApp.use(router)
myApp.use(createPinia())
myApp.use(Quasar, {
    plugins: {
        Notify,
        Dialog
    }, // import Quasar plugins and add here
    lang: quasarLang,
})

// Assumes you have a <div id="app"></div> in your index.html
myApp.mount('#app')

connectWebSocket()
