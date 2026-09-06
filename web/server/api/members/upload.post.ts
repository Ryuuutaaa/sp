import { promises as fs } from "node:fs"
import { join } from "node:path"

const MAX_SIZE = 2 * 1024 * 1024 // 2MB
const ALLOWED_EXT = ["jpg", "jpeg", "png", "webp"]

export default defineEventHandler(async (event) => {
  const form = await readMultipartFormData(event)
  const file = form?.find((p) => p.name === "file")

  if (!file?.filename || !file.data) {
    throw createError({ statusCode: 400, message: "File wajib diisi" })
  }
  if (!file.type?.startsWith("image/")) {
    throw createError({ statusCode: 400, message: "Hanya file gambar yang diizinkan" })
  }
  if (file.data.length > MAX_SIZE) {
    throw createError({ statusCode: 400, message: "Ukuran file maksimal 2MB" })
  }

  const rawExt = file.filename.split(".").pop()?.toLowerCase() || "jpg"
  const ext = ALLOWED_EXT.includes(rawExt) ? rawExt : "jpg"
  const name = `${Date.now()}-${Math.random().toString(36).slice(2)}.${ext}`

  const dir = join(process.cwd(), "public", "uploads")
  await fs.mkdir(dir, { recursive: true })
  await fs.writeFile(join(dir, name), file.data)

  // TODO: pindah ke Cloudflare R2 (simpan R2 key, kembalikan public URL)
  return { url: `/uploads/${name}` }
})
