import {Config, PixelStreaming} from '@thingue/lib-pixelstreamingfrontend';
import {Application, PixelStreamingApplicationStyle} from '@thingue/lib-pixelstreamingfrontend-ui';

const PixelStreamingApplicationStyles = new PixelStreamingApplicationStyle();
PixelStreamingApplicationStyles.applyStyleSheet();

document.body.onload = function () {
    const config = new Config({
        initialSettings: {
            AutoPlayVideo: true,
            AutoConnect: true,
            OfferToReceive: true,
            HoveringMouse: true,
            StartVideoMuted: true,
            MatchViewportRes: true,
        }
    });
    const stream = new PixelStreaming(config, {
        playerUrlBuilder: playerUrlBuilder
    });
    const application = new Application({
        stream,
        onColorModeChanged: (isLightMode) => PixelStreamingApplicationStyles.setColorMode(isLightMode)
    });
    document.body.appendChild(application.rootElement);
}

async function playerUrlBuilder() {
    const origin = window.location.origin.replace('http://', 'ws://').replace('https://', 'wss://');
    const path = window.location.pathname.slice(0, location.pathname.lastIndexOf("/")).replace("/static", "");
    const urlParams = new URLSearchParams(window.location.search);
    if (urlParams.has("ticket")) {
        return `${origin}${path}/ws/player/${urlParams.get("ticket")}`;
    }
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const response = await fetch(path + "/api/instance/ticketSelect", {
        method: 'POST',
        headers: myHeaders,
        body: JSON.stringify({
            sid: urlParams.get("sid"),
            name: urlParams.get("name"),
            playerCount: urlParams.get("playerCount") ? Number(urlParams.get("playerCount")) : -1,
            labelSelector: urlParams.get("labelSelector"),
        }),
    })
    const resJson = await response.json()
    if (resJson.code === 200) {
        return `${origin}${path}/ws/player/${resJson.data.ticket}`;
    } else {
        throw new Error(resJson.msg);
    }
}