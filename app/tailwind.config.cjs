/** @type {import("tailwindcss").Config} */
module.exports = {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {

    extend: {
      colors: {
        primary: {
          DEFAULT: "#12013D",
          "50": "#9A73FD",
          "100": "#8C5FFC",
          "200": "#6F37FC",
          "300": "#520FFB",
          "400": "#4104DE",
          "500": "#3603B5",
          "600": "#2A028D",
          "700": "#1E0265",
          "800": "#12013D",
          "900": "#020006"
        }
      }
    }
  },
  plugins: []
};
