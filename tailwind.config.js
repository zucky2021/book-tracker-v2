/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,ts,jsx,tsx}", // 必要なファイルを指定
  ],
  theme: {
    extend: {
      animation: {
        spin: "spin 0.8s linear infinite", // カスタムスピンアニメーション
      },
      keyframes: {
        spin: {
          from: { transform: "rotate(0deg)" },
          to: { transform: "rotate(359deg)" },
        },
      },
      colors: {
        cyan: {
          100: "rgba(63, 249, 220, 0.1)", // カスタムカラー
          500: "rgb(63, 249, 220)",
        },
      },
    },
  },
  plugins: [],
};
