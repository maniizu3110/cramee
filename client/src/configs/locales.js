const locale = "ja";

export default {
  // current locale
  locale,

  // when translation is not available fallback to that locale
  fallbackLocale: "ja",

  // availabled locales for user selection
  availableLocales: [
    {
      code: "en",
      flag: "us",
      name: "English",
      file: "en.js"
    },
    {
      code: "ja",
      flag: "jp",
      name: "日本語",
      file: "ja.js"
    }
  ]
};
