module.exports = {
  root: true,
  env: {
    browser: true,
    "vue/setup-compiler-macros": true,
  },
  extends: [
    // "plugin:import/recommended",
    // "plugin:import/errors",
    // "plugin:import/warnings",
    // "plugin:import/typescript",
    "plugin:vue/vue3-essential",
    "plugin:vue/vue3-recommended",
    "plugin:vue/vue3-strongly-recommended",
    "eslint:recommended",
    "@vue/typescript/recommended",
    "plugin:@typescript-eslint/eslint-recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-type-checked",
    "plugin:@typescript-eslint/strict",
    "plugin:@typescript-eslint/strict-type-checked",
    "plugin:@typescript-eslint/stylistic",
    "plugin:@typescript-eslint/stylistic-type-checked",
    "plugin:unicorn/recommended",
    "prettier",
  ],
  // plugins: ["simple-import-sort"],
  parser: "vue-eslint-parser",
  parserOptions: {
    ecmaVersion: 2022,
    project: ["./dashboard/**/tsconfig.json"],
    parser: "@typescript-eslint/parser",
  },
  rules: {
    // "no-console": process.env.NODE_ENV === "production" ? ["warn", {allow: ["warn", "error"]}] : "off",
    "unicorn/prevent-abbreviations": "off",
    "unicorn/filename-case": "off",
    "@typescript-eslint/no-empty-function": ["error", { allow: ["arrowFunctions"] }],
    "unicorn/switch-case-braces": "off",
    "unicorn/no-null": "off",
    "unicorn/no-magic-array-flat-depth": "off",
    "unicorn/numeric-separators-style": "off",
    "unicorn/consistent-function-scoping": ["error", { checkArrowFunctions: false }],
    "unicorn/no-new-array": "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "off",
    "max-len": ["error", { code: 300 }],
    "object-shorthand": ["error", "always", { avoidExplicitReturnArrows: true }],
    quotes: ["error", "double", { avoidEscape: true }],
    "@typescript-eslint/no-unused-vars": "off",
    "@typescript-eslint/prefer-regexp-exec": "off",
    // "import/order": [process.env.NODE_ENV === "production" ? "error" : "warn", { alphabetize: { order: "asc" } }],
    // "import/no-unresolved": "off",
    "@typescript-eslint/restrict-template-expressions": ["error", { allowNullish: true }],
    "@typescript-eslint/no-inferrable-types": ["error", { ignoreParameters: true }],
    "@typescript-eslint/no-unsafe-enum-comparison": "off",
    "@typescript-eslint/non-nullable-type-assertion-style": "off",
    "vue/html-quotes": ["error", "double", { avoidEscape: true }],
    "vue/multi-word-component-names": [
      "error",
      {
        ignores: ["Dashboard", "Report", "Divider"],
      },
    ],
    "vue/no-setup-props-destructure": "off",
    "vue/no-deprecated-filter": "off",
    "unicorn/prefer-global-this": "off",
  },
}
