import configs from "../../configs";

const { product, time, theme, currencies } = configs;

const {
  globalTheme,
  menuTheme,
  toolbarTheme,
  isToolbarDetached,
  isContentBoxed,
  isRTL,
} = theme;
const { currency, availableCurrencies } = currencies;

// state initial values
export const state = () => ({
  product,

  time,
  // currency
  currency,
  availableCurrencies,

  // themes and layout configurations
  globalTheme,
  menuTheme,
  toolbarTheme,
  isToolbarDetached,
  isContentBoxed,
  isRTL,
});

export const mutations = {
  setGlobalTheme: function (state, theme) {
    this.app.vuetify.framework.theme.dark = theme === "dark";
    state.globalTheme = theme;
  },
  setRTL: function (state, isRTL) {
    this.app.vuetify.framework.rtl = isRTL;
    state.isRTL = isRTL;
  },
  setContentBoxed: (state, isBoxed) => {
    state.isContentBoxed = isBoxed;
  },
  setMenuTheme: (state, theme) => {
    state.menuTheme = theme;
  },
  setToolbarTheme: (state, theme) => {
    state.toolbarTheme = theme;
  },
  setTimeZone: (state, zone) => {
    state.time.zone = zone;
  },
  setTimeFormat: (state, format) => {
    state.time.format = format;
  },
  setCurrency: (state, currency) => {
    state.currency = currency;
  },
  setToolbarDetached: (state, isDetached) => {
    state.isToolbarDetached = isDetached;
  },
};
