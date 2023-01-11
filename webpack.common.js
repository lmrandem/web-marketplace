import path from 'path';
import * as url from 'url';
import MiniCssExtractPlugin from 'mini-css-extract-plugin';

const mainConfig = {
  entry: {
    main: [
      './public/src/scripts/main.mjs',
      './public/src/styles/style.css',
    ],
  },
  output: {
    path: path.resolve(url.fileURLToPath(new URL('.', import.meta.url)), 'public/dist')
  },
  module: {
    rules: [
      {
        test: /\.(m|)js$/i,
        use: 'babel-loader'
      },
      {
        test: /\.css$/i,
        use: [MiniCssExtractPlugin.loader, 'css-loader'],
      }
    ]
  },
  resolve: {
    extensions: ['.js', '.mjs', '.css']
  }
};

const swConfig = {
  entry: {
    'caching-service-worker': './public/src/scripts/caching-service-worker.js'
  },
  output: {
    path: path.resolve(url.fileURLToPath(new URL('.', import.meta.url)), 'public/dist'),
    filename: '[name].js'
  },
  module: {
    rules: [
      {
        test: /\.(m|)js$/i,
        use: 'babel-loader'
      }
    ]
  },
  resolve: {
    extensions: ['.js']
  }
}

export default [mainConfig, swConfig];