import { createApp } from 'vue';
import App from './App.vue';
import { Quasar,Notify } from 'quasar';

// Import icon libraries
import '@quasar/extras/material-icons/material-icons.css';
// Import Quasar css
import 'quasar/src/css/index.sass';
import quasarLang from 'quasar/lang/zh-CN'

// import stompClient from './userStomp';
import { Client } from '@stomp/stompjs';

var origin = window.location.origin.replace('http://', 'ws://').replace('https://', 'wss://')
var path = window.location.pathname.slice(0, location.pathname.lastIndexOf("/"))
var brokerURL = `${origin}${path}/ws/stomp`
const stompClient = new Client({
  brokerURL,
  heartbeatIncoming: 4000,
  heartbeatOutgoing: 4000
});

let app;

stompClient.onConnect = function (frame) {
  console.log('已建立STOMP连接');
  if (!app) {
    app = createApp(App);
    app.use(Quasar, {
      lang: quasarLang,
      plugins: {
        Notify
      } // import Quasar plugins and add here
    });

    app.provide('stompClient', stompClient);

    app.mount('#app');
  }
};

stompClient.onStompError = function (frame) {
  console.log('Broker reported error: ' + frame.headers['message']);
  console.log('Additional details: ' + frame.body);
}

stompClient.activate();
