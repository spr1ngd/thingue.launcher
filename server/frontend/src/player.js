import {Config, PixelStreaming} from '@thingue/lib-pixelstreamingfrontend';
import {Application, PixelStreamingApplicationStyle} from '@thingue/lib-pixelstreamingfrontend-ui';
import {v4 as uuidv4} from "uuid";

const PixelStreamingApplicationStyles = new PixelStreamingApplicationStyle();
PixelStreamingApplicationStyles.applyStyleSheet();
const urlParams = new URLSearchParams(window.location.search);

const idleController = {
    timer: null,
    stream: null,
    timeout: null,
    resetTimer: function () {
        clearTimeout(idleController.timer);
        idleController.timer = setTimeout(idleController.disconnectStream, idleController.timeout);
    },
    disconnectStream: function () {
        idleController.stream.disconnect()
    },
    startListener: function () {
        window.addEventListener('mousemove', idleController.resetTimer);
        window.addEventListener('keydown', idleController.resetTimer);
        idleController.resetTimer()
    },
    stopListener: function () {
        window.removeEventListener('mousemove', idleController.resetTimer);
        window.removeEventListener('keydown', idleController.resetTimer);
        clearTimeout(idleController.timer);
    }
}

document.body.onload = function () {
    const config = new Config({
        initialSettings: {
            AutoPlayVideo: true,
            AutoConnect: true,
            OfferToReceive: true,
            HoveringMouse: true,
            StartVideoMuted: true,
            MatchViewportRes: false,
            SuppressBrowserKeys: false,
        },
        useUrlParams: true
    });
    const stream = new PixelStreaming(config, {
        playerUrlBuilder: playerUrlBuilder
    });

    idleController.stream = stream

    window.onmessage = async function (e) {
        const uuid = uuidv4();
        stream.emitUIInteraction({
            type: "UserCommand",
            uuid,
            command: "ToUEMessage",
            param: e.data
        })
    }

    stream.addResponseEventListener("user_handler", function (response) {
        if (window.top !== window) {
            window.top.postMessage(JSON.parse(response).param, "*");
        }
    })

    const hideUi = urlParams.has("HideUI") && (urlParams.get("HideUI") === "true" || urlParams.get("HideUI") === "True")

    const application = new Application({
        stream,
        onColorModeChanged: (isLightMode) => PixelStreamingApplicationStyles.setColorMode(isLightMode),
        // 隐藏UI上的控制元素配置
        settingsPanelConfig: {
            isEnabled: true,
            visibilityButtonConfig: {
                creationMode: hideUi ? 2 : 0
            }
        },
        statsPanelConfig: {
            isEnabled: true,
            visibilityButtonConfig: {
                creationMode: hideUi ? 2 : 0
            }
        },
        fullScreenControlsConfig: {
            creationMode: hideUi ? 2 : 0
        },
        videoQpIndicatorConfig: {
            disableIndicator: hideUi
        },
    });
    document.body.appendChild(application.rootElement);

}

async function playerUrlBuilder() {
    const origin = window.location.origin.replace('http://', 'ws://').replace('https://', 'wss://');
    const path = window.location.pathname.slice(0, location.pathname.lastIndexOf("/")).replace("/static", "");
    if (urlParams.has("ticket")) {
        return `${origin}${path}/ws/player/${urlParams.get("ticket")}`;
    }
    const response = await fetch(path + "/api/instance/ticketSelect", {
        method: 'POST',
        headers: [["Content-Type", "application/json"]],
        body: JSON.stringify({
            sid: urlParams.get("sid"),
            name: urlParams.get("name"),
            playerCount: urlParams.get("playerCount") ? Number(urlParams.get("playerCount")) : -1,
            labelSelector: urlParams.get("labelSelector"),
        }),
    })
    const resJson = await response.json()
    if (resJson.code === 200) {
        // 获取player配置
        const playerConfig = resJson.data.playerConfig
        idleController.stream.config.setFlagEnabled('MatchViewportRes', playerConfig.matchViewportRes)
        document.getElementById("controls").style.visibility = playerConfig.hideUI ? "hidden" : "visible"
        document.getElementById("connection").style.visibility = playerConfig.hideUI ? "hidden" : "visible"
        if (playerConfig.idleDisconnect) {
            idleController.timeout = playerConfig.idleTimeout * 60000
            idleController.startListener()
        } else {
            idleController.stopListener()
        }
        return `${origin}${path}/ws/player/${resJson.data.ticket}`;
    } else {
        throw new Error(resJson.msg);
    }
}