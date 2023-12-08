
//copy from thingjs
var lut = [];

for (var i$1 = 0; i$1 < 256; i$1++) {
    lut[i$1] = (i$1 < 16 ? '0' : '') + i$1.toString(16).toUpperCase();
}
function generateUUID() {
    
    var simpilfy = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : false;
    var d0 = Math.random() * 0xffffffff | 0;
    var d1 = Math.random() * 0xffffffff | 0;
    var d2 = Math.random() * 0xffffffff | 0;
    var d3 = Math.random() * 0xffffffff | 0;
    var uuid = lut[d0 & 0xff] + lut[d0 >> 8 & 0xff] + lut[d0 >> 16 & 0xff] + lut[d0 >> 24 & 0xff] + '-' + lut[d1 & 0xff] + lut[d1 >> 8 & 0xff] + '-' + lut[d1 >> 16 & 0x0f | 0x40] + lut[d1 >> 24 & 0xff] + '-' + lut[d2 & 0x3f | 0x80] + lut[d2 >> 8 & 0xff] + '-' + lut[d2 >> 16 & 0xff] + lut[d2 >> 24 & 0xff] + lut[d3 & 0xff] + lut[d3 >> 8 & 0xff] + lut[d3 >> 16 & 0xff] + lut[d3 >> 24 & 0xff];

    if (simpilfy) {
        // @ts-ignore
        uuid = uuid._replaceAll('-', '');
    }

    return uuid;
}

var CommandMapCache = {};

async function SendUserCommand(command, param){
    return new Promise(function(resolve, reject){
        var uuid = generateUUID();
        console.debug({
            type: "UserCommand",
            uuid: uuid,
            command: command,
            param: param
        })
        emitUIInteraction({
            type: "UserCommand",
            uuid: uuid,
            command: command,
            param: param
        });
        CommandMapCache[uuid] = {resolve: resolve, reject: reject};
    });
}

window.addEventListener("load", function(){
    addResponseEventListener("thingjs", function(response){
        var obj = JSON.parse(response);
        if(obj.type == "UserCommandRet"){
            var uuid = obj.uuid;
            if(CommandMapCache[uuid]){
                CommandMapCache[uuid].resolve(obj.param);
                delete CommandMapCache[uuid];
            }
        }else if(obj.type == "UserEvent"){
            var command = obj.command;
            var param = obj.param;
            if(OnUserEvent){
                OnUserEvent(command, param);
            }
        }
    })
})