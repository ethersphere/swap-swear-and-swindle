module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 2020,
    sourceType: 'module',
  },
  extends: [
    'plugin:@typescript-eslint/recommended',
    'plugin:prettier/recommended', // Puts Prettier last to override other settings
  ],
  rules: {
    '@typescript-eslint/no-non-null-assertion': 'off',
    // other rules...
  },
};
