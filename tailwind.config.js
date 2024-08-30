/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./cmd/web/**/*.html", "./cmd/web/**/*.templ"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      {
        light: {
          primary: "#0070e3",
          secondary: "#13c4a3",
          accent: "#232c33",
          "accent-content": "#f3f4f6",
          neutral: "#232C33",
          "base-100": "#f3f4f6",
          info: "#a8a29e",
          success: "#a3e635",
          warning: "#fbbf24",
          error: "#dc2626",

          "--rounded-box": "1rem",
          "--rounded-btn": "0.4rem",
          "--rounded-badge": "1.9rem",
          "--animation-btn": "0.25s",
          "--btn-focus-scale": "0.95",
          "--border-btn": "1px",
          "--tab-border": "1px",
          "--tab-radius": "0.5rem",
        },
      },
    ],
  },
  plugins: [require("@tailwindcss/forms"), require("daisyui")],
};
