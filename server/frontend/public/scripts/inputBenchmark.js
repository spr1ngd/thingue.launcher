
var intervalId = null;
var runningInputBenchmarking = false;
const mouseButtons = [0, 1, 2];

const randomInputPosition = function () {

    const left = 600;
    const top = 300;
    let x = left + Math.random() * 600;
    let y = top + Math.random() * 300;
    return {
        x,
        y
    };
}

const beginInputBenchmark = function () {

    runningInputBenchmarking = true;

    intervalId = setInterval(() => {

        if (runningInputBenchmarking) {

            let buttons = [];
            const num = Math.floor(Math.random() * 3);
            for (var i = 0; i < num; i++) {
                buttons.push(...mouseButtons);
            }

            buttons.forEach(button => {
                let { x, y } = randomInputPosition();
                console.log(`input ${x},${y}`);
                emitMouseDown(button, x, y);

                let { x1, y1 } = randomInputPosition();
                emitMouseMove(button, x1, y1);

                emitMouseUp(button, x, y);
            });
        }
    }, 8);
}

const endInputBenchmark = function () {

    runningInputBenchmarking = false;
    clearInterval(intervalId);
}