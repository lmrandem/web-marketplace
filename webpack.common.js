import path from 'path';
import * as url from 'url';
import MiniCssExtractPlugin from 'mini-css-extract-plugin';

export default {
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
      },
    ],
  },
  resolve: {
    extensions: ['.js', '.mjs', '.css']
  }
};