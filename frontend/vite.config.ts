import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    host: "0.0.0.0", // コンテナ内でアクセスを可能に
    port: 3000,
    watch: {
      usePolling: true, // ファイルの変更検知方法としていpollingを使用
      interval: 1000,
    },
  },
});
