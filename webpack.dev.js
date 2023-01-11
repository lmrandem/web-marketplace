import path from 'path';
import * as url from 'url';
import { merge } from 'webpack-merge';
import MiniCssExtractPlugin from 'mini-css-extract-plugin';
import CssMinimizerPlugin from 'css-minimizer-webpack-plugin';
import common from './webpack.common.js';

const [mainCommonConfig, swConfig] = common;

const mainConfig = merge(mainCommonConfig, {
  mode: 'development',
  output: {
    filename: '[name].mjs',
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: '[name].css',
    }),
    new CssMinimizerPlugin()
  ]
});

export default [mainConfig, swConfig];
