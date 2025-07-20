import path from 'node:path';
import { defineConfig } from '@rsbuild/core';
import { pluginReact } from '@rsbuild/plugin-react';

export default defineConfig({
  dev: {
    writeToDisk: (file: string) => file.includes('.html'),
  },
  source: {
    entry: {
      index: 'frontend/src/index.tsx',
    },
    define: {
      'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV),
      'process.env.API_URL': JSON.stringify(process.env.API_URL),
    },
  },
  output: {
    assetPrefix: process.env.FRONTEND_CDN_PATH,
    distPath: {
      root: 'frontend/dist',
    },
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'frontend/src'),
    },
  },
  plugins: [pluginReact()],
});
