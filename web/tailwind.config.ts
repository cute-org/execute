/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,ts}'],
  theme: {
    extend: {
      colors: {
        navBar: '#271834',
        activeWindowNavBar: '#4d3168',
        borderColor: '#3C2650',
        fillingInfo: '#201B1B',
        infoBg: '#100F0F',
        error: '#cc0000',
        accepted: '#33ff33',
        descriptionDialog: '#201B1B',
      },

      spacing: {
        '10xl': '10rem',
        '11xl': '12rem',
        '12xl': '14rem',
      },
      fontFamily: {
        adlam: ['"ADLaM Display"', 'sans-serif'],
      },
    },
  },
  plugins: [],
}
