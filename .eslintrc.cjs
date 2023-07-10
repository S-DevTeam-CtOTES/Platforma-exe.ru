/* eslint-env node */

module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:@typescript-eslint/recommended-requiring-type-checking',
    'plugin:react-hooks/recommended',
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
    project: true,
    tsconfigRootDir: __dirname,
  },
  settings: {
    'import/resolver': {
      alias: {
        map: [
          ['@', '/src'],
          ['@components', '/src/Components'],
          ['@assets', '/src/assets'],
          ['@imgs', '/src/assets/imgs'],
        ],
        extensions: ['.ts', '.tsx', '.js', '.jsx', '.json'],
      },
    },
    'import/no-unresolved': [
      'error',
      {
        ignore: ['@components/*', '@assets/*', '@imgs/*', '@/*'],
      },
    ],
  },
  plugins: ['react-refresh',],
  rules: {
    'react-refresh/only-export-components': [
      'warn',
      { allowConstantExport: true },
    ],
    '@typescript-eslint/no-non-null-assertion': 'off',
  },
}
