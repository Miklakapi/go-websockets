import WebSocket from "ws"
import readline from "readline"

const ws = new WebSocket("ws://localhost:8000/ws")

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
})

ws.on("open", () => {
    console.log("Connected to WebSocket")
    console.log("Type messages. Type 'exit' to close the client")

    rl.on("line", input => {
        if (input === "exit") {
            ws.close(1000, "Client exit")
            rl.close()
            return
        }
        ws.send(input)
    })
})

ws.on("message", data => {
    const msg = data.toString()

    const padding = 40 - msg.length
    const spaces = padding > 0 ? " ".repeat(padding) : " "
    console.log(spaces + msg)
})

ws.on("close", (code, reason) => {
    const textReason = reason instanceof Buffer ? reason.toString() || "<empty>" : String(reason || "<empty>")
    console.log("Connection closed. Code =", code, "Reason =", textReason)
    process.exit()
})

ws.on("error", err => {
    console.error("Error:", err)
    process.exit()
})
