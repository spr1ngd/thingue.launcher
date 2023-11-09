import {createApp} from 'vue'
import {Dialog, Notify, Quasar} from 'quasar'
import quasarLang from 'quasar/lang/zh-CN'

import '@quasar/extras/material-icons/material-icons.css'
import '@quasar/extras/material-symbols-outlined/material-symbols-outlined.css'
import '@quasar/extras/material-symbols-rounded/material-symbols-rounded.css'
import '@quasar/extras/material-symbols-sharp/material-symbols-sharp.css'

import 'quasar/src/css/index.sass'

import App from './App.vue'

const myApp = createApp(App)

myApp.use(Quasar, {
    plugins: {Dialog, Notify},
    lang: quasarLang
})

myApp.mount('#app')
