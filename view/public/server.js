import { Application } from 'https://deno.land/x/abc/mod.ts'

const PORT = 5000
const app = new Application()

app.static("/assets", "assets")
app.static("/build", "build")
app.file("/", "index.html")

app.start({ port: PORT })
console.log(`Listening on port ${PORT}`)